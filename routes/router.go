// SPDX-License-Identifier: EUPL-1.2

package routes

import (
	api "github.com/digimaks/api-audit"
	"github.com/digimaks/api-audit/openapi"

	"azugo.io/azugo"
	"github.com/digimaks/go-idauth"
	oa "github.com/lx-lib/go-openapi"
	"golang.org/x/text/language"
)

type router struct {
	*api.App
	openapi *oa.OpenAPI
	lang    language.Matcher
}

func Init(app *api.App) error {
	r := &router{
		App: app,
		// Supported languages.
		lang: language.NewMatcher([]language.Tag{
			language.Latvian, // The first language is used as fallback.
			language.English,
		}),
	}

	if app.Env().IsDevelopment() {
		r.openapi = oa.NewDefaultOpenAPIHandler(openapi.OpenAPIDefinition, app.App)
	}

	app.Get("/healthz", r.healthz)

	v1 := app.Group("/1.0")
	v1.Use(idauth.Authentication(app.App, app.Config().Auth))

	v1.Post("/audit", r.audit)

	return nil
}

// Language returns the requested language.
func (r *router) Language(ctx *azugo.Context) language.Tag {
	langs, _, _ := language.ParseAcceptLanguage(ctx.Header.Get("Accept-Language"))
	lang, _, _ := r.lang.Match(langs...)

	return lang
}
