package migrations

import (
	"pgxs.io/chassis"
	"pgxs.io/chassis/config"
	mysqlMigrations "pgxs.io/powerdoc/pkg/migrations/mysql"
	pgMigrations "pgxs.io/powerdoc/pkg/migrations/pg"
	sqliteMigrations "pgxs.io/powerdoc/pkg/migrations/sqlite"
	"strings"
	"sync"
)

var migrateOnce sync.Once

//Run start migrations
func Run() (err error) {

	dbURL := config.Databases()[0].DSN
	if strings.HasPrefix(dbURL, "postgres://") {
		dbURL = dbURL[11:] //remove prefix
	}
	migrateOnce.Do(func() {

		switch config.Databases()[0].Dialect {
		case "postgres":
			//数据库版本管理
			err = chassis.Migrate(pgMigrations.AssetNames(), func(name string) ([]byte, error) {
				return pgMigrations.Asset(name)
			}, dbURL, config.Database().Dialect)
		case "mysql":
			err = chassis.Migrate(mysqlMigrations.AssetNames(), func(name string) ([]byte, error) {
				return mysqlMigrations.Asset(name)
			}, dbURL, config.Database().Dialect)
		case "sqlite3":
			err = chassis.Migrate(sqliteMigrations.AssetNames(), func(name string) ([]byte, error) {
				return sqliteMigrations.Asset(name)
			}, dbURL, config.Database().Dialect)
		}

	})
	return
}
