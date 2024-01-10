// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/windowsfunctionappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsFunctionAppSlotAuthSettingsV2OutputReference interface {
	cdktf.ComplexObject
	ActiveDirectoryV2() WindowsFunctionAppSlotAuthSettingsV2ActiveDirectoryV2OutputReference
	ActiveDirectoryV2Input() *WindowsFunctionAppSlotAuthSettingsV2ActiveDirectoryV2
	AppleV2() WindowsFunctionAppSlotAuthSettingsV2AppleV2OutputReference
	AppleV2Input() *WindowsFunctionAppSlotAuthSettingsV2AppleV2
	AuthEnabled() interface{}
	SetAuthEnabled(val interface{})
	AuthEnabledInput() interface{}
	AzureStaticWebAppV2() WindowsFunctionAppSlotAuthSettingsV2AzureStaticWebAppV2OutputReference
	AzureStaticWebAppV2Input() *WindowsFunctionAppSlotAuthSettingsV2AzureStaticWebAppV2
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
	CustomOidcV2() WindowsFunctionAppSlotAuthSettingsV2CustomOidcV2List
	CustomOidcV2Input() interface{}
	DefaultProvider() *string
	SetDefaultProvider(val *string)
	DefaultProviderInput() *string
	ExcludedPaths() *[]*string
	SetExcludedPaths(val *[]*string)
	ExcludedPathsInput() *[]*string
	FacebookV2() WindowsFunctionAppSlotAuthSettingsV2FacebookV2OutputReference
	FacebookV2Input() *WindowsFunctionAppSlotAuthSettingsV2FacebookV2
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
	GithubV2() WindowsFunctionAppSlotAuthSettingsV2GithubV2OutputReference
	GithubV2Input() *WindowsFunctionAppSlotAuthSettingsV2GithubV2
	GoogleV2() WindowsFunctionAppSlotAuthSettingsV2GoogleV2OutputReference
	GoogleV2Input() *WindowsFunctionAppSlotAuthSettingsV2GoogleV2
	HttpRouteApiPrefix() *string
	SetHttpRouteApiPrefix(val *string)
	HttpRouteApiPrefixInput() *string
	InternalValue() *WindowsFunctionAppSlotAuthSettingsV2
	SetInternalValue(val *WindowsFunctionAppSlotAuthSettingsV2)
	Login() WindowsFunctionAppSlotAuthSettingsV2LoginOutputReference
	LoginInput() *WindowsFunctionAppSlotAuthSettingsV2Login
	MicrosoftV2() WindowsFunctionAppSlotAuthSettingsV2MicrosoftV2OutputReference
	MicrosoftV2Input() *WindowsFunctionAppSlotAuthSettingsV2MicrosoftV2
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
	TwitterV2() WindowsFunctionAppSlotAuthSettingsV2TwitterV2OutputReference
	TwitterV2Input() *WindowsFunctionAppSlotAuthSettingsV2TwitterV2
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
	PutActiveDirectoryV2(value *WindowsFunctionAppSlotAuthSettingsV2ActiveDirectoryV2)
	PutAppleV2(value *WindowsFunctionAppSlotAuthSettingsV2AppleV2)
	PutAzureStaticWebAppV2(value *WindowsFunctionAppSlotAuthSettingsV2AzureStaticWebAppV2)
	PutCustomOidcV2(value interface{})
	PutFacebookV2(value *WindowsFunctionAppSlotAuthSettingsV2FacebookV2)
	PutGithubV2(value *WindowsFunctionAppSlotAuthSettingsV2GithubV2)
	PutGoogleV2(value *WindowsFunctionAppSlotAuthSettingsV2GoogleV2)
	PutLogin(value *WindowsFunctionAppSlotAuthSettingsV2Login)
	PutMicrosoftV2(value *WindowsFunctionAppSlotAuthSettingsV2MicrosoftV2)
	PutTwitterV2(value *WindowsFunctionAppSlotAuthSettingsV2TwitterV2)
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

