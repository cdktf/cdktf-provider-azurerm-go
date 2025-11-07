// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storageobjectreplication

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageObjectReplicationRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StorageObjectReplicationRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageObjectReplicationRulesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageObjectReplicationRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageObjectReplicationRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageObjectReplicationRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageObjectReplicationRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageObjectReplicationRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

