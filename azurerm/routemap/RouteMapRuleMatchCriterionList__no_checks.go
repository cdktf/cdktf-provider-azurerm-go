// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package routemap

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouteMapRuleMatchCriterionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRouteMapRuleMatchCriterionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

