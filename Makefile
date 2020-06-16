test:
	POWERDOC_CONF=${PWD}/conf/app.yml go test -v ./... -count=1
migrate:
	@echo "Run migrations"
	@cd scripts/sql/migrations && go-bindata -pkg migrations -o ../../../pkg/migrations/bindata.go . && cd ../../../ \
	&& POWERDOC_CONF=$(PWD)/conf/app.sqlite.yml \
	PG_TEST_DATA_FILE=$(PWD)/scripts/sql/test/data.sql \
	go test -v ./pkg/migrations/. -count=1