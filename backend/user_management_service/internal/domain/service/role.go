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
	GetAll() ([]*model.Role, error)
}

type roleService struct {
	rp repository.IRoleRepository
}

func NewRoleService(rp repository.IRoleRepository) IRoleService {
	return &roleService{rp}
}

func (rs *roleService) CreateRole(roleName string) (*model.Role, error) {
	if _, err := rs.rp.FindByName(roleName); err == nil {
		return nil, errors.New("role already exists")
	}
	role := &model.Role{
		RoleName: roleName,
	}
	newRole, err := rs.rp.Create(role)
	if err != nil {
		return nil, err
	}
	return newRole, nil
}

func (rs *roleService) UpdateRole(roleID uint, roleName string) (*model.Role, error) {
	role, err := rs.rp.FindByID(roleID)
	if err != nil {
		return nil, errors.New("role not found")
	}
	role.RoleName = roleName
	return rs.rp.Update(role)
}

func (rs *roleService) DeleteRole(roleID uint) error {
	if _, err := rs.rp.FindByID(roleID); err != nil {
		return errors.New("role not found")
	}
	if err := rs.rp.Delete(roleID); err != nil {
		return err
	}
	return nil
}

func (rs *roleService) FindByID(roleID uint) (*model.Role, error) {
	return rs.rp.FindByID(roleID)
}

func (rs *roleService) FindByName(roleName string) (*model.Role, error) {
	return rs.rp.FindByName(roleName)
}

func (rs *roleService) GetAll() ([]*model.Role, error) {
	return rs.rp.GetAll()
}
