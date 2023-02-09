NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
BINARY_SRC=./cmd/server.go
BINARY_DST=./bin/server

build:
	@printf "$(OK_COLOR)==> Building binary$(NO_COLOR)\n"
	go build -o ${BINARY_DST} ${BINARY_SRC}
	@printf "$(OK_COLOR)==> Succesfully built artifacts to ${BINARY_DST} $(NO_COLOR)\n"

run:
	${BINARY_DST}

test:
	@printf "$(OK_COLOR)==> Running tests$(NO_COLOR)\n"
	go test ./... -cover -coverprofile=cover.out
