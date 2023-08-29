// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package nginxdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NginxDeploymentFrontendPrivateList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NginxDeploymentFrontendPrivateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NginxDeploymentFrontendPrivateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NginxDeploymentFrontendPrivateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NginxDeploymentFrontendPrivateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NginxDeploymentFrontendPrivateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNginxDeploymentFrontendPrivateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

