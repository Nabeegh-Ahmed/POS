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

func (server *Server) CreateCollection(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	// Unmarshal the body into a payload
	collection := models.Collection{}
	err = json.Unmarshal(body, &collection)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = collection.CreateCollection(server.db)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(res, http.StatusCreated, collection)
}

func (server *Server) GetCollections(res http.ResponseWriter, req *http.Request) {
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

	collection := models.Collection{}
	collections, err := collection.GetCollections(offset, limit, server.db)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(res, http.StatusOK, map[string]interface{}{"collections": collections})
}

func (server *Server) UpdateCollection(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	// Unmarshal the body into a payload
	collection := models.Collection{}
	err = json.Unmarshal(body, &collection)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	fmt.Println("Here")
	_, err = collection.UpdateCollection(server.db)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(res, http.StatusOK, map[string]interface{}{"collection": collection})
}

func (server *Server) GetCollection(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	collectionId := params["id"]

	collection := models.Collection{}

	foundCollection, err := collection.GetCollection(collectionId, server.db)

	if err != nil {
		fmt.Println(err)
		responses.ERROR(res, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(res, http.StatusOK, map[string]interface{}{"collection": foundCollection})
}
