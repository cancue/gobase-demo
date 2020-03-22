# [gobase](https://github.com/cancue/gobase)-demo

## 0. Setup

### Go
```bash
$ go build
$ go run app.go
```

### DB(postgresql)
#### prerequisite
```
docker
```

#### run container
```bash
$ mkdir -p ~/workspace/data/postgresql
$ export DOCKER_PGDATA=$(cd ~/workspace/data/postgresql && pwd)
$ docker run --name "dev-pg" -d -p "5432:5432" -v "$DOCKER_PGDATA:/var/lib/postgresql/data" postgres:11.6
```

#### create user and database (example)
```bash
$ docker exec -it dev-pg bash
root# psql --username=postgres
postgres=# create user local_demo_user with password 'local_demo_pw';
postgres=# create database local_demo owner local_demo_user;
postgres=# \q
root# exit
```

#### migration
```bash
$ go run db/migrate.go up
```

## 1. Usages
### New Stage
```bash
$ vi config/{stage}.yaml
$ STAGE={stage} && go run app.go
```

### New Table
```bash
$ vi db/migrations/{migration}.(sql|go) # based on [goose](https://github.com/pressly/goose)
$ go run db/goose.go up
$ vi model/{package}/dao.go
```

### New Endpoint
```bash
$ vi controller/{package}/{package}.go
$ vi router/router.go
$ vi router/router_test.go
```

## 2. Test
### TDD
```bash
$ go get github.com/dnephin/filewatcher # watch
$ go get gotest.tools/gotestsum         # format

$ filewatcher gotestsum ./...
```

### Coverage
#### summary
```bash
$ go test ./... -cover -coverprofile coverage.out; go tool cover -func coverage.out   # print console
$ go test ./... -cover -coverprofile=coverage.out && go tool cover -html=coverage.out # open browser
```

#### detail
```bash
$ go get github.com/axw/gocov/gocov

$ gocov test ./... | gocov-html > coverage_report.html
