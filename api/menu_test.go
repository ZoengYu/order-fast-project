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

func TestCreateStoreMenuAPI(t *testing.T) {
	store := randomStore()
	menu := randomStoreMenu(store)
	testCases := []struct {
		name			string
		body 			gin.H
		buildStubs 		func(mock_db *mockdb.MockDBService)
		checkResponse 	func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"store_id":		menu.StoreID,
				"menu_name": 	menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.CreateStoreMenuParams{
					StoreID: 	store.ID,
					MenuName: 	menu.MenuName,
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
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "InvalidID",
			body: gin.H{
				"store_id":		0,
				"menu_name": 	menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					CreateStoreMenu(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "NotFound",
			body: gin.H{
				"store_id":		menu.StoreID,
				"menu_name": 	menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).
					Return(db.Store{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErrCreateMenu",
			body: gin.H{
				"store_id":		menu.StoreID,
				"menu_name": 	menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(menu.ID)).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErrCreateMenu",
			body: gin.H{
				"store_id":		menu.StoreID,
				"menu_name": 	menu.MenuName,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.CreateStoreMenuParams{
					StoreID: 	store.ID,
					MenuName: 	menu.MenuName,
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

			url := "/v1/menu"
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
		name 			string
		menuID			int64
		body			gin.H
		buildStubs 		func(mock_db *mockdb.MockDBService)
		checkResponse 	func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"store_id":	menu.StoreID,
				"menu_id": 	menu.ID,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.GetStoreMenuParams{
					StoreID: 	store.ID,
					ID:			menu.ID,
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
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "NotFound",
			body: gin.H{
				"store_id":	menu.StoreID,
				"menu_id": 	menu.ID,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.GetStoreMenuParams{
					StoreID: 	store.ID,
					ID:			menu.ID,
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
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "InvalidID",
			body: gin.H{
				"store_id":	0,
				"menu_id": 	menu.ID,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				GetStore(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "DBError",
			body: gin.H{
				"store_id":	menu.StoreID,
				"menu_id": 	menu.ID,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.GetStoreMenuParams{
					StoreID: 	store.ID,
					ID:			menu.ID,
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
			url := "/v1/menu"
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomStoreMenu(store db.Store) db.Menu {
	return db.Menu{
		ID:	1,
		StoreID: store.ID,
		MenuName: util.RandomMenuName(),
	}
}
