// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dnstxtrecord

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsTxtRecordRecordList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DnsTxtRecordRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DnsTxtRecordRecordList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DnsTxtRecordRecordList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsTxtRecordRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsTxtRecordRecordList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DnsTxtRecordRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDnsTxtRecordRecordListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

