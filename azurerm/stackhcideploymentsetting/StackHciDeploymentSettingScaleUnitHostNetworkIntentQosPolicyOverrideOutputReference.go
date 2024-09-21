// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/stackhcideploymentsetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference interface {
	cdktf.ComplexObject
	BandwidthPercentageSmb() *string
	SetBandwidthPercentageSmb(val *string)
	BandwidthPercentageSmbInput() *string
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
	InternalValue() *StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride
	SetInternalValue(val *StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride)
	PriorityValue8021ActionCluster() *string
	SetPriorityValue8021ActionCluster(val *string)
	PriorityValue8021ActionClusterInput() *string
	PriorityValue8021ActionSmb() *string
	SetPriorityValue8021ActionSmb(val *string)
	PriorityValue8021ActionSmbInput() *string
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
	ResetBandwidthPercentageSmb()
	ResetPriorityValue8021ActionCluster()
	ResetPriorityValue8021ActionSmb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference
type jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) BandwidthPercentageSmb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthPercentageSmb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) BandwidthPercentageSmbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthPercentageSmbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) InternalValue() *StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride {
	var returns *StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) PriorityValue8021ActionCluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityValue8021ActionCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) PriorityValue8021ActionClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityValue8021ActionClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) PriorityValue8021ActionSmb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityValue8021ActionSmb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) PriorityValue8021ActionSmbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityValue8021ActionSmbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference {
	_init_.Initialize()

	if err := validateNewStackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciDeploymentSetting.StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference_Override(s StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciDeploymentSetting.StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference)SetBandwidthPercentageSmb(val *string) {
	if err := j.validateSetBandwidthPercentageSmbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthPercentageSmb",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference)SetInternalValue(val *StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference)SetPriorityValue8021ActionCluster(val *string) {
	if err := j.validateSetPriorityValue8021ActionClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priorityValue8021ActionCluster",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference)SetPriorityValue8021ActionSmb(val *string) {
	if err := j.validateSetPriorityValue8021ActionSmbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priorityValue8021ActionSmb",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) ResetBandwidthPercentageSmb() {
	_jsii_.InvokeVoid(
		s,
		"resetBandwidthPercentageSmb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) ResetPriorityValue8021ActionCluster() {
	_jsii_.InvokeVoid(
		s,
		"resetPriorityValue8021ActionCluster",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) ResetPriorityValue8021ActionSmb() {
	_jsii_.InvokeVoid(
		s,
		"resetPriorityValue8021ActionSmb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverrideOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

