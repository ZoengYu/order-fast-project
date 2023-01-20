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

type CreateMenuItemRequest struct {
	MenuID    int64    `json:"menu_id" binding:"required,min=1"`
	ItemName  string   `json:"name" binding:"required"`
	ItemPrice int32    `json:"price" binding:"required"`
	ItemTag   []string `json:"tag"`
}

func (server *Server) CreateMenuItem(ctx *gin.Context) {
	var req CreateMenuItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// check menu exist
	menu, err := server.db_service.GetMenu(ctx, req.MenuID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("menu %d does not exist", menu.StoreID)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	store, err := server.db_service.GetStore(ctx, menu.StoreID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Username != store.Owner {
		err := errors.New("the menu doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// list all the item of menu
	menu_items, err := server.db_service.ListAllMenuItem(ctx, menu.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, item := range menu_items {
		if req.ItemName == item.Name {
			err = fmt.Errorf("cannot create item, the item name %s already exist", req.ItemName)
			ctx.JSON(http.StatusUnprocessableEntity, errorResponse(err))
			return
		}
	}

	arg := db.CreateMenuItemParams{
		MenuID: menu.ID,
		Name:   req.ItemName,
		Price:  req.ItemPrice,
	}
	menu_item, err := server.db_service.CreateMenuItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if len(req.ItemTag) > 0 {
		for _, tag := range req.ItemTag {
			arg := db.CreateMenuItemTagParams{
				ItemID:  menu_item.ID,
				ItemTag: tag,
			}
			_, err := server.db_service.CreateMenuItemTag(ctx, arg)
			if err != nil {
				err = fmt.Errorf("item tag %s created fail", arg.ItemTag)
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
	}
	ctx.JSON(http.StatusOK, menu_item.ID)
}

type delMenuItemRequest struct {
	MenuID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) DeleteMenuItem(ctx *gin.Context) {
	var req delMenuItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	item, err := server.db_service.GetItem(ctx, req.MenuID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("cannot find the item ID %d", req.MenuID)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// check menu exist
	menu, err := server.db_service.GetMenu(ctx, item.MenuID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("menu %d does not exist", menu.StoreID)
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	store, err := server.db_service.GetStore(ctx, menu.StoreID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Username != store.Owner {
		err := errors.New("the menu doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.DeleteMenuItemParams{
		ID:     menu.ID,
		MenuID: menu.ID,
	}
	// delete the item
	err = server.db_service.DeleteMenuItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

type ListMenuItemsRequest struct {
	MenuID   int64 `form:"menu_id" binding:"required,min=1"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listMenuItems(ctx *gin.Context) {
	var req ListMenuItemsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	menu, err := server.db_service.GetMenu(ctx, req.MenuID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("menu %d does not exist", menu.StoreID)
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	store, err := server.db_service.GetStore(ctx, menu.StoreID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Username != store.Owner {
		err := errors.New("the menu doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.ListMenuItemParams{
		MenuID: req.MenuID,
		Limit:  req.PageSize,
		Offset: calculate_offset(req.PageID, req.PageSize),
	}
	menu_item, err := server.db_service.ListMenuItem(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			menu_item = []db.Item{}
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	ctx.JSON(http.StatusOK, menu_item)
}

type updateMenuItemRequest struct {
	ItemID   int64  `json:"item_id" binding:"required,min=1"`
	ItemName string `json:"item_name" binding:"required"`
	Price    int32  `json:"price" binding:"required"`
}

func (server *Server) updateMenuItem(ctx *gin.Context) {
	var req updateMenuItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	item, err := server.db_service.GetItem(ctx, req.ItemID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("item does not exist")
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	menu, err := server.db_service.GetMenu(ctx, item.MenuID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	store, err := server.db_service.GetStore(ctx, menu.StoreID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Username != store.Owner {
		err := errors.New("the menu doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	arg := db.UpdateMenuItemParams{
		ID:    item.ID,
		Name:  req.ItemName,
		Price: req.Price,
	}
	item, err = server.db_service.UpdateMenuItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, item)
}
