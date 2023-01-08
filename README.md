## Golang: new SQLite driver + SQLC compiler

In this example i didn't use ORMs, Instead i used a compiled go files from sqlc project: https://github.com/kyleconroy/sqlc

I also used this new SQLite driver with No CGO from https://gitlab.com/cznic/sqlite

In the example in `app.go` , I added a new function by `sqlite driver` called `new_name()`, And i called it from SQL query directly! *it's just return "zaki!"*

## Installation

```bash
go mod tidy
```

You can add your sql schema into `schema.sql` file, Then add your queries into `query.sql`,

To generate go files for your *schema/query*, run:
```
sqlc generate
```

Then run the app:
```
go run .
```

Follow me on [twitter](https://twitter.com/zaki_chahboun) 
------
2023-01-08 03:37
