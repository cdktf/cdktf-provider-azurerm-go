// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arckubernetesfluxconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/arckubernetesfluxconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArcKubernetesFluxConfigurationKustomizationsOutputReference interface {
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

// The jsii proxy struct for ArcKubernetesFluxConfigurationKustomizationsOutputReference
type jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) DependsOnInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) GarbageCollectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"garbageCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) GarbageCollectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"garbageCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) RecreatingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recreatingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) RecreatingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recreatingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) RetryIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) RetryIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) SyncIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) SyncIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}


func NewArcKubernetesFluxConfigurationKustomizationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ArcKubernetesFluxConfigurationKustomizationsOutputReference {
	_init_.Initialize()

	if err := validateNewArcKubernetesFluxConfigurationKustomizationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.arcKubernetesFluxConfiguration.ArcKubernetesFluxConfigurationKustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewArcKubernetesFluxConfigurationKustomizationsOutputReference_Override(a ArcKubernetesFluxConfigurationKustomizationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.arcKubernetesFluxConfiguration.ArcKubernetesFluxConfigurationKustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetDependsOn(val *[]*string) {
	if err := j.validateSetDependsOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetGarbageCollectionEnabled(val interface{}) {
	if err := j.validateSetGarbageCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"garbageCollectionEnabled",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetRecreatingEnabled(val interface{}) {
	if err := j.validateSetRecreatingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recreatingEnabled",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetRetryIntervalInSeconds(val *float64) {
	if err := j.validateSetRetryIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetSyncIntervalInSeconds(val *float64) {
	if err := j.validateSetSyncIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		a,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) ResetGarbageCollectionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetGarbageCollectionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) ResetRecreatingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRecreatingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) ResetRetryIntervalInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryIntervalInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) ResetSyncIntervalInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncIntervalInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationKustomizationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

