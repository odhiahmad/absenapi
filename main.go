package main

import (
	"bri-rece/api/connect"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type briApp struct {
	connect connect.Connect
	router  *mux.Router
}

func (app *briApp) run() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	h := app.connect.ApiServer()
	log.Println("Listening on", h)
	NewAppRouter(app).InitMainRoutes()
	err := http.ListenAndServe(h, handlers.CORS(originsOk, headersOk, methodsOk)(app.router))
	if err != nil {
		log.Fatalln(err)
	}

}

func AbsensiApp() *briApp {
	r := mux.NewRouter()
	var appConnect = connect.NewConnect()
	return &briApp{
		connect: appConnect,
		router:  r,
	}
}

func main() {
	AbsensiApp().run()

}
