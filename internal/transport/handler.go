package transport

import (
	"net/http"
	"strings"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
)

// Handler В инициализации использую интерфейсы, дабы при изменении
// реализации сервисов не изменять реализацию InitRouter
type Handler struct {
	authTransport   AuthInterface
	postTransport   PostInterface
	accessTransport AccessInterface
}

func NewHandler(authTransport AuthInterface, postTransport PostInterface, accessTransport AccessInterface) *Handler {
	return &Handler{authTransport: authTransport, postTransport: postTransport, accessTransport: accessTransport}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathServ := r.URL.Path
	method := r.Method

	signUp := "/v1/auth/sing-up"

	getAllPosts := "/v1/api/post/"
	getPostById := "/v1/api/post/"
	uploadPost := "/v1/api/post/"
	updatePost := "/v1/api/post/"
	deletePost := "/v1/api/post/"

	getAccess := "/v1/api/post/access/"
	postAccess := "/v1/api/post/access/"

	if pathServ == signUp {
		h.authTransport.Signup(w, r)
	} else if pathServ == getAllPosts && method == http.MethodGet {
		h.postTransport.GetAllPosts(w, r)

	} else if pathServ == uploadPost && method == http.MethodPost {
		h.postTransport.UploadPost(w, r)

	} else if strings.Contains(pathServ, updatePost) && method == http.MethodPut {
		h.postTransport.UpdatePostByID(w, r)

	} else if strings.Contains(pathServ, deletePost) && method == http.MethodDelete {
		h.postTransport.DeletePostByID(w, r)

	} else if strings.Contains(pathServ, getAccess) && method == http.MethodGet {
		h.accessTransport.GetAccessPost(w, r)

	} else if strings.Contains(pathServ, postAccess) && method == http.MethodPost {
		h.accessTransport.SetAccessPost(w, r)

	} else if strings.Contains(pathServ, getPostById) && method == http.MethodGet {
		h.postTransport.GetPostByID(w, r)
	}
}

// InitRouter конструктор роута с эндпоинтами
func (h *Handler) InitRouter() http.Handler {
	mux := NewMiddlewareMux(h.authTransport)
	mux.AppendMiddleware(mux.authMiddleware)
	mux.Handle("/", h)
	return mux
}
