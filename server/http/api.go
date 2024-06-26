package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.notify/app/command"
	"github.com/turistikrota/service.notify/app/query"
	"github.com/turistikrota/service.notify/domains/actor_config"
	"github.com/turistikrota/service.notify/pkg/utils"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_business"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h srv) UserConfigAdd(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigAddCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigAdd(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) UserConfigUpdate(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigUpdateCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigUpdate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) UserConfigRemove(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigRemoveCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigRemove(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessConfigAdd(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigAddCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigAdd(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessConfigUpdate(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigUpdateCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigUpdate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessConfigRemove(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigRemoveCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigRemove(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AdminFilter(ctx *fiber.Ctx) error {
	pagi := utils.Pagination{}
	h.parseQuery(ctx, &pagi)
	filter := actor_config.FilterEntity{}
	h.parseQuery(ctx, &filter)
	query := query.ActorConfigFilterQuery{
		FilterEntity: &filter,
		Pagination:   &pagi,
	}
	res, err := h.app.Queries.ActorConfigFilter(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) AdminDetail(ctx *fiber.Ctx) error {
	query := query.ActorConfigGetByUUIDQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.ActorConfigGetByUUID(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) AdminDetailByUser(ctx *fiber.Ctx) error {
	query := query.ActorConfigGetByUserNameQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.ActorConfigGetByUserName(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) AdminDetailByBusiness(ctx *fiber.Ctx) error {
	query := query.ActorConfigGetByBusinessUUIDQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.ActorConfigGetByBusinessUUID(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) GetBySelectedUser(ctx *fiber.Ctx) error {
	query := query.ActorConfigGetOrCreateByUserQuery{}
	query.UserName = current_account.Parse(ctx).Name
	query.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Queries.ActorConfigGetOrCreateByUser(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) GetBySelectedBusiness(ctx *fiber.Ctx) error {
	query := query.ActorConfigGetOrCreateByBusinessQuery{}
	business := current_business.Parse(ctx)
	query.BusinessName = business.NickName
	query.BusinessUUID = business.UUID
	res, err := h.app.Queries.ActorConfigGetOrCreateByBusiness(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) NotifyTestEmail(ctx *fiber.Ctx) error {
	cmd := command.NotifyTestMailCmd{}
	h.parseBody(ctx, &cmd)
	l, a := i18n.ParseLocales(ctx)
	cmd.Locale = l
	res, err := h.app.Commands.NotifyTestMail(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) NotifyTestSms(ctx *fiber.Ctx) error {
	cmd := command.NotifyTestSmsCmd{}
	h.parseBody(ctx, &cmd)
	l, a := i18n.ParseLocales(ctx)
	cmd.Locale = l
	res, err := h.app.Commands.NotifyTestSms(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) NotifyTestTelegram(ctx *fiber.Ctx) error {
	cmd := command.NotifyTestTelegramCmd{}
	h.parseBody(ctx, &cmd)
	l, a := i18n.ParseLocales(ctx)
	cmd.Locale = l
	res, err := h.app.Commands.NotifyTestTelegram(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}
