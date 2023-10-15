package firebase

import (
	"context"
	"log"

	"firebase.google.com/go/messaging"
	"github.com/mixarchitecture/i18np"
	"github.com/sirupsen/logrus"
	"github.com/turistikrota/service.notify/src/domain/push"
)

func (r *repo) Send(ctx context.Context, entity *push.Entity, token string) *i18np.Error {
	message := &messaging.Message{
		Data: map[string]string{
			"score": entity.Data.Score,
			"time":  entity.Data.Time,
		},
		Notification: &messaging.Notification{
			ImageURL: entity.Notification.Image,
			Title:    entity.Notification.Title,
			Body:     entity.Notification.Body,
		},
		Token: token,
	}
	_, err := r.client.Send(ctx, message)
	if err != nil {
		logrus.Error(err)
		return r.factory.Errors.Failed(err.Error())
	}
	return nil
}

func (r *repo) SendAll(ctx context.Context, entity *push.Entity, tokens []string) *i18np.Error {
	message := &messaging.MulticastMessage{
		Data: map[string]string{
			"score": entity.Data.Score,
			"time":  entity.Data.Time,
		},
		Notification: &messaging.Notification{
			ImageURL: entity.Notification.Image,
			Title:    entity.Notification.Title,
			Body:     entity.Notification.Body,
		},
		Tokens: tokens,
	}

	br, err := r.client.SendMulticast(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	if br.FailureCount > 0 {
		logrus.Error(br)
		return r.factory.Errors.Failed("failed to send push")
	}
	return nil
}
