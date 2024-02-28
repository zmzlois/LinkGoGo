package monitor

import (
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
)

func EnhancedSentryEvent(ctx *fiber.Ctx, key string, value string) error {
	if hub := fibersentry.GetHubFromContext(ctx); hub != nil {
		hub.Scope().SetTag(key, value)
	}
	return ctx.Next()
}
