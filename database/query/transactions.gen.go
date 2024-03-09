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

func newTransactions(db *gorm.DB, opts ...gen.DOOption) transactions {
	_transactions := transactions{}

	_transactions.transactionsDo.UseDB(db, opts...)
	_transactions.transactionsDo.UseModel(&models.Transactions{})

	tableName := _transactions.transactionsDo.TableName()
	_transactions.ALL = field.NewAsterisk(tableName)
	_transactions.ID = field.NewUint(tableName, "id")
	_transactions.UserId = field.NewUint(tableName, "user_id")
	_transactions.TerminalId = field.NewUint(tableName, "terminal_id")
	_transactions.TypeId = field.NewUint(tableName, "type_id")
	_transactions.Amount = field.NewFloat64(tableName, "amount")
	_transactions.Charge = field.NewFloat64(tableName, "charge")
	_transactions.TotalAmount = field.NewFloat64(tableName, "total_amount")
	_transactions.Reference = field.NewString(tableName, "reference")
	_transactions.ResponseCode = field.NewString(tableName, "response_code")
	_transactions.Stan = field.NewString(tableName, "stan")
	_transactions.BankName = field.NewString(tableName, "bank_name")
	_transactions.BankCode = field.NewString(tableName, "bank_code")
	_transactions.AccountNumber = field.NewString(tableName, "account_number")
	_transactions.AccountName = field.NewString(tableName, "account_name")
	_transactions.Info = field.NewString(tableName, "info")
	_transactions.PowerToken = field.NewString(tableName, "power_token")
	_transactions.Status = field.NewString(tableName, "status")
	_transactions.Channel = field.NewString(tableName, "channel")
	_transactions.Provider = field.NewString(tableName, "provider")
	_transactions.Meta = field.NewField(tableName, "meta")
	_transactions.WalletDebited = field.NewBool(tableName, "wallet_debited")
	_transactions.WalletCredited = field.NewBool(tableName, "wallet_credited")
	_transactions.Version = field.NewString(tableName, "version")
	_transactions.Device = field.NewString(tableName, "device")
	_transactions.CreatedAt = field.NewTime(tableName, "created_at")
	_transactions.UpdatedAt = field.NewTime(tableName, "updated_at")

	_transactions.fillFieldMap()

	return _transactions
}

type transactions struct {
	transactionsDo

	ALL            field.Asterisk
	ID             field.Uint
	UserId         field.Uint
	TerminalId     field.Uint
	TypeId         field.Uint
	Amount         field.Float64
	Charge         field.Float64
	TotalAmount    field.Float64
	Reference      field.String
	ResponseCode   field.String
	Stan           field.String
	BankName       field.String
	BankCode       field.String
	AccountNumber  field.String
	AccountName    field.String
	Info           field.String
	PowerToken     field.String
	Status         field.String
	Channel        field.String
	Provider       field.String
	Meta           field.Field
	WalletDebited  field.Bool
	WalletCredited field.Bool
	Version        field.String
	Device         field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time

	fieldMap map[string]field.Expr
}

func (t transactions) Table(newTableName string) *transactions {
	t.transactionsDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t transactions) As(alias string) *transactions {
	t.transactionsDo.DO = *(t.transactionsDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *transactions) updateTableName(table string) *transactions {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.UserId = field.NewUint(table, "user_id")
	t.TerminalId = field.NewUint(table, "terminal_id")
	t.TypeId = field.NewUint(table, "type_id")
	t.Amount = field.NewFloat64(table, "amount")
	t.Charge = field.NewFloat64(table, "charge")
	t.TotalAmount = field.NewFloat64(table, "total_amount")
	t.Reference = field.NewString(table, "reference")
	t.ResponseCode = field.NewString(table, "response_code")
	t.Stan = field.NewString(table, "stan")
	t.BankName = field.NewString(table, "bank_name")
	t.BankCode = field.NewString(table, "bank_code")
	t.AccountNumber = field.NewString(table, "account_number")
	t.AccountName = field.NewString(table, "account_name")
	t.Info = field.NewString(table, "info")
	t.PowerToken = field.NewString(table, "power_token")
	t.Status = field.NewString(table, "status")
	t.Channel = field.NewString(table, "channel")
	t.Provider = field.NewString(table, "provider")
	t.Meta = field.NewField(table, "meta")
	t.WalletDebited = field.NewBool(table, "wallet_debited")
	t.WalletCredited = field.NewBool(table, "wallet_credited")
	t.Version = field.NewString(table, "version")
	t.Device = field.NewString(table, "device")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")

	t.fillFieldMap()

	return t
}

func (t *transactions) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *transactions) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 26)
	t.fieldMap["id"] = t.ID
	t.fieldMap["user_id"] = t.UserId
	t.fieldMap["terminal_id"] = t.TerminalId
	t.fieldMap["type_id"] = t.TypeId
	t.fieldMap["amount"] = t.Amount
	t.fieldMap["charge"] = t.Charge
	t.fieldMap["total_amount"] = t.TotalAmount
	t.fieldMap["reference"] = t.Reference
	t.fieldMap["response_code"] = t.ResponseCode
	t.fieldMap["stan"] = t.Stan
	t.fieldMap["bank_name"] = t.BankName
	t.fieldMap["bank_code"] = t.BankCode
	t.fieldMap["account_number"] = t.AccountNumber
	t.fieldMap["account_name"] = t.AccountName
	t.fieldMap["info"] = t.Info
	t.fieldMap["power_token"] = t.PowerToken
	t.fieldMap["status"] = t.Status
	t.fieldMap["channel"] = t.Channel
	t.fieldMap["provider"] = t.Provider
	t.fieldMap["meta"] = t.Meta
	t.fieldMap["wallet_debited"] = t.WalletDebited
	t.fieldMap["wallet_credited"] = t.WalletCredited
	t.fieldMap["version"] = t.Version
	t.fieldMap["device"] = t.Device
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
}

