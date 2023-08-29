// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databricksworkspace

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabricksWorkspaceManagedDiskIdentityList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabricksWorkspaceManagedDiskIdentityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabricksWorkspaceManagedDiskIdentityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabricksWorkspaceManagedDiskIdentityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabricksWorkspaceManagedDiskIdentityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabricksWorkspaceManagedDiskIdentityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

