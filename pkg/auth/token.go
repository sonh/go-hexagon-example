package auth

import (
	"context"

	"hexagon-example/core/valueobj"
)

type Token struct {
}

func (token Token) OrgID() string {
	return "randomToken"
}

func GetOrgIDFromCtx(ctx context.Context) valueobj.HasOrgID {
	return Token{}
}
