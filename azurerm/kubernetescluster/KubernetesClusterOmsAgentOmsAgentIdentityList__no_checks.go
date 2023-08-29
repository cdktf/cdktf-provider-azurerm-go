// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kubernetescluster

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKubernetesClusterOmsAgentOmsAgentIdentityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

