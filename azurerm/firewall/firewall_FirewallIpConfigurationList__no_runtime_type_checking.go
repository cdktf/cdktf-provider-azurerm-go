//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package firewall

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirewallIpConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FirewallIpConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FirewallIpConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FirewallIpConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FirewallIpConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FirewallIpConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFirewallIpConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

