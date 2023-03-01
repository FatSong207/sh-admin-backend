package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IApiService interface {
	Core.IService[models.Api, models.ApiOutDto]
	GetAllApiTree() (result []models.ApiForTree, err error)
}
