//go:build no_runtime_type_checking

package applicationgateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsAuthenticationCertificateList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsAuthenticationCertificateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsAuthenticationCertificateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsAuthenticationCertificateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsAuthenticationCertificateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsAuthenticationCertificateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationGatewayBackendHttpSettingsAuthenticationCertificateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

