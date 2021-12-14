package handler

import (
	"fmt"
	"modules/enum"
	"modules/pojo"
	"modules/query"
	"modules/resp"
	"modules/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	USC service.UserService
}

// user object convert to json-send object
func (h *UserHandler) GetEntity(result pojo.User) resp.User {
	return resp.User{
		UserId:   result.UserId,
		UserName: result.UserName,
		Mobile:   result.Mobile,
		Address:  result.Address,
		IsDelete: result.IsDeleted,
		IsLocked: result.IsLocked,
	}
}

func (h *UserHandler) UserInfoHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}

	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	user := pojo.User{
		UserId: userId,
	}
	result, err := h.USC.Get(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	r := h.GetEntity(*result)
	entity = resp.Entity{
		Code:      int(enum.OperateOK),
		Msg:       "OK",
		Total:     0,
		TotalPage: 0,
		Data:      r,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})

}

func (h *UserHandler) UserListHandler(c *gin.Context) {
	var q query.ListQuery
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	result, err := h.USC.List(&q)
	total, err := h.USC.GetTotal(&q)
	fmt.Sprintf("total user:%d", total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
	}
	if q.PageSize == 0 {
		q.PageSize = 5
	}
	ret := int(total % int64(q.PageSize))
	ret2 := int(total / int64(q.PageSize))

	totalPage := 0
	if ret == 0 {
		totalPage = ret2
	} else {
		totalPage = ret2 + 1
	}
	var userList []*resp.User

	for _, item := range result {
		user := h.GetEntity(item)
		userList = append(userList, &user)
	}

	entity = resp.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     int(total),
		TotalPage: totalPage,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *UserHandler) AddUserHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	u := pojo.User{}
	err := c.ShouldBindQuery(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	result, err := h.USC.Add(u)
	if err != nil {
		entity.Msg = err.Error()
		return
	}

	if result.UserId == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	entity.Code = int(enum.OperateOK)
	entity.Msg = enum.OperateOK.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *UserHandler) EditUserHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	u := pojo.User{}
	err := c.ShouldBindQuery(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	result, err := h.USC.Edit(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if result {
		entity.Code = int(enum.OperateOK)
		entity.Msg = enum.OperateOK.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}

	result, err := h.USC.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if result {
		entity.Code = int(enum.OperateOK)
		entity.Msg = enum.OperateOK.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})

	} else {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}
