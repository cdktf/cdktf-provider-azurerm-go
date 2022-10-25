//go:build no_runtime_type_checking

package networksecuritygroup

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkSecurityGroupSecurityRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkSecurityGroupSecurityRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkSecurityGroupSecurityRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

