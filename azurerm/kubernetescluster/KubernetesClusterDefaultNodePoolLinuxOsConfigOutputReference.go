// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterDefaultNodePoolLinuxOsConfig
	SetInternalValue(val *KubernetesClusterDefaultNodePoolLinuxOsConfig)
	SwapFileSizeMb() *float64
	SetSwapFileSizeMb(val *float64)
	SwapFileSizeMbInput() *float64
	SysctlConfig() KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference
	SysctlConfigInput() *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransparentHugePageDefrag() *string
	SetTransparentHugePageDefrag(val *string)
	TransparentHugePageDefragInput() *string
	TransparentHugePageEnabled() *string
	SetTransparentHugePageEnabled(val *string)
	TransparentHugePageEnabledInput() *string
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
	PutSysctlConfig(value *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig)
	ResetSwapFileSizeMb()
	ResetSysctlConfig()
	ResetTransparentHugePageDefrag()
	ResetTransparentHugePageEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference
type jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) InternalValue() *KubernetesClusterDefaultNodePoolLinuxOsConfig {
	var returns *KubernetesClusterDefaultNodePoolLinuxOsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SwapFileSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapFileSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SwapFileSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapFileSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SysctlConfig() KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference {
	var returns KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference
	_jsii_.Get(
		j,
		"sysctlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SysctlConfigInput() *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig {
	var returns *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig
	_jsii_.Get(
		j,
		"sysctlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TransparentHugePageDefrag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageDefrag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TransparentHugePageDefragInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageDefragInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TransparentHugePageEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TransparentHugePageEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageEnabledInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterDefaultNodePoolLinuxOsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference_Override(k KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference)SetInternalValue(val *KubernetesClusterDefaultNodePoolLinuxOsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference)SetSwapFileSizeMb(val *float64) {
	if err := j.validateSetSwapFileSizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"swapFileSizeMb",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference)SetTransparentHugePageDefrag(val *string) {
	if err := j.validateSetTransparentHugePageDefragParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentHugePageDefrag",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference)SetTransparentHugePageEnabled(val *string) {
	if err := j.validateSetTransparentHugePageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentHugePageEnabled",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) PutSysctlConfig(value *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig) {
	if err := k.validatePutSysctlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSysctlConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ResetSwapFileSizeMb() {
	_jsii_.InvokeVoid(
		k,
		"resetSwapFileSizeMb",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ResetSysctlConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetSysctlConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ResetTransparentHugePageDefrag() {
	_jsii_.InvokeVoid(
		k,
		"resetTransparentHugePageDefrag",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ResetTransparentHugePageEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetTransparentHugePageEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

