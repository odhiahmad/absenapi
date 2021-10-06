package handler

import (
	"bri-rece/api/controllers"
	"bri-rece/api/middlewares"
	"bri-rece/api/models"
	"bri-rece/api/usecase"
	"bri-rece/api/utils/httpParse"
	"bri-rece/api/utils/httpResponse"
	"bri-rece/api/utils/status"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type WalletController struct {
	router         *mux.Router
	parseJson      *httpParse.JsonParse
	responder      httpResponse.IResponder
	service        usecase.IWalletUsecase
	serviceHistory usecase.IWalletHistoryUsecase
}

func (t *WalletController) InitRoute(mdw ...mux.MiddlewareFunc) {
	wallet := t.router.PathPrefix("/wallet").Subrouter()
	wallet.Use(mdw...)
	wallet.HandleFunc("/topup/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(t.TopUp))).Methods("PUT")
	wallet.HandleFunc("/withdraw/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(t.WithDraw))).Methods("PUT")
	wallet.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(t.GetWalletById))).Methods("GET")
}

func NewWalletController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, service usecase.IWalletUsecase, serviceHistory usecase.IWalletHistoryUsecase) controllers.IDelivery {
	return &WalletController{
		router,
		parse,
		responder,
		service,
		serviceHistory,
	}
}

func (t *WalletController) TopUp(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	var id, _ = uuid.FromString(param["id"])
	var wallet models.Wallet
	if err := t.parseJson.Parse(r, &wallet); err != nil {
		t.responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err := t.service.TopUp(&wallet, param["id"])
	if err != nil {
		t.responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	walletHistory := models.WalletHistory{
		TransactionDate: time.Now(),
		Amount:          wallet.Balance,
		WalletID:        id,
		TypeTransaction: models.TOPUP,
		CreatedAt:       time.Now(),
	}
	t.serviceHistory.CreateHistory(&walletHistory)
	t.responder.Data(w, status.Success, status.StatusText(status.Success), "You have successfully top up")
}

func (t *WalletController) WithDraw(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	var id, _ = uuid.FromString(param["id"])
	var wallet models.Wallet
	if err := t.parseJson.Parse(r, &wallet); err != nil {
		t.responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	wallet.Prepare()
	_, err := t.service.WithDraw(&wallet, param["id"])
	if err != nil {
		t.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	walletHistory := models.WalletHistory{
		TransactionDate: time.Now(),
		Amount:          wallet.Balance,
		WalletID:        id,
		TypeTransaction: models.WITHDRAW,
		CreatedAt:       time.Now(),
	}
	t.serviceHistory.CreateHistory(&walletHistory)
	t.responder.Data(w, status.Success, status.StatusText(status.Success), "You have successfully withdraw")
}

func (t *WalletController) GetWalletById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	walletId, err := t.service.GetWalletById(param["id"])
	if err != nil {
		t.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	t.responder.Data(w, status.Success, status.StatusText(status.Success), walletId)
}
