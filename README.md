# film-lib

- Using Goa to "generate docs faster" was a terrible mistake
- OpenAPI 3 docs are at gen/http/openapi3.json
- It barely works, it's ugly, but it was fun

## TODO
- [x] Goa API for Actors + JWT Auth
- [x] API for Films
- [x] API for Querying
- [x] Auth Impl
- [x] Build commands adjustments
- [ ] Logger setup
- [x] DB setup
  - pgx
- [x] Docker setup
- [ ] Get rid of non-std-http frameworks; might break everything ye whatever
- [ ] Make project structure less ugly
- [x] Fix linter errors
- [ ] Make sure that the project is "eventually test-covered" (idk when maybe a day or two or three after deadline)
- [ ] getAllFilms => getFilms + limit
- [ ] Make auth less retarded
  - Store users in-memory for now
  - Users table later
- [x] Fix docs not working as intended, probably will need to regenerate the whole thing
- [ ] Error handling sucks
- [ ] No tests
