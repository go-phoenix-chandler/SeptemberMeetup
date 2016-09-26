# Phoenix Golang Meetup Image API Challenge

This is my submission to the Phoenix GoLang challenge #2

> Problem 2:

> Create an API with an endpoint that returns a PNG image containing the numeric representation of the number of times that endpoint has been hit.

I basically built an API that can easily be configured for more endpoints than what I am displaying

### Endpoints

Return Binary Image URL
```
http://<host>:<port>/api/v1/bin
```

Return DataURI string
```
http://<host>:<port>/api/v1/uri
```

## TODO:

- Create MakeFile for pkgs
- Test `go get`
- Have the root parse and show this README
- Comment code
- Finish tests (figure out best way to test with httprouter)
- Test EVERYTHING by actually running it
