// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	github "github.com/google/go-github/v52/github"

	mock "github.com/stretchr/testify/mock"
)

// GHAppClient is an autogenerated mock type for the GHAppClient type
type GHAppClient struct {
	mock.Mock
}

type GHAppClient_Expecter struct {
	mock *mock.Mock
}

func (_m *GHAppClient) EXPECT() *GHAppClient_Expecter {
	return &GHAppClient_Expecter{mock: &_m.Mock}
}

// CloneRepo provides a mock function with given fields: ctx, owner, repo, branch, dir, isPublic
func (_m *GHAppClient) CloneRepo(ctx context.Context, owner string, repo string, branch string, dir string, isPublic bool) error {
	ret := _m.Called(ctx, owner, repo, branch, dir, isPublic)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, bool) error); ok {
		r0 = rf(ctx, owner, repo, branch, dir, isPublic)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GHAppClient_CloneRepo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloneRepo'
type GHAppClient_CloneRepo_Call struct {
	*mock.Call
}

// CloneRepo is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - branch string
//   - dir string
//   - isPublic bool
func (_e *GHAppClient_Expecter) CloneRepo(ctx interface{}, owner interface{}, repo interface{}, branch interface{}, dir interface{}, isPublic interface{}) *GHAppClient_CloneRepo_Call {
	return &GHAppClient_CloneRepo_Call{Call: _e.mock.On("CloneRepo", ctx, owner, repo, branch, dir, isPublic)}
}

func (_c *GHAppClient_CloneRepo_Call) Run(run func(ctx context.Context, owner string, repo string, branch string, dir string, isPublic bool)) *GHAppClient_CloneRepo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string), args[5].(bool))
	})
	return _c
}

func (_c *GHAppClient_CloneRepo_Call) Return(_a0 error) *GHAppClient_CloneRepo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GHAppClient_CloneRepo_Call) RunAndReturn(run func(context.Context, string, string, string, string, bool) error) *GHAppClient_CloneRepo_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCommitStatus provides a mock function with given fields: ctx, owner, repo, sha, state, targetURL
func (_m *GHAppClient) CreateCommitStatus(ctx context.Context, owner string, repo string, sha string, state string, targetURL *string) error {
	ret := _m.Called(ctx, owner, repo, sha, state, targetURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, *string) error); ok {
		r0 = rf(ctx, owner, repo, sha, state, targetURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GHAppClient_CreateCommitStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCommitStatus'
type GHAppClient_CreateCommitStatus_Call struct {
	*mock.Call
}

// CreateCommitStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - sha string
//   - state string
//   - targetURL *string
func (_e *GHAppClient_Expecter) CreateCommitStatus(ctx interface{}, owner interface{}, repo interface{}, sha interface{}, state interface{}, targetURL interface{}) *GHAppClient_CreateCommitStatus_Call {
	return &GHAppClient_CreateCommitStatus_Call{Call: _e.mock.On("CreateCommitStatus", ctx, owner, repo, sha, state, targetURL)}
}

func (_c *GHAppClient_CreateCommitStatus_Call) Run(run func(ctx context.Context, owner string, repo string, sha string, state string, targetURL *string)) *GHAppClient_CreateCommitStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string), args[5].(*string))
	})
	return _c
}

func (_c *GHAppClient_CreateCommitStatus_Call) Return(_a0 error) *GHAppClient_CreateCommitStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GHAppClient_CreateCommitStatus_Call) RunAndReturn(run func(context.Context, string, string, string, string, *string) error) *GHAppClient_CreateCommitStatus_Call {
	_c.Call.Return(run)
	return _c
}

