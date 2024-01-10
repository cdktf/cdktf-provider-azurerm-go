// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package subnet

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SubnetDelegationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SubnetDelegationList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SubnetDelegationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SubnetDelegationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SubnetDelegationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SubnetDelegationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SubnetDelegationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSubnetDelegationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

