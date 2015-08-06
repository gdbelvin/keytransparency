// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/google/e2e-key-server/rest/handlers"
	"github.com/gorilla/mux"

	v2pb "github.com/google/e2e-key-server/proto/v2"
	context "golang.org/x/net/context"
)

const (
	valid_ts               = "2015-05-18T23:58:36.000Z"
	invalid_ts             = "Mon May 18 23:58:36 UTC 2015"
	ts_seconds             = 1431993516
	primary_test_epoch     = "2367"
	primary_test_page_size = "653"
	primary_test_sequence  = "8626"
	primary_test_email     = "e2eshare.test@gmail.com"
	primary_test_app_id    = "gmail"
	primary_test_key_id    = "mykey"
)

type fakeJSONParserReader struct {
	*bytes.Buffer
}

func (pr fakeJSONParserReader) Close() error {
	return nil
}

type FakeServer struct {
}

func Fake_Handler(srv interface{}, ctx context.Context, w http.ResponseWriter, r *http.Request, info *handlers.HandlerInfo) error {
	w.Write([]byte("hi"))
	return nil
}

func Fake_Initializer(rInfo handlers.RouteInfo) *handlers.HandlerInfo {
	return nil
}

func Fake_RequestHandler(srv interface{}, ctx context.Context, arg interface{}) (*interface{}, error) {
	b := true
	i := new(interface{})
	*i = b
	return i, nil
}

func TestFoo(t *testing.T) {
	v1 := &FakeServer{}
	s := New(v1)
	rInfo := handlers.RouteInfo{
		"/hi",
		"GET",
		Fake_Initializer,
		Fake_RequestHandler,
	}
	s.AddHandler(rInfo, Fake_Handler, v1)

	server := httptest.NewServer(s.Handlers())
	defer server.Close()
	res, err := http.Get(fmt.Sprintf("%s/hi", server.URL))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res.StatusCode, http.StatusOK; got != want {
		t.Errorf("GET: %v = %v, want %v", res.Request.URL, got, want)
	}
}

func TestGetUserV1_InitiateHandlerInfo(t *testing.T) {
	mx := mux.NewRouter()
	mx.KeepContext = true
	mx.HandleFunc("/v1/users/{"+handlers.USER_ID_KEYWORD+"}", Fake_HTTPHandler)

	i, _ := strconv.ParseUint(primary_test_epoch, 10, 64)
	var tests = []struct {
		path         string
		userId       string
		appId        string
		epoch        uint64
		parserNilErr bool
	}{
		{"/v1/users/" + primary_test_email + "?app_id=" + primary_test_app_id +
			"&epoch=" + primary_test_epoch,
			primary_test_email, primary_test_app_id, i, true},
		{"/v1/users/" + primary_test_email + "?epoch=" + primary_test_epoch,
			primary_test_email, "", i, true},
		{"/v1/users/" + primary_test_email + "?app_id=" + primary_test_app_id,
			primary_test_email, primary_test_app_id, 0, true},
		{"/v1/users/" + primary_test_email, primary_test_email, "", 0, true},
		// Invalid epoch format.
		{"/v1/users/" + primary_test_email + "?epoch=-2587", primary_test_email,
			"", 0, false},
		{"/v1/users/" + primary_test_email + "?epoch=greatepoch", primary_test_email,
			"", 0, false},
	}

	for _, test := range tests {
		rInfo := handlers.RouteInfo{
			test.path,
			"GET",
			Fake_Initializer,
			Fake_RequestHandler,
		}
		// Body is empty when invoking get user API.
		jsonBody := "{}"

		info := GetUserV1_InitializeHandlerInfo(rInfo)

		if _, ok := info.Arg.(*v2pb.GetUserRequest); !ok {
			t.Errorf("info.Arg is not of type v2pb.GetUserRequest")
		}

		r, _ := http.NewRequest(rInfo.Method, rInfo.Path, fakeJSONParserReader{bytes.NewBufferString(jsonBody)})
		mx.ServeHTTP(nil, r)
		err := info.Parser(r, &info.Arg)
		if got, want := (err == nil), test.parserNilErr; got != want {
			t.Errorf("Unexpected parser err = (%v), want nil = %v", err, test.parserNilErr)
		}
		// If there's an error parsing, the test cannot be completed.
		// The parsing error might be expected though.
		if err != nil {
			continue
		}

		// Call JSONDecoder to simulate decoding JSON -> Proto.
		err = JSONDecoder(r, &info.Arg)
		if err != nil {
			t.Errorf("Error while calling JSONDecoder, this should not happen. err: %v", err)
		}

		if got, want := info.Arg.(*v2pb.GetUserRequest).UserId, test.userId; got != want {
			t.Errorf("UserId = %v, want %v", got, want)
		}
		if got, want := info.Arg.(*v2pb.GetUserRequest).AppId, test.appId; got != want {
			t.Errorf("AppId = %v, want %v", got, want)
		}
		if got, want := info.Arg.(*v2pb.GetUserRequest).Epoch, test.epoch; got != want {
			t.Errorf("Epoch = %v, want %v", got, want)
		}

		v1 := &FakeServer{}
		srv := New(v1)
		resp, err := info.H(srv, nil, nil)
		if err != nil {
			t.Errorf("Error while calling Fake_RequestHandler, this should not happen.")
		}
		if got, want := (*resp).(bool), true; got != want {
			t.Errorf("resp = %v, want %v.", got, want)
		}
	}
}

