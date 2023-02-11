package controllers_test

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/kinbiko/jsonassert"
	"github.com/rmargar/website/pkg/domain"
	"github.com/rmargar/website/pkg/web/controllers"
	"github.com/stretchr/testify/mock"
)

type PostServiceMock struct {
	mock.Mock
}

func (m *PostServiceMock) Create(title string, content string, tags []string, summary string, urlPath string) (*domain.Post, error) {
	args := m.Called(title, content, tags)
	return args.Get(0).(*domain.Post), args.Error(1)
}

func (m *PostServiceMock) GetOneByTitle(title string) (*domain.Post, error) {
	args := m.Called(title)
	return args.Get(0).(*domain.Post), args.Error(1)
}

func (m *PostServiceMock) GetAll() ([]domain.Post, error) {
	args := m.Called()
	return args.Get(0).([]domain.Post), args.Error(1)
}

func (m *PostServiceMock) GetByUrlPath(urlPath string) (domain.Post, error) {
	args := m.Called()
	return args.Get(0).(domain.Post), args.Error(1)
}

func TestAddPost_Success(t *testing.T) {
	mockService := &PostServiceMock{}
	nowTime := time.Now()
	mockService.
		On("Create", "Test", "Test", []string(nil)).
		Return(
			&domain.Post{
				ID:       1,
				Added:    nowTime,
				Modified: nowTime,
				Author:   "rmargar",
				Tags:     []string{},
				Title:    "Test",
				Content:  "Test",
				URLPath:  "test",
			},
			nil,
		)

	controller := controllers.Posts{PostService: mockService}
	reader := strings.NewReader(`{
		"title": "Test",
		"content": "Test",
		"urlPath": "test"
	  }`,
	)
	request := httptest.NewRequest("POST", "http://localhost/api/posts", reader)
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()

	controller.AddPost(writer, request)

	assert := jsonassert.New(t)
	const layout string = "2006-01-02T15:04:05.999999999-07:00"
	expectedResponse := fmt.Sprintf(`{"msg":"Created","data":{"id":1,"title":"Test","content":"Test","tags":[],"author":"rmargar","added":"%s","modified":"%s","urlPath":"test", "summary":"", "imgUrl":""}}`, nowTime.Format(time.RFC3339Nano), nowTime.Format(time.RFC3339Nano))
	assert.Assertf(writer.Body.String(), expectedResponse)
}

func TestAddPost_ThrowsValidationError(t *testing.T) {
	mockService := &PostServiceMock{}

	controller := controllers.Posts{PostService: mockService}
	reader := strings.NewReader(`{
		"title": "Test"
	  }`,
	)
	request := httptest.NewRequest("POST", "http://localhost/api/posts", reader)
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()

	controller.AddPost(writer, request)
	assert := jsonassert.New(t)
	expectedResponse := `{"errors":[{"title":"POST Payload is invalid","detail":["content is required","urlPath is required"],"code":"invalid-payload","source":""}]}`
	assert.Assertf(writer.Body.String(), expectedResponse)
}
