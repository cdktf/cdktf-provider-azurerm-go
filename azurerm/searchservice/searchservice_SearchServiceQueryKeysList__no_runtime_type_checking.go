//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package searchservice

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SearchServiceQueryKeysList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SearchServiceQueryKeysList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SearchServiceQueryKeysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SearchServiceQueryKeysList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SearchServiceQueryKeysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSearchServiceQueryKeysListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

