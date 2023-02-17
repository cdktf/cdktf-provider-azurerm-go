//go:build no_runtime_type_checking

package iothub

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IothubRouteList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IothubRouteList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IothubRouteList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IothubRouteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IothubRouteList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IothubRouteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIothubRouteListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

