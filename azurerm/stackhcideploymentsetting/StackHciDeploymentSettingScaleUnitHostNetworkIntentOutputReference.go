// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/stackhcideploymentsetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference interface {
	cdktf.ComplexObject
	Adapter() *[]*string
	SetAdapter(val *[]*string)
	AdapterInput() *[]*string
	AdapterPropertyOverride() StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference
	AdapterPropertyOverrideEnabled() interface{}
	SetAdapterPropertyOverrideEnabled(val interface{})
	AdapterPropertyOverrideEnabledInput() interface{}
	AdapterPropertyOverrideInput() *StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverride
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	QosPolicyOverride() StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference
	QosPolicyOverrideEnabled() interface{}
	SetQosPolicyOverrideEnabled(val interface{})
	QosPolicyOverrideEnabledInput() interface{}
	QosPolicyOverrideInput() *StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrafficType() *[]*string
	SetTrafficType(val *[]*string)
	TrafficTypeInput() *[]*string
	VirtualSwitchConfigurationOverride() StackHciDeploymentSettingScaleUnitHostNetworkIntentVirtualSwitchConfigurationOverrideOutputReference
	VirtualSwitchConfigurationOverrideEnabled() interface{}
	SetVirtualSwitchConfigurationOverrideEnabled(val interface{})
	VirtualSwitchConfigurationOverrideEnabledInput() interface{}
	VirtualSwitchConfigurationOverrideInput() *StackHciDeploymentSettingScaleUnitHostNetworkIntentVirtualSwitchConfigurationOverride
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
	PutAdapterPropertyOverride(value *StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverride)
	PutQosPolicyOverride(value *StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride)
	PutVirtualSwitchConfigurationOverride(value *StackHciDeploymentSettingScaleUnitHostNetworkIntentVirtualSwitchConfigurationOverride)
	ResetAdapterPropertyOverride()
	ResetAdapterPropertyOverrideEnabled()
	ResetQosPolicyOverride()
	ResetQosPolicyOverrideEnabled()
	ResetVirtualSwitchConfigurationOverride()
	ResetVirtualSwitchConfigurationOverrideEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference
type jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) Adapter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adapter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) AdapterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adapterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) AdapterPropertyOverride() StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference {
	var returns StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference
	_jsii_.Get(
		j,
		"adapterPropertyOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) AdapterPropertyOverrideEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adapterPropertyOverrideEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) AdapterPropertyOverrideEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adapterPropertyOverrideEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) AdapterPropertyOverrideInput() *StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverride {
	var returns *StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverride
	_jsii_.Get(
		j,
		"adapterPropertyOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) QosPolicyOverride() StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference {
	var returns StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference
	_jsii_.Get(
		j,
		"qosPolicyOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) QosPolicyOverrideEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"qosPolicyOverrideEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) QosPolicyOverrideEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"qosPolicyOverrideEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) QosPolicyOverrideInput() *StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride {
	var returns *StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride
	_jsii_.Get(
		j,
		"qosPolicyOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) TrafficType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trafficType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) TrafficTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trafficTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) VirtualSwitchConfigurationOverride() StackHciDeploymentSettingScaleUnitHostNetworkIntentVirtualSwitchConfigurationOverrideOutputReference {
	var returns StackHciDeploymentSettingScaleUnitHostNetworkIntentVirtualSwitchConfigurationOverrideOutputReference
	_jsii_.Get(
		j,
		"virtualSwitchConfigurationOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) VirtualSwitchConfigurationOverrideEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualSwitchConfigurationOverrideEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) VirtualSwitchConfigurationOverrideEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualSwitchConfigurationOverrideEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) VirtualSwitchConfigurationOverrideInput() *StackHciDeploymentSettingScaleUnitHostNetworkIntentVirtualSwitchConfigurationOverride {
	var returns *StackHciDeploymentSettingScaleUnitHostNetworkIntentVirtualSwitchConfigurationOverride
	_jsii_.Get(
		j,
		"virtualSwitchConfigurationOverrideInput",
		&returns,
	)
	return returns
}


func NewStackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference {
	_init_.Initialize()

	if err := validateNewStackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciDeploymentSetting.StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference_Override(s StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciDeploymentSetting.StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference)SetAdapter(val *[]*string) {
	if err := j.validateSetAdapterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adapter",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference)SetAdapterPropertyOverrideEnabled(val interface{}) {
	if err := j.validateSetAdapterPropertyOverrideEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adapterPropertyOverrideEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference)SetQosPolicyOverrideEnabled(val interface{}) {
	if err := j.validateSetQosPolicyOverrideEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qosPolicyOverrideEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference)SetTrafficType(val *[]*string) {
	if err := j.validateSetTrafficTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficType",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference)SetVirtualSwitchConfigurationOverrideEnabled(val interface{}) {
	if err := j.validateSetVirtualSwitchConfigurationOverrideEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualSwitchConfigurationOverrideEnabled",
		val,
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) PutAdapterPropertyOverride(value *StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverride) {
	if err := s.validatePutAdapterPropertyOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdapterPropertyOverride",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) PutQosPolicyOverride(value *StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride) {
	if err := s.validatePutQosPolicyOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putQosPolicyOverride",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) PutVirtualSwitchConfigurationOverride(value *StackHciDeploymentSettingScaleUnitHostNetworkIntentVirtualSwitchConfigurationOverride) {
	if err := s.validatePutVirtualSwitchConfigurationOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVirtualSwitchConfigurationOverride",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) ResetAdapterPropertyOverride() {
	_jsii_.InvokeVoid(
		s,
		"resetAdapterPropertyOverride",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) ResetAdapterPropertyOverrideEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetAdapterPropertyOverrideEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) ResetQosPolicyOverride() {
	_jsii_.InvokeVoid(
		s,
		"resetQosPolicyOverride",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) ResetQosPolicyOverrideEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetQosPolicyOverrideEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) ResetVirtualSwitchConfigurationOverride() {
	_jsii_.InvokeVoid(
		s,
		"resetVirtualSwitchConfigurationOverride",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) ResetVirtualSwitchConfigurationOverrideEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetVirtualSwitchConfigurationOverrideEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

