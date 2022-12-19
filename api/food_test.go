package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_db "github.com/ZoengYu/order-fast-project/db/mock"
	mockdb "github.com/ZoengYu/order-fast-project/db/mock"
	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAddMenuFoodAPI(t *testing.T) {
	store := randomStore()
	menu := randomStoreMenu(store)
	food := randomMenuFood(menu)
	food_tag := randomFoodTag(food)

	testCases := []struct {
		name			string
		body 			gin.H
		buildStubs 		func(mock_db *mockdb.MockDBService)
		checkResponse 	func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"menu_id":		menu.ID,
				"name": 		food.Name,
				"price":		food.Price,
				"tag":			[]string{food_tag.FoodTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.CreateMenuFoodParams{
					MenuID: menu.ID,
					Name: 	food.Name,
					Price: 	food.Price,
				}
				mockdb.EXPECT().
					CreateMenuFood(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(food, nil)
				tag_arg := db.CreateMenuFoodTagParams{
					FoodID: 	food.ID,
					FoodTag: 	food_tag.FoodTag,
				}
				mockdb.EXPECT().
					CreateMenuFoodTag(gomock.Any(), gomock.Eq(tag_arg)).
					Times(1).
					Return(food_tag, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "InvalidID",
			body: gin.H{
				"menu_id":		0,
				"food_name": 	food.Name,
				"price":		food.Price,
				"tag":			[]string{food_tag.FoodTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					CreateMenuFood(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "NotFound",
			body: gin.H{
				"menu_id":		menu.ID,
				"name": 		food.Name,
				"price":		food.Price,
				"tag":			[]string{food_tag.FoodTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.CreateMenuFoodParams{
					MenuID: 	menu.ID,
					Name: 		food.Name,
					Price: 		food.Price,
				}
				mockdb.EXPECT().
					CreateMenuFood(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(food, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErrCreateMenu",
			body: gin.H{
				"menu_id":		menu.ID,
				"name": 		food.Name,
				"price":		food.Price,
				"tag":			[]string{food_tag.FoodTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					CreateMenuFood(gomock.Any(), gomock.Any()).
					Times(1).
					Return(food, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockdb_service := mock_db.NewMockDBService(ctrl)
			tc.buildStubs(mockdb_service)

			server := newTestServer(t, mockdb_service)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/v1/menu/food"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomMenuFood(menu db.Menu) db.Food {
	return db.Food{
		ID:		1,
		MenuID: menu.ID,
		Name: 	util.RandomFoodName(),
		Price: 	int32(util.RandomInt(20, 100)),
	}
}

func randomFoodTag(food db.Food) db.FoodTag {
	return db.FoodTag{
		ID: 1,
		FoodID: food.ID,
		FoodTag: util.RandomFoodTag(),
	}
}
