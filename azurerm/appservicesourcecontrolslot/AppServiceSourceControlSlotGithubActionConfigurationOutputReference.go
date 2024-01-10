// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservicesourcecontrolslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/appservicesourcecontrolslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppServiceSourceControlSlotGithubActionConfigurationOutputReference interface {
	cdktf.ComplexObject
	CodeConfiguration() AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference
	CodeConfigurationInput() *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration
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
	ContainerConfiguration() AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference
	ContainerConfigurationInput() *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GenerateWorkflowFile() interface{}
	SetGenerateWorkflowFile(val interface{})
	GenerateWorkflowFileInput() interface{}
	InternalValue() *AppServiceSourceControlSlotGithubActionConfiguration
	SetInternalValue(val *AppServiceSourceControlSlotGithubActionConfiguration)
	LinuxAction() cdktf.IResolvable
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
	PutCodeConfiguration(value *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration)
	PutContainerConfiguration(value *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration)
	ResetCodeConfiguration()
	ResetContainerConfiguration()
	ResetGenerateWorkflowFile()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSourceControlSlotGithubActionConfigurationOutputReference
type jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) CodeConfiguration() AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference {
	var returns AppServiceSourceControlSlotGithubActionConfigurationCodeConfigurationOutputReference
	_jsii_.Get(
		j,
		"codeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) CodeConfigurationInput() *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration {
	var returns *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration
	_jsii_.Get(
		j,
		"codeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ContainerConfiguration() AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference {
	var returns AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference
	_jsii_.Get(
		j,
		"containerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ContainerConfigurationInput() *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration {
	var returns *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration
	_jsii_.Get(
		j,
		"containerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GenerateWorkflowFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateWorkflowFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GenerateWorkflowFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateWorkflowFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) InternalValue() *AppServiceSourceControlSlotGithubActionConfiguration {
	var returns *AppServiceSourceControlSlotGithubActionConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) LinuxAction() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"linuxAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSourceControlSlotGithubActionConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSourceControlSlotGithubActionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAppServiceSourceControlSlotGithubActionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotGithubActionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSourceControlSlotGithubActionConfigurationOutputReference_Override(a AppServiceSourceControlSlotGithubActionConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotGithubActionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference)SetGenerateWorkflowFile(val interface{}) {
	if err := j.validateSetGenerateWorkflowFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generateWorkflowFile",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference)SetInternalValue(val *AppServiceSourceControlSlotGithubActionConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) PutCodeConfiguration(value *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration) {
	if err := a.validatePutCodeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) PutContainerConfiguration(value *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration) {
	if err := a.validatePutContainerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContainerConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ResetCodeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ResetContainerConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ResetGenerateWorkflowFile() {
	_jsii_.InvokeVoid(
		a,
		"resetGenerateWorkflowFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

