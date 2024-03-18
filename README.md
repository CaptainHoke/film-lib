# film-lib

## TODO
- [ ] Create the app lul
- [ ] Write launch instructions
- [ ] Roll out std-http-based server

## Milestones
- [x] Goa API for Actors + JWT Auth
- [x] API for Films
- [x] API for Querying

### API Reference:

Auth:
- [x] POST sign-in/auth?username=...password=...

Role >= User:
- [x] GET /actors
- [x] GET /films?sort=title/rating/release-date
- [x] GET /search?query=search-query

Admin-only:
- [x] POST /actors?name=...sex=...birthdate=...
- [x] PUT /actors/{actor-id}?name=...sex=...birthdate=...
- [x] DELETE /actors/{actor-id}
- [x] POST /films?title=...desc=...release-date=...rating=...actors=[id1, id2, ...]
- [x] PUT /films?title=...desc=...release-date=...rating=...actors=[id1, id2, ...]
- [x] DELETE /films/{film-id}
