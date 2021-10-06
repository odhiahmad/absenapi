package handler

import (
	"bri-rece/api/controllers"
	"bri-rece/api/middlewares"
	"bri-rece/api/usecase"
	"bri-rece/api/utils/httpParse"
	"bri-rece/api/utils/httpResponse"
	"bri-rece/api/utils/status"
	"encoding/json"
	_ "errors"
	"io/ioutil"
	"net/http"
	_ "strconv"

	"bri-rece/api/models"
	"bri-rece/api/utils/formaterror"

	"github.com/gorilla/mux"
)

type UserController struct {
	router    *mux.Router
	parseJson *httpParse.JsonParse
	responder httpResponse.IResponder
	service   usecase.IUserUseCase
}

func NewUserController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, service usecase.IUserUseCase) controllers.IDelivery {
	return &UserController{
		router, parse, responder, service,
	}
}

func (s *UserController) InitRoute(mdw ...mux.MiddlewareFunc) {
	u := s.router.PathPrefix("/users").Subrouter()
	u.Use(mdw...)
	u.HandleFunc("", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	u.HandleFunc("/edit", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.EditUser))).Methods("PUT")
	u.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.FindUserByID))).Methods("GET")
	u.HandleFunc("/ping", s.Home).Methods("POST")
}

func (u *UserController) Home(w http.ResponseWriter, r *http.Request) {
	u.responder.Data(w, http.StatusOK, status.StatusText(status.CREATED),"Welcome To Rece")

}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	var user models.UserRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = user.Validate("")
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	userCreated, _ := u.service.Register(&user)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		u.responder.Error(w, http.StatusInternalServerError, formattedError.Error())
		return
	}
	u.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), userCreated)
}

func (u *UserController) EditUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userEdit, err := u.service.UpdateInfo(&user)
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}

	u.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), userEdit)
}

func (u *UserController) FindUserByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userId, err := u.service.FindUserById(param["id"])
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	u.responder.Data(w, status.Success, status.StatusText(status.Success), userId)
}
