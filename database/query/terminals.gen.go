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

func newTerminal(db *gorm.DB, opts ...gen.DOOption) terminal {
	_terminal := terminal{}

	_terminal.terminalDo.UseDB(db, opts...)
	_terminal.terminalDo.UseModel(&models.Terminal{})

	tableName := _terminal.terminalDo.TableName()
	_terminal.ALL = field.NewAsterisk(tableName)
	_terminal.ID = field.NewUint(tableName, "id")
	_terminal.CreatedAt = field.NewTime(tableName, "created_at")
	_terminal.UpdatedAt = field.NewTime(tableName, "updated_at")
	_terminal.UserId = field.NewUint(tableName, "user_id")
	_terminal.GroupId = field.NewUint(tableName, "group_id")
	_terminal.Device = field.NewString(tableName, "device")
	_terminal.Serial = field.NewString(tableName, "serial")
	_terminal.Status = field.NewString(tableName, "status")
	_terminal.Tid = field.NewString(tableName, "tid")
	_terminal.Mid = field.NewString(tableName, "mid")
	_terminal.Tmk = field.NewString(tableName, "tmk")
	_terminal.Tsk = field.NewString(tableName, "tsk")
	_terminal.Tpk = field.NewString(tableName, "tpk")
	_terminal.DateTime = field.NewString(tableName, "date_time")
	_terminal.Timeout = field.NewInt(tableName, "timeout")
	_terminal.CurrencyCode = field.NewString(tableName, "currency_code")
	_terminal.CountryCode = field.NewString(tableName, "country_code")
	_terminal.CategoryCode = field.NewString(tableName, "category_code")
	_terminal.NameLocation = field.NewString(tableName, "name_location")
	_terminal.AdminPin = field.NewString(tableName, "admin_pin")
	_terminal.Pin = field.NewString(tableName, "pin")
	_terminal.WrongPinCount = field.NewInt(tableName, "wrong_pin_count")
	_terminal.HasChangedPin = field.NewBool(tableName, "has_changed_pin")
	_terminal.User = terminalBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "models.User"),
	}

	_terminal.fillFieldMap()

	return _terminal
}

type terminal struct {
	terminalDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	UserId        field.Uint
	GroupId       field.Uint
	Device        field.String
	Serial        field.String
	Status        field.String
	Tid           field.String
	Mid           field.String
	Tmk           field.String
	Tsk           field.String
	Tpk           field.String
	DateTime      field.String
	Timeout       field.Int
	CurrencyCode  field.String
	CountryCode   field.String
	CategoryCode  field.String
	NameLocation  field.String
	AdminPin      field.String
	Pin           field.String
	WrongPinCount field.Int
	HasChangedPin field.Bool
	User          terminalBelongsToUser

	fieldMap map[string]field.Expr
}

func (t terminal) Table(newTableName string) *terminal {
	t.terminalDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t terminal) As(alias string) *terminal {
	t.terminalDo.DO = *(t.terminalDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *terminal) updateTableName(table string) *terminal {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.UserId = field.NewUint(table, "user_id")
	t.GroupId = field.NewUint(table, "group_id")
	t.Device = field.NewString(table, "device")
	t.Serial = field.NewString(table, "serial")
	t.Status = field.NewString(table, "status")
	t.Tid = field.NewString(table, "tid")
	t.Mid = field.NewString(table, "mid")
	t.Tmk = field.NewString(table, "tmk")
	t.Tsk = field.NewString(table, "tsk")
	t.Tpk = field.NewString(table, "tpk")
	t.DateTime = field.NewString(table, "date_time")
	t.Timeout = field.NewInt(table, "timeout")
	t.CurrencyCode = field.NewString(table, "currency_code")
	t.CountryCode = field.NewString(table, "country_code")
	t.CategoryCode = field.NewString(table, "category_code")
	t.NameLocation = field.NewString(table, "name_location")
	t.AdminPin = field.NewString(table, "admin_pin")
	t.Pin = field.NewString(table, "pin")
	t.WrongPinCount = field.NewInt(table, "wrong_pin_count")
	t.HasChangedPin = field.NewBool(table, "has_changed_pin")

	t.fillFieldMap()

	return t
}

func (t *terminal) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *terminal) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 24)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["user_id"] = t.UserId
	t.fieldMap["group_id"] = t.GroupId
	t.fieldMap["device"] = t.Device
	t.fieldMap["serial"] = t.Serial
	t.fieldMap["status"] = t.Status
	t.fieldMap["tid"] = t.Tid
	t.fieldMap["mid"] = t.Mid
	t.fieldMap["tmk"] = t.Tmk
	t.fieldMap["tsk"] = t.Tsk
	t.fieldMap["tpk"] = t.Tpk
	t.fieldMap["date_time"] = t.DateTime
	t.fieldMap["timeout"] = t.Timeout
	t.fieldMap["currency_code"] = t.CurrencyCode
	t.fieldMap["country_code"] = t.CountryCode
	t.fieldMap["category_code"] = t.CategoryCode
	t.fieldMap["name_location"] = t.NameLocation
	t.fieldMap["admin_pin"] = t.AdminPin
	t.fieldMap["pin"] = t.Pin
	t.fieldMap["wrong_pin_count"] = t.WrongPinCount
	t.fieldMap["has_changed_pin"] = t.HasChangedPin

}

func (t terminal) clone(db *gorm.DB) terminal {
	t.terminalDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t terminal) replaceDB(db *gorm.DB) terminal {
	t.terminalDo.ReplaceDB(db)
	return t
}

type terminalBelongsToUser struct {
	db *gorm.DB

	field.RelationField
}