func (t transactions) clone(db *gorm.DB) transactions {
	t.transactionsDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t transactions) replaceDB(db *gorm.DB) transactions {
	t.transactionsDo.ReplaceDB(db)
	return t
}

type transactionsDo struct{ gen.DO }

type ITransactionsDo interface {
	gen.SubQuery
	Debug() ITransactionsDo
	WithContext(ctx context.Context) ITransactionsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITransactionsDo
	WriteDB() ITransactionsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITransactionsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITransactionsDo
	Not(conds ...gen.Condition) ITransactionsDo
	Or(conds ...gen.Condition) ITransactionsDo
	Select(conds ...field.Expr) ITransactionsDo
	Where(conds ...gen.Condition) ITransactionsDo
	Order(conds ...field.Expr) ITransactionsDo
	Distinct(cols ...field.Expr) ITransactionsDo
	Omit(cols ...field.Expr) ITransactionsDo
	Join(table schema.Tabler, on ...field.Expr) ITransactionsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITransactionsDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITransactionsDo
	Group(cols ...field.Expr) ITransactionsDo
	Having(conds ...gen.Condition) ITransactionsDo
	Limit(limit int) ITransactionsDo
	Offset(offset int) ITransactionsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITransactionsDo
	Unscoped() ITransactionsDo
	Create(values ...*models.Transactions) error
	CreateInBatches(values []*models.Transactions, batchSize int) error
	Save(values ...*models.Transactions) error
	First() (*models.Transactions, error)
	Take() (*models.Transactions, error)
	Last() (*models.Transactions, error)
	Find() ([]*models.Transactions, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Transactions, err error)
	FindInBatches(result *[]*models.Transactions, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Transactions) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITransactionsDo
	Assign(attrs ...field.AssignExpr) ITransactionsDo
	Joins(fields ...field.RelationField) ITransactionsDo
	Preload(fields ...field.RelationField) ITransactionsDo
	FirstOrInit() (*models.Transactions, error)
	FirstOrCreate() (*models.Transactions, error)
	FindByPage(offset int, limit int) (result []*models.Transactions, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITransactionsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t transactionsDo) Debug() ITransactionsDo {
	return t.withDO(t.DO.Debug())
}

func (t transactionsDo) WithContext(ctx context.Context) ITransactionsDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t transactionsDo) ReadDB() ITransactionsDo {
	return t.Clauses(dbresolver.Read)
}

func (t transactionsDo) WriteDB() ITransactionsDo {
	return t.Clauses(dbresolver.Write)
}

func (t transactionsDo) Session(config *gorm.Session) ITransactionsDo {
	return t.withDO(t.DO.Session(config))
}

func (t transactionsDo) Clauses(conds ...clause.Expression) ITransactionsDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t transactionsDo) Returning(value interface{}, columns ...string) ITransactionsDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t transactionsDo) Not(conds ...gen.Condition) ITransactionsDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t transactionsDo) Or(conds ...gen.Condition) ITransactionsDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t transactionsDo) Select(conds ...field.Expr) ITransactionsDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t transactionsDo) Where(conds ...gen.Condition) ITransactionsDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t transactionsDo) Order(conds ...field.Expr) ITransactionsDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t transactionsDo) Distinct(cols ...field.Expr) ITransactionsDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t transactionsDo) Omit(cols ...field.Expr) ITransactionsDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t transactionsDo) Join(table schema.Tabler, on ...field.Expr) ITransactionsDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t transactionsDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITransactionsDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t transactionsDo) RightJoin(table schema.Tabler, on ...field.Expr) ITransactionsDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t transactionsDo) Group(cols ...field.Expr) ITransactionsDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t transactionsDo) Having(conds ...gen.Condition) ITransactionsDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t transactionsDo) Limit(limit int) ITransactionsDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t transactionsDo) Offset(offset int) ITransactionsDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t transactionsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITransactionsDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t transactionsDo) Unscoped() ITransactionsDo {
	return t.withDO(t.DO.Unscoped())
}

func (t transactionsDo) Create(values ...*models.Transactions) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t transactionsDo) CreateInBatches(values []*models.Transactions, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t transactionsDo) Save(values ...*models.Transactions) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t transactionsDo) First() (*models.Transactions, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Transactions), nil
	}
}

func (t transactionsDo) Take() (*models.Transactions, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Transactions), nil
	}
}

func (t transactionsDo) Last() (*models.Transactions, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Transactions), nil
	}
}

func (t transactionsDo) Find() ([]*models.Transactions, error) {
	result, err := t.DO.Find()
	return result.([]*models.Transactions), err
}

func (t transactionsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Transactions, err error) {
	buf := make([]*models.Transactions, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t transactionsDo) FindInBatches(result *[]*models.Transactions, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t transactionsDo) Attrs(attrs ...field.AssignExpr) ITransactionsDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t transactionsDo) Assign(attrs ...field.AssignExpr) ITransactionsDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t transactionsDo) Joins(fields ...field.RelationField) ITransactionsDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t transactionsDo) Preload(fields ...field.RelationField) ITransactionsDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t transactionsDo) FirstOrInit() (*models.Transactions, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Transactions), nil
	}
}

func (t transactionsDo) FirstOrCreate() (*models.Transactions, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Transactions), nil
	}
}

func (t transactionsDo) FindByPage(offset int, limit int) (result []*models.Transactions, count int64, err error) {
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

func (t transactionsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t transactionsDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t transactionsDo) Delete(models ...*models.Transactions) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *transactionsDo) withDO(do gen.Dao) *transactionsDo {
	t.DO = *do.(*gen.DO)
	return t
}
