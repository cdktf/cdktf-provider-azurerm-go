// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialiveevent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/medialiveevent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaLiveEventPreviewOutputReference interface {
	cdktf.ComplexObject
	AlternativeMediaId() *string
	SetAlternativeMediaId(val *string)
	AlternativeMediaIdInput() *string
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
	Endpoint() MediaLiveEventPreviewEndpointList
	// Experimental.
	Fqn() *string
	InternalValue() *MediaLiveEventPreview
	SetInternalValue(val *MediaLiveEventPreview)
	IpAccessControlAllow() MediaLiveEventPreviewIpAccessControlAllowList
	IpAccessControlAllowInput() interface{}
	PreviewLocator() *string
	SetPreviewLocator(val *string)
	PreviewLocatorInput() *string
	StreamingPolicyName() *string
	SetStreamingPolicyName(val *string)
	StreamingPolicyNameInput() *string
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
	PutIpAccessControlAllow(value interface{})
	ResetAlternativeMediaId()
	ResetIpAccessControlAllow()
	ResetPreviewLocator()
	ResetStreamingPolicyName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaLiveEventPreviewOutputReference
type jsiiProxy_MediaLiveEventPreviewOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) AlternativeMediaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternativeMediaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) AlternativeMediaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternativeMediaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) Endpoint() MediaLiveEventPreviewEndpointList {
	var returns MediaLiveEventPreviewEndpointList
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) InternalValue() *MediaLiveEventPreview {
	var returns *MediaLiveEventPreview
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) IpAccessControlAllow() MediaLiveEventPreviewIpAccessControlAllowList {
	var returns MediaLiveEventPreviewIpAccessControlAllowList
	_jsii_.Get(
		j,
		"ipAccessControlAllow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) IpAccessControlAllowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipAccessControlAllowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) PreviewLocator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"previewLocator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) PreviewLocatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"previewLocatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) StreamingPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamingPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) StreamingPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamingPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaLiveEventPreviewOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaLiveEventPreviewOutputReference {
	_init_.Initialize()

	if err := validateNewMediaLiveEventPreviewOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaLiveEventPreviewOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEventPreviewOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaLiveEventPreviewOutputReference_Override(m MediaLiveEventPreviewOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaLiveEvent.MediaLiveEventPreviewOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference)SetAlternativeMediaId(val *string) {
	if err := j.validateSetAlternativeMediaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternativeMediaId",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference)SetInternalValue(val *MediaLiveEventPreview) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference)SetPreviewLocator(val *string) {
	if err := j.validateSetPreviewLocatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"previewLocator",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference)SetStreamingPolicyName(val *string) {
	if err := j.validateSetStreamingPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingPolicyName",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventPreviewOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) PutIpAccessControlAllow(value interface{}) {
	if err := m.validatePutIpAccessControlAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIpAccessControlAllow",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) ResetAlternativeMediaId() {
	_jsii_.InvokeVoid(
		m,
		"resetAlternativeMediaId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) ResetIpAccessControlAllow() {
	_jsii_.InvokeVoid(
		m,
		"resetIpAccessControlAllow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) ResetPreviewLocator() {
	_jsii_.InvokeVoid(
		m,
		"resetPreviewLocator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) ResetStreamingPolicyName() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamingPolicyName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaLiveEventPreviewOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

