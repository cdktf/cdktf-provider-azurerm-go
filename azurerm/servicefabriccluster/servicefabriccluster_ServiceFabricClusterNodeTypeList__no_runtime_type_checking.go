//go:build no_runtime_type_checking

package servicefabriccluster

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceFabricClusterNodeTypeList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceFabricClusterNodeTypeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

