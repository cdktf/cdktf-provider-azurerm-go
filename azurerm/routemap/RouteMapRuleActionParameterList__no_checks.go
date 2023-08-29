// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package routemap

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouteMapRuleActionParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RouteMapRuleActionParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleActionParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleActionParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleActionParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleActionParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRouteMapRuleActionParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

