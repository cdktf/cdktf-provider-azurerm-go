// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediaassetfilter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/mediaassetfilter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaAssetFilterPresentationTimeRangeOutputReference interface {
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
	EndInUnits() *float64
	SetEndInUnits(val *float64)
	EndInUnitsInput() *float64
	ForceEnd() interface{}
	SetForceEnd(val interface{})
	ForceEndInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *MediaAssetFilterPresentationTimeRange
	SetInternalValue(val *MediaAssetFilterPresentationTimeRange)
	LiveBackoffInUnits() *float64
	SetLiveBackoffInUnits(val *float64)
	LiveBackoffInUnitsInput() *float64
	PresentationWindowInUnits() *float64
	SetPresentationWindowInUnits(val *float64)
	PresentationWindowInUnitsInput() *float64
	StartInUnits() *float64
	SetStartInUnits(val *float64)
	StartInUnitsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnitTimescaleInMiliseconds() *float64
	SetUnitTimescaleInMiliseconds(val *float64)
	UnitTimescaleInMilisecondsInput() *float64
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
	ResetEndInUnits()
	ResetForceEnd()
	ResetLiveBackoffInUnits()
	ResetPresentationWindowInUnits()
	ResetStartInUnits()
	ResetUnitTimescaleInMiliseconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaAssetFilterPresentationTimeRangeOutputReference
type jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) EndInUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endInUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) EndInUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endInUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ForceEnd() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ForceEndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) InternalValue() *MediaAssetFilterPresentationTimeRange {
	var returns *MediaAssetFilterPresentationTimeRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) LiveBackoffInUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"liveBackoffInUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) LiveBackoffInUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"liveBackoffInUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) PresentationWindowInUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"presentationWindowInUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) PresentationWindowInUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"presentationWindowInUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) StartInUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startInUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) StartInUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startInUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) UnitTimescaleInMiliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unitTimescaleInMiliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) UnitTimescaleInMilisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unitTimescaleInMilisecondsInput",
		&returns,
	)
	return returns
}


func NewMediaAssetFilterPresentationTimeRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaAssetFilterPresentationTimeRangeOutputReference {
	_init_.Initialize()

	if err := validateNewMediaAssetFilterPresentationTimeRangeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaAssetFilter.MediaAssetFilterPresentationTimeRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaAssetFilterPresentationTimeRangeOutputReference_Override(m MediaAssetFilterPresentationTimeRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaAssetFilter.MediaAssetFilterPresentationTimeRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference)SetEndInUnits(val *float64) {
	if err := j.validateSetEndInUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endInUnits",
		val,
	)
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference)SetForceEnd(val interface{}) {
	if err := j.validateSetForceEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceEnd",
		val,
	)
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference)SetInternalValue(val *MediaAssetFilterPresentationTimeRange) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference)SetLiveBackoffInUnits(val *float64) {
	if err := j.validateSetLiveBackoffInUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveBackoffInUnits",
		val,
	)
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference)SetPresentationWindowInUnits(val *float64) {
	if err := j.validateSetPresentationWindowInUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"presentationWindowInUnits",
		val,
	)
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference)SetStartInUnits(val *float64) {
	if err := j.validateSetStartInUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startInUnits",
		val,
	)
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference)SetUnitTimescaleInMiliseconds(val *float64) {
	if err := j.validateSetUnitTimescaleInMilisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitTimescaleInMiliseconds",
		val,
	)
}

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ResetEndInUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetEndInUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ResetForceEnd() {
	_jsii_.InvokeVoid(
		m,
		"resetForceEnd",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ResetLiveBackoffInUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetLiveBackoffInUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ResetPresentationWindowInUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetPresentationWindowInUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ResetStartInUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetStartInUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ResetUnitTimescaleInMiliseconds() {
	_jsii_.InvokeVoid(
		m,
		"resetUnitTimescaleInMiliseconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaAssetFilterPresentationTimeRangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

