package authapi

import (
	"github.com/kataras/iris"
	"github.com/minhlucvan/gotodo/API"
	"github.com/minhlucvan/gotodo/utils/validate"
)


type Login struct {
	UserName string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"username" validate:"required"`
}

func (l Login) Validate() error{
	valid := validate.New()
	
	errs := valid.Struct(l)
	
	if errs != nil {
		//fuck you bitch!
	}
	
	return errs
}


func (api Login) Serve(ctx *iris.Context){
	response := API.JSON()
	form := Login{}
	formErr := ctx.ReadForm(&form)
	if formErr != nil {
		response = API.JSON().Err(formErr).Msg("bad request")
		ctx.JSON(iris.StatusBadRequest, response)
		return
	}
	validateErr := form.Validate()
	
	if validateErr != nil {
		response = API.JSON().Err(validateErr).Msg("validation error found")
		ctx.JSON(iris.StatusBadRequest, response)
		return
	}
	
	ctx.Session().Set("token", "secret")
	ctx.JSON(iris.StatusOK, response)
	
	response = API.JSON(ctx.Session().GetAll()).Msg("login sucessfull")
	ctx.JSON(iris.StatusBadRequest, response)
}
