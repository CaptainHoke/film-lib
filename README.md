# film-lib

## TODO
- [ ] Create the app lul
- [ ] Write launch instructions
- [ ] Setup golangci as a separate gh action

## Milestones
- [ ] Goa API for Actors + JWT Auth
  - GET /actors
  - POST /actors/{actor-id}?name=...sex=...birthdate=...
  - PUT /actors/{actor-id}?name=...sex=...birthdate=...
- [ ] API for Films
- [ ] API for Querying

### API Reference:

Auth:
- POST admins/sign-in
- POST users/sign-in

Role >= User:
- GET /actors
- GET /films?sort=title/rating/release-date
- GET /search?query=search-query

Admin-only:
- POST /actors?name=...sex=...birthdate=...
- PUT /actors/{actor-id}?name=...sex=...birthdate=...
- DELETE /actors/{actor-id}
- POST /films?title=...desc=...release-date=...rating=...actors=[id1, id2, ...]
- PUT /films?title=...desc=...release-date=...rating=...actors=[id1, id2, ...]
- DELETE /films/{film-id}



