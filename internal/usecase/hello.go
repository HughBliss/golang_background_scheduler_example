package usecase

import (
	"context"
	"time"

	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/logger"
)

func NewHelloUsecase(log logger.Logger) HelloUsecase {
	log, _ = log.Named(context.Background(), "hello_usecase")
	defer log.End()

	log.Info("Initializing hello usecase")

	return &helloUsecase{
		log: log,
	}
}

type HelloUsecase interface {
	SayHello(ctx context.Context, t time.Time) error
}

type helloUsecase struct {
	log logger.Logger
}

// SayHello implements HelloUsecase.
func (h *helloUsecase) SayHello(ctx context.Context, t time.Time) error {
	log, _ := h.log.Named(ctx, "say_hello")
	defer log.End()


	log.Info("Hello, world!", logger.Time("time", t))


	return nil
}
