package http

import (
	"awesomeProject3/internal/domain"
	"awesomeProject3/internal/repositories"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type PostHandler interface {
	GetPosts(c echo.Context) error
	GetPost(c echo.Context) error
	StorePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}

type PostHandlerImpl struct {
	Repo repositories.PostRepository
}

func NewPostHandler(r repositories.PostRepository) *PostHandlerImpl {
	return &PostHandlerImpl{Repo: r}
}

func (h *PostHandlerImpl) GetPosts(c echo.Context) error {
	posts, err := h.Repo.GetPosts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *PostHandlerImpl) GetPost(c echo.Context) error {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	post, err := h.Repo.GetPost(intId)
	if err != nil {
		if err.Error() == fmt.Sprintf("пост с ID %d не найден", intId) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}

func (h *PostHandlerImpl) StorePost(c echo.Context) error {
	var post domain.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	storePost, err := h.Repo.StorePost(post)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, storePost)
}

func (h *PostHandlerImpl) UpdatePost(c echo.Context) error {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var post domain.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	updatePost, err := h.Repo.UpdatePost(post, intId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updatePost)
}

func (h *PostHandlerImpl) DeletePost(c echo.Context) error {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.Repo.DeletePost(intId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
