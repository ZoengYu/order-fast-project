package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createMenuRequest struct {
	StoreID int64 	`json:"store_id" binding:"required,min=1"`
	MenuName string `json:"menu_name" binding:"required"`
}

func (server *Server) createStoreMenu(ctx *gin.Context){
	var req createMenuRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	store, err := server.db_service.GetStore(ctx, req.StoreID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("store ID %d is not exist", req.StoreID)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateStoreMenuParams{
		StoreID: store.ID,
		MenuName: req.MenuName,
	}
	menu, err := server.db_service.CreateStoreMenu(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, menu.ID)
}

type getStoreMenuRequest struct {
	StoreID int64 	`json:"store_id" binding:"required,min=1"`
	MenuID 	int64 	`json:"menu_id" binding:"required,min=1"`
}

func (server *Server) getStoreMenu(ctx *gin.Context) {
	var req getStoreMenuRequest

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	store, err := server.db_service.GetStore(ctx, req.StoreID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("store ID %d is not exist", req.StoreID)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.GetStoreMenuParams{
		StoreID:	store.ID,
		ID:			req.MenuID,
	}
	menu, err := server.db_service.GetStoreMenu(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("store menu ID %d is not exist", req.MenuID)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, menu)
}

type updateStoreMenuRequest struct {
	StoreID 	int64 	`json:"store_id" binding:"required,min=1"`
	MenuID 		int64 	`json:"menu_id" binding:"required,min=1"`
	MenuName 	string	`json:"menu_name" binding:"required"`
}

func (server *Server) updateStoreMenu(ctx *gin.Context){
	var req updateStoreMenuRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	store, err := server.db_service.GetStore(ctx, req.StoreID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("store ID %d is not exist", req.StoreID)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.UpdateStoreMenuParams{
		StoreID:	store.ID,
		ID:			req.MenuID,
		MenuName:	req.MenuName,
	}
	menu, err := server.db_service.UpdateStoreMenu(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("store menu id %d is not exist", req.MenuID)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, menu)
}

type delStoreMenuRequest struct{
	MenuID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteStoreMenu(ctx *gin.Context){
	var req delStoreMenuRequest
	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.db_service.DeleteMenu(ctx, req.MenuID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("cannot find menu id: %d", req.MenuID)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
