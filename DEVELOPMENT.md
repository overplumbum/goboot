### Dev Env bootstrap ###

```bash
# cd project root
brew install golang govendor
export GOPATH=$PWD
cd src/api
govendor sync # can take several minutes
```

### Run / Build ###

```bash
make run
```

### New dependency ###

```bash
# cd project root
export GOPATH=$PWD
cd src/api
govendor fetch -v github.com/bugsnag/bugsnag-go/gin # with package name
```

### Docker Testing ###

```bash
docker-compose up --build
open http://localhost:8001/robots.txt
```
