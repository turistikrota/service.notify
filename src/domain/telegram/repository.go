package telegram

import (
	"context"

	"github.com/mixarchitecture/i18np"
)

type SendConfig struct {
	Recipient string
	Data      *Data
}

type Repository interface {
	Send(ctx context.Context, config SendConfig) *i18np.Error
}
