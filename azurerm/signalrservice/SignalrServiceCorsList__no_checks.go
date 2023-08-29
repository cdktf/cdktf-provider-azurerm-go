// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package signalrservice

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SignalrServiceCorsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SignalrServiceCorsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceCorsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceCorsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceCorsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceCorsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSignalrServiceCorsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

