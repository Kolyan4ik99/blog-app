package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/Kolyan4ik99/blog-app/internal/logger"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

type Posts struct {
	Repo *repository.Posts
	Ctx  *context.Context
}

func (p *Posts) GetPosts(c *gin.Context) {
	postsInfo, err := p.Repo.GetAll(p.Ctx)
	if err != nil {
		InternalServerError(c)
		logger.Logger.Errorln(err)
	}

	logger.Logger.Infof("Return posts %s", postsInfo)
	c.JSON(http.StatusOK, postsInfo)
}

func (p *Posts) GetPostByID(c *gin.Context) {
	id, err := p.findPostById(c)
	if err != nil {
		BadRequest(c)
		logger.Logger.Warningf("Invalid id = [%s]\n", err)
		return
	}

	post, err := p.Repo.GetById(p.Ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		NotFound(c)
		logger.Logger.Warningf("Post by id=%d not found\n", id)
		return
	}
	if err != nil {
		InternalServerError(c)
		logger.Logger.Errorf("Something wrong %s\n", err)
		return
	}

	logger.Logger.Infof("Return post %s", post)
	c.JSON(http.StatusOK, post)
}

func (p *Posts) UploadPost(c *gin.Context) {
	// TODO добавить обработку загрузки поста не авторизованным пользователем
	newPost, err := p.parseBodyPost(c)
	if err != nil {
		BadRequest(c)
		logger.Logger.Warningf("Bad body: %s\n", err)
		return
	}

	if err = p.Repo.Save(p.Ctx, newPost); err != nil {
		c.JSON(http.StatusNotFound, "Author not exists")
		logger.Logger.Warningf("Author not exist %s\n", newPost)
		return
	}

	logger.Logger.Infof("Post was succesful created: %s", newPost)
	c.Status(http.StatusCreated)
}

func (p *Posts) UpdatePostByID(c *gin.Context) {
	id, err := p.findPostById(c)
	if err != nil {
		logger.Logger.Warningf("Invalid id = [%s]", err)
		BadRequest(c)
		return
	}

	newPost, err := p.parseBodyPost(c)
	if err != nil {
		logger.Logger.Warningf("Bad body: %s\n", err)
		BadRequest(c)
		return
	}

	err = p.Repo.UpdateById(p.Ctx, newPost, id)
	if err != nil {
		logger.Logger.Errorf("Something wrong %s\n", err)
		InternalServerError(c)
		return
	}

	c.Status(http.StatusOK)
	logger.Logger.Infof("User-id was successful update, id = %d\n", id)
}

func (p *Posts) DeletePostByID(c *gin.Context) {
	id, err := p.findPostById(c)
	if err != nil {
		BadRequest(c)
		logger.Logger.Warningf("Invalid id = [%s]", err)
		return
	}

	err = p.Repo.DeleteById(p.Ctx, id)
	if err != nil {
		InternalServerError(c)
		logger.Logger.Errorln(err.Error())
		return
	}

	c.Status(http.StatusOK)
	logger.Logger.Infof("Post was succesful delete, id=%d", id)
}

func (p *Posts) parseBodyPost(c *gin.Context) (*repository.PostInfo, error) {
	arr, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	c.Request.Body.Close()
	var newPost repository.PostInfo
	err = json.Unmarshal(arr, &newPost)
	if err != nil {
		return nil, err
	}
	return &newPost, nil
}

func (p *Posts) findPostById(c *gin.Context) (int, error) {
	strId := c.Param("id")
	if strId == "" {
		return -1, errors.New("bad post id")
	}
	num, err := strconv.Atoi(strId)
	if err != nil {
		return -1, err
	}

	return num, nil
}
