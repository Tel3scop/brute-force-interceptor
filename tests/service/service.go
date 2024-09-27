// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/service/service.go
//
// Generated by this command:
//
//	mockgen -source=../internal/service/service.go -destination=service/service.go -package=serviceMocks
//

// Package serviceMocks is a generated GoMock package.
package serviceMocks

import (
	context "context"
	reflect "reflect"
	time "time"

	model "github.com/Tel3scop/brute-force-interceptor/internal/model"
	redis "github.com/go-redis/redis/v8"
	gomock "go.uber.org/mock/gomock"
)

// MockBucketService is a mock of BucketService interface.
type MockBucketService struct {
	ctrl     *gomock.Controller
	recorder *MockBucketServiceMockRecorder
}

// MockBucketServiceMockRecorder is the mock recorder for MockBucketService.
type MockBucketServiceMockRecorder struct {
	mock *MockBucketService
}

// NewMockBucketService creates a new mock instance.
func NewMockBucketService(ctrl *gomock.Controller) *MockBucketService {
	mock := &MockBucketService{ctrl: ctrl}
	mock.recorder = &MockBucketServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBucketService) EXPECT() *MockBucketServiceMockRecorder {
	return m.recorder
}

// AddRequestTimestamps mocks base method.
func (m *MockBucketService) AddRequestTimestamps(ctx context.Context, auth model.Auth, time time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRequestTimestamps", ctx, auth, time)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRequestTimestamps indicates an expected call of AddRequestTimestamps.
func (mr *MockBucketServiceMockRecorder) AddRequestTimestamps(ctx, auth, time any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRequestTimestamps", reflect.TypeOf((*MockBucketService)(nil).AddRequestTimestamps), ctx, auth, time)
}

// CheckBucketLimit mocks base method.
func (m *MockBucketService) CheckBucketLimit(ctx context.Context, bucketKey string, time time.Time, limit int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBucketLimit", ctx, bucketKey, time, limit)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckBucketLimit indicates an expected call of CheckBucketLimit.
func (mr *MockBucketServiceMockRecorder) CheckBucketLimit(ctx, bucketKey, time, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBucketLimit", reflect.TypeOf((*MockBucketService)(nil).CheckBucketLimit), ctx, bucketKey, time, limit)
}

// GetRequestTimestamps mocks base method.
func (m *MockBucketService) GetRequestTimestamps(ctx context.Context, bucketKey string) ([]time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestTimestamps", ctx, bucketKey)
	ret0, _ := ret[0].([]time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestTimestamps indicates an expected call of GetRequestTimestamps.
func (mr *MockBucketServiceMockRecorder) GetRequestTimestamps(ctx, bucketKey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestTimestamps", reflect.TypeOf((*MockBucketService)(nil).GetRequestTimestamps), ctx, bucketKey)
}

// Reset mocks base method.
func (m *MockBucketService) Reset(ctx context.Context, bucket string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset", ctx, bucket)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset.
func (mr *MockBucketServiceMockRecorder) Reset(ctx, bucket any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockBucketService)(nil).Reset), ctx, bucket)
}

// ResetBuckets mocks base method.
func (m *MockBucketService) ResetBuckets(ctx context.Context, bucketKeys []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetBuckets", ctx, bucketKeys)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetBuckets indicates an expected call of ResetBuckets.
func (mr *MockBucketServiceMockRecorder) ResetBuckets(ctx, bucketKeys any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetBuckets", reflect.TypeOf((*MockBucketService)(nil).ResetBuckets), ctx, bucketKeys)
}

// UsePipeline mocks base method.
func (m *MockBucketService) UsePipeline(ctx context.Context, fn func(redis.Pipeliner) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsePipeline", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// UsePipeline indicates an expected call of UsePipeline.
func (mr *MockBucketServiceMockRecorder) UsePipeline(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsePipeline", reflect.TypeOf((*MockBucketService)(nil).UsePipeline), ctx, fn)
}

// MockAccessService is a mock of AccessService interface.
type MockAccessService struct {
	ctrl     *gomock.Controller
	recorder *MockAccessServiceMockRecorder
}

// MockAccessServiceMockRecorder is the mock recorder for MockAccessService.
type MockAccessServiceMockRecorder struct {
	mock *MockAccessService
}

// NewMockAccessService creates a new mock instance.
func NewMockAccessService(ctrl *gomock.Controller) *MockAccessService {
	mock := &MockAccessService{ctrl: ctrl}
	mock.recorder = &MockAccessServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessService) EXPECT() *MockAccessServiceMockRecorder {
	return m.recorder
}

// AddToBlackList mocks base method.
func (m *MockAccessService) AddToBlackList(ctx context.Context, subnet string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToBlackList", ctx, subnet)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToBlackList indicates an expected call of AddToBlackList.
func (mr *MockAccessServiceMockRecorder) AddToBlackList(ctx, subnet any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToBlackList", reflect.TypeOf((*MockAccessService)(nil).AddToBlackList), ctx, subnet)
}

// AddToWhiteList mocks base method.
func (m *MockAccessService) AddToWhiteList(ctx context.Context, subnet string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToWhiteList", ctx, subnet)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToWhiteList indicates an expected call of AddToWhiteList.
func (mr *MockAccessServiceMockRecorder) AddToWhiteList(ctx, subnet any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToWhiteList", reflect.TypeOf((*MockAccessService)(nil).AddToWhiteList), ctx, subnet)
}

// Check mocks base method.
func (m *MockAccessService) Check(ctx context.Context, auth model.Auth) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", ctx, auth)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockAccessServiceMockRecorder) Check(ctx, auth any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockAccessService)(nil).Check), ctx, auth)
}