// CreatePullRequest provides a mock function with given fields: ctx, owner, repo, branchPrefix, changes, title, description
func (_m *GHAppClient) CreatePullRequest(ctx context.Context, owner string, repo string, branchPrefix string, changes map[string]string, title string, description string) (*github.PullRequest, error) {
	ret := _m.Called(ctx, owner, repo, branchPrefix, changes, title, description)

	var r0 *github.PullRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]string, string, string) (*github.PullRequest, error)); ok {
		return rf(ctx, owner, repo, branchPrefix, changes, title, description)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]string, string, string) *github.PullRequest); ok {
		r0 = rf(ctx, owner, repo, branchPrefix, changes, title, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.PullRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, map[string]string, string, string) error); ok {
		r1 = rf(ctx, owner, repo, branchPrefix, changes, title, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_CreatePullRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePullRequest'
type GHAppClient_CreatePullRequest_Call struct {
	*mock.Call
}

// CreatePullRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - branchPrefix string
//   - changes map[string]string
//   - title string
//   - description string
func (_e *GHAppClient_Expecter) CreatePullRequest(ctx interface{}, owner interface{}, repo interface{}, branchPrefix interface{}, changes interface{}, title interface{}, description interface{}) *GHAppClient_CreatePullRequest_Call {
	return &GHAppClient_CreatePullRequest_Call{Call: _e.mock.On("CreatePullRequest", ctx, owner, repo, branchPrefix, changes, title, description)}
}

func (_c *GHAppClient_CreatePullRequest_Call) Run(run func(ctx context.Context, owner string, repo string, branchPrefix string, changes map[string]string, title string, description string)) *GHAppClient_CreatePullRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(map[string]string), args[5].(string), args[6].(string))
	})
	return _c
}

func (_c *GHAppClient_CreatePullRequest_Call) Return(_a0 *github.PullRequest, _a1 error) *GHAppClient_CreatePullRequest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_CreatePullRequest_Call) RunAndReturn(run func(context.Context, string, string, string, map[string]string, string, string) (*github.PullRequest, error)) *GHAppClient_CreatePullRequest_Call {
	_c.Call.Return(run)
	return _c
}

// DoesBranchExist provides a mock function with given fields: ctx, owner, repo, branch, branchOwner
func (_m *GHAppClient) DoesBranchExist(ctx context.Context, owner string, repo string, branch string, branchOwner string) (bool, error) {
	ret := _m.Called(ctx, owner, repo, branch, branchOwner)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) (bool, error)); ok {
		return rf(ctx, owner, repo, branch, branchOwner)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) bool); ok {
		r0 = rf(ctx, owner, repo, branch, branchOwner)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, owner, repo, branch, branchOwner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_DoesBranchExist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DoesBranchExist'
type GHAppClient_DoesBranchExist_Call struct {
	*mock.Call
}

// DoesBranchExist is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - branch string
//   - branchOwner string
func (_e *GHAppClient_Expecter) DoesBranchExist(ctx interface{}, owner interface{}, repo interface{}, branch interface{}, branchOwner interface{}) *GHAppClient_DoesBranchExist_Call {
	return &GHAppClient_DoesBranchExist_Call{Call: _e.mock.On("DoesBranchExist", ctx, owner, repo, branch, branchOwner)}
}

func (_c *GHAppClient_DoesBranchExist_Call) Run(run func(ctx context.Context, owner string, repo string, branch string, branchOwner string)) *GHAppClient_DoesBranchExist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *GHAppClient_DoesBranchExist_Call) Return(_a0 bool, _a1 error) *GHAppClient_DoesBranchExist_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_DoesBranchExist_Call) RunAndReturn(run func(context.Context, string, string, string, string) (bool, error)) *GHAppClient_DoesBranchExist_Call {
	_c.Call.Return(run)
	return _c
}

