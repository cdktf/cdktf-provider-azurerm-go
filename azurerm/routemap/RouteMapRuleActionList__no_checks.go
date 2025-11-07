// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package routemap

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouteMapRuleActionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RouteMapRuleActionList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RouteMapRuleActionList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRouteMapRuleActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

