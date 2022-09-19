package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Kolyan4ik99/blog-app/internal/logger"
	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/service"
)

type AuthInterface interface {
	Signup(w http.ResponseWriter, r *http.Request)
	Signin(w http.ResponseWriter, r *http.Request)
	CheckToken(token string) bool
}

type Auth struct {
	ctx         context.Context
	authService service.AuthInterface
}

func NewAuth(ctx context.Context, authService service.AuthInterface) *Auth {
	return &Auth{ctx: ctx, authService: authService}
}

// Signup godoc
// @Summary      Auth new user
// @Description  Аутентификация пользователя. Для дальнейшей работы,
// @Description скопировать всё тело токена в Authorize -> Пользоваться сервисом
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input       body     model.AuthInput true "account info"
// @Success      200  {object}  model.AuthOutput
// @Failure      400,401,404 {object} transport.Response
// @Router       /auth/sing-up [post]
func (a *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	var user model.AuthInput
	arr, err := io.ReadAll(r.Body)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(arr, &user)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		logger.Logger.Warningf("Bad request: %s\n", err)
		return
	}

	auth, err := a.authService.SignUp(a.ctx, &user)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		logger.Logger.Warningf("Bad request: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeResponseBody(w, fmt.Sprintf("id=[%d], token=[%s]", auth.Id, auth.Token))
}

func (a *Auth) Signin(w http.ResponseWriter, r *http.Request) {
	// TODO реализовать метод
	w.WriteHeader(http.StatusInternalServerError)
}

func (a *Auth) CheckToken(token string) bool {
	if a.authService.CheckToken(a.ctx, token) != nil {
		return false
	}
	return true
}
