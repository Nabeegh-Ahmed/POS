package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pos/api/models"
	"pos/api/responses"
	"strconv"
)

func (server *Server) CreateItem(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	// Unmarshal the body into a payload
	item := models.Item{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = item.CreateItem(server.db)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(res, http.StatusCreated, item)
}

func (server *Server) GetItems(res http.ResponseWriter, req *http.Request) {
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

	item := models.Item{}
	items, err := item.GetItems(offset, limit, server.db)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(res, http.StatusOK, map[string]interface{}{"items": items})
}

func (server *Server) UpdateItem(res http.ResponseWriter, req *http.Request) {

}

func (server *Server) GetItem(res http.ResponseWriter, req *http.Request) {

}
