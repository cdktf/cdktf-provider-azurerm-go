// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermwindowswebapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/dataazurermwindowswebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermWindowsWebAppAuthSettingsV2OutputReference interface {
	cdktf.ComplexObject
	ActiveDirectoryV2() DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2List
	AppleV2() DataAzurermWindowsWebAppAuthSettingsV2AppleV2List
	AuthEnabled() cdktf.IResolvable
	AzureStaticWebAppV2() DataAzurermWindowsWebAppAuthSettingsV2AzureStaticWebAppV2List
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
	ConfigFilePath() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomOidcV2() DataAzurermWindowsWebAppAuthSettingsV2CustomOidcV2List
	DefaultProvider() *string
	ExcludedPaths() *[]*string
	FacebookV2() DataAzurermWindowsWebAppAuthSettingsV2FacebookV2List
	ForwardProxyConvention() *string
	ForwardProxyCustomHostHeaderName() *string
	ForwardProxyCustomSchemeHeaderName() *string
	// Experimental.
	Fqn() *string
	GithubV2() DataAzurermWindowsWebAppAuthSettingsV2GithubV2List
	GoogleV2() DataAzurermWindowsWebAppAuthSettingsV2GoogleV2List
	HttpRouteApiPrefix() *string
	InternalValue() *DataAzurermWindowsWebAppAuthSettingsV2
	SetInternalValue(val *DataAzurermWindowsWebAppAuthSettingsV2)
	Login() DataAzurermWindowsWebAppAuthSettingsV2LoginList
	MicrosoftV2() DataAzurermWindowsWebAppAuthSettingsV2MicrosoftV2List
	RequireAuthentication() cdktf.IResolvable
	RequireHttps() cdktf.IResolvable
	RuntimeVersion() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TwitterV2() DataAzurermWindowsWebAppAuthSettingsV2TwitterV2List
	UnauthenticatedAction() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurermWindowsWebAppAuthSettingsV2OutputReference
type jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) ActiveDirectoryV2() DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2List {
	var returns DataAzurermWindowsWebAppAuthSettingsV2ActiveDirectoryV2List
	_jsii_.Get(
		j,
		"activeDirectoryV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) AppleV2() DataAzurermWindowsWebAppAuthSettingsV2AppleV2List {
	var returns DataAzurermWindowsWebAppAuthSettingsV2AppleV2List
	_jsii_.Get(
		j,
		"appleV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) AuthEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"authEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) AzureStaticWebAppV2() DataAzurermWindowsWebAppAuthSettingsV2AzureStaticWebAppV2List {
	var returns DataAzurermWindowsWebAppAuthSettingsV2AzureStaticWebAppV2List
	_jsii_.Get(
		j,
		"azureStaticWebAppV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) ConfigFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) CustomOidcV2() DataAzurermWindowsWebAppAuthSettingsV2CustomOidcV2List {
	var returns DataAzurermWindowsWebAppAuthSettingsV2CustomOidcV2List
	_jsii_.Get(
		j,
		"customOidcV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) DefaultProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) ExcludedPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) FacebookV2() DataAzurermWindowsWebAppAuthSettingsV2FacebookV2List {
	var returns DataAzurermWindowsWebAppAuthSettingsV2FacebookV2List
	_jsii_.Get(
		j,
		"facebookV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) ForwardProxyConvention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyConvention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) ForwardProxyCustomHostHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomHostHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) ForwardProxyCustomSchemeHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomSchemeHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) GithubV2() DataAzurermWindowsWebAppAuthSettingsV2GithubV2List {
	var returns DataAzurermWindowsWebAppAuthSettingsV2GithubV2List
	_jsii_.Get(
		j,
		"githubV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) GoogleV2() DataAzurermWindowsWebAppAuthSettingsV2GoogleV2List {
	var returns DataAzurermWindowsWebAppAuthSettingsV2GoogleV2List
	_jsii_.Get(
		j,
		"googleV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) HttpRouteApiPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouteApiPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) InternalValue() *DataAzurermWindowsWebAppAuthSettingsV2 {
	var returns *DataAzurermWindowsWebAppAuthSettingsV2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) Login() DataAzurermWindowsWebAppAuthSettingsV2LoginList {
	var returns DataAzurermWindowsWebAppAuthSettingsV2LoginList
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) MicrosoftV2() DataAzurermWindowsWebAppAuthSettingsV2MicrosoftV2List {
	var returns DataAzurermWindowsWebAppAuthSettingsV2MicrosoftV2List
	_jsii_.Get(
		j,
		"microsoftV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) RequireAuthentication() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requireAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) RequireHttps() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requireHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) TwitterV2() DataAzurermWindowsWebAppAuthSettingsV2TwitterV2List {
	var returns DataAzurermWindowsWebAppAuthSettingsV2TwitterV2List
	_jsii_.Get(
		j,
		"twitterV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) UnauthenticatedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedAction",
		&returns,
	)
	return returns
}


func NewDataAzurermWindowsWebAppAuthSettingsV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermWindowsWebAppAuthSettingsV2OutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermWindowsWebAppAuthSettingsV2OutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermWindowsWebApp.DataAzurermWindowsWebAppAuthSettingsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermWindowsWebAppAuthSettingsV2OutputReference_Override(d DataAzurermWindowsWebAppAuthSettingsV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermWindowsWebApp.DataAzurermWindowsWebAppAuthSettingsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference)SetInternalValue(val *DataAzurermWindowsWebAppAuthSettingsV2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsWebAppAuthSettingsV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

