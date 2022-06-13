package query

import (
	"context"
	"time"

	"github.com/Ranxy/gormfields/generate/example"
	"github.com/Ranxy/gormfields/generate/example/example_fields"
	"github.com/Ranxy/gormfields/query"
	"gorm.io/gorm"
)

type service struct {
	userOperator query.Operator[example.User]
	db           *gorm.DB
}

func (s *service) QueryUser(ctx context.Context) ([]*example.User, error) {
	db := s.db.WithContext(ctx)

	// select * from users where phone = 13412 or user_name = 'foo' limit 10 offset 20
	return s.userOperator.Find(ctx,
		db,
		example_fields.UserPhone(13412),
		example_fields.UserUserName("foo", query.Or()),
		query.Limit[example.User](10),
		query.Offset[example.User](20),
	)
}
func (s *service) UpdateUser(ctx context.Context) error {
	db := s.db.WithContext(ctx)

	// update users set updated_at = 'now-time',user_name='bar' where phone = 13412
	return s.userOperator.Update(ctx,
		db,
		[]query.Field[example.User]{example_fields.UserPhone(13412)},
		example_fields.UserUpdatedAt(time.Now()),
		example_fields.UserUserName("bar"),
	)
}
