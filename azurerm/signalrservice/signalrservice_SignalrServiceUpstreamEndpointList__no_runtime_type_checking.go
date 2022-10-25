//go:build no_runtime_type_checking

package signalrservice

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SignalrServiceUpstreamEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SignalrServiceUpstreamEndpointList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceUpstreamEndpointList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceUpstreamEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceUpstreamEndpointList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceUpstreamEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSignalrServiceUpstreamEndpointListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

