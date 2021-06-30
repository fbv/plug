projects=\
	plug

plugins=\
	en\
	ru

all: build test

clean:
	@rm -rf bin/*

build: $(projects) $(plugins)

$(projects):
	go build -o bin/$@ cmd/$@/main.go

$(plugins):
	go build -buildmode=plugin -o lib/$@.so plugins/$@/main.go

test:
	go test ./...

.PHONY: all clean build test $(projects) $(plugins)