func TestGetUserV2_InitiateHandlerInfo(t *testing.T) {
	mx := mux.NewRouter()
	mx.KeepContext = true
	mx.HandleFunc("/v2/users/{"+handlers.USER_ID_KEYWORD+"}", Fake_HTTPHandler)

	i, _ := strconv.ParseUint(primary_test_epoch, 10, 64)
	var tests = []struct {
		path         string
		userId       string
		appId        string
		epoch        uint64
		parserNilErr bool
	}{
		{"/v2/users/" + primary_test_email + "?app_id=" + primary_test_app_id +
			"&epoch=" + primary_test_epoch,
			primary_test_email, primary_test_app_id, i, true},
		{"/v2/users/" + primary_test_email + "?epoch=" + primary_test_epoch,
			primary_test_email, "", i, true},
		{"/v2/users/" + primary_test_email + "?app_id=" + primary_test_app_id,
			primary_test_email, primary_test_app_id, 0, true},
		{"/v2/users/" + primary_test_email, primary_test_email, "", 0, true},
		// Invalid epoch format.
		{"/v2/users/" + primary_test_email + "?epoch=-2587", primary_test_email,
			"", 0, false},
		{"/v2/users/" + primary_test_email + "?epoch=greatepoch", primary_test_email,
			"", 0, false},
	}

	for _, test := range tests {
		rInfo := handlers.RouteInfo{
			test.path,
			"GET",
			Fake_Initializer,
			Fake_RequestHandler,
		}
		// Body is empty when invoking get user API.
		jsonBody := "{}"

		info := GetUserV2_InitializeHandlerInfo(rInfo)

		if _, ok := info.Arg.(*v2pb.GetUserRequest); !ok {
			t.Errorf("info.Arg is not of type v2pb.GetUserRequest")
		}

		r, _ := http.NewRequest(rInfo.Method, rInfo.Path, fakeJSONParserReader{bytes.NewBufferString(jsonBody)})
		mx.ServeHTTP(nil, r)
		err := info.Parser(r, &info.Arg)
		if got, want := (err == nil), test.parserNilErr; got != want {
			t.Errorf("Unexpected parser err = (%v), want nil = %v", err, test.parserNilErr)
		}
		// If there's an error parsing, the test cannot be completed.
		// The parsing error might be expected though.
		if err != nil {
			continue
		}

		// Call JSONDecoder to simulate decoding JSON -> Proto.
		err = JSONDecoder(r, &info.Arg)
		if err != nil {
			t.Errorf("Error while calling JSONDecoder, this should not happen. err: %v", err)
		}

		if got, want := info.Arg.(*v2pb.GetUserRequest).UserId, test.userId; got != want {
			t.Errorf("UserId = %v, want %v", got, want)
		}
		if got, want := info.Arg.(*v2pb.GetUserRequest).AppId, test.appId; got != want {
			t.Errorf("AppId = %v, want %v", got, want)
		}
		if got, want := info.Arg.(*v2pb.GetUserRequest).Epoch, test.epoch; got != want {
			t.Errorf("Epoch = %v, want %v", got, want)
		}

		v2 := &FakeServer{}
		srv := New(v2)
		resp, err := info.H(srv, nil, nil)
		if err != nil {
			t.Errorf("Error while calling Fake_RequestHandler, this should not happen.")
		}
		if got, want := (*resp).(bool), true; got != want {
			t.Errorf("resp = %v, want %v.", got, want)
		}
	}
}

