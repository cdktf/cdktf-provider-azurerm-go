// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermlinuxwebapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/dataazurermlinuxwebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermLinuxWebAppAuthSettingsV2OutputReference interface {
	cdktf.ComplexObject
	ActiveDirectoryV2() DataAzurermLinuxWebAppAuthSettingsV2ActiveDirectoryV2List
	AppleV2() DataAzurermLinuxWebAppAuthSettingsV2AppleV2List
	AuthEnabled() cdktf.IResolvable
	AzureStaticWebAppV2() DataAzurermLinuxWebAppAuthSettingsV2AzureStaticWebAppV2List
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
	CustomOidcV2() DataAzurermLinuxWebAppAuthSettingsV2CustomOidcV2List
	DefaultProvider() *string
	ExcludedPaths() *[]*string
	FacebookV2() DataAzurermLinuxWebAppAuthSettingsV2FacebookV2List
	ForwardProxyConvention() *string
	ForwardProxyCustomHostHeaderName() *string
	ForwardProxyCustomSchemeHeaderName() *string
	// Experimental.
	Fqn() *string
	GithubV2() DataAzurermLinuxWebAppAuthSettingsV2GithubV2List
	GoogleV2() DataAzurermLinuxWebAppAuthSettingsV2GoogleV2List
	HttpRouteApiPrefix() *string
	InternalValue() *DataAzurermLinuxWebAppAuthSettingsV2
	SetInternalValue(val *DataAzurermLinuxWebAppAuthSettingsV2)
	Login() DataAzurermLinuxWebAppAuthSettingsV2LoginList
	MicrosoftV2() DataAzurermLinuxWebAppAuthSettingsV2MicrosoftV2List
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
	TwitterV2() DataAzurermLinuxWebAppAuthSettingsV2TwitterV2List
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

// The jsii proxy struct for DataAzurermLinuxWebAppAuthSettingsV2OutputReference
type jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) ActiveDirectoryV2() DataAzurermLinuxWebAppAuthSettingsV2ActiveDirectoryV2List {
	var returns DataAzurermLinuxWebAppAuthSettingsV2ActiveDirectoryV2List
	_jsii_.Get(
		j,
		"activeDirectoryV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) AppleV2() DataAzurermLinuxWebAppAuthSettingsV2AppleV2List {
	var returns DataAzurermLinuxWebAppAuthSettingsV2AppleV2List
	_jsii_.Get(
		j,
		"appleV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) AuthEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"authEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) AzureStaticWebAppV2() DataAzurermLinuxWebAppAuthSettingsV2AzureStaticWebAppV2List {
	var returns DataAzurermLinuxWebAppAuthSettingsV2AzureStaticWebAppV2List
	_jsii_.Get(
		j,
		"azureStaticWebAppV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) ConfigFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) CustomOidcV2() DataAzurermLinuxWebAppAuthSettingsV2CustomOidcV2List {
	var returns DataAzurermLinuxWebAppAuthSettingsV2CustomOidcV2List
	_jsii_.Get(
		j,
		"customOidcV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) DefaultProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) ExcludedPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) FacebookV2() DataAzurermLinuxWebAppAuthSettingsV2FacebookV2List {
	var returns DataAzurermLinuxWebAppAuthSettingsV2FacebookV2List
	_jsii_.Get(
		j,
		"facebookV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) ForwardProxyConvention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyConvention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) ForwardProxyCustomHostHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomHostHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) ForwardProxyCustomSchemeHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomSchemeHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) GithubV2() DataAzurermLinuxWebAppAuthSettingsV2GithubV2List {
	var returns DataAzurermLinuxWebAppAuthSettingsV2GithubV2List
	_jsii_.Get(
		j,
		"githubV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) GoogleV2() DataAzurermLinuxWebAppAuthSettingsV2GoogleV2List {
	var returns DataAzurermLinuxWebAppAuthSettingsV2GoogleV2List
	_jsii_.Get(
		j,
		"googleV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) HttpRouteApiPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouteApiPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) InternalValue() *DataAzurermLinuxWebAppAuthSettingsV2 {
	var returns *DataAzurermLinuxWebAppAuthSettingsV2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) Login() DataAzurermLinuxWebAppAuthSettingsV2LoginList {
	var returns DataAzurermLinuxWebAppAuthSettingsV2LoginList
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) MicrosoftV2() DataAzurermLinuxWebAppAuthSettingsV2MicrosoftV2List {
	var returns DataAzurermLinuxWebAppAuthSettingsV2MicrosoftV2List
	_jsii_.Get(
		j,
		"microsoftV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) RequireAuthentication() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requireAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) RequireHttps() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requireHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) TwitterV2() DataAzurermLinuxWebAppAuthSettingsV2TwitterV2List {
	var returns DataAzurermLinuxWebAppAuthSettingsV2TwitterV2List
	_jsii_.Get(
		j,
		"twitterV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) UnauthenticatedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedAction",
		&returns,
	)
	return returns
}


func NewDataAzurermLinuxWebAppAuthSettingsV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermLinuxWebAppAuthSettingsV2OutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermLinuxWebAppAuthSettingsV2OutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermLinuxWebApp.DataAzurermLinuxWebAppAuthSettingsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermLinuxWebAppAuthSettingsV2OutputReference_Override(d DataAzurermLinuxWebAppAuthSettingsV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermLinuxWebApp.DataAzurermLinuxWebAppAuthSettingsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference)SetInternalValue(val *DataAzurermLinuxWebAppAuthSettingsV2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAzurermLinuxWebAppAuthSettingsV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

