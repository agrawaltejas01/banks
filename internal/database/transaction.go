package database

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type CtxKey string

const (
	// TransactionKey is the key used to store the transaction in the context
	TransactionKey CtxKey = "transaction"
)

func StartTransaction(ctx context.Context) context.Context {
	// Start a new transaction
	tx := DB.Begin()
	if tx.Error != nil {
		return nil // Handle error appropriately in your application
	}
	fmt.Print(tx)
	ctx = context.WithValue(ctx, TransactionKey, tx)
	return ctx
}

func CommitTransaction(ctx context.Context) error {

	// Retrieve the transaction from the context
	tx, ok := ctx.Value(TransactionKey).(*gorm.DB)
	if !ok || tx == nil {
		return errors.New("no transaction found in context")
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func RollbackTransaction(ctx context.Context) error {

	// Retrieve the transaction from the context
	tx, ok := ctx.Value("transaction").(*gorm.DB)
	if !ok || tx == nil {
		return errors.New("no transaction found in context")
	}

	// Rollback the transaction
	if err := tx.Rollback().Error; err != nil {
		return err
	}
	return nil
}

func GetDbInstanceFromContextOrDB(ctx context.Context) *gorm.DB {
	// Retrieve the transaction from the context
	tx, ok := ctx.Value(TransactionKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}
	return DB // Fallback to the global DB instance if no transaction is found
}
