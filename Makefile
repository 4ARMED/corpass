EXECUTABLE := "corpass"
GITVERSION := $(shell git describe --dirty --always --tags --long)

build:
	go build -ldflags "-X main.version=${GITVERSION}" -o ${EXECUTABLE} -v