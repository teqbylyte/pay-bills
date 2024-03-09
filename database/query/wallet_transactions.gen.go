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

func newWalletTransaction(db *gorm.DB, opts ...gen.DOOption) walletTransaction {
	_walletTransaction := walletTransaction{}

	_walletTransaction.walletTransactionDo.UseDB(db, opts...)
	_walletTransaction.walletTransactionDo.UseModel(&models.WalletTransaction{})

	tableName := _walletTransaction.walletTransactionDo.TableName()
	_walletTransaction.ALL = field.NewAsterisk(tableName)
	_walletTransaction.ID = field.NewUint(tableName, "id")
	_walletTransaction.CreatedAt = field.NewTime(tableName, "created_at")
	_walletTransaction.UpdatedAt = field.NewTime(tableName, "updated_at")
	_walletTransaction.ProductId = field.NewUint(tableName, "product_id")
	_walletTransaction.WalletId = field.NewUint(tableName, "wallet_id")
	_walletTransaction.Reference = field.NewString(tableName, "reference")
	_walletTransaction.Amount = field.NewFloat64(tableName, "amount")
	_walletTransaction.PrevBalance = field.NewFloat64(tableName, "prev_balance")
	_walletTransaction.NewBalance = field.NewFloat64(tableName, "new_balance")
	_walletTransaction.Status = field.NewString(tableName, "status")
	_walletTransaction.Action = field.NewString(tableName, "action")
	_walletTransaction.Type = field.NewString(tableName, "type")
	_walletTransaction.Info = field.NewString(tableName, "info")

	_walletTransaction.fillFieldMap()

	return _walletTransaction
}

type walletTransaction struct {
	walletTransactionDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	ProductId   field.Uint
	WalletId    field.Uint
	Reference   field.String
	Amount      field.Float64
	PrevBalance field.Float64
	NewBalance  field.Float64
	Status      field.String
	Action      field.String
	Type        field.String
	Info        field.String

	fieldMap map[string]field.Expr
}

func (w walletTransaction) Table(newTableName string) *walletTransaction {
	w.walletTransactionDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w walletTransaction) As(alias string) *walletTransaction {
	w.walletTransactionDo.DO = *(w.walletTransactionDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *walletTransaction) updateTableName(table string) *walletTransaction {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewUint(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.ProductId = field.NewUint(table, "product_id")
	w.WalletId = field.NewUint(table, "wallet_id")
	w.Reference = field.NewString(table, "reference")
	w.Amount = field.NewFloat64(table, "amount")
	w.PrevBalance = field.NewFloat64(table, "prev_balance")
	w.NewBalance = field.NewFloat64(table, "new_balance")
	w.Status = field.NewString(table, "status")
	w.Action = field.NewString(table, "action")
	w.Type = field.NewString(table, "type")
	w.Info = field.NewString(table, "info")

	w.fillFieldMap()

	return w
}

func (w *walletTransaction) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *walletTransaction) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 13)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["product_id"] = w.ProductId
	w.fieldMap["wallet_id"] = w.WalletId
	w.fieldMap["reference"] = w.Reference
	w.fieldMap["amount"] = w.Amount
	w.fieldMap["prev_balance"] = w.PrevBalance
	w.fieldMap["new_balance"] = w.NewBalance
	w.fieldMap["status"] = w.Status
	w.fieldMap["action"] = w.Action
	w.fieldMap["type"] = w.Type
	w.fieldMap["info"] = w.Info
}

func (w walletTransaction) clone(db *gorm.DB) walletTransaction {
	w.walletTransactionDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w walletTransaction) replaceDB(db *gorm.DB) walletTransaction {
	w.walletTransactionDo.ReplaceDB(db)
	return w
}

type walletTransactionDo struct{ gen.DO }

