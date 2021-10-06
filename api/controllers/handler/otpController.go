package handler

import (
	"bri-rece/api/controllers"
	"bri-rece/api/middlewares"
	"bri-rece/api/models"
	"bri-rece/api/usecase"
	"bri-rece/api/utils/httpParse"
	"bri-rece/api/utils/httpResponse"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type otpController struct {
	router    *mux.Router
	parseJson *httpParse.JsonParse
	responder httpResponse.IResponder
	service   usecase.IOtpUseCase
}

func NewOtpController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, service usecase.IOtpUseCase) controllers.IDelivery {
	return &otpController{
		router,
		parse,
		responder,
		service,
	}
}

func (a *otpController) InitRoute(mdw ...mux.MiddlewareFunc) {
	u := a.router.PathPrefix("/otp").Subrouter()
	u.Use(mdw...)
	u.HandleFunc("/verification", middlewares.SetMiddlewareJSON(a.Verfication)).Methods("POST")
}

func (a *otpController) Verfication(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var otpRequest models.OtpRequest
	err = json.Unmarshal(body, &otpRequest)
	if err != nil {
		a.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	otp, err := a.service.Verification(otpRequest.UserId, otpRequest.OtpCode)
	if err != nil {
		a.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}

	if otp.OtpCode == "" {
		a.responder.Error(w, http.StatusUnauthorized, "OTP Verification Unsuccessfull")
		return
	}

	a.responder.Data(w, http.StatusOK, "OTP Verification Successfull", nil)
}
