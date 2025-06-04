// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkslice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/mobilenetworkslice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference interface {
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
	InternalValue() *MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformation
	SetInternalValue(val *MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformation)
	SliceDifferentiator() *string
	SetSliceDifferentiator(val *string)
	SliceDifferentiatorInput() *string
	SliceServiceType() *float64
	SetSliceServiceType(val *float64)
	SliceServiceTypeInput() *float64
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
	ResetSliceDifferentiator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference
type jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) InternalValue() *MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformation {
	var returns *MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) SliceDifferentiator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sliceDifferentiator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) SliceDifferentiatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sliceDifferentiatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) SliceServiceType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sliceServiceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) SliceServiceTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sliceServiceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference {
	_init_.Initialize()

	if err := validateNewMobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkSlice.MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference_Override(m MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkSlice.MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference)SetInternalValue(val *MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference)SetSliceDifferentiator(val *string) {
	if err := j.validateSetSliceDifferentiatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sliceDifferentiator",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference)SetSliceServiceType(val *float64) {
	if err := j.validateSetSliceServiceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sliceServiceType",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) ResetSliceDifferentiator() {
	_jsii_.InvokeVoid(
		m,
		"resetSliceDifferentiator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

