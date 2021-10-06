package main

import (
	"bri-rece/api/connect"
	"bri-rece/api/controllers"
	"bri-rece/api/controllers/handler"
	"bri-rece/api/manager"
	"bri-rece/api/middlewares"
	"bri-rece/api/utils/httpParse"
	"bri-rece/api/utils/httpResponse"

	"github.com/gorilla/mux"
)

type appRouter struct {
	app                  *briApp
	parse                *httpParse.JsonParse
	responder            httpResponse.IResponder
	connect              connect.Connect
	logRequestMiddleware *middlewares.LogRequestMiddleware
}

type appRoutes struct {
	centerRoutes controllers.IDelivery
	mdw          []mux.MiddlewareFunc
}

func (r *appRouter) InitMainRoutes() {
	r.app.router.Use(r.logRequestMiddleware.Log)
	serviceManager := manager.NewServiceManager(r.connect)
	appRoutes := []appRoutes{
		{
			centerRoutes: handler.NewUserController(r.app.router, r.parse, r.responder, serviceManager.UserUseCase()),
			mdw:          nil,
		},
		{
			centerRoutes: handler.NewAccountController(r.app.router, r.parse, r.responder, serviceManager.AccountUseCase(), serviceManager.WalletUseCase()),
			mdw:          nil,
		},
		{
			centerRoutes: handler.NewWalletController(r.app.router, r.parse, r.responder, serviceManager.WalletUseCase(), serviceManager.WalletHistoryUseCase()),
			mdw:          nil,
		},
		{
			centerRoutes: handler.NewOtpController(r.app.router, r.parse, r.responder, serviceManager.OtpUseCase()),
			mdw:          nil,
		},
	}

	for _, r := range appRoutes {
		r.centerRoutes.InitRoute(r.mdw...)
	}
}

func NewAppRouter(app *briApp) *appRouter {
	return &appRouter{
		app,
		httpParse.NewJsonParse(),
		httpResponse.NewJSONResponder(),
		app.connect,
		middlewares.NewLogRequestMiddleware(),
	}
}
