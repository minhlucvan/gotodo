package auth

import (
	"github.com/kataras/iris"
	"strings"
)

type AuthUser struct {
	UserName string
	Password string
}


type AuthMiddleware struct {
	Roles string
	Ctx *iris.Context
}

func (m AuthMiddleware) Serve(ctx *iris.Context){
	m.Ctx = ctx
	if(m.Verify() == true){
		m.onAuthPass()
	} else {
		m.onAuthFail()
	}
}


func (m AuthMiddleware) onAuthPass(){
	m.Ctx.Session().Set("login", "true")
	m.Ctx.Session().Set("user", struct {
		Name string
	}{Name: "minh"})
	m.Ctx.Next()
}

func (m AuthMiddleware) onAuthFail(){
	m.Ctx.EmitError(iris.StatusUnauthorized)
}

func (m AuthMiddleware) Verify() bool{
	return false
}


func (m AuthMiddleware) verifyToken()  error{
	strTooken := m.Ctx.Session().GetString(SESSION_TOKEN_SECERET)
	token, err  := Token(strTooken).Decode()
	if(err != nil){
		m.Ctx.Text(iris.StatusForbidden, MISS_TOKEN_MSG)
		return err
	}
	m.Ctx.Session().Set(SESSION_TOKEN_KEY, string(*token))
	
	return  err
}

func (m AuthMiddleware) verifySession(){
	
}

func (m AuthMiddleware) verifyUser(){
	
}

func (m AuthMiddleware) verifyRole(){

}

func (m AuthMiddleware) bind(){
	
}


func RequireRole(roles ...string) iris.Handler{
	roleStr := strings.Join(roles, ",")
	instance := AuthMiddleware{Roles: roleStr}
	return instance
}