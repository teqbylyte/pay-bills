// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	model "martpay/app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newWallet(db *gorm.DB, opts ...gen.DOOption) wallet {
	_wallet := wallet{}

	_wallet.walletDo.UseDB(db, opts...)
	_wallet.walletDo.UseModel(&model.Wallet{})

	tableName := _wallet.walletDo.TableName()
	_wallet.ALL = field.NewAsterisk(tableName)
	_wallet.ID = field.NewUint(tableName, "id")
	_wallet.CreatedAt = field.NewTime(tableName, "created_at")
	_wallet.UpdatedAt = field.NewTime(tableName, "updated_at")
	_wallet.Uwid = field.NewString(tableName, "uwid")
	_wallet.UserId = field.NewUint(tableName, "user_id")
	_wallet.AccountNumber = field.NewString(tableName, "account_number")
	_wallet.Balance = field.NewFloat64(tableName, "balance")
	_wallet.Commission = field.NewFloat64(tableName, "commission")
	_wallet.Status = field.NewString(tableName, "status")
	_wallet.DisableDebit = field.NewBool(tableName, "disable_debit")
	_wallet.Meta = field.NewField(tableName, "meta")

	_wallet.fillFieldMap()

	return _wallet
}

type wallet struct {
	walletDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	Uwid          field.String
	UserId        field.Uint
	AccountNumber field.String
	Balance       field.Float64
	Commission    field.Float64
	Status        field.String
	DisableDebit  field.Bool
	Meta          field.Field

	fieldMap map[string]field.Expr
}

func (w wallet) Table(newTableName string) *wallet {
	w.walletDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wallet) As(alias string) *wallet {
	w.walletDo.DO = *(w.walletDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wallet) updateTableName(table string) *wallet {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewUint(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.Uwid = field.NewString(table, "uwid")
	w.UserId = field.NewUint(table, "user_id")
	w.AccountNumber = field.NewString(table, "account_number")
	w.Balance = field.NewFloat64(table, "balance")
	w.Commission = field.NewFloat64(table, "commission")
	w.Status = field.NewString(table, "status")
	w.DisableDebit = field.NewBool(table, "disable_debit")
	w.Meta = field.NewField(table, "meta")

	w.fillFieldMap()

	return w
}

func (w *wallet) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wallet) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["uwid"] = w.Uwid
	w.fieldMap["user_id"] = w.UserId
	w.fieldMap["account_number"] = w.AccountNumber
	w.fieldMap["balance"] = w.Balance
	w.fieldMap["commission"] = w.Commission
	w.fieldMap["status"] = w.Status
	w.fieldMap["disable_debit"] = w.DisableDebit
	w.fieldMap["meta"] = w.Meta
}

func (w wallet) clone(db *gorm.DB) wallet {
	w.walletDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wallet) replaceDB(db *gorm.DB) wallet {
	w.walletDo.ReplaceDB(db)
	return w
}

type walletDo struct{ gen.DO }

