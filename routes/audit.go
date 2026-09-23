// SPDX-License-Identifier: EUPL-1.2

package routes

import (
	"azugo.io/azugo"
	"github.com/nobid-lsp-latvia/go-audit"
)

// @operationId AuditPersonData
// @title Audit person data
// @description Audits person data access.
// @param AuditRequest body audit.AuditRequest true "Audit data"
// @success 200 {empty} "Audit response"
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 404 {empty} "Not Found"
// @failure 500 string string "Internal server error"
// @resource Audit
// @route /1.0/audit [post].
func (r *router) audit(ctx *azugo.Context) {
	auditReq := &audit.AuditRequest{}
	if err := ctx.Body.JSON(auditReq); err != nil {
		ctx.Error(err)

		return
	}

	if err := r.Store().Exec(ctx, "audit.audit_person_data", auditReq, nil); err != nil {
		ctx.Error(err)
	}
}
