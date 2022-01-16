package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	conf := config
	return &APIServer{
		config: conf,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}

}

// Start ...
func (s *APIServer) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("Starting API server ...")
	s.configureRouter()
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	// variables
	// types
	type request struct {
		name string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// ...
		io.WriteString(w, "Hello")

	}

}
