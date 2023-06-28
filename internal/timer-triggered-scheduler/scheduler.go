package timertriggeredtasks

import (
	"context"
	"time"

	"go.uber.org/zap"
)

// TimerTriggeredScheduler ...
type TimerTriggeredScheduler struct {
	logger       *zap.SugaredLogger
	ctx          context.Context
	paymentLogic PaymentLogic
	paymentSync  *time.Ticker
}

// PaymentLogic ...
type PaymentLogic interface {
	SyncAllPayments(ctx context.Context) error
}

// NewTimerTriggeredScheduler ...
func NewTimerTriggeredScheduler(logger *zap.SugaredLogger, paymentLogic PaymentLogic, paymentSync *time.Ticker) (TimerTriggeredScheduler, error) {
	return TimerTriggeredScheduler{
		ctx:          context.Background(),
		paymentLogic: paymentLogic,
		paymentSync:  paymentSync,
	}, nil
}

// Run ...
func (h *TimerTriggeredScheduler) Run() {
	go func() {
		for {
			select {
			case <-h.paymentSync.C:
				err := h.paymentLogic.SyncAllPayments(h.ctx)
				if err != nil {
					h.logger.Error(err)
				}
			case <-h.ctx.Done():
				return
			}
		}
	}()
}

// Stop
func (h *TimerTriggeredScheduler) Close() {
	h.ctx.Done()
}
