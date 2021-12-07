package liveview

import (
	_ "embed"

	"github.com/brendonmatos/golive"
)

//go:embed templates/sidebar.html
var sidebarTemplate string

type SideBar struct {
	golive.LiveComponentWrapper
	Root     *Root
	Username string
}

func NewSideBar(root *Root) *golive.LiveComponent {
	return golive.NewLiveComponent("SideBar", &SideBar{Root: root})
}

func (h *SideBar) Mounted(_ *golive.LiveComponent) {

}

func (h *SideBar) TemplateHandler(_ *golive.LiveComponent) string {
	return sidebarTemplate
}
