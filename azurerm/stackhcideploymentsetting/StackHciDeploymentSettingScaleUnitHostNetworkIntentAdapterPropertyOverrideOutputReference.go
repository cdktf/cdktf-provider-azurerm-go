// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/stackhcideploymentsetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference interface {
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
	InternalValue() *StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverride
	SetInternalValue(val *StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverride)
	JumboPacket() *string
	SetJumboPacket(val *string)
	JumboPacketInput() *string
	NetworkDirect() *string
	SetNetworkDirect(val *string)
	NetworkDirectInput() *string
	NetworkDirectTechnology() *string
	SetNetworkDirectTechnology(val *string)
	NetworkDirectTechnologyInput() *string
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
	ResetJumboPacket()
	ResetNetworkDirect()
	ResetNetworkDirectTechnology()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference
type jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) InternalValue() *StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverride {
	var returns *StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverride
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) JumboPacket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jumboPacket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) JumboPacketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jumboPacketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) NetworkDirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkDirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) NetworkDirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkDirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) NetworkDirectTechnology() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkDirectTechnology",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) NetworkDirectTechnologyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkDirectTechnologyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference {
	_init_.Initialize()

	if err := validateNewStackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciDeploymentSetting.StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference_Override(s StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciDeploymentSetting.StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference)SetInternalValue(val *StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverride) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference)SetJumboPacket(val *string) {
	if err := j.validateSetJumboPacketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jumboPacket",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference)SetNetworkDirect(val *string) {
	if err := j.validateSetNetworkDirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkDirect",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference)SetNetworkDirectTechnology(val *string) {
	if err := j.validateSetNetworkDirectTechnologyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkDirectTechnology",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) ResetJumboPacket() {
	_jsii_.InvokeVoid(
		s,
		"resetJumboPacket",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) ResetNetworkDirect() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDirect",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) ResetNetworkDirectTechnology() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkDirectTechnology",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentAdapterPropertyOverrideOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

