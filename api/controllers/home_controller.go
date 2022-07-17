package controllers

import (
	"net/http"

	"zayed/Work/go-jwt-restapi-postgres/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
