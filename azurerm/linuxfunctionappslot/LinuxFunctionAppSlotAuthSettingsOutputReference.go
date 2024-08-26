// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/linuxfunctionappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxFunctionAppSlotAuthSettingsOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectory() LinuxFunctionAppSlotAuthSettingsActiveDirectoryOutputReference
	ActiveDirectoryInput() *LinuxFunctionAppSlotAuthSettingsActiveDirectory
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
	Facebook() LinuxFunctionAppSlotAuthSettingsFacebookOutputReference
	FacebookInput() *LinuxFunctionAppSlotAuthSettingsFacebook
	// Experimental.
	Fqn() *string
	Github() LinuxFunctionAppSlotAuthSettingsGithubOutputReference
	GithubInput() *LinuxFunctionAppSlotAuthSettingsGithub
	Google() LinuxFunctionAppSlotAuthSettingsGoogleOutputReference
	GoogleInput() *LinuxFunctionAppSlotAuthSettingsGoogle
	InternalValue() *LinuxFunctionAppSlotAuthSettings
	SetInternalValue(val *LinuxFunctionAppSlotAuthSettings)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	Microsoft() LinuxFunctionAppSlotAuthSettingsMicrosoftOutputReference
	MicrosoftInput() *LinuxFunctionAppSlotAuthSettingsMicrosoft
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
	Twitter() LinuxFunctionAppSlotAuthSettingsTwitterOutputReference
	TwitterInput() *LinuxFunctionAppSlotAuthSettingsTwitter
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
	PutActiveDirectory(value *LinuxFunctionAppSlotAuthSettingsActiveDirectory)
	PutFacebook(value *LinuxFunctionAppSlotAuthSettingsFacebook)
	PutGithub(value *LinuxFunctionAppSlotAuthSettingsGithub)
	PutGoogle(value *LinuxFunctionAppSlotAuthSettingsGoogle)
	PutMicrosoft(value *LinuxFunctionAppSlotAuthSettingsMicrosoft)
	PutTwitter(value *LinuxFunctionAppSlotAuthSettingsTwitter)
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