func TestListUserHistoryV2_InitiateHandlerInfo(t *testing.T) {
	mx := mux.NewRouter()
	mx.KeepContext = true
	mx.HandleFunc("/v2/users/{"+handlers.USER_ID_KEYWORD+"}/history", Fake_HTTPHandler)

	e, _ := strconv.ParseUint(primary_test_epoch, 10, 64)
	ps, _ := strconv.ParseUint(primary_test_page_size, 10, 32)
	var tests = []struct {
		path         string
		userId       string
		startEpoch   uint64
		pageSize     int32
		parserNilErr bool
	}{
		{"/v2/users/" + primary_test_email + "/history?start_epoch=" + primary_test_epoch +
			"&page_size=" + primary_test_page_size,
			primary_test_email, e, int32(ps), true},
		{"/v2/users/" + primary_test_email + "/history?start_epoch=" + primary_test_epoch,
			primary_test_email, e, 0, true},
		{"/v2/users/" + primary_test_email + "/history?page_size=" + primary_test_page_size,
			primary_test_email, 0, int32(ps), true},
		{"/v2/users/" + primary_test_email + "/history", primary_test_email, 0, 0, true},
		// Invalid start_epoch format.
		{"/v2/users/" + primary_test_email + "/history?start_epoch=-2587", primary_test_email,
			0, 0, false},
		{"/v2/users/" + primary_test_email + "/history?start_epoch=greatepoch", primary_test_email,
			0, 0, false},
		// Invalid page_size format.
		{"/v2/users/" + primary_test_email + "/history?page_size=bigpagesize", primary_test_email,
			0, 0, false},
	}

	for _, test := range tests {
		rInfo := handlers.RouteInfo{
			test.path,
			"GET",
			Fake_Initializer,
			Fake_RequestHandler,
		}
		// Body is empty when invoking list user history API.
		jsonBody := "{}"

		info := ListUserHistoryV2_InitializeHandlerInfo(rInfo)

		if _, ok := info.Arg.(*v2pb.ListUserHistoryRequest); !ok {
			t.Errorf("info.Arg is not of type v2pb.ListUserHistoryRequest")
		}

		r, _ := http.NewRequest(rInfo.Method, rInfo.Path, fakeJSONParserReader{bytes.NewBufferString(jsonBody)})
		mx.ServeHTTP(nil, r)
		err := info.Parser(r, &info.Arg)
		if got, want := (err == nil), test.parserNilErr; got != want {
			t.Errorf("Unexpected parser err = (%v), want nil = %v", err, test.parserNilErr)
		}
		// If there's an error parsing, the test cannot be completed.
		// The parsing error might be expected though.
		if err != nil {
			continue
		}

		// Call JSONDecoder to simulate decoding JSON -> Proto.
		err = JSONDecoder(r, &info.Arg)
		if err != nil {
			t.Errorf("Error while calling JSONDecoder, this should not happen. err: %v", err)
		}

		if got, want := info.Arg.(*v2pb.ListUserHistoryRequest).UserId, test.userId; got != want {
			t.Errorf("UserId = %v, want %v", got, want)
		}
		if got, want := info.Arg.(*v2pb.ListUserHistoryRequest).StartEpoch, test.startEpoch; got != want {
			t.Errorf("StartEpoch = %v, want %v", got, want)
		}
		if got, want := info.Arg.(*v2pb.ListUserHistoryRequest).PageSize, test.pageSize; got != want {
			t.Errorf("PageSize = %v, want %v", got, want)
		}

		v2 := &FakeServer{}
		srv := New(v2)
		resp, err := info.H(srv, nil, nil)
		if err != nil {
			t.Errorf("Error while calling Fake_RequestHandler, this should not happen.")
		}
		if got, want := (*resp).(bool), true; got != want {
			t.Errorf("resp = %v, want %v.", got, want)
		}
	}
}

