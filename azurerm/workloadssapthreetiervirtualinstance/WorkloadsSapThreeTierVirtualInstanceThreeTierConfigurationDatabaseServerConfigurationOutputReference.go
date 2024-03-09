// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/workloadssapthreetiervirtualinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference interface {
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
	DatabaseType() *string
	SetDatabaseType(val *string)
	DatabaseTypeInput() *string
	DiskVolumeConfiguration() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationDiskVolumeConfigurationList
	DiskVolumeConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InternalValue() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfiguration
	SetInternalValue(val *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfiguration)
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
	VirtualMachineConfiguration() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationVirtualMachineConfigurationOutputReference
	VirtualMachineConfigurationInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationVirtualMachineConfiguration
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
	PutDiskVolumeConfiguration(value interface{})
	PutVirtualMachineConfiguration(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationVirtualMachineConfiguration)
	ResetDatabaseType()
	ResetDiskVolumeConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference
type jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) DatabaseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) DatabaseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) DiskVolumeConfiguration() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationDiskVolumeConfigurationList {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationDiskVolumeConfigurationList
	_jsii_.Get(
		j,
		"diskVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) DiskVolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) InternalValue() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfiguration {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) VirtualMachineConfiguration() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationVirtualMachineConfigurationOutputReference {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationVirtualMachineConfigurationOutputReference
	_jsii_.Get(
		j,
		"virtualMachineConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) VirtualMachineConfigurationInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationVirtualMachineConfiguration {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationVirtualMachineConfiguration
	_jsii_.Get(
		j,
		"virtualMachineConfigurationInput",
		&returns,
	)
	return returns
}


func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference_Override(w WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference)SetDatabaseType(val *string) {
	if err := j.validateSetDatabaseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseType",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference)SetInternalValue(val *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) PutDiskVolumeConfiguration(value interface{}) {
	if err := w.validatePutDiskVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDiskVolumeConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) PutVirtualMachineConfiguration(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationVirtualMachineConfiguration) {
	if err := w.validatePutVirtualMachineConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putVirtualMachineConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) ResetDatabaseType() {
	_jsii_.InvokeVoid(
		w,
		"resetDatabaseType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) ResetDiskVolumeConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetDiskVolumeConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

