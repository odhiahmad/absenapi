package handler

import (
	"bri-rece/api/controllers"
	"bri-rece/api/middlewares"
	"bri-rece/api/models"
	"bri-rece/api/models/dto"
	"bri-rece/api/usecase"
	"bri-rece/api/utils/formaterror"
	"bri-rece/api/utils/httpParse"
	"bri-rece/api/utils/httpResponse"
	"bri-rece/api/utils/status"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountController struct {
	router        *mux.Router
	parseJson     *httpParse.JsonParse
	responder     httpResponse.IResponder
	service       usecase.IAccountUsecase
	serviceWallet usecase.IWalletUsecase
}

func NewAccountController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, service usecase.IAccountUsecase, serviceWallet usecase.IWalletUsecase) controllers.IDelivery {
	return &AccountController{
		router,
		parse,
		responder,
		service,
		serviceWallet,
	}
}

func (a *AccountController) InitRoute(mdw ...mux.MiddlewareFunc) {
	u := a.router.PathPrefix("/account").Subrouter()
	u.Use(mdw...)
	u.HandleFunc("", middlewares.SetMiddlewareJSON(a.SaveAccount)).Methods("POST")
	u.HandleFunc("/login", middlewares.SetMiddlewareJSON(a.Login)).Methods("POST")
	u.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(a.GetAccountById))).Methods("GET")
	u.HandleFunc("/unactive/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(a.UnActiveAccount))).Methods("DELETE")
	u.HandleFunc("/activated/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(a.ActivatedAccount))).Methods("PUT")
}

func (a *AccountController) SaveAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	account := models.Account{
		IsActive: true,
	}
	err = json.Unmarshal(body, &account)
	if err != nil {
		a.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	account.Prepare()
	err = account.Validate("")

	if err != nil {
		a.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	accountCreated, err := a.service.SaveAccount(&account)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		a.responder.Error(w, http.StatusInternalServerError, formattedError.Error())
		return

	}

	a.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), accountCreated)
}

func (a *AccountController) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var account dto.Login
	fmt.Println("account", &account)
	err = json.Unmarshal(body, &account)
	if err != nil {
		a.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	token, err := a.service.LoginByUsername(&account)
	if err != nil {
		a.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	a.responder.Data(w, http.StatusOK, status.StatusText(status.Success), token)
}

func (a *AccountController) UnActiveAccount(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	accountId, err := a.service.UnActiveAccount(param["id"])

	if err != nil {
		a.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	a.responder.Data(w, http.StatusOK, http.StatusText(status.Success), accountId)

}

func (a *AccountController) ActivatedAccount(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	accountId, err := a.service.ActivatedAccount(param["id"])

	if err != nil {
		a.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	a.responder.Data(w, http.StatusOK, http.StatusText(status.Success), accountId)

}

func (a *AccountController) GetAccountById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	accountId, err := a.service.FindByIdAccount(param["id"])
	if err != nil {
		a.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	a.responder.Data(w, status.Success, status.StatusText(status.Success), accountId)
}
