package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
)

type IGroupService interface {
	CreateGroup(groupName, description string) (*model.Group, error)
	UpdateGroup(id uuid.UUID, groupName, description string) (*model.Group, error)
	DeleteGroup(id uuid.UUID) error
	FindByID(id uuid.UUID) (*model.Group, error)
	FindByName(name string) (*model.Group, error)
	GetAllGroups() ([]*model.Group, error)
}

type groupService struct {
	gr repository.IGroupRepository
}

func NewGroupService(gr repository.IGroupRepository) IGroupService {
	return &groupService{gr}
}

func (gs *groupService) CreateGroup(groupName, description string) (*model.Group, error) {
	group := &model.Group{
		ID:          uuid.New(),
		GroupName:   groupName,
		Description: description,
	}
	newGroup, err := gs.gr.Create(group)
	if err != nil {
		return nil, err
	}
	return newGroup, nil
}

func (gs *groupService) UpdateGroup(id uuid.UUID, groupName, description string) (*model.Group, error) {
	group, err := gs.gr.FindByID(id)
	if err != nil {
		return nil, err
	}

	if groupName != "" {
		group.GroupName = groupName
	}

	if description != "" {
		group.Description = description
	}

	if groupName != "" || description != "" {
		group.UpdatedAt = time.Now()
		updatedUser, err := gs.gr.Update(group)
		if err != nil {
			return nil, err
		}
		return updatedUser, nil
	}
	return group, nil
}

func (gs *groupService) DeleteGroup(id uuid.UUID) error {
	_, err := gs.gr.FindByID(id)
	if err != nil {
		return err
	}
	if err := gs.gr.Delete(id); err != nil {
		return err
	}
	return nil
}

func (gs *groupService) FindByID(id uuid.UUID) (*model.Group, error) {
	return gs.gr.FindByID(id)
}

func (gs *groupService) FindByName(name string) (*model.Group, error) {
	return gs.gr.FindByName(name)
}

func (gs *groupService) GetAllGroups() ([]*model.Group, error) {
	return gs.gr.GetAll()
}
