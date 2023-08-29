// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kubernetescluster

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKubernetesClusterMaintenanceWindowNotAllowedListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

