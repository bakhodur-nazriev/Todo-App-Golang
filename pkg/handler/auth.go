package handler

import (
	"github.com/bakhodur-nazriev/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signIn(ctx *gin.Context) {
	var input todo_app.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONP(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signUp(ctx *gin.Context) {

}
