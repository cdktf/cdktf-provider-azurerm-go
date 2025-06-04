// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/orchestratedvirtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference interface {
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
	CrossZoneUpgradesEnabled() interface{}
	SetCrossZoneUpgradesEnabled(val interface{})
	CrossZoneUpgradesEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *OrchestratedVirtualMachineScaleSetRollingUpgradePolicy
	SetInternalValue(val *OrchestratedVirtualMachineScaleSetRollingUpgradePolicy)
	MaxBatchInstancePercent() *float64
	SetMaxBatchInstancePercent(val *float64)
	MaxBatchInstancePercentInput() *float64
	MaximumSurgeInstancesEnabled() interface{}
	SetMaximumSurgeInstancesEnabled(val interface{})
	MaximumSurgeInstancesEnabledInput() interface{}
	MaxUnhealthyInstancePercent() *float64
	SetMaxUnhealthyInstancePercent(val *float64)
	MaxUnhealthyInstancePercentInput() *float64
	MaxUnhealthyUpgradedInstancePercent() *float64
	SetMaxUnhealthyUpgradedInstancePercent(val *float64)
	MaxUnhealthyUpgradedInstancePercentInput() *float64
	PauseTimeBetweenBatches() *string
	SetPauseTimeBetweenBatches(val *string)
	PauseTimeBetweenBatchesInput() *string
	PrioritizeUnhealthyInstancesEnabled() interface{}
	SetPrioritizeUnhealthyInstancesEnabled(val interface{})
	PrioritizeUnhealthyInstancesEnabledInput() interface{}
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
	ResetCrossZoneUpgradesEnabled()
	ResetMaximumSurgeInstancesEnabled()
	ResetPrioritizeUnhealthyInstancesEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference
type jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) CrossZoneUpgradesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossZoneUpgradesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) CrossZoneUpgradesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossZoneUpgradesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) InternalValue() *OrchestratedVirtualMachineScaleSetRollingUpgradePolicy {
	var returns *OrchestratedVirtualMachineScaleSetRollingUpgradePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxBatchInstancePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchInstancePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxBatchInstancePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchInstancePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaximumSurgeInstancesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maximumSurgeInstancesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaximumSurgeInstancesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maximumSurgeInstancesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyInstancePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyInstancePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyInstancePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyInstancePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyUpgradedInstancePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyUpgradedInstancePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyUpgradedInstancePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyUpgradedInstancePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PauseTimeBetweenBatches() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseTimeBetweenBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PauseTimeBetweenBatchesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseTimeBetweenBatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PrioritizeUnhealthyInstancesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prioritizeUnhealthyInstancesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PrioritizeUnhealthyInstancesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prioritizeUnhealthyInstancesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewOrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference_Override(o OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetCrossZoneUpgradesEnabled(val interface{}) {
	if err := j.validateSetCrossZoneUpgradesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossZoneUpgradesEnabled",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetInternalValue(val *OrchestratedVirtualMachineScaleSetRollingUpgradePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetMaxBatchInstancePercent(val *float64) {
	if err := j.validateSetMaxBatchInstancePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBatchInstancePercent",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetMaximumSurgeInstancesEnabled(val interface{}) {
	if err := j.validateSetMaximumSurgeInstancesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumSurgeInstancesEnabled",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetMaxUnhealthyInstancePercent(val *float64) {
	if err := j.validateSetMaxUnhealthyInstancePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnhealthyInstancePercent",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetMaxUnhealthyUpgradedInstancePercent(val *float64) {
	if err := j.validateSetMaxUnhealthyUpgradedInstancePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnhealthyUpgradedInstancePercent",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetPauseTimeBetweenBatches(val *string) {
	if err := j.validateSetPauseTimeBetweenBatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pauseTimeBetweenBatches",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetPrioritizeUnhealthyInstancesEnabled(val interface{}) {
	if err := j.validateSetPrioritizeUnhealthyInstancesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prioritizeUnhealthyInstancesEnabled",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ResetCrossZoneUpgradesEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetCrossZoneUpgradesEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ResetMaximumSurgeInstancesEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetMaximumSurgeInstancesEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ResetPrioritizeUnhealthyInstancesEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetPrioritizeUnhealthyInstancesEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

