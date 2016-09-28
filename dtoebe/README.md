# Phoenix Golang Meetup Image API Challenge

This is my submission to the Phoenix GoLang challenge #2

> Problem 2:

> Create an API with an endpoint that returns a PNG image containing the numeric representation of the number of times that endpoint has been hit.

I basically built an API that can easily be configured for more endpoints than what I am displaying

I was going to go simple with this but ended up starting a proof of consept for an easily extendable api platform.  While chances are I will not continue to update this repo I will start working on a hot loading API system in the near future. 

For the purposes of this challenge, don't make too many changes to the server-config.json, accept the absolute path to the font-path, and image-dir.

### To run

Clone the repo: `git clone git@github.com:dtoebe/gophx-img-api.git`

#### Makefile commands

Get, Build and run 


Get all 3rd party dependancies
``` Bash
make get
```

Run `go run`
``` Bash
make ARGS="/path/to/server-config.json" run-dev
```

Compile and run
``` Bash
make ARGS="/path/to/server-config.json" run
```
Build
``` Bash
make build
```

Clean up the compiled Binary and the images folder
``` Bash
make clean
```



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

- Finish tests (figure out best way to test with httprouter)
    - I have found a way for basic testing of httprouter I don't have time to add the appropriate tests to my dynamic routing scheme.
