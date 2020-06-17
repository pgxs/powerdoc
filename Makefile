.PHONY: migratea bindata migrate-mysql
test:
	POWERDOC_CONF=${PWD}/conf/app.yml go test -v ./... -count=1

migrate-all: migrate-sqlite migrate-mysql

migrate-sqlite: bindata sqlite-migrations
	@echo "Run migrations to sqlite"
sqlite-migrations:
	@POWERDOC_CONF=$(PWD)/conf/app.sqlite.yml \
	PG_TEST_DATA_FILE=$(PWD)/scripts/sql/test/data.sql \
	go test -v ./pkg/migrations/. -count=1

migrate-mysql: bindata mysql-migrations
bindata:
	@echo "bindata"
	@cd scripts/sql/migrations && go-bindata -pkg migrations -o ../../../pkg/migrations/bindata.go .
mysql-migrations:
	@POWERDOC_CONF=$(PWD)/conf/app.yml \
	PG_TEST_DATA_FILE=$(PWD)/scripts/sql/test/data.sql \
	go test -v ./pkg/migrations/. -count=1