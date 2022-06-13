package query

import (
	"gorm.io/gorm"
)

func OrderBy[V TableModel](val string) orderBy[V] {
	return orderBy[V]{
		orderBy: val,
	}
}

type orderBy[V TableModel] struct {
	orderBy string
}

func (i orderBy[V]) Do(db *gorm.DB) *gorm.DB {
	return db.Order(i.orderBy)
}
func (i orderBy[V]) DoUpdate(UpdateReq) {
	panic("orderby does not allow update")
}
func (i orderBy[V]) Table() V {
	var res V
	return res
}

func Limit[V TableModel](val int) limit[V] {
	return limit[V]{
		limit: val,
	}
}

type limit[V TableModel] struct {
	limit int
}

func (i limit[V]) Do(db *gorm.DB) *gorm.DB {
	return db.Limit(i.limit)
}
func (i limit[V]) DoUpdate(UpdateReq) {
	panic("limit does not allow update")
}
func (i limit[V]) Table() V {
	var res V
	return res
}

func Offset[V TableModel](val int) offset[V] {
	return offset[V]{
		offset: val,
	}
}

type offset[V TableModel] struct {
	offset int
}

func (i offset[V]) Do(db *gorm.DB) *gorm.DB {
	return db.Offset(i.offset)
}
func (i offset[V]) DoUpdate(UpdateReq) {
	panic("offset does not allow update")
}
func (i offset[V]) Table() V {
	var res V
	return res
}

func CustomQuery[V TableModel](or bool, sql string, args ...any) Field[V] {
	return &customSql[V]{
		Sql:  sql,
		Args: args,
		or:   orCond(or),
	}
}

type customSql[V TableModel] struct {
	Sql  string
	or   orCond
	Args []any
}

func (c *customSql[V]) Do(db *gorm.DB) *gorm.DB {
	return c.or.Do(db)(c.Sql, c.Args...)
}

func (i customSql[V]) DoUpdate(UpdateReq) {
	panic("customQuery does not allow update")
}
func (i customSql[V]) Table() V {
	var res V
	return res
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
