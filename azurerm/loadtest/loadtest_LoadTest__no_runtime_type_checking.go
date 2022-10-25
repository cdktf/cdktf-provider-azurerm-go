//go:build no_runtime_type_checking

package loadtest

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadTest) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validatePutIdentityParameters(value *LoadTestIdentity) error {
	return nil
}

func (l *jsiiProxy_LoadTest) validatePutTimeoutsParameters(value *LoadTestTimeouts) error {
	return nil
}

func validateLoadTest_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadTest) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadTest) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadTest) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadTest) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_LoadTest) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadTest) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadTest) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadTest) validateSetResourceGroupNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadTest) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func validateNewLoadTestParameters(scope constructs.Construct, id *string, config *LoadTestConfig) error {
	return nil
}

