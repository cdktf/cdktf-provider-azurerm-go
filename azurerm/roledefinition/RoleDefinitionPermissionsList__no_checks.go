// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package roledefinition

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RoleDefinitionPermissionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RoleDefinitionPermissionsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RoleDefinitionPermissionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RoleDefinitionPermissionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RoleDefinitionPermissionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RoleDefinitionPermissionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RoleDefinitionPermissionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoleDefinitionPermissionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

