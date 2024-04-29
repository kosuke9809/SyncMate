package service

import "github.com/kosuke9809/SyncMate/internal/domain/repository"

type IRolePermissionService interface {
}

type rolePermissionService struct {
	rpr repository.IRolePermissionRepository
}

func NewRolePermissionService(rpr repository.IRolePermissionRepository) IRolePermissionService {
	return &rolePermissionService{rpr}
}
