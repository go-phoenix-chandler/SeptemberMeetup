CC=go
GET=get -u -v
RUN=run
BUILD=build
BUILDFLAGS=-o
IN=main.go
OUT=img-api

all: get run

get: 
	$(CC) $(GET) github.com/julienschmidt/httprouter
	$(CC) $(GET) github.com/comail/colog
	$(CC) $(GET) github.com/golang/freetype
	$(CC) $(GET) github.com/golang/freetype/truetype
	$(CC) $(GET) golang.org/x/image/font

run-dev:
	$(CC) $(RUN) $(IN) ${ARGS}

run: build
	./$(OUT) ${ARGS}

build:
	$(CC) $(BUILD) $(BUILDFLAGS) $(OUT) $(IN)

clean:
	rm -rf $(OUT) images
