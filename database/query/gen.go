// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                 = new(Query)
	Fee               *fee
	GeneralLedger     *generalLedger
	GlTransaction     *glTransaction
	KycLevel          *kycLevel
	Loan              *loan
	Service           *service
	ServiceProvider   *serviceProvider
	Terminal          *terminal
	TerminalGroup     *terminalGroup
	Transactions      *transactions
	User              *user
	Wallet            *wallet
	WalletTransaction *walletTransaction
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Fee = &Q.Fee
	GeneralLedger = &Q.GeneralLedger
	GlTransaction = &Q.GlTransaction
	KycLevel = &Q.KycLevel
	Loan = &Q.Loan
	Service = &Q.Service
	ServiceProvider = &Q.ServiceProvider
	Terminal = &Q.Terminal
	TerminalGroup = &Q.TerminalGroup
	Transactions = &Q.Transactions
	User = &Q.User
	Wallet = &Q.Wallet
	WalletTransaction = &Q.WalletTransaction
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                db,
		Fee:               newFee(db, opts...),
		GeneralLedger:     newGeneralLedger(db, opts...),
		GlTransaction:     newGlTransaction(db, opts...),
		KycLevel:          newKycLevel(db, opts...),
		Loan:              newLoan(db, opts...),
		Service:           newService(db, opts...),
		ServiceProvider:   newServiceProvider(db, opts...),
		Terminal:          newTerminal(db, opts...),
		TerminalGroup:     newTerminalGroup(db, opts...),
		Transactions:      newTransactions(db, opts...),
		User:              newUser(db, opts...),
		Wallet:            newWallet(db, opts...),
		WalletTransaction: newWalletTransaction(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Fee               fee
	GeneralLedger     generalLedger
	GlTransaction     glTransaction
	KycLevel          kycLevel
	Loan              loan
	Service           service
	ServiceProvider   serviceProvider
	Terminal          terminal
	TerminalGroup     terminalGroup
	Transactions      transactions
	User              user
	Wallet            wallet
	WalletTransaction walletTransaction
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		Fee:               q.Fee.clone(db),
		GeneralLedger:     q.GeneralLedger.clone(db),
		GlTransaction:     q.GlTransaction.clone(db),
		KycLevel:          q.KycLevel.clone(db),
		Loan:              q.Loan.clone(db),
		Service:           q.Service.clone(db),
		ServiceProvider:   q.ServiceProvider.clone(db),
		Terminal:          q.Terminal.clone(db),
		TerminalGroup:     q.TerminalGroup.clone(db),
		Transactions:      q.Transactions.clone(db),
		User:              q.User.clone(db),
		Wallet:            q.Wallet.clone(db),
		WalletTransaction: q.WalletTransaction.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		Fee:               q.Fee.replaceDB(db),
		GeneralLedger:     q.GeneralLedger.replaceDB(db),
		GlTransaction:     q.GlTransaction.replaceDB(db),
		KycLevel:          q.KycLevel.replaceDB(db),
		Loan:              q.Loan.replaceDB(db),
		Service:           q.Service.replaceDB(db),
		ServiceProvider:   q.ServiceProvider.replaceDB(db),
		Terminal:          q.Terminal.replaceDB(db),
		TerminalGroup:     q.TerminalGroup.replaceDB(db),
		Transactions:      q.Transactions.replaceDB(db),
		User:              q.User.replaceDB(db),
		Wallet:            q.Wallet.replaceDB(db),
		WalletTransaction: q.WalletTransaction.replaceDB(db),
	}
}

type queryCtx struct {
	Fee               IFeeDo
	GeneralLedger     IGeneralLedgerDo
	GlTransaction     IGlTransactionDo
	KycLevel          IKycLevelDo
	Loan              ILoanDo
	Service           IServiceDo
	ServiceProvider   IServiceProviderDo
	Terminal          ITerminalDo
	TerminalGroup     ITerminalGroupDo
	Transactions      ITransactionsDo
	User              IUserDo
	Wallet            IWalletDo
	WalletTransaction IWalletTransactionDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Fee:               q.Fee.WithContext(ctx),
		GeneralLedger:     q.GeneralLedger.WithContext(ctx),
		GlTransaction:     q.GlTransaction.WithContext(ctx),
		KycLevel:          q.KycLevel.WithContext(ctx),
		Loan:              q.Loan.WithContext(ctx),
		Service:           q.Service.WithContext(ctx),
		ServiceProvider:   q.ServiceProvider.WithContext(ctx),
		Terminal:          q.Terminal.WithContext(ctx),
		TerminalGroup:     q.TerminalGroup.WithContext(ctx),
		Transactions:      q.Transactions.WithContext(ctx),
		User:              q.User.WithContext(ctx),
		Wallet:            q.Wallet.WithContext(ctx),
		WalletTransaction: q.WalletTransaction.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}