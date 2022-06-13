// Code generated by gorm_column_gen. DO NOT EDIT.
package example_fields
	
import (
	. "github.com/Ranxy/gormfields/generate/example"
	"github.com/Ranxy/gormfields/query"
	"gorm.io/gorm"
	"time"
	
) 



func UserCreatedAt(CreatedAt time.Time, opts ...query.WithOpt) *hUserCreatedAt {
	res := hUserCreatedAt{
		CreatedAt: CreatedAt,
	}
	for _, opt := range opts {
		opt(&res.opt)
	}
	return &res

}

type hUserCreatedAt struct {
	CreatedAt  time.Time
	opt query.Opt
}

func (i *hUserCreatedAt) Do(db *gorm.DB) *gorm.DB {
	
	var zero time.Time
	if i.opt.CheckZero.CheckZero() && i.CreatedAt == zero {
		return db
	}
	
	return i.opt.Or.Do(db)("created_at = ?", i.CreatedAt)
}

func (i *hUserCreatedAt) DoUpdate(req query.UpdateReq){
	req["created_at"] = i.CreatedAt
}

func (i *hUserCreatedAt) Table()User{
	return User{}
}



func UserUpdatedAt(UpdatedAt time.Time, opts ...query.WithOpt) *hUserUpdatedAt {
	res := hUserUpdatedAt{
		UpdatedAt: UpdatedAt,
	}
	for _, opt := range opts {
		opt(&res.opt)
	}
	return &res

}

type hUserUpdatedAt struct {
	UpdatedAt  time.Time
	opt query.Opt
}

func (i *hUserUpdatedAt) Do(db *gorm.DB) *gorm.DB {
	
	var zero time.Time
	if i.opt.CheckZero.CheckZero() && i.UpdatedAt == zero {
		return db
	}
	
	return i.opt.Or.Do(db)("updated_at = ?", i.UpdatedAt)
}

func (i *hUserUpdatedAt) DoUpdate(req query.UpdateReq){
	req["updated_at"] = i.UpdatedAt
}

func (i *hUserUpdatedAt) Table()User{
	return User{}
}



func UserDeletedAt(DeletedAt gorm.DeletedAt, opts ...query.WithOpt) *hUserDeletedAt {
	res := hUserDeletedAt{
		DeletedAt: DeletedAt,
	}
	for _, opt := range opts {
		opt(&res.opt)
	}
	return &res

}

type hUserDeletedAt struct {
	DeletedAt  gorm.DeletedAt
	opt query.Opt
}

func (i *hUserDeletedAt) Do(db *gorm.DB) *gorm.DB {
	
	var zero gorm.DeletedAt
	if i.opt.CheckZero.CheckZero() && i.DeletedAt == zero {
		return db
	}
	
	return i.opt.Or.Do(db)("deleted_at = ?", i.DeletedAt)
}

func (i *hUserDeletedAt) DoUpdate(req query.UpdateReq){
	req["deleted_at"] = i.DeletedAt
}

func (i *hUserDeletedAt) Table()User{
	return User{}
}



func UserUserName(UserName string, opts ...query.WithOpt) *hUserUserName {
	res := hUserUserName{
		UserName: UserName,
	}
	for _, opt := range opts {
		opt(&res.opt)
	}
	return &res

}

type hUserUserName struct {
	UserName  string
	opt query.Opt
}

func (i *hUserUserName) Do(db *gorm.DB) *gorm.DB {
	
	var zero string
	if i.opt.CheckZero.CheckZero() && i.UserName == zero {
		return db
	}
	
	return i.opt.Or.Do(db)("user_name = ?", i.UserName)
}

func (i *hUserUserName) DoUpdate(req query.UpdateReq){
	req["user_name"] = i.UserName
}

func (i *hUserUserName) Table()User{
	return User{}
}



func UserUserDisplay(UserDisplay *string, opts ...query.WithOpt) *hUserUserDisplay {
	res := hUserUserDisplay{
		UserDisplay: UserDisplay,
	}
	for _, opt := range opts {
		opt(&res.opt)
	}
	return &res

}

type hUserUserDisplay struct {
	UserDisplay  *string
	opt query.Opt
}

func (i *hUserUserDisplay) Do(db *gorm.DB) *gorm.DB {
	
	var zero *string
	if i.opt.CheckZero.CheckZero() && i.UserDisplay == zero {
		return db
	}
	
	return i.opt.Or.Do(db)("user_display = ?", i.UserDisplay)
}

func (i *hUserUserDisplay) DoUpdate(req query.UpdateReq){
	req["user_display"] = i.UserDisplay
}

func (i *hUserUserDisplay) Table()User{
	return User{}
}



func UserPhone(Phone int64, opts ...query.WithOpt) *hUserPhone {
	res := hUserPhone{
		Phone: Phone,
	}
	for _, opt := range opts {
		opt(&res.opt)
	}
	return &res

}

type hUserPhone struct {
	Phone  int64
	opt query.Opt
}

func (i *hUserPhone) Do(db *gorm.DB) *gorm.DB {
	
	var zero int64
	if i.opt.CheckZero.CheckZero() && i.Phone == zero {
		return db
	}
	
	return i.opt.Or.Do(db)("phone = ?", i.Phone)
}

func (i *hUserPhone) DoUpdate(req query.UpdateReq){
	req["phone"] = i.Phone
}

func (i *hUserPhone) Table()User{
	return User{}
}



func UserStatus(Status uint, opts ...query.WithOpt) *hUserStatus {
	res := hUserStatus{
		Status: Status,
	}
	for _, opt := range opts {
		opt(&res.opt)
	}
	return &res

}

type hUserStatus struct {
	Status  uint
	opt query.Opt
}

func (i *hUserStatus) Do(db *gorm.DB) *gorm.DB {
	
	var zero uint
	if i.opt.CheckZero.CheckZero() && i.Status == zero {
		return db
	}
	
	return i.opt.Or.Do(db)("status = ?", i.Status)
}

func (i *hUserStatus) DoUpdate(req query.UpdateReq){
	req["status"] = i.Status
}

func (i *hUserStatus) Table()User{
	return User{}
}



func UserID(ID uint, opts ...query.WithOpt) *hUserID {
	res := hUserID{
		ID: ID,
	}
	for _, opt := range opts {
		opt(&res.opt)
	}
	return &res

}

type hUserID struct {
	ID  uint
	opt query.Opt
}

func (i *hUserID) Do(db *gorm.DB) *gorm.DB {
	
	var zero uint
	if i.opt.CheckZero.CheckZero() && i.ID == zero {
		return db
	}
	
	return i.opt.Or.Do(db)("id = ?", i.ID)
}

func (i *hUserID) DoUpdate(req query.UpdateReq){
	req["id"] = i.ID
}

func (i *hUserID) Table()User{
	return User{}
}




	