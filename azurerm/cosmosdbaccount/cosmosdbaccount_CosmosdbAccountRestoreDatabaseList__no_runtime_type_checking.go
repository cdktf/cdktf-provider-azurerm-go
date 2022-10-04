//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cosmosdbaccount

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCosmosdbAccountRestoreDatabaseListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

