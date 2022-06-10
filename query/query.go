package query

import (
	"gorm.io/gorm"
)

type GormQueryReq interface {
	Do(db *gorm.DB) *gorm.DB
}

type OrderBy string

func (i OrderBy) Do(db *gorm.DB) *gorm.DB {
	return db.Order(i)
}

type Limit int

func (i Limit) Do(db *gorm.DB) *gorm.DB {
	return db.Limit(int(i))
}

type Offset int

func (i Offset) Do(db *gorm.DB) *gorm.DB {
	return db.Offset(int(i))
}

func CustomSql(or bool, sql string, args ...any) GormQueryReq {
	return &customSql{
		Sql:  sql,
		Args: args,
		or:   orCond(or),
	}
}

type customSql struct {
	Sql  string
	or   orCond
	Args []any
}

func (c *customSql) Do(db *gorm.DB) *gorm.DB {
	return c.or.Do(db)(c.Sql, c.Args...)
}

type orCond bool

func (o orCond) Do(db *gorm.DB) func(query interface{}, args ...interface{}) (tx *gorm.DB) {
	if o {
		return db.Or
	} else {
		return db.Where
	}
}

type zeroCond bool

func (o zeroCond) CheckZero() bool {
	return bool(o)
}

type Opt struct {
	Or        orCond
	CheckZero zeroCond
}

type WithOpt func(o *Opt)

func Or() WithOpt {
	return func(o *Opt) {
		o.Or = true
	}
}

// CheckZero if add this opt, cond will check the zero value ,it means `where id = 0` will not exec and it will be just return
func CheckZero() WithOpt {
	return func(o *Opt) {
		o.CheckZero = true
	}
}
