package services

import (
	"blog-api/repository"

)

type PostService struct {
	Repo repository.PostRepo
}