package auth

import "github.com/kataras/iris"

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
	m.Ctx.Text(iris.StatusUnauthorized, AUTH_FAILURE_MSG)
	m.Ctx.SetConnectionClose()
}

func (m AuthMiddleware) Verify() bool{
	return true
}


func (m AuthMiddleware) verifyToken()  error{
	token, err  := Token(m.Ctx.Session().GetString(SESSION_TOKEN_SECERET)).Decode()
	if(err != nil){
		m.Ctx.Text(iris.StatusForbidden, MISS_TOKEN_MSG)
		return err
	}
	m.Ctx.Session().Set(SESSION_TOKEN_KEY, string(token))
	
	return  token, err
}

func (m AuthMiddleware) verifySession(){
	
}

func (m AuthMiddleware) verifyUser(){
	
}

func (m AuthMiddleware) verifyRole(){

}

func (m AuthMiddleware) bind(){
	
}


func RequireRole(roles ...string){
	instance := AuthMiddleware{roles}
	return instance
}