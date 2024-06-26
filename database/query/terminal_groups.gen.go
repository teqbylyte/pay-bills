// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	model "pay-bills/app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newTerminalGroup(db *gorm.DB, opts ...gen.DOOption) terminalGroup {
	_terminalGroup := terminalGroup{}

	_terminalGroup.terminalGroupDo.UseDB(db, opts...)
	_terminalGroup.terminalGroupDo.UseModel(&model.TerminalGroup{})

	tableName := _terminalGroup.terminalGroupDo.TableName()
	_terminalGroup.ALL = field.NewAsterisk(tableName)
	_terminalGroup.ID = field.NewUint(tableName, "id")
	_terminalGroup.CreatedAt = field.NewTime(tableName, "created_at")
	_terminalGroup.UpdatedAt = field.NewTime(tableName, "updated_at")
	_terminalGroup.Name = field.NewString(tableName, "name")
	_terminalGroup.Info = field.NewString(tableName, "info")

	_terminalGroup.fillFieldMap()

	return _terminalGroup
}

type terminalGroup struct {
	terminalGroupDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	Name      field.String
	Info      field.String

	fieldMap map[string]field.Expr
}

func (t terminalGroup) Table(newTableName string) *terminalGroup {
	t.terminalGroupDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t terminalGroup) As(alias string) *terminalGroup {
	t.terminalGroupDo.DO = *(t.terminalGroupDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *terminalGroup) updateTableName(table string) *terminalGroup {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.Name = field.NewString(table, "name")
	t.Info = field.NewString(table, "info")

	t.fillFieldMap()

	return t
}

func (t *terminalGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *terminalGroup) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["name"] = t.Name
	t.fieldMap["info"] = t.Info
}

func (t terminalGroup) clone(db *gorm.DB) terminalGroup {
	t.terminalGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t terminalGroup) replaceDB(db *gorm.DB) terminalGroup {
	t.terminalGroupDo.ReplaceDB(db)
	return t
}

type terminalGroupDo struct{ gen.DO }

type ITerminalGroupDo interface {
	gen.SubQuery
	Debug() ITerminalGroupDo
	WithContext(ctx context.Context) ITerminalGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITerminalGroupDo
	WriteDB() ITerminalGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITerminalGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITerminalGroupDo
	Not(conds ...gen.Condition) ITerminalGroupDo
	Or(conds ...gen.Condition) ITerminalGroupDo
	Select(conds ...field.Expr) ITerminalGroupDo
	Where(conds ...gen.Condition) ITerminalGroupDo
	Order(conds ...field.Expr) ITerminalGroupDo
	Distinct(cols ...field.Expr) ITerminalGroupDo
	Omit(cols ...field.Expr) ITerminalGroupDo
	Join(table schema.Tabler, on ...field.Expr) ITerminalGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITerminalGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITerminalGroupDo
	Group(cols ...field.Expr) ITerminalGroupDo
	Having(conds ...gen.Condition) ITerminalGroupDo
	Limit(limit int) ITerminalGroupDo
	Offset(offset int) ITerminalGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITerminalGroupDo
	Unscoped() ITerminalGroupDo
	Create(values ...*model.TerminalGroup) error
	CreateInBatches(values []*model.TerminalGroup, batchSize int) error
	Save(values ...*model.TerminalGroup) error
	First() (*model.TerminalGroup, error)
	Take() (*model.TerminalGroup, error)
	Last() (*model.TerminalGroup, error)
	Find() ([]*model.TerminalGroup, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TerminalGroup, err error)
	FindInBatches(result *[]*model.TerminalGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TerminalGroup) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITerminalGroupDo
	Assign(attrs ...field.AssignExpr) ITerminalGroupDo
	Joins(fields ...field.RelationField) ITerminalGroupDo
	Preload(fields ...field.RelationField) ITerminalGroupDo
	FirstOrInit() (*model.TerminalGroup, error)
	FirstOrCreate() (*model.TerminalGroup, error)
	FindByPage(offset int, limit int) (result []*model.TerminalGroup, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITerminalGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t terminalGroupDo) Debug() ITerminalGroupDo {
	return t.withDO(t.DO.Debug())
}

func (t terminalGroupDo) WithContext(ctx context.Context) ITerminalGroupDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t terminalGroupDo) ReadDB() ITerminalGroupDo {
	return t.Clauses(dbresolver.Read)
}

