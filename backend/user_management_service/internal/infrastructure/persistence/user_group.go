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

func (ugp *userGroupPersistence) AddUserToGroup(userID uuid.UUID, groupID uuid.UUID, roleID uint) error {
	userGroup := model.UserGroup{
		UserID:  userID,
		GroupID: groupID,
		RoleID:  roleID,
	}
	if err := ugp.db.Create(&userGroup).Error; err != nil {
		return err
	}
	return nil
}

func (ugp *userGroupPersistence) RemoveUserFromGroup(userID uuid.UUID, groupID uuid.UUID) error {
	if err := ugp.db.Delete(&model.UserGroup{}, "user_id = ? AND group_id = ?", userID, groupID).Error; err != nil {
		return err
	}
	return nil
}

func (ugp *userGroupPersistence) GetGroupsByUserID(userID uuid.UUID) ([]*model.Group, error) {
	groups := []*model.Group{}
	if err := ugp.db.Table("user_groups").Select("groups.*").Joins("JOIN groups ON user_groups.group_id = groups.id").Where("user_id = ?", userID).Scan(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (ugp *userGroupPersistence) FindUserRoleInGroup(userID uuid.UUID, groupID uuid.UUID) (*model.Role, error) {
	role := model.Role{}
	if err := ugp.db.Table("user_groups").Select("roles.*").Joins("JOIN roles ON user_groups.role_id = roles.id").Where("user_id = ? AND group_id = ?", userID, groupID).Scan(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (ugp *userGroupPersistence) GetMembersByGroupID(groupID uuid.UUID) ([]*model.User, error) {
	users := []*model.User{}
	if err := ugp.db.Table("user_groups").Select("users.*").Joins("JOIN users ON user_groups.user_id = users.id").Where("group_id = ?", groupID).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ugp *userGroupPersistence) FindGroupOwner(groupID uuid.UUID) (*model.User, error) {
	user := model.User{}
	if err := ugp.db.Table("user_groups").Select("users.*").Joins("JOIN users ON user_groups.user_id = users.id").Joins("JOIN roles ON user_groups.role_id = roles.id").Where("group_id = ? AND roles.name = ?", groupID, "Owner").Scan(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (usp *userGroupPersistence) IsUserInGroup(userID, groupID uuid.UUID) (bool, error) {
	var count int64
	if err := usp.db.Model(&model.UserGroup{}).Where("user_id = ? AND group_id = ?", userID, groupID).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
