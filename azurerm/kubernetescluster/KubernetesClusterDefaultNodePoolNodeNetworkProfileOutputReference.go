// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference interface {
	cdktf.ComplexObject
	AllowedHostPorts() KubernetesClusterDefaultNodePoolNodeNetworkProfileAllowedHostPortsList
	AllowedHostPortsInput() interface{}
	ApplicationSecurityGroupIds() *[]*string
	SetApplicationSecurityGroupIds(val *[]*string)
	ApplicationSecurityGroupIdsInput() *[]*string
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
	InternalValue() *KubernetesClusterDefaultNodePoolNodeNetworkProfile
	SetInternalValue(val *KubernetesClusterDefaultNodePoolNodeNetworkProfile)
	NodePublicIpTags() *map[string]*string
	SetNodePublicIpTags(val *map[string]*string)
	NodePublicIpTagsInput() *map[string]*string
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
	PutAllowedHostPorts(value interface{})
	ResetAllowedHostPorts()
	ResetApplicationSecurityGroupIds()
	ResetNodePublicIpTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference
type jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) AllowedHostPorts() KubernetesClusterDefaultNodePoolNodeNetworkProfileAllowedHostPortsList {
	var returns KubernetesClusterDefaultNodePoolNodeNetworkProfileAllowedHostPortsList
	_jsii_.Get(
		j,
		"allowedHostPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) AllowedHostPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedHostPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) ApplicationSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) ApplicationSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) InternalValue() *KubernetesClusterDefaultNodePoolNodeNetworkProfile {
	var returns *KubernetesClusterDefaultNodePoolNodeNetworkProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) NodePublicIpTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodePublicIpTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) NodePublicIpTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodePublicIpTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference_Override(k KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference)SetApplicationSecurityGroupIds(val *[]*string) {
	if err := j.validateSetApplicationSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference)SetInternalValue(val *KubernetesClusterDefaultNodePoolNodeNetworkProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference)SetNodePublicIpTags(val *map[string]*string) {
	if err := j.validateSetNodePublicIpTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodePublicIpTags",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) PutAllowedHostPorts(value interface{}) {
	if err := k.validatePutAllowedHostPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAllowedHostPorts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) ResetAllowedHostPorts() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowedHostPorts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) ResetApplicationSecurityGroupIds() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationSecurityGroupIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) ResetNodePublicIpTags() {
	_jsii_.InvokeVoid(
		k,
		"resetNodePublicIpTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolNodeNetworkProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