func TestUpdateUserV2_InitiateHandlerInfo(t *testing.T) {
	mx := mux.NewRouter()
	mx.KeepContext = true
	mx.HandleFunc("/v2/users/{"+handlers.USER_ID_KEYWORD+"}", Fake_HTTPHandler)

	var tests = []struct {
		path         string
		userId       string
		parserNilErr bool
	}{
		{"/v2/users/" + primary_test_email, primary_test_email, true},
	}

	for _, test := range tests {
		rInfo := handlers.RouteInfo{
			test.path,
			"PUT",
			Fake_Initializer,
			Fake_RequestHandler,
		}
		// Body is empty because it is irrelevant in this test.
		jsonBody := "{}"

		info := UpdateUserV2_InitializeHandlerInfo(rInfo)

		if _, ok := info.Arg.(*v2pb.UpdateUserRequest); !ok {
			t.Errorf("info.Arg is not of type v2pb.UpdateUserRequest")
		}

		r, _ := http.NewRequest(rInfo.Method, rInfo.Path, fakeJSONParserReader{bytes.NewBufferString(jsonBody)})
		mx.ServeHTTP(nil, r)
		err := info.Parser(r, &info.Arg)
		if got, want := (err == nil), test.parserNilErr; got != want {
			t.Errorf("Unexpected parser err = (%v), want nil = %v", err, test.parserNilErr)
		}
		// If there's an error parsing, the test cannot be completed.
		// The parsing error might be expected though.
		if err != nil {
			continue
		}

		// Call JSONDecoder to simulate decoding JSON -> Proto.
		err = JSONDecoder(r, &info.Arg)
		if err != nil {
			t.Errorf("Error while calling JSONDecoder, this should not happen. err: %v", err)
		}

		if got, want := info.Arg.(*v2pb.UpdateUserRequest).UserId, test.userId; got != want {
			t.Errorf("UserId = %v, want %v", got, want)
		}

		v2 := &FakeServer{}
		srv := New(v2)
		resp, err := info.H(srv, nil, nil)
		if err != nil {
			t.Errorf("Error while calling Fake_RequestHandler, this should not happen.")
		}
		if got, want := (*resp).(bool), true; got != want {
			t.Errorf("resp = %v, want %v.", got, want)
		}
	}
}

func TestListSEHV2_InitiateHandlerInfo(t *testing.T) {
	e, _ := strconv.ParseUint(primary_test_epoch, 10, 64)
	ps, _ := strconv.ParseUint(primary_test_page_size, 10, 32)
	var tests = []struct {
		path         string
		startEpoch   uint64
		pageSize     int32
		parserNilErr bool
	}{
		{"/v2/seh?start_epoch=" + primary_test_epoch + "&page_size=" + primary_test_page_size,
			e, int32(ps), true},
		{"/v2/seh?start_epoch=" + primary_test_epoch,
			e, 0, true},
		{"/v2/seh?page_size=" + primary_test_page_size,
			0, int32(ps), true},
		{"/v2/seh", 0, 0, true},
		// Invalid start_epoch format.
		{"/v2/seh?start_epoch=-2587",
			0, 0, false},
		{"/v2/seh?start_epoch=greatepoch",
			0, 0, false},
		// Invalid page_size format.
		{"/v2/seh?page_size=bigpagesize",
			0, 0, false},
	}

	for _, test := range tests {
		rInfo := handlers.RouteInfo{
			test.path,
			"GET",
			Fake_Initializer,
			Fake_RequestHandler,
		}
		// Body is empty when invoking list SEH API.
		jsonBody := "{}"

		info := ListSEHV2_InitializeHandlerInfo(rInfo)

		if _, ok := info.Arg.(*v2pb.ListSEHRequest); !ok {
			t.Errorf("info.Arg is not of type v2pb.ListSEHRequest")
		}

		r, _ := http.NewRequest(rInfo.Method, rInfo.Path, fakeJSONParserReader{bytes.NewBufferString(jsonBody)})
		err := info.Parser(r, &info.Arg)
		if got, want := (err == nil), test.parserNilErr; got != want {
			t.Errorf("Unexpected parser err = (%v), want nil = %v", err, test.parserNilErr)
		}
		// If there's an error parsing, the test cannot be completed.
		// The parsing error might be expected though.
		if err != nil {
			continue
		}

		// Call JSONDecoder to simulate decoding JSON -> Proto.
		err = JSONDecoder(r, &info.Arg)
		if err != nil {
			t.Errorf("Error while calling JSONDecoder, this should not happen. err: %v", err)
		}

		if got, want := info.Arg.(*v2pb.ListSEHRequest).StartEpoch, test.startEpoch; got != want {
			t.Errorf("StartEpoch = %v, want %v", got, want)
		}
		if got, want := info.Arg.(*v2pb.ListSEHRequest).PageSize, test.pageSize; got != want {
			t.Errorf("PageSize = %v, want %v", got, want)
		}

		v2 := &FakeServer{}
		srv := New(v2)
		resp, err := info.H(srv, nil, nil)
		if err != nil {
			t.Errorf("Error while calling Fake_RequestHandler, this should not happen.")
		}
		if got, want := (*resp).(bool), true; got != want {
			t.Errorf("resp = %v, want %v.", got, want)
		}
	}
}