// The jsii proxy struct for WindowsFunctionAppSlotAuthSettingsV2OutputReference
type jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ActiveDirectoryV2() WindowsFunctionAppSlotAuthSettingsV2ActiveDirectoryV2OutputReference {
	var returns WindowsFunctionAppSlotAuthSettingsV2ActiveDirectoryV2OutputReference
	_jsii_.Get(
		j,
		"activeDirectoryV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ActiveDirectoryV2Input() *WindowsFunctionAppSlotAuthSettingsV2ActiveDirectoryV2 {
	var returns *WindowsFunctionAppSlotAuthSettingsV2ActiveDirectoryV2
	_jsii_.Get(
		j,
		"activeDirectoryV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) AppleV2() WindowsFunctionAppSlotAuthSettingsV2AppleV2OutputReference {
	var returns WindowsFunctionAppSlotAuthSettingsV2AppleV2OutputReference
	_jsii_.Get(
		j,
		"appleV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) AppleV2Input() *WindowsFunctionAppSlotAuthSettingsV2AppleV2 {
	var returns *WindowsFunctionAppSlotAuthSettingsV2AppleV2
	_jsii_.Get(
		j,
		"appleV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) AuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) AuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) AzureStaticWebAppV2() WindowsFunctionAppSlotAuthSettingsV2AzureStaticWebAppV2OutputReference {
	var returns WindowsFunctionAppSlotAuthSettingsV2AzureStaticWebAppV2OutputReference
	_jsii_.Get(
		j,
		"azureStaticWebAppV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) AzureStaticWebAppV2Input() *WindowsFunctionAppSlotAuthSettingsV2AzureStaticWebAppV2 {
	var returns *WindowsFunctionAppSlotAuthSettingsV2AzureStaticWebAppV2
	_jsii_.Get(
		j,
		"azureStaticWebAppV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ConfigFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ConfigFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) CustomOidcV2() WindowsFunctionAppSlotAuthSettingsV2CustomOidcV2List {
	var returns WindowsFunctionAppSlotAuthSettingsV2CustomOidcV2List
	_jsii_.Get(
		j,
		"customOidcV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) CustomOidcV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOidcV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) DefaultProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) DefaultProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ExcludedPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ExcludedPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) FacebookV2() WindowsFunctionAppSlotAuthSettingsV2FacebookV2OutputReference {
	var returns WindowsFunctionAppSlotAuthSettingsV2FacebookV2OutputReference
	_jsii_.Get(
		j,
		"facebookV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) FacebookV2Input() *WindowsFunctionAppSlotAuthSettingsV2FacebookV2 {
	var returns *WindowsFunctionAppSlotAuthSettingsV2FacebookV2
	_jsii_.Get(
		j,
		"facebookV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ForwardProxyConvention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyConvention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ForwardProxyConventionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyConventionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ForwardProxyCustomHostHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomHostHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ForwardProxyCustomHostHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomHostHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ForwardProxyCustomSchemeHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomSchemeHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ForwardProxyCustomSchemeHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomSchemeHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GithubV2() WindowsFunctionAppSlotAuthSettingsV2GithubV2OutputReference {
	var returns WindowsFunctionAppSlotAuthSettingsV2GithubV2OutputReference
	_jsii_.Get(
		j,
		"githubV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GithubV2Input() *WindowsFunctionAppSlotAuthSettingsV2GithubV2 {
	var returns *WindowsFunctionAppSlotAuthSettingsV2GithubV2
	_jsii_.Get(
		j,
		"githubV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GoogleV2() WindowsFunctionAppSlotAuthSettingsV2GoogleV2OutputReference {
	var returns WindowsFunctionAppSlotAuthSettingsV2GoogleV2OutputReference
	_jsii_.Get(
		j,
		"googleV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GoogleV2Input() *WindowsFunctionAppSlotAuthSettingsV2GoogleV2 {
	var returns *WindowsFunctionAppSlotAuthSettingsV2GoogleV2
	_jsii_.Get(
		j,
		"googleV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) HttpRouteApiPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouteApiPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) HttpRouteApiPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouteApiPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) InternalValue() *WindowsFunctionAppSlotAuthSettingsV2 {
	var returns *WindowsFunctionAppSlotAuthSettingsV2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) Login() WindowsFunctionAppSlotAuthSettingsV2LoginOutputReference {
	var returns WindowsFunctionAppSlotAuthSettingsV2LoginOutputReference
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) LoginInput() *WindowsFunctionAppSlotAuthSettingsV2Login {
	var returns *WindowsFunctionAppSlotAuthSettingsV2Login
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) MicrosoftV2() WindowsFunctionAppSlotAuthSettingsV2MicrosoftV2OutputReference {
	var returns WindowsFunctionAppSlotAuthSettingsV2MicrosoftV2OutputReference
	_jsii_.Get(
		j,
		"microsoftV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) MicrosoftV2Input() *WindowsFunctionAppSlotAuthSettingsV2MicrosoftV2 {
	var returns *WindowsFunctionAppSlotAuthSettingsV2MicrosoftV2
	_jsii_.Get(
		j,
		"microsoftV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) RequireAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) RequireAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) RequireHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) RequireHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) TwitterV2() WindowsFunctionAppSlotAuthSettingsV2TwitterV2OutputReference {
	var returns WindowsFunctionAppSlotAuthSettingsV2TwitterV2OutputReference
	_jsii_.Get(
		j,
		"twitterV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) TwitterV2Input() *WindowsFunctionAppSlotAuthSettingsV2TwitterV2 {
	var returns *WindowsFunctionAppSlotAuthSettingsV2TwitterV2
	_jsii_.Get(
		j,
		"twitterV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) UnauthenticatedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) UnauthenticatedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedActionInput",
		&returns,
	)
	return returns
}


func NewWindowsFunctionAppSlotAuthSettingsV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsFunctionAppSlotAuthSettingsV2OutputReference {
	_init_.Initialize()

	if err := validateNewWindowsFunctionAppSlotAuthSettingsV2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsFunctionAppSlot.WindowsFunctionAppSlotAuthSettingsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsFunctionAppSlotAuthSettingsV2OutputReference_Override(w WindowsFunctionAppSlotAuthSettingsV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsFunctionAppSlot.WindowsFunctionAppSlotAuthSettingsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetAuthEnabled(val interface{}) {
	if err := j.validateSetAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetConfigFilePath(val *string) {
	if err := j.validateSetConfigFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configFilePath",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetDefaultProvider(val *string) {
	if err := j.validateSetDefaultProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProvider",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetExcludedPaths(val *[]*string) {
	if err := j.validateSetExcludedPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedPaths",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetForwardProxyConvention(val *string) {
	if err := j.validateSetForwardProxyConventionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardProxyConvention",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetForwardProxyCustomHostHeaderName(val *string) {
	if err := j.validateSetForwardProxyCustomHostHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardProxyCustomHostHeaderName",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetForwardProxyCustomSchemeHeaderName(val *string) {
	if err := j.validateSetForwardProxyCustomSchemeHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardProxyCustomSchemeHeaderName",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetHttpRouteApiPrefix(val *string) {
	if err := j.validateSetHttpRouteApiPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRouteApiPrefix",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetInternalValue(val *WindowsFunctionAppSlotAuthSettingsV2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetRequireAuthentication(val interface{}) {
	if err := j.validateSetRequireAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAuthentication",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetRequireHttps(val interface{}) {
	if err := j.validateSetRequireHttpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireHttps",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference)SetUnauthenticatedAction(val *string) {
	if err := j.validateSetUnauthenticatedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unauthenticatedAction",
		val,
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) PutActiveDirectoryV2(value *WindowsFunctionAppSlotAuthSettingsV2ActiveDirectoryV2) {
	if err := w.validatePutActiveDirectoryV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putActiveDirectoryV2",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) PutAppleV2(value *WindowsFunctionAppSlotAuthSettingsV2AppleV2) {
	if err := w.validatePutAppleV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAppleV2",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) PutAzureStaticWebAppV2(value *WindowsFunctionAppSlotAuthSettingsV2AzureStaticWebAppV2) {
	if err := w.validatePutAzureStaticWebAppV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAzureStaticWebAppV2",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) PutCustomOidcV2(value interface{}) {
	if err := w.validatePutCustomOidcV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCustomOidcV2",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) PutFacebookV2(value *WindowsFunctionAppSlotAuthSettingsV2FacebookV2) {
	if err := w.validatePutFacebookV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putFacebookV2",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) PutGithubV2(value *WindowsFunctionAppSlotAuthSettingsV2GithubV2) {
	if err := w.validatePutGithubV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGithubV2",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) PutGoogleV2(value *WindowsFunctionAppSlotAuthSettingsV2GoogleV2) {
	if err := w.validatePutGoogleV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGoogleV2",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) PutLogin(value *WindowsFunctionAppSlotAuthSettingsV2Login) {
	if err := w.validatePutLoginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLogin",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) PutMicrosoftV2(value *WindowsFunctionAppSlotAuthSettingsV2MicrosoftV2) {
	if err := w.validatePutMicrosoftV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMicrosoftV2",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) PutTwitterV2(value *WindowsFunctionAppSlotAuthSettingsV2TwitterV2) {
	if err := w.validatePutTwitterV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTwitterV2",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetActiveDirectoryV2() {
	_jsii_.InvokeVoid(
		w,
		"resetActiveDirectoryV2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetAppleV2() {
	_jsii_.InvokeVoid(
		w,
		"resetAppleV2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetAuthEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetAuthEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetAzureStaticWebAppV2() {
	_jsii_.InvokeVoid(
		w,
		"resetAzureStaticWebAppV2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetConfigFilePath() {
	_jsii_.InvokeVoid(
		w,
		"resetConfigFilePath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetCustomOidcV2() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomOidcV2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetDefaultProvider() {
	_jsii_.InvokeVoid(
		w,
		"resetDefaultProvider",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetExcludedPaths() {
	_jsii_.InvokeVoid(
		w,
		"resetExcludedPaths",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetFacebookV2() {
	_jsii_.InvokeVoid(
		w,
		"resetFacebookV2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetForwardProxyConvention() {
	_jsii_.InvokeVoid(
		w,
		"resetForwardProxyConvention",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetForwardProxyCustomHostHeaderName() {
	_jsii_.InvokeVoid(
		w,
		"resetForwardProxyCustomHostHeaderName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetForwardProxyCustomSchemeHeaderName() {
	_jsii_.InvokeVoid(
		w,
		"resetForwardProxyCustomSchemeHeaderName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetGithubV2() {
	_jsii_.InvokeVoid(
		w,
		"resetGithubV2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetGoogleV2() {
	_jsii_.InvokeVoid(
		w,
		"resetGoogleV2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetHttpRouteApiPrefix() {
	_jsii_.InvokeVoid(
		w,
		"resetHttpRouteApiPrefix",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetMicrosoftV2() {
	_jsii_.InvokeVoid(
		w,
		"resetMicrosoftV2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetRequireAuthentication() {
	_jsii_.InvokeVoid(
		w,
		"resetRequireAuthentication",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetRequireHttps() {
	_jsii_.InvokeVoid(
		w,
		"resetRequireHttps",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetTwitterV2() {
	_jsii_.InvokeVoid(
		w,
		"resetTwitterV2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ResetUnauthenticatedAction() {
	_jsii_.InvokeVoid(
		w,
		"resetUnauthenticatedAction",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WindowsFunctionAppSlotAuthSettingsV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

