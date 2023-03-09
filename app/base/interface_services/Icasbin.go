package interface_services

import (
	"SH-admin/app/models"
)

type ICasbinService interface {
	GetAccessApiByRoleId(roleId string) [][]string
	UpdateCasbin(updateParam *models.UpdateCasbinParam) error
	UpdateUserRole(oldRules []string, newRoles []string) error
}
