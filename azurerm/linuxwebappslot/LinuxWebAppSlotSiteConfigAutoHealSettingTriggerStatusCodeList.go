// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/linuxwebappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList
type jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewLinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList {
	_init_.Initialize()

	if err := validateNewLinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxWebAppSlot.LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewLinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList_Override(l LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxWebAppSlot.LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		l,
	)
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList) Get(index *float64) LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeOutputReference {
	if err := l.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeOutputReference

	_jsii_.Invoke(
		l,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCodeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

