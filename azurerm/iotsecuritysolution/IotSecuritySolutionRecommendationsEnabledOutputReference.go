// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotsecuritysolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/iotsecuritysolution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotSecuritySolutionRecommendationsEnabledOutputReference interface {
	cdktf.ComplexObject
	AcrAuthentication() interface{}
	SetAcrAuthentication(val interface{})
	AcrAuthenticationInput() interface{}
	AgentSendUnutilizedMsg() interface{}
	SetAgentSendUnutilizedMsg(val interface{})
	AgentSendUnutilizedMsgInput() interface{}
	Baseline() interface{}
	SetBaseline(val interface{})
	BaselineInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EdgeHubMemOptimize() interface{}
	SetEdgeHubMemOptimize(val interface{})
	EdgeHubMemOptimizeInput() interface{}
	EdgeLoggingOption() interface{}
	SetEdgeLoggingOption(val interface{})
	EdgeLoggingOptionInput() interface{}
	// Experimental.
	Fqn() *string
	InconsistentModuleSettings() interface{}
	SetInconsistentModuleSettings(val interface{})
	InconsistentModuleSettingsInput() interface{}
	InstallAgent() interface{}
	SetInstallAgent(val interface{})
	InstallAgentInput() interface{}
	InternalValue() *IotSecuritySolutionRecommendationsEnabled
	SetInternalValue(val *IotSecuritySolutionRecommendationsEnabled)
	IpFilterDenyAll() interface{}
	SetIpFilterDenyAll(val interface{})
	IpFilterDenyAllInput() interface{}
	IpFilterPermissiveRule() interface{}
	SetIpFilterPermissiveRule(val interface{})
	IpFilterPermissiveRuleInput() interface{}
	OpenPorts() interface{}
	SetOpenPorts(val interface{})
	OpenPortsInput() interface{}
	PermissiveFirewallPolicy() interface{}
	SetPermissiveFirewallPolicy(val interface{})
	PermissiveFirewallPolicyInput() interface{}
	PermissiveInputFirewallRules() interface{}
	SetPermissiveInputFirewallRules(val interface{})
	PermissiveInputFirewallRulesInput() interface{}
	PermissiveOutputFirewallRules() interface{}
	SetPermissiveOutputFirewallRules(val interface{})
	PermissiveOutputFirewallRulesInput() interface{}
	PrivilegedDockerOptions() interface{}
	SetPrivilegedDockerOptions(val interface{})
	PrivilegedDockerOptionsInput() interface{}
	SharedCredentials() interface{}
	SetSharedCredentials(val interface{})
	SharedCredentialsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VulnerableTlsCipherSuite() interface{}
	SetVulnerableTlsCipherSuite(val interface{})
	VulnerableTlsCipherSuiteInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAcrAuthentication()
	ResetAgentSendUnutilizedMsg()
	ResetBaseline()
	ResetEdgeHubMemOptimize()
	ResetEdgeLoggingOption()
	ResetInconsistentModuleSettings()
	ResetInstallAgent()
	ResetIpFilterDenyAll()
	ResetIpFilterPermissiveRule()
	ResetOpenPorts()
	ResetPermissiveFirewallPolicy()
	ResetPermissiveInputFirewallRules()
	ResetPermissiveOutputFirewallRules()
	ResetPrivilegedDockerOptions()
	ResetSharedCredentials()
	ResetVulnerableTlsCipherSuite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotSecuritySolutionRecommendationsEnabledOutputReference
type jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) AcrAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acrAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) AcrAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acrAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) AgentSendUnutilizedMsg() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentSendUnutilizedMsg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) AgentSendUnutilizedMsgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentSendUnutilizedMsgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) Baseline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) BaselineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) EdgeHubMemOptimize() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeHubMemOptimize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) EdgeHubMemOptimizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeHubMemOptimizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) EdgeLoggingOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeLoggingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) EdgeLoggingOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeLoggingOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InconsistentModuleSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inconsistentModuleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InconsistentModuleSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inconsistentModuleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InstallAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InstallAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InternalValue() *IotSecuritySolutionRecommendationsEnabled {
	var returns *IotSecuritySolutionRecommendationsEnabled
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) IpFilterDenyAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterDenyAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) IpFilterDenyAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterDenyAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) IpFilterPermissiveRule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterPermissiveRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) IpFilterPermissiveRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterPermissiveRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) OpenPorts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) OpenPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveFirewallPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveFirewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveFirewallPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveFirewallPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveInputFirewallRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveInputFirewallRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveInputFirewallRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveInputFirewallRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveOutputFirewallRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveOutputFirewallRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PermissiveOutputFirewallRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveOutputFirewallRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PrivilegedDockerOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedDockerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) PrivilegedDockerOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedDockerOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SharedCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) SharedCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) VulnerableTlsCipherSuite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerableTlsCipherSuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) VulnerableTlsCipherSuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerableTlsCipherSuiteInput",
		&returns,
	)
	return returns
}


