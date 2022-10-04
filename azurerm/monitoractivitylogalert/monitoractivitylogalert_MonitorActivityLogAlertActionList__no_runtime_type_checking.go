//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package monitoractivitylogalert

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitorActivityLogAlertActionList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitorActivityLogAlertActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitorActivityLogAlertActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

