// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/workloadssapthreetiervirtualinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference interface {
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
	DataDisk() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineDataDiskList
	DataDiskInput() interface{}
	// Experimental.
	Fqn() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NetworkInterfaceNames() *[]*string
	SetNetworkInterfaceNames(val *[]*string)
	NetworkInterfaceNamesInput() *[]*string
	OsDiskName() *string
	SetOsDiskName(val *string)
	OsDiskNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualMachineName() *string
	SetVirtualMachineName(val *string)
	VirtualMachineNameInput() *string
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
	PutDataDisk(value interface{})
	ResetDataDisk()
	ResetHostName()
	ResetNetworkInterfaceNames()
	ResetOsDiskName()
	ResetVirtualMachineName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference
type jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) DataDisk() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineDataDiskList {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineDataDiskList
	_jsii_.Get(
		j,
		"dataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) DataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) NetworkInterfaceNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) NetworkInterfaceNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) OsDiskName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) OsDiskNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) VirtualMachineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) VirtualMachineNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineNameInput",
		&returns,
	)
	return returns
}


func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference {
	_init_.Initialize()

	if err := validateNewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference_Override(w WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference)SetHostName(val *string) {
	if err := j.validateSetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference)SetNetworkInterfaceNames(val *[]*string) {
	if err := j.validateSetNetworkInterfaceNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceNames",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference)SetOsDiskName(val *string) {
	if err := j.validateSetOsDiskNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskName",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference)SetVirtualMachineName(val *string) {
	if err := j.validateSetVirtualMachineNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualMachineName",
		val,
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) PutDataDisk(value interface{}) {
	if err := w.validatePutDataDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDataDisk",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) ResetDataDisk() {
	_jsii_.InvokeVoid(
		w,
		"resetDataDisk",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) ResetHostName() {
	_jsii_.InvokeVoid(
		w,
		"resetHostName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) ResetNetworkInterfaceNames() {
	_jsii_.InvokeVoid(
		w,
		"resetNetworkInterfaceNames",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) ResetOsDiskName() {
	_jsii_.InvokeVoid(
		w,
		"resetOsDiskName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) ResetVirtualMachineName() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualMachineName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServerVirtualMachineOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

