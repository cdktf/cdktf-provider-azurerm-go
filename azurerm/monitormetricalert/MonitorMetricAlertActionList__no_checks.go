// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package monitormetricalert

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitorMetricAlertActionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MonitorMetricAlertActionList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitorMetricAlertActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitorMetricAlertActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

