// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference interface {
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
	EffectiveOutboundIps() *[]*string
	// Experimental.
	Fqn() *string
	IdleTimeoutInMinutes() *float64
	SetIdleTimeoutInMinutes(val *float64)
	IdleTimeoutInMinutesInput() *float64
	InternalValue() *KubernetesClusterNetworkProfileLoadBalancerProfile
	SetInternalValue(val *KubernetesClusterNetworkProfileLoadBalancerProfile)
	ManagedOutboundIpCount() *float64
	SetManagedOutboundIpCount(val *float64)
	ManagedOutboundIpCountInput() *float64
	ManagedOutboundIpv6Count() *float64
	SetManagedOutboundIpv6Count(val *float64)
	ManagedOutboundIpv6CountInput() *float64
	OutboundIpAddressIds() *[]*string
	SetOutboundIpAddressIds(val *[]*string)
	OutboundIpAddressIdsInput() *[]*string
	OutboundIpPrefixIds() *[]*string
	SetOutboundIpPrefixIds(val *[]*string)
	OutboundIpPrefixIdsInput() *[]*string
	OutboundPortsAllocated() *float64
	SetOutboundPortsAllocated(val *float64)
	OutboundPortsAllocatedInput() *float64
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
	ResetIdleTimeoutInMinutes()
	ResetManagedOutboundIpCount()
	ResetManagedOutboundIpv6Count()
	ResetOutboundIpAddressIds()
	ResetOutboundIpPrefixIds()
	ResetOutboundPortsAllocated()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference
type jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) EffectiveOutboundIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectiveOutboundIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) IdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) IdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) InternalValue() *KubernetesClusterNetworkProfileLoadBalancerProfile {
	var returns *KubernetesClusterNetworkProfileLoadBalancerProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ManagedOutboundIpCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedOutboundIpCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ManagedOutboundIpCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedOutboundIpCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ManagedOutboundIpv6Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedOutboundIpv6Count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ManagedOutboundIpv6CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedOutboundIpv6CountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundIpAddressIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundIpAddressIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundIpPrefixIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpPrefixIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundIpPrefixIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpPrefixIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundPortsAllocated() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outboundPortsAllocated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundPortsAllocatedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outboundPortsAllocatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterNetworkProfileLoadBalancerProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterNetworkProfileLoadBalancerProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterNetworkProfileLoadBalancerProfileOutputReference_Override(k KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference)SetIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference)SetInternalValue(val *KubernetesClusterNetworkProfileLoadBalancerProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference)SetManagedOutboundIpCount(val *float64) {
	if err := j.validateSetManagedOutboundIpCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedOutboundIpCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference)SetManagedOutboundIpv6Count(val *float64) {
	if err := j.validateSetManagedOutboundIpv6CountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedOutboundIpv6Count",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference)SetOutboundIpAddressIds(val *[]*string) {
	if err := j.validateSetOutboundIpAddressIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundIpAddressIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference)SetOutboundIpPrefixIds(val *[]*string) {
	if err := j.validateSetOutboundIpPrefixIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundIpPrefixIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference)SetOutboundPortsAllocated(val *float64) {
	if err := j.validateSetOutboundPortsAllocatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundPortsAllocated",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ResetIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		k,
		"resetIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ResetManagedOutboundIpCount() {
	_jsii_.InvokeVoid(
		k,
		"resetManagedOutboundIpCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ResetManagedOutboundIpv6Count() {
	_jsii_.InvokeVoid(
		k,
		"resetManagedOutboundIpv6Count",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ResetOutboundIpAddressIds() {
	_jsii_.InvokeVoid(
		k,
		"resetOutboundIpAddressIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ResetOutboundIpPrefixIds() {
	_jsii_.InvokeVoid(
		k,
		"resetOutboundIpPrefixIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ResetOutboundPortsAllocated() {
	_jsii_.InvokeVoid(
		k,
		"resetOutboundPortsAllocated",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

