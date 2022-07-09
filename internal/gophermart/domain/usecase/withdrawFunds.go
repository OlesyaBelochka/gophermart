package usecase

import (
	"context"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	WithdrawUserFundsRepository interface {
		FindAccountByID(context.Context, string) (core.Account, error)
		SaveAccount(context.Context, core.Account) error
	}

	WithdrawFundsInputPort interface {
		Execute(context.Context, *sharedkernel.User, WithdrawUserFundsInputDTO) error
	}

	// WithdrawUserFundsInputDTO Example of DTO with json at usecase level, which not quite correct.
	WithdrawUserFundsInputDTO struct {
		Order int                `json:"order"`
		Sum   sharedkernel.Money `json:"sum"`
	}

	WithdrawUserFunds struct {
		Repo WithdrawUserFundsRepository
	}
)

func NewWithdrawUserFunds(repo WithdrawUserFundsRepository) *WithdrawUserFunds {
	return &WithdrawUserFunds{
		Repo: repo,
	}
}

func (w *WithdrawUserFunds) Execute(ctx context.Context, user *sharedkernel.User, dto WithdrawUserFundsInputDTO) error { // 5
	account, err := w.Repo.FindAccountByID(ctx, user.ID())
	if err != nil {
		return err // nolint:wrapcheck // ok
	}
	// do work with account
	err = account.WithdrawPoints(dto.Order, dto.Sum)
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	err = w.Repo.SaveAccount(ctx, account)

	return nil
}

// 5
