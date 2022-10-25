//go:build no_runtime_type_checking

package purviewaccount

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PurviewAccountManagedResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PurviewAccountManagedResourcesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PurviewAccountManagedResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PurviewAccountManagedResourcesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PurviewAccountManagedResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPurviewAccountManagedResourcesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

