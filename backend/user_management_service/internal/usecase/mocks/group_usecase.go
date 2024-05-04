package mocks

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type IGroupUsecase interface {
	CreateNewGroup(groupName, description string, creatorID uuid.UUID) (*model.Group, error)
	InviteUserToGroup(inviterID, groupID uuid.UUID, inviteeEmail string) (*model.Invitation, error)
	AcceptInvitation(inviteeID, invitationID uuid.UUID) error
	RejectInvitation(inviteeID, invitationID uuid.UUID) error
	CancelInvitation(inviterID, invitationID uuid.UUID) error
	RemoveUserFromGroup(inviterID, removeUserID, groupID uuid.UUID) error
	GetGroupDetails(groupID uuid.UUID) (*model.Group, error)
	GetGroupMembers(groupID uuid.UUID) ([]*model.User, error)
	DeleteGroup(groupID uuid.UUID) error
}

type MockGroupUsecase struct {
	mock.Mock
}

func (m *MockGroupUsecase) CreateNewGroup(groupName, description string, creatorID uuid.UUID) (*model.Group, error) {
	args := m.Called(groupName, description, creatorID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Group), args.Error(1)
}

func (m *MockGroupUsecase) InviteUserToGroup(inviterID, groupID uuid.UUID, inviteeEmail string) (*model.Invitation, error) {
	args := m.Called(inviterID, groupID, inviteeEmail)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Invitation), args.Error(1)
}

func (m *MockGroupUsecase) AcceptInvitation(inviteeID, invitationID uuid.UUID) error {
	args := m.Called(inviteeID, invitationID)
	return args.Error(0)
}

func (m *MockGroupUsecase) RejectInvitation(inviteeID, invitationID uuid.UUID) error {
	args := m.Called(inviteeID, invitationID)
	return args.Error(0)
}

func (m *MockGroupUsecase) CancelInvitation(inviterID, invitationID uuid.UUID) error {
	args := m.Called(inviterID, invitationID)
	return args.Error(0)
}

func (m *MockGroupUsecase) RemoveUserFromGroup(inviterID, removeUserID, groupID uuid.UUID) error {
	args := m.Called(inviterID, removeUserID, groupID)
	return args.Error(0)
}

func (m *MockGroupUsecase) GetGroupDetails(groupID uuid.UUID) (*model.Group, error) {
	args := m.Called(groupID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Group), args.Error(1)
}

func (m *MockGroupUsecase) GetGroupMembers(groupID uuid.UUID) ([]*model.User, error) {
	args := m.Called(groupID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.User), args.Error(1)
}

func (m *MockGroupUsecase) DeleteGroup(groupID uuid.UUID) error {
	args := m.Called(groupID)
	return args.Error(0)
}
