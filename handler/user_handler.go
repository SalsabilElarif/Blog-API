package handler

import (
	"blog-api/models"
	"blog-api/services"
	"encoding/json"
	"net/http"
)

// handler layer (handles (http requests& response) and call the service layer)
//  	|
// service layer (business logic and calls the repository layer )
//  	|
// repository layer (handle direct database operations)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	  var user models.User
	  err := json.NewDecoder(r.Body).Decode(&user)
	  if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	  }

	  // call the service layer
	  err = h.Service.RegisterUser(&user)
	  if err != nil {
		http.Error(w, "could not register user", http.StatusInternalServerError)
		return
	  }

	  // response
	  w.WriteHeader(http.StatusCreated)
	  json.NewEncoder(w).Encode(&user)
}
