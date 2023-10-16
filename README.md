Go Echo Setup
=========================================


### How To Run Application
```shell
docker-compose up -d --build
go run main.go runserver
```
or
```shell
# support live reload
go install github.com/cosmtrek/air@latest
# then run
air runserver
```

### How to Run Migration
```shell
go run main.go migrate -e up
```
another script:
```shell
go run main.go makemigrations -f "create_new_migration_file"
go run main.go migrate -e up 0001
go run main.go migrate -e down
go run main.go migrate -e down 0001
go run main.go migrate -e undo
```

### How to Generate Swagger
```shell
swag init --parseDependency --parseInternal
```