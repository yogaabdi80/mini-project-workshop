// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	"workshop-golang-psm/food-app-service/internal/domain/entity"

	gomock "github.com/golang/mock/gomock"
)

type MockIFoodRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIFoodRepositoryMockRecorder
}

type MockIFoodRepositoryMockRecorder struct {
	mock *MockIFoodRepository
}

func NewMockIFoodRepository(ctrl *gomock.Controller) *MockIFoodRepository {
	mock := &MockIFoodRepository{ctrl: ctrl}
	mock.recorder = &MockIFoodRepositoryMockRecorder{mock}
	return mock
}

func (m *MockIFoodRepository) EXPECT() *MockIFoodRepositoryMockRecorder {
	return m.recorder
}

func (m *MockIFoodRepository) SaveFood(food *entity.Food) (*entity.Food, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFood", food)
	ret0, _ := ret[0].(*entity.Food)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIFoodRepositoryMockRecorder) SaveFood(food interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFood", reflect.TypeOf((*MockIFoodRepository)(nil).SaveFood), food)
}

func (m *MockIFoodRepository) GetDetailFood(id int) (*entity.Food, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailFood", id)
	ret0, _ := ret[0].(*entity.Food)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIFoodRepositoryMockRecorder) GetDetailFood(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailFood", reflect.TypeOf((*MockIFoodRepository)(nil).GetDetailFood), id)
}

func (m *MockIFoodRepository) GetAllFood() ([]entity.Food, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFood")
	ret0, _ := ret[0].([]entity.Food)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIFoodRepositoryMockRecorder) GetAllFood() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFood", reflect.TypeOf((*MockIFoodRepository)(nil).GetAllFood))
}

func (m *MockIFoodRepository) UpdateFood(food *entity.Food) (*entity.Food, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFood", food)
	ret0, _ := ret[0].(*entity.Food)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIFoodRepositoryMockRecorder) UpdateFood(food interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFood", reflect.TypeOf((*MockIFoodRepository)(nil).UpdateFood), food)
}

func (m *MockIFoodRepository) DeleteFood(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFood", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockIFoodRepositoryMockRecorder) DeleteFood(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFood", reflect.TypeOf((*MockIFoodRepository)(nil).DeleteFood), id)
}
