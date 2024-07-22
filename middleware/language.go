package middleware

import (
	"github.com/labstack/echo/v4"
	i18n2 "github.com/mauriciomartinezc/real-estate-mc-auth/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func LanguageHandler() echo.MiddlewareFunc {
	bundle := i18n2.NewLocalization()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			supportedLangs := []language.Tag{
				language.English, // "en"
				language.Spanish, // "es"
			}
			lang := c.Request().Header.Get("Accept-Language")
			matcher := language.NewMatcher(supportedLangs)
			bestMatch, _ := language.MatchStrings(matcher, lang)

			localize := i18n.NewLocalizer(bundle, bestMatch.String())

			c.Set("localize", localize)

			return next(c)
		}
	}
}
