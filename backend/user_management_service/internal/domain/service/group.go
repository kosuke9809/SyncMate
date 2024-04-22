package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
)

type IGroupService interface {
	CreateGroup(name, description string) (*model.Group, error)
	UpdateGroup(id, name, description string) (*model.Group, error)
	DeleteGroup(id string) error
}

type groupService struct {
	gr repository.IGroupRepository
}

func NewGroupService(gr repository.IGroupRepository) *groupService {
	return &groupService{gr}
}

func (gs *groupService) CreateGroup(groupName, description string) (*model.Group, error) {
	group := &model.Group{
		ID:          uuid.New(),
		GroupName:   groupName,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	newGroup, err := gs.gr.Create(group)
	if err != nil {
		return nil, err
	}
	return newGroup, nil
}
