// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package resourcedeploymentscriptazurecli

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourceDeploymentScriptAzureCliEnvironmentVariableList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ResourceDeploymentScriptAzureCliEnvironmentVariableList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCliEnvironmentVariableList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCliEnvironmentVariableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCliEnvironmentVariableList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ResourceDeploymentScriptAzureCliEnvironmentVariableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewResourceDeploymentScriptAzureCliEnvironmentVariableListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

