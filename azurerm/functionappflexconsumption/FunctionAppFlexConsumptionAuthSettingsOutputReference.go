// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/functionappflexconsumption/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionAppFlexConsumptionAuthSettingsOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectory() FunctionAppFlexConsumptionAuthSettingsActiveDirectoryOutputReference
	ActiveDirectoryInput() *FunctionAppFlexConsumptionAuthSettingsActiveDirectory
	AdditionalLoginParameters() *map[string]*string
	SetAdditionalLoginParameters(val *map[string]*string)
	AdditionalLoginParametersInput() *map[string]*string
	AllowedExternalRedirectUrls() *[]*string
	SetAllowedExternalRedirectUrls(val *[]*string)
	AllowedExternalRedirectUrlsInput() *[]*string
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
	DefaultProvider() *string
	SetDefaultProvider(val *string)
	DefaultProviderInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Facebook() FunctionAppFlexConsumptionAuthSettingsFacebookOutputReference
	FacebookInput() *FunctionAppFlexConsumptionAuthSettingsFacebook
	// Experimental.
	Fqn() *string
	Github() FunctionAppFlexConsumptionAuthSettingsGithubOutputReference
	GithubInput() *FunctionAppFlexConsumptionAuthSettingsGithub
	Google() FunctionAppFlexConsumptionAuthSettingsGoogleOutputReference
	GoogleInput() *FunctionAppFlexConsumptionAuthSettingsGoogle
	InternalValue() *FunctionAppFlexConsumptionAuthSettings
	SetInternalValue(val *FunctionAppFlexConsumptionAuthSettings)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	Microsoft() FunctionAppFlexConsumptionAuthSettingsMicrosoftOutputReference
	MicrosoftInput() *FunctionAppFlexConsumptionAuthSettingsMicrosoft
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
	TokenRefreshExtensionHours() *float64
	SetTokenRefreshExtensionHours(val *float64)
	TokenRefreshExtensionHoursInput() *float64
	TokenStoreEnabled() interface{}
	SetTokenStoreEnabled(val interface{})
	TokenStoreEnabledInput() interface{}
	Twitter() FunctionAppFlexConsumptionAuthSettingsTwitterOutputReference
	TwitterInput() *FunctionAppFlexConsumptionAuthSettingsTwitter
	UnauthenticatedClientAction() *string
	SetUnauthenticatedClientAction(val *string)
	UnauthenticatedClientActionInput() *string
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
	PutActiveDirectory(value *FunctionAppFlexConsumptionAuthSettingsActiveDirectory)
	PutFacebook(value *FunctionAppFlexConsumptionAuthSettingsFacebook)
	PutGithub(value *FunctionAppFlexConsumptionAuthSettingsGithub)
	PutGoogle(value *FunctionAppFlexConsumptionAuthSettingsGoogle)
	PutMicrosoft(value *FunctionAppFlexConsumptionAuthSettingsMicrosoft)
	PutTwitter(value *FunctionAppFlexConsumptionAuthSettingsTwitter)
	ResetActiveDirectory()
	ResetAdditionalLoginParameters()
	ResetAllowedExternalRedirectUrls()
	ResetDefaultProvider()
	ResetFacebook()
	ResetGithub()
	ResetGoogle()
	ResetIssuer()
	ResetMicrosoft()
	ResetRuntimeVersion()
	ResetTokenRefreshExtensionHours()
	ResetTokenStoreEnabled()
	ResetTwitter()
	ResetUnauthenticatedClientAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionAppFlexConsumptionAuthSettingsOutputReference
type jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ActiveDirectory() FunctionAppFlexConsumptionAuthSettingsActiveDirectoryOutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ActiveDirectoryInput() *FunctionAppFlexConsumptionAuthSettingsActiveDirectory {
	var returns *FunctionAppFlexConsumptionAuthSettingsActiveDirectory
	_jsii_.Get(
		j,
		"activeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) AdditionalLoginParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) AdditionalLoginParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) AllowedExternalRedirectUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) AllowedExternalRedirectUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) DefaultProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) DefaultProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) Facebook() FunctionAppFlexConsumptionAuthSettingsFacebookOutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsFacebookOutputReference
	_jsii_.Get(
		j,
		"facebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) FacebookInput() *FunctionAppFlexConsumptionAuthSettingsFacebook {
	var returns *FunctionAppFlexConsumptionAuthSettingsFacebook
	_jsii_.Get(
		j,
		"facebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) Github() FunctionAppFlexConsumptionAuthSettingsGithubOutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsGithubOutputReference
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) GithubInput() *FunctionAppFlexConsumptionAuthSettingsGithub {
	var returns *FunctionAppFlexConsumptionAuthSettingsGithub
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) Google() FunctionAppFlexConsumptionAuthSettingsGoogleOutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsGoogleOutputReference
	_jsii_.Get(
		j,
		"google",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) GoogleInput() *FunctionAppFlexConsumptionAuthSettingsGoogle {
	var returns *FunctionAppFlexConsumptionAuthSettingsGoogle
	_jsii_.Get(
		j,
		"googleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) InternalValue() *FunctionAppFlexConsumptionAuthSettings {
	var returns *FunctionAppFlexConsumptionAuthSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) Microsoft() FunctionAppFlexConsumptionAuthSettingsMicrosoftOutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsMicrosoftOutputReference
	_jsii_.Get(
		j,
		"microsoft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) MicrosoftInput() *FunctionAppFlexConsumptionAuthSettingsMicrosoft {
	var returns *FunctionAppFlexConsumptionAuthSettingsMicrosoft
	_jsii_.Get(
		j,
		"microsoftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) TokenRefreshExtensionHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) TokenRefreshExtensionHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) TokenStoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) TokenStoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) Twitter() FunctionAppFlexConsumptionAuthSettingsTwitterOutputReference {
	var returns FunctionAppFlexConsumptionAuthSettingsTwitterOutputReference
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) TwitterInput() *FunctionAppFlexConsumptionAuthSettingsTwitter {
	var returns *FunctionAppFlexConsumptionAuthSettingsTwitter
	_jsii_.Get(
		j,
		"twitterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) UnauthenticatedClientAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) UnauthenticatedClientActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientActionInput",
		&returns,
	)
	return returns
}


func NewFunctionAppFlexConsumptionAuthSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FunctionAppFlexConsumptionAuthSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewFunctionAppFlexConsumptionAuthSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumptionAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFunctionAppFlexConsumptionAuthSettingsOutputReference_Override(f FunctionAppFlexConsumptionAuthSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionAppFlexConsumption.FunctionAppFlexConsumptionAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetAdditionalLoginParameters(val *map[string]*string) {
	if err := j.validateSetAdditionalLoginParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalLoginParameters",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetAllowedExternalRedirectUrls(val *[]*string) {
	if err := j.validateSetAllowedExternalRedirectUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedExternalRedirectUrls",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetDefaultProvider(val *string) {
	if err := j.validateSetDefaultProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProvider",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetInternalValue(val *FunctionAppFlexConsumptionAuthSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetTokenRefreshExtensionHours(val *float64) {
	if err := j.validateSetTokenRefreshExtensionHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenRefreshExtensionHours",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetTokenStoreEnabled(val interface{}) {
	if err := j.validateSetTokenStoreEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenStoreEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference)SetUnauthenticatedClientAction(val *string) {
	if err := j.validateSetUnauthenticatedClientActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unauthenticatedClientAction",
		val,
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) PutActiveDirectory(value *FunctionAppFlexConsumptionAuthSettingsActiveDirectory) {
	if err := f.validatePutActiveDirectoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putActiveDirectory",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) PutFacebook(value *FunctionAppFlexConsumptionAuthSettingsFacebook) {
	if err := f.validatePutFacebookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFacebook",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) PutGithub(value *FunctionAppFlexConsumptionAuthSettingsGithub) {
	if err := f.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putGithub",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) PutGoogle(value *FunctionAppFlexConsumptionAuthSettingsGoogle) {
	if err := f.validatePutGoogleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putGoogle",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) PutMicrosoft(value *FunctionAppFlexConsumptionAuthSettingsMicrosoft) {
	if err := f.validatePutMicrosoftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putMicrosoft",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) PutTwitter(value *FunctionAppFlexConsumptionAuthSettingsTwitter) {
	if err := f.validatePutTwitterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTwitter",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetActiveDirectory() {
	_jsii_.InvokeVoid(
		f,
		"resetActiveDirectory",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetAdditionalLoginParameters() {
	_jsii_.InvokeVoid(
		f,
		"resetAdditionalLoginParameters",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetAllowedExternalRedirectUrls() {
	_jsii_.InvokeVoid(
		f,
		"resetAllowedExternalRedirectUrls",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetDefaultProvider() {
	_jsii_.InvokeVoid(
		f,
		"resetDefaultProvider",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetFacebook() {
	_jsii_.InvokeVoid(
		f,
		"resetFacebook",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		f,
		"resetGithub",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetGoogle() {
	_jsii_.InvokeVoid(
		f,
		"resetGoogle",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		f,
		"resetIssuer",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetMicrosoft() {
	_jsii_.InvokeVoid(
		f,
		"resetMicrosoft",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetTokenRefreshExtensionHours() {
	_jsii_.InvokeVoid(
		f,
		"resetTokenRefreshExtensionHours",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetTokenStoreEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetTokenStoreEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetTwitter() {
	_jsii_.InvokeVoid(
		f,
		"resetTwitter",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ResetUnauthenticatedClientAction() {
	_jsii_.InvokeVoid(
		f,
		"resetUnauthenticatedClientAction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppFlexConsumptionAuthSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