func TestListUpdateV2_InitiateHandlerInfo(t *testing.T) {
	e, _ := strconv.ParseUint(primary_test_sequence, 10, 64)
	ps, _ := strconv.ParseUint(primary_test_page_size, 10, 32)
	var tests = []struct {
		path          string
		startSequence uint64
		pageSize      int32
		parserNilErr  bool
	}{
		{"/v2/seh?start_sequence=" + primary_test_sequence + "&page_size=" + primary_test_page_size,
			e, int32(ps), true},
		{"/v2/seh?start_sequence=" + primary_test_sequence,
			e, 0, true},
		{"/v2/seh?page_size=" + primary_test_page_size,
			0, int32(ps), true},
		{"/v2/seh", 0, 0, true},
		// Invalid start_sequence format.
		{"/v2/seh?start_sequence=-2587",
			0, 0, false},
		{"/v2/seh?start_sequence=greatsequence",
			0, 0, false},
		// Invalid page_size format.
		{"/v2/seh?page_size=bigpagesize",
			0, 0, false},
	}

	for _, test := range tests {
		rInfo := handlers.RouteInfo{
			test.path,
			"GET",
			Fake_Initializer,
			Fake_RequestHandler,
		}
		// Body is empty when invoking list update API.
		jsonBody := "{}"

		info := ListUpdateV2_InitializeHandlerInfo(rInfo)

		if _, ok := info.Arg.(*v2pb.ListUpdateRequest); !ok {
			t.Errorf("info.Arg is not of type v2pb.ListUpdateRequest")
		}

		r, _ := http.NewRequest(rInfo.Method, rInfo.Path, fakeJSONParserReader{bytes.NewBufferString(jsonBody)})
		err := info.Parser(r, &info.Arg)
		if got, want := (err == nil), test.parserNilErr; got != want {
			t.Errorf("Unexpected parser err = (%v), want nil = %v", err, test.parserNilErr)
		}
		// If there's an error parsing, the test cannot be completed.
		// The parsing error might be expected though.
		if err != nil {
			continue
		}

		// Call JSONDecoder to simulate decoding JSON -> Proto.
		err = JSONDecoder(r, &info.Arg)
		if err != nil {
			t.Errorf("Error while calling JSONDecoder, this should not happen. err: %v", err)
		}

		if got, want := info.Arg.(*v2pb.ListUpdateRequest).StartSequence, test.startSequence; got != want {
			t.Errorf("StartSequence = %v, want %v", got, want)
		}
		if got, want := info.Arg.(*v2pb.ListUpdateRequest).PageSize, test.pageSize; got != want {
			t.Errorf("PageSize = %v, want %v", got, want)
		}

		v2 := &FakeServer{}
		srv := New(v2)
		resp, err := info.H(srv, nil, nil)
		if err != nil {
			t.Errorf("Error while calling Fake_RequestHandler, this should not happen.")
		}
		if got, want := (*resp).(bool), true; got != want {
			t.Errorf("resp = %v, want %v.", got, want)
		}
	}
}

