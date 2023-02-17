//go:build no_runtime_type_checking

package virtualhubroutetable

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualHubRouteTableRouteList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualHubRouteTableRouteList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteTableRouteList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteTableRouteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteTableRouteList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteTableRouteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualHubRouteTableRouteListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

