package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	log "gopkg.in/logger.v1"
)

func init() {
	log.SetOutputLevel(0)
}
func TestHashPassword(t *testing.T) {
	hashPwd, ok := HashHelper().HashWithPrefix("123456", SHA1)
	assert.True(t, ok)
	assert.NotEmpty(t, hashPwd)
	assert.Equal(t, "{sha1}7c4a8d09ca3762af61e59520943dc26494f8941b", hashPwd)
	log.Debugf("hash password [%d] :%s", len(hashPwd), hashPwd)

	log.SetOutputLevel(0)
	hashPwd256, ok256 := HashHelper().HashWithPrefix("123456", SHA256)
	assert.True(t, ok256)
	assert.NotEmpty(t, hashPwd256)
	assert.Equal(t, "{sha256}8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", hashPwd256)
	log.Debugf("hash password [%d] :%s", len(hashPwd256), hashPwd256)
}
func TestPasswordHashAlgo(t *testing.T) {
	algo, ok := HashHelper().HashAlgorithm("{sha1}7c4a8d09ca3762af61e59520943dc26494f8941b")
	assert.True(t, ok)
	assert.Equal(t, SHA1, algo)

	algo2, ok2 := HashHelper().HashAlgorithm("{sha256}8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92")
	assert.True(t, ok2)
	assert.Equal(t, SHA256, algo2)
}