func (t terminalGroupDo) WriteDB() ITerminalGroupDo {
	return t.Clauses(dbresolver.Write)
}

func (t terminalGroupDo) Session(config *gorm.Session) ITerminalGroupDo {
	return t.withDO(t.DO.Session(config))
}

func (t terminalGroupDo) Clauses(conds ...clause.Expression) ITerminalGroupDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t terminalGroupDo) Returning(value interface{}, columns ...string) ITerminalGroupDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t terminalGroupDo) Not(conds ...gen.Condition) ITerminalGroupDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t terminalGroupDo) Or(conds ...gen.Condition) ITerminalGroupDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t terminalGroupDo) Select(conds ...field.Expr) ITerminalGroupDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t terminalGroupDo) Where(conds ...gen.Condition) ITerminalGroupDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t terminalGroupDo) Order(conds ...field.Expr) ITerminalGroupDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t terminalGroupDo) Distinct(cols ...field.Expr) ITerminalGroupDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t terminalGroupDo) Omit(cols ...field.Expr) ITerminalGroupDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t terminalGroupDo) Join(table schema.Tabler, on ...field.Expr) ITerminalGroupDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t terminalGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITerminalGroupDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t terminalGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) ITerminalGroupDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t terminalGroupDo) Group(cols ...field.Expr) ITerminalGroupDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t terminalGroupDo) Having(conds ...gen.Condition) ITerminalGroupDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t terminalGroupDo) Limit(limit int) ITerminalGroupDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t terminalGroupDo) Offset(offset int) ITerminalGroupDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t terminalGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITerminalGroupDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t terminalGroupDo) Unscoped() ITerminalGroupDo {
	return t.withDO(t.DO.Unscoped())
}

func (t terminalGroupDo) Create(values ...*model.TerminalGroup) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t terminalGroupDo) CreateInBatches(values []*model.TerminalGroup, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t terminalGroupDo) Save(values ...*model.TerminalGroup) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t terminalGroupDo) First() (*model.TerminalGroup, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TerminalGroup), nil
	}
}

func (t terminalGroupDo) Take() (*model.TerminalGroup, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TerminalGroup), nil
	}
}

func (t terminalGroupDo) Last() (*model.TerminalGroup, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TerminalGroup), nil
	}
}

func (t terminalGroupDo) Find() ([]*model.TerminalGroup, error) {
	result, err := t.DO.Find()
	return result.([]*model.TerminalGroup), err
}

func (t terminalGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TerminalGroup, err error) {
	buf := make([]*model.TerminalGroup, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t terminalGroupDo) FindInBatches(result *[]*model.TerminalGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t terminalGroupDo) Attrs(attrs ...field.AssignExpr) ITerminalGroupDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t terminalGroupDo) Assign(attrs ...field.AssignExpr) ITerminalGroupDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t terminalGroupDo) Joins(fields ...field.RelationField) ITerminalGroupDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t terminalGroupDo) Preload(fields ...field.RelationField) ITerminalGroupDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t terminalGroupDo) FirstOrInit() (*model.TerminalGroup, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TerminalGroup), nil
	}
}

func (t terminalGroupDo) FirstOrCreate() (*model.TerminalGroup, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TerminalGroup), nil
	}
}

func (t terminalGroupDo) FindByPage(offset int, limit int) (result []*model.TerminalGroup, count int64, err error) {
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

func (t terminalGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t terminalGroupDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t terminalGroupDo) Delete(models ...*model.TerminalGroup) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *terminalGroupDo) withDO(do gen.Dao) *terminalGroupDo {
	t.DO = *do.(*gen.DO)
	return t
}
