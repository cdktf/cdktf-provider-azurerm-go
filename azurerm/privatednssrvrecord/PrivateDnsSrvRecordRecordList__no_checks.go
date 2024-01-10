// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package privatednssrvrecord

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateDnsSrvRecordRecordList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrivateDnsSrvRecordRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrivateDnsSrvRecordRecordList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsSrvRecordRecordList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsSrvRecordRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsSrvRecordRecordList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsSrvRecordRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrivateDnsSrvRecordRecordListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

