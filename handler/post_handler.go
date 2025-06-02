package handler

import "blog-api/services"

type PostHandler struct {
	Service *services.PostService
}