// GetBranchSHA provides a mock function with given fields: ctx, owner, repo, branch
func (_m *GHAppClient) GetBranchSHA(ctx context.Context, owner string, repo string, branch string) (string, error) {
	ret := _m.Called(ctx, owner, repo, branch)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (string, error)); ok {
		return rf(ctx, owner, repo, branch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) string); ok {
		r0 = rf(ctx, owner, repo, branch)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, owner, repo, branch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_GetBranchSHA_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBranchSHA'
type GHAppClient_GetBranchSHA_Call struct {
	*mock.Call
}

// GetBranchSHA is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - branch string
func (_e *GHAppClient_Expecter) GetBranchSHA(ctx interface{}, owner interface{}, repo interface{}, branch interface{}) *GHAppClient_GetBranchSHA_Call {
	return &GHAppClient_GetBranchSHA_Call{Call: _e.mock.On("GetBranchSHA", ctx, owner, repo, branch)}
}

func (_c *GHAppClient_GetBranchSHA_Call) Run(run func(ctx context.Context, owner string, repo string, branch string)) *GHAppClient_GetBranchSHA_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *GHAppClient_GetBranchSHA_Call) Return(_a0 string, _a1 error) *GHAppClient_GetBranchSHA_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_GetBranchSHA_Call) RunAndReturn(run func(context.Context, string, string, string) (string, error)) *GHAppClient_GetBranchSHA_Call {
	_c.Call.Return(run)
	return _c
}

// GetCloneParams provides a mock function with given fields:
func (_m *GHAppClient) GetCloneParams() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GHAppClient_GetCloneParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCloneParams'
type GHAppClient_GetCloneParams_Call struct {
	*mock.Call
}

// GetCloneParams is a helper method to define mock.On call
func (_e *GHAppClient_Expecter) GetCloneParams() *GHAppClient_GetCloneParams_Call {
	return &GHAppClient_GetCloneParams_Call{Call: _e.mock.On("GetCloneParams")}
}

func (_c *GHAppClient_GetCloneParams_Call) Run(run func()) *GHAppClient_GetCloneParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GHAppClient_GetCloneParams_Call) Return(_a0 []string) *GHAppClient_GetCloneParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GHAppClient_GetCloneParams_Call) RunAndReturn(run func() []string) *GHAppClient_GetCloneParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetCloneToken provides a mock function with given fields: ctx, owner, repo
func (_m *GHAppClient) GetCloneToken(ctx context.Context, owner string, repo string) (string, error) {
	ret := _m.Called(ctx, owner, repo)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, owner, repo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, owner, repo)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, owner, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_GetCloneToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCloneToken'
type GHAppClient_GetCloneToken_Call struct {
	*mock.Call
}

// GetCloneToken is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
func (_e *GHAppClient_Expecter) GetCloneToken(ctx interface{}, owner interface{}, repo interface{}) *GHAppClient_GetCloneToken_Call {
	return &GHAppClient_GetCloneToken_Call{Call: _e.mock.On("GetCloneToken", ctx, owner, repo)}
}

func (_c *GHAppClient_GetCloneToken_Call) Run(run func(ctx context.Context, owner string, repo string)) *GHAppClient_GetCloneToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *GHAppClient_GetCloneToken_Call) Return(_a0 string, _a1 error) *GHAppClient_GetCloneToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_GetCloneToken_Call) RunAndReturn(run func(context.Context, string, string) (string, error)) *GHAppClient_GetCloneToken_Call {
	_c.Call.Return(run)
	return _c
}

// GetCloneUrl provides a mock function with given fields:
func (_m *GHAppClient) GetCloneUrl() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GHAppClient_GetCloneUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCloneUrl'
type GHAppClient_GetCloneUrl_Call struct {
	*mock.Call
}

// GetCloneUrl is a helper method to define mock.On call
func (_e *GHAppClient_Expecter) GetCloneUrl() *GHAppClient_GetCloneUrl_Call {
	return &GHAppClient_GetCloneUrl_Call{Call: _e.mock.On("GetCloneUrl")}
}

