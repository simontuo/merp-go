package subscriber

import (
	"context"

	"github.com/prometheus/common/log"
)

type UserService struct {
}

func (s *UserService) Handle(ctx context.Context) error {
	log.Info("Handler Received message: ")
	return nil
}

func Handler(ctx context.Context) error {
	log.Info("Function Received message: ")
	return nil
}
