package filters

import (
	"io"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful/v3"
)

var (
	sharedSecret = []byte("shared-token")
)

func secretJWT(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "42")
}

func validJWT(authHeader string) bool {
	if !strings.HasPrefix(authHeader, "JWT ") {
		return false
	}

	jwtToken := strings.Split(authHeader, " ")
	if len(jwtToken) < 2 {
		return false
	}
	parts := strings.Split(jwtToken[1], ".")
	err := jwt.SigningMethodHS512.Verify(strings.Join(parts[0:2], "."), parts[2], sharedSecret)
	if err != nil {
		return false
	}

	return true
}

func authJWT(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	authHeader := req.HeaderParameter("Authorization")

	if !validJWT(authHeader) {
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}

	chain.ProcessFilter(req, resp)
}
