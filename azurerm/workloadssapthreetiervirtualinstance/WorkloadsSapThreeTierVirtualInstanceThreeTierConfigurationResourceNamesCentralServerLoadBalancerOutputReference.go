// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/workloadssapthreetiervirtualinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference interface {
	cdktf.ComplexObject
	BackendPoolNames() *[]*string
	SetBackendPoolNames(val *[]*string)
	BackendPoolNamesInput() *[]*string
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
	FrontendIpConfigurationNames() *[]*string
	SetFrontendIpConfigurationNames(val *[]*string)
	FrontendIpConfigurationNamesInput() *[]*string
	HealthProbeNames() *[]*string
	SetHealthProbeNames(val *[]*string)
	HealthProbeNamesInput() *[]*string
	InternalValue() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancer
	SetInternalValue(val *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancer)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ResetBackendPoolNames()
	ResetFrontendIpConfigurationNames()
	ResetHealthProbeNames()
	ResetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference
type jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) BackendPoolNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backendPoolNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) BackendPoolNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backendPoolNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) FrontendIpConfigurationNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frontendIpConfigurationNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) FrontendIpConfigurationNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frontendIpConfigurationNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) HealthProbeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthProbeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) HealthProbeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthProbeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) InternalValue() *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancer {
	var returns *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference {
	_init_.Initialize()

	if err := validateNewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference_Override(w WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapThreeTierVirtualInstance.WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference)SetBackendPoolNames(val *[]*string) {
	if err := j.validateSetBackendPoolNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendPoolNames",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference)SetFrontendIpConfigurationNames(val *[]*string) {
	if err := j.validateSetFrontendIpConfigurationNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendIpConfigurationNames",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference)SetHealthProbeNames(val *[]*string) {
	if err := j.validateSetHealthProbeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthProbeNames",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference)SetInternalValue(val *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) ResetBackendPoolNames() {
	_jsii_.InvokeVoid(
		w,
		"resetBackendPoolNames",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) ResetFrontendIpConfigurationNames() {
	_jsii_.InvokeVoid(
		w,
		"resetFrontendIpConfigurationNames",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) ResetHealthProbeNames() {
	_jsii_.InvokeVoid(
		w,
		"resetHealthProbeNames",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		w,
		"resetName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

