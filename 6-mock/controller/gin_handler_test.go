package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	usecasemock "github.com/vivaldy22/go-unit-test/6-mock/mocks/usecase"
	"github.com/vivaldy22/go-unit-test/6-mock/usecase"
)

func TestUserGinHandler_GetByID(t *testing.T) {
	type mockRes struct {
		res usecase.User
		err error
	}

	tests := []struct {
		name            string
		paramID         int64
		mockRes         mockRes
		expectedJSONRes string
		expectedErrCode int
		expectErr       bool
	}{
		{
			name:    "UserHandler_GetByID success",
			paramID: 1,
			mockRes: mockRes{
				res: usecase.User{
					ID:   1,
					Name: "test 1",
					Age:  21,
				},
				err: nil,
			},
			expectedJSONRes: "{\"user\":{\"id\":1,\"name\":\"test 1\",\"age\":21}}\n",
			expectErr:       false,
		},
		{
			name:    "UserHandler_GetByID return 500",
			paramID: 1,
			mockRes: mockRes{
				err: errors.New("general-error"),
			},
			expectedErrCode: http.StatusInternalServerError,
			expectErr:       true,
		},
		{
			name: "UserHandler_GetByID return 400",
			mockRes: mockRes{
				err: errors.New("general-error"),
			},
			expectedErrCode: http.StatusBadRequest,
			expectErr:       true,
		},
	}

	for _, tt := range tests {
		gin.SetMode(gin.TestMode)
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := &usecasemock.UserService{}
			mockUsecase.On("GetByID", mock.Anything).Return(tt.mockRes.res, tt.mockRes.err)
			rec := httptest.NewRecorder()
			r := gin.Default()
			if tt.expectedErrCode != http.StatusBadRequest {
				r.Use(func(c *gin.Context) {
					c.Set("id", tt.paramID)
				})
			}

			h := NewUserGinHandler(&UserGinHandlerConfig{
				R:           r,
				UserService: mockUsecase,
			})
			r.GET("/users", h.GetByID)
			req, err := http.NewRequest(http.MethodGet, "/users", nil)
			assert.NoError(t, err)
			r.ServeHTTP(rec, req)

			if tt.expectErr {
				assert.Equal(t, tt.expectedErrCode, rec.Code)
			} else {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.Equal(t, tt.expectedJSONRes, rec.Body.String())
			}
		})
	}
}
