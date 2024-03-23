// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package nginxdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NginxDeploymentConfigurationConfigFileList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NginxDeploymentConfigurationConfigFileList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NginxDeploymentConfigurationConfigFileList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NginxDeploymentConfigurationConfigFileList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NginxDeploymentConfigurationConfigFileList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NginxDeploymentConfigurationConfigFileList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NginxDeploymentConfigurationConfigFileList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNginxDeploymentConfigurationConfigFileListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

