package liveview

import (
	_ "embed"
	"fmt"

	"github.com/atticuss/chefconnect/models"
	"github.com/brendonmatos/golive"
)

//go:embed templates/login.html
var loginTemplate string

type LoginModal struct {
	golive.LiveComponentWrapper
	Root *Root

	Error    string
	Username string
	Password string
}

func NewLoginModal(root *Root) *golive.LiveComponent {
	return golive.NewLiveComponent("LoginModal", &LoginModal{Root: root})
}

func (h *LoginModal) Mounted(_ *golive.LiveComponent) {

}

func (h *LoginModal) TemplateHandler(_ *golive.LiveComponent) string {
	return loginTemplate
}

func (l *LoginModal) HandleLogin(data map[string]string) {
	user := &models.User{
		Username: l.Username,
		Password: l.Password,
	}
	user, svcErr := (*service).GenerateJwtTokens(user)
	if svcErr.Error != nil {
		l.Error = svcErr.Error.Error()
		return
	}

	fmt.Printf("user: %+v\n", user)
}
