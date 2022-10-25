//go:build no_runtime_type_checking

package applicationgateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationGatewayUrlPathMapList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationGatewayUrlPathMapList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayUrlPathMapList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationGatewayUrlPathMapListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

