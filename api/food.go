package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateMenuFoodRequest struct {
	MenuID 		int64 		`json:"menu_id" binding:"required,min=1"`
	FoodName	string		`json:"name" binding:"required"`
	FoodPrice	int32		`json:"price" binding:"required"`
	FoodTag		[]string	`json:"tag"`
}

func (server *Server) createMenuFood(ctx *gin.Context){
	var req CreateMenuFoodRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// check menu exist
	menu, err := server.db_service.GetMenu(ctx, req.MenuID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// list all the food of menu
	menu_foods, err := server.db_service.ListMenuFood(ctx, menu.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, food := range(menu_foods) {
		if req.FoodName == food.Name {
			err = fmt.Errorf("cannot create menu, the menu name %s already exist", req.FoodName)
			ctx.JSON(http.StatusUnprocessableEntity, errorResponse(err))
			return
		}
	}

	arg := db.CreateMenuFoodParams{
		MenuID: 	menu.ID,
		Name: 		req.FoodName,
		Price: 		req.FoodPrice,
	}
	menu_food, err := server.db_service.CreateMenuFood(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
		}

	if len(req.FoodTag) > 0{
		for _, tag := range(req.FoodTag) {
			arg := db.CreateMenuFoodTagParams{
				FoodID: menu_food.ID,
				FoodTag: tag,
			}
			_, err := server.db_service.CreateMenuFoodTag(ctx, arg)
			if err != nil {
				err = fmt.Errorf("food tag %s created fail", arg.FoodTag)
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
	}
	ctx.JSON(http.StatusOK, menu_food.ID)
}

type delMenuFoodRequest struct{
	FoodID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteMenuFood(ctx *gin.Context){
	var req delMenuFoodRequest
	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	food, err := server.db_service.GetFood(ctx, req.FoodID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("cannot find the food ID %d", req.FoodID)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// check menu exist
	menu, err := server.db_service.GetMenu(ctx, food.MenuID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("something went wrong, the food exist but the food menu is not")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.DeleteMenuFoodParams{
		ID:		food.ID,
		MenuID: menu.ID,
	}
	// delete the food
	err = server.db_service.DeleteMenuFood(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
