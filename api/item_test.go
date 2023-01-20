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

func TestAddMenuItemAPI(t *testing.T) {
	user, _ := randomUser(t)
	store := randomStore(user)
	menu := randomStoreMenu(store)
	existed_item := randomMenuItem(menu)
	item := randomMenuItem(menu)
	item_tag := RandomItemTag(item)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(mock_db *mockdb.MockDBService)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"menu_id": menu.ID,
				"name":    item.Name,
				"price":   item.Price,
				"tag":     []string{item_tag.ItemTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.CreateMenuItemParams{
					MenuID: menu.ID,
					Name:   item.Name,
					Price:  item.Price,
				}
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(menu, nil)
				mockdb.EXPECT().
					ListAllMenuItem(gomock.Any(), menu.ID).
					Times(1).Return([]db.Item{existed_item}, nil)
				mockdb.EXPECT().
					CreateMenuItem(gomock.Any(), gomock.Eq(arg)).
					Times(1).Return(item, nil)
				tag_arg := db.CreateMenuItemTagParams{
					ItemID:  item.ID,
					ItemTag: item_tag.ItemTag,
				}
				mockdb.EXPECT().
					CreateMenuItemTag(gomock.Any(), gomock.Eq(tag_arg)).
					Times(1).
					Return(item_tag, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "InvalidID",
			body: gin.H{
				"menu_id":   0,
				"item_name": item.Name,
				"price":     item.Price,
				"tag":       []string{item_tag.ItemTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "MenuNotFound",
			body: gin.H{
				"menu_id": menu.ID,
				"name":    item.Name,
				"price":   item.Price,
				"tag":     []string{item_tag.ItemTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(db.Menu{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErrCreateMenu",
			body: gin.H{
				"menu_id": menu.ID,
				"name":    item.Name,
				"price":   item.Price,
				"tag":     []string{item_tag.ItemTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(db.Menu{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "DuplicatedItemOnSameMenuReturn422",
			body: gin.H{
				"menu_id": menu.ID,
				"name":    existed_item.Name,
				"price":   existed_item.Price,
				"tag":     []string{item_tag.ItemTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(menu, nil)
				mockdb.EXPECT().
					ListAllMenuItem(gomock.Any(), menu.ID).
					Times(1).Return([]db.Item{existed_item}, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recoder.Code)
			},
		},
		{
			name: "DuplicatedItemOnSameMenuShouldReturn422",
			body: gin.H{
				"menu_id": menu.ID,
				"name":    existed_item.Name,
				"price":   existed_item.Price,
				"tag":     []string{item_tag.ItemTag},
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(menu, nil)
				mockdb.EXPECT().
					ListAllMenuItem(gomock.Any(), menu.ID).
					Times(1).Return([]db.Item{existed_item}, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
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

			url := "/v1/store/menu/item"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestDelMenuItemAPI(t *testing.T) {
	user, _ := randomUser(t)
	store := randomStore(user)
	menu := randomStoreMenu(store)
	item := randomMenuItem(menu)

	testCases := []struct {
		name          string
		ItemID        int64
		buildStubs    func(mock_db *mockdb.MockDBService)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			ItemID: item.ID,
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetItem(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(item, nil)
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(menu, nil)
				arg := db.DeleteMenuItemParams{
					ID:     item.ID,
					MenuID: menu.ID,
				}
				mockdb.EXPECT().
					DeleteMenuItem(gomock.Any(), gomock.Eq(arg)).
					Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recoder.Code)
			},
		},
		{
			name:   "NotFound",
			ItemID: item.ID,
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetItem(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(db.Item{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name:   "InvalidID",
			ItemID: 0,
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					DeleteMenu(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name:   "DBError",
			ItemID: item.ID,
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetItem(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(db.Item{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name:   "HaveItemInDBButNotHaveMenuShouldReturn500",
			ItemID: item.ID,
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetItem(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(item, nil)
				mockdb.EXPECT().
					GetMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).Return(db.Menu{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
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
			url := fmt.Sprintf("/v1/store/menu/item/%d", tc.ItemID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListMenuItemAPI(t *testing.T) {
	user, _ := randomUser(t)
	store := randomStore(user)
	menu := randomStoreMenu(store)

	n := 10
	item_list := make([]db.Item, n)
	for i := 0; i < n; i++ {
		item_list[i] = randomMenuItem(menu)
	}

	type ListQuery struct {
		menuID   int64
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         ListQuery
		buildStubs    func(mock_db *mockdb.MockDBService)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: ListQuery{
				menuID:   menu.ID,
				pageID:   1,
				pageSize: n / 2,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.ListMenuItemParams{
					MenuID: menu.ID,
					Limit:  int32(n / 2),
					Offset: 0,
				}
				mockdb.EXPECT().
					ListMenuItem(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(item_list, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: ListQuery{
				menuID:   menu.ID,
				pageID:   -1,
				pageSize: n,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					ListMenuItem(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: ListQuery{
				menuID:   menu.ID,
				pageID:   1,
				pageSize: 20,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					ListMenuItem(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "NotFoundReturnStatusOK",
			query: ListQuery{
				menuID:   menu.ID,
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.ListMenuItemParams{
					MenuID: menu.ID,
					Limit:  int32(n),
					Offset: 0,
				}
				mockdb.EXPECT().
					ListMenuItem(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return([]db.Item{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErr",
			query: ListQuery{
				menuID:   menu.ID,
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.ListMenuItemParams{
					MenuID: menu.ID,
					Limit:  int32(n),
					Offset: 0,
				}
				mockdb.EXPECT().
					ListMenuItem(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return([]db.Item{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
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

func TestUpdateMenuItemAPI(t *testing.T) {
	user, _ := randomUser(t)
	store := randomStore(user)
	menu := randomStoreMenu(store)
	item := randomMenuItem(menu)
	updated_item := randomMenuItem(menu)
	updated_item.ID = item.ID

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(mock_db *mockdb.MockDBService)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"item_id":   item.ID,
				"item_name": updated_item.Name,
				"price":     updated_item.Price,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				updated_arg := db.UpdateMenuItemParams{
					ID:    item.ID,
					Name:  updated_item.Name,
					Price: updated_item.Price,
				}
				mockdb.EXPECT().
					GetItem(gomock.Any(), gomock.Eq(item.ID)).
					Times(1).
					Return(item, nil)

				mockdb.EXPECT().
					UpdateMenuItem(gomock.Any(), gomock.Eq(updated_arg)).
					Times(1).
					Return(updated_item, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "BadRequestPayload",
			body: gin.H{
				"item_id":   item.ID,
				"item_name": updated_item.Name,
				"price":     "15",
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetItem(gomock.Any(), gomock.Eq(store.ID)).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "NotFound",
			body: gin.H{
				"item_id":   item.ID,
				"item_name": updated_item.Name,
				"price":     updated_item.Price,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				updated_arg := db.UpdateMenuItemParams{
					ID:    item.ID,
					Name:  updated_item.Name,
					Price: updated_item.Price,
				}
				mockdb.EXPECT().
					GetItem(gomock.Any(), gomock.Eq(item.ID)).
					Times(1).
					Return(db.Item{}, sql.ErrNoRows)

				mockdb.EXPECT().
					UpdateStoreMenu(gomock.Any(), gomock.Eq(updated_arg)).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErrGetStore",
			body: gin.H{
				"item_id":   item.ID,
				"item_name": updated_item.Name,
				"price":     updated_item.Price,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				updated_arg := db.UpdateMenuItemParams{
					ID:    item.ID,
					Name:  updated_item.Name,
					Price: updated_item.Price,
				}
				mockdb.EXPECT().
					GetItem(gomock.Any(), gomock.Eq(item.ID)).
					Times(1).
					Return(item, nil)

				mockdb.EXPECT().
					UpdateMenuItem(gomock.Any(), gomock.Eq(updated_arg)).
					Times(1).
					Return(db.Item{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
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

			url := "/v1/store/menu/item"
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomMenuItem(menu db.Menu) db.Item {
	return db.Item{
		ID:     1,
		MenuID: menu.ID,
		Name:   util.RandomItemName(),
		Price:  int32(util.RandomInt(20, 100)),
	}
}

func RandomItemTag(item db.Item) db.ItemTag {
	return db.ItemTag{
		ID:      1,
		ItemID:  item.ID,
		ItemTag: util.RandomItemTag(),
	}
}
