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

func (h srv) UserAddMail(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigAddMailCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigAddMail(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) UserAddSms(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigAddSmsCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigAddSms(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) UserAddTelegram(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigAddTelegramCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigAddTelegram(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) UserUpdateMail(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigUpdateMailCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigUpdateMail(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) UserUpdateSms(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigUpdateSmsCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigUpdateSms(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) UserUpdateTelegram(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigUpdateTelegramCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigUpdateTelegram(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) UserRemoveMail(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigRemoveMailCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigRemoveMail(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) UserRemoveSms(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigRemoveSmsCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigRemoveSms(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) UserRemoveTelegram(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigRemoveTelegramCmd{}
	h.parseBody(ctx, &cmd)
	cmd.ActorUUID = current_user.Parse(ctx).UUID
	cmd.ActorName = current_account.Parse(ctx).Name
	cmd.ActorType = actor_config.ActorTypeUser
	res, err := h.app.Commands.ActorConfigRemoveTelegram(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessAddMail(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigAddMailCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigAddMail(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessAddSms(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigAddSmsCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigAddSms(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessAddTelegram(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigAddTelegramCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigAddTelegram(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessUpdateMail(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigUpdateMailCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigUpdateMail(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessUpdateSms(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigUpdateSmsCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigUpdateSms(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessUpdateTelegram(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigUpdateTelegramCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigUpdateTelegram(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessRemoveMail(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigRemoveMailCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigRemoveMail(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessRemoveSms(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigRemoveSmsCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigRemoveSms(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessRemoveTelegram(ctx *fiber.Ctx) error {
	cmd := command.ActorConfigRemoveTelegramCmd{}
	h.parseBody(ctx, &cmd)
	business := current_business.Parse(ctx)
	cmd.ActorUUID = business.UUID
	cmd.ActorName = business.NickName
	cmd.ActorType = actor_config.ActorTypeBusiness
	res, err := h.app.Commands.ActorConfigRemoveTelegram(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
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
	res, err := h.app.Queries.ActorConfigFilter(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) AdminDetail(ctx *fiber.Ctx) error {
	query := query.ActorConfigGetByUUIDQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.ActorConfigGetByUUID(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) AdminDetailByUser(ctx *fiber.Ctx) error {
	query := query.ActorConfigGetByUserNameQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.ActorConfigGetByUserName(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) AdminDetailByBusiness(ctx *fiber.Ctx) error {
	query := query.ActorConfigGetByBusinessUUIDQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.ActorConfigGetByBusinessUUID(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) GetBySelectedUser(ctx *fiber.Ctx) error {
	query := query.ActorConfigGetByUserQuery{}
	query.UserName = current_account.Parse(ctx).Name
	query.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Queries.ActorConfigGetByUser(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) GetBySelectedBusiness(ctx *fiber.Ctx) error {
	query := query.ActorConfigGetByBusinessQuery{}
	business := current_business.Parse(ctx)
	query.BusinessName = business.NickName
	query.BusinessUUID = business.UUID
	res, err := h.app.Queries.ActorConfigGetByBusiness(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}
