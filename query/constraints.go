package query

import "gorm.io/gorm"

type Field[V TableModel] interface {
	Do(db *gorm.DB) *gorm.DB
	DoUpdate(UpdateReq)
	Table() V
}
