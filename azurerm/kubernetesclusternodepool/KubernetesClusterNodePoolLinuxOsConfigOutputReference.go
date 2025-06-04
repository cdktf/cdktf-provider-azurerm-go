// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusternodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/kubernetesclusternodepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterNodePoolLinuxOsConfigOutputReference interface {
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
	InternalValue() *KubernetesClusterNodePoolLinuxOsConfig
	SetInternalValue(val *KubernetesClusterNodePoolLinuxOsConfig)
	SwapFileSizeMb() *float64
	SetSwapFileSizeMb(val *float64)
	SwapFileSizeMbInput() *float64
	SysctlConfig() KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference
	SysctlConfigInput() *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig
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
	PutSysctlConfig(value *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig)
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

// The jsii proxy struct for KubernetesClusterNodePoolLinuxOsConfigOutputReference
type jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) InternalValue() *KubernetesClusterNodePoolLinuxOsConfig {
	var returns *KubernetesClusterNodePoolLinuxOsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SwapFileSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapFileSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SwapFileSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapFileSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SysctlConfig() KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference {
	var returns KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference
	_jsii_.Get(
		j,
		"sysctlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SysctlConfigInput() *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig {
	var returns *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig
	_jsii_.Get(
		j,
		"sysctlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TransparentHugePageDefrag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageDefrag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TransparentHugePageDefragInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageDefragInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TransparentHugePageEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TransparentHugePageEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageEnabledInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterNodePoolLinuxOsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterNodePoolLinuxOsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterNodePoolLinuxOsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolLinuxOsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterNodePoolLinuxOsConfigOutputReference_Override(k KubernetesClusterNodePoolLinuxOsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolLinuxOsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference)SetInternalValue(val *KubernetesClusterNodePoolLinuxOsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference)SetSwapFileSizeMb(val *float64) {
	if err := j.validateSetSwapFileSizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"swapFileSizeMb",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference)SetTransparentHugePageDefrag(val *string) {
	if err := j.validateSetTransparentHugePageDefragParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentHugePageDefrag",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference)SetTransparentHugePageEnabled(val *string) {
	if err := j.validateSetTransparentHugePageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentHugePageEnabled",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) PutSysctlConfig(value *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig) {
	if err := k.validatePutSysctlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSysctlConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ResetSwapFileSizeMb() {
	_jsii_.InvokeVoid(
		k,
		"resetSwapFileSizeMb",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ResetSysctlConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetSysctlConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ResetTransparentHugePageDefrag() {
	_jsii_.InvokeVoid(
		k,
		"resetTransparentHugePageDefrag",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ResetTransparentHugePageEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetTransparentHugePageEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

