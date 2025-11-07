// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package aiservices

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AiServicesStorageList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AiServicesStorageList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AiServicesStorageList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AiServicesStorageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AiServicesStorageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AiServicesStorageList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AiServicesStorageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAiServicesStorageListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

