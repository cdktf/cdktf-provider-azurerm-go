// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkpacketcapture

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkPacketCaptureFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkPacketCaptureFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkPacketCaptureFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkPacketCaptureFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkPacketCaptureFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkPacketCaptureFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkPacketCaptureFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

