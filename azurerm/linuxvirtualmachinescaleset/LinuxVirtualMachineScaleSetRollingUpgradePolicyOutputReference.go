// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/linuxvirtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference interface {
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
	InternalValue() *LinuxVirtualMachineScaleSetRollingUpgradePolicy
	SetInternalValue(val *LinuxVirtualMachineScaleSetRollingUpgradePolicy)
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

// The jsii proxy struct for LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference
type jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) CrossZoneUpgradesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossZoneUpgradesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) CrossZoneUpgradesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossZoneUpgradesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) InternalValue() *LinuxVirtualMachineScaleSetRollingUpgradePolicy {
	var returns *LinuxVirtualMachineScaleSetRollingUpgradePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxBatchInstancePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchInstancePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxBatchInstancePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchInstancePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaximumSurgeInstancesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maximumSurgeInstancesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaximumSurgeInstancesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maximumSurgeInstancesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyInstancePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyInstancePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyInstancePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyInstancePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyUpgradedInstancePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyUpgradedInstancePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) MaxUnhealthyUpgradedInstancePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyUpgradedInstancePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PauseTimeBetweenBatches() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseTimeBetweenBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PauseTimeBetweenBatchesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseTimeBetweenBatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PrioritizeUnhealthyInstancesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prioritizeUnhealthyInstancesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) PrioritizeUnhealthyInstancesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prioritizeUnhealthyInstancesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewLinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxVirtualMachineScaleSet.LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference_Override(l LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxVirtualMachineScaleSet.LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetCrossZoneUpgradesEnabled(val interface{}) {
	if err := j.validateSetCrossZoneUpgradesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossZoneUpgradesEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetInternalValue(val *LinuxVirtualMachineScaleSetRollingUpgradePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetMaxBatchInstancePercent(val *float64) {
	if err := j.validateSetMaxBatchInstancePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBatchInstancePercent",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetMaximumSurgeInstancesEnabled(val interface{}) {
	if err := j.validateSetMaximumSurgeInstancesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumSurgeInstancesEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetMaxUnhealthyInstancePercent(val *float64) {
	if err := j.validateSetMaxUnhealthyInstancePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnhealthyInstancePercent",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetMaxUnhealthyUpgradedInstancePercent(val *float64) {
	if err := j.validateSetMaxUnhealthyUpgradedInstancePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnhealthyUpgradedInstancePercent",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetPauseTimeBetweenBatches(val *string) {
	if err := j.validateSetPauseTimeBetweenBatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pauseTimeBetweenBatches",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetPrioritizeUnhealthyInstancesEnabled(val interface{}) {
	if err := j.validateSetPrioritizeUnhealthyInstancesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prioritizeUnhealthyInstancesEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ResetCrossZoneUpgradesEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetCrossZoneUpgradesEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ResetMaximumSurgeInstancesEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetMaximumSurgeInstancesEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ResetPrioritizeUnhealthyInstancesEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetPrioritizeUnhealthyInstancesEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