func NewIotSecuritySolutionRecommendationsEnabledOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotSecuritySolutionRecommendationsEnabledOutputReference {
	_init_.Initialize()

	if err := validateNewIotSecuritySolutionRecommendationsEnabledOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionRecommendationsEnabledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotSecuritySolutionRecommendationsEnabledOutputReference_Override(i IotSecuritySolutionRecommendationsEnabledOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecuritySolution.IotSecuritySolutionRecommendationsEnabledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetAcrAuthentication(val interface{}) {
	if err := j.validateSetAcrAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acrAuthentication",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetAgentSendUnutilizedMsg(val interface{}) {
	if err := j.validateSetAgentSendUnutilizedMsgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentSendUnutilizedMsg",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetBaseline(val interface{}) {
	if err := j.validateSetBaselineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseline",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetEdgeHubMemOptimize(val interface{}) {
	if err := j.validateSetEdgeHubMemOptimizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeHubMemOptimize",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetEdgeLoggingOption(val interface{}) {
	if err := j.validateSetEdgeLoggingOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeLoggingOption",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetInconsistentModuleSettings(val interface{}) {
	if err := j.validateSetInconsistentModuleSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inconsistentModuleSettings",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetInstallAgent(val interface{}) {
	if err := j.validateSetInstallAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installAgent",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetInternalValue(val *IotSecuritySolutionRecommendationsEnabled) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetIpFilterDenyAll(val interface{}) {
	if err := j.validateSetIpFilterDenyAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFilterDenyAll",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetIpFilterPermissiveRule(val interface{}) {
	if err := j.validateSetIpFilterPermissiveRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFilterPermissiveRule",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetOpenPorts(val interface{}) {
	if err := j.validateSetOpenPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openPorts",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetPermissiveFirewallPolicy(val interface{}) {
	if err := j.validateSetPermissiveFirewallPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissiveFirewallPolicy",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetPermissiveInputFirewallRules(val interface{}) {
	if err := j.validateSetPermissiveInputFirewallRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissiveInputFirewallRules",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetPermissiveOutputFirewallRules(val interface{}) {
	if err := j.validateSetPermissiveOutputFirewallRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissiveOutputFirewallRules",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetPrivilegedDockerOptions(val interface{}) {
	if err := j.validateSetPrivilegedDockerOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedDockerOptions",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetSharedCredentials(val interface{}) {
	if err := j.validateSetSharedCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedCredentials",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference)SetVulnerableTlsCipherSuite(val interface{}) {
	if err := j.validateSetVulnerableTlsCipherSuiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vulnerableTlsCipherSuite",
		val,
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetAcrAuthentication() {
	_jsii_.InvokeVoid(
		i,
		"resetAcrAuthentication",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetAgentSendUnutilizedMsg() {
	_jsii_.InvokeVoid(
		i,
		"resetAgentSendUnutilizedMsg",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetBaseline() {
	_jsii_.InvokeVoid(
		i,
		"resetBaseline",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetEdgeHubMemOptimize() {
	_jsii_.InvokeVoid(
		i,
		"resetEdgeHubMemOptimize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetEdgeLoggingOption() {
	_jsii_.InvokeVoid(
		i,
		"resetEdgeLoggingOption",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetInconsistentModuleSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetInconsistentModuleSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetInstallAgent() {
	_jsii_.InvokeVoid(
		i,
		"resetInstallAgent",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetIpFilterDenyAll() {
	_jsii_.InvokeVoid(
		i,
		"resetIpFilterDenyAll",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetIpFilterPermissiveRule() {
	_jsii_.InvokeVoid(
		i,
		"resetIpFilterPermissiveRule",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetOpenPorts() {
	_jsii_.InvokeVoid(
		i,
		"resetOpenPorts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetPermissiveFirewallPolicy() {
	_jsii_.InvokeVoid(
		i,
		"resetPermissiveFirewallPolicy",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetPermissiveInputFirewallRules() {
	_jsii_.InvokeVoid(
		i,
		"resetPermissiveInputFirewallRules",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetPermissiveOutputFirewallRules() {
	_jsii_.InvokeVoid(
		i,
		"resetPermissiveOutputFirewallRules",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetPrivilegedDockerOptions() {
	_jsii_.InvokeVoid(
		i,
		"resetPrivilegedDockerOptions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetSharedCredentials() {
	_jsii_.InvokeVoid(
		i,
		"resetSharedCredentials",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ResetVulnerableTlsCipherSuite() {
	_jsii_.InvokeVoid(
		i,
		"resetVulnerableTlsCipherSuite",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsEnabledOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

