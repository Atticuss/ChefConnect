package liveview

import (
	_ "embed"

	"github.com/atticuss/chefconnect/services"
	"github.com/brendonmatos/golive"
)

//go:embed templates/index.html
var indexTemplate string

type Root struct {
	golive.LiveComponentWrapper
	NavBar     *golive.LiveComponent
	SideBar    *golive.LiveComponent
	LoginModal *golive.LiveComponent

	Service *services.Service

	AuthToken string
	Username  string
}

func NewRoot() *golive.LiveComponent {
	root := &Root{}
	root.NavBar = NewNavBar(root)
	root.SideBar = NewSideBar(root)
	root.LoginModal = NewLoginModal(root)

	return golive.NewLiveComponent("Root", root)
}

func (h *Root) Mounted(_ *golive.LiveComponent) {

}

func (h *Root) TemplateHandler(_ *golive.LiveComponent) string {
	return indexTemplate
}
