package sql

import (
	"os"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"

	"pgxs.io/powerdoc/pkg/config"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
func setup() {
	config.Load()
}

func TestUserSqlMapper_FindByLoginName(t *testing.T) {
	user := UserSqlMapperInstance().FindByLoginName("admin")
	assert.Equal(t, "测试", user.Name)
}
