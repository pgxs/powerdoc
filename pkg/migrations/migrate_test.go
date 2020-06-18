package migrations

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"pgxs.io/chassis"
	cfg "pgxs.io/powerdoc/pkg/config"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(-1)
	}
	exitVal := m.Run()

	//tearDown
	os.Exit(exitVal)
}
func setup() error {
	cfg.Load()
	//
	//if config.Databases()[0].Dialect == "postgres" {
	//	args, err := pq.ParseURL(config.Databases()[0].DSN)
	//	if err != nil {
	//		panic(err)
	//		os.Exit(-1)
	//	}
	//	config.Databases()[0].DSN = args
	//}
	//config.LoadFromEnvFile()
	return nil
}
func Test_ImportData(t *testing.T) {
	err := Run()
	assert.NoError(t, err)
	count := 0
	db, err := chassis.DB()
	assert.NoError(t, err)
	db.Table("users").Count(&count)
	assert.NotEmpty(t, count)
	assert.Equal(t, 1, count)
}
