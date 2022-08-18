package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/Kolyan4ik99/blog-app/internal"
	"github.com/Kolyan4ik99/blog-app/pkg/repository"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Posts struct {
	Repo *repository.Posts
	Ctx  *context.Context
}

func (p *Posts) GetPosts(w http.ResponseWriter, r *http.Request) {
	postsInfo, err := p.Repo.GetAll(p.Ctx)
	if err != nil {
		InternalServerError(w)
		internal.Logger.Errorln(err)
	}

	retBytes, err := p.getBytes(postsInfo...)
	if err != nil {
		InternalServerError(w)
		internal.Logger.Errorln(err)
		return
	}

	internal.Logger.Infof("Return posts %s", string(retBytes))
	w.WriteHeader(http.StatusOK)
	w.Write(retBytes)
}

func (p *Posts) GetPostByID(w http.ResponseWriter, r *http.Request) {
	id, err := p.findPostById(r)
	if err != nil {
		BadRequest(w)
		internal.Logger.Warningf("Invalid id = [%s]\n", err)
		return
	}

	post, err := p.Repo.GetById(p.Ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		NotFound(w)
		internal.Logger.Warningf("Post by id=%d not found\n", id)
		return
	}
	if err != nil {
		InternalServerError(w)
		internal.Logger.Errorf("Something wrong %s\n", err)
		return
	}

	retBytes, err := p.getBytes(post)
	if err != nil {
		InternalServerError(w)
		return
	}

	internal.Logger.Infof("Return post %s", string(retBytes))
	w.WriteHeader(http.StatusOK)
	w.Write(retBytes)
}

func (p *Posts) UploadPost(w http.ResponseWriter, r *http.Request) {
	// TODO добавить обработку загрузки поста не авторизованным пользователем
	newPost, err := p.parseBodyPost(w, r)
	if err != nil {
		BadRequest(w)
		internal.Logger.Warningf("Bad body: %s\n", err)
		return
	}

	if err = p.Repo.Save(p.Ctx, newPost); err != nil {
		BadRequest(w)
		w.Write([]byte("Author not exists"))
		internal.Logger.Warningf("Author not exist %s\n", newPost)
		return
	}

	w.WriteHeader(http.StatusCreated)
	internal.Logger.Infof("Post was succesful created: %s", newPost)
}

func (p *Posts) UpdatePostByID(w http.ResponseWriter, r *http.Request) {
	id, err := p.findPostById(r)
	if err != nil {
		internal.Logger.Warningf("Invalid id = [%s]", err)
		BadRequest(w)
		return
	}

	newPost, err := p.parseBodyPost(w, r)
	if err != nil {
		internal.Logger.Warningf("Bad body: %s\n", err)
		BadRequest(w)
		return
	}

	err = p.Repo.UpdateById(p.Ctx, newPost, id)
	if err != nil {
		internal.Logger.Errorf("Something wrong %s\n", err)
		InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	internal.Logger.Infof("User-id was successful update, id = %d\n", id)
}

func (p *Posts) DeletePostByID(w http.ResponseWriter, r *http.Request) {
	id, err := p.findPostById(r)
	if err != nil {
		BadRequest(w)
		internal.Logger.Warningf("Invalid id = [%s]", err)
		return
	}

	err = p.Repo.DeleteById(p.Ctx, id)
	if err != nil {
		InternalServerError(w)
		internal.Logger.Errorln(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	internal.Logger.Infof("Post was succesful delete, id=%d", id)
}

func (p *Posts) parseBodyPost(w http.ResponseWriter, r *http.Request) (*repository.PostInfo, error) {
	arr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body.Close()
	var newPost repository.PostInfo
	err = json.Unmarshal(arr, &newPost)
	if err != nil {
		return nil, err
	}
	return &newPost, nil
}

func (p *Posts) findPostById(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return -1, errors.New("bad post id")
	}
	num, err := strconv.Atoi(id)
	if err != nil {
		return -1, err
	}
	if id == "" {
		return -1, errors.New("bad post id")
	}

	return num, nil
}

func (p *Posts) getBytes(postInfo ...*repository.PostInfo) ([]byte, error) {
	retBytes, err := json.Marshal(postInfo)
	if err != nil {
		internal.Logger.Errorln(err)
		return retBytes, err
	}
	return retBytes, nil
}
