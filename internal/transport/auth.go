package transport

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Kolyan4ik99/blog-app/internal/logger"
	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
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
func (a *Auth) Signup(c *gin.Context) {
	var user model.AuthInput
	arr, err := io.ReadAll(c.Request.Body)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer c.Request.Body.Close()
	err = json.Unmarshal(arr, &user)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		logger.Logger.Warningf("Bad request: %s\n", err)
		return
	}

	id, err := a.authService.SignUp(a.ctx, &user)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		logger.Logger.Warningf("Bad request: %s\n", err)
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (a *Auth) Signin(c *gin.Context) {
	// TODO реализовать метод
	c.JSON(http.StatusInternalServerError, "")
}

func (a *Auth) CheckToken(token string) bool {
	if a.authService.CheckToken(a.ctx, token) != nil {
		return false
	}
	return true
}
