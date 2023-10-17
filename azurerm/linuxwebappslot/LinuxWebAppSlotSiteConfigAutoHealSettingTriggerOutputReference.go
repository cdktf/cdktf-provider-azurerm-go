// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/linuxwebappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *LinuxWebAppSlotSiteConfigAutoHealSettingTrigger
	SetInternalValue(val *LinuxWebAppSlotSiteConfigAutoHealSettingTrigger)
	Requests() LinuxWebAppSlotSiteConfigAutoHealSettingTriggerRequestsOutputReference
	RequestsInput() *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerRequests
	SlowRequest() LinuxWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequestOutputReference
	SlowRequestInput() *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest
	StatusCode() LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList
	StatusCodeInput() interface{}
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
	PutRequests(value *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerRequests)
	PutSlowRequest(value *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest)
	PutStatusCode(value interface{})
	ResetRequests()
	ResetSlowRequest()
	ResetStatusCode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference
type jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) InternalValue() *LinuxWebAppSlotSiteConfigAutoHealSettingTrigger {
	var returns *LinuxWebAppSlotSiteConfigAutoHealSettingTrigger
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) Requests() LinuxWebAppSlotSiteConfigAutoHealSettingTriggerRequestsOutputReference {
	var returns LinuxWebAppSlotSiteConfigAutoHealSettingTriggerRequestsOutputReference
	_jsii_.Get(
		j,
		"requests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) RequestsInput() *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerRequests {
	var returns *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerRequests
	_jsii_.Get(
		j,
		"requestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) SlowRequest() LinuxWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequestOutputReference {
	var returns LinuxWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequestOutputReference
	_jsii_.Get(
		j,
		"slowRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) SlowRequestInput() *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest {
	var returns *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest
	_jsii_.Get(
		j,
		"slowRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) StatusCode() LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList {
	var returns LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) StatusCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference {
	_init_.Initialize()

	if err := validateNewLinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxWebAppSlot.LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference_Override(l LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxWebAppSlot.LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference)SetInternalValue(val *LinuxWebAppSlotSiteConfigAutoHealSettingTrigger) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) PutRequests(value *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerRequests) {
	if err := l.validatePutRequestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRequests",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) PutSlowRequest(value *LinuxWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest) {
	if err := l.validatePutSlowRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSlowRequest",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) PutStatusCode(value interface{}) {
	if err := l.validatePutStatusCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStatusCode",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ResetRequests() {
	_jsii_.InvokeVoid(
		l,
		"resetRequests",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ResetSlowRequest() {
	_jsii_.InvokeVoid(
		l,
		"resetSlowRequest",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ResetStatusCode() {
	_jsii_.InvokeVoid(
		l,
		"resetStatusCode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