type IWalletTransactionDo interface {
	gen.SubQuery
	Debug() IWalletTransactionDo
	WithContext(ctx context.Context) IWalletTransactionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWalletTransactionDo
	WriteDB() IWalletTransactionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWalletTransactionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWalletTransactionDo
	Not(conds ...gen.Condition) IWalletTransactionDo
	Or(conds ...gen.Condition) IWalletTransactionDo
	Select(conds ...field.Expr) IWalletTransactionDo
	Where(conds ...gen.Condition) IWalletTransactionDo
	Order(conds ...field.Expr) IWalletTransactionDo
	Distinct(cols ...field.Expr) IWalletTransactionDo
	Omit(cols ...field.Expr) IWalletTransactionDo
	Join(table schema.Tabler, on ...field.Expr) IWalletTransactionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWalletTransactionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWalletTransactionDo
	Group(cols ...field.Expr) IWalletTransactionDo
	Having(conds ...gen.Condition) IWalletTransactionDo
	Limit(limit int) IWalletTransactionDo
	Offset(offset int) IWalletTransactionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWalletTransactionDo
	Unscoped() IWalletTransactionDo
	Create(values ...*models.WalletTransaction) error
	CreateInBatches(values []*models.WalletTransaction, batchSize int) error
	Save(values ...*models.WalletTransaction) error
	First() (*models.WalletTransaction, error)
	Take() (*models.WalletTransaction, error)
	Last() (*models.WalletTransaction, error)
	Find() ([]*models.WalletTransaction, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.WalletTransaction, err error)
	FindInBatches(result *[]*models.WalletTransaction, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.WalletTransaction) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWalletTransactionDo
	Assign(attrs ...field.AssignExpr) IWalletTransactionDo
	Joins(fields ...field.RelationField) IWalletTransactionDo
	Preload(fields ...field.RelationField) IWalletTransactionDo
	FirstOrInit() (*models.WalletTransaction, error)
	FirstOrCreate() (*models.WalletTransaction, error)
	FindByPage(offset int, limit int) (result []*models.WalletTransaction, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWalletTransactionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w walletTransactionDo) Debug() IWalletTransactionDo {
	return w.withDO(w.DO.Debug())
}

func (w walletTransactionDo) WithContext(ctx context.Context) IWalletTransactionDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w walletTransactionDo) ReadDB() IWalletTransactionDo {
	return w.Clauses(dbresolver.Read)
}

func (w walletTransactionDo) WriteDB() IWalletTransactionDo {
	return w.Clauses(dbresolver.Write)
}

func (w walletTransactionDo) Session(config *gorm.Session) IWalletTransactionDo {
	return w.withDO(w.DO.Session(config))
}

func (w walletTransactionDo) Clauses(conds ...clause.Expression) IWalletTransactionDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w walletTransactionDo) Returning(value interface{}, columns ...string) IWalletTransactionDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w walletTransactionDo) Not(conds ...gen.Condition) IWalletTransactionDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w walletTransactionDo) Or(conds ...gen.Condition) IWalletTransactionDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w walletTransactionDo) Select(conds ...field.Expr) IWalletTransactionDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w walletTransactionDo) Where(conds ...gen.Condition) IWalletTransactionDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w walletTransactionDo) Order(conds ...field.Expr) IWalletTransactionDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w walletTransactionDo) Distinct(cols ...field.Expr) IWalletTransactionDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w walletTransactionDo) Omit(cols ...field.Expr) IWalletTransactionDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w walletTransactionDo) Join(table schema.Tabler, on ...field.Expr) IWalletTransactionDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w walletTransactionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWalletTransactionDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w walletTransactionDo) RightJoin(table schema.Tabler, on ...field.Expr) IWalletTransactionDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w walletTransactionDo) Group(cols ...field.Expr) IWalletTransactionDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w walletTransactionDo) Having(conds ...gen.Condition) IWalletTransactionDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w walletTransactionDo) Limit(limit int) IWalletTransactionDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w walletTransactionDo) Offset(offset int) IWalletTransactionDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w walletTransactionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWalletTransactionDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w walletTransactionDo) Unscoped() IWalletTransactionDo {
	return w.withDO(w.DO.Unscoped())
}

func (w walletTransactionDo) Create(values ...*models.WalletTransaction) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w walletTransactionDo) CreateInBatches(values []*models.WalletTransaction, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w walletTransactionDo) Save(values ...*models.WalletTransaction) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w walletTransactionDo) First() (*models.WalletTransaction, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.WalletTransaction), nil
	}
}

func (w walletTransactionDo) Take() (*models.WalletTransaction, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.WalletTransaction), nil
	}
}

func (w walletTransactionDo) Last() (*models.WalletTransaction, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.WalletTransaction), nil
	}
}

func (w walletTransactionDo) Find() ([]*models.WalletTransaction, error) {
	result, err := w.DO.Find()
	return result.([]*models.WalletTransaction), err
}

func (w walletTransactionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.WalletTransaction, err error) {
	buf := make([]*models.WalletTransaction, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w walletTransactionDo) FindInBatches(result *[]*models.WalletTransaction, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w walletTransactionDo) Attrs(attrs ...field.AssignExpr) IWalletTransactionDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w walletTransactionDo) Assign(attrs ...field.AssignExpr) IWalletTransactionDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w walletTransactionDo) Joins(fields ...field.RelationField) IWalletTransactionDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w walletTransactionDo) Preload(fields ...field.RelationField) IWalletTransactionDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w walletTransactionDo) FirstOrInit() (*models.WalletTransaction, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.WalletTransaction), nil
	}
}

func (w walletTransactionDo) FirstOrCreate() (*models.WalletTransaction, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.WalletTransaction), nil
	}
}

func (w walletTransactionDo) FindByPage(offset int, limit int) (result []*models.WalletTransaction, count int64, err error) {
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

func (w walletTransactionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w walletTransactionDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w walletTransactionDo) Delete(models ...*models.WalletTransaction) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *walletTransactionDo) withDO(do gen.Dao) *walletTransactionDo {
	w.DO = *do.(*gen.DO)
	return w
}
