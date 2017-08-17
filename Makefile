build:
	GOPATH=`pwd` go build -o /tmp/api-cli src/api/cli/cli.go

run:
	TZ=UTC GOPATH=`pwd` go run src/api/cli/cli.go web

# creates missing tables, don't modifies any existing tables
migrate:
	GOPATH=`pwd` go run src/api/cli/cli.go migrate

schema:
	GOPATH=`pwd` go run src/api/cli/cli.go schema

psql: bin/cli
	./bin/cli psql

bin/cli: src/api/config/config.go
	GOPATH=`pwd` go build -o bin/cli src/api/cli/cli.go

fmt:
	find src/api -name '*.go' | grep -vE '^/*src/api/vendor' | xargs -L1 -P4 go fmt
