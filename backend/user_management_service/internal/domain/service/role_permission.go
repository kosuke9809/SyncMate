package service

import (
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
)

type IRolePermissionService interface {
	HasPermission(roleID uint, permissionName string) (bool, error)
	AddPermissionToRole(roleID uint, permissionName string) error
	RemovePermissionFromRole(roleID uint, permissionName string) error
}

type rolePermissionService struct {
	rpr repository.IRolePermissionRepository
	rr  repository.IRoleRepository
	pr  repository.IPermissionRepository
}

func NewRolePermissionService(rpr repository.IRolePermissionRepository, rr repository.IRoleRepository, pr repository.IPermissionRepository) IRolePermissionService {
	return &rolePermissionService{
		rpr: rpr,
		rr:  rr,
		pr:  pr,
	}
}

func (rps *rolePermissionService) HasPermission(roleID uint, permissionName string) (bool, error) {
	permission, err := rps.pr.FindByName(permissionName)
	if err != nil {
		return false, err
	}
	return rps.rpr.HasPermission(roleID, permission.ID)
}

func (rps *rolePermissionService) AddPermissionToRole(roleID uint, permissionName string) error {
	permission, err := rps.rr.FindByName(permissionName)
	if err != nil {
		return err
	}
	return rps.rpr.AddPermissionToRole(roleID, permission.ID)
}

func (rps *rolePermissionService) RemovePermissionFromRole(roleID uint, permissionName string) error {
	permission, err := rps.rr.FindByName(permissionName)
	if err != nil {
		return err
	}
	return rps.rpr.RemovePermissionFromRole(roleID, permission.ID)
}