func (a terminalBelongsToUser) Where(conds ...field.Expr) *terminalBelongsToUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a terminalBelongsToUser) WithContext(ctx context.Context) *terminalBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a terminalBelongsToUser) Session(session *gorm.Session) *terminalBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a terminalBelongsToUser) Model(m *models.Terminal) *terminalBelongsToUserTx {
	return &terminalBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

type terminalBelongsToUserTx struct{ tx *gorm.Association }

func (a terminalBelongsToUserTx) Find() (result *models.User, err error) {
	return result, a.tx.Find(&result)
}

func (a terminalBelongsToUserTx) Append(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a terminalBelongsToUserTx) Replace(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a terminalBelongsToUserTx) Delete(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a terminalBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a terminalBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

type terminalDo struct{ gen.DO }

type ITerminalDo interface {
	gen.SubQuery
	Debug() ITerminalDo
	WithContext(ctx context.Context) ITerminalDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITerminalDo
	WriteDB() ITerminalDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITerminalDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITerminalDo
	Not(conds ...gen.Condition) ITerminalDo
	Or(conds ...gen.Condition) ITerminalDo
	Select(conds ...field.Expr) ITerminalDo
	Where(conds ...gen.Condition) ITerminalDo
	Order(conds ...field.Expr) ITerminalDo
	Distinct(cols ...field.Expr) ITerminalDo
	Omit(cols ...field.Expr) ITerminalDo
	Join(table schema.Tabler, on ...field.Expr) ITerminalDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITerminalDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITerminalDo
	Group(cols ...field.Expr) ITerminalDo
	Having(conds ...gen.Condition) ITerminalDo
	Limit(limit int) ITerminalDo
	Offset(offset int) ITerminalDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITerminalDo
	Unscoped() ITerminalDo
	Create(values ...*models.Terminal) error
	CreateInBatches(values []*models.Terminal, batchSize int) error
	Save(values ...*models.Terminal) error
	First() (*models.Terminal, error)
	Take() (*models.Terminal, error)
	Last() (*models.Terminal, error)
	Find() ([]*models.Terminal, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Terminal, err error)
	FindInBatches(result *[]*models.Terminal, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Terminal) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITerminalDo
	Assign(attrs ...field.AssignExpr) ITerminalDo
	Joins(fields ...field.RelationField) ITerminalDo
	Preload(fields ...field.RelationField) ITerminalDo
	FirstOrInit() (*models.Terminal, error)
	FirstOrCreate() (*models.Terminal, error)
	FindByPage(offset int, limit int) (result []*models.Terminal, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITerminalDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t terminalDo) Debug() ITerminalDo {
	return t.withDO(t.DO.Debug())
}

func (t terminalDo) WithContext(ctx context.Context) ITerminalDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t terminalDo) ReadDB() ITerminalDo {
	return t.Clauses(dbresolver.Read)
}

func (t terminalDo) WriteDB() ITerminalDo {
	return t.Clauses(dbresolver.Write)
}

func (t terminalDo) Session(config *gorm.Session) ITerminalDo {
	return t.withDO(t.DO.Session(config))
}

func (t terminalDo) Clauses(conds ...clause.Expression) ITerminalDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t terminalDo) Returning(value interface{}, columns ...string) ITerminalDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t terminalDo) Not(conds ...gen.Condition) ITerminalDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t terminalDo) Or(conds ...gen.Condition) ITerminalDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t terminalDo) Select(conds ...field.Expr) ITerminalDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t terminalDo) Where(conds ...gen.Condition) ITerminalDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t terminalDo) Order(conds ...field.Expr) ITerminalDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t terminalDo) Distinct(cols ...field.Expr) ITerminalDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t terminalDo) Omit(cols ...field.Expr) ITerminalDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t terminalDo) Join(table schema.Tabler, on ...field.Expr) ITerminalDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t terminalDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITerminalDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t terminalDo) RightJoin(table schema.Tabler, on ...field.Expr) ITerminalDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t terminalDo) Group(cols ...field.Expr) ITerminalDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t terminalDo) Having(conds ...gen.Condition) ITerminalDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t terminalDo) Limit(limit int) ITerminalDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t terminalDo) Offset(offset int) ITerminalDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t terminalDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITerminalDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t terminalDo) Unscoped() ITerminalDo {
	return t.withDO(t.DO.Unscoped())
}

func (t terminalDo) Create(values ...*models.Terminal) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t terminalDo) CreateInBatches(values []*models.Terminal, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t terminalDo) Save(values ...*models.Terminal) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t terminalDo) First() (*models.Terminal, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Terminal), nil
	}
}

func (t terminalDo) Take() (*models.Terminal, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Terminal), nil
	}
}

func (t terminalDo) Last() (*models.Terminal, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Terminal), nil
	}
}

func (t terminalDo) Find() ([]*models.Terminal, error) {
	result, err := t.DO.Find()
	return result.([]*models.Terminal), err
}

func (t terminalDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Terminal, err error) {
	buf := make([]*models.Terminal, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t terminalDo) FindInBatches(result *[]*models.Terminal, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t terminalDo) Attrs(attrs ...field.AssignExpr) ITerminalDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t terminalDo) Assign(attrs ...field.AssignExpr) ITerminalDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t terminalDo) Joins(fields ...field.RelationField) ITerminalDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t terminalDo) Preload(fields ...field.RelationField) ITerminalDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t terminalDo) FirstOrInit() (*models.Terminal, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Terminal), nil
	}
}

func (t terminalDo) FirstOrCreate() (*models.Terminal, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Terminal), nil
	}
}

func (t terminalDo) FindByPage(offset int, limit int) (result []*models.Terminal, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t terminalDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t terminalDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t terminalDo) Delete(models ...*models.Terminal) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *terminalDo) withDO(do gen.Dao) *terminalDo {
	t.DO = *do.(*gen.DO)
	return t
}