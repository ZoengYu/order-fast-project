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

func TestCreateStoreMenuAPI(t *testing.T) {
	store := randomStore()
	menu := randomStoreMenu(store)
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(mock_db *mockdb.MockDBService)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"store_id":  menu.StoreID,
				"menu_name": menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.CreateStoreMenuParams{
					StoreID:  store.ID,
					MenuName: menu.MenuName,
				}
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).
					Return(store, nil)
				mockdb.EXPECT().
					CreateStoreMenu(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(menu, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "InvalidID",
			body: gin.H{
				"store_id":  0,
				"menu_name": menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					CreateStoreMenu(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "NotFound",
			body: gin.H{
				"store_id":  menu.StoreID,
				"menu_name": menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).
					Return(db.Store{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErrCreateMenu",
			body: gin.H{
				"store_id":  menu.StoreID,
				"menu_name": menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErrCreateMenu",
			body: gin.H{
				"store_id":  menu.StoreID,
				"menu_name": menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.CreateStoreMenuParams{
					StoreID:  store.ID,
					MenuName: menu.MenuName,
				}
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).
					Return(store, nil)
				mockdb.EXPECT().
					CreateStoreMenu(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.Menu{}, sql.ErrConnDone)
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

			url := "/v1/store/menu"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func TestGetStoreMenuAPI(t *testing.T) {
	store := randomStore()
	menu := randomStoreMenu(store)

	testCases := []struct {
		name          string
		menuID        int64
		body          gin.H
		buildStubs    func(mock_db *mockdb.MockDBService)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"store_id": menu.StoreID,
				"menu_id":  menu.ID,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.GetStoreMenuParams{
					StoreID: store.ID,
					ID:      menu.ID,
				}
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil)
				mockdb.EXPECT().
					GetStoreMenu(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(menu, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "NotFound",
			body: gin.H{
				"store_id": menu.StoreID,
				"menu_id":  menu.ID,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.GetStoreMenuParams{
					StoreID: store.ID,
					ID:      menu.ID,
				}
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil)
				mockdb.EXPECT().
					GetStoreMenu(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.Menu{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "InvalidID",
			body: gin.H{
				"store_id": 0,
				"menu_id":  menu.ID,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "DBError",
			body: gin.H{
				"store_id": menu.StoreID,
				"menu_id":  menu.ID,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.GetStoreMenuParams{
					StoreID: store.ID,
					ID:      menu.ID,
				}
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil)
				mockdb.EXPECT().
					GetStoreMenu(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.Menu{}, sql.ErrConnDone)
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
			url := "/v1/store/menu"
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestUpdateStoreMenuAPI(t *testing.T) {
	store := randomStore()
	menu := randomStoreMenu(store)
	update_menu := randomStoreMenu(store)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(mock_db *mockdb.MockDBService)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"store_id":  store.ID,
				"menu_id":   menu.ID,
				"menu_name": update_menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				updated_arg := db.UpdateStoreMenuParams{
					StoreID:  store.ID,
					ID:       menu.ID,
					MenuName: update_menu.MenuName,
				}
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil)

				mockdb.EXPECT().
					UpdateStoreMenu(gomock.Any(), gomock.Eq(updated_arg)).
					Times(1).
					Return(update_menu, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "BadRequestPayload",
			body: gin.H{
				"store_id":  "15",
				"menu_id":   menu.ID,
				"menu_name": update_menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "NotFound",
			body: gin.H{
				"store_id":  store.ID,
				"menu_id":   menu.ID,
				"menu_name": update_menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				updated_arg := db.UpdateStoreMenuParams{
					StoreID:  store.ID,
					ID:       menu.ID,
					MenuName: update_menu.MenuName,
				}
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil)

				mockdb.EXPECT().
					UpdateStoreMenu(gomock.Any(), gomock.Eq(updated_arg)).
					Times(1).
					Return(db.Menu{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErrGetStore",
			body: gin.H{
				"store_id":  store.ID,
				"menu_id":   menu.ID,
				"menu_name": update_menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				updated_arg := db.UpdateStoreMenuParams{
					StoreID:  store.ID,
					ID:       menu.ID,
					MenuName: update_menu.MenuName,
				}
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil)

				mockdb.EXPECT().
					UpdateStoreMenu(gomock.Any(), gomock.Eq(updated_arg)).
					Times(1).
					Return(db.Menu{}, sql.ErrConnDone)
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

			url := "/v1/store/menu"
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestDelMenuAPI(t *testing.T) {
	store := randomStore()
	menu := randomStoreMenu(store)

	testCases := []struct {
		name          string
		menuID        int64
		buildStubs    func(mock_db *mockdb.MockDBService)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			menuID: menu.ID,
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					DeleteMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recoder.Code)
			},
		},
		{
			name:   "NotFound",
			menuID: menu.ID,
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					DeleteMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name:   "InvalidID",
			menuID: 0,
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
			menuID: menu.ID,
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					DeleteMenu(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).
					Return(sql.ErrConnDone)
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
			url := fmt.Sprintf("/v1/store/menu/%d", tc.menuID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListMenuAPI(t *testing.T) {
	store := randomStore()
	n := 5
	menu_list := make([]db.Menu, n)
	for i := 0; i < n; i++ {
		menu_list[i] = randomStoreMenu(store)
	}

	type ListQuery struct {
		storeID  int64
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
				storeID:  store.ID,
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.ListStoreMenuParams{
					StoreID: store.ID,
					Limit:   int32(n),
					Offset:  0,
				}
				mockdb.EXPECT().
					ListStoreMenu(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(menu_list, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: ListQuery{
				storeID:  store.ID,
				pageID:   -1,
				pageSize: n,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					ListStoreMenu(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: ListQuery{
				storeID:  store.ID,
				pageID:   1,
				pageSize: 20,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				mockdb.EXPECT().
					ListStoreMenu(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "NotFoundReturnStatusOK",
			query: ListQuery{
				storeID:  store.ID,
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.ListStoreMenuParams{
					StoreID: store.ID,
					Limit:   int32(n),
					Offset:  0,
				}
				mockdb.EXPECT().
					ListStoreMenu(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(menu_list, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErr",
			query: ListQuery{
				storeID:  store.ID,
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(mockdb *mockdb.MockDBService) {
				arg := db.ListStoreMenuParams{
					StoreID: store.ID,
					Limit:   int32(n),
					Offset:  0,
				}
				mockdb.EXPECT().
					ListStoreMenu(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(menu_list, sql.ErrConnDone)
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
			url := "/v1/store/menu_list"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			q := request.URL.Query()
			q.Add("store_id", fmt.Sprintf("%d", tc.query.storeID))
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomStoreMenu(store db.Store) db.Menu {
	return db.Menu{
		ID:       1,
		StoreID:  store.ID,
		MenuName: util.RandomMenuName(),
	}
}
