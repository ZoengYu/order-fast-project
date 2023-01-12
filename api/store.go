package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createStoreRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Owner   string `json:"owner" binding:"required"`
	Manager string `json:"manager"`
}

func (server *Server) createStore(ctx *gin.Context) {
	var req createStoreRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateStoreParams{
		StoreName:    req.Name,
		StoreAddress: req.Address,
		StorePhone:   req.Phone,
		StoreOwner:   req.Owner,
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
	StoreID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getStore(ctx *gin.Context) {
	var req getStoreRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	store, err := server.db_service.GetStore(ctx, req.StoreID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("cannot find store id: %d", req.StoreID)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, store)
}

type getStoreByNameRequest struct {
	StoreName string `json:"name" binding:"required"`
}

func (server *Server) getStoreByName(ctx *gin.Context) {
	var req getStoreByNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	store, err := server.db_service.GetStoreByName(ctx, req.StoreName)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("cannot find store name: %s", req.StoreName)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, store)
}

type updateStoreRequest struct {
	StoreID      int64  `json:"store_id" binding:"required"`
	StoreName    string `json:"store_name" binding:"required"`
	StoreAddress string `json:"store_address" binding:"required"`
	StorePhone   string `json:"store_phone" binding:"required"`
	StoreOwner   string `json:"store_owner" binding:"required"`
	StoreManager string `json:"store_manager"`
}

func (server *Server) updateStore(ctx *gin.Context) {
	var req updateStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	store, err := server.db_service.GetStore(ctx, req.StoreID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("store %s is not exist", req.StoreName)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdateStoreParams{
		ID:           store.ID,
		StoreName:    req.StoreName,
		StoreAddress: req.StoreAddress,
		StorePhone:   req.StorePhone,
		StoreOwner:   req.StoreOwner,
		StoreManager: req.StoreManager,
	}
	updated_store, err := server.db_service.UpdateStore(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updated_store.ID)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

type delStoreRequest struct {
	StoreID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) delStore(ctx *gin.Context) {
	var req delStoreRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.db_service.DeleteStore(ctx, req.StoreID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("cannot find store id: %d", req.StoreID)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
