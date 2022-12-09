//go:build no_runtime_type_checking

package networkmanager

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkManagerCrossTenantScopesList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkManagerCrossTenantScopesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkManagerCrossTenantScopesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkManagerCrossTenantScopesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkManagerCrossTenantScopesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkManagerCrossTenantScopesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

