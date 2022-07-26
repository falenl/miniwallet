// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/interfaces.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	entity "github.com/falenl/miniwallet/entity"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAccountRepository) Authenticate(ctx context.Context, token string) (entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, token)
	ret0, _ := ret[0].(entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAccountRepositoryMockRecorder) Authenticate(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAccountRepository)(nil).Authenticate), ctx, token)
}

// Create mocks base method.
func (m *MockAccountRepository) Create(ctx context.Context, customerID uuid.UUID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, customerID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountRepositoryMockRecorder) Create(ctx, customerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountRepository)(nil).Create), ctx, customerID)
}

// MockCustomerService is a mock of CustomerService interface.
type MockCustomerService struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerServiceMockRecorder
}

// MockCustomerServiceMockRecorder is the mock recorder for MockCustomerService.
type MockCustomerServiceMockRecorder struct {
	mock *MockCustomerService
}

// NewMockCustomerService creates a new mock instance.
func NewMockCustomerService(ctrl *gomock.Controller) *MockCustomerService {
	mock := &MockCustomerService{ctrl: ctrl}
	mock.recorder = &MockCustomerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerService) EXPECT() *MockCustomerServiceMockRecorder {
	return m.recorder
}

// Verify mocks base method.
func (m *MockCustomerService) Verify(ctx context.Context, customerID uuid.UUID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, customerID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockCustomerServiceMockRecorder) Verify(ctx, customerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockCustomerService)(nil).Verify), ctx, customerID)
}

// MockWalletRepository is a mock of WalletRepository interface.
type MockWalletRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWalletRepositoryMockRecorder
}

// MockWalletRepositoryMockRecorder is the mock recorder for MockWalletRepository.
type MockWalletRepositoryMockRecorder struct {
	mock *MockWalletRepository
}

// NewMockWalletRepository creates a new mock instance.
func NewMockWalletRepository(ctrl *gomock.Controller) *MockWalletRepository {
	mock := &MockWalletRepository{ctrl: ctrl}
	mock.recorder = &MockWalletRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletRepository) EXPECT() *MockWalletRepositoryMockRecorder {
	return m.recorder
}

// Enable mocks base method.
func (m *MockWalletRepository) Enable(ctx context.Context, wallet *entity.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", ctx, wallet)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockWalletRepositoryMockRecorder) Enable(ctx, wallet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockWalletRepository)(nil).Enable), ctx, wallet)
}

// GetByAccountID mocks base method.
func (m *MockWalletRepository) GetByAccountID(ctx context.Context, accountID string) (entity.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAccountID", ctx, accountID)
	ret0, _ := ret[0].(entity.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAccountID indicates an expected call of GetByAccountID.
func (mr *MockWalletRepositoryMockRecorder) GetByAccountID(ctx, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAccountID", reflect.TypeOf((*MockWalletRepository)(nil).GetByAccountID), ctx, accountID)
}

// Update mocks base method.
func (m *MockWalletRepository) Update(ctx context.Context, wallet *entity.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, wallet)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWalletRepositoryMockRecorder) Update(ctx, wallet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWalletRepository)(nil).Update), ctx, wallet)
}

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// Deposit mocks base method.
func (m *MockTransactionRepository) Deposit(ctx context.Context, trx *entity.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deposit", ctx, trx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deposit indicates an expected call of Deposit.
func (mr *MockTransactionRepositoryMockRecorder) Deposit(ctx, trx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deposit", reflect.TypeOf((*MockTransactionRepository)(nil).Deposit), ctx, trx)
}

// Withdraw mocks base method.
func (m *MockTransactionRepository) Withdraw(ctx context.Context, trx *entity.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", ctx, trx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockTransactionRepositoryMockRecorder) Withdraw(ctx, trx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockTransactionRepository)(nil).Withdraw), ctx, trx)
}
