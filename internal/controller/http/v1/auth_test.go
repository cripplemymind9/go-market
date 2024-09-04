package v1

import (
	"bytes"
	"errors"
	"context"
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/cripplemymind9/go-market/internal/mocks/servicemocks"
	"github.com/cripplemymind9/go-market/internal/service"
	"github.com/cripplemymind9/go-market/internal/service/serviceerrs"
	"github.com/cripplemymind9/go-market/internal/service/types"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/golang/mock/gomock"
	"github.com/gin-gonic/gin"
)

func TestAuthRoutes_SignUp(t *testing.T) {
	type args struct {
		ctx context.Context
		input types.AuthRegisterUserInput
	}
	
	type MockBehaviour func(m *servicemocks.MockAuth, args args)

	testCases := []struct {
		name                 string
		args 				 args
		inputBody            string
		mockBehaviour         MockBehaviour
		wantStatusCode       int
		wantRequestBody      string
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				input: types.AuthRegisterUserInput{
					Username: "test",
					Password: "Qwerty!1",
					Email: 	  "test@example.com",
					},
				},
			inputBody: `{"username":"test","password":"Qwerty!1","email":"test@example.com"}`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {
				m.EXPECT().RegisterUser(args.ctx, args.input).Return(1, nil)
			},
			wantStatusCode:  201,
			wantRequestBody: `{"id":1}` + "\n",
		},
		{
			name: "Invalid password: not provided",
			args: args{},
			inputBody: `{"username":"test","email":"test@example.com"}`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {},
			wantStatusCode:  400,
			wantRequestBody: `{"error":"Key: 'signUpInput.Password' Error:Field validation for 'Password' failed on the 'required' tag"}` + "\n",
		},
		{
			name: "Invalid username: not provided",
			args: args{},
			inputBody: `{"password":"Qwerty!1","email":"test@example.com"}`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {},
			wantStatusCode:  400,
			wantRequestBody: `{"error":"Key: 'signUpInput.Username' Error:Field validation for 'Username' failed on the 'required' tag"}` + "\n",
		},
		{
			name: "Invalid email: not provided",
			args: args{},
			inputBody: `{"username":"test","password":"Qwerty!1"}`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {},
			wantStatusCode:  400,
			wantRequestBody: `{"error":"Key: 'signUpInput.Email' Error:Field validation for 'Email' failed on the 'required' tag"}` + "\n",
		},
		{
			name: "Invalid request body",
			args: args{},
			inputBody: `{"username" test","password":"Qwerty!1"`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {},
			wantStatusCode:  400,
			wantRequestBody: `{"error":"invalid request body"}` + "\n",
		},
		{
			name: "Auth service error",
			args: args{
				ctx: context.Background(),
				input: types.AuthRegisterUserInput{
					Username: "test",
					Password: "Qwerty!1",
					Email:    "test@example.com",
				},
			},
			inputBody: `{"username":"test","password":"Qwerty!1","email":"test@example.com"}`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {
				m.EXPECT().RegisterUser(args.ctx, args.input).Return(0, serviceerrs.ErrUserAlreadyExists)
			},
			wantStatusCode:  400,
			wantRequestBody: `{"error":"user already exists"}` + "\n",
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {
			// Init deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Init service mock
			auth := servicemocks.NewMockAuth(ctrl)
			tc.mockBehaviour(auth, tc.args)
			services := &service.Services{Auth: auth}

			// Create router
			router := gin.Default()
			authRoutes := &authRoutes{
				authService: services.Auth,
				validator: 	validator.New(),
			}
			router.POST("/auth/sign-up", authRoutes.signUp)

			// Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewBufferString(tc.inputBody))
			req.Header.Set("Content-Type", "application/json")

			// Execute request
			router.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tc.wantStatusCode, w.Code)
			assert.JSONEq(t, tc.wantRequestBody, w.Body.String())
		})
	}
}

func TestAuthRoutes_SignIn(t *testing.T) {
	type args struct {
		ctx 	context.Context
		input 	types.AuthGenerateTokenInput
	}

	type mockBehaviour func(m *servicemocks.MockAuth, args args)

	testCases := []struct {
		name 				string
		args				args
		inputBody 			string
		mockBehaviour 		mockBehaviour
		wantStatusCode 		int
		wantRequestBody 	string
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				input: types.AuthGenerateTokenInput{
					Username: "test",
					Password: "Qwerty1!",
					},
				},
			inputBody: `{"username":"test","password":"Qwerty1!"}`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {
				m.EXPECT().GenerateToken(args.ctx, args.input).Return("token", nil)
			},
			wantStatusCode: 201,
			wantRequestBody: `{"token":"token"}` + "\n",
		},
		{
			name: "Invalid username: not provided",
			args: args{},
			inputBody: `{"password":"Qwerty1!"}`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {},
			wantStatusCode: 400,
			wantRequestBody: `{"error":"Key: 'signInInput.Username' Error:Field validation for 'Username' failed on the 'required' tag"}` + "\n",
		},
		{
			name: "Invalid password: not provided",
			args: args{},
			inputBody: `{"username":"test"}`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {},
			wantStatusCode: 400,
			wantRequestBody: `{"error":"Key: 'signInInput.Password' Error:Field validation for 'Password' failed on the 'required' tag"}` + "\n",
		},
		{
			name: "Wrong username or password",
			args: args{
				ctx: context.Background(),
				input: types.AuthGenerateTokenInput{
					Username: "test",
					Password: "Qwerty1!",
					},
				},
			inputBody: `{"username":"test","password":"Qwerty1!"}`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {
				m.EXPECT().GenerateToken(args.ctx, args.input).Return("", serviceerrs.ErrUserNotFound)
			},
			wantStatusCode: 400,
			wantRequestBody: `{"error":"Invalid credentials"}` + "\n",
		},
		{
			name: "Invalid request body",
			args: args{},
			inputBody: `({Qwerty1!)`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {},
			wantStatusCode: 400,
			wantRequestBody: `{"error":"invalid request body"}` + "\n",
		},
		{
			name: "Internal server error",
			args: args{
				ctx: context.Background(),
				input: types.AuthGenerateTokenInput{
					Username: "test",
					Password: "Qwerty1!",
					},
				},
			inputBody: `{"username":"test","password":"Qwerty1!"}`,
			mockBehaviour: func(m *servicemocks.MockAuth, args args) {
				m.EXPECT().GenerateToken(args.ctx, args.input).Return("", errors.New("some error"))
			},
			wantStatusCode: 500,
			wantRequestBody: `{"error":"Internal server error"}` + "\n",
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {
			// Init deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Init service mock
			auth := servicemocks.NewMockAuth(ctrl)
			tc.mockBehaviour(auth, tc.args)
			services := &service.Services{Auth: auth}

			// Create router
			router := gin.Default()
			authRoutes := &authRoutes{
				authService: services.Auth,
				validator: 	validator.New(),
			}
			router.POST("/auth/sign-in", authRoutes.signIn)

			// Create request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/auth/sign-in", bytes.NewBufferString(tc.inputBody))
			req.Header.Set("Content-Type", "application/json")

			// Execute request
			router.ServeHTTP(w, req)

			// Check response
			assert.Equal(t, tc.wantStatusCode, w.Code)
			assert.JSONEq(t, tc.wantRequestBody, w.Body.String())
		})
	}
}