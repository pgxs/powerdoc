package api

import (
	"pgxs.io/powerdoc/pkg/types"
)

//UserMapper user db map interface
type UserMapper interface {
	FindByLoginName(loginName string) *types.UserDO
	FindUsersByKeywords(keywords string, pageIndex, pageSize int)
}
