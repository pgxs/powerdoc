package sql

import (
	"github.com/jinzhu/gorm"
	"pgxs.io/chassis"
	"sync"

	logx "pgxs.io/chassis/log"
	"pgxs.io/powerdoc/pkg/types"
)

type UserSqlMapper struct {
	log *logx.Entry
	db  *gorm.DB
}

var (
	userSqlMapper *UserSqlMapper
	usmInitOnce   sync.Once
)

func UserSqlMapperInstance() *UserSqlMapper {
	usmInitOnce.Do(func() {
		userSqlMapper = new(UserSqlMapper)
		userSqlMapper.log = logx.New().Category("mapper").Component("user")
		var err error
		userSqlMapper.db, err = chassis.DB()
		if err != nil {
			userSqlMapper.log.Error(err)
		}
	})
	return userSqlMapper
}
func (usm *UserSqlMapper) FindByLoginName(loginName string) *types.UserDO {

	usm.log.Debugf("Find user by name: [%s]", loginName)

	var user types.UserDO
	err := usm.db.Where("username = ?", loginName).First(&user).Error
	if err != nil {
		usm.log.Error(err)
	}
	return &user
}
