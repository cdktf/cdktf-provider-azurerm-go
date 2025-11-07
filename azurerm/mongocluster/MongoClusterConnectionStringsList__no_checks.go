// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mongocluster

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MongoClusterConnectionStringsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MongoClusterConnectionStringsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MongoClusterConnectionStringsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MongoClusterConnectionStringsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MongoClusterConnectionStringsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MongoClusterConnectionStringsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMongoClusterConnectionStringsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

