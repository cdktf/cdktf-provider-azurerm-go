// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lb

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Lb) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_Lb) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateImportFromParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (l *jsiiProxy_Lb) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (l *jsiiProxy_Lb) validatePutFrontendIpConfigurationParameters(value interface{}) error {
	return nil
}

func (l *jsiiProxy_Lb) validatePutTimeoutsParameters(value *LbTimeouts) error {
	return nil
}

func validateLb_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateLb_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLb_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateLb_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetEdgeZoneParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetResourceGroupNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetSkuParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetSkuTierParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Lb) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func validateNewLbParameters(scope constructs.Construct, id *string, config *LbConfig) error {
	return nil
}

