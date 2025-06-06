include local.env
export $(shell sed 's/=.*//' local.env)

run: run_broker
	
run_broker: 
	go run .

build:
	go get && go build -ldflags "-s -w" -trimpath -buildvcs=false -o bin/opentalaria

run_test:
	go run ./_examples/client/confluent/main.go

race: 
	go run -race .

test:
	go test -v ./...

cover:
	go test -coverprofile=coverage.out ./.../... ; go tool cover -html=coverage.out
