// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package vpnsite

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpnSiteLinkList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VpnSiteLinkList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VpnSiteLinkList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VpnSiteLinkList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VpnSiteLinkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VpnSiteLinkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VpnSiteLinkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVpnSiteLinkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

