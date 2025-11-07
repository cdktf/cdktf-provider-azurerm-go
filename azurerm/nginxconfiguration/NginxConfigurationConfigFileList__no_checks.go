// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package nginxconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NginxConfigurationConfigFileList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NginxConfigurationConfigFileList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NginxConfigurationConfigFileList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationConfigFileList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationConfigFileList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationConfigFileList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationConfigFileList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNginxConfigurationConfigFileListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

