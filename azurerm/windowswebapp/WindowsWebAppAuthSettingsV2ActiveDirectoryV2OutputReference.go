// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/windowswebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference interface {
	cdktf.ComplexObject
	AllowedApplications() *[]*string
	SetAllowedApplications(val *[]*string)
	AllowedApplicationsInput() *[]*string
	AllowedAudiences() *[]*string
	SetAllowedAudiences(val *[]*string)
	AllowedAudiencesInput() *[]*string
	AllowedGroups() *[]*string
	SetAllowedGroups(val *[]*string)
	AllowedGroupsInput() *[]*string
	AllowedIdentities() *[]*string
	SetAllowedIdentities(val *[]*string)
	AllowedIdentitiesInput() *[]*string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecretCertificateThumbprint() *string
	SetClientSecretCertificateThumbprint(val *string)
	ClientSecretCertificateThumbprintInput() *string
	ClientSecretSettingName() *string
	SetClientSecretSettingName(val *string)
	ClientSecretSettingNameInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppAuthSettingsV2ActiveDirectoryV2
	SetInternalValue(val *WindowsWebAppAuthSettingsV2ActiveDirectoryV2)
	JwtAllowedClientApplications() *[]*string
	SetJwtAllowedClientApplications(val *[]*string)
	JwtAllowedClientApplicationsInput() *[]*string
	JwtAllowedGroups() *[]*string
	SetJwtAllowedGroups(val *[]*string)
	JwtAllowedGroupsInput() *[]*string
	LoginParameters() *map[string]*string
	SetLoginParameters(val *map[string]*string)
	LoginParametersInput() *map[string]*string
	TenantAuthEndpoint() *string
	SetTenantAuthEndpoint(val *string)
	TenantAuthEndpointInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WwwAuthenticationDisabled() interface{}
	SetWwwAuthenticationDisabled(val interface{})
	WwwAuthenticationDisabledInput() interface{}
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
	ResetAllowedApplications()
	ResetAllowedAudiences()
	ResetAllowedGroups()
	ResetAllowedIdentities()
	ResetClientSecretCertificateThumbprint()
	ResetClientSecretSettingName()
	ResetJwtAllowedClientApplications()
	ResetJwtAllowedGroups()
	ResetLoginParameters()
	ResetWwwAuthenticationDisabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference
type jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedApplications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedApplicationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedAudiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedAudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedIdentities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIdentities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) AllowedIdentitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIdentitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ClientSecretCertificateThumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretCertificateThumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ClientSecretCertificateThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretCertificateThumbprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ClientSecretSettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ClientSecretSettingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretSettingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) InternalValue() *WindowsWebAppAuthSettingsV2ActiveDirectoryV2 {
	var returns *WindowsWebAppAuthSettingsV2ActiveDirectoryV2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) JwtAllowedClientApplications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwtAllowedClientApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) JwtAllowedClientApplicationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwtAllowedClientApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) JwtAllowedGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwtAllowedGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) JwtAllowedGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwtAllowedGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) LoginParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"loginParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) LoginParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"loginParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) TenantAuthEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantAuthEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) TenantAuthEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantAuthEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) WwwAuthenticationDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wwwAuthenticationDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) WwwAuthenticationDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wwwAuthenticationDisabledInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference {
	_init_.Initialize()

	if err := validateNewWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference_Override(w WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetAllowedApplications(val *[]*string) {
	if err := j.validateSetAllowedApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedApplications",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetAllowedAudiences(val *[]*string) {
	if err := j.validateSetAllowedAudiencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAudiences",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetAllowedGroups(val *[]*string) {
	if err := j.validateSetAllowedGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedGroups",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetAllowedIdentities(val *[]*string) {
	if err := j.validateSetAllowedIdentitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIdentities",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetClientSecretCertificateThumbprint(val *string) {
	if err := j.validateSetClientSecretCertificateThumbprintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretCertificateThumbprint",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetClientSecretSettingName(val *string) {
	if err := j.validateSetClientSecretSettingNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretSettingName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetInternalValue(val *WindowsWebAppAuthSettingsV2ActiveDirectoryV2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetJwtAllowedClientApplications(val *[]*string) {
	if err := j.validateSetJwtAllowedClientApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtAllowedClientApplications",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetJwtAllowedGroups(val *[]*string) {
	if err := j.validateSetJwtAllowedGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtAllowedGroups",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetLoginParameters(val *map[string]*string) {
	if err := j.validateSetLoginParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginParameters",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetTenantAuthEndpoint(val *string) {
	if err := j.validateSetTenantAuthEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantAuthEndpoint",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference)SetWwwAuthenticationDisabled(val interface{}) {
	if err := j.validateSetWwwAuthenticationDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wwwAuthenticationDisabled",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ResetAllowedApplications() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowedApplications",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ResetAllowedAudiences() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowedAudiences",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ResetAllowedGroups() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowedGroups",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ResetAllowedIdentities() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowedIdentities",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ResetClientSecretCertificateThumbprint() {
	_jsii_.InvokeVoid(
		w,
		"resetClientSecretCertificateThumbprint",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ResetClientSecretSettingName() {
	_jsii_.InvokeVoid(
		w,
		"resetClientSecretSettingName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ResetJwtAllowedClientApplications() {
	_jsii_.InvokeVoid(
		w,
		"resetJwtAllowedClientApplications",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ResetJwtAllowedGroups() {
	_jsii_.InvokeVoid(
		w,
		"resetJwtAllowedGroups",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ResetLoginParameters() {
	_jsii_.InvokeVoid(
		w,
		"resetLoginParameters",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ResetWwwAuthenticationDisabled() {
	_jsii_.InvokeVoid(
		w,
		"resetWwwAuthenticationDisabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2ActiveDirectoryV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

