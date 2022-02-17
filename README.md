# GoPI
A visual client to easily create RESTful APIs and execute CRUD opperations

## TODO:
1. In order to allow different behavior to be dynamically bound to endpoints
we need to build a set of functions which handles different requests and then
passes the writer and request to another function which has context on which
endpoint we are handling. These functions will then return the action or something
along those lines
2. UPDATE: we can read the endpoint name from the url path r.URL.Path
