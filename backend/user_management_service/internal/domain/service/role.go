package service

import (
	"errors"

	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
)

type IRoleService interface {
	CreateRole(roleName string) (*model.Role, error)
	UpdateRole(roleID uint, roleName string) (*model.Role, error)
	DeleteRole(roleID uint) error
	FindByID(roleID uint) (*model.Role, error)
	FindByName(roleName string) (*model.Role, error)
	GetAllRoles() ([]*model.Role, error)
}

type roleService struct {
	rr repository.IRoleRepository
}

func NewRoleService(rr repository.IRoleRepository) IRoleService {
	return &roleService{rr}
}

func (rs *roleService) CreateRole(roleName string) (*model.Role, error) {
	if _, err := rs.rr.FindByName(roleName); err == nil {
		return nil, errors.New("role already exists")
	}
	role := &model.Role{
		RoleName: roleName,
	}
	newRole, err := rs.rr.Create(role)
	if err != nil {
		return nil, err
	}
	return newRole, nil
}

func (rs *roleService) UpdateRole(roleID uint, roleName string) (*model.Role, error) {
	role, err := rs.rr.FindByID(roleID)
	if err != nil {
		return nil, errors.New("role not found")
	}
	role.RoleName = roleName
	return rs.rr.Update(role)
}

func (rs *roleService) DeleteRole(roleID uint) error {
	if _, err := rs.rr.FindByID(roleID); err != nil {
		return errors.New("role not found")
	}
	if err := rs.rr.Delete(roleID); err != nil {
		return err
	}
	return nil
}

func (rs *roleService) FindByID(roleID uint) (*model.Role, error) {
	return rs.rr.FindByID(roleID)
}

func (rs *roleService) FindByName(roleName string) (*model.Role, error) {
	return rs.rr.FindByName(roleName)
}

func (rs *roleService) GetAllRoles() ([]*model.Role, error) {
	return rs.rr.GetAll()
}