func (_c *GHAppClient_GetCloneUrl_Call) Run(run func()) *GHAppClient_GetCloneUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GHAppClient_GetCloneUrl_Call) Return(_a0 string) *GHAppClient_GetCloneUrl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GHAppClient_GetCloneUrl_Call) RunAndReturn(run func() string) *GHAppClient_GetCloneUrl_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultBranch provides a mock function with given fields: ctx, owner, repo, branchOwner
func (_m *GHAppClient) GetDefaultBranch(ctx context.Context, owner string, repo string, branchOwner string) (string, error) {
	ret := _m.Called(ctx, owner, repo, branchOwner)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (string, error)); ok {
		return rf(ctx, owner, repo, branchOwner)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) string); ok {
		r0 = rf(ctx, owner, repo, branchOwner)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, owner, repo, branchOwner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_GetDefaultBranch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultBranch'
type GHAppClient_GetDefaultBranch_Call struct {
	*mock.Call
}

// GetDefaultBranch is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - branchOwner string
func (_e *GHAppClient_Expecter) GetDefaultBranch(ctx interface{}, owner interface{}, repo interface{}, branchOwner interface{}) *GHAppClient_GetDefaultBranch_Call {
	return &GHAppClient_GetDefaultBranch_Call{Call: _e.mock.On("GetDefaultBranch", ctx, owner, repo, branchOwner)}
}

func (_c *GHAppClient_GetDefaultBranch_Call) Run(run func(ctx context.Context, owner string, repo string, branchOwner string)) *GHAppClient_GetDefaultBranch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *GHAppClient_GetDefaultBranch_Call) Return(_a0 string, _a1 error) *GHAppClient_GetDefaultBranch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_GetDefaultBranch_Call) RunAndReturn(run func(context.Context, string, string, string) (string, error)) *GHAppClient_GetDefaultBranch_Call {
	_c.Call.Return(run)
	return _c
}

// GetInstallation provides a mock function with given fields: ctx, installationID
func (_m *GHAppClient) GetInstallation(ctx context.Context, installationID int64) (*github.Installation, error) {
	ret := _m.Called(ctx, installationID)

	var r0 *github.Installation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*github.Installation, error)); ok {
		return rf(ctx, installationID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *github.Installation); ok {
		r0 = rf(ctx, installationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Installation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, installationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_GetInstallation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInstallation'
type GHAppClient_GetInstallation_Call struct {
	*mock.Call
}

// GetInstallation is a helper method to define mock.On call
//   - ctx context.Context
//   - installationID int64
func (_e *GHAppClient_Expecter) GetInstallation(ctx interface{}, installationID interface{}) *GHAppClient_GetInstallation_Call {
	return &GHAppClient_GetInstallation_Call{Call: _e.mock.On("GetInstallation", ctx, installationID)}
}

func (_c *GHAppClient_GetInstallation_Call) Run(run func(ctx context.Context, installationID int64)) *GHAppClient_GetInstallation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *GHAppClient_GetInstallation_Call) Return(_a0 *github.Installation, _a1 error) *GHAppClient_GetInstallation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_GetInstallation_Call) RunAndReturn(run func(context.Context, int64) (*github.Installation, error)) *GHAppClient_GetInstallation_Call {
	_c.Call.Return(run)
	return _c
}

// IsOwnerInstalled provides a mock function with given fields: ctx, owner
func (_m *GHAppClient) IsOwnerInstalled(ctx context.Context, owner string) (bool, error) {
	ret := _m.Called(ctx, owner)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, owner)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, owner)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_IsOwnerInstalled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsOwnerInstalled'
type GHAppClient_IsOwnerInstalled_Call struct {
	*mock.Call
}

// IsOwnerInstalled is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
func (_e *GHAppClient_Expecter) IsOwnerInstalled(ctx interface{}, owner interface{}) *GHAppClient_IsOwnerInstalled_Call {
	return &GHAppClient_IsOwnerInstalled_Call{Call: _e.mock.On("IsOwnerInstalled", ctx, owner)}
}

