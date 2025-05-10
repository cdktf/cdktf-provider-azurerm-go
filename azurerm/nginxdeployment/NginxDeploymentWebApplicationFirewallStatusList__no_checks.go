// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package nginxdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NginxDeploymentWebApplicationFirewallStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNginxDeploymentWebApplicationFirewallStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

