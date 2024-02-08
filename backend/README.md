# MIGRATION

Command to create a new migration:

```console
migrate create -seq -ext=.sql -dir=./pkg/db/migrations create_ratings_table
```

If migrate-cli not installed:

check the installation doc [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

Explanation

    -seq: to use number for ordering 
    -ext= specify the migration file extension
    -dir= directory to the migrations directory

Command line to migrate:

```console
go run . <migration_options>
```

options:
    - up
    - down
    - upall
    - downall
    - to (on hold)

Note: All Scheme related to this project are in migrations files
