// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediacontentkeypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/mediacontentkeypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference interface {
	cdktf.ComplexObject
	Ask() *string
	SetAsk(val *string)
	AskInput() *string
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
	InternalValue() *MediaContentKeyPolicyPolicyOptionFairplayConfiguration
	SetInternalValue(val *MediaContentKeyPolicyPolicyOptionFairplayConfiguration)
	OfflineRentalConfiguration() MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference
	OfflineRentalConfigurationInput() *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration
	Pfx() *string
	SetPfx(val *string)
	PfxInput() *string
	PfxPassword() *string
	SetPfxPassword(val *string)
	PfxPasswordInput() *string
	RentalAndLeaseKeyType() *string
	SetRentalAndLeaseKeyType(val *string)
	RentalAndLeaseKeyTypeInput() *string
	RentalDurationSeconds() *float64
	SetRentalDurationSeconds(val *float64)
	RentalDurationSecondsInput() *float64
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
	PutOfflineRentalConfiguration(value *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration)
	ResetAsk()
	ResetOfflineRentalConfiguration()
	ResetPfx()
	ResetPfxPassword()
	ResetRentalAndLeaseKeyType()
	ResetRentalDurationSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) Ask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) AskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"askInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) InternalValue() *MediaContentKeyPolicyPolicyOptionFairplayConfiguration {
	var returns *MediaContentKeyPolicyPolicyOptionFairplayConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) OfflineRentalConfiguration() MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference
	_jsii_.Get(
		j,
		"offlineRentalConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) OfflineRentalConfigurationInput() *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration {
	var returns *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration
	_jsii_.Get(
		j,
		"offlineRentalConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) Pfx() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) PfxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) PfxPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) PfxPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) RentalAndLeaseKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rentalAndLeaseKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) RentalAndLeaseKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rentalAndLeaseKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) RentalDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rentalDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) RentalDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rentalDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference_Override(m MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference)SetAsk(val *string) {
	if err := j.validateSetAskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ask",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference)SetInternalValue(val *MediaContentKeyPolicyPolicyOptionFairplayConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference)SetPfx(val *string) {
	if err := j.validateSetPfxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pfx",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference)SetPfxPassword(val *string) {
	if err := j.validateSetPfxPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pfxPassword",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference)SetRentalAndLeaseKeyType(val *string) {
	if err := j.validateSetRentalAndLeaseKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rentalAndLeaseKeyType",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference)SetRentalDurationSeconds(val *float64) {
	if err := j.validateSetRentalDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rentalDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) PutOfflineRentalConfiguration(value *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration) {
	if err := m.validatePutOfflineRentalConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putOfflineRentalConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetAsk() {
	_jsii_.InvokeVoid(
		m,
		"resetAsk",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetOfflineRentalConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetOfflineRentalConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetPfx() {
	_jsii_.InvokeVoid(
		m,
		"resetPfx",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetPfxPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetPfxPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetRentalAndLeaseKeyType() {
	_jsii_.InvokeVoid(
		m,
		"resetRentalAndLeaseKeyType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetRentalDurationSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetRentalDurationSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

