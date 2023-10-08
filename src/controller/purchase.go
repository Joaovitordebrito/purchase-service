package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Joaovitordebrito/wex-purchase-service/src/db"
	"github.com/Joaovitordebrito/wex-purchase-service/src/model"
	"github.com/Joaovitordebrito/wex-purchase-service/src/repository"
	response "github.com/Joaovitordebrito/wex-purchase-service/src/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CreatePurchase(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var purchase model.Purchase
	err = json.Unmarshal(requestBody, &purchase)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err = purchase.Prepare()
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	uuid := uuid.New().String()
	purchase.UUID = uuid

	purchaseRepo := repository.NewPurchaseRepo(db)
	_, err = purchaseRepo.Create(purchase)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, purchase)
}

func GetConvertedCurrency(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	purchaseUUID := params["uuid"]

	country := params["country"]
	titleCase := cases.Title(language.AmericanEnglish)
	country = titleCase.String(country)

	db, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}
	repo := repository.NewPurchaseRepo(db)
	purchase, err := repo.GetConvertedCurrency(purchaseUUID, country)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusOK, purchase)
	defer db.Close()
}
