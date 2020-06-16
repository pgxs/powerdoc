package sql

import (
	logx "pgxs.io/chassis/log"
	"pgxs.io/powerdoc/pkg/types"
)

type UserSqlMapper struct {
	log logx.Entry
}

func (usm *UserSqlMapper) FindByLoginName(loginName string) *types.UserDO {
	return nil
}
