// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package routefilter

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouteFilterRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RouteFilterRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RouteFilterRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RouteFilterRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RouteFilterRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RouteFilterRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RouteFilterRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRouteFilterRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

