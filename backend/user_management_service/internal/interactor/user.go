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
	NewUserGroupRepository() repository.IUserGroupRepository
}

type userInteractor struct {
	Conn *gorm.DB
}

func NewInteractor(conn *gorm.DB) IInteractor {
	return &userInteractor{Conn: conn}
}

func (i *userInteractor) NewUserPersistence() repository.IUserRepository {
	return persistence.NewUserPersistence(i.Conn)
}

func (i *userInteractor) NewUserService() service.IUserService {
	return service.NewUserService(i.NewUserPersistence())
}

func (i *userInteractor) NewUserUsecase() usecase.IUserUsecase {
	return usecase.NewUserUsecase(i.NewUserService())
}

func (i *userInteractor) NewUserHandler() handler.IUserHandler {
	return handler.NewUserHandler(i.NewUserUsecase())
}

func (i *userInteractor) NewUserGroupRepository() repository.IUserGroupRepository {
	return persistence.NewUserGroupPersistence(i.Conn)
}
