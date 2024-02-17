// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/orchestratedvirtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference interface {
	cdktf.ComplexObject
	AdditionalUnattendContent() OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationAdditionalUnattendContentList
	AdditionalUnattendContentInput() interface{}
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
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
	ComputerNamePrefix() *string
	SetComputerNamePrefix(val *string)
	ComputerNamePrefixInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableAutomaticUpdates() interface{}
	SetEnableAutomaticUpdates(val interface{})
	EnableAutomaticUpdatesInput() interface{}
	// Experimental.
	Fqn() *string
	HotpatchingEnabled() interface{}
	SetHotpatchingEnabled(val interface{})
	HotpatchingEnabledInput() interface{}
	InternalValue() *OrchestratedVirtualMachineScaleSetOsProfileWindowsConfiguration
	SetInternalValue(val *OrchestratedVirtualMachineScaleSetOsProfileWindowsConfiguration)
	PatchAssessmentMode() *string
	SetPatchAssessmentMode(val *string)
	PatchAssessmentModeInput() *string
	PatchMode() *string
	SetPatchMode(val *string)
	PatchModeInput() *string
	ProvisionVmAgent() interface{}
	SetProvisionVmAgent(val interface{})
	ProvisionVmAgentInput() interface{}
	Secret() OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretList
	SecretInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	WinrmListener() OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListenerList
	WinrmListenerInput() interface{}
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
	PutAdditionalUnattendContent(value interface{})
	PutSecret(value interface{})
	PutWinrmListener(value interface{})
	ResetAdditionalUnattendContent()
	ResetComputerNamePrefix()
	ResetEnableAutomaticUpdates()
	ResetHotpatchingEnabled()
	ResetPatchAssessmentMode()
	ResetPatchMode()
	ResetProvisionVmAgent()
	ResetSecret()
	ResetTimezone()
	ResetWinrmListener()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference
type jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) AdditionalUnattendContent() OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationAdditionalUnattendContentList {
	var returns OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationAdditionalUnattendContentList
	_jsii_.Get(
		j,
		"additionalUnattendContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) AdditionalUnattendContentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalUnattendContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ComputerNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ComputerNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) EnableAutomaticUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) EnableAutomaticUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) HotpatchingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hotpatchingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) HotpatchingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hotpatchingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) InternalValue() *OrchestratedVirtualMachineScaleSetOsProfileWindowsConfiguration {
	var returns *OrchestratedVirtualMachineScaleSetOsProfileWindowsConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) PatchAssessmentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchAssessmentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) PatchAssessmentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchAssessmentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) PatchMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) PatchModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ProvisionVmAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ProvisionVmAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) Secret() OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretList {
	var returns OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) WinrmListener() OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListenerList {
	var returns OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListenerList
	_jsii_.Get(
		j,
		"winrmListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) WinrmListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmListenerInput",
		&returns,
	)
	return returns
}


func NewOrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewOrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference_Override(o OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetAdminUsername(val *string) {
	if err := j.validateSetAdminUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetComputerNamePrefix(val *string) {
	if err := j.validateSetComputerNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computerNamePrefix",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetEnableAutomaticUpdates(val interface{}) {
	if err := j.validateSetEnableAutomaticUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutomaticUpdates",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetHotpatchingEnabled(val interface{}) {
	if err := j.validateSetHotpatchingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hotpatchingEnabled",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetInternalValue(val *OrchestratedVirtualMachineScaleSetOsProfileWindowsConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetPatchAssessmentMode(val *string) {
	if err := j.validateSetPatchAssessmentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchAssessmentMode",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetPatchMode(val *string) {
	if err := j.validateSetPatchModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchMode",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetProvisionVmAgent(val interface{}) {
	if err := j.validateSetProvisionVmAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionVmAgent",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) PutAdditionalUnattendContent(value interface{}) {
	if err := o.validatePutAdditionalUnattendContentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAdditionalUnattendContent",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) PutSecret(value interface{}) {
	if err := o.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSecret",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) PutWinrmListener(value interface{}) {
	if err := o.validatePutWinrmListenerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWinrmListener",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ResetAdditionalUnattendContent() {
	_jsii_.InvokeVoid(
		o,
		"resetAdditionalUnattendContent",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ResetComputerNamePrefix() {
	_jsii_.InvokeVoid(
		o,
		"resetComputerNamePrefix",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ResetEnableAutomaticUpdates() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableAutomaticUpdates",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ResetHotpatchingEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetHotpatchingEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ResetPatchAssessmentMode() {
	_jsii_.InvokeVoid(
		o,
		"resetPatchAssessmentMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ResetPatchMode() {
	_jsii_.InvokeVoid(
		o,
		"resetPatchMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ResetProvisionVmAgent() {
	_jsii_.InvokeVoid(
		o,
		"resetProvisionVmAgent",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		o,
		"resetSecret",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		o,
		"resetTimezone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ResetWinrmListener() {
	_jsii_.InvokeVoid(
		o,
		"resetWinrmListener",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

