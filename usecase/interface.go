package usecase

import "context"

type Transaction interface {
	RunTxn(ctx context.Context, txFunc func(ctx context.Context) error) error
}
