// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworksimpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/mobilenetworksimpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MobileNetworkSimPolicySliceDataNetworkOutputReference interface {
	cdktf.ComplexObject
	AdditionalAllowedSessionTypes() *[]*string
	SetAdditionalAllowedSessionTypes(val *[]*string)
	AdditionalAllowedSessionTypesInput() *[]*string
	AllocationAndRetentionPriorityLevel() *float64
	SetAllocationAndRetentionPriorityLevel(val *float64)
	AllocationAndRetentionPriorityLevelInput() *float64
	AllowedServicesIds() *[]*string
	SetAllowedServicesIds(val *[]*string)
	AllowedServicesIdsInput() *[]*string
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
	DataNetworkId() *string
	SetDataNetworkId(val *string)
	DataNetworkIdInput() *string
	DefaultSessionType() *string
	SetDefaultSessionType(val *string)
	DefaultSessionTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxBufferedPackets() *float64
	SetMaxBufferedPackets(val *float64)
	MaxBufferedPacketsInput() *float64
	PreemptionCapability() *string
	SetPreemptionCapability(val *string)
	PreemptionCapabilityInput() *string
	PreemptionVulnerability() *string
	SetPreemptionVulnerability(val *string)
	PreemptionVulnerabilityInput() *string
	QosIndicator() *float64
	SetQosIndicator(val *float64)
	QosIndicatorInput() *float64
	SessionAggregateMaximumBitRate() MobileNetworkSimPolicySliceDataNetworkSessionAggregateMaximumBitRateOutputReference
	SessionAggregateMaximumBitRateInput() *MobileNetworkSimPolicySliceDataNetworkSessionAggregateMaximumBitRate
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
	PutSessionAggregateMaximumBitRate(value *MobileNetworkSimPolicySliceDataNetworkSessionAggregateMaximumBitRate)
	ResetAdditionalAllowedSessionTypes()
	ResetAllocationAndRetentionPriorityLevel()
	ResetDefaultSessionType()
	ResetMaxBufferedPackets()
	ResetPreemptionCapability()
	ResetPreemptionVulnerability()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MobileNetworkSimPolicySliceDataNetworkOutputReference
type jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) AdditionalAllowedSessionTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalAllowedSessionTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) AdditionalAllowedSessionTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalAllowedSessionTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) AllocationAndRetentionPriorityLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationAndRetentionPriorityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) AllocationAndRetentionPriorityLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationAndRetentionPriorityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) AllowedServicesIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedServicesIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) AllowedServicesIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedServicesIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) DataNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) DataNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) DefaultSessionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSessionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) DefaultSessionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSessionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) MaxBufferedPackets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBufferedPackets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) MaxBufferedPacketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBufferedPacketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) PreemptionCapability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptionCapability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) PreemptionCapabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptionCapabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) PreemptionVulnerability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptionVulnerability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) PreemptionVulnerabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptionVulnerabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) QosIndicator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qosIndicator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) QosIndicatorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qosIndicatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) SessionAggregateMaximumBitRate() MobileNetworkSimPolicySliceDataNetworkSessionAggregateMaximumBitRateOutputReference {
	var returns MobileNetworkSimPolicySliceDataNetworkSessionAggregateMaximumBitRateOutputReference
	_jsii_.Get(
		j,
		"sessionAggregateMaximumBitRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) SessionAggregateMaximumBitRateInput() *MobileNetworkSimPolicySliceDataNetworkSessionAggregateMaximumBitRate {
	var returns *MobileNetworkSimPolicySliceDataNetworkSessionAggregateMaximumBitRate
	_jsii_.Get(
		j,
		"sessionAggregateMaximumBitRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMobileNetworkSimPolicySliceDataNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MobileNetworkSimPolicySliceDataNetworkOutputReference {
	_init_.Initialize()

	if err := validateNewMobileNetworkSimPolicySliceDataNetworkOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkSimPolicy.MobileNetworkSimPolicySliceDataNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMobileNetworkSimPolicySliceDataNetworkOutputReference_Override(m MobileNetworkSimPolicySliceDataNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkSimPolicy.MobileNetworkSimPolicySliceDataNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetAdditionalAllowedSessionTypes(val *[]*string) {
	if err := j.validateSetAdditionalAllowedSessionTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalAllowedSessionTypes",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetAllocationAndRetentionPriorityLevel(val *float64) {
	if err := j.validateSetAllocationAndRetentionPriorityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationAndRetentionPriorityLevel",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetAllowedServicesIds(val *[]*string) {
	if err := j.validateSetAllowedServicesIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedServicesIds",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetDataNetworkId(val *string) {
	if err := j.validateSetDataNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataNetworkId",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetDefaultSessionType(val *string) {
	if err := j.validateSetDefaultSessionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSessionType",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetMaxBufferedPackets(val *float64) {
	if err := j.validateSetMaxBufferedPacketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBufferedPackets",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetPreemptionCapability(val *string) {
	if err := j.validateSetPreemptionCapabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptionCapability",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetPreemptionVulnerability(val *string) {
	if err := j.validateSetPreemptionVulnerabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptionVulnerability",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetQosIndicator(val *float64) {
	if err := j.validateSetQosIndicatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qosIndicator",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) PutSessionAggregateMaximumBitRate(value *MobileNetworkSimPolicySliceDataNetworkSessionAggregateMaximumBitRate) {
	if err := m.validatePutSessionAggregateMaximumBitRateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSessionAggregateMaximumBitRate",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) ResetAdditionalAllowedSessionTypes() {
	_jsii_.InvokeVoid(
		m,
		"resetAdditionalAllowedSessionTypes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) ResetAllocationAndRetentionPriorityLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetAllocationAndRetentionPriorityLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) ResetDefaultSessionType() {
	_jsii_.InvokeVoid(
		m,
		"resetDefaultSessionType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) ResetMaxBufferedPackets() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBufferedPackets",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) ResetPreemptionCapability() {
	_jsii_.InvokeVoid(
		m,
		"resetPreemptionCapability",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) ResetPreemptionVulnerability() {
	_jsii_.InvokeVoid(
		m,
		"resetPreemptionVulnerability",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MobileNetworkSimPolicySliceDataNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

