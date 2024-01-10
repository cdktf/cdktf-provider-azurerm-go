// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sqlfailovergroup

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlFailoverGroupPartnerServersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SqlFailoverGroupPartnerServersList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlFailoverGroupPartnerServersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlFailoverGroupPartnerServersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SqlFailoverGroupPartnerServersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlFailoverGroupPartnerServersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlFailoverGroupPartnerServersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlFailoverGroupPartnerServersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

