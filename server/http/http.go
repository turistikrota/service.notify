package http

import (
	"strings"
	"time"

	"github.com/cilloparch/cillop/helpers/http"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/server"
	"github.com/cilloparch/cillop/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/timeout"

	"github.com/turistikrota/service.notify/app"
	"github.com/turistikrota/service.notify/config"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	httpServer "github.com/turistikrota/service.shared/server/http"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/claim_guard"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_business"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/required_access"
)

type srv struct {
	config      config.App
	app         app.Application
	i18n        *i18np.I18n
	validator   validation.Validator
	tknSrv      token.Service
	sessionSrv  session.Service
	httpHeaders config.HttpHeaders
}

type Config struct {
	Env         config.App
	App         app.Application
	I18n        *i18np.I18n
	Validator   validation.Validator
	HttpHeaders config.HttpHeaders
	TokenSrv    token.Service
	SessionSrv  session.Service
}

func New(config Config) server.Server {
	return srv{
		config:      config.Env,
		app:         config.App,
		i18n:        config.I18n,
		validator:   config.Validator,
		tknSrv:      config.TokenSrv,
		sessionSrv:  config.SessionSrv,
		httpHeaders: config.HttpHeaders,
	}
}

func (h srv) Listen() error {
	return http.RunServer(http.Config{
		Host:        h.config.Http.Host,
		Port:        h.config.Http.Port,
		I18n:        h.i18n,
		Debug:       true,
		AcceptLangs: []string{},
		CreateHandler: func(router fiber.Router) fiber.Router {
			router.Use(h.cors(), h.deviceUUID())

			// business routes
			business := router.Group("/business", h.rateLimit(), h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess())
			business.Get("/", h.currentBusinessAccess(config.Roles.ActorConfig.Super, config.Roles.ActorConfig.View), h.wrapWithTimeout(h.GetBySelectedBusiness))
			business.Post("/config", h.currentBusinessAccess(config.Roles.ActorConfig.Super, config.Roles.ActorConfig.Add), h.wrapWithTimeout(h.BusinessConfigAdd))
			business.Put("/config", h.currentBusinessAccess(config.Roles.ActorConfig.Super, config.Roles.ActorConfig.Edit), h.wrapWithTimeout(h.BusinessConfigUpdate))
			business.Patch("/config", h.currentBusinessAccess(config.Roles.ActorConfig.Super, config.Roles.ActorConfig.Remove), h.wrapWithTimeout(h.BusinessConfigRemove))

			// user routes
			user := router.Group("/user", h.rateLimit(), h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess())
			user.Get("/", h.wrapWithTimeout(h.GetBySelectedUser))
			user.Post("/config", h.wrapWithTimeout(h.UserConfigAdd))
			user.Put("/config", h.wrapWithTimeout(h.UserConfigUpdate))
			user.Patch("/config", h.wrapWithTimeout(h.UserConfigRemove))

			// admin routes
			admin := router.Group("/admin", h.currentUserAccess(), h.requiredAccess())
			admin.Get("/", h.adminRoute(config.Roles.ActorConfig.Super, config.Roles.ActorConfig.List), h.wrapWithTimeout(h.AdminFilter))
			admin.Get("/:uuid", h.adminRoute(config.Roles.ActorConfig.Super, config.Roles.ActorConfig.View), h.wrapWithTimeout(h.AdminDetail))
			admin.Get("/business/:uuid", h.adminRoute(config.Roles.ActorConfig.Super, config.Roles.ActorConfig.View), h.wrapWithTimeout(h.AdminDetailByBusiness))
			admin.Get("/user/:name", h.adminRoute(config.Roles.ActorConfig.Super, config.Roles.ActorConfig.View), h.wrapWithTimeout(h.AdminDetailByUser))

			admin.Post("/test/mail", h.adminRoute(config.Roles.ActorConfig.Super, config.Roles.Notify.TestEmail), h.wrapWithTimeout(h.NotifyTestEmail))
			admin.Post("/test/sms", h.adminRoute(config.Roles.ActorConfig.Super, config.Roles.Notify.TestSms), h.wrapWithTimeout(h.NotifyTestSms))
			admin.Post("/test/telegram", h.adminRoute(config.Roles.ActorConfig.Super, config.Roles.Notify.TestTelegram), h.wrapWithTimeout(h.NotifyTestTelegram))

			return router
		},
	})
}

func (h srv) currentBusinessAccess(roles ...string) fiber.Handler {
	roles = append(roles, config.Roles.Business.Super)
	return current_business.New(current_business.Config{
		Roles: roles,
	})
}

func (h srv) currentAccountAccess() fiber.Handler {
	return current_account.New(current_account.Config{})
}

func (h srv) parseBody(c *fiber.Ctx, d interface{}) {
	http.ParseBody(c, h.validator, *h.i18n, d)
}

func (h srv) parseParams(c *fiber.Ctx, d interface{}) {
	http.ParseParams(c, h.validator, *h.i18n, d)
}

func (h srv) parseQuery(c *fiber.Ctx, d interface{}) {
	http.ParseQuery(c, h.validator, *h.i18n, d)
}

func (h srv) currentUserAccess() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
		I18n:       h.i18n,
		MsgKey:     Messages.Error.CurrentUserAccess,
		HeaderKey:  httpServer.Headers.Authorization,
		CookieKey:  auth.Cookies.AccessToken,
		UseCookie:  true,
		UseBearer:  true,
		IsRefresh:  false,
		IsAccess:   true,
	})
}

func (h srv) rateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        50,
		Expiration: 1 * time.Minute,
	})
}

func (h srv) deviceUUID() fiber.Handler {
	return device_uuid.New(device_uuid.Config{
		Domain: h.httpHeaders.Domain,
	})
}

func (h srv) requiredAccess() fiber.Handler {
	return required_access.New(required_access.Config{
		I18n:   *h.i18n,
		MsgKey: Messages.Error.RequiredAuth,
	})
}

func (h srv) adminRoute(extra ...string) fiber.Handler {
	claims := []string{config.Roles.Admin}
	if len(extra) > 0 {
		claims = append(claims, extra...)
	}
	return claim_guard.New(claim_guard.Config{
		Claims: claims,
		I18n:   *h.i18n,
		MsgKey: Messages.Error.AdminRoute,
	})
}

func (h srv) cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowMethods:     h.httpHeaders.AllowedMethods,
		AllowHeaders:     h.httpHeaders.AllowedHeaders,
		AllowCredentials: h.httpHeaders.AllowCredentials,
		AllowOriginsFunc: func(origin string) bool {
			origins := strings.Split(h.httpHeaders.AllowedOrigins, ",")
			for _, o := range origins {
				if strings.Contains(origin, o) {
					return true
				}
			}
			return false
		},
	})
}

func (h srv) wrapWithTimeout(fn fiber.Handler) fiber.Handler {
	return timeout.NewWithContext(fn, 10*time.Second)
}
