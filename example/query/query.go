package query

import (
	"context"

	"github.com/Ranxy/formfieldexample/models"
	"github.com/Ranxy/formfieldexample/models/models_fields"
	"github.com/Ranxy/gormfields/query"
	"gorm.io/gorm"
)

type service struct {
	userOperator query.Operator[models.User]
	db           *gorm.DB
}

func (s *service) QueryUser(ctx context.Context) ([]*models.User, error) {
	db := s.db.WithContext(ctx)

	// select * from users where phone = 13412 or user_name = 'foo' limit 10 offset 20
	return s.userOperator.Find(ctx,
		db,
		models_fields.UserPhone(13412),
		models_fields.UserUserName("foo", query.Or()),
		query.Limit(10),
		query.Offset(20),
	)
}
