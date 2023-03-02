package IServices

import "SH-admin/models"

type ICasbinService interface {
	GetAccessApiByRoleId(roleId string) [][]string
	UpdateCasbin(updateParam *models.UpdateCasbinParam) error
	UpdateUserRole(oldRules []string, newRoles []string) error
}