package service

import (
	"errors"

	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
)

type IPermissionService interface {
	CreatePermission(permissionName string) (*model.Permission, error)
	UpdatePermission(permissionID uint, permissionName string) (*model.Permission, error)
	DeletePermission(permissionID uint) error
	FindByID(permissionID uint) (*model.Permission, error)
	FindByName(permissionName string) (*model.Permission, error)
	GetAllPermissions() ([]*model.Permission, error)
}

type permissionService struct {
	pr repository.IPermissionRepository
}

func NewPermissionService(pr repository.IPermissionRepository) IPermissionService {
	return &permissionService{pr}
}

func (ps *permissionService) CreatePermission(permissionName string) (*model.Permission, error) {
	if _, err := ps.pr.FindByName(permissionName); err == nil {
		return nil, errors.New("permission already exists")
	}
	permission := &model.Permission{
		PermissionName: permissionName,
	}
	newPermission, err := ps.pr.Create(permission)
	if err != nil {
		return nil, err
	}
	return newPermission, nil
}

func (ps *permissionService) UpdatePermission(permissionID uint, permissionName string) (*model.Permission, error) {
	permission, err := ps.pr.FindByID(permissionID)
	if err != nil {
		return nil, errors.New("permission not found")
	}
	permission.PermissionName = permissionName
	updatedPermission, err := ps.pr.Update(permission)
	if err != nil {
		return nil, err
	}
	return updatedPermission, nil
}

func (ps *permissionService) DeletePermission(permissionID uint) error {
	if _, err := ps.pr.FindByID(permissionID); err != nil {
		return errors.New("permission not found")
	}
	if err := ps.pr.Delete(permissionID); err != nil {
		return err
	}
	return nil
}

func (ps *permissionService) FindByID(permissionID uint) (*model.Permission, error) {
	return ps.pr.FindByID(permissionID)
}

func (ps *permissionService) FindByName(permissionName string) (*model.Permission, error) {
	return ps.pr.FindByName(permissionName)
}

func (ps *permissionService) GetAllPermissions() ([]*model.Permission, error) {
	return ps.pr.GetAll()
}
