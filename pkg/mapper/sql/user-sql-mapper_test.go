package sql

import (
	"github.com/stretchr/testify/assert"
	"os"
	"pgxs.io/powerdoc/pkg/config"
	"testing"
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
