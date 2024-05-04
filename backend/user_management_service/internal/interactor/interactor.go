package interactor

import (
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
	"github.com/kosuke9809/SyncMate/internal/domain/service"
	"github.com/kosuke9809/SyncMate/internal/infrastructure/persistence"
	"github.com/kosuke9809/SyncMate/internal/presentation/http/handler"
	"github.com/kosuke9809/SyncMate/internal/usecase"
	"gorm.io/gorm"
)

type IInteractor interface {
	NewUserPersistence() repository.IUserRepository
	NewUserService() service.IUserService
	NewUserUsecase() usecase.IUserUsecase
	NewUserHandler() handler.IUserHandler

	NewGroupPersistence() repository.IGroupRepository
	NewGroupService() service.IGroupService
	NewGroupUsecase() usecase.IGroupUsecase
	NewGroupHandler() handler.IGroupHandler

	NewUserGroupPersistence() repository.IUserGroupRepository
	NewUserGroupService() service.IUserGroupService

	NewRolePersistence() repository.IRoleRepository
	NewRoleService() service.IRoleService

	NewPermissionPersistence() repository.IPermissionRepository
	NewPermissionService() service.IPermissionService

	NewRolePermissionPersistence() repository.IRolePermissionRepository
	NewRolePermissionService() service.IRolePermissionService

	NewInvitationPersistence() repository.IInvitationRepository
	NewInvitationService() service.IInvitationService
}

type interactor struct {
	db *gorm.DB
}

func NewInteractor(db *gorm.DB) IInteractor {
	return &interactor{db}
}

// user
func (i *interactor) NewUserPersistence() repository.IUserRepository {
	return persistence.NewUserPersistence(i.db)
}

func (i *interactor) NewUserService() service.IUserService {
	return service.NewUserService(i.NewUserPersistence())
}

func (i *interactor) NewUserUsecase() usecase.IUserUsecase {
	return usecase.NewUserUsecase(i.NewUserService())
}

func (i *interactor) NewUserHandler() handler.IUserHandler {
	return handler.NewUserHandler(i.NewUserUsecase())
}

func (i *interactor) NewUserGroupRepository() repository.IUserGroupRepository {
	return persistence.NewUserGroupPersistence(i.db)
}

// group
func (i *interactor) NewGroupPersistence() repository.IGroupRepository {
	return persistence.NewGroupPersistence(i.db)
}

func (i *interactor) NewGroupService() service.IGroupService {
	return service.NewGroupService(i.NewGroupPersistence())
}

func (i *interactor) NewGroupUsecase() usecase.IGroupUsecase {
	return usecase.NewGroupUsecase(
		i.NewGroupService(),
		i.NewUserService(),
		i.NewUserGroupService(),
		i.NewRoleService(),
		i.NewRolePermissionService(),
		i.NewInvitationService(),
	)
}

func (i *interactor) NewGroupHandler() handler.IGroupHandler {
	return handler.NewGroupHandler(i.NewGroupUsecase())
}

// user group
func (i *interactor) NewUserGroupPersistence() repository.IUserGroupRepository {
	return persistence.NewUserGroupPersistence(i.db)
}

func (i *interactor) NewUserGroupService() service.IUserGroupService {
	return service.NewUserGroupService(i.NewUserGroupPersistence())
}

// role
func (i *interactor) NewRolePersistence() repository.IRoleRepository {
	return persistence.NewRolePersistence(i.db)
}

func (i *interactor) NewRoleService() service.IRoleService {
	return service.NewRoleService(i.NewRolePersistence())
}

// permission
func (i *interactor) NewPermissionPersistence() repository.IPermissionRepository {
	return persistence.NewPermissionPersistence(i.db)
}

func (i *interactor) NewPermissionService() service.IPermissionService {
	return service.NewPermissionService(i.NewPermissionPersistence())
}

// role permission
func (i *interactor) NewRolePermissionPersistence() repository.IRolePermissionRepository {
	return persistence.NewRolePermissionPersistence(i.db)
}

func (i *interactor) NewRolePermissionService() service.IRolePermissionService {
	return service.NewRolePermissionService(
		i.NewRolePermissionPersistence(),
		i.NewRolePersistence(),
		i.NewPermissionPersistence(),
	)
}

// invitation
func (i *interactor) NewInvitationPersistence() repository.IInvitationRepository {
	return persistence.NewInvitationPersistence(i.db)
}

func (i *interactor) NewInvitationService() service.IInvitationService {
	return service.NewInvitationService(
		i.NewInvitationPersistence(),
		i.NewUserPersistence(),
		i.NewGroupPersistence(),
	)
}
