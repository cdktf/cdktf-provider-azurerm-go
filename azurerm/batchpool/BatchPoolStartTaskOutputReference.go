// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/batchpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchPoolStartTaskOutputReference interface {
	cdktf.ComplexObject
	CommandLine() *string
	SetCommandLine(val *string)
	CommandLineInput() *string
	CommonEnvironmentProperties() *map[string]*string
	SetCommonEnvironmentProperties(val *map[string]*string)
	CommonEnvironmentPropertiesInput() *map[string]*string
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
	Container() BatchPoolStartTaskContainerList
	ContainerInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *BatchPoolStartTask
	SetInternalValue(val *BatchPoolStartTask)
	ResourceFile() BatchPoolStartTaskResourceFileList
	ResourceFileInput() interface{}
	TaskRetryMaximum() *float64
	SetTaskRetryMaximum(val *float64)
	TaskRetryMaximumInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserIdentity() BatchPoolStartTaskUserIdentityOutputReference
	UserIdentityInput() *BatchPoolStartTaskUserIdentity
	WaitForSuccess() interface{}
	SetWaitForSuccess(val interface{})
	WaitForSuccessInput() interface{}
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
	PutContainer(value interface{})
	PutResourceFile(value interface{})
	PutUserIdentity(value *BatchPoolStartTaskUserIdentity)
	ResetCommonEnvironmentProperties()
	ResetContainer()
	ResetResourceFile()
	ResetTaskRetryMaximum()
	ResetWaitForSuccess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolStartTaskOutputReference
type jsiiProxy_BatchPoolStartTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) CommandLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commandLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) CommandLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commandLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) CommonEnvironmentProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"commonEnvironmentProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) CommonEnvironmentPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"commonEnvironmentPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) Container() BatchPoolStartTaskContainerList {
	var returns BatchPoolStartTaskContainerList
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) ContainerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) InternalValue() *BatchPoolStartTask {
	var returns *BatchPoolStartTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) ResourceFile() BatchPoolStartTaskResourceFileList {
	var returns BatchPoolStartTaskResourceFileList
	_jsii_.Get(
		j,
		"resourceFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) ResourceFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) TaskRetryMaximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskRetryMaximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) TaskRetryMaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskRetryMaximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) UserIdentity() BatchPoolStartTaskUserIdentityOutputReference {
	var returns BatchPoolStartTaskUserIdentityOutputReference
	_jsii_.Get(
		j,
		"userIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) UserIdentityInput() *BatchPoolStartTaskUserIdentity {
	var returns *BatchPoolStartTaskUserIdentity
	_jsii_.Get(
		j,
		"userIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) WaitForSuccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference) WaitForSuccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSuccessInput",
		&returns,
	)
	return returns
}


func NewBatchPoolStartTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchPoolStartTaskOutputReference {
	_init_.Initialize()

	if err := validateNewBatchPoolStartTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchPoolStartTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchPoolStartTaskOutputReference_Override(b BatchPoolStartTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchPool.BatchPoolStartTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference)SetCommandLine(val *string) {
	if err := j.validateSetCommandLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commandLine",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference)SetCommonEnvironmentProperties(val *map[string]*string) {
	if err := j.validateSetCommonEnvironmentPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonEnvironmentProperties",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference)SetInternalValue(val *BatchPoolStartTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference)SetTaskRetryMaximum(val *float64) {
	if err := j.validateSetTaskRetryMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskRetryMaximum",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolStartTaskOutputReference)SetWaitForSuccess(val interface{}) {
	if err := j.validateSetWaitForSuccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForSuccess",
		val,
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) PutContainer(value interface{}) {
	if err := b.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putContainer",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) PutResourceFile(value interface{}) {
	if err := b.validatePutResourceFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putResourceFile",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) PutUserIdentity(value *BatchPoolStartTaskUserIdentity) {
	if err := b.validatePutUserIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putUserIdentity",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ResetCommonEnvironmentProperties() {
	_jsii_.InvokeVoid(
		b,
		"resetCommonEnvironmentProperties",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		b,
		"resetContainer",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ResetResourceFile() {
	_jsii_.InvokeVoid(
		b,
		"resetResourceFile",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ResetTaskRetryMaximum() {
	_jsii_.InvokeVoid(
		b,
		"resetTaskRetryMaximum",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ResetWaitForSuccess() {
	_jsii_.InvokeVoid(
		b,
		"resetWaitForSuccess",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolStartTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

