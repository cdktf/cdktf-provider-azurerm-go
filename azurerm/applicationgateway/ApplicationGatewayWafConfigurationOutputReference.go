// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewayWafConfigurationOutputReference interface {
	cdktf.ComplexObject
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
	DisabledRuleGroup() ApplicationGatewayWafConfigurationDisabledRuleGroupList
	DisabledRuleGroupInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Exclusion() ApplicationGatewayWafConfigurationExclusionList
	ExclusionInput() interface{}
	FileUploadLimitMb() *float64
	SetFileUploadLimitMb(val *float64)
	FileUploadLimitMbInput() *float64
	FirewallMode() *string
	SetFirewallMode(val *string)
	FirewallModeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ApplicationGatewayWafConfiguration
	SetInternalValue(val *ApplicationGatewayWafConfiguration)
	MaxRequestBodySizeKb() *float64
	SetMaxRequestBodySizeKb(val *float64)
	MaxRequestBodySizeKbInput() *float64
	RequestBodyCheck() interface{}
	SetRequestBodyCheck(val interface{})
	RequestBodyCheckInput() interface{}
	RuleSetType() *string
	SetRuleSetType(val *string)
	RuleSetTypeInput() *string
	RuleSetVersion() *string
	SetRuleSetVersion(val *string)
	RuleSetVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutDisabledRuleGroup(value interface{})
	PutExclusion(value interface{})
	ResetDisabledRuleGroup()
	ResetExclusion()
	ResetFileUploadLimitMb()
	ResetMaxRequestBodySizeKb()
	ResetRequestBodyCheck()
	ResetRuleSetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewayWafConfigurationOutputReference
type jsiiProxy_ApplicationGatewayWafConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) DisabledRuleGroup() ApplicationGatewayWafConfigurationDisabledRuleGroupList {
	var returns ApplicationGatewayWafConfigurationDisabledRuleGroupList
	_jsii_.Get(
		j,
		"disabledRuleGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) DisabledRuleGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledRuleGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) Exclusion() ApplicationGatewayWafConfigurationExclusionList {
	var returns ApplicationGatewayWafConfigurationExclusionList
	_jsii_.Get(
		j,
		"exclusion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) ExclusionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exclusionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) FileUploadLimitMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileUploadLimitMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) FileUploadLimitMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileUploadLimitMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) FirewallMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) FirewallModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) InternalValue() *ApplicationGatewayWafConfiguration {
	var returns *ApplicationGatewayWafConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) MaxRequestBodySizeKb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRequestBodySizeKb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) MaxRequestBodySizeKbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRequestBodySizeKbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) RequestBodyCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBodyCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) RequestBodyCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBodyCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) RuleSetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) RuleSetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) RuleSetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) RuleSetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationGatewayWafConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationGatewayWafConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewayWafConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewayWafConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayWafConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationGatewayWafConfigurationOutputReference_Override(a ApplicationGatewayWafConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayWafConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetFileUploadLimitMb(val *float64) {
	if err := j.validateSetFileUploadLimitMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileUploadLimitMb",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetFirewallMode(val *string) {
	if err := j.validateSetFirewallModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallMode",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetInternalValue(val *ApplicationGatewayWafConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetMaxRequestBodySizeKb(val *float64) {
	if err := j.validateSetMaxRequestBodySizeKbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRequestBodySizeKb",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetRequestBodyCheck(val interface{}) {
	if err := j.validateSetRequestBodyCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBodyCheck",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetRuleSetType(val *string) {
	if err := j.validateSetRuleSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleSetType",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetRuleSetVersion(val *string) {
	if err := j.validateSetRuleSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleSetVersion",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) PutDisabledRuleGroup(value interface{}) {
	if err := a.validatePutDisabledRuleGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDisabledRuleGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) PutExclusion(value interface{}) {
	if err := a.validatePutExclusionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExclusion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) ResetDisabledRuleGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetDisabledRuleGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) ResetExclusion() {
	_jsii_.InvokeVoid(
		a,
		"resetExclusion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) ResetFileUploadLimitMb() {
	_jsii_.InvokeVoid(
		a,
		"resetFileUploadLimitMb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) ResetMaxRequestBodySizeKb() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxRequestBodySizeKb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) ResetRequestBodyCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestBodyCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) ResetRuleSetType() {
	_jsii_.InvokeVoid(
		a,
		"resetRuleSetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayWafConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

