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

func (server *Server) createMenu(ctx *gin.Context){
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
