package persistence

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
	"gorm.io/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

type userGroupPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.IUserRepository {
	return &userPersistence{db}
}

func NewUserGroupPersistence(db *gorm.DB) repository.IUserGroupRepository {
	return &userGroupPersistence{db}
}

func (up *userPersistence) Create(user *model.User) error {
	if err := up.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (up *userPersistence) Update(user *model.User) error {
	if err := up.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (up *userPersistence) Delete(id uuid.UUID) error {
	if err := up.db.Delete(&model.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (up *userPersistence) FindByEmail(email string) (*model.User, error) {
	user := model.User{}
	if err := up.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (up *userPersistence) FindById(id uuid.UUID) (*model.User, error) {
	user := model.User{}
	if err := up.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (up *userPersistence) FindProfileById(userId uuid.UUID) (*model.Profile, error) {
	profile := model.Profile{}
	if err := up.db.Where("user_id = ?", userId).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (up *userPersistence) GetAll() ([]*model.User, error) {
	users := []*model.User{}
	if err := up.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (up *userPersistence) FindByPasswordResetToken(token string) (*model.User, error) {
	user := model.User{}
	if err := up.db.Where("password_reset_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ugp *userGroupPersistence) AddUserToGroup(userId uuid.UUID, groupId uuid.UUID, roleId uint) error {
	userGroup := model.UserGroup{
		UserID:  userId,
		GroupID: groupId,
		RoleID:  roleId,
	}
	if err := ugp.db.Create(&userGroup).Error; err != nil {
		return err
	}
	return nil
}

func (ugp *userGroupPersistence) RemoveUserFromGroup(userId uuid.UUID, groupId uuid.UUID) error {
	if err := ugp.db.Delete(&model.UserGroup{}, "user_id = ? AND group_id = ?", userId, groupId).Error; err != nil {
		return err
	}
	return nil
}

func (ugp *userGroupPersistence) GetGroupsByUserId(userId uuid.UUID) ([]*model.Group, error) {
	groups := []*model.Group{}
	if err := ugp.db.Table("user_groups").Select("groups.*").Joins("JOIN groups ON user_groups.group_id = groups.id").Where("user_id = ?", userId).Scan(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (ugp *userGroupPersistence) FindUserGroupRole(userId uuid.UUID, groupId uuid.UUID) (*model.Role, error) {
	role := model.Role{}
	if err := ugp.db.Table(("user_groups")).Select("roles.*").Joins("JOIN roles ON user_groups.role_id = roles.id").Where("user_id = ? AND group_id = ?", userId, groupId).Scan(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
