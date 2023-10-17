// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialiveevent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/medialiveevent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaLiveEventCrossSiteAccessPolicyOutputReference interface {
	cdktf.ComplexObject
	ClientAccessPolicy() *string
	SetClientAccessPolicy(val *string)
	ClientAccessPolicyInput() *string
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
	CrossDomainPolicy() *string
	SetCrossDomainPolicy(val *string)
	CrossDomainPolicyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MediaLiveEventCrossSiteAccessPolicy
	SetInternalValue(val *MediaLiveEventCrossSiteAccessPolicy)
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
	ResetClientAccessPolicy()
	ResetCrossDomainPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaLiveEventCrossSiteAccessPolicyOutputReference
type jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) ClientAccessPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) ClientAccessPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) CrossDomainPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossDomainPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) CrossDomainPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossDomainPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) InternalValue() *MediaLiveEventCrossSiteAccessPolicy {
	var returns *MediaLiveEventCrossSiteAccessPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaLiveEventCrossSiteAccessPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaLiveEventCrossSiteAccessPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewMediaLiveEventCrossSiteAccessPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEventCrossSiteAccessPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaLiveEventCrossSiteAccessPolicyOutputReference_Override(m MediaLiveEventCrossSiteAccessPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEventCrossSiteAccessPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference)SetClientAccessPolicy(val *string) {
	if err := j.validateSetClientAccessPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAccessPolicy",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference)SetCrossDomainPolicy(val *string) {
	if err := j.validateSetCrossDomainPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossDomainPolicy",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference)SetInternalValue(val *MediaLiveEventCrossSiteAccessPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) ResetClientAccessPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetClientAccessPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) ResetCrossDomainPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetCrossDomainPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventCrossSiteAccessPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