func TestListStepsV2_InitiateHandlerInfo(t *testing.T) {
	e, _ := strconv.ParseUint(primary_test_sequence, 10, 64)
	ps, _ := strconv.ParseUint(primary_test_page_size, 10, 32)
	var tests = []struct {
		path          string
		startSequence uint64
		pageSize      int32
		parserNilErr  bool
	}{
		{"/v2/seh?start_sequence=" + primary_test_sequence + "&page_size=" + primary_test_page_size,
			e, int32(ps), true},
		{"/v2/seh?start_sequence=" + primary_test_sequence,
			e, 0, true},
		{"/v2/seh?page_size=" + primary_test_page_size,
			0, int32(ps), true},
		{"/v2/seh", 0, 0, true},
		// Invalid start_sequence format.
		{"/v2/seh?start_sequence=-2587",
			0, 0, false},
		{"/v2/seh?start_sequence=greatsequence",
			0, 0, false},
		// Invalid page_size format.
		{"/v2/seh?page_size=bigpagesize",
			0, 0, false},
	}

	for _, test := range tests {
		rInfo := handlers.RouteInfo{
			test.path,
			"GET",
			Fake_Initializer,
			Fake_RequestHandler,
		}
		// Body is empty when invoking list steps API.
		jsonBody := "{}"

		info := ListStepsV2_InitializeHandlerInfo(rInfo)

		if _, ok := info.Arg.(*v2pb.ListStepsRequest); !ok {
			t.Errorf("info.Arg is not of type v2pb.ListStepsRequest")
		}

		r, _ := http.NewRequest(rInfo.Method, rInfo.Path, fakeJSONParserReader{bytes.NewBufferString(jsonBody)})
		err := info.Parser(r, &info.Arg)
		if got, want := (err == nil), test.parserNilErr; got != want {
			t.Errorf("Unexpected parser err = (%v), want nil = %v", err, test.parserNilErr)
		}
		// If there's an error parsing, the test cannot be completed.
		// The parsing error might be expected though.
		if err != nil {
			continue
		}

		// Call JSONDecoder to simulate decoding JSON -> Proto.
		err = JSONDecoder(r, &info.Arg)
		if err != nil {
			t.Errorf("Error while calling JSONDecoder, this should not happen. err: %v", err)
		}

		if got, want := info.Arg.(*v2pb.ListStepsRequest).StartSequence, test.startSequence; got != want {
			t.Errorf("StartSequence = %v, want %v", got, want)
		}
		if got, want := info.Arg.(*v2pb.ListStepsRequest).PageSize, test.pageSize; got != want {
			t.Errorf("PageSize = %v, want %v", got, want)
		}

		v2 := &FakeServer{}
		srv := New(v2)
		resp, err := info.H(srv, nil, nil)
		if err != nil {
			t.Errorf("Error while calling Fake_RequestHandler, this should not happen.")
		}
		if got, want := (*resp).(bool), true; got != want {
			t.Errorf("resp = %v, want %v.", got, want)
		}
	}
}

func JSONDecoder(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}

func TestParseURLComponent(t *testing.T) {
	mx := mux.NewRouter()
	mx.KeepContext = true
	mx.HandleFunc("/v1/users/{"+handlers.USER_ID_KEYWORD+"}", Fake_HTTPHandler)

	var tests = []struct {
		path    string
		keyword string
		out     string
		nilErr  bool
	}{
		{"/v1/users/" + primary_test_email, handlers.USER_ID_KEYWORD, primary_test_email, true},
		{"/v1/users/" + primary_test_email, "random_keyword", "", false},
	}
	for _, test := range tests {
		r, _ := http.NewRequest("GET", test.path, nil)
		mx.ServeHTTP(nil, r)
		gots, gote := parseURLVariable(r, test.keyword)
		wants := test.out
		wante := test.nilErr
		if gots != wants || wante != (gote == nil) {
			t.Errorf("Error while parsing User ID. Input = (%v, %v), got ('%v', %v), want ('%v', nil = %v)", test.path, test.keyword, gots, gote, wants, wante)
		}

	}
}

func Fake_HTTPHandler(w http.ResponseWriter, r *http.Request) {
}

