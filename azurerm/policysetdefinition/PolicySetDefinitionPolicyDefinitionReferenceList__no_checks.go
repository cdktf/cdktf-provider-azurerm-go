// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package policysetdefinition

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicySetDefinitionPolicyDefinitionReferenceList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PolicySetDefinitionPolicyDefinitionReferenceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PolicySetDefinitionPolicyDefinitionReferenceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicySetDefinitionPolicyDefinitionReferenceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PolicySetDefinitionPolicyDefinitionReferenceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PolicySetDefinitionPolicyDefinitionReferenceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPolicySetDefinitionPolicyDefinitionReferenceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

