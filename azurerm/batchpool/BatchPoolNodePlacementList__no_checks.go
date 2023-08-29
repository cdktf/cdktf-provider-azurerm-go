// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package batchpool

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchPoolNodePlacementList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchPoolNodePlacementList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchPoolNodePlacementList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchPoolNodePlacementList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchPoolNodePlacementList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchPoolNodePlacementList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchPoolNodePlacementListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

