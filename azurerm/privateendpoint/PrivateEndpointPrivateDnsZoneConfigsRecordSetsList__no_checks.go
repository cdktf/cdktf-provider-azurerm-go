// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package privateendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateEndpointPrivateDnsZoneConfigsRecordSetsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrivateEndpointPrivateDnsZoneConfigsRecordSetsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrivateEndpointPrivateDnsZoneConfigsRecordSetsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointPrivateDnsZoneConfigsRecordSetsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointPrivateDnsZoneConfigsRecordSetsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointPrivateDnsZoneConfigsRecordSetsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrivateEndpointPrivateDnsZoneConfigsRecordSetsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

