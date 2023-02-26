package handler

// func TestAddTask(t *testing.T) {
// 	t.Parallel()
// 	type want struct {
// 		status  int
// 		rspFile string
// 	}
// 	tests := map[string]struct {
// 		reqFile string
// 		want    want
// 	}{
// 		"ok": {
// 			reqFile: "testdata/add_task/ok_req.golden.json",
// 			want: want{
// 				status:  http.StatusOK,
// 				rspFile: "testdata/add_task/ok_rsp.golden.json",
// 			},
// 		},
// 		"badRequest": {
// 			reqFile: "testdata/add_task/bad_req.golden.json",
// 			want: want{
// 				status:  http.StatusBadRequest,
// 				rspFile: "testdata/add_task/bad_req_rsp.golden.json",
// 			},
// 		},
// 	}
// 	for n, tt := range tests {
// 		tt := tt
// 		t.Run(n, func(t *testing.T) {
// 			t.Parallel()
//
// 			w := httptest.NewRecorder()
// 			r := httptest.NewRequest(
// 				http.MethodPost,
// 				"/tasks",
// 				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
// 			)
//
// 			sut := AddTask{
// 				Store: &store.TaskStore{
// 					Tasks: map[entity.TaskID]*entity.Task{},
// 				},
// 				Validator: validator.New(),
// 			}
// 			sut.ServeHTTP(w, r)
//
// 			resp := w.Result()
// 			testutil.AssertResponse(t, resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile))
// 		})
// 	}
// }
