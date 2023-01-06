package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
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
	existed_food := randomMenuFood(menu)
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
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(menu, nil)
				mockdb.EXPECT().
					ListAllMenuFood(gomock.Any(), menu.ID).
					Times(1).Return([]db.Food{existed_food}, nil)
				mockdb.EXPECT().
					CreateMenuFood(gomock.Any(), gomock.Eq(arg)).
					Times(1).Return(food, nil)
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
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "MenuNotFound",
			body: gin.H{
				"menu_id":		menu.ID,
				"name": 		food.Name,
				"price":		food.Price,
				"tag":			[]string{food_tag.FoodTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(db.Menu{}, sql.ErrNoRows)
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
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(db.Menu{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "DuplicatedFoodOnSameMenuReturn422",
			body: gin.H{
				"menu_id":		menu.ID,
				"name": 		existed_food.Name,
				"price":		existed_food.Price,
				"tag":			[]string{food_tag.FoodTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(menu, nil)
				mockdb.EXPECT().
					ListAllMenuFood(gomock.Any(), menu.ID).
					Times(1).Return([]db.Food{existed_food}, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusUnprocessableEntity, recoder.Code)
			},
		},
		{
			name: "DuplicatedFoodOnSameMenuShouldReturn422",
			body: gin.H{
				"menu_id":		menu.ID,
				"name": 		existed_food.Name,
				"price":		existed_food.Price,
				"tag":			[]string{food_tag.FoodTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(menu, nil)
				mockdb.EXPECT().
					ListAllMenuFood(gomock.Any(), menu.ID).
					Times(1).Return([]db.Food{existed_food}, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusUnprocessableEntity, recoder.Code)
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

			url := "/v1/store/menu/food"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestDelMenuFoodAPI(t *testing.T) {
	store := randomStore()
	menu := randomStoreMenu(store)
	food := randomMenuFood(menu)

	testCases := []struct {
		name 			string
		foodID			int64
		buildStubs 		func(mock_db *mockdb.MockDBService)
		checkResponse 	func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			foodID: food.ID,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				GetFood(gomock.Any(), gomock.Eq(food.ID)).
					Times(1).Return(food, nil)
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(menu, nil)
				arg := db.DeleteMenuFoodParams{
					ID:		food.ID,
					MenuID: menu.ID,
				}
				mockdb.EXPECT().
					DeleteMenuFood(gomock.Any(), gomock.Eq(arg)).
					Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNoContent, recoder.Code)
			},
		},
		{
			name: "NotFound",
			foodID: food.ID,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				GetFood(gomock.Any(), gomock.Eq(food.ID)).
					Times(1).Return(db.Food{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "InvalidID",
			foodID: 0,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				DeleteMenu(gomock.Any(), gomock.Any()).
				Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "DBError",
			foodID: food.ID,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				GetFood(gomock.Any(), gomock.Eq(food.ID)).
				Times(1).Return(db.Food{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "HaveFoodInDBButNotHaveMenuShouldReturn500",
			foodID: food.ID,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				GetFood(gomock.Any(), gomock.Eq(food.ID)).
					Times(1).Return(food, nil)
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(db.Menu{}, sql.ErrNoRows)
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
			url := fmt.Sprintf("/v1/store/menu/food/%d", tc.foodID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListMenuFoodAPI(t *testing.T){
	store := randomStore()
	menu := randomStoreMenu(store)

	n := 10
	food_list := make([]db.Food, n)
	for i := 0; i < n; i++ {
		food_list[i] = randomMenuFood(menu)
	}

	type ListQuery struct{
		menuID		int64
		pageID		int
		pageSize	int
	}

	testCases := []struct {
		name 			string
		query			ListQuery
		buildStubs 		func(mock_db *mockdb.MockDBService)
		checkResponse 	func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: ListQuery{
				menuID: 	menu.ID,
				pageID:		1,
				pageSize:	n/2,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.ListMenuFoodParams{
					MenuID: menu.ID,
					Limit:	int32(n/2),
					Offset: 0,
				}
				mockdb.EXPECT().
				ListMenuFood(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(food_list, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: ListQuery{
				menuID: 	menu.ID,
				pageID:		-1,
				pageSize:	n,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				ListMenuFood(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: ListQuery{
				menuID: 	menu.ID,
				pageID:		1,
				pageSize:	20,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				ListMenuFood(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "NotFoundReturnStatusOK",
			query: ListQuery{
				menuID: 	menu.ID,
				pageID:		1,
				pageSize:	n,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.ListMenuFoodParams{
					MenuID: menu.ID,
					Limit:	int32(n),
					Offset: 0,
				}
				mockdb.EXPECT().
				ListMenuFood(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return([]db.Food{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErr",
			query: ListQuery{
				menuID: 	menu.ID,
				pageID:		1,
				pageSize:	n,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.ListMenuFoodParams{
					MenuID: menu.ID,
					Limit:	int32(n),
					Offset: 0,
				}
				mockdb.EXPECT().
				ListMenuFood(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return([]db.Food{}, sql.ErrConnDone)
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
				url := "/v1/store/menu/list_items"
				request, err := http.NewRequest(http.MethodGet, url, nil)
				require.NoError(t, err)

				q := request.URL.Query()
				q.Add("menu_id", fmt.Sprintf("%d", tc.query.menuID))
				q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
				q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
				request.URL.RawQuery = q.Encode()

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