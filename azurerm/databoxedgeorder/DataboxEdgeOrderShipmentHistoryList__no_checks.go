// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databoxedgeorder

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataboxEdgeOrderShipmentHistoryList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataboxEdgeOrderShipmentHistoryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataboxEdgeOrderShipmentHistoryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataboxEdgeOrderShipmentHistoryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataboxEdgeOrderShipmentHistoryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataboxEdgeOrderShipmentHistoryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