func (_c *GHAppClient_IsOwnerInstalled_Call) Run(run func(ctx context.Context, owner string)) *GHAppClient_IsOwnerInstalled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *GHAppClient_IsOwnerInstalled_Call) Return(_a0 bool, _a1 error) *GHAppClient_IsOwnerInstalled_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_IsOwnerInstalled_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *GHAppClient_IsOwnerInstalled_Call {
	_c.Call.Return(run)
	return _c
}

// IsRepoPrivate provides a mock function with given fields: ctx, owner, repo
func (_m *GHAppClient) IsRepoPrivate(ctx context.Context, owner string, repo string) (bool, error) {
	ret := _m.Called(ctx, owner, repo)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, owner, repo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, owner, repo)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, owner, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_IsRepoPrivate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRepoPrivate'
type GHAppClient_IsRepoPrivate_Call struct {
	*mock.Call
}

// IsRepoPrivate is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
func (_e *GHAppClient_Expecter) IsRepoPrivate(ctx interface{}, owner interface{}, repo interface{}) *GHAppClient_IsRepoPrivate_Call {
	return &GHAppClient_IsRepoPrivate_Call{Call: _e.mock.On("IsRepoPrivate", ctx, owner, repo)}
}

func (_c *GHAppClient_IsRepoPrivate_Call) Run(run func(ctx context.Context, owner string, repo string)) *GHAppClient_IsRepoPrivate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *GHAppClient_IsRepoPrivate_Call) Return(_a0 bool, _a1 error) *GHAppClient_IsRepoPrivate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_IsRepoPrivate_Call) RunAndReturn(run func(context.Context, string, string) (bool, error)) *GHAppClient_IsRepoPrivate_Call {
	_c.Call.Return(run)
	return _c
}

// ListBranches provides a mock function with given fields: ctx, owner, repo
func (_m *GHAppClient) ListBranches(ctx context.Context, owner string, repo string) ([]string, error) {
	ret := _m.Called(ctx, owner, repo)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]string, error)); ok {
		return rf(ctx, owner, repo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []string); ok {
		r0 = rf(ctx, owner, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, owner, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_ListBranches_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListBranches'
type GHAppClient_ListBranches_Call struct {
	*mock.Call
}

// ListBranches is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
func (_e *GHAppClient_Expecter) ListBranches(ctx interface{}, owner interface{}, repo interface{}) *GHAppClient_ListBranches_Call {
	return &GHAppClient_ListBranches_Call{Call: _e.mock.On("ListBranches", ctx, owner, repo)}
}

func (_c *GHAppClient_ListBranches_Call) Run(run func(ctx context.Context, owner string, repo string)) *GHAppClient_ListBranches_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *GHAppClient_ListBranches_Call) Return(_a0 []string, _a1 error) *GHAppClient_ListBranches_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_ListBranches_Call) RunAndReturn(run func(context.Context, string, string) ([]string, error)) *GHAppClient_ListBranches_Call {
	_c.Call.Return(run)
	return _c
}

// ListInstalledOwners provides a mock function with given fields: ctx
func (_m *GHAppClient) ListInstalledOwners(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_ListInstalledOwners_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListInstalledOwners'
type GHAppClient_ListInstalledOwners_Call struct {
	*mock.Call
}

// ListInstalledOwners is a helper method to define mock.On call
//   - ctx context.Context
func (_e *GHAppClient_Expecter) ListInstalledOwners(ctx interface{}) *GHAppClient_ListInstalledOwners_Call {
	return &GHAppClient_ListInstalledOwners_Call{Call: _e.mock.On("ListInstalledOwners", ctx)}
}

func (_c *GHAppClient_ListInstalledOwners_Call) Run(run func(ctx context.Context)) *GHAppClient_ListInstalledOwners_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *GHAppClient_ListInstalledOwners_Call) Return(_a0 []string, _a1 error) *GHAppClient_ListInstalledOwners_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_ListInstalledOwners_Call) RunAndReturn(run func(context.Context) ([]string, error)) *GHAppClient_ListInstalledOwners_Call {
	_c.Call.Return(run)
	return _c
}

