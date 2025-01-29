package service

import (
	"context"
)

type ServiceRunner interface {
	Run(ctx context.Context)
}
