package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/oklog/run"
	log "go.uber.org/zap"

	"intouche-back-core/internal/config"
)

const (
	headerXRequestID      = "X-Request-Id"
	constAuthHeader       = "Authorization"
	constBearerAuthPrefix = "Bearer"
)

type Server struct {
	*http.Server
	rm     *ResponseManager
	cfg    *config.Config
	logger *log.SugaredLogger
}

//TODO: add minio support
func NewServer(cfg *config.Config, logger *log.SugaredLogger,

//minioStore db.MinioStore
) *Server {
	srv := &Server{
		Server: &http.Server{
			Addr:         cfg.API.Address,
			ReadTimeout:  time.Duration(cfg.API.ReadTimeout),
			WriteTimeout: time.Duration(cfg.API.WriteTimeout),
		},
		rm:     NewResponseManager(logger),
		cfg:    cfg,
		logger: logger,
		//minioStore: minioStore,
	}

	r := chi.NewRouter()

	//r.Handle("/", graphql)
	srv.Handler = r
	return srv
}

func (s *Server) Run(g *run.Group) {
	g.Add(func() error {
		s.logger.Info("[http-server] started")
		return s.ListenAndServe()
	}, func(err error) {
		s.logger.Error("[http-server] terminated", err)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.cfg.API.ShutdownTimeout))
		defer cancel()

		s.logger.Error("[http-server] stopped", s.Shutdown(ctx))
	})
}
