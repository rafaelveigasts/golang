package routes

import (
	"api-devbook/src/controllers"
)

var LoginRoute = Route{
  URI: "/login",
  Method: "POST",
  Function: controllers.Login,
  RequiresAuth: false,
}