package transport

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/pkg/util"
	"github.com/stretchr/testify/require"
)

func createPost(t *testing.T, postInput *model.PostInfoInput) string {
	byteArr, err := json.Marshal(postInput)
	require.NoError(t, err)

	path := "/v1/api/post/"
	request := defaultRequest(http.MethodPost, path, byteArr)
	request.Header.Set("Authorization", testAuthorToken)

	response := httptest.NewRecorder()
	mockServer.ServeHTTP(response, request)

	res := response.Result()

	require.Equal(t, res.StatusCode, http.StatusCreated)

	body, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	var retResponse Response
	err = json.Unmarshal(body, &retResponse)
	require.NoError(t, err)

	result := strings.Split(retResponse.Message, " ")
	require.Equal(t, "post_id:", result[0])
	require.NotZero(t, result[1])

	return result[1]
}

func TestPost_UploadPost(t *testing.T) {

	testCases := []struct {
		name          string
		postInput     *model.PostInfoInput
		createRequest func(t *testing.T, method, path string, postInput *model.PostInfoInput) *http.Request
		checkResponse func(t *testing.T, res *http.Response)
	}{
		{
			name: "Successful create",
			postInput: &model.PostInfoInput{
				Header: util.RandomString(10),
				Text:   util.RandomString(10),
				Author: testAuthorId,
				TTL:    time.Now().Add(time.Minute).Format(time.RFC3339),
			},
			// Create test Request
			createRequest: func(t *testing.T, method, path string, postInput *model.PostInfoInput) *http.Request {
				byteArr, err := json.Marshal(postInput)
				require.NoError(t, err)

				request := defaultRequest(method, path, byteArr)
				request.Header.Set("Authorization", testAuthorToken)
				return request
			},
			// Check Response
			checkResponse: func(t *testing.T, res *http.Response) {
				require.Equal(t, res.StatusCode, http.StatusCreated)

				body, err := io.ReadAll(res.Body)
				require.NoError(t, err)

				var retResponse Response
				err = json.Unmarshal(body, &retResponse)
				require.NoError(t, err)

				result := strings.Split(retResponse.Message, " ")
				require.Equal(t, "post_id:", result[0])
				require.NotZero(t, result[1])
			},
		},
		{
			name: "Unauthorized error",
			postInput: &model.PostInfoInput{
				Header: util.RandomString(10),
				Text:   util.RandomString(10),
				Author: testAuthorId,
				TTL:    time.Now().Add(time.Minute).Format(time.RFC3339),
			},
			createRequest: func(t *testing.T, method, path string, postInput *model.PostInfoInput) *http.Request {
				byteArr, err := json.Marshal(postInput)
				require.NoError(t, err)

				request := defaultRequest(method, path, byteArr)
				return request
			},
			checkResponse: func(t *testing.T, res *http.Response) {
				require.Equal(t, res.StatusCode, http.StatusUnauthorized)

				body, err := io.ReadAll(res.Body)
				require.NoError(t, err)

				var retResponse Response
				err = json.Unmarshal(body, &retResponse)
				require.NoError(t, err)
				require.NotEmpty(t, retResponse)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		path := "/v1/api/post/"

		t.Run(tc.name, func(t *testing.T) {
			request := tc.createRequest(t, http.MethodPost, path, tc.postInput)

			response := httptest.NewRecorder()
			mockServer.ServeHTTP(response, request)

			res := response.Result()

			tc.checkResponse(t, res)
		})

	}
}

func TestPost_GetPostByID(t *testing.T) {
	testCases := []struct {
		name          string
		postInput     *model.PostInfoInput
		createRequest func(t *testing.T, method, path string) *http.Request
		checkResponse func(t *testing.T, res *http.Response, input *model.PostInfoInput)
	}{
		{
			name: "OK",
			postInput: &model.PostInfoInput{
				Header: util.RandomString(10),
				Text:   util.RandomString(10),
				Author: testAuthorId,
				TTL:    time.Now().Add(time.Minute).Format(time.RFC3339),
			},
			// Create test Request
			createRequest: func(t *testing.T, method, path string) *http.Request {
				request := defaultRequest(method, path, nil)
				request.Header.Set("Authorization", testAuthorToken)
				return request
			},
			// Check Response
			checkResponse: func(t *testing.T, res *http.Response, input *model.PostInfoInput) {
				require.Equal(t, res.StatusCode, http.StatusOK)

				body, err := io.ReadAll(res.Body)
				require.NoError(t, err)

				var findPost model.PostInfo
				err = json.Unmarshal(body, &findPost)
				require.NoError(t, err)

				// TODO разобраться как проверить type time.Time
				//require.Equal(t, expectedTTL, findPost.TTL)

				require.Equal(t, input.Header, findPost.Header)
				require.Equal(t, input.Author, findPost.Author)
				require.Equal(t, input.Text, findPost.Text)
			},
		},
		{
			name: "Not found",
			postInput: &model.PostInfoInput{
				Header: util.RandomString(10),
				Text:   util.RandomString(10),
				Author: testAuthorId,
				TTL:    time.Now().Add(time.Minute).Format(time.RFC3339),
			},
			createRequest: func(t *testing.T, method, path string) *http.Request {
				path += "2361"
				request := defaultRequest(method, path, nil)
				request.Header.Set("Authorization", testAuthorToken)
				return request
			},
			checkResponse: func(t *testing.T, res *http.Response, input *model.PostInfoInput) {
				require.Equal(t, http.StatusNotFound, res.StatusCode)

				body, err := io.ReadAll(res.Body)
				require.NoError(t, err)

				var retResponse Response
				err = json.Unmarshal(body, &retResponse)
				require.NoError(t, err)
				require.NotEmpty(t, retResponse)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		postId := createPost(t, tc.postInput)
		path := "/v1/api/post/" + postId

		t.Run(tc.name, func(t *testing.T) {
			request := tc.createRequest(t, http.MethodGet, path)

			response := httptest.NewRecorder()
			mockServer.ServeHTTP(response, request)

			res := response.Result()

			tc.checkResponse(t, res, tc.postInput)
		})

	}
}

func TestPost_UpdatePostByID(t *testing.T) {
	testCases := []struct {
		name          string
		postInput     *model.PostInfoInput
		postUpdate    *model.PostInfoUpdate
		createRequest func(t *testing.T, method, path string, updatePost *model.PostInfoUpdate) *http.Request
		checkResponse func(t *testing.T, res *http.Response, update *model.PostInfoUpdate)
	}{
		{
			name: "OK",
			postInput: &model.PostInfoInput{
				Header: util.RandomString(10),
				Text:   util.RandomString(10),
				Author: testAuthorId,
				TTL:    time.Now().Add(time.Minute).Format(time.RFC3339),
			},
			postUpdate: &model.PostInfoUpdate{
				Header: util.RandomString(10),
				Text:   util.RandomString(10),
				TTL:    time.Now().Add(time.Minute * 2).Format(time.RFC3339),
			},
			// Create test Request
			createRequest: func(t *testing.T, method, path string, updatePost *model.PostInfoUpdate) *http.Request {
				byteArr, err := json.Marshal(updatePost)
				require.NoError(t, err)

				request := defaultRequest(method, path, byteArr)
				request.Header.Set("Authorization", testAuthorToken)
				return request
			},
			// Check Response
			checkResponse: func(t *testing.T, res *http.Response, update *model.PostInfoUpdate) {
				require.Equal(t, res.StatusCode, http.StatusOK)

				body, err := io.ReadAll(res.Body)
				require.NoError(t, err)

				var findPost model.PostInfo
				err = json.Unmarshal(body, &findPost)
				require.NoError(t, err)

				// TODO разобраться как проверить type time.Time
				//require.Equal(t, update.TTL, findPost.TTL)

				require.Equal(t, update.Header, findPost.Header)
				require.Equal(t, update.Text, findPost.Text)
			},
		},
		{
			name: "TTL less than was",
			postInput: &model.PostInfoInput{
				Header: util.RandomString(10),
				Text:   util.RandomString(10),
				Author: testAuthorId,
				TTL:    time.Now().Add(time.Minute).Format(time.RFC3339),
			},
			postUpdate: &model.PostInfoUpdate{
				Header: util.RandomString(10),
				Text:   util.RandomString(10),
				TTL:    time.Now().Add(-(time.Minute * 2)).Format(time.RFC3339), // Time expiry
			},
			// Create test Request
			createRequest: func(t *testing.T, method, path string, updatePost *model.PostInfoUpdate) *http.Request {
				byteArr, err := json.Marshal(updatePost)
				require.NoError(t, err)

				request := defaultRequest(method, path, byteArr)
				request.Header.Set("Authorization", testAuthorToken)
				return request
			},
			// Check Response
			checkResponse: func(t *testing.T, res *http.Response, update *model.PostInfoUpdate) {
				require.Equal(t, res.StatusCode, http.StatusOK)

				body, err := io.ReadAll(res.Body)
				require.NoError(t, err)

				var findPost model.PostInfo
				err = json.Unmarshal(body, &findPost)
				require.NoError(t, err)
				require.Empty(t, findPost)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		postId := createPost(t, tc.postInput)

		path := "/v1/api/post/" + postId

		t.Run(tc.name, func(t *testing.T) {
			request := tc.createRequest(t, http.MethodPut, path, tc.postUpdate)

			response := httptest.NewRecorder()
			mockServer.ServeHTTP(response, request)

			res := response.Result()

			tc.checkResponse(t, res, tc.postUpdate)
		})

	}
}

func defaultRequest(method, path string, byteArr []byte) *http.Request {
	request := httptest.NewRequest(method, path, bytes.NewBuffer(byteArr))
	request.Header.Set("Content-Type", "application/json")
	return request
}
