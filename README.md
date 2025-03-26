# in.orbit API in Go

Same Api that you can find in [in.orbit](https://www.github.com/nathancamolez-dev/in.orbit-back-end-/) but writed in Go.

## Install

Very important to have Go, sqlc , tern, docker installed on your machine

Before everything use `docker compose -d up` for creating the database

```bash
go get github.com/nathancamolez-dev/in.orbit-go
go get -u ./...
```

After that you have to use the wrappe for tern because of the env file.

```bash
go run cmd/terdotenv/main.go
```

Now you can generate the queries by using the sqlc command on the location of the sqlc.yml

```bash
sqlc generate
```
