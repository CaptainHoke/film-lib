FROM golang:1.22.1-alpine AS base
WORKDIR /src

COPY go.* .
RUN go mod download
COPY . .

FROM base AS build
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /out/app ./cmd/film_lib && go build -o /out/cli ./cmd/film_lib-cli

#FROM base AS test
#RUN --mount=target=. \
#    --mount=type=cache,target=/root/.cache/go-build \
#    go test -v .

FROM golangci/golangci-lint:v1.56.2-alpine AS lint-base

FROM base AS lint
RUN --mount=target=. \
    --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
    golangci-lint run --timeout 1m0s ./...

FROM scratch AS bin
COPY --from=build /out/app /

FROM gcr.io/distroless/base-debian12:latest-amd64
COPY ./db/migrations ./db/migrations
COPY --from=build /out/app /
CMD ["./app", "--debug=true", "--host=docker"]
