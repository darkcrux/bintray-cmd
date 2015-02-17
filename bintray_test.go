package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockBintrayClient struct {
	mock.Mock
}

func (mock *MockBintrayClient) PackageExists(subject, repo, pkg string) (bool, error) {
	args := mock.Called(subject, repo, pkg)
	return args.Bool(0), args.Error(1)
}

func (mock *MockBintrayClient) GetVersions(subject, repo, pkg string) ([]string, error) {
	args := mock.Called(subject, repo, pkg)
	return nil, args.Error(0)
}

func (mock *MockBintrayClient) CreateVersion(subject, repo, pkg, version string) error {
	args := mock.Called(subject, repo, pkg, version)
	return args.Error(0)
}

func (mock *MockBintrayClient) UploadFile(subject, repo, pkg, version, prjId, prjName, filePath string, mavenRepo bool) error {
	args := mock.Called(subject, repo, pkg, version, prjId, prjName, filePath, mavenRepo)
	return args.Error(0)
}

func (mock *MockBintrayClient) Publish(subject, repo, pkg, version string) error {
	args := mock.Called(subject, repo, pkg, version)
	return args.Error(0)
}

func TestExists(t *testing.T) {
	mockClient := new(MockBintrayClient)
	bp := &bintrayPackage{bintrayClient: mockClient}

	mockClient.On("PackageExists", mock.Anything, mock.Anything, "exists").Return(true, nil)
	mockClient.On("PackageExists", mock.Anything, mock.Anything, "not-exists").Return(false, nil)
	mockClient.On("PackageExists", mock.Anything, mock.Anything, mock.Anything).Return(false, errors.New("error"))

	bp.pkg = "exists"
	bp.exists()

	bp.pkg = "not-exists"
	bp.exists()

	mockClient.AssertExpectations(t)
}

func TestListVersions(t *testing.T) {
	mockClient := new(MockBintrayClient)
	bp := &bintrayPackage{bintrayClient: mockClient}

	mockClient.On("GetVersions", mock.Anything, mock.Anything, "exists").Return(nil)
	mockClient.On("GetVersions", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
	bp.pkg = "exists"
	bp.listVersions()
	mockClient.AssertExpectations(t)
}

func TestCreateVersion(t *testing.T) {
	mockClient := new(MockBintrayClient)
	bp := &bintrayPackage{bintrayClient: mockClient}

	mockClient.On("CreateVersion", mock.Anything, mock.Anything, mock.Anything, "v1.2").Return(nil)
	mockClient.On("CreateVersion", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
	bp.createVersion("v1.2")
	mockClient.AssertExpectations(t)
}

func TestUploadFile(t *testing.T) {
	mockClient := new(MockBintrayClient)
	bp := &bintrayPackage{bintrayClient: mockClient}

	mockClient.On("UploadFile", mock.Anything, mock.Anything, mock.Anything, "v1.2", "hello", "kitty", "/somewhere/there", false).Return(nil)
	mockClient.On("UploadFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, false).Return(errors.New("error"))
	bp.uploadFile("v1.2", "hello", "kitty", "/somewhere/there", false)
	mockClient.AssertExpectations(t)
}

func TestPublish(t *testing.T) {
	mockClient := new(MockBintrayClient)
	bp := &bintrayPackage{bintrayClient: mockClient}

	mockClient.On("Publish", mock.Anything, mock.Anything, mock.Anything, "v1.2").Return(nil)
	mockClient.On("Publish", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
	bp.publish("v1.2")
	mockClient.AssertExpectations(t)
}
