// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/repository/repository.go
//
// Generated by this command:
//
//	mockgen -source=../internal/repository/repository.go -destination=repository/repository.go -package=repositoryMocks
//

// Package repositoryMocks is a generated GoMock package.
package repositoryMocks

import (
	context "context"
	reflect "reflect"
	time "time"

	model "github.com/Tel3scop/brute-force-interceptor/internal/model"
	redis "github.com/go-redis/redis/v8"
	gomock "go.uber.org/mock/gomock"
)

// MockWhiteListRepository is a mock of WhiteListRepository interface.
type MockWhiteListRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWhiteListRepositoryMockRecorder
}

// MockWhiteListRepositoryMockRecorder is the mock recorder for MockWhiteListRepository.
type MockWhiteListRepositoryMockRecorder struct {
	mock *MockWhiteListRepository
}

// NewMockWhiteListRepository creates a new mock instance.
func NewMockWhiteListRepository(ctrl *gomock.Controller) *MockWhiteListRepository {
	mock := &MockWhiteListRepository{ctrl: ctrl}
	mock.recorder = &MockWhiteListRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWhiteListRepository) EXPECT() *MockWhiteListRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWhiteListRepository) Create(ctx context.Context, dto model.WhiteList) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, dto)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWhiteListRepositoryMockRecorder) Create(ctx, dto any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWhiteListRepository)(nil).Create), ctx, dto)
}

// Delete mocks base method.
func (m *MockWhiteListRepository) Delete(ctx context.Context, subnet string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, subnet)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWhiteListRepositoryMockRecorder) Delete(ctx, subnet any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWhiteListRepository)(nil).Delete), ctx, subnet)
}

// IsInList mocks base method.
func (m *MockWhiteListRepository) IsInList(ctx context.Context, ip string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInList", ctx, ip)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInList indicates an expected call of IsInList.
func (mr *MockWhiteListRepositoryMockRecorder) IsInList(ctx, ip any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInList", reflect.TypeOf((*MockWhiteListRepository)(nil).IsInList), ctx, ip)
}

// MockBlackListRepository is a mock of BlackListRepository interface.
type MockBlackListRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBlackListRepositoryMockRecorder
}

// MockBlackListRepositoryMockRecorder is the mock recorder for MockBlackListRepository.
type MockBlackListRepositoryMockRecorder struct {
	mock *MockBlackListRepository
}

// NewMockBlackListRepository creates a new mock instance.
func NewMockBlackListRepository(ctrl *gomock.Controller) *MockBlackListRepository {
	mock := &MockBlackListRepository{ctrl: ctrl}
	mock.recorder = &MockBlackListRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlackListRepository) EXPECT() *MockBlackListRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBlackListRepository) Create(ctx context.Context, dto model.BlackList) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, dto)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBlackListRepositoryMockRecorder) Create(ctx, dto any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBlackListRepository)(nil).Create), ctx, dto)
}

// Delete mocks base method.
func (m *MockBlackListRepository) Delete(ctx context.Context, subnet string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, subnet)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBlackListRepositoryMockRecorder) Delete(ctx, subnet any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBlackListRepository)(nil).Delete), ctx, subnet)
}

// IsInList mocks base method.
func (m *MockBlackListRepository) IsInList(ctx context.Context, ip string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInList", ctx, ip)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInList indicates an expected call of IsInList.
func (mr *MockBlackListRepositoryMockRecorder) IsInList(ctx, ip any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInList", reflect.TypeOf((*MockBlackListRepository)(nil).IsInList), ctx, ip)
}

// MockBucketRepository is a mock of BucketRepository interface.
type MockBucketRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBucketRepositoryMockRecorder
}

// MockBucketRepositoryMockRecorder is the mock recorder for MockBucketRepository.
type MockBucketRepositoryMockRecorder struct {
	mock *MockBucketRepository
}

// NewMockBucketRepository creates a new mock instance.
func NewMockBucketRepository(ctrl *gomock.Controller) *MockBucketRepository {
	mock := &MockBucketRepository{ctrl: ctrl}
	mock.recorder = &MockBucketRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBucketRepository) EXPECT() *MockBucketRepositoryMockRecorder {
	return m.recorder
}

// AddRequestTimestamp mocks base method.
func (m *MockBucketRepository) AddRequestTimestamp(ctx context.Context, bucketKey string, timestamp time.Time, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRequestTimestamp", ctx, bucketKey, timestamp, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRequestTimestamp indicates an expected call of AddRequestTimestamp.
func (mr *MockBucketRepositoryMockRecorder) AddRequestTimestamp(ctx, bucketKey, timestamp, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRequestTimestamp", reflect.TypeOf((*MockBucketRepository)(nil).AddRequestTimestamp), ctx, bucketKey, timestamp, ttl)
}

// Delete mocks base method.
func (m *MockBucketRepository) Delete(ctx context.Context, bucketKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, bucketKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBucketRepositoryMockRecorder) Delete(ctx, bucketKey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBucketRepository)(nil).Delete), ctx, bucketKey)
}

// GetRequestTimestamps mocks base method.
func (m *MockBucketRepository) GetRequestTimestamps(ctx context.Context, bucketKey string) ([]time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestTimestamps", ctx, bucketKey)
	ret0, _ := ret[0].([]time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestTimestamps indicates an expected call of GetRequestTimestamps.
func (mr *MockBucketRepositoryMockRecorder) GetRequestTimestamps(ctx, bucketKey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestTimestamps", reflect.TypeOf((*MockBucketRepository)(nil).GetRequestTimestamps), ctx, bucketKey)
}

// UsePipeline mocks base method.
func (m *MockBucketRepository) UsePipeline(ctx context.Context, fn func(redis.Pipeliner) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsePipeline", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// UsePipeline indicates an expected call of UsePipeline.
func (mr *MockBucketRepositoryMockRecorder) UsePipeline(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsePipeline", reflect.TypeOf((*MockBucketRepository)(nil).UsePipeline), ctx, fn)
}
