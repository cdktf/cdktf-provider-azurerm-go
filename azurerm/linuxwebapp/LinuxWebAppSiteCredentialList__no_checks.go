// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package linuxwebapp

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxWebAppSiteCredentialList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LinuxWebAppSiteCredentialList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LinuxWebAppSiteCredentialList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LinuxWebAppSiteCredentialList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LinuxWebAppSiteCredentialList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LinuxWebAppSiteCredentialList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLinuxWebAppSiteCredentialListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

