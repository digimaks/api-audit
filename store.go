// SPDX-License-Identifier: EUPL-1.2

package api

import (
	"context"

	"azugo.io/azugo"
	"azugo.io/core/instrumenter"
	"github.com/lx-lib/lx-go-jsondb"
	"go.uber.org/zap"
)

// storeInstrumenter is a instrumenter to pass session data to database.
func (a *App) storeInstrumenter() instrumenter.Instrumenter {
	return instrumenter.Instrumenter(func(ctx context.Context, op string, _ ...interface{}) func(err error) {
		if op != jsondb.InstrumentationExec {
			return func(_ error) {}
		}

		tctx, tx := jsondb.FetchTxCtx(ctx)
		if rctx, ok := tctx.(*azugo.Context); ok && tx != nil {
			if rctx.User().Authorized() {
				ip := rctx.IP()
				if ipv4 := ip.To4(); ipv4 != nil {
					ip = ipv4
				}

				_, err := tx.Exec(ctx, `CALL "session"."create"($1, $2, $3, $4, $5, $6, $7)`,
					rctx.User().ClaimValue("sid"),
					rctx.User().ID(),
					rctx.User().ClaimValue("code"),
					rctx.User().GivenName(),
					rctx.User().FamilyName(),
					rctx.User().ClaimValue("role"),
					ip,
				)
				if err != nil {
					a.Log().Error("failed to set database session context", zap.Error(err))
				}
			}
		}

		return func(_ error) {}
	})
}