func TestParseJson(t *testing.T) {
	var tests = []struct {
		inJSON    string
		outJSON   string
		outNilErr bool
	}{
		// Empty string
		{"", "", true},
		// Basic cases.
		{"\"creation_time\": \"" + valid_ts + "\"",
			"\"creation_time\": {\"seconds\": " +
				strconv.Itoa(ts_seconds) + ", \"nanos\": 0}", true},
		{"{\"creation_time\": \"" + valid_ts + "\"}",
			"{\"creation_time\": {\"seconds\": " +
				strconv.Itoa(ts_seconds) + ", \"nanos\": 0}}", true},
		// Nested case.
		{"{\"signed_key\":{\"key\": {\"creation_time\": \"" + valid_ts + "\"}}}",
			"{\"signed_key\":{\"key\": {\"creation_time\": {\"seconds\": " +
				strconv.Itoa(ts_seconds) + ", \"nanos\": 0}}}}", true},
		// Nothing to be changed.
		{"nothing to be changed here", "nothing to be changed here", true},
		// Multiple keywords.
		{"\"creation_time\": \"" + valid_ts + "\", \"creation_time\": \"" +
			valid_ts + "\"",
			"\"creation_time\": {\"seconds\": " + strconv.Itoa(ts_seconds) +
				", \"nanos\": 0}, \"creation_time\": {\"seconds\": " +
				strconv.Itoa(ts_seconds) + ", \"nanos\": 0}", true},
		// Invalid timestamp.
		{"\"creation_time\": \"invalid\"", "\"creation_time\": \"invalid\"", false},
		// Empty timestamp.
		{"\"creation_time\": \"\"", "\"creation_time\": \"\"", false},
		{"\"creation_time\": \"\", \"creation_time\": \"\"",
			"\"creation_time\": \"\", \"creation_time\": \"\"", false},
		// Malformed JSON, missing " at the beginning of invalid
		// timestamp.
		{"\"creation_time\": invalid\"", "\"creation_time\": invalid\"", true},
		// Malformed JSON, missing " at the end of invalid timestamp.
		{"\"creation_time\": \"invalid", "\"creation_time\": \"invalid", true},
		// Malformed JSON, missing " at the beginning and end of
		// invalid timestamp.
		{"\"creation_time\": invalid", "\"creation_time\": invalid", true},
		// Malformed JSON, missing " at the end of valid timestamp.
		{"\"creation_time\": \"" + valid_ts, "\"creation_time\": \"" + valid_ts, true},
		// keyword is not surrounded by "", in four cases: invalid
		// timestamp, basic, nested and multiple keywords.
		{"creation_time: \"invalid\"", "creation_time: \"invalid\"", false},
		{"{creation_time: \"" + valid_ts + "\"}",
			"{creation_time: {\"seconds\": " +
				strconv.Itoa(ts_seconds) + ", \"nanos\": 0}}", true},
		{"{\"signed_key\":{\"key\": {creation_time: \"" + valid_ts + "\"}}}",
			"{\"signed_key\":{\"key\": {creation_time: {\"seconds\": " +
				strconv.Itoa(ts_seconds) + ", \"nanos\": 0}}}}", true},
		// Only first keyword is not surrounded by "".
		{"creation_time: \"" + valid_ts + "\", \"creation_time\": \"" +
			valid_ts + "\"",
			"creation_time: {\"seconds\": " + strconv.Itoa(ts_seconds) +
				", \"nanos\": 0}, \"creation_time\": {\"seconds\": " +
				strconv.Itoa(ts_seconds) + ", \"nanos\": 0}", true},
		// Timestamp is not surrounded by "" and there's other keys and
		// values after.
		{"{\"signed_key\":{\"key\": {\"creation_time\": " + valid_ts +
			", app_id: \"" + primary_test_app_id + "\"}}}",
			"{\"signed_key\":{\"key\": {\"creation_time\": " + valid_ts +
				", app_id: \"" + primary_test_app_id + "\"}}}", true},
	}

	for _, test := range tests {
		r, _ := http.NewRequest("", "", fakeJSONParserReader{bytes.NewBufferString(test.inJSON)})
		err := parseJSON(r, "creation_time")
		if test.outNilErr != (err == nil) {
			t.Errorf("Unexpected JSON parser err = (%v), want nil = %v", err, test.outNilErr)
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		if got, want := buf.String(), test.outJSON; got != want {
			t.Errorf("Out JSON = (%v), want (%v)", got, want)
		}
	}
}