// ListOwnerInstalledRepos provides a mock function with given fields: ctx, owner
func (_m *GHAppClient) ListOwnerInstalledRepos(ctx context.Context, owner string) ([]*github.Repository, error) {
	ret := _m.Called(ctx, owner)

	var r0 []*github.Repository
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*github.Repository, error)); ok {
		return rf(ctx, owner)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*github.Repository); ok {
		r0 = rf(ctx, owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.Repository)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_ListOwnerInstalledRepos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOwnerInstalledRepos'
type GHAppClient_ListOwnerInstalledRepos_Call struct {
	*mock.Call
}

// ListOwnerInstalledRepos is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
func (_e *GHAppClient_Expecter) ListOwnerInstalledRepos(ctx interface{}, owner interface{}) *GHAppClient_ListOwnerInstalledRepos_Call {
	return &GHAppClient_ListOwnerInstalledRepos_Call{Call: _e.mock.On("ListOwnerInstalledRepos", ctx, owner)}
}

func (_c *GHAppClient_ListOwnerInstalledRepos_Call) Run(run func(ctx context.Context, owner string)) *GHAppClient_ListOwnerInstalledRepos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *GHAppClient_ListOwnerInstalledRepos_Call) Return(_a0 []*github.Repository, _a1 error) *GHAppClient_ListOwnerInstalledRepos_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_ListOwnerInstalledRepos_Call) RunAndReturn(run func(context.Context, string) ([]*github.Repository, error)) *GHAppClient_ListOwnerInstalledRepos_Call {
	_c.Call.Return(run)
	return _c
}

// UpsertComment provides a mock function with given fields: ctx, owner, repo, prNumber, commentID, comment
func (_m *GHAppClient) UpsertComment(ctx context.Context, owner string, repo string, prNumber int, commentID int64, comment string) (*github.IssueComment, error) {
	ret := _m.Called(ctx, owner, repo, prNumber, commentID, comment)

	var r0 *github.IssueComment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int64, string) (*github.IssueComment, error)); ok {
		return rf(ctx, owner, repo, prNumber, commentID, comment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int64, string) *github.IssueComment); ok {
		r0 = rf(ctx, owner, repo, prNumber, commentID, comment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.IssueComment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, int64, string) error); ok {
		r1 = rf(ctx, owner, repo, prNumber, commentID, comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GHAppClient_UpsertComment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertComment'
type GHAppClient_UpsertComment_Call struct {
	*mock.Call
}

// UpsertComment is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - prNumber int
//   - commentID int64
//   - comment string
func (_e *GHAppClient_Expecter) UpsertComment(ctx interface{}, owner interface{}, repo interface{}, prNumber interface{}, commentID interface{}, comment interface{}) *GHAppClient_UpsertComment_Call {
	return &GHAppClient_UpsertComment_Call{Call: _e.mock.On("UpsertComment", ctx, owner, repo, prNumber, commentID, comment)}
}

func (_c *GHAppClient_UpsertComment_Call) Run(run func(ctx context.Context, owner string, repo string, prNumber int, commentID int64, comment string)) *GHAppClient_UpsertComment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(int), args[4].(int64), args[5].(string))
	})
	return _c
}

func (_c *GHAppClient_UpsertComment_Call) Return(_a0 *github.IssueComment, _a1 error) *GHAppClient_UpsertComment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GHAppClient_UpsertComment_Call) RunAndReturn(run func(context.Context, string, string, int, int64, string) (*github.IssueComment, error)) *GHAppClient_UpsertComment_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewGHAppClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewGHAppClient creates a new instance of GHAppClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGHAppClient(t mockConstructorTestingTNewGHAppClient) *GHAppClient {
	mock := &GHAppClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
