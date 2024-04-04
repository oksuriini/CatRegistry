# Cat Registry App

Simple API to register cats and fetch them. This API supports only MongoDB type database.

Available on Dockerhub from:

- image: oksuriini/catregistry
- tags: `main` or `nightly`

Uses envars:

`MONGODB_URI`, uri for the MongoDB, defaults to "127.0.0.1"

`MONGODB_PORT`, port that MongoDB exposes, defaults to "27017"

`HOSTNAME`, hostname for the API, defaults to "127.0.0.1"

`HOSTPORT`, port that API exposes, defaults to "8080"

`USERNAME`, username to authenticate for database, default not provided

`PASSWORD`, password for the authenticating username, default not provided

## Endpoints:

`/getcats`: Gets all the cats in database. GET Method only, returns JSON.

`/filtercats`: Gets cats by filtering them. GET Method only, returns JSON. Filtering is done by passing JSON information to it.

`/insertcat`: Insert one cat. POST Method only, returns JSON. Inserting is done by passing JSON information to it.

## Endpoint calls:

Endpoints are called with Get or Post method with JSON body.

Only endpoints `/filtercats` and `/insertcat` cares about JSON details, as `/getcats` it ignores the body and only returns information.

### JSON field has following properties:

"name", name of the cat, string

"breed", breed of the cat, string

"fur", fur color, fuzziness etc, string

"lives", the amount of lives the cat has, int

"age", age of the cat, int

"favorite_foods", favorite foods of the cat, []string

Filtering only uses fields "name", "breed", "lives" and "age", whereas when inserting a cat to database, all of these properties should be used.
