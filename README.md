# film-lib

## TODO
- [ ] Create the app lul
- [ ] Write launch instructions
- [ ] Roll out std-http-based server

## Milestones
- [x] Goa API for Actors + JWT Auth
- [ ] API for Films
- [ ] API for Querying

### API Reference:

Auth:
- POST admins/sign-in
- POST users/sign-in

Role >= User:
- [x] GET /actors
- [ ] GET /films?sort=title/rating/release-date
- [ ] GET /search?query=search-query

Admin-only:
- [x] POST /actors?name=...sex=...birthdate=...
- [x] PUT /actors/{actor-id}?name=...sex=...birthdate=...
- [x] DELETE /actors/{actor-id}
- POST /films?title=...desc=...release-date=...rating=...actors=[id1, id2, ...]
- PUT /films?title=...desc=...release-date=...rating=...actors=[id1, id2, ...]
- DELETE /films/{film-id}
