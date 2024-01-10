// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediaservicesaccountfilter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/mediaservicesaccountfilter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaServicesAccountFilterPresentationTimeRangeOutputReference interface {
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
	InternalValue() *MediaServicesAccountFilterPresentationTimeRange
	SetInternalValue(val *MediaServicesAccountFilterPresentationTimeRange)
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
	UnitTimescaleInMilliseconds() *float64
	SetUnitTimescaleInMilliseconds(val *float64)
	UnitTimescaleInMillisecondsInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaServicesAccountFilterPresentationTimeRangeOutputReference
type jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) EndInUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endInUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) EndInUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endInUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) ForceEnd() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) ForceEndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) InternalValue() *MediaServicesAccountFilterPresentationTimeRange {
	var returns *MediaServicesAccountFilterPresentationTimeRange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) LiveBackoffInUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"liveBackoffInUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) LiveBackoffInUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"liveBackoffInUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) PresentationWindowInUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"presentationWindowInUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) PresentationWindowInUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"presentationWindowInUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) StartInUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startInUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) StartInUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startInUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) UnitTimescaleInMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unitTimescaleInMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) UnitTimescaleInMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unitTimescaleInMillisecondsInput",
		&returns,
	)
	return returns
}


func NewMediaServicesAccountFilterPresentationTimeRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaServicesAccountFilterPresentationTimeRangeOutputReference {
	_init_.Initialize()

	if err := validateNewMediaServicesAccountFilterPresentationTimeRangeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaServicesAccountFilter.MediaServicesAccountFilterPresentationTimeRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaServicesAccountFilterPresentationTimeRangeOutputReference_Override(m MediaServicesAccountFilterPresentationTimeRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaServicesAccountFilter.MediaServicesAccountFilterPresentationTimeRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference)SetEndInUnits(val *float64) {
	if err := j.validateSetEndInUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endInUnits",
		val,
	)
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference)SetForceEnd(val interface{}) {
	if err := j.validateSetForceEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceEnd",
		val,
	)
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference)SetInternalValue(val *MediaServicesAccountFilterPresentationTimeRange) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference)SetLiveBackoffInUnits(val *float64) {
	if err := j.validateSetLiveBackoffInUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveBackoffInUnits",
		val,
	)
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference)SetPresentationWindowInUnits(val *float64) {
	if err := j.validateSetPresentationWindowInUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"presentationWindowInUnits",
		val,
	)
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference)SetStartInUnits(val *float64) {
	if err := j.validateSetStartInUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startInUnits",
		val,
	)
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference)SetUnitTimescaleInMilliseconds(val *float64) {
	if err := j.validateSetUnitTimescaleInMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitTimescaleInMilliseconds",
		val,
	)
}

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) ResetEndInUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetEndInUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) ResetForceEnd() {
	_jsii_.InvokeVoid(
		m,
		"resetForceEnd",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) ResetLiveBackoffInUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetLiveBackoffInUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) ResetPresentationWindowInUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetPresentationWindowInUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) ResetStartInUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetStartInUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaServicesAccountFilterPresentationTimeRangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

