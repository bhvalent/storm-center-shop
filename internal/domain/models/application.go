package models

import "log"

type Application struct {
	Config Config
	Logger *log.Logger
}
