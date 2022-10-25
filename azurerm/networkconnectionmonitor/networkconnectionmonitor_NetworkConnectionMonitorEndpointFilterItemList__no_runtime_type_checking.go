//go:build no_runtime_type_checking

package networkconnectionmonitor

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkConnectionMonitorEndpointFilterItemList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkConnectionMonitorEndpointFilterItemList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointFilterItemList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointFilterItemList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointFilterItemList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkConnectionMonitorEndpointFilterItemList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkConnectionMonitorEndpointFilterItemListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

