package persistence

import (
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/domain/model"
	"github.com/kosuke9809/SyncMate/internal/domain/repository"
	"gorm.io/gorm"
)

type userGroupPersistence struct {
	db *gorm.DB
}

func NewUserGroupPersistence(db *gorm.DB) repository.IUserGroupRepository {
	return &userGroupPersistence{db}
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

func (ugp *userGroupPersistence) GetMembersByGroupId(groupId uuid.UUID) ([]*model.User, error) {
	return nil, nil
}

func (ugp *userGroupPersistence) FindGroupOwner(groupId uuid.UUID) (*model.User, error) {
	return nil, nil
}
