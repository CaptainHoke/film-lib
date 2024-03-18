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
- POST sign-in/auth?username=...password=...

Role >= User:
- [x] GET /actors
- [x] GET /films?sort=title/rating/release-date
- [ ] GET /search?query=search-query

Admin-only:
- [x] POST /actors?name=...sex=...birthdate=...
- [x] PUT /actors/{actor-id}?name=...sex=...birthdate=...
- [x] DELETE /actors/{actor-id}
- [x] POST /films?title=...desc=...release-date=...rating=...actors=[id1, id2, ...]
- [x] PUT /films?title=...desc=...release-date=...rating=...actors=[id1, id2, ...]
- [x] DELETE /films/{film-id}
