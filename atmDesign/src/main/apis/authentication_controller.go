package apis

import (
	"atmDesign/src/main/services"
	"net/http"
)

type AuthenticationController struct {
	authenticationService services.IAuthenticationOrch
}

func NewAuthenticationController(authenticationService services.IAuthenticationOrch) *AuthenticationController {
	return &AuthenticationController{authenticationService:authenticationService};
}

func(controller *AuthenticationController) authenticate(req *http.Request) *Response {
	return controller.authenticationService.Authenticate()
}