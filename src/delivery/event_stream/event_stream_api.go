package event_stream

import (
	"encoding/json"
)

func (s Server) ListenSendNotifyEmail(data []byte) {
	d := s.dto.NotifyMail()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	s.app.Commands.SendMail.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenSendNotifySMS(data []byte) {
	d := s.dto.NotifySMS()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	s.app.Commands.SendSms.Handle(s.ctx, d.ToCommand())
}

func (s Server) ListenSendNotifyTelegram(data []byte) {
	d := s.dto.NotifyTelegram()
	err := json.Unmarshal(data, &d)
	if err != nil {
		return
	}
	s.app.Commands.SendTelegram.Handle(s.ctx, d.ToCommand())
}
