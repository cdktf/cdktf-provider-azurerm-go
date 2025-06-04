// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/linuxfunctionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxFunctionAppAuthSettingsV2OutputReference interface {
	cdktf.ComplexObject
	ActiveDirectoryV2() LinuxFunctionAppAuthSettingsV2ActiveDirectoryV2OutputReference
	ActiveDirectoryV2Input() *LinuxFunctionAppAuthSettingsV2ActiveDirectoryV2
	AppleV2() LinuxFunctionAppAuthSettingsV2AppleV2OutputReference
	AppleV2Input() *LinuxFunctionAppAuthSettingsV2AppleV2
	AuthEnabled() interface{}
	SetAuthEnabled(val interface{})
	AuthEnabledInput() interface{}
	AzureStaticWebAppV2() LinuxFunctionAppAuthSettingsV2AzureStaticWebAppV2OutputReference
	AzureStaticWebAppV2Input() *LinuxFunctionAppAuthSettingsV2AzureStaticWebAppV2
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
	SetConfigFilePath(val *string)
	ConfigFilePathInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomOidcV2() LinuxFunctionAppAuthSettingsV2CustomOidcV2List
	CustomOidcV2Input() interface{}
	DefaultProvider() *string
	SetDefaultProvider(val *string)
	DefaultProviderInput() *string
	ExcludedPaths() *[]*string
	SetExcludedPaths(val *[]*string)
	ExcludedPathsInput() *[]*string
	FacebookV2() LinuxFunctionAppAuthSettingsV2FacebookV2OutputReference
	FacebookV2Input() *LinuxFunctionAppAuthSettingsV2FacebookV2
	ForwardProxyConvention() *string
	SetForwardProxyConvention(val *string)
	ForwardProxyConventionInput() *string
	ForwardProxyCustomHostHeaderName() *string
	SetForwardProxyCustomHostHeaderName(val *string)
	ForwardProxyCustomHostHeaderNameInput() *string
	ForwardProxyCustomSchemeHeaderName() *string
	SetForwardProxyCustomSchemeHeaderName(val *string)
	ForwardProxyCustomSchemeHeaderNameInput() *string
	// Experimental.
	Fqn() *string
	GithubV2() LinuxFunctionAppAuthSettingsV2GithubV2OutputReference
	GithubV2Input() *LinuxFunctionAppAuthSettingsV2GithubV2
	GoogleV2() LinuxFunctionAppAuthSettingsV2GoogleV2OutputReference
	GoogleV2Input() *LinuxFunctionAppAuthSettingsV2GoogleV2
	HttpRouteApiPrefix() *string
	SetHttpRouteApiPrefix(val *string)
	HttpRouteApiPrefixInput() *string
	InternalValue() *LinuxFunctionAppAuthSettingsV2
	SetInternalValue(val *LinuxFunctionAppAuthSettingsV2)
	Login() LinuxFunctionAppAuthSettingsV2LoginOutputReference
	LoginInput() *LinuxFunctionAppAuthSettingsV2Login
	MicrosoftV2() LinuxFunctionAppAuthSettingsV2MicrosoftV2OutputReference
	MicrosoftV2Input() *LinuxFunctionAppAuthSettingsV2MicrosoftV2
	RequireAuthentication() interface{}
	SetRequireAuthentication(val interface{})
	RequireAuthenticationInput() interface{}
	RequireHttps() interface{}
	SetRequireHttps(val interface{})
	RequireHttpsInput() interface{}
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TwitterV2() LinuxFunctionAppAuthSettingsV2TwitterV2OutputReference
	TwitterV2Input() *LinuxFunctionAppAuthSettingsV2TwitterV2
	UnauthenticatedAction() *string
	SetUnauthenticatedAction(val *string)
	UnauthenticatedActionInput() *string
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
	PutActiveDirectoryV2(value *LinuxFunctionAppAuthSettingsV2ActiveDirectoryV2)
	PutAppleV2(value *LinuxFunctionAppAuthSettingsV2AppleV2)
	PutAzureStaticWebAppV2(value *LinuxFunctionAppAuthSettingsV2AzureStaticWebAppV2)
	PutCustomOidcV2(value interface{})
	PutFacebookV2(value *LinuxFunctionAppAuthSettingsV2FacebookV2)
	PutGithubV2(value *LinuxFunctionAppAuthSettingsV2GithubV2)
	PutGoogleV2(value *LinuxFunctionAppAuthSettingsV2GoogleV2)
	PutLogin(value *LinuxFunctionAppAuthSettingsV2Login)
	PutMicrosoftV2(value *LinuxFunctionAppAuthSettingsV2MicrosoftV2)
	PutTwitterV2(value *LinuxFunctionAppAuthSettingsV2TwitterV2)
	ResetActiveDirectoryV2()
	ResetAppleV2()
	ResetAuthEnabled()
	ResetAzureStaticWebAppV2()
	ResetConfigFilePath()
	ResetCustomOidcV2()
	ResetDefaultProvider()
	ResetExcludedPaths()
	ResetFacebookV2()
	ResetForwardProxyConvention()
	ResetForwardProxyCustomHostHeaderName()
	ResetForwardProxyCustomSchemeHeaderName()
	ResetGithubV2()
	ResetGoogleV2()
	ResetHttpRouteApiPrefix()
	ResetMicrosoftV2()
	ResetRequireAuthentication()
	ResetRequireHttps()
	ResetRuntimeVersion()
	ResetTwitterV2()
	ResetUnauthenticatedAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LinuxFunctionAppAuthSettingsV2OutputReference
type jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ActiveDirectoryV2() LinuxFunctionAppAuthSettingsV2ActiveDirectoryV2OutputReference {
	var returns LinuxFunctionAppAuthSettingsV2ActiveDirectoryV2OutputReference
	_jsii_.Get(
		j,
		"activeDirectoryV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ActiveDirectoryV2Input() *LinuxFunctionAppAuthSettingsV2ActiveDirectoryV2 {
	var returns *LinuxFunctionAppAuthSettingsV2ActiveDirectoryV2
	_jsii_.Get(
		j,
		"activeDirectoryV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) AppleV2() LinuxFunctionAppAuthSettingsV2AppleV2OutputReference {
	var returns LinuxFunctionAppAuthSettingsV2AppleV2OutputReference
	_jsii_.Get(
		j,
		"appleV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) AppleV2Input() *LinuxFunctionAppAuthSettingsV2AppleV2 {
	var returns *LinuxFunctionAppAuthSettingsV2AppleV2
	_jsii_.Get(
		j,
		"appleV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) AuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) AuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) AzureStaticWebAppV2() LinuxFunctionAppAuthSettingsV2AzureStaticWebAppV2OutputReference {
	var returns LinuxFunctionAppAuthSettingsV2AzureStaticWebAppV2OutputReference
	_jsii_.Get(
		j,
		"azureStaticWebAppV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) AzureStaticWebAppV2Input() *LinuxFunctionAppAuthSettingsV2AzureStaticWebAppV2 {
	var returns *LinuxFunctionAppAuthSettingsV2AzureStaticWebAppV2
	_jsii_.Get(
		j,
		"azureStaticWebAppV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ConfigFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ConfigFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) CustomOidcV2() LinuxFunctionAppAuthSettingsV2CustomOidcV2List {
	var returns LinuxFunctionAppAuthSettingsV2CustomOidcV2List
	_jsii_.Get(
		j,
		"customOidcV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) CustomOidcV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOidcV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) DefaultProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) DefaultProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ExcludedPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ExcludedPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) FacebookV2() LinuxFunctionAppAuthSettingsV2FacebookV2OutputReference {
	var returns LinuxFunctionAppAuthSettingsV2FacebookV2OutputReference
	_jsii_.Get(
		j,
		"facebookV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) FacebookV2Input() *LinuxFunctionAppAuthSettingsV2FacebookV2 {
	var returns *LinuxFunctionAppAuthSettingsV2FacebookV2
	_jsii_.Get(
		j,
		"facebookV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ForwardProxyConvention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyConvention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ForwardProxyConventionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyConventionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ForwardProxyCustomHostHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomHostHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ForwardProxyCustomHostHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomHostHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ForwardProxyCustomSchemeHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomSchemeHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ForwardProxyCustomSchemeHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomSchemeHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GithubV2() LinuxFunctionAppAuthSettingsV2GithubV2OutputReference {
	var returns LinuxFunctionAppAuthSettingsV2GithubV2OutputReference
	_jsii_.Get(
		j,
		"githubV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GithubV2Input() *LinuxFunctionAppAuthSettingsV2GithubV2 {
	var returns *LinuxFunctionAppAuthSettingsV2GithubV2
	_jsii_.Get(
		j,
		"githubV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GoogleV2() LinuxFunctionAppAuthSettingsV2GoogleV2OutputReference {
	var returns LinuxFunctionAppAuthSettingsV2GoogleV2OutputReference
	_jsii_.Get(
		j,
		"googleV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GoogleV2Input() *LinuxFunctionAppAuthSettingsV2GoogleV2 {
	var returns *LinuxFunctionAppAuthSettingsV2GoogleV2
	_jsii_.Get(
		j,
		"googleV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) HttpRouteApiPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouteApiPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) HttpRouteApiPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouteApiPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) InternalValue() *LinuxFunctionAppAuthSettingsV2 {
	var returns *LinuxFunctionAppAuthSettingsV2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) Login() LinuxFunctionAppAuthSettingsV2LoginOutputReference {
	var returns LinuxFunctionAppAuthSettingsV2LoginOutputReference
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) LoginInput() *LinuxFunctionAppAuthSettingsV2Login {
	var returns *LinuxFunctionAppAuthSettingsV2Login
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) MicrosoftV2() LinuxFunctionAppAuthSettingsV2MicrosoftV2OutputReference {
	var returns LinuxFunctionAppAuthSettingsV2MicrosoftV2OutputReference
	_jsii_.Get(
		j,
		"microsoftV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) MicrosoftV2Input() *LinuxFunctionAppAuthSettingsV2MicrosoftV2 {
	var returns *LinuxFunctionAppAuthSettingsV2MicrosoftV2
	_jsii_.Get(
		j,
		"microsoftV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) RequireAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) RequireAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) RequireHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) RequireHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) TwitterV2() LinuxFunctionAppAuthSettingsV2TwitterV2OutputReference {
	var returns LinuxFunctionAppAuthSettingsV2TwitterV2OutputReference
	_jsii_.Get(
		j,
		"twitterV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) TwitterV2Input() *LinuxFunctionAppAuthSettingsV2TwitterV2 {
	var returns *LinuxFunctionAppAuthSettingsV2TwitterV2
	_jsii_.Get(
		j,
		"twitterV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) UnauthenticatedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) UnauthenticatedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedActionInput",
		&returns,
	)
	return returns
}


func NewLinuxFunctionAppAuthSettingsV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LinuxFunctionAppAuthSettingsV2OutputReference {
	_init_.Initialize()

	if err := validateNewLinuxFunctionAppAuthSettingsV2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionAppAuthSettingsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLinuxFunctionAppAuthSettingsV2OutputReference_Override(l LinuxFunctionAppAuthSettingsV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionApp.LinuxFunctionAppAuthSettingsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetAuthEnabled(val interface{}) {
	if err := j.validateSetAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetConfigFilePath(val *string) {
	if err := j.validateSetConfigFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configFilePath",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetDefaultProvider(val *string) {
	if err := j.validateSetDefaultProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProvider",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetExcludedPaths(val *[]*string) {
	if err := j.validateSetExcludedPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedPaths",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetForwardProxyConvention(val *string) {
	if err := j.validateSetForwardProxyConventionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardProxyConvention",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetForwardProxyCustomHostHeaderName(val *string) {
	if err := j.validateSetForwardProxyCustomHostHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardProxyCustomHostHeaderName",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetForwardProxyCustomSchemeHeaderName(val *string) {
	if err := j.validateSetForwardProxyCustomSchemeHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardProxyCustomSchemeHeaderName",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetHttpRouteApiPrefix(val *string) {
	if err := j.validateSetHttpRouteApiPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRouteApiPrefix",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetInternalValue(val *LinuxFunctionAppAuthSettingsV2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetRequireAuthentication(val interface{}) {
	if err := j.validateSetRequireAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAuthentication",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetRequireHttps(val interface{}) {
	if err := j.validateSetRequireHttpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireHttps",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference)SetUnauthenticatedAction(val *string) {
	if err := j.validateSetUnauthenticatedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unauthenticatedAction",
		val,
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) PutActiveDirectoryV2(value *LinuxFunctionAppAuthSettingsV2ActiveDirectoryV2) {
	if err := l.validatePutActiveDirectoryV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putActiveDirectoryV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) PutAppleV2(value *LinuxFunctionAppAuthSettingsV2AppleV2) {
	if err := l.validatePutAppleV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAppleV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) PutAzureStaticWebAppV2(value *LinuxFunctionAppAuthSettingsV2AzureStaticWebAppV2) {
	if err := l.validatePutAzureStaticWebAppV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAzureStaticWebAppV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) PutCustomOidcV2(value interface{}) {
	if err := l.validatePutCustomOidcV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCustomOidcV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) PutFacebookV2(value *LinuxFunctionAppAuthSettingsV2FacebookV2) {
	if err := l.validatePutFacebookV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFacebookV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) PutGithubV2(value *LinuxFunctionAppAuthSettingsV2GithubV2) {
	if err := l.validatePutGithubV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGithubV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) PutGoogleV2(value *LinuxFunctionAppAuthSettingsV2GoogleV2) {
	if err := l.validatePutGoogleV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGoogleV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) PutLogin(value *LinuxFunctionAppAuthSettingsV2Login) {
	if err := l.validatePutLoginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLogin",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) PutMicrosoftV2(value *LinuxFunctionAppAuthSettingsV2MicrosoftV2) {
	if err := l.validatePutMicrosoftV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMicrosoftV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) PutTwitterV2(value *LinuxFunctionAppAuthSettingsV2TwitterV2) {
	if err := l.validatePutTwitterV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTwitterV2",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetActiveDirectoryV2() {
	_jsii_.InvokeVoid(
		l,
		"resetActiveDirectoryV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetAppleV2() {
	_jsii_.InvokeVoid(
		l,
		"resetAppleV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetAuthEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetAuthEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetAzureStaticWebAppV2() {
	_jsii_.InvokeVoid(
		l,
		"resetAzureStaticWebAppV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetConfigFilePath() {
	_jsii_.InvokeVoid(
		l,
		"resetConfigFilePath",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetCustomOidcV2() {
	_jsii_.InvokeVoid(
		l,
		"resetCustomOidcV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetDefaultProvider() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultProvider",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetExcludedPaths() {
	_jsii_.InvokeVoid(
		l,
		"resetExcludedPaths",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetFacebookV2() {
	_jsii_.InvokeVoid(
		l,
		"resetFacebookV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetForwardProxyConvention() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardProxyConvention",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetForwardProxyCustomHostHeaderName() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardProxyCustomHostHeaderName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetForwardProxyCustomSchemeHeaderName() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardProxyCustomSchemeHeaderName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetGithubV2() {
	_jsii_.InvokeVoid(
		l,
		"resetGithubV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetGoogleV2() {
	_jsii_.InvokeVoid(
		l,
		"resetGoogleV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetHttpRouteApiPrefix() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpRouteApiPrefix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetMicrosoftV2() {
	_jsii_.InvokeVoid(
		l,
		"resetMicrosoftV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetRequireAuthentication() {
	_jsii_.InvokeVoid(
		l,
		"resetRequireAuthentication",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetRequireHttps() {
	_jsii_.InvokeVoid(
		l,
		"resetRequireHttps",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetTwitterV2() {
	_jsii_.InvokeVoid(
		l,
		"resetTwitterV2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ResetUnauthenticatedAction() {
	_jsii_.InvokeVoid(
		l,
		"resetUnauthenticatedAction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppAuthSettingsV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

