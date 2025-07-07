package convertor

import (
	"auth/biz/model/consts"
	"auth/biz/model/dto"
)

var RoleStatus2DO = map[dto.RoleStatus]string{
	dto.RoleStatusValid:   consts.ResourceStatusValid,
	dto.RoleStatusInvalid: consts.ResourceStatusInvalid,
}

var RoleStatus2DTO = map[string]dto.RoleStatus{
	consts.ResourceStatusValid:   dto.RoleStatusValid,
	consts.ResourceStatusInvalid: dto.RoleStatusInvalid,
}

var ResourceStatus2DO = map[dto.ResourceStatus]string{
	dto.ResourceStatusValid:   consts.ResourceStatusValid,
	dto.ResourceStatusInvalid: consts.ResourceStatusInvalid,
}

var ResourceStatus2DTO = map[string]dto.ResourceStatus{
	consts.ResourceStatusValid:   dto.ResourceStatusValid,
	consts.ResourceStatusInvalid: dto.ResourceStatusInvalid,
}

var PermissionStatus2DO = map[dto.PermissionStatus]string{
	dto.PermissionStatusValid:   consts.PermissionStatusValid,
	dto.PermissionStatusInvalid: consts.PermissionStatusInvalid,
}

var PermissionStatus2DTO = map[string]dto.PermissionStatus{
	consts.PermissionStatusValid:   dto.PermissionStatusValid,
	consts.PermissionStatusInvalid: dto.PermissionStatusInvalid,
}

var PermissionEffect2DO = map[dto.PermissionEffect]string{
	dto.PermissionEffectAllow: consts.PermissionEffectAllow,
	dto.PermissionEffectDeny:  consts.PermissionEffectDeny,
}

var PermissionEffect2DTO = map[string]dto.PermissionEffect{
	consts.PermissionEffectAllow: dto.PermissionEffectAllow,
	consts.PermissionEffectDeny:  dto.PermissionEffectDeny,
}
