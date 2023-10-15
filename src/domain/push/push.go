package push

import "github.com/turistikrota/service.notify/src/domain/notify"

type Entity struct {
	Token        string        `json:"token"`
	Data         *Data         `json:"data"`
	Notification *Notification `json:"notification"`
	Notify       *notify.Notify
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Image string `json:"image"`
}

type Data struct {
	Time  string `json:"time"`
	Score string `json:"score"`
}
