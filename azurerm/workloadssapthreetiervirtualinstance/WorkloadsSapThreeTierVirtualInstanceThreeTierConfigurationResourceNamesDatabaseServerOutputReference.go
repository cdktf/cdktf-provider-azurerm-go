// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/workloadssapthreetiervirtualinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference interface {
	cdktf.ComplexObject
	AvailabilitySetName() *string
	SetAvailabilitySetName(val *string)
	AvailabilitySetNameInput() *string
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
	InternalValue() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServer
	SetInternalValue(val *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServer)
	LoadBalancer() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerLoadBalancerOutputReference
	LoadBalancerInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerLoadBalancer
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualMachine() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerVirtualMachineList
	VirtualMachineInput() interface{}
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
	PutLoadBalancer(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerLoadBalancer)
	PutVirtualMachine(value interface{})
	ResetAvailabilitySetName()
	ResetLoadBalancer()
	ResetVirtualMachine()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference
type jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) AvailabilitySetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilitySetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) AvailabilitySetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilitySetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) InternalValue() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServer {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) LoadBalancer() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerLoadBalancerOutputReference {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) LoadBalancerInput() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerLoadBalancer {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) VirtualMachine() WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerVirtualMachineList {
	var returns WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerVirtualMachineList
	_jsii_.Get(
		j,
		"virtualMachine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) VirtualMachineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualMachineInput",
		&returns,
	)
	return returns
}


func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference {
	_init_.Initialize()

	if err := validateNewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference_Override(w WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference)SetAvailabilitySetName(val *string) {
	if err := j.validateSetAvailabilitySetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilitySetName",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference)SetInternalValue(val *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) PutLoadBalancer(value *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerLoadBalancer) {
	if err := w.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) PutVirtualMachine(value interface{}) {
	if err := w.validatePutVirtualMachineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putVirtualMachine",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) ResetAvailabilitySetName() {
	_jsii_.InvokeVoid(
		w,
		"resetAvailabilitySetName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		w,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) ResetVirtualMachine() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualMachine",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

