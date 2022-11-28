package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createStoreRequest struct {
	Name string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Owner string `json:"owner" binding:"required"`
	Manager string `json:"manager"`
}

func (server *Server) createStore(ctx *gin.Context){
	var req createStoreRequest

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateStoreParams{
		StoreName: req.Name,
		StoreAddress: req.Address,
		StorePhone: req.Phone,
		StoreOwner: req.Owner,
		StoreManager: req.Manager,
	}
	store, err := server.db_service.CreateStore(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, store.ID)
}

type getStoreRequest struct {
	StoreName 	string 	`json:"name" binding:"required"`
}

func (server *Server) getStore(ctx *gin.Context){
	var req getStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	store, err := server.db_service.GetStoreByName(ctx, req.StoreName)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("cannot find store name: %s", req.StoreName)
		}
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, store)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
