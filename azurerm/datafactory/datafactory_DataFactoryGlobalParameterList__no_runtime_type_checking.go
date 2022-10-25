//go:build no_runtime_type_checking

package datafactory

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataFactoryGlobalParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataFactoryGlobalParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataFactoryGlobalParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataFactoryGlobalParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataFactoryGlobalParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataFactoryGlobalParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataFactoryGlobalParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

