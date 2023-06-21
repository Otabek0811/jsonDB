package controller

import (
	"github.com/Udevs/mini_Project/config"
	"github.com/Udevs/mini_Project/storage"
)

type Controller struct {
	Cfg  *config.Config
	Strg storage.StorageI
}

func NewController(cfg *config.Config, storage storage.StorageI) *Controller {
	return &Controller{
		Cfg:  cfg,
		Strg: storage,
	}
}
