package migrations

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
	"pgxs.io/chassis"
	cfg "pgxs.io/powerdoc/pkg/config"

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
	//config.LoadFromEnvFile()
	return nil
}
func Test_ImportData(t *testing.T) {
	err := Run()
	assert.NoError(t, err)
	if !chassis.EnvIsProd() {
		count := 0
		db, err := chassis.DB()
		assert.NoError(t, err)
		db.Table("users").Count(&count)
		assert.NotEmpty(t, count)
		assert.Equal(t, 1, count)
	}
}
