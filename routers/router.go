package routers

import (
	"firstDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/toLogin", &controllers.UserController{}, "get:Get")
	beego.Router("/Login", &controllers.UserController{}, "POST:Login")
	beego.Router("/regiser", &controllers.RegisterController{})
	beego.Router("/doregister", &controllers.RegisterController{}, "POST:Doregister")
	beego.Router("/deleteuser/:id", &controllers.UserListController{}, "DELETE:DeleteUser")
	beego.Router("/userlist", &controllers.UserListController{})
	beego.Router("/userlist/GetByCustomerID/:id", &controllers.UserListController{}, "get:GetByCustomerID")
	beego.Router("/getall", &controllers.UserListController{}, "GET:GetAll")
	beego.Router("/add", &controllers.AddController{})
	beego.Router("/doadd", &controllers.AddController{}, "POST:Doadd")
	beego.Router("/competitions", &controllers.CompetitionController{})
	beego.Router("/competitions/GetByCustomerID/:id", &controllers.CompetitionController{}, "get:GetByCustomerID")
	beego.Router("/tocompetitions", &controllers.CompetitionController{}, "GET:GetAll")
	beego.Router("/DeleteCompetition/:id", &controllers.CompetitionController{}, "DELETE:DeleteCompetition")
	beego.Router("/CompetitionCategory", &controllers.CompetitionCategoryController{})
	beego.Router("/toCompetitionCategory", &controllers.CompetitionCategoryController{}, "GET:GetAll")
	beego.Router("/page", &controllers.UserListController{}, "POST:ShowUserByPage")
	beego.Router("/UpdateUser", &controllers.UpdateUserController{}, "get:Get;PUT:UpdateUser")
	beego.Router("/UpdateCompetition", &controllers.UpdatecompetitionsController{}, "get:Get;POST:UpdateCompetition")
	beego.Router("/updacomprtition/:id", &controllers.UpdatecompetitionsController{})
	beego.Router("/Updatepage/:id", &controllers.UpdateUserController{})
	beego.Router("/Registration", &controllers.RegistrationsController{})
	beego.Router("/Registration/GetByCustomerID/:id", &controllers.RegistrationsController{}, "get:GetByCustomerID")
	beego.Router("/DeleteRegistration/:id", &controllers.RegistrationsController{}, "DELETE:DeleteRegistration")
	beego.Router("/toRegistrations", &controllers.RegistrationsController{}, "GET:GetAll")
	beego.Router("/addregistration", &controllers.AddRegistrationController{})
	beego.Router("/doaddregistration", &controllers.AddRegistrationController{}, "post:DoAddRegistration")
	beego.Router("/updateRegistration/:id", &controllers.UpdateRegistrationsController{})
	beego.Router("/UpdateRegistration", &controllers.UpdateRegistrationsController{}, "get:Get;POST:UpdateRegistration")
	beego.Router("/HOMEPAGE", &controllers.HOMEPAGEController{})
	beego.Router("/notifications", &controllers.NotificationController{})
	beego.Router("/tonotifications", &controllers.NotificationController{}, "get:GetAll")
}
