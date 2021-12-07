package liveview

import (
	_ "embed"

	"github.com/brendonmatos/golive"
)

//go:embed templates/navbar.html
var navbarTemplate string

type NavBar struct {
	golive.LiveComponentWrapper
	Root *Root

	Username string
}

func NewNavBar(root *Root) *golive.LiveComponent {
	return golive.NewLiveComponent("NavBar", &NavBar{Root: root})
}

func (h *NavBar) Mounted(_ *golive.LiveComponent) {

}

func (h *NavBar) TemplateHandler(_ *golive.LiveComponent) string {
	return navbarTemplate
}
