package handler

import (
	"blog-api/middleware"
	"blog-api/models"
	"blog-api/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type PostHandler struct {
	Service *services.PostService
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "could not get user id", http.StatusInternalServerError)
		return
	}

	post.UserID = userID

	err = h.Service.CreatePost(&post)
	if err != nil {
		http.Error(w, "could not create post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "post ID is required", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := h.Service.GetPostByID(uint(idInt))
	if err != nil {
		http.Error(w, "could not fetch post", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(post)

}
