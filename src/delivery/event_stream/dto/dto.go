package dto

type Dto interface {
	NotifyMail() *NotifyMail
	NotifySMS() *NotifySMS
	NotifyTelegram() *NotifyTelegram
}

type dto struct{}

func New() Dto {
	return &dto{}
}

func (d *dto) NotifyMail() *NotifyMail {
	return &NotifyMail{}
}

func (d *dto) NotifySMS() *NotifySMS {
	return &NotifySMS{}
}

func (d *dto) NotifyTelegram() *NotifyTelegram {
	return &NotifyTelegram{}
}
