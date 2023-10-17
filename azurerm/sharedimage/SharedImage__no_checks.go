// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sharedimage

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SharedImage) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validatePutIdentifierParameters(value *SharedImageIdentifier) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validatePutPurchasePlanParameters(value *SharedImagePurchasePlan) error {
	return nil
}

func (s *jsiiProxy_SharedImage) validatePutTimeoutsParameters(value *SharedImageTimeouts) error {
	return nil
}

func validateSharedImage_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSharedImage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSharedImage_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSharedImage_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetAcceleratedNetworkSupportEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetArchitectureParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetConfidentialVmEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetConfidentialVmSupportedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetDiskTypesNotAllowedParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetEndOfLifeDateParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetEulaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetGalleryNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetHyperVGenerationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetMaxRecommendedMemoryInGbParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetMaxRecommendedVcpuCountParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetMinRecommendedMemoryInGbParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetMinRecommendedVcpuCountParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetOsTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetPrivacyStatementUriParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetReleaseNoteUriParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetResourceGroupNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetSpecializedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_SharedImage) validateSetTrustedLaunchEnabledParameters(val interface{}) error {
	return nil
}

func validateNewSharedImageParameters(scope constructs.Construct, id *string, config *SharedImageConfig) error {
	return nil
}

