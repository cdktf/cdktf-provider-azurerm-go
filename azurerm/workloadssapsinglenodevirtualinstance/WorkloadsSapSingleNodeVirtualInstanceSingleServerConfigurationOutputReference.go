// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapsinglenodevirtualinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/workloadssapsinglenodevirtualinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference interface {
	cdktf.ComplexObject
	AppResourceGroupName() *string
	SetAppResourceGroupName(val *string)
	AppResourceGroupNameInput() *string
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
	DatabaseType() *string
	SetDatabaseType(val *string)
	DatabaseTypeInput() *string
	DiskVolumeConfiguration() WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationDiskVolumeConfigurationList
	DiskVolumeConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfiguration
	SetInternalValue(val *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfiguration)
	SecondaryIpEnabled() interface{}
	SetSecondaryIpEnabled(val interface{})
	SecondaryIpEnabledInput() interface{}
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualMachineConfiguration() WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfigurationOutputReference
	VirtualMachineConfigurationInput() *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfiguration
	VirtualMachineResourceNames() WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineResourceNamesOutputReference
	VirtualMachineResourceNamesInput() *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineResourceNames
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDiskVolumeConfiguration(value interface{})
	PutVirtualMachineConfiguration(value *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfiguration)
	PutVirtualMachineResourceNames(value *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineResourceNames)
	ResetDatabaseType()
	ResetDiskVolumeConfiguration()
	ResetSecondaryIpEnabled()
	ResetVirtualMachineResourceNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference
type jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) AppResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) AppResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) DatabaseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) DatabaseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) DiskVolumeConfiguration() WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationDiskVolumeConfigurationList {
	var returns WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationDiskVolumeConfigurationList
	_jsii_.Get(
		j,
		"diskVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) DiskVolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) InternalValue() *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfiguration {
	var returns *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) SecondaryIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) SecondaryIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) VirtualMachineConfiguration() WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfigurationOutputReference {
	var returns WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfigurationOutputReference
	_jsii_.Get(
		j,
		"virtualMachineConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) VirtualMachineConfigurationInput() *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfiguration {
	var returns *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfiguration
	_jsii_.Get(
		j,
		"virtualMachineConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) VirtualMachineResourceNames() WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineResourceNamesOutputReference {
	var returns WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineResourceNamesOutputReference
	_jsii_.Get(
		j,
		"virtualMachineResourceNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) VirtualMachineResourceNamesInput() *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineResourceNames {
	var returns *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineResourceNames
	_jsii_.Get(
		j,
		"virtualMachineResourceNamesInput",
		&returns,
	)
	return returns
}


func NewWorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapSingleNodeVirtualInstance.WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference_Override(w WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapSingleNodeVirtualInstance.WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference)SetAppResourceGroupName(val *string) {
	if err := j.validateSetAppResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference)SetDatabaseType(val *string) {
	if err := j.validateSetDatabaseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseType",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference)SetInternalValue(val *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference)SetSecondaryIpEnabled(val interface{}) {
	if err := j.validateSetSecondaryIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryIpEnabled",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) PutDiskVolumeConfiguration(value interface{}) {
	if err := w.validatePutDiskVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDiskVolumeConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) PutVirtualMachineConfiguration(value *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfiguration) {
	if err := w.validatePutVirtualMachineConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putVirtualMachineConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) PutVirtualMachineResourceNames(value *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineResourceNames) {
	if err := w.validatePutVirtualMachineResourceNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putVirtualMachineResourceNames",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) ResetDatabaseType() {
	_jsii_.InvokeVoid(
		w,
		"resetDatabaseType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) ResetDiskVolumeConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetDiskVolumeConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) ResetSecondaryIpEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetSecondaryIpEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) ResetVirtualMachineResourceNames() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualMachineResourceNames",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

