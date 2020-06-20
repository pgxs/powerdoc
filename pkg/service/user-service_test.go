package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"pgxs.io/powerdoc/pkg/config"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	os.Exit(code)
}
func setup() {
	config.Load()
	UserServiceInstance()
}

func TestUserService_UserSignIn(t *testing.T) {
	token, err := userService.SignIn("admin", "powerdoc")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
