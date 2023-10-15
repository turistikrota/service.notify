package notify

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
)

type ListConfig struct {
	Offset    int64
	Limit     int64
	StartDate time.Time
	EndDate   time.Time
}

type ListResult struct {
	List          []DetailResult
	Total         int64
	FilteredTotal int64
	IsNext        bool
	IsPrev        bool
}

type DetailResult struct {
	UUID      string
	Type      Channel
	Recipient string
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      interface{}
}

type Repository interface {
	Log(ctx context.Context, notify *Notify, data interface{}) *i18np.Error
	GetByUUID(ctx context.Context, uuid string) (*DetailResult, *i18np.Error)
	GetAllByRecipient(ctx context.Context, recipient string, config ListConfig) (*ListResult, *i18np.Error)
	GetAllByChannel(ctx context.Context, ch Channel, config ListConfig) (*ListResult, *i18np.Error)
}
