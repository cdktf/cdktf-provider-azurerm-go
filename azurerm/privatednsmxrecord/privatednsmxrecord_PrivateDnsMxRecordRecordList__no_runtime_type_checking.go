//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package privatednsmxrecord

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateDnsMxRecordRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrivateDnsMxRecordRecordList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsMxRecordRecordList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsMxRecordRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsMxRecordRecordList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsMxRecordRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrivateDnsMxRecordRecordListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

