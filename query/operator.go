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

func (d Operator[V]) Limit(l int) Field[V] {
	return Limit[V](l)
}

func (d Operator[V]) Offset(o int) Field[V] {
	return Offset[V](o)
}

func (d Operator[V]) OrderBy(ob string) Field[V] {
	return OrderBy[V](ob)
}
