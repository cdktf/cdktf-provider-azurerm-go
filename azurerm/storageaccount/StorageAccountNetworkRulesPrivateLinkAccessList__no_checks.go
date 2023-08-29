// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storageaccount

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageAccountNetworkRulesPrivateLinkAccessList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageAccountNetworkRulesPrivateLinkAccessListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

