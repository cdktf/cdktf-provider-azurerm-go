// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediacontentkeypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/mediacontentkeypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference interface {
	cdktf.ComplexObject
	AgcAndColorStripeRestriction() *float64
	SetAgcAndColorStripeRestriction(val *float64)
	AgcAndColorStripeRestrictionInput() *float64
	AllowPassingVideoContentToUnknownOutput() *string
	SetAllowPassingVideoContentToUnknownOutput(val *string)
	AllowPassingVideoContentToUnknownOutputInput() *string
	AnalogVideoOpl() *float64
	SetAnalogVideoOpl(val *float64)
	AnalogVideoOplInput() *float64
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
	CompressedDigitalAudioOpl() *float64
	SetCompressedDigitalAudioOpl(val *float64)
	CompressedDigitalAudioOplInput() *float64
	CompressedDigitalVideoOpl() *float64
	SetCompressedDigitalVideoOpl(val *float64)
	CompressedDigitalVideoOplInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DigitalVideoOnlyContentRestriction() interface{}
	SetDigitalVideoOnlyContentRestriction(val interface{})
	DigitalVideoOnlyContentRestrictionInput() interface{}
	ExplicitAnalogTelevisionOutputRestriction() MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightExplicitAnalogTelevisionOutputRestrictionOutputReference
	ExplicitAnalogTelevisionOutputRestrictionInput() *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightExplicitAnalogTelevisionOutputRestriction
	FirstPlayExpiration() *string
	SetFirstPlayExpiration(val *string)
	FirstPlayExpirationInput() *string
	// Experimental.
	Fqn() *string
	ImageConstraintForAnalogComponentVideoRestriction() interface{}
	SetImageConstraintForAnalogComponentVideoRestriction(val interface{})
	ImageConstraintForAnalogComponentVideoRestrictionInput() interface{}
	ImageConstraintForAnalogComputerMonitorRestriction() interface{}
	SetImageConstraintForAnalogComputerMonitorRestriction(val interface{})
	ImageConstraintForAnalogComputerMonitorRestrictionInput() interface{}
	InternalValue() *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight
	SetInternalValue(val *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight)
	ScmsRestriction() *float64
	SetScmsRestriction(val *float64)
	ScmsRestrictionInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UncompressedDigitalAudioOpl() *float64
	SetUncompressedDigitalAudioOpl(val *float64)
	UncompressedDigitalAudioOplInput() *float64
	UncompressedDigitalVideoOpl() *float64
	SetUncompressedDigitalVideoOpl(val *float64)
	UncompressedDigitalVideoOplInput() *float64
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
	PutExplicitAnalogTelevisionOutputRestriction(value *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightExplicitAnalogTelevisionOutputRestriction)
	ResetAgcAndColorStripeRestriction()
	ResetAllowPassingVideoContentToUnknownOutput()
	ResetAnalogVideoOpl()
	ResetCompressedDigitalAudioOpl()
	ResetCompressedDigitalVideoOpl()
	ResetDigitalVideoOnlyContentRestriction()
	ResetExplicitAnalogTelevisionOutputRestriction()
	ResetFirstPlayExpiration()
	ResetImageConstraintForAnalogComponentVideoRestriction()
	ResetImageConstraintForAnalogComputerMonitorRestriction()
	ResetScmsRestriction()
	ResetUncompressedDigitalAudioOpl()
	ResetUncompressedDigitalVideoOpl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AgcAndColorStripeRestriction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"agcAndColorStripeRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AgcAndColorStripeRestrictionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"agcAndColorStripeRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AllowPassingVideoContentToUnknownOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowPassingVideoContentToUnknownOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AllowPassingVideoContentToUnknownOutputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowPassingVideoContentToUnknownOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AnalogVideoOpl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"analogVideoOpl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AnalogVideoOplInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"analogVideoOplInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) CompressedDigitalAudioOpl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"compressedDigitalAudioOpl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) CompressedDigitalAudioOplInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"compressedDigitalAudioOplInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) CompressedDigitalVideoOpl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"compressedDigitalVideoOpl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) CompressedDigitalVideoOplInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"compressedDigitalVideoOplInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) DigitalVideoOnlyContentRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"digitalVideoOnlyContentRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) DigitalVideoOnlyContentRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"digitalVideoOnlyContentRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ExplicitAnalogTelevisionOutputRestriction() MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightExplicitAnalogTelevisionOutputRestrictionOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightExplicitAnalogTelevisionOutputRestrictionOutputReference
	_jsii_.Get(
		j,
		"explicitAnalogTelevisionOutputRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ExplicitAnalogTelevisionOutputRestrictionInput() *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightExplicitAnalogTelevisionOutputRestriction {
	var returns *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightExplicitAnalogTelevisionOutputRestriction
	_jsii_.Get(
		j,
		"explicitAnalogTelevisionOutputRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) FirstPlayExpiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstPlayExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) FirstPlayExpirationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstPlayExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ImageConstraintForAnalogComponentVideoRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConstraintForAnalogComponentVideoRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ImageConstraintForAnalogComponentVideoRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConstraintForAnalogComponentVideoRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ImageConstraintForAnalogComputerMonitorRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConstraintForAnalogComputerMonitorRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ImageConstraintForAnalogComputerMonitorRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConstraintForAnalogComputerMonitorRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) InternalValue() *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight {
	var returns *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ScmsRestriction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scmsRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ScmsRestrictionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scmsRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) UncompressedDigitalAudioOpl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uncompressedDigitalAudioOpl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) UncompressedDigitalAudioOplInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uncompressedDigitalAudioOplInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) UncompressedDigitalVideoOpl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uncompressedDigitalVideoOpl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) UncompressedDigitalVideoOplInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uncompressedDigitalVideoOplInput",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference {
	_init_.Initialize()

	if err := validateNewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference_Override(m MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetAgcAndColorStripeRestriction(val *float64) {
	if err := j.validateSetAgcAndColorStripeRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agcAndColorStripeRestriction",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetAllowPassingVideoContentToUnknownOutput(val *string) {
	if err := j.validateSetAllowPassingVideoContentToUnknownOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPassingVideoContentToUnknownOutput",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetAnalogVideoOpl(val *float64) {
	if err := j.validateSetAnalogVideoOplParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analogVideoOpl",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetCompressedDigitalAudioOpl(val *float64) {
	if err := j.validateSetCompressedDigitalAudioOplParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressedDigitalAudioOpl",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetCompressedDigitalVideoOpl(val *float64) {
	if err := j.validateSetCompressedDigitalVideoOplParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressedDigitalVideoOpl",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetDigitalVideoOnlyContentRestriction(val interface{}) {
	if err := j.validateSetDigitalVideoOnlyContentRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digitalVideoOnlyContentRestriction",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetFirstPlayExpiration(val *string) {
	if err := j.validateSetFirstPlayExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstPlayExpiration",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetImageConstraintForAnalogComponentVideoRestriction(val interface{}) {
	if err := j.validateSetImageConstraintForAnalogComponentVideoRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageConstraintForAnalogComponentVideoRestriction",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetImageConstraintForAnalogComputerMonitorRestriction(val interface{}) {
	if err := j.validateSetImageConstraintForAnalogComputerMonitorRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageConstraintForAnalogComputerMonitorRestriction",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetInternalValue(val *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetScmsRestriction(val *float64) {
	if err := j.validateSetScmsRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmsRestriction",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetUncompressedDigitalAudioOpl(val *float64) {
	if err := j.validateSetUncompressedDigitalAudioOplParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uncompressedDigitalAudioOpl",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference)SetUncompressedDigitalVideoOpl(val *float64) {
	if err := j.validateSetUncompressedDigitalVideoOplParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uncompressedDigitalVideoOpl",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) PutExplicitAnalogTelevisionOutputRestriction(value *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightExplicitAnalogTelevisionOutputRestriction) {
	if err := m.validatePutExplicitAnalogTelevisionOutputRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putExplicitAnalogTelevisionOutputRestriction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetAgcAndColorStripeRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetAgcAndColorStripeRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetAllowPassingVideoContentToUnknownOutput() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowPassingVideoContentToUnknownOutput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetAnalogVideoOpl() {
	_jsii_.InvokeVoid(
		m,
		"resetAnalogVideoOpl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetCompressedDigitalAudioOpl() {
	_jsii_.InvokeVoid(
		m,
		"resetCompressedDigitalAudioOpl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetCompressedDigitalVideoOpl() {
	_jsii_.InvokeVoid(
		m,
		"resetCompressedDigitalVideoOpl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetDigitalVideoOnlyContentRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetDigitalVideoOnlyContentRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetExplicitAnalogTelevisionOutputRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetExplicitAnalogTelevisionOutputRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetFirstPlayExpiration() {
	_jsii_.InvokeVoid(
		m,
		"resetFirstPlayExpiration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetImageConstraintForAnalogComponentVideoRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetImageConstraintForAnalogComponentVideoRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetImageConstraintForAnalogComputerMonitorRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetImageConstraintForAnalogComputerMonitorRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetScmsRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetScmsRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetUncompressedDigitalAudioOpl() {
	_jsii_.InvokeVoid(
		m,
		"resetUncompressedDigitalAudioOpl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetUncompressedDigitalVideoOpl() {
	_jsii_.InvokeVoid(
		m,
		"resetUncompressedDigitalVideoOpl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

