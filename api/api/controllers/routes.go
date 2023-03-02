package controllers

import (
	"pos/api/middlewares"
)

func (server *Server) initializeRoutes() {
	// Auth Routes
	server.router.HandleFunc("/api/v1/collections", middlewares.SetMiddlewareJSON(server.CreateCollection)).Methods("POST")
	server.router.HandleFunc("/api/v1/collections", middlewares.SetMiddlewareJSON(server.GetCollections)).Methods("GET")
	server.router.HandleFunc("/api/v1/collections", middlewares.SetMiddlewareJSON(server.UpdateCollection)).Methods("PUT")
	server.router.HandleFunc("/api/v1/collections/{id}", middlewares.SetMiddlewareJSON(server.GetCollection)).Methods("GET")

	server.router.HandleFunc("/api/v1/items", middlewares.SetMiddlewareJSON(server.CreateItem)).Methods("POST")
	server.router.HandleFunc("/api/v1/items", middlewares.SetMiddlewareJSON(server.GetItems)).Methods("GET")
	server.router.HandleFunc("/api/v1/items", middlewares.SetMiddlewareJSON(server.UpdateItem)).Methods("PUT")
	server.router.HandleFunc("/api/v1/items/{id}", middlewares.SetMiddlewareJSON(server.GetItem)).Methods("GET")
	server.router.HandleFunc("/api/v1/items/name/{name}", middlewares.SetMiddlewareJSON(server.GetItemsByName)).Methods("GET")
	server.router.HandleFunc("/api/v1/items/barcode/{code}", middlewares.SetMiddlewareJSON(server.GetItemByBarcode)).Methods("GET")

	server.router.HandleFunc("/api/v1/orders", middlewares.SetMiddlewareJSON(server.CreateOrder)).Methods("POST")
	server.router.HandleFunc("/api/v1/orders", middlewares.SetMiddlewareJSON(server.GetOrders)).Methods("GET")
	server.router.HandleFunc("/api/v1/orders/{id}", middlewares.SetMiddlewareJSON(server.GetOrder)).Methods("GET")
}
