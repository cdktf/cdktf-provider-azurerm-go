// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package automationaccount

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutomationAccountPrivateEndpointConnectionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AutomationAccountPrivateEndpointConnectionList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AutomationAccountPrivateEndpointConnectionList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AutomationAccountPrivateEndpointConnectionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AutomationAccountPrivateEndpointConnectionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AutomationAccountPrivateEndpointConnectionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAutomationAccountPrivateEndpointConnectionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

