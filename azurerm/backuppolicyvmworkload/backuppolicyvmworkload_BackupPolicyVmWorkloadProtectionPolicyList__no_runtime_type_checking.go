//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package backuppolicyvmworkload

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupPolicyVmWorkloadProtectionPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

