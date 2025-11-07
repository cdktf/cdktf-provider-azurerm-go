// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package privatednstxtrecord

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateDnsTxtRecordRecordList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrivateDnsTxtRecordRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrivateDnsTxtRecordRecordList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsTxtRecordRecordList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsTxtRecordRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsTxtRecordRecordList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsTxtRecordRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrivateDnsTxtRecordRecordListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

