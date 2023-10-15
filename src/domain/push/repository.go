package push

import (
	"context"

	"github.com/mixarchitecture/i18np"
)

type Repository interface {
	Send(ctx context.Context, e *Entity, token string) *i18np.Error
	SendAll(ctx context.Context, e *Entity, tokens []string) *i18np.Error
}
