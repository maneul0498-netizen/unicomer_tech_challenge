package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maneul0498-netizen/unicomer_tech_challenge/internal/interfaces/http/handler"
)

type Server struct {
	Eng *gin.Engine

	srvHTTP  *http.Server
	srvHTTPS *http.Server

	pathCertHTTPS string
	pathKeyHTTPS  string
}

func New() (*Server, error) {

	serverConf, err := ConfigFromEnv()

	if err != nil {
		return nil, err
	}

	eng := gin.New()
	eng.Use(gin.Logger())
	eng.Use(gin.Recovery())

	handler := handler.NewHandler()

	r := Router{
		Eng:     eng,
		Handler: handler,
	}

	r.InitRouters()

	server := &Server{
		Eng: eng,
		srvHTTP: &http.Server{
			Addr:           serverConf.AddressHTTP,
			Handler:        eng,
			ReadTimeout:    time.Duration(serverConf.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(serverConf.WriteTimeout) * time.Second,
			MaxHeaderBytes: 0,
		},
		srvHTTPS: &http.Server{
			Addr:           serverConf.AddressHTTPS,
			Handler:        eng,
			ReadTimeout:    time.Duration(serverConf.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(serverConf.WriteTimeout) * time.Second,
			MaxHeaderBytes: 0,
		},
		pathCertHTTPS: serverConf.PathCertHTTPS,
		pathKeyHTTPS:  serverConf.PathKeyHTTPS,
	}

	return server, nil

}

func (s *Server) Run(errCh chan<- error) {
	// HTTP
	go func() {
		log.Printf("http server listening on %s", s.srvHTTP.Addr)
		if err := s.srvHTTP.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	// HTTPS (solo si hay certificados)
	if len(s.pathCertHTTPS) != 0 && len(s.pathKeyHTTPS) != 0 {
		go func() {
			log.Printf("https server listening on %s", s.srvHTTPS.Addr)
			if err := s.srvHTTPS.ListenAndServeTLS(
				s.pathCertHTTPS,
				s.pathKeyHTTPS,
			); err != nil && !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()
	} else {
		log.Println("https server not started (cert or key missing)")
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.srvHTTP.Shutdown(ctx); err != nil {
		return err
	}

	if s.srvHTTPS != nil {
		if err := s.srvHTTPS.Shutdown(ctx); err != nil {
			return err
		}
	}

	return nil
}
