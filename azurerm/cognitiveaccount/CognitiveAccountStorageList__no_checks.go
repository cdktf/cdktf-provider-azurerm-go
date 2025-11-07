// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cognitiveaccount

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CognitiveAccountStorageList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CognitiveAccountStorageList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CognitiveAccountStorageList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CognitiveAccountStorageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CognitiveAccountStorageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CognitiveAccountStorageList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CognitiveAccountStorageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCognitiveAccountStorageListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

