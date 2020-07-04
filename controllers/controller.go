package controllers

type Controller interface {
	Start() error
	Stop() error
}
