// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"martpay/app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newUser(db *gorm.DB, opts ...gen.DOOption) user {
	_user := user{}

	_user.userDo.UseDB(db, opts...)
	_user.userDo.UseModel(&models.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewAsterisk(tableName)
	_user.ID = field.NewUint(tableName, "id")
	_user.CreatedAt = field.NewTime(tableName, "created_at")
	_user.UpdatedAt = field.NewTime(tableName, "updated_at")
	_user.LevelId = field.NewUint(tableName, "level_id")
	_user.SuperAgentId = field.NewUint(tableName, "super_agent_id")
	_user.FirstName = field.NewString(tableName, "first_name")
	_user.OtherNames = field.NewString(tableName, "other_names")
	_user.BusinessName = field.NewString(tableName, "business_name")
	_user.Email = field.NewString(tableName, "email")
	_user.Phone = field.NewString(tableName, "phone")
	_user.Gender = field.NewString(tableName, "gender")
	_user.Dob = field.NewString(tableName, "dob")
	_user.State = field.NewString(tableName, "state")
	_user.Country = field.NewString(tableName, "country")
	_user.Address = field.NewString(tableName, "address")
	_user.Status = field.NewString(tableName, "status")
	_user.Avatar = field.NewString(tableName, "avatar")
	_user.Bvn = field.NewString(tableName, "bvn")
	_user.Nin = field.NewString(tableName, "nin")
	_user.ReferralCode = field.NewString(tableName, "referral_code")
	_user.EmailVerifiedAt = field.NewTime(tableName, "email_verified_at")
	_user.Password = field.NewString(tableName, "password")
	_user.DeletedAt = field.NewField(tableName, "deleted_at")

	_user.fillFieldMap()

	return _user
}

type user struct {
	userDo

	ALL             field.Asterisk
	ID              field.Uint
	CreatedAt       field.Time
	UpdatedAt       field.Time
	LevelId         field.Uint
	SuperAgentId    field.Uint
	FirstName       field.String
	OtherNames      field.String
	BusinessName    field.String
	Email           field.String
	Phone           field.String
	Gender          field.String
	Dob             field.String
	State           field.String
	Country         field.String
	Address         field.String
	Status          field.String
	Avatar          field.String
	Bvn             field.String
	Nin             field.String
	ReferralCode    field.String
	EmailVerifiedAt field.Time
	Password        field.String
	DeletedAt       field.Field

	fieldMap map[string]field.Expr
}

func (u user) Table(newTableName string) *user {
	u.userDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *user) updateTableName(table string) *user {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.LevelId = field.NewUint(table, "level_id")
	u.SuperAgentId = field.NewUint(table, "super_agent_id")
	u.FirstName = field.NewString(table, "first_name")
	u.OtherNames = field.NewString(table, "other_names")
	u.BusinessName = field.NewString(table, "business_name")
	u.Email = field.NewString(table, "email")
	u.Phone = field.NewString(table, "phone")
	u.Gender = field.NewString(table, "gender")
	u.Dob = field.NewString(table, "dob")
	u.State = field.NewString(table, "state")
	u.Country = field.NewString(table, "country")
	u.Address = field.NewString(table, "address")
	u.Status = field.NewString(table, "status")
	u.Avatar = field.NewString(table, "avatar")
	u.Bvn = field.NewString(table, "bvn")
	u.Nin = field.NewString(table, "nin")
	u.ReferralCode = field.NewString(table, "referral_code")
	u.EmailVerifiedAt = field.NewTime(table, "email_verified_at")
	u.Password = field.NewString(table, "password")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 23)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["level_id"] = u.LevelId
	u.fieldMap["super_agent_id"] = u.SuperAgentId
	u.fieldMap["first_name"] = u.FirstName
	u.fieldMap["other_names"] = u.OtherNames
	u.fieldMap["business_name"] = u.BusinessName
	u.fieldMap["email"] = u.Email
	u.fieldMap["phone"] = u.Phone
	u.fieldMap["gender"] = u.Gender
	u.fieldMap["dob"] = u.Dob
	u.fieldMap["state"] = u.State
	u.fieldMap["country"] = u.Country
	u.fieldMap["address"] = u.Address
	u.fieldMap["status"] = u.Status
	u.fieldMap["avatar"] = u.Avatar
	u.fieldMap["bvn"] = u.Bvn
	u.fieldMap["nin"] = u.Nin
	u.fieldMap["referral_code"] = u.ReferralCode
	u.fieldMap["email_verified_at"] = u.EmailVerifiedAt
	u.fieldMap["password"] = u.Password
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u user) replaceDB(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userDo struct{ gen.DO }

type IUserDo interface {
	gen.SubQuery
	Debug() IUserDo
	WithContext(ctx context.Context) IUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserDo
	WriteDB() IUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserDo
	Not(conds ...gen.Condition) IUserDo
	Or(conds ...gen.Condition) IUserDo
	Select(conds ...field.Expr) IUserDo
	Where(conds ...gen.Condition) IUserDo
	Order(conds ...field.Expr) IUserDo
	Distinct(cols ...field.Expr) IUserDo
	Omit(cols ...field.Expr) IUserDo
	Join(table schema.Tabler, on ...field.Expr) IUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserDo
	Group(cols ...field.Expr) IUserDo
	Having(conds ...gen.Condition) IUserDo
	Limit(limit int) IUserDo
	Offset(offset int) IUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo
	Unscoped() IUserDo
	Create(values ...*models.User) error
	CreateInBatches(values []*models.User, batchSize int) error
	Save(values ...*models.User) error
	First() (*models.User, error)
	Take() (*models.User, error)
	Last() (*models.User, error)
	Find() ([]*models.User, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.User, err error)
	FindInBatches(result *[]*models.User, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.User) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserDo
	Assign(attrs ...field.AssignExpr) IUserDo
	Joins(fields ...field.RelationField) IUserDo
	Preload(fields ...field.RelationField) IUserDo
	FirstOrInit() (*models.User, error)
	FirstOrCreate() (*models.User, error)
	FindByPage(offset int, limit int) (result []*models.User, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userDo) Debug() IUserDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) IUserDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) ReadDB() IUserDo {
	return u.Clauses(dbresolver.Read)
}

func (u userDo) WriteDB() IUserDo {
	return u.Clauses(dbresolver.Write)
}

func (u userDo) Session(config *gorm.Session) IUserDo {
	return u.withDO(u.DO.Session(config))
}

func (u userDo) Clauses(conds ...clause.Expression) IUserDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Returning(value interface{}, columns ...string) IUserDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userDo) Not(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) IUserDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) IUserDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() IUserDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*models.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*models.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*models.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*models.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.User), nil
	}
}

func (u userDo) Take() (*models.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.User), nil
	}
}

func (u userDo) Last() (*models.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.User), nil
	}
}

func (u userDo) Find() ([]*models.User, error) {
	result, err := u.DO.Find()
	return result.([]*models.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.User, err error) {
	buf := make([]*models.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*models.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userDo) Preload(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userDo) FirstOrInit() (*models.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.User), nil
	}
}

func (u userDo) FirstOrCreate() (*models.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*models.User, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userDo) Delete(models ...*models.User) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}
