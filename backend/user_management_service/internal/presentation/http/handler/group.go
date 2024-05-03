package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kosuke9809/SyncMate/internal/usecase"
	"github.com/kosuke9809/SyncMate/internal/utils"
	"github.com/labstack/echo/v4"
)

type IGroupHandler interface {
	CreateNewGroup(ctx echo.Context) error
	InviteUserToGroup(ctx echo.Context) error
	AcceptInvitation(ctx echo.Context) error
	RejectInvitation(ctx echo.Context) error
	CancelInvitation(ctx echo.Context) error
	RemoveUserFromGroup(ctx echo.Context) error
	GetGroupDetails(ctx echo.Context) error
	GetGroupMembers(ctx echo.Context) error
	DeleteGroup(ctx echo.Context) error
}

type groupHandler struct {
	gu usecase.IGroupUsecase
}

func NewGroupHandler(gu usecase.IGroupUsecase) IGroupHandler {
	return &groupHandler{gu}
}

// CreateNewGroup godoc
// @Summary Create a new group
// @Description Create a new group with the provided details
// @Tags Group
// @Accept json
// @Produce json
// @Success 200 {object} model.Group
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /group/create [post]
func (gh *groupHandler) CreateNewGroup(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.Claims)
	creatorID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse user ID")
	}
	var req struct {
		GroupName   string `json:"group_name"`
		Description string `json:"description"`
	}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}
	group, err := gh.gu.CreateNewGroup(req.GroupName, req.Description, creatorID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create new group")
	}
	return ctx.JSON(http.StatusOK, group)
}

// InviteUserToGroup godoc
// @Summary Invite a user to a group
// @Description Invite a user to a group by email
// @Tags Group
// @Accept json
// @Produce json
// @Success 200 {object} model.Invitation
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /group/invite [post]
func (gh *groupHandler) InviteUserToGroup(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.Claims)
	inviterID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse user ID")
	}
	var req struct {
		GroupID      uuid.UUID `json:"group_id"`
		InviteeEmail string    `json:"invitee_email"`
	}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}
	invitation, err := gh.gu.InviteUserToGroup(inviterID, req.GroupID, req.InviteeEmail)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to invite user to group")
	}
	return ctx.JSON(http.StatusOK, invitation)
}

// AcceptInvitation godoc
// @Summary Accept a group invitation
// @Description Accept a pending group invitation
// @Tags Group
// @Accept json
// @Produce json
// @Param invitationID path string true "Invitation ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /group/invitation/accept [post]
func (gh *groupHandler) AcceptInvitation(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.Claims)
	inviteeID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse user ID")
	}
	var req struct {
		InvitationID uuid.UUID `json:"invitation_id"`
	}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}
	err = gh.gu.AcceptInvitation(inviteeID, req.InvitationID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to accept invitation")
	}
	return ctx.NoContent(http.StatusOK)
}

// RejectInvitation godoc
// @Summary Reject a group invitation
// @Description Reject a pending group invitation
// @Tags Group
// @Accept json
// @Produce json
// @Param invitationID path string true "Invitation ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /group/invitation/reject [post]
func (gh *groupHandler) RejectInvitation(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.Claims)
	inviteeID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse user ID")
	}
	var req struct {
		InvitationID uuid.UUID `json:"invitation_id"`
	}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}
	err = gh.gu.RejectInvitation(inviteeID, req.InvitationID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to reject invitation")
	}
	return ctx.NoContent(http.StatusOK)
}

// CancelInvitation godoc
// @Summary Cancel a group invitation
// @Description Cancel a pending group invitation
// @Tags Group
// @Accept json
// @Produce json
// @Param invitationID path string true "Invitation ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /grous/invitation/cancel [post]
func (gh *groupHandler) CancelInvitation(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.Claims)
	inviterID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse user ID")
	}
	var req struct {
		InvitationID uuid.UUID `json:"invitation_id"`
	}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}
	err = gh.gu.CancelInvitation(inviterID, req.InvitationID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to cancel invitation")
	}
	return ctx.NoContent(http.StatusOK)
}

// RemoveUserFromGroup godoc
// @Summary Remove a user from a group
// @Description Remove a user from a group
// @Tags Group
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /group/member/remove [post]
func (gh *groupHandler) RemoveUserFromGroup(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.Claims)
	inviterID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse user ID")
	}
	var req struct {
		RemoveUserID uuid.UUID `json:"remove_user_id"`
		GroupID      uuid.UUID `json:"group_id"`
	}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}
	err = gh.gu.RemoveUserFromGroup(inviterID, req.RemoveUserID, req.GroupID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to remove user from group")
	}
	return ctx.NoContent(http.StatusOK)
}

// DeleteGroup godoc
// @Summary Delete a group
// @Description Delete a specific group
// @Tags Group
// @Accept json
// @Produce json
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /group/delete [post]
func (gh *groupHandler) DeleteGroup(ctx echo.Context) error {
	var req struct {
		GroupID uuid.UUID `json:"group_id"`
	}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}
	err := gh.gu.DeleteGroup(req.GroupID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete group")
	}
	return ctx.NoContent(http.StatusOK)
}

// GetGroupDetails godoc
// @Summary Get group details
// @Description Get the details of a specific group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} model.Group
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /group/{id}/details [get]
func (gh *groupHandler) GetGroupDetails(ctx echo.Context) error {
	groupID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid group ID")
	}
	group, err := gh.gu.GetGroupDetails(groupID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get group details")
	}
	return ctx.JSON(http.StatusOK, group)
}

// GetGroupMembers godoc
// @Summary Get group members
// @Description Get the members of a specific group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {array} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /group/{id}/members [get]
func (gh *groupHandler) GetGroupMembers(ctx echo.Context) error {
	groupID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid group ID")
	}
	members, err := gh.gu.GetGroupMembers(groupID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get group members")
	}
	return ctx.JSON(http.StatusOK, members)
}
