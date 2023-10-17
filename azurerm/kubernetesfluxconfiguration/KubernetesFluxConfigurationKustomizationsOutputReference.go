// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfluxconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/kubernetesfluxconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesFluxConfigurationKustomizationsOutputReference interface {
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
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DependsOnInput() *[]*string
	// Experimental.
	Fqn() *string
	GarbageCollectionEnabled() interface{}
	SetGarbageCollectionEnabled(val interface{})
	GarbageCollectionEnabledInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	RecreatingEnabled() interface{}
	SetRecreatingEnabled(val interface{})
	RecreatingEnabledInput() interface{}
	RetryIntervalInSeconds() *float64
	SetRetryIntervalInSeconds(val *float64)
	RetryIntervalInSecondsInput() *float64
	SyncIntervalInSeconds() *float64
	SetSyncIntervalInSeconds(val *float64)
	SyncIntervalInSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutInSeconds() *float64
	SetTimeoutInSeconds(val *float64)
	TimeoutInSecondsInput() *float64
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
	ResetDependsOn()
	ResetGarbageCollectionEnabled()
	ResetPath()
	ResetRecreatingEnabled()
	ResetRetryIntervalInSeconds()
	ResetSyncIntervalInSeconds()
	ResetTimeoutInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesFluxConfigurationKustomizationsOutputReference
type jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) DependsOnInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) GarbageCollectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"garbageCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) GarbageCollectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"garbageCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) RecreatingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recreatingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) RecreatingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recreatingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) RetryIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) RetryIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) SyncIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) SyncIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}


func NewKubernetesFluxConfigurationKustomizationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KubernetesFluxConfigurationKustomizationsOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesFluxConfigurationKustomizationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesFluxConfiguration.KubernetesFluxConfigurationKustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKubernetesFluxConfigurationKustomizationsOutputReference_Override(k KubernetesFluxConfigurationKustomizationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesFluxConfiguration.KubernetesFluxConfigurationKustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetDependsOn(val *[]*string) {
	if err := j.validateSetDependsOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetGarbageCollectionEnabled(val interface{}) {
	if err := j.validateSetGarbageCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"garbageCollectionEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetRecreatingEnabled(val interface{}) {
	if err := j.validateSetRecreatingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recreatingEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetRetryIntervalInSeconds(val *float64) {
	if err := j.validateSetRetryIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetSyncIntervalInSeconds(val *float64) {
	if err := j.validateSetSyncIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		k,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) ResetGarbageCollectionEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetGarbageCollectionEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		k,
		"resetPath",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) ResetRecreatingEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetRecreatingEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) ResetRetryIntervalInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetRetryIntervalInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) ResetSyncIntervalInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetSyncIntervalInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesFluxConfigurationKustomizationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

