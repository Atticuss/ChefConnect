package controllers

// Controller is an interface that all controller implementations must follow
type Controller interface {
	SetupController() error
	Run() error
	Stop() error
}
