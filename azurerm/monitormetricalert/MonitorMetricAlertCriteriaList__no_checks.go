// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package monitormetricalert

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitorMetricAlertCriteriaList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitorMetricAlertCriteriaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertCriteriaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertCriteriaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertCriteriaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertCriteriaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitorMetricAlertCriteriaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

