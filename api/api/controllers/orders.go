package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"pos/api/models"
	"pos/api/responses"
	"strconv"
)

func (server *Server) CreateOrder(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	// Unmarshal the body into a payload
	order := models.Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = order.CreateOrder(server.db)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(res, http.StatusCreated, order)
}

func (server *Server) GetOrders(res http.ResponseWriter, req *http.Request) {
	_limit := req.URL.Query().Get("limit")
	_offset := req.URL.Query().Get("offset")

	if _limit == "" {
		_limit = "10"
	}

	if _offset == "" {
		_offset = "0"
	}

	limit, err := strconv.Atoi(_limit)
	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}

	offset, err := strconv.Atoi(_offset)
	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}

	order := models.Order{}
	orders, err := order.GetOrders(offset, limit, server.db)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(res, http.StatusOK, map[string]interface{}{"orders": orders})
}

func (server *Server) UpdateOrder(res http.ResponseWriter, req *http.Request) {

}

func (server *Server) GetOrder(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	orderId := params["id"]

	order := models.Order{}

	foundOrder, err := order.GetOrder(orderId, server.db)

	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(res, http.StatusOK, map[string]interface{}{"order": foundOrder})
}
