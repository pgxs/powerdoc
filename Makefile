.PHONY: migratea bindata migrate-mysql
test:
	POWERDOC_CONF=${PWD}/conf/app.yml go test -v ./... -count=1

migrate-all: migrate-sqlite migrate-mysql migrate-postgres
	@ehco "Run all migrations"
migrate-sqlite: sqlite-bindata sqlite-migrations
	@echo "Run migrations to sqlite"
sqlite-bindata:
	@echo "sqlite-bindata"
	@cd scripts/sql/migrations/sqlite && go-bindata -pkg sqliteMigrations -o ../../../../pkg/migrations/sqlite/bindata.go .
sqlite-migrations:
	@POWERDOC_CONF=$(PWD)/conf/app.sqlite.yml \
	PG_TEST_DATA_FILE=$(PWD)/scripts/sql/test/data.sql \
	go test -v ./pkg/migrations/. -count=1

migrate-mysql: mysql-bindata mysql-migrations
mysql-bindata:
	@echo "mysql-bindata"
	@cd scripts/sql/migrations/mysql && go-bindata -pkg mysqlMigrations -o ../../../../pkg/migrations/mysql/bindata.go .
mysql-migrations:
	@POWERDOC_CONF=$(PWD)/conf/app.yml \
	PG_TEST_DATA_FILE=$(PWD)/scripts/sql/test/data.sql \
	go test -v ./pkg/migrations/. -count=1
migrate-postgres: postgres-bindata postgre-migrations
postgres-bindata:
	@echo "bindata"
	@cd scripts/sql/migrations/pg && go-bindata -pkg pgMigrations -o ../../../../pkg/migrations/pg/bindata.go .
postgre-migrations:
	@POWERDOC_CONF=$(PWD)/conf/app.pg.yml \
     	PG_TEST_DATA_FILE=$(PWD)/scripts/sql/test/data.sql \
     	go test -v ./pkg/migrations/. -count=1