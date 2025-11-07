// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package iothubdps

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IothubDpsIpFilterRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IothubDpsIpFilterRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IothubDpsIpFilterRuleList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IothubDpsIpFilterRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IothubDpsIpFilterRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IothubDpsIpFilterRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IothubDpsIpFilterRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIothubDpsIpFilterRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

