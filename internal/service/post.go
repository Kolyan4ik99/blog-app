package service

import (
	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
)

type PostInterface interface {
	GetPosts() []model.PostInfo
	GetPostByID(id int64) model.PostInfo
	UploadPost(newPost model.PostInfo) (idNewPost int64, err error)
	UpdatePostByID(id int64, newPost model.PostInfo) (updatePost model.PostInfo, err error)
	DeletePostByID(id int64) error
}

type Post struct {
	repo repository.PostInterface
}

func NewPost(repo repository.PostInterface) *Post {
	return &Post{repo: repo}
}

func (p *Post) GetPosts() []model.PostInfo {
	//TODO implement me
	panic("implement me")
}

func (p *Post) GetPostByID(id int64) model.PostInfo {
	//TODO implement me
	panic("implement me")
}

func (p *Post) UploadPost(newPost model.PostInfo) (idNewPost int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (p *Post) UpdatePostByID(id int64, newPost model.PostInfo) (updatePost model.PostInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (p *Post) DeletePostByID(id int64) error {
	//TODO implement me
	panic("implement me")
}

//func (p *Post) GetPosts(c *gin.Context) {
//	postsInfo, err := p.repo.GetAll(p.ctx)
//	if err != nil {
//		InternalServerError(c)
//		logger.Logger.Errorln(err)
//	}
//
//	logger.Logger.Infof("Return posts %s", postsInfo)
//	c.JSON(http.StatusOK, postsInfo)
//}
//
//func (p *Post) GetPostByID(c *gin.Context) {
//	id, err := p.findPostById(c)
//	if err != nil {
//		BadRequest(c)
//		logger.Logger.Warningf("Invalid id = [%s]\n", err)
//		return
//	}
//
//	post, err := p.repo.GetById(p.ctx, id)
//	if errors.Is(err, sql.ErrNoRows) {
//		NotFound(c)
//		logger.Logger.Warningf("Post by id=%d not found\n", id)
//		return
//	}
//	if err != nil {
//		InternalServerError(c)
//		logger.Logger.Errorf("Something wrong %s\n", err)
//		return
//	}
//
//	logger.Logger.Infof("Return post %s", post)
//	c.JSON(http.StatusOK, post)
//}
//
//func (p *Post) UploadPost(c *gin.Context) {
//	// TODO добавить обработку загрузки поста не авторизованным пользователем
//	newPost, err := p.parseBodyPost(c)
//	if err != nil {
//		BadRequest(c)
//		logger.Logger.Warningf("Bad body: %s\n", err)
//		return
//	}
//
//	if err = p.repo.Save(p.ctx, newPost); err != nil {
//		c.JSON(http.StatusNotFound, "Author not exists")
//		logger.Logger.Warningf("Author not exist %s\n", newPost)
//		return
//	}
//
//	logger.Logger.Infof("Post was succesful created: %s", newPost)
//	c.Status(http.StatusCreated)
//}
//
//func (p *Post) UpdatePostByID(c *gin.Context) {
//	id, err := p.findPostById(c)
//	if err != nil {
//		logger.Logger.Warningf("Invalid id = [%s]", err)
//		BadRequest(c)
//		return
//	}
//
//	newPost, err := p.parseBodyPost(c)
//	if err != nil {
//		logger.Logger.Warningf("Bad body: %s\n", err)
//		BadRequest(c)
//		return
//	}
//
//	err = p.repo.UpdateById(p.ctx, newPost, id)
//	if err != nil {
//		logger.Logger.Errorf("Something wrong %s\n", err)
//		InternalServerError(c)
//		return
//	}
//
//	c.Status(http.StatusOK)
//	logger.Logger.Infof("User-id was successful update, id = %d\n", id)
//}
//
//func (p *Post) DeletePostByID(c *gin.Context) {
//	id, err := p.findPostById(c)
//	if err != nil {
//		BadRequest(c)
//		logger.Logger.Warningf("Invalid id = [%s]", err)
//		return
//	}
//
//	err = p.repo.DeleteById(p.ctx, id)
//	if err != nil {
//		InternalServerError(c)
//		logger.Logger.Errorln(err.Error())
//		return
//	}
//
//	c.Status(http.StatusOK)
//	logger.Logger.Infof("Post was succesful delete, id=%d", id)
//}
//
//func (p *Post) parseBodyPost(c *gin.Context) (*model.PostInfo, error) {
//	arr, err := io.ReadAll(c.Request.Body)
//	if err != nil {
//		return nil, err
//	}
//	c.Request.Body.Close()
//	var newPost model.PostInfo
//	err = json.Unmarshal(arr, &newPost)
//	if err != nil {
//		return nil, err
//	}
//	return &newPost, nil
//}
//
//func (p *Post) findPostById(c *gin.Context) (int, error) {
//	strId := c.Param("id")
//	if strId == "" {
//		return -1, errors.New("bad post id")
//	}
//	num, err := strconv.Atoi(strId)
//	if err != nil {
//		return -1, err
//	}
//
//	return num, nil
//}
