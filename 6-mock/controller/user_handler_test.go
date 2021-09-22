package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	usecasemock "github.com/vivaldy22/go-unit-test/6-mock/mocks/usecase"
	"github.com/vivaldy22/go-unit-test/6-mock/usecase"
)

func TestNewUserHandler(t *testing.T) {

}

func TestUserHandler_GetByID(t *testing.T) {
	type mockRes struct {
		res usecase.User
		err error
	}

	tests := []struct {
		name            string
		paramID         string
		mockRes         mockRes
		expectedJSONRes string
		expectedErrCode int
		expectErr       bool
	}{
		{
			name:    "UserHandler_GetByID success",
			paramID: "1",
			mockRes: mockRes{
				res: usecase.User{
					ID:   1,
					Name: "test 1",
					Age:  21,
				},
				err: nil,
			},
			expectedJSONRes: "{\"id\":1,\"name\":\"test 1\",\"age\":21}\n",
			expectErr:       false,
		},
		{
			name:    "UserHandler_GetByID return 500",
			paramID: "1",
			mockRes: mockRes{
				err: errors.New("general-error"),
			},
			expectedErrCode: http.StatusInternalServerError,
			expectErr:       true,
		},
		{
			name:    "UserHandler_GetByID return 400",
			paramID: "a",
			mockRes: mockRes{
				err: errors.New("general-error"),
			},
			expectedErrCode: http.StatusBadRequest,
			expectErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.paramID)

			mockUsecase := &usecasemock.UserService{}
			mockUsecase.On("GetByID", mock.Anything).Return(tt.mockRes.res, tt.mockRes.err)

			handler := NewUserHandler(mockUsecase)

			err := handler.GetByID(c)
			if tt.expectErr {
				assert.NoError(err)
				assert.Equal(tt.expectedErrCode, rec.Code)
			} else {
				assert.NoError(err)
				assert.Equal(http.StatusOK, rec.Code)
				assert.Equal(tt.expectedJSONRes, rec.Body.String())
			}
		})
	}
}
