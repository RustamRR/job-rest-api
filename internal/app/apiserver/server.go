package apiserver

import (
	"github.com/RustamRR/job-rest-api/internal/app/apiserver/controller/user"
	"github.com/RustamRR/job-rest-api/internal/store"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Server struct {
	router *echo.Echo
	logger *zap.Logger
	store  store.Store
}

func New(l *zap.Logger, s store.Store) *Server {
	server := &Server{
		router: echo.New(),
		logger: l,
		store:  s,
	}

	server.configureRouter()
	return server
}

func (s *Server) Run(port string) {
	s.router.Logger.Fatal(s.router.Start(port))
}

func (s *Server) configureRouter() {
	user.InitRoutes(s.router, s.store)
}
