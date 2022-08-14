package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Kolyan4ik99/blog-app/pkg/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Posts struct {
	Repo repository.Posts
}

func (p *Posts) GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, p.Repo.GetAll())
}

func (p *Posts) GetPostByID(c *gin.Context) {
	str := c.Param("id")
	id, err := p.findPostById(str)
	if err != nil {
		BadRequest(c, fmt.Sprintf("Invalid id = [%s]", str))
		return
	}

	post, err := p.Repo.GetById(id)
	if errors.Is(err, sql.ErrNoRows) {
		NotFound(c, fmt.Sprintf("Post by id=%d not found", id))
		return
	}
	if err != nil {
		InternalServerError(c, "")
		return
	}

	c.JSON(http.StatusOK, post)
}

func (p *Posts) UploadPost(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p *Posts) UpdatePostByID(c *gin.Context) {
	str := c.Param("id")
	id, err := p.findPostById(str)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid id = [%s]", str))
		return
	}

	//TODO implement me
	c.JSON(http.StatusOK, fmt.Sprintf("User-id = %d", id))
}

func (p *Posts) DeletePostByID(c *gin.Context) {
	str := c.Param("id")
	id, err := p.findPostById(str)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid id = [%s]", str))
		return
	}
	//TODO implement me
	c.JSON(http.StatusOK, fmt.Sprintf("User-id = %d", id))
}

func (p *Posts) findPostById(id string) (int, error) {
	num, err := strconv.Atoi(id)
	if err != nil {
		return -1, err
	}
	if id == "" {
		return -1, errors.New("bad post id")
	}

	// Подумать над проверка существования записи с определенным id
	return num, nil
}
