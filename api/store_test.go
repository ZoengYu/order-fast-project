package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mock_db "github.com/ZoengYu/order-fast-project/db/mock"
	mockdb "github.com/ZoengYu/order-fast-project/db/mock"
	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetStoreAPI(t *testing.T) {
	store := randomStore()

	testCases := []struct {
		name 			string
		storeID			int64
		buildStubs 		func(mock_db *mockdb.MockDBService)
		checkResponse 	func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			storeID: store.ID,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "NotFound",
			storeID: store.ID,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(db.Store{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "InvalidID",
			storeID: 0,
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
			storeID: store.ID,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
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
			url := fmt.Sprintf("/v1/store/%d", tc.storeID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetStoreByNameAPI(t *testing.T) {
	store := randomStore()

	testCases := []struct {
		name 			string
		body 			[]byte
		buildStubs 		func(mock_db *mockdb.MockDBService)
		checkResponse 	func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body : []byte(fmt.Sprintf(`{"name": "%s"}`, store.StoreName)),
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStoreByName(gomock.Any(), gomock.Eq(store.StoreName)).
					Times(1).
					Return(store, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "NotFound",
			body : []byte(fmt.Sprintf(`{"name": "%s"}`, store.StoreName)),
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStoreByName(gomock.Any(), gomock.Eq(store.StoreName)).
					Times(1).
					Return(db.Store{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "BadRequestPayload",
			body : []byte(fmt.Sprintf(`{"wrong": "%s"}`, store.StoreName)),
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStoreByName(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "DBError",
			body : []byte(fmt.Sprintf(`{"name": "%s"}`, store.StoreName)),
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStoreByName(gomock.Any(), gomock.Eq(store.StoreName)).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
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
			url := "/v1/store/name"
			request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(tc.body))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreateStoreAPI(t *testing.T) {
	store := randomStore()
	testCases := []struct {
		name			string
		body 			gin.H
		buildStubs 		func(mock_db *mockdb.MockDBService)
		checkResponse 	func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name": 	store.StoreName,
				"address": 	store.StoreAddress,
				"phone": 	store.StorePhone,
				"owner": 	store.StoreOwner,
				"manager": 	store.StoreManager,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.CreateStoreParams{
					StoreName: 		store.StoreName,
					StoreAddress: 	store.StoreAddress,
					StorePhone: 	store.StorePhone,
					StoreOwner: 	store.StoreOwner,
					StoreManager: 	store.StoreManager,
				}
				mockdb.EXPECT().
					CreateStore(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(store, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "BadRequestPayload",
			body: gin.H{
				"address": 	store.StoreAddress,
				"phone": 	store.StorePhone,
				"owner": 	store.StoreOwner,
				"manager": 	store.StoreManager,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					CreateStore(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErr",
			body: gin.H{
				"name": 	store.StoreName,
				"address": 	store.StoreAddress,
				"phone": 	store.StorePhone,
				"owner": 	store.StoreOwner,
				"manager": 	store.StoreManager,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				arg := db.CreateStoreParams{
					StoreName: 		store.StoreName,
					StoreAddress: 	store.StoreAddress,
					StorePhone: 	store.StorePhone,
					StoreOwner: 	store.StoreOwner,
					StoreManager: 	store.StoreManager,
				}
				mockdb.EXPECT().
					CreateStore(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
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

			url := "/v1/store"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func TestUpdateStoreAPI(t *testing.T) {
	store := randomStore()
	updated_store := randomStore()
	updated_store.ID = store.ID
	testCases := []struct {
		name			string
		body 			gin.H
		buildStubs 		func(mock_db *mockdb.MockDBService)
		checkResponse 	func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"store_id":			store.ID,
				"store_name": 		updated_store.StoreName,
				"store_address": 	updated_store.StoreAddress,
				"store_phone": 		updated_store.StorePhone,
				"store_owner": 		updated_store.StoreOwner,
				"store_manager": 	updated_store.StoreManager,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				updated_arg := db.UpdateStoreParams{
					ID:				updated_store.ID,
					StoreName: 		updated_store.StoreName,
					StoreAddress: 	updated_store.StoreAddress,
					StorePhone: 	updated_store.StorePhone,
					StoreOwner: 	updated_store.StoreOwner,
					StoreManager: 	updated_store.StoreManager,
				}
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil)

				mockdb.EXPECT().
					UpdateStore(gomock.Any(), gomock.Eq(updated_arg)).
					Times(1).
					Return(updated_store, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "BadRequestPayload",
			body: gin.H{
				"address": 	store.StoreAddress,
				"phone": 	store.StorePhone,
				"owner": 	store.StoreOwner,
				"manager": 	store.StoreManager,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					UpdateStore(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "NotFound",
			body: gin.H{
				"store_id":			updated_store.ID,
				"store_name": 		updated_store.StoreName,
				"store_address": 	updated_store.StoreAddress,
				"store_phone": 		updated_store.StorePhone,
				"store_owner": 		updated_store.StoreOwner,
				"store_manager": 	updated_store.StoreManager,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErrGetStore",
			body: gin.H{
				"store_id":			updated_store.ID,
				"store_name": 		updated_store.StoreName,
				"store_address": 	updated_store.StoreAddress,
				"store_phone": 		updated_store.StorePhone,
				"store_owner": 		updated_store.StoreOwner,
				"store_manager": 	updated_store.StoreManager,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "UnexpectedDBErrUpdateStore",
			body: gin.H{
				"store_id":			updated_store.ID,
				"store_name": 		updated_store.StoreName,
				"store_address": 	updated_store.StoreAddress,
				"store_phone": 		updated_store.StorePhone,
				"store_owner": 		updated_store.StoreOwner,
				"store_manager": 	updated_store.StoreManager,
			},
			buildStubs: func(mockdb *mockdb.MockDBService){
				updated_arg := db.UpdateStoreParams{
					ID:				updated_store.ID,
					StoreName: 		updated_store.StoreName,
					StoreAddress: 	updated_store.StoreAddress,
					StorePhone: 	updated_store.StorePhone,
					StoreOwner: 	updated_store.StoreOwner,
					StoreManager: 	updated_store.StoreManager,
				}
				mockdb.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil)

				mockdb.EXPECT().
					UpdateStore(gomock.Any(), gomock.Eq(updated_arg)).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
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

			url := "/v1/store"
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestDelStoreAPI(t *testing.T) {
	store := randomStore()

	testCases := []struct {
		name 			string
		storeID			int64
		buildStubs 		func(mock_db *mockdb.MockDBService)
		checkResponse 	func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			storeID: store.ID,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				DeleteStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNoContent, recoder.Code)
			},
		},
		{
			name: "NotFound",
			storeID: store.ID,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					DeleteStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "InvalidID",
			storeID: 0,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				DeleteStore(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "DBError",
			storeID: store.ID,
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
				DeleteStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(sql.ErrConnDone)
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
			url := fmt.Sprintf("/v1/store/%d", tc.storeID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomStore() db.Store {
	return db.Store{
		ID:	1,
		StoreName: util.RandomStoreName(),
		StoreAddress: util.RandomStoreAddress(),
		StorePhone: util.RandomPhone(),
		StoreOwner: util.RandomOwner(),
		StoreManager: util.RandomManager(),
		CreatedAt: time.Now(),
	}
}
