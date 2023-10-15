package http

import (
	"github.com/gofiber/fiber/v2"
	httpI18n "github.com/mixarchitecture/microp/server/http/i18n"
	"github.com/mixarchitecture/microp/server/http/result"
	"github.com/turistikrota/service.notify/src/delivery/http/dto"
)

func (h Server) GetByUUID(ctx *fiber.Ctx) error {
	d := dto.Request.GetByUUID()
	h.parseParams(ctx, d)
	l, a := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	res, err := h.app.Queries.GetByUUID.Handle(ctx.UserContext(), d.ToQuery())
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(h.i18n.Translate(Messages.Success.GetByUUID, l, a), dto.Response.Detail(&res.Notify))
}

func (h Server) GetAllByRecipient(ctx *fiber.Ctx) error {
	d := dto.Request.GetAllByRecipient()
	h.parseQuery(ctx, d)
	l, a := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	err := d.Default()
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	res, err := h.app.Queries.GetAllByRecipient.Handle(ctx.UserContext(), d.ToQuery())
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(h.i18n.Translate(Messages.Success.GetAllByRecipient, l, a), dto.Response.GetAllByRecipient(res))
}

func (h Server) GetAllByChannel(ctx *fiber.Ctx) error {
	d := dto.Request.GetAllByChannel()
	h.parseQuery(ctx, d)
	err := d.Default()
	l, a := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	res, err := h.app.Queries.GetAllByChannel.Handle(ctx.UserContext(), d.ToQuery())
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(h.i18n.Translate(Messages.Success.GetAllByChannel, l, a), dto.Response.GetAllByChannel(res))
}
