//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dataazurermdatafactory

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermDataFactoryVstsConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermDataFactoryVstsConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermDataFactoryVstsConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermDataFactoryVstsConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermDataFactoryVstsConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermDataFactoryVstsConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

