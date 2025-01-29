package service

import (
	"context"
)

type Service interface {
	Run(ctx context.Context) error
	Stop(ctx context.Context)
}
