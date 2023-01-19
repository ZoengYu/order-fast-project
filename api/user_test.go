package api

// func TestGetUserAPI(t *testing.T) {
// 	user := randomUser()

// 	testCases := []struct {
// 		name          string
// 		accountID     int64
// 		buildStubs    func(mock_db *mockdb.MockDBService)
// 		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name:      "OK",
// 			accountID: account.ID,
// 			buildStubs: func(mockdb *mockdb.MockDBService) {
// 				mockdb.EXPECT().
// 					GetUser(gomock.Any(), gomock.Eq(account.ID)).
// 					Times(1).
// 					Return(user, nil)
// 			},
// 			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recoder.Code)
// 			},
// 		},
// 		{
// 			name:      "NotFound",
// 			accountID: account.ID,
// 			buildStubs: func(mockdb *mockdb.MockDBService) {
// 				mockdb.EXPECT().
// 					GetUser(gomock.Any(), gomock.Eq(account.ID)).
// 					Times(1).
// 					Return(db.User{}, sql.ErrNoRows)
// 			},
// 			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recoder.Code)
// 			},
// 		},
// 		{
// 			name:      "InvalidID",
// 			accountID: 0,
// 			buildStubs: func(mockdb *mockdb.MockDBService) {
// 				mockdb.EXPECT().
// 					GetUser(gomock.Any(), gomock.Any()).
// 					Times(0)
// 			},
// 			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recoder.Code)
// 			},
// 		},
// 		{
// 			name:      "DBError",
// 			accountID: account.ID,
// 			buildStubs: func(mockdb *mockdb.MockDBService) {
// 				mockdb.EXPECT().
// 					GetUser(gomock.Any(), gomock.Eq(account.ID)).
// 					Times(1).
// 					Return(db.User{}, sql.ErrConnDone)
// 			},
// 			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recoder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
// 			mockdb_service := mock_db.NewMockDBService(ctrl)
// 			tc.buildStubs(mockdb_service)
// 			server := newTestServer(t, mockdb_service)
// 			recorder := httptest.NewRecorder()
// 			url := fmt.Sprintf("/v1/account/%d", tc.accountID)
// 			request, err := http.NewRequest(http.MethodGet, url, nil)
// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(t, recorder)
// 		})
// 	}
// }

// func TestCreateUserAPI(t *testing.T) {
// 	user := randomUser()

// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		buildStubs    func(store *mockdb.MockDBService)
// 		checkResponse func(recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"owner": user.Username,
// 			},
// 			buildStubs: func(store *mockdb.MockDBService) {

// 				store.EXPECT().
// 					CreateUser(gomock.Any(), gomock.Eq(user.Username)).
// 					Times(1).
// 					Return(user, nil)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "DBError",
// 			body: gin.H{
// 				"owner": user.Username,
// 			},
// 			buildStubs: func(store *mockdb.MockDBService) {

// 				store.EXPECT().
// 					CreateUser(gomock.Any(), gomock.Eq(user.Username)).
// 					Times(1).
// 					Return(db.User{}, sql.ErrConnDone)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockDBService(ctrl)
// 			tc.buildStubs(store)

// 			server := newTestServer(t, store)
// 			recorder := httptest.NewRecorder()

// 			// Marshal body data to JSON
// 			data, err := json.Marshal(tc.body)
// 			require.NoError(t, err)

// 			url := "/v1/account"
// 			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(recorder)
// 		})
// 	}
// }
