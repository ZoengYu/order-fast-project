package api

import (
	"database/sql"
	"net/http"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateMenuFoodRequest struct {
	MenuID 		int64 	`json:"menu_id" binding:"required,min=1"`
	FoodName	string	`json:"food_name" binding:"required"`
}

func (server *Server) addMenuFood(ctx *gin.Context){
	var req CreateMenuFoodRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMenuFoodParams{
		MenuID: 	req.MenuID,
		Name: 		req.FoodName,
	}
	menu_list, err := server.db_service.CreateMenuFood(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	ctx.JSON(http.StatusOK, menu_list)
}
