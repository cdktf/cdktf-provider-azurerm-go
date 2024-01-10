// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apimanagement

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiManagementAdditionalLocationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApiManagementAdditionalLocationList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApiManagementAdditionalLocationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApiManagementAdditionalLocationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApiManagementAdditionalLocationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

