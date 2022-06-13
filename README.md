# gormfields
Simple and fast generation of strongly typed gorm parameters.


### Usage

#### Install
Run `go install github.com/Ranxy/gormfields/gormfields@latest`

#### Generate
First, add "// gormfields:query" to the structure to be generated, or use `-all` in the next command to provide the structure to be generated

Then, use the command `gormfields [--all] [path-to-model]`,as in the example, using `gormfields -all . /example/models`.

#### example usage
```go
s.userOperator.Find(ctx,
		db,
		models_fields.UserPhone(13412),
		models_fields.UserUserName("foo", query.Or()),
		query.Limit(10),
		query.Offset(20),
	)
```
Detailed examples can be found in the examples folder.
