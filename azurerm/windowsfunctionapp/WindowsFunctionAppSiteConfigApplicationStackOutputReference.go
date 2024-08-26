// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/windowsfunctionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsFunctionAppSiteConfigApplicationStackOutputReference interface {
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
	DotnetVersion() *string
	SetDotnetVersion(val *string)
	DotnetVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsFunctionAppSiteConfigApplicationStack
	SetInternalValue(val *WindowsFunctionAppSiteConfigApplicationStack)
	JavaVersion() *string
	SetJavaVersion(val *string)
	JavaVersionInput() *string
	NodeVersion() *string
	SetNodeVersion(val *string)
	NodeVersionInput() *string
	PowershellCoreVersion() *string
	SetPowershellCoreVersion(val *string)
	PowershellCoreVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseCustomRuntime() interface{}
	SetUseCustomRuntime(val interface{})
	UseCustomRuntimeInput() interface{}
	UseDotnetIsolatedRuntime() interface{}
	SetUseDotnetIsolatedRuntime(val interface{})
	UseDotnetIsolatedRuntimeInput() interface{}
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
	ResetDotnetVersion()
	ResetJavaVersion()
	ResetNodeVersion()
	ResetPowershellCoreVersion()
	ResetUseCustomRuntime()
	ResetUseDotnetIsolatedRuntime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsFunctionAppSiteConfigApplicationStackOutputReference
type jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) DotnetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) DotnetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) InternalValue() *WindowsFunctionAppSiteConfigApplicationStack {
	var returns *WindowsFunctionAppSiteConfigApplicationStack
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) JavaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) JavaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) NodeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) PowershellCoreVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powershellCoreVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) PowershellCoreVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powershellCoreVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) UseCustomRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) UseCustomRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) UseDotnetIsolatedRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDotnetIsolatedRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) UseDotnetIsolatedRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDotnetIsolatedRuntimeInput",
		&returns,
	)
	return returns
}


func NewWindowsFunctionAppSiteConfigApplicationStackOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsFunctionAppSiteConfigApplicationStackOutputReference {
	_init_.Initialize()

	if err := validateNewWindowsFunctionAppSiteConfigApplicationStackOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsFunctionApp.WindowsFunctionAppSiteConfigApplicationStackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsFunctionAppSiteConfigApplicationStackOutputReference_Override(w WindowsFunctionAppSiteConfigApplicationStackOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsFunctionApp.WindowsFunctionAppSiteConfigApplicationStackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference)SetDotnetVersion(val *string) {
	if err := j.validateSetDotnetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dotnetVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference)SetInternalValue(val *WindowsFunctionAppSiteConfigApplicationStack) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference)SetJavaVersion(val *string) {
	if err := j.validateSetJavaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference)SetNodeVersion(val *string) {
	if err := j.validateSetNodeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference)SetPowershellCoreVersion(val *string) {
	if err := j.validateSetPowershellCoreVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powershellCoreVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference)SetUseCustomRuntime(val interface{}) {
	if err := j.validateSetUseCustomRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCustomRuntime",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference)SetUseDotnetIsolatedRuntime(val interface{}) {
	if err := j.validateSetUseDotnetIsolatedRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDotnetIsolatedRuntime",
		val,
	)
}

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) ResetDotnetVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetDotnetVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) ResetJavaVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetJavaVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) ResetNodeVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetNodeVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) ResetPowershellCoreVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetPowershellCoreVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) ResetUseCustomRuntime() {
	_jsii_.InvokeVoid(
		w,
		"resetUseCustomRuntime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) ResetUseDotnetIsolatedRuntime() {
	_jsii_.InvokeVoid(
		w,
		"resetUseDotnetIsolatedRuntime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WindowsFunctionAppSiteConfigApplicationStackOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

