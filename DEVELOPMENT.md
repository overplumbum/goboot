### Dev Env bootstrap ###

```bash
# cd project root
brew install golang dep
export GOPATH=$PWD
cd src/api
dep ensure # can take several minutes
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
dep ensure -add github.com/bugsnag/bugsnag-go/gin # with package name
```

### Docker Testing ###

```bash
docker-compose up --build
open http://localhost:8001/robots.txt
```
