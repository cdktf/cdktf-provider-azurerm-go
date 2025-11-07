// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package netappvolume

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetappVolumeExportPolicyRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetappVolumeExportPolicyRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

