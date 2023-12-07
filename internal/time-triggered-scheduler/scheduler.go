package timetriggeredtasks

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type TimeTriggeredScheduler struct {
	logger      *zap.SugaredLogger
	ctx         context.Context
	paymentSvc  PaymentService
	paymentSync *time.Ticker
}

type PaymentService interface {
	SyncAllPayments(ctx context.Context) error
}

func NewTimeTriggeredScheduler(logger *zap.SugaredLogger, paymentService PaymentService, paymentSync *time.Ticker) (TimeTriggeredScheduler, error) {
	return TimeTriggeredScheduler{
		ctx:         context.Background(),
		paymentSvc:  paymentService,
		paymentSync: paymentSync,
	}, nil
}

func (h *TimeTriggeredScheduler) Run() {
	go func() {
		for {
			select {
			case <-h.paymentSync.C:
				err := h.paymentSvc.SyncAllPayments(h.ctx)
				if err != nil {
					h.logger.Error(err)
				}
			case <-h.ctx.Done():
				return
			}
		}
	}()
}

func (h *TimeTriggeredScheduler) Close() {
	h.ctx.Done()
}
