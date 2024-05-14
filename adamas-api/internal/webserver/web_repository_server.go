package webserver

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
)

type WebRepoHandler struct {
	RepoService *service.RepositoryService
}

func NewRepoHandler(repoService *service.RepositoryService) *WebRepoHandler {
	return &WebRepoHandler{
		RepoService: repoService,
	}
}
