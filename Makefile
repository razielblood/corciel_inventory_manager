build:
	go build -o bin/corciel_api

run: build
	./bin/corciel_api

test:
	go test -v ./... -count=1
