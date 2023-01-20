package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/ZoengYu/order-fast-project/token"
	"github.com/gin-gonic/gin"
)

type createStoreRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Manager string `json:"manager"`
}

func (server *Server) createStore(ctx *gin.Context) {
	var req createStoreRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateStoreParams{
		Owner:   authPayload.Username,
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		Manager: req.Manager,
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Username != store.Owner {
		err := errors.New("store doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, store)
}

type listStoresByNameRequest struct {
	Name     string `form:"name" binding:"required"`
	PageID   int32  `form:"page_id" binding:"required"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listStoresByName(ctx *gin.Context) {
	var req listStoresByNameRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListStoresByNameParams{
		Owner:  authPayload.Username,
		Name:   req.Name,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1),
	}

	stores, err := server.db_service.ListStoresByName(ctx, arg)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	ctx.JSON(http.StatusOK, stores)
}

type updateStoreRequest struct {
	StoreID int64  `json:"store_id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Manager string `json:"manager"`
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
			err = fmt.Errorf("store %s is not exist", req.Name)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Username != store.Owner {
		err := errors.New("store doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.UpdateStoreParams{
		ID:      store.ID,
		Owner:   authPayload.Username,
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		Manager: req.Manager,
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Username != store.Owner {
		err := errors.New("store doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	err = server.db_service.DeleteStore(ctx, req.StoreID)
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
