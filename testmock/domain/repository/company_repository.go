// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/company_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/kohge2/upsdct-server/domain/models"
)

// MockCompanyRepository is a mock of CompanyRepository interface.
type MockCompanyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyRepositoryMockRecorder
}

// MockCompanyRepositoryMockRecorder is the mock recorder for MockCompanyRepository.
type MockCompanyRepositoryMockRecorder struct {
	mock *MockCompanyRepository
}

// NewMockCompanyRepository creates a new mock instance.
func NewMockCompanyRepository(ctrl *gomock.Controller) *MockCompanyRepository {
	mock := &MockCompanyRepository{ctrl: ctrl}
	mock.recorder = &MockCompanyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompanyRepository) EXPECT() *MockCompanyRepositoryMockRecorder {
	return m.recorder
}

// FindByCompanyID mocks base method.
func (m *MockCompanyRepository) FindByCompanyID(ctx context.Context, companyID string) (*models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCompanyID", ctx, companyID)
	ret0, _ := ret[0].(*models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCompanyID indicates an expected call of FindByCompanyID.
func (mr *MockCompanyRepositoryMockRecorder) FindByCompanyID(ctx, companyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCompanyID", reflect.TypeOf((*MockCompanyRepository)(nil).FindByCompanyID), ctx, companyID)
}
