// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package redisenterprisedatabase

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedisEnterpriseDatabaseModuleList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedisEnterpriseDatabaseModuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedisEnterpriseDatabaseModuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RedisEnterpriseDatabaseModuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedisEnterpriseDatabaseModuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedisEnterpriseDatabaseModuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedisEnterpriseDatabaseModuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

