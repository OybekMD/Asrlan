package v1

import (
	"asrlan-monolight/api/tokens"
	"asrlan-monolight/config"
	"asrlan-monolight/storage"
	"log"

	"github.com/casbin/casbin/v2"
)

type handlerV1 struct {
	cfg        *config.Config
	storage    storage.StorageI
	jwtHandler tokens.JWTHandler
	enforcer   *casbin.Enforcer
}

type HandlerV1Options struct {
	Cfg        *config.Config
	Storage    storage.StorageI
	JWTHandler tokens.JWTHandler
	Log        *log.Logger
	Enforcer   *casbin.Enforcer
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:        options.Cfg,
		storage:    options.Storage,
		jwtHandler: options.JWTHandler,
		enforcer:   options.Enforcer,
	}
}