// The jsii proxy struct for LinuxFunctionAppSlotAuthSettingsOutputReference
type jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ActiveDirectory() LinuxFunctionAppSlotAuthSettingsActiveDirectoryOutputReference {
	var returns LinuxFunctionAppSlotAuthSettingsActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ActiveDirectoryInput() *LinuxFunctionAppSlotAuthSettingsActiveDirectory {
	var returns *LinuxFunctionAppSlotAuthSettingsActiveDirectory
	_jsii_.Get(
		j,
		"activeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) AdditionalLoginParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) AdditionalLoginParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) AllowedExternalRedirectUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) AllowedExternalRedirectUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) DefaultProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) DefaultProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) Facebook() LinuxFunctionAppSlotAuthSettingsFacebookOutputReference {
	var returns LinuxFunctionAppSlotAuthSettingsFacebookOutputReference
	_jsii_.Get(
		j,
		"facebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) FacebookInput() *LinuxFunctionAppSlotAuthSettingsFacebook {
	var returns *LinuxFunctionAppSlotAuthSettingsFacebook
	_jsii_.Get(
		j,
		"facebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) Github() LinuxFunctionAppSlotAuthSettingsGithubOutputReference {
	var returns LinuxFunctionAppSlotAuthSettingsGithubOutputReference
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) GithubInput() *LinuxFunctionAppSlotAuthSettingsGithub {
	var returns *LinuxFunctionAppSlotAuthSettingsGithub
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) Google() LinuxFunctionAppSlotAuthSettingsGoogleOutputReference {
	var returns LinuxFunctionAppSlotAuthSettingsGoogleOutputReference
	_jsii_.Get(
		j,
		"google",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) GoogleInput() *LinuxFunctionAppSlotAuthSettingsGoogle {
	var returns *LinuxFunctionAppSlotAuthSettingsGoogle
	_jsii_.Get(
		j,
		"googleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) InternalValue() *LinuxFunctionAppSlotAuthSettings {
	var returns *LinuxFunctionAppSlotAuthSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) Microsoft() LinuxFunctionAppSlotAuthSettingsMicrosoftOutputReference {
	var returns LinuxFunctionAppSlotAuthSettingsMicrosoftOutputReference
	_jsii_.Get(
		j,
		"microsoft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) MicrosoftInput() *LinuxFunctionAppSlotAuthSettingsMicrosoft {
	var returns *LinuxFunctionAppSlotAuthSettingsMicrosoft
	_jsii_.Get(
		j,
		"microsoftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) TokenRefreshExtensionHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) TokenRefreshExtensionHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) TokenStoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) TokenStoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) Twitter() LinuxFunctionAppSlotAuthSettingsTwitterOutputReference {
	var returns LinuxFunctionAppSlotAuthSettingsTwitterOutputReference
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) TwitterInput() *LinuxFunctionAppSlotAuthSettingsTwitter {
	var returns *LinuxFunctionAppSlotAuthSettingsTwitter
	_jsii_.Get(
		j,
		"twitterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) UnauthenticatedClientAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) UnauthenticatedClientActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientActionInput",
		&returns,
	)
	return returns
}


func NewLinuxFunctionAppSlotAuthSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LinuxFunctionAppSlotAuthSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewLinuxFunctionAppSlotAuthSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionAppSlot.LinuxFunctionAppSlotAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLinuxFunctionAppSlotAuthSettingsOutputReference_Override(l LinuxFunctionAppSlotAuthSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxFunctionAppSlot.LinuxFunctionAppSlotAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetAdditionalLoginParameters(val *map[string]*string) {
	if err := j.validateSetAdditionalLoginParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalLoginParameters",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetAllowedExternalRedirectUrls(val *[]*string) {
	if err := j.validateSetAllowedExternalRedirectUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedExternalRedirectUrls",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetDefaultProvider(val *string) {
	if err := j.validateSetDefaultProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProvider",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetInternalValue(val *LinuxFunctionAppSlotAuthSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetTokenRefreshExtensionHours(val *float64) {
	if err := j.validateSetTokenRefreshExtensionHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenRefreshExtensionHours",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetTokenStoreEnabled(val interface{}) {
	if err := j.validateSetTokenStoreEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenStoreEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference)SetUnauthenticatedClientAction(val *string) {
	if err := j.validateSetUnauthenticatedClientActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unauthenticatedClientAction",
		val,
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) PutActiveDirectory(value *LinuxFunctionAppSlotAuthSettingsActiveDirectory) {
	if err := l.validatePutActiveDirectoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putActiveDirectory",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) PutFacebook(value *LinuxFunctionAppSlotAuthSettingsFacebook) {
	if err := l.validatePutFacebookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFacebook",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) PutGithub(value *LinuxFunctionAppSlotAuthSettingsGithub) {
	if err := l.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGithub",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) PutGoogle(value *LinuxFunctionAppSlotAuthSettingsGoogle) {
	if err := l.validatePutGoogleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGoogle",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) PutMicrosoft(value *LinuxFunctionAppSlotAuthSettingsMicrosoft) {
	if err := l.validatePutMicrosoftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMicrosoft",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) PutTwitter(value *LinuxFunctionAppSlotAuthSettingsTwitter) {
	if err := l.validatePutTwitterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTwitter",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetActiveDirectory() {
	_jsii_.InvokeVoid(
		l,
		"resetActiveDirectory",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetAdditionalLoginParameters() {
	_jsii_.InvokeVoid(
		l,
		"resetAdditionalLoginParameters",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetAllowedExternalRedirectUrls() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowedExternalRedirectUrls",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetDefaultProvider() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultProvider",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetFacebook() {
	_jsii_.InvokeVoid(
		l,
		"resetFacebook",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		l,
		"resetGithub",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetGoogle() {
	_jsii_.InvokeVoid(
		l,
		"resetGoogle",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		l,
		"resetIssuer",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetMicrosoft() {
	_jsii_.InvokeVoid(
		l,
		"resetMicrosoft",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetTokenRefreshExtensionHours() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenRefreshExtensionHours",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetTokenStoreEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenStoreEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetTwitter() {
	_jsii_.InvokeVoid(
		l,
		"resetTwitter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ResetUnauthenticatedClientAction() {
	_jsii_.InvokeVoid(
		l,
		"resetUnauthenticatedClientAction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LinuxFunctionAppSlotAuthSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

