// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/functionappflexconsumption/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionAppFlexConsumptionAuthSettingsV2OutputReference interface {
	cdktf.ComplexObject
	ActiveDirectoryV2() FunctionAppFlexConsumptionAuthSettingsV2ActiveDirectoryV2OutputReference
	ActiveDirectoryV2Input() *FunctionAppFlexConsumptionAuthSettingsV2ActiveDirectoryV2
	AppleV2() FunctionAppFlexConsumptionAuthSettingsV2AppleV2OutputReference
	AppleV2Input() *FunctionAppFlexConsumptionAuthSettingsV2AppleV2
	AuthEnabled() interface{}
	SetAuthEnabled(val interface{})
	AuthEnabledInput() interface{}
	AzureStaticWebAppV2() FunctionAppFlexConsumptionAuthSettingsV2AzureStaticWebAppV2OutputReference
	AzureStaticWebAppV2Input() *FunctionAppFlexConsumptionAuthSettingsV2AzureStaticWebAppV2
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
	CustomOidcV2() FunctionAppFlexConsumptionAuthSettingsV2CustomOidcV2List
	CustomOidcV2Input() interface{}
	DefaultProvider() *string
	SetDefaultProvider(val *string)
	DefaultProviderInput() *string
	ExcludedPaths() *[]*string
	SetExcludedPaths(val *[]*string)
	ExcludedPathsInput() *[]*string
	FacebookV2() FunctionAppFlexConsumptionAuthSettingsV2FacebookV2OutputReference
	FacebookV2Input() *FunctionAppFlexConsumptionAuthSettingsV2FacebookV2
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
	GithubV2() FunctionAppFlexConsumptionAuthSettingsV2GithubV2OutputReference
	GithubV2Input() *FunctionAppFlexConsumptionAuthSettingsV2GithubV2
	GoogleV2() FunctionAppFlexConsumptionAuthSettingsV2GoogleV2OutputReference
	GoogleV2Input() *FunctionAppFlexConsumptionAuthSettingsV2GoogleV2
	HttpRouteApiPrefix() *string
	SetHttpRouteApiPrefix(val *string)
	HttpRouteApiPrefixInput() *string
	InternalValue() *FunctionAppFlexConsumptionAuthSettingsV2
	SetInternalValue(val *FunctionAppFlexConsumptionAuthSettingsV2)
	Login() FunctionAppFlexConsumptionAuthSettingsV2LoginOutputReference
	LoginInput() *FunctionAppFlexConsumptionAuthSettingsV2Login
	MicrosoftV2() FunctionAppFlexConsumptionAuthSettingsV2MicrosoftV2OutputReference
	MicrosoftV2Input() *FunctionAppFlexConsumptionAuthSettingsV2MicrosoftV2
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
	TwitterV2() FunctionAppFlexConsumptionAuthSettingsV2TwitterV2OutputReference
	TwitterV2Input() *FunctionAppFlexConsumptionAuthSettingsV2TwitterV2
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutActiveDirectoryV2(value *FunctionAppFlexConsumptionAuthSettingsV2ActiveDirectoryV2)
	PutAppleV2(value *FunctionAppFlexConsumptionAuthSettingsV2AppleV2)
	PutAzureStaticWebAppV2(value *FunctionAppFlexConsumptionAuthSettingsV2AzureStaticWebAppV2)
	PutCustomOidcV2(value interface{})
	PutFacebookV2(value *FunctionAppFlexConsumptionAuthSettingsV2FacebookV2)
	PutGithubV2(value *FunctionAppFlexConsumptionAuthSettingsV2GithubV2)
	PutGoogleV2(value *FunctionAppFlexConsumptionAuthSettingsV2GoogleV2)
	PutLogin(value *FunctionAppFlexConsumptionAuthSettingsV2Login)
	PutMicrosoftV2(value *FunctionAppFlexConsumptionAuthSettingsV2MicrosoftV2)
	PutTwitterV2(value *FunctionAppFlexConsumptionAuthSettingsV2TwitterV2)
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
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionAppFlexConsumptionAuthSettingsV2OutputReference
type jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ActiveDirectoryV2() FunctionAppFlexConsumptionAuthSettingsV2ActiveDirectoryV2OutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsV2ActiveDirectoryV2OutputReference
	_jsii_.Get(
		j,
		"activeDirectoryV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ActiveDirectoryV2Input() *FunctionAppFlexConsumptionAuthSettingsV2ActiveDirectoryV2 {
	var returns *FunctionAppFlexConsumptionAuthSettingsV2ActiveDirectoryV2
	_jsii_.Get(
		j,
		"activeDirectoryV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) AppleV2() FunctionAppFlexConsumptionAuthSettingsV2AppleV2OutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsV2AppleV2OutputReference
	_jsii_.Get(
		j,
		"appleV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) AppleV2Input() *FunctionAppFlexConsumptionAuthSettingsV2AppleV2 {
	var returns *FunctionAppFlexConsumptionAuthSettingsV2AppleV2
	_jsii_.Get(
		j,
		"appleV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) AuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) AuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) AzureStaticWebAppV2() FunctionAppFlexConsumptionAuthSettingsV2AzureStaticWebAppV2OutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsV2AzureStaticWebAppV2OutputReference
	_jsii_.Get(
		j,
		"azureStaticWebAppV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) AzureStaticWebAppV2Input() *FunctionAppFlexConsumptionAuthSettingsV2AzureStaticWebAppV2 {
	var returns *FunctionAppFlexConsumptionAuthSettingsV2AzureStaticWebAppV2
	_jsii_.Get(
		j,
		"azureStaticWebAppV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ConfigFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ConfigFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) CustomOidcV2() FunctionAppFlexConsumptionAuthSettingsV2CustomOidcV2List {
	var returns FunctionAppFlexConsumptionAuthSettingsV2CustomOidcV2List
	_jsii_.Get(
		j,
		"customOidcV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) CustomOidcV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOidcV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) DefaultProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) DefaultProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ExcludedPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ExcludedPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) FacebookV2() FunctionAppFlexConsumptionAuthSettingsV2FacebookV2OutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsV2FacebookV2OutputReference
	_jsii_.Get(
		j,
		"facebookV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) FacebookV2Input() *FunctionAppFlexConsumptionAuthSettingsV2FacebookV2 {
	var returns *FunctionAppFlexConsumptionAuthSettingsV2FacebookV2
	_jsii_.Get(
		j,
		"facebookV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ForwardProxyConvention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyConvention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ForwardProxyConventionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyConventionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ForwardProxyCustomHostHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomHostHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ForwardProxyCustomHostHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomHostHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ForwardProxyCustomSchemeHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomSchemeHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ForwardProxyCustomSchemeHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardProxyCustomSchemeHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GithubV2() FunctionAppFlexConsumptionAuthSettingsV2GithubV2OutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsV2GithubV2OutputReference
	_jsii_.Get(
		j,
		"githubV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GithubV2Input() *FunctionAppFlexConsumptionAuthSettingsV2GithubV2 {
	var returns *FunctionAppFlexConsumptionAuthSettingsV2GithubV2
	_jsii_.Get(
		j,
		"githubV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GoogleV2() FunctionAppFlexConsumptionAuthSettingsV2GoogleV2OutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsV2GoogleV2OutputReference
	_jsii_.Get(
		j,
		"googleV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GoogleV2Input() *FunctionAppFlexConsumptionAuthSettingsV2GoogleV2 {
	var returns *FunctionAppFlexConsumptionAuthSettingsV2GoogleV2
	_jsii_.Get(
		j,
		"googleV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) HttpRouteApiPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouteApiPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) HttpRouteApiPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRouteApiPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) InternalValue() *FunctionAppFlexConsumptionAuthSettingsV2 {
	var returns *FunctionAppFlexConsumptionAuthSettingsV2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) Login() FunctionAppFlexConsumptionAuthSettingsV2LoginOutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsV2LoginOutputReference
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) LoginInput() *FunctionAppFlexConsumptionAuthSettingsV2Login {
	var returns *FunctionAppFlexConsumptionAuthSettingsV2Login
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) MicrosoftV2() FunctionAppFlexConsumptionAuthSettingsV2MicrosoftV2OutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsV2MicrosoftV2OutputReference
	_jsii_.Get(
		j,
		"microsoftV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) MicrosoftV2Input() *FunctionAppFlexConsumptionAuthSettingsV2MicrosoftV2 {
	var returns *FunctionAppFlexConsumptionAuthSettingsV2MicrosoftV2
	_jsii_.Get(
		j,
		"microsoftV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) RequireAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) RequireAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) RequireHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) RequireHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) TwitterV2() FunctionAppFlexConsumptionAuthSettingsV2TwitterV2OutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsV2TwitterV2OutputReference
	_jsii_.Get(
		j,
		"twitterV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) TwitterV2Input() *FunctionAppFlexConsumptionAuthSettingsV2TwitterV2 {
	var returns *FunctionAppFlexConsumptionAuthSettingsV2TwitterV2
	_jsii_.Get(
		j,
		"twitterV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) UnauthenticatedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) UnauthenticatedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedActionInput",
		&returns,
	)
	return returns
}


func NewFunctionAppFlexConsumptionAuthSettingsV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FunctionAppFlexConsumptionAuthSettingsV2OutputReference {
	_init_.Initialize()

	if err := validateNewFunctionAppFlexConsumptionAuthSettingsV2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumptionAuthSettingsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFunctionAppFlexConsumptionAuthSettingsV2OutputReference_Override(f FunctionAppFlexConsumptionAuthSettingsV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumptionAuthSettingsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetAuthEnabled(val interface{}) {
	if err := j.validateSetAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetConfigFilePath(val *string) {
	if err := j.validateSetConfigFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configFilePath",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetDefaultProvider(val *string) {
	if err := j.validateSetDefaultProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProvider",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetExcludedPaths(val *[]*string) {
	if err := j.validateSetExcludedPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedPaths",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetForwardProxyConvention(val *string) {
	if err := j.validateSetForwardProxyConventionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardProxyConvention",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetForwardProxyCustomHostHeaderName(val *string) {
	if err := j.validateSetForwardProxyCustomHostHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardProxyCustomHostHeaderName",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetForwardProxyCustomSchemeHeaderName(val *string) {
	if err := j.validateSetForwardProxyCustomSchemeHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardProxyCustomSchemeHeaderName",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetHttpRouteApiPrefix(val *string) {
	if err := j.validateSetHttpRouteApiPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRouteApiPrefix",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetInternalValue(val *FunctionAppFlexConsumptionAuthSettingsV2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetRequireAuthentication(val interface{}) {
	if err := j.validateSetRequireAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAuthentication",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetRequireHttps(val interface{}) {
	if err := j.validateSetRequireHttpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireHttps",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference)SetUnauthenticatedAction(val *string) {
	if err := j.validateSetUnauthenticatedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unauthenticatedAction",
		val,
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) PutActiveDirectoryV2(value *FunctionAppFlexConsumptionAuthSettingsV2ActiveDirectoryV2) {
	if err := f.validatePutActiveDirectoryV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putActiveDirectoryV2",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) PutAppleV2(value *FunctionAppFlexConsumptionAuthSettingsV2AppleV2) {
	if err := f.validatePutAppleV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAppleV2",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) PutAzureStaticWebAppV2(value *FunctionAppFlexConsumptionAuthSettingsV2AzureStaticWebAppV2) {
	if err := f.validatePutAzureStaticWebAppV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAzureStaticWebAppV2",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) PutCustomOidcV2(value interface{}) {
	if err := f.validatePutCustomOidcV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCustomOidcV2",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) PutFacebookV2(value *FunctionAppFlexConsumptionAuthSettingsV2FacebookV2) {
	if err := f.validatePutFacebookV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFacebookV2",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) PutGithubV2(value *FunctionAppFlexConsumptionAuthSettingsV2GithubV2) {
	if err := f.validatePutGithubV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putGithubV2",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) PutGoogleV2(value *FunctionAppFlexConsumptionAuthSettingsV2GoogleV2) {
	if err := f.validatePutGoogleV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putGoogleV2",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) PutLogin(value *FunctionAppFlexConsumptionAuthSettingsV2Login) {
	if err := f.validatePutLoginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putLogin",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) PutMicrosoftV2(value *FunctionAppFlexConsumptionAuthSettingsV2MicrosoftV2) {
	if err := f.validatePutMicrosoftV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putMicrosoftV2",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) PutTwitterV2(value *FunctionAppFlexConsumptionAuthSettingsV2TwitterV2) {
	if err := f.validatePutTwitterV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTwitterV2",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetActiveDirectoryV2() {
	_jsii_.InvokeVoid(
		f,
		"resetActiveDirectoryV2",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetAppleV2() {
	_jsii_.InvokeVoid(
		f,
		"resetAppleV2",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetAuthEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetAuthEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetAzureStaticWebAppV2() {
	_jsii_.InvokeVoid(
		f,
		"resetAzureStaticWebAppV2",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetConfigFilePath() {
	_jsii_.InvokeVoid(
		f,
		"resetConfigFilePath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetCustomOidcV2() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomOidcV2",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetDefaultProvider() {
	_jsii_.InvokeVoid(
		f,
		"resetDefaultProvider",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetExcludedPaths() {
	_jsii_.InvokeVoid(
		f,
		"resetExcludedPaths",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetFacebookV2() {
	_jsii_.InvokeVoid(
		f,
		"resetFacebookV2",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetForwardProxyConvention() {
	_jsii_.InvokeVoid(
		f,
		"resetForwardProxyConvention",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetForwardProxyCustomHostHeaderName() {
	_jsii_.InvokeVoid(
		f,
		"resetForwardProxyCustomHostHeaderName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetForwardProxyCustomSchemeHeaderName() {
	_jsii_.InvokeVoid(
		f,
		"resetForwardProxyCustomSchemeHeaderName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetGithubV2() {
	_jsii_.InvokeVoid(
		f,
		"resetGithubV2",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetGoogleV2() {
	_jsii_.InvokeVoid(
		f,
		"resetGoogleV2",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetHttpRouteApiPrefix() {
	_jsii_.InvokeVoid(
		f,
		"resetHttpRouteApiPrefix",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetMicrosoftV2() {
	_jsii_.InvokeVoid(
		f,
		"resetMicrosoftV2",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetRequireAuthentication() {
	_jsii_.InvokeVoid(
		f,
		"resetRequireAuthentication",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetRequireHttps() {
	_jsii_.InvokeVoid(
		f,
		"resetRequireHttps",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetTwitterV2() {
	_jsii_.InvokeVoid(
		f,
		"resetTwitterV2",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ResetUnauthenticatedAction() {
	_jsii_.InvokeVoid(
		f,
		"resetUnauthenticatedAction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

