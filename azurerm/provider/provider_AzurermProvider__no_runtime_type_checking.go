//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AzurermProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AzurermProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAzurermProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AzurermProvider) validateSetDisableCorrelationRequestIdParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzurermProvider) validateSetDisableTerraformPartnerIdParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzurermProvider) validateSetFeaturesParameters(val *AzurermProviderFeatures) error {
	return nil
}

func (j *jsiiProxy_AzurermProvider) validateSetSkipProviderRegistrationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzurermProvider) validateSetStorageUseAzureadParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzurermProvider) validateSetUseMsiParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzurermProvider) validateSetUseOidcParameters(val interface{}) error {
	return nil
}

func validateNewAzurermProviderParameters(scope constructs.Construct, id *string, config *AzurermProviderConfig) error {
	return nil
}

