// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\Lenovo\Desktop\go-market\internal\service\service.go

// Package servicemocks is a generated GoMock package.
package servicemocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/cripplemymind9/go-market/internal/entity"
	types "github.com/cripplemymind9/go-market/internal/service/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockAuth) GenerateToken(ctx context.Context, input types.AuthGenerateTokenInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", ctx, input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthMockRecorder) GenerateToken(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuth)(nil).GenerateToken), ctx, input)
}

// ParseToken mocks base method.
func (m *MockAuth) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuth)(nil).ParseToken), token)
}

// RegisterUser mocks base method.
func (m *MockAuth) RegisterUser(ctx context.Context, input types.AuthRegisterUserInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockAuthMockRecorder) RegisterUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockAuth)(nil).RegisterUser), ctx, input)
}

// MockProduct is a mock of Product interface.
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct.
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance.
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// AddProduct mocks base method.
func (m *MockProduct) AddProduct(ctx context.Context, input types.ProductAddProductInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", ctx, input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockProductMockRecorder) AddProduct(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockProduct)(nil).AddProduct), ctx, input)
}

// DeleteProduct mocks base method.
func (m *MockProduct) DeleteProduct(ctx context.Context, productId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", ctx, productId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductMockRecorder) DeleteProduct(ctx, productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProduct)(nil).DeleteProduct), ctx, productId)
}

// GetAllProducts mocks base method.
func (m *MockProduct) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts", ctx)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockProductMockRecorder) GetAllProducts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockProduct)(nil).GetAllProducts), ctx)
}

// GetProductById mocks base method.
func (m *MockProduct) GetProductById(ctx context.Context, productId int) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", ctx, productId)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockProductMockRecorder) GetProductById(ctx, productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockProduct)(nil).GetProductById), ctx, productId)
}

// UpdateProduct mocks base method.
func (m *MockProduct) UpdateProduct(ctx context.Context, input types.ProductUpdateProductInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductMockRecorder) UpdateProduct(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProduct)(nil).UpdateProduct), ctx, input)
}

// MockPurchase is a mock of Purchase interface.
type MockPurchase struct {
	ctrl     *gomock.Controller
	recorder *MockPurchaseMockRecorder
}

// MockPurchaseMockRecorder is the mock recorder for MockPurchase.
type MockPurchaseMockRecorder struct {
	mock *MockPurchase
}

// NewMockPurchase creates a new mock instance.
func NewMockPurchase(ctrl *gomock.Controller) *MockPurchase {
	mock := &MockPurchase{ctrl: ctrl}
	mock.recorder = &MockPurchaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPurchase) EXPECT() *MockPurchaseMockRecorder {
	return m.recorder
}

// GetProductPurchases mocks base method.
func (m *MockPurchase) GetProductPurchases(ctx context.Context, productId int) ([]entity.Purchase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductPurchases", ctx, productId)
	ret0, _ := ret[0].([]entity.Purchase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductPurchases indicates an expected call of GetProductPurchases.
func (mr *MockPurchaseMockRecorder) GetProductPurchases(ctx, productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductPurchases", reflect.TypeOf((*MockPurchase)(nil).GetProductPurchases), ctx, productId)
}

// GetUserPurchases mocks base method.
func (m *MockPurchase) GetUserPurchases(ctx context.Context, userId int) ([]entity.Purchase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPurchases", ctx, userId)
	ret0, _ := ret[0].([]entity.Purchase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPurchases indicates an expected call of GetUserPurchases.
func (mr *MockPurchaseMockRecorder) GetUserPurchases(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPurchases", reflect.TypeOf((*MockPurchase)(nil).GetUserPurchases), ctx, userId)
}

// MakePurchase mocks base method.
func (m *MockPurchase) MakePurchase(ctx context.Context, input types.PurchaseMakePurchaseInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakePurchase", ctx, input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakePurchase indicates an expected call of MakePurchase.
func (mr *MockPurchaseMockRecorder) MakePurchase(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakePurchase", reflect.TypeOf((*MockPurchase)(nil).MakePurchase), ctx, input)
}
