// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/apimanagement/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementHostnameConfigurationOutputReference interface {
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
	DeveloperPortal() ApiManagementHostnameConfigurationDeveloperPortalList
	DeveloperPortalInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ApiManagementHostnameConfiguration
	SetInternalValue(val *ApiManagementHostnameConfiguration)
	Management() ApiManagementHostnameConfigurationManagementList
	ManagementInput() interface{}
	Portal() ApiManagementHostnameConfigurationPortalList
	PortalInput() interface{}
	Proxy() ApiManagementHostnameConfigurationProxyList
	ProxyInput() interface{}
	Scm() ApiManagementHostnameConfigurationScmList
	ScmInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDeveloperPortal(value interface{})
	PutManagement(value interface{})
	PutPortal(value interface{})
	PutProxy(value interface{})
	PutScm(value interface{})
	ResetDeveloperPortal()
	ResetManagement()
	ResetPortal()
	ResetProxy()
	ResetScm()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationOutputReference
type jsiiProxy_ApiManagementHostnameConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) DeveloperPortal() ApiManagementHostnameConfigurationDeveloperPortalList {
	var returns ApiManagementHostnameConfigurationDeveloperPortalList
	_jsii_.Get(
		j,
		"developerPortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) DeveloperPortalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerPortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) InternalValue() *ApiManagementHostnameConfiguration {
	var returns *ApiManagementHostnameConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Management() ApiManagementHostnameConfigurationManagementList {
	var returns ApiManagementHostnameConfigurationManagementList
	_jsii_.Get(
		j,
		"management",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ManagementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Portal() ApiManagementHostnameConfigurationPortalList {
	var returns ApiManagementHostnameConfigurationPortalList
	_jsii_.Get(
		j,
		"portal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PortalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Proxy() ApiManagementHostnameConfigurationProxyList {
	var returns ApiManagementHostnameConfigurationProxyList
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ProxyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Scm() ApiManagementHostnameConfigurationScmList {
	var returns ApiManagementHostnameConfigurationScmList
	_jsii_.Get(
		j,
		"scm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ScmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementHostnameConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewApiManagementHostnameConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementHostnameConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationOutputReference_Override(a ApiManagementHostnameConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference)SetInternalValue(val *ApiManagementHostnameConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PutDeveloperPortal(value interface{}) {
	if err := a.validatePutDeveloperPortalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeveloperPortal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PutManagement(value interface{}) {
	if err := a.validatePutManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PutPortal(value interface{}) {
	if err := a.validatePutPortalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPortal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PutProxy(value interface{}) {
	if err := a.validatePutProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProxy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) PutScm(value interface{}) {
	if err := a.validatePutScmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ResetDeveloperPortal() {
	_jsii_.InvokeVoid(
		a,
		"resetDeveloperPortal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ResetManagement() {
	_jsii_.InvokeVoid(
		a,
		"resetManagement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ResetPortal() {
	_jsii_.InvokeVoid(
		a,
		"resetPortal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ResetProxy() {
	_jsii_.InvokeVoid(
		a,
		"resetProxy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ResetScm() {
	_jsii_.InvokeVoid(
		a,
		"resetScm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

