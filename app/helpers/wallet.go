package helper

import (
	"fmt"
	"pay-bills/app/enums"
	"pay-bills/app/models"
	"pay-bills/database"
	"pay-bills/database/query"
)

func ProcessWalletDebit(
	wallet *model.Wallet,
	service *model.Service,
	amount float64,
	charge float64,
	reference string,
	info *string,
	glAmount float64,
	allowNegative bool,
) error {
	var err error

	if err = wallet.AllowsTransaction(amount, true, false); err != nil {
		return err
	}

	// Debit amount
	err = DebitWallet(wallet, service, amount, enums.TRANSACTION, reference, info, glAmount)
	if err != nil {
		return err
	}

	chargeRef := NewWalletTnxRef()
	chargeInfo := fmt.Sprintf("Charge on tnx: %s", reference)

	// Debit charge
	err = DebitWallet(wallet, service, charge, enums.CHARGE, chargeRef, &chargeInfo, charge)
	if err != nil {
		return err
	}

	return err
}

// DebitWallet - Debit the wallet for the service.
func DebitWallet(
	wallet *model.Wallet,
	service *model.Service,
	amount float64,
	transType enums.TransType,
	reference string,
	info *string,
	glAmount float64,
) error {
	dbTxErr := query.Use(database.Db).Transaction(func(tx *query.Query) error {
		var err error
		prevBalance := wallet.Balance
		wallet.Balance -= amount
		_, err = tx.Wallet.Where(tx.Wallet.ID.Eq(wallet.ID)).Update(query.Wallet.Balance, wallet.Balance)

		if err != nil {
			return err
		}

		err = tx.WalletTransaction.Create(&model.WalletTransaction{
			WalletId:    wallet.ID,
			ProductId:   service.ID,
			Type:        transType,
			Amount:      amount,
			PrevBalance: prevBalance,
			NewBalance:  wallet.Balance,
			Status:      enums.SUCCESSFUL,
			Action:      enums.DEBIT,
			Info:        info,
			Reference:   reference,
		})

		if err != nil {
			return err
		}

		// TODO: Impact service ledger.

		return nil
	})

	return dbTxErr
}

// CreditWallet - Credit the wallet for the service.
func CreditWallet(
	wallet *model.Wallet,
	service *model.Service,
	amount float64,
	transType enums.TransType,
	reference string,
	info *string,
) error {
	dbTxErr := query.Use(database.Db).Transaction(func(tx *query.Query) error {
		var err error
		prevBalance := wallet.Balance
		wallet.Balance += amount
		_, err = tx.Wallet.Update(query.Wallet.Balance, wallet.Balance)

		if err != nil {
			return err
		}

		err = tx.WalletTransaction.Create(&model.WalletTransaction{
			WalletId:    wallet.ID,
			ProductId:   service.ID,
			Type:        transType,
			Amount:      amount,
			PrevBalance: prevBalance,
			NewBalance:  wallet.Balance,
			Status:      enums.SUCCESSFUL,
			Action:      enums.CREDIT,
			Info:        info,
			Reference:   reference,
		})

		if err != nil {
			return err
		}

		// TODO: Impact service ledger.

		return nil
	})

	return dbTxErr
}
