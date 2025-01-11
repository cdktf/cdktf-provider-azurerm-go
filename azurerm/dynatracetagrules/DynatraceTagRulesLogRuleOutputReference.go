// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynatracetagrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/dynatracetagrules/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DynatraceTagRulesLogRuleOutputReference interface {
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
	FilteringTag() DynatraceTagRulesLogRuleFilteringTagList
	FilteringTagInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DynatraceTagRulesLogRule
	SetInternalValue(val *DynatraceTagRulesLogRule)
	SendActivityLogsEnabled() interface{}
	SetSendActivityLogsEnabled(val interface{})
	SendActivityLogsEnabledInput() interface{}
	SendAzureActiveDirectoryLogsEnabled() interface{}
	SetSendAzureActiveDirectoryLogsEnabled(val interface{})
	SendAzureActiveDirectoryLogsEnabledInput() interface{}
	SendSubscriptionLogsEnabled() interface{}
	SetSendSubscriptionLogsEnabled(val interface{})
	SendSubscriptionLogsEnabledInput() interface{}
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
	PutFilteringTag(value interface{})
	ResetSendActivityLogsEnabled()
	ResetSendAzureActiveDirectoryLogsEnabled()
	ResetSendSubscriptionLogsEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynatraceTagRulesLogRuleOutputReference
type jsiiProxy_DynatraceTagRulesLogRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) FilteringTag() DynatraceTagRulesLogRuleFilteringTagList {
	var returns DynatraceTagRulesLogRuleFilteringTagList
	_jsii_.Get(
		j,
		"filteringTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) FilteringTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filteringTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) InternalValue() *DynatraceTagRulesLogRule {
	var returns *DynatraceTagRulesLogRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) SendActivityLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendActivityLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) SendActivityLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendActivityLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) SendAzureActiveDirectoryLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendAzureActiveDirectoryLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) SendAzureActiveDirectoryLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendAzureActiveDirectoryLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) SendSubscriptionLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendSubscriptionLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) SendSubscriptionLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendSubscriptionLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDynatraceTagRulesLogRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DynatraceTagRulesLogRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDynatraceTagRulesLogRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynatraceTagRulesLogRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dynatraceTagRules.DynatraceTagRulesLogRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDynatraceTagRulesLogRuleOutputReference_Override(d DynatraceTagRulesLogRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dynatraceTagRules.DynatraceTagRulesLogRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference)SetInternalValue(val *DynatraceTagRulesLogRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference)SetSendActivityLogsEnabled(val interface{}) {
	if err := j.validateSetSendActivityLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendActivityLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference)SetSendAzureActiveDirectoryLogsEnabled(val interface{}) {
	if err := j.validateSetSendAzureActiveDirectoryLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendAzureActiveDirectoryLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference)SetSendSubscriptionLogsEnabled(val interface{}) {
	if err := j.validateSetSendSubscriptionLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendSubscriptionLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DynatraceTagRulesLogRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) PutFilteringTag(value interface{}) {
	if err := d.validatePutFilteringTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilteringTag",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) ResetSendActivityLogsEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetSendActivityLogsEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) ResetSendAzureActiveDirectoryLogsEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetSendAzureActiveDirectoryLogsEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) ResetSendSubscriptionLogsEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetSendSubscriptionLogsEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynatraceTagRulesLogRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

