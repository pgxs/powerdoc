package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"regexp"
	"sync"

	log "gopkg.in/logger.v1"
)

const (
	//SHA1 sha1算法hash
	SHA1 = iota
	//SHA256 sha256 算法求hash值
	SHA256
	SHA512
	MD5
)

//HashHelper
type HashHelperImpl struct {
}

var (
	hashHelper         *HashHelperImpl
	hashHelperInitOnce sync.Once
)

//HashHelper get hash helper instance
func HashHelper() *HashHelperImpl {
	hashHelperInitOnce.Do(func() {
		hashHelper = new(HashHelperImpl)
	})
	return hashHelper
}

//PasswordHashAlgorithm {sha1}xialvaiao 提取密码hash算法 support sha1 sha256 md4
func (hh HashHelperImpl) HashAlgorithm(hasPrefixHashStr string) (method string, ok bool) {
	pwdReg := regexp.MustCompile(`^{(sha1|sha256|sha512|md5)}([a-zA-Z0-9]+)*$`)
	params := pwdReg.FindStringSubmatch(hasPrefixHashStr)

	log.Debug(params)
	if len(params) == 3 {
		method = params[1]
		ok = true
	}
	return
}

//Md5String md5 data
func (hh HashHelperImpl) Md5String(data string) string {
	sha := md5.New()
	sha.Write([]byte(data))
	return hex.EncodeToString(sha.Sum([]byte(nil)))
}

//Sha1String sha1 data
func (hh HashHelperImpl) Sha1String(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}

//Sha256String sha1 data
func (hh HashHelperImpl) Sha256String(data string) string {
	sha1 := sha256.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}

//Sha256WithPrefix sha1 data string with {sh256} prefix
func (hh HashHelperImpl) Sha256WithPrefix(data string) string {
	return "{sha256}" + hh.Sha256String(data)
}

//Sha512String sha1 data
func (hh HashHelperImpl) Sha512String(data string) string {
	sha := sha512.New()
	sha.Write([]byte(data))
	return hex.EncodeToString(sha.Sum([]byte(nil)))
}

//Sha512String sha1 data
func (hh HashHelperImpl) Sha512WithPrefix(data string) string {
	return fmt.Sprintf("{sha512}%s", hh.Sha512String(data))
}

//Md5 md5 byte data
func (hh HashHelperImpl) Md5(data []byte) string {
	md5 := md5.New()
	md5.Write(data)
	return hex.EncodeToString(md5.Sum([]byte(nil)))
}

//Sha1 sha1 byte data
func (hh HashHelperImpl) Sha1(data []byte) string {
	sha1 := sha1.New()
	sha1.Write(data)
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}

//Sha256 sha256 byte data
func (hh HashHelperImpl) Sha256(data []byte) string {
	sha1 := sha256.New()
	sha1.Write(data)
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}

//Sha512 sha256 byte data
func (hh HashHelperImpl) Sha512(data []byte) string {
	sha := sha512.New()
	sha.Write(data)
	return hex.EncodeToString(sha.Sum([]byte(nil)))
}

//HashWithPrefix hash password
func (hh HashHelperImpl) HashWithPrefix(plainTxt string, algo int) (hashString string, ok bool) {
	switch algo {
	case MD5:
		ok = true
		hashString = "{md5}" + hh.Md5String(plainTxt)
	case SHA1:
		ok = true
		hashString = "{sha1}" + hh.Sha1String(plainTxt)
	case SHA256:
		ok = true
		hashString = "{sha256}" + hh.Sha256String(plainTxt)
	case SHA512:
		ok = true
		hashString = "{sha512}" + hh.Sha512String(plainTxt)
	default:
		ok = false
	}

	return
}
