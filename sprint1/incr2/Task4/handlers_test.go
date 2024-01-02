package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserViewHandler(t *testing.T) {
	type args struct {
		users map[string]User
	}
	type want struct {
		code        int
		user        User
		contentType string
	}
	tests := []struct {
		name   string
		args   args
		reqUrl string
		want   want
	}{
		{
			name: "no body",
			args: args{
				users: map[string]User{
					"u1": User{
						ID:        "u1",
						FirstName: "ss",
						LastName:  "ssp"}}},
			reqUrl: "/users",
			want: want{
				code:        http.StatusBadRequest,
				user:        User{},
				contentType: "application/json"},
		},
		{
			name: "no user",
			args: args{
				users: map[string]User{
					"u1": User{
						ID:        "u1",
						FirstName: "ss",
						LastName:  "ssp"}}},
			reqUrl: "/users?user_id=u2",
			want: want{
				code:        http.StatusNotFound,
				user:        User{},
				contentType: "application/json"},
		},

		{
			name: "ok",
			args: args{
				users: map[string]User{
					"u1": User{
						ID:        "u1",
						FirstName: "ss",
						LastName:  "ssp"}}},
			reqUrl: "/users?user_id=u1",
			want: want{
				code: http.StatusOK,
				user: User{ID: "u1",
					FirstName: "ss",
					LastName:  "ssp"},
				contentType: "application/json"},
		},
		// reqUrl: "/users?user_id="
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.reqUrl, nil)
			// создаём новый Recorder
			w := httptest.NewRecorder()

			UserViewHandler(tt.args.users)(w, request)
			res := w.Result()

			assert.Equal(t, tt.want.code, res.StatusCode, "Code is %d, wanted %d", res.StatusCode, tt.want.code)
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"), "contentType is %s, wanted %s", res.Header.Get("Content-Type"), tt.want.contentType)

			// defer res.Body.Close()
			resBody, err := ioutil.ReadAll(res.Body)

			require.NoError(t, err)
			err = res.Body.Close()
			require.NoError(t, err)

			var user User
			err = json.Unmarshal(resBody, &user)
			require.NoError(t, err)

			assert.Equal(t, tt.want.user, user)

		})
	}
}
