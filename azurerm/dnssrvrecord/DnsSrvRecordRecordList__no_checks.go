// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dnssrvrecord

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsSrvRecordRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DnsSrvRecordRecordList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DnsSrvRecordRecordList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsSrvRecordRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsSrvRecordRecordList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DnsSrvRecordRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDnsSrvRecordRecordListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

