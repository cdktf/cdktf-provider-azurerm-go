//go:build no_runtime_type_checking

package privateendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateEndpointIpConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrivateEndpointIpConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointIpConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointIpConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointIpConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointIpConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrivateEndpointIpConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
