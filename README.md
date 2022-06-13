# gormfields
Simple and fast generation of strongly typed gorm parameters.


### Usage

#### Install
Run `go install github.com/Ranxy/gormfields/gormfields@latest`

#### Generate
First, add `// gormfields:query` to the structure to be generated, or use `-all` in the next command to provide the structure to be generated

Then, use the command `gormfields [--all] [path-to-model]`,as in the example, using `gormfields -all . /example/models`.

#### example usage
```go
var userOperator query.Operator[models.User]

userOperator.Find(ctx,
		db,
		models_fields.UserPhone(13412),
		models_fields.UserUserName("foo", query.Or()),
		query.Limit[models.User](10),
		query.Offset[models.User](20),
	)
```


#### Examples of errors
If we use a field from another table, we will get a compile error.
```go
userOperator.Find(ctx,
		db,
		models_fields.UserPhone(13412),
		models_fields.RoleRoleInfo("foo", query.Or()),
		query.Limit[models.User](10),
		query.Offset[models.User](20),
	)
```
We will got a compile error like this 
```
cannot use models_fields.RoleRoleInfo("foo", query.Or()) (value of type *models_fields.hRoleRoleInfo) as "github.com/Ranxy/gormfields/query".Field[models.User] value in argument to s.userOperator.Find: *models_fields.hRoleRoleInfo does not implement "github.com/Ranxy/gormfields/query".Field[models.User] (wrong type for method Table)
		have Table() models.Role
		want Table() models.User
```
Detailed examples can be found in the examples folder.

#### Problem

```
query.Limit[models.User](10),
```
If we don't add `[models.User]` manually, go can't automatically infer the type here.