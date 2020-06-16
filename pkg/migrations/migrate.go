package migrations

import (
	"sync"

	"pgxs.io/chassis"
	"pgxs.io/chassis/config"
)

var migrateOnce sync.Once

//Run start migrations
func Run() (err error) {
	migrateOnce.Do(func() {
		//数据库版本管理
		err = chassis.Migrate(AssetNames(), func(name string) ([]byte, error) {
			return Asset(name)
		}, config.Databases()[0].DSN, config.Databases()[0].Dialect)
	})
	return
}
