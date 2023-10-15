package req

import (
	"time"

	"github.com/mixarchitecture/i18np"
)

type ListRequest struct {
	StartDate *time.Time `query:"start_date" validate:"omitempty"`
	EndDate   *time.Time `query:"end_date" validate:"omitempty"`
}

func (r *ListRequest) validateDatetime() bool {
	return r.StartDate.Before(*r.EndDate) && r.EndDate.Sub(*r.StartDate).Hours() <= 24*14
}

func (r *ListRequest) validate() *i18np.Error {
	if !r.validateDatetime() {
		return i18np.NewError("invalid_datetime")
	}
	return nil
}

func (r *ListRequest) Default() *i18np.Error {
	if r.StartDate == nil {
		r.StartDate = new(time.Time)
		*r.StartDate = time.Now().AddDate(0, 0, -7)
	}
	if r.EndDate == nil {
		r.EndDate = new(time.Time)
		*r.EndDate = time.Now()
	}
	if err := r.validate(); err != nil {
		return err
	}
	return nil
}

func (r *request) ListRequest() *ListRequest {
	return &ListRequest{}
}
