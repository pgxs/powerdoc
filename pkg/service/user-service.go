package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"pgxs.io/powerdoc/pkg/config"
	"pgxs.io/powerdoc/pkg/types"
	"pgxs.io/powerdoc/pkg/util"
	"sync"
	"time"

	logx "pgxs.io/chassis/log"
	"pgxs.io/powerdoc/pkg/mapper"
	"pgxs.io/powerdoc/pkg/mapper/api"
)

//UserService user service for sign in/up/out and manage user info.
type UserService struct {
	userMapper api.UserMapper
	log        *logx.Entry
}

var (
	userService     *UserService
	userSvcInitOnce sync.Once
)

//UserServiceInstance user service singleton.
func UserServiceInstance() *UserService {
	userSvcInitOnce.Do(func() {
		userService = new(UserService)
		userService.userMapper = mapper.UserMapper()
		userService.log = logx.New().Category("service").Component("user")
	})
	return userService
}

//UserLogin user login successfully return true or false
func (svc *UserService) SignIn(username, password string) (token string, err error) {

	user := mapper.UserMapper().FindByLoginName(username)
	//check password
	if user == nil {
		err = errors.New("user not found")
		return
	}

	ok := svc.verifyPwd(password, user.Password)
	if ok {
		token, err = svc.buildToken(user)
		if err != nil {
			svc.log.Debug(err)
			err = errors.New("build token failed")
		}
		svc.log.Debugf("user %s token: %s", username, token)
		return
	}
	err = errors.New("user name or password is wrong")
	return
}

func (svc *UserService) verifyPwd(inputPwd, pwdInDB string) (ok bool) {
	algo, hasPrefix := util.HashHelper().HashAlgorithm(pwdInDB)
	if !hasPrefix {
		return
	}

	inputPwdHash := ""

	switch algo {
	case "sha256":
		inputPwdHash = util.HashHelper().Sha256WithPrefix(inputPwd)
	case "sha512":
		inputPwdHash = util.HashHelper().Sha512WithPrefix(inputPwd)
	}

	if inputPwdHash == pwdInDB {
		ok = true
	}
	svc.log.Debugf("inputPwd hash %s ,db pwd:%s", inputPwdHash, pwdInDB)

	return
}

func (svc *UserService) buildToken(user *types.UserDO) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"exp":      time.Now().Add(2 * 3600 * time.Second),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err = token.SignedString([]byte(config.Server().JWT.JWTHmacSignSecret))
	svc.log.Debugf("token info:\nclaims:%+v\nkey: %s\ntoken:%s", token, config.Server().JWT.JWTHmacSignSecret, tokenStr)

	return
}

//SignUp user sign up
func (svc *UserService) SignUp() error {

	//todo
	return nil
}