type IWalletDo interface {
	gen.SubQuery
	Debug() IWalletDo
	WithContext(ctx context.Context) IWalletDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWalletDo
	WriteDB() IWalletDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWalletDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWalletDo
	Not(conds ...gen.Condition) IWalletDo
	Or(conds ...gen.Condition) IWalletDo
	Select(conds ...field.Expr) IWalletDo
	Where(conds ...gen.Condition) IWalletDo
	Order(conds ...field.Expr) IWalletDo
	Distinct(cols ...field.Expr) IWalletDo
	Omit(cols ...field.Expr) IWalletDo
	Join(table schema.Tabler, on ...field.Expr) IWalletDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWalletDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWalletDo
	Group(cols ...field.Expr) IWalletDo
	Having(conds ...gen.Condition) IWalletDo
	Limit(limit int) IWalletDo
	Offset(offset int) IWalletDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWalletDo
	Unscoped() IWalletDo
	Create(values ...*model.Wallet) error
	CreateInBatches(values []*model.Wallet, batchSize int) error
	Save(values ...*model.Wallet) error
	First() (*model.Wallet, error)
	Take() (*model.Wallet, error)
	Last() (*model.Wallet, error)
	Find() ([]*model.Wallet, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Wallet, err error)
	FindInBatches(result *[]*model.Wallet, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Wallet) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWalletDo
	Assign(attrs ...field.AssignExpr) IWalletDo
	Joins(fields ...field.RelationField) IWalletDo
	Preload(fields ...field.RelationField) IWalletDo
	FirstOrInit() (*model.Wallet, error)
	FirstOrCreate() (*model.Wallet, error)
	FindByPage(offset int, limit int) (result []*model.Wallet, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWalletDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w walletDo) Debug() IWalletDo {
	return w.withDO(w.DO.Debug())
}

func (w walletDo) WithContext(ctx context.Context) IWalletDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w walletDo) ReadDB() IWalletDo {
	return w.Clauses(dbresolver.Read)
}

func (w walletDo) WriteDB() IWalletDo {
	return w.Clauses(dbresolver.Write)
}

func (w walletDo) Session(config *gorm.Session) IWalletDo {
	return w.withDO(w.DO.Session(config))
}

func (w walletDo) Clauses(conds ...clause.Expression) IWalletDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w walletDo) Returning(value interface{}, columns ...string) IWalletDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w walletDo) Not(conds ...gen.Condition) IWalletDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w walletDo) Or(conds ...gen.Condition) IWalletDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w walletDo) Select(conds ...field.Expr) IWalletDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w walletDo) Where(conds ...gen.Condition) IWalletDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w walletDo) Order(conds ...field.Expr) IWalletDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w walletDo) Distinct(cols ...field.Expr) IWalletDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w walletDo) Omit(cols ...field.Expr) IWalletDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w walletDo) Join(table schema.Tabler, on ...field.Expr) IWalletDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w walletDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWalletDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w walletDo) RightJoin(table schema.Tabler, on ...field.Expr) IWalletDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w walletDo) Group(cols ...field.Expr) IWalletDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w walletDo) Having(conds ...gen.Condition) IWalletDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w walletDo) Limit(limit int) IWalletDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w walletDo) Offset(offset int) IWalletDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w walletDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWalletDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w walletDo) Unscoped() IWalletDo {
	return w.withDO(w.DO.Unscoped())
}

func (w walletDo) Create(values ...*model.Wallet) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w walletDo) CreateInBatches(values []*model.Wallet, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w walletDo) Save(values ...*model.Wallet) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w walletDo) First() (*model.Wallet, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Wallet), nil
	}
}

func (w walletDo) Take() (*model.Wallet, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Wallet), nil
	}
}

func (w walletDo) Last() (*model.Wallet, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Wallet), nil
	}
}

func (w walletDo) Find() ([]*model.Wallet, error) {
	result, err := w.DO.Find()
	return result.([]*model.Wallet), err
}

func (w walletDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Wallet, err error) {
	buf := make([]*model.Wallet, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w walletDo) FindInBatches(result *[]*model.Wallet, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w walletDo) Attrs(attrs ...field.AssignExpr) IWalletDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w walletDo) Assign(attrs ...field.AssignExpr) IWalletDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w walletDo) Joins(fields ...field.RelationField) IWalletDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w walletDo) Preload(fields ...field.RelationField) IWalletDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w walletDo) FirstOrInit() (*model.Wallet, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Wallet), nil
	}
}

func (w walletDo) FirstOrCreate() (*model.Wallet, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Wallet), nil
	}
}

func (w walletDo) FindByPage(offset int, limit int) (result []*model.Wallet, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w walletDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w walletDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w walletDo) Delete(models ...*model.Wallet) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *walletDo) withDO(do gen.Dao) *walletDo {
	w.DO = *do.(*gen.DO)
	return w
}
