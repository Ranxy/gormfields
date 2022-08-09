package query

import (
	"context"

	"gorm.io/gorm"
)

func NewOperator[V TableModel]() Operator[V] {
	return Operator[V]{}
}

type Operator[V TableModel] struct {
}

func (o Operator[V]) Insert(ctx context.Context, db *gorm.DB, val *V) error {
	return db.Save(val).Error
}

func (o Operator[V]) Get(ctx context.Context, db *gorm.DB, finds ...Field[V]) (*V, error) {
	for _, f := range finds {
		db = f.Do(db)
	}

	res := new(V)

	err := db.First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (o Operator[V]) Find(ctx context.Context, db *gorm.DB, finds ...Field[V]) ([]*V, error) {
	for _, f := range finds {
		db = f.Do(db)
	}

	res := make([]*V, 0)

	err := db.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (o Operator[V]) FindAndCount(ctx context.Context, db *gorm.DB, finds ...Field[V]) ([]*V, int64, error) {
	for _, f := range finds {
		db = f.Do(db)
	}

	res := make([]*V, 0)

	err := db.Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	var count int64 = 0
	err = db.Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

func (d Operator[V]) Delete(ctx context.Context, db *gorm.DB, finds ...Field[V]) error {
	for _, f := range finds {
		db = f.Do(db)
	}
	return db.Delete(new(V)).Error
}

func (d Operator[V]) Update(ctx context.Context, db *gorm.DB, finds []Field[V], updates ...Field[V]) error {
	updateReq := UpdateReq{}
	for _, update := range updates {
		update.DoUpdate(updateReq)
	}

	for _, f := range finds {
		db = f.Do(db)
	}
	return db.Updates(updateReq).Error
}
func (o Operator[V]) Count(ctx context.Context, db *gorm.DB, finds ...Field[V]) (int64, error) {
	for _, f := range finds {
		db = f.Do(db)
	}
	var cnt int64
	err := db.Model(new(V)).Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
func (d Operator[V]) Limit(l int) Field[V] {
	return Limit[V](l)
}

func (d Operator[V]) Offset(o int) Field[V] {
	return Offset[V](o)
}

func (d Operator[V]) OrderBy(ob string) Field[V] {
	return OrderBy[V](ob)
}

func (d Operator[V]) Projection(list ...string) Field[V] {
	return Projection[V](list...)
}

func (d Operator[V]) CustomQuery(or bool, sql string, args ...any) Field[V] {
	return CustomQuery[V](or, sql, args...)
}
