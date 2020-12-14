.PHONY: all
# TODO: add env to bin name
BINARY=todos

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

install:
	@echo "Installing air" 
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
	make dependencies
	make lint-prepare

integration-test: 
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up --abort-on-container-exit
	docker-compose -f docker-compose.yml -f docker-compose.test.yml down

run-migrations: 
	# migrations script
run-integration-test: 
	go test --tags="integration" -v ./...

unit-test:
	go test --tags="unit" -v ./...

test: 
	go test --tags="unit integration" -cover -covermode=atomic ./...

build: dependencies build-todos-cmd

# TODO: use tags
build-todos-cmd:
	go build -tags release -o ./build/${BINARY} ./cmd/todos/main.go

dependencies:
	go mod download

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

# send binary name to dockerfile?
docker:
	docker build -t todos .

start:
	docker-compose up -d

stop:
	docker-compose down

serve:
	air -c .air.conf

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

# read more about linting in: https://github.com/golangci/golangci-lint#config-file - maybe use config file?
lint:
	./bin/golangci-lint run
