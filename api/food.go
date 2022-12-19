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

func (server *Server) addMenuFood(ctx *gin.Context){
	var req CreateMenuFoodRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMenuFoodParams{
		MenuID: 	req.MenuID,
		Name: 		req.FoodName,
		Price: 		req.FoodPrice,
	}
	menu_food, err := server.db_service.CreateMenuFood(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	if req.FoodTag != nil{
		for i := range(req.FoodTag) {
			fmt.Printf(req.FoodTag[i])
			arg := db.CreateMenuFoodTagParams{
				FoodID: menu_food.ID,
				FoodTag: req.FoodTag[i],
			}
			_, err := server.db_service.CreateMenuFoodTag(ctx, arg)
			if err != nil {
				err = fmt.Errorf("food tag %s cannot be created", arg.FoodTag)
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			}
		}
	}
	ctx.JSON(http.StatusOK, menu_food.ID)
}
