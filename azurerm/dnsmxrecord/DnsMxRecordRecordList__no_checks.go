// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dnsmxrecord

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsMxRecordRecordList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DnsMxRecordRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DnsMxRecordRecordList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DnsMxRecordRecordList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsMxRecordRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsMxRecordRecordList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DnsMxRecordRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDnsMxRecordRecordListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

