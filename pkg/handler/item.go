package handler

import (
	todo_app "github.com/bakhodur-nazriev/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid id param")
		return
	}

	var input todo_app.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid id param")
		return
	}

	var input todo_app.UpdateItemInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteItem(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, statusResponse{
		"ok",
	})
}
