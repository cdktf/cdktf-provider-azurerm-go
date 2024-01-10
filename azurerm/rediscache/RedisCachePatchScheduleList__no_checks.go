// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rediscache

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedisCachePatchScheduleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RedisCachePatchScheduleList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedisCachePatchScheduleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedisCachePatchScheduleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RedisCachePatchScheduleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedisCachePatchScheduleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedisCachePatchScheduleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedisCachePatchScheduleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

