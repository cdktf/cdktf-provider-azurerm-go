// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/apimanagementapi/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/api_management_api azurerm_api_management_api}.
type ApiManagementApi interface {
	cdktf.TerraformResource
	ApiManagementName() *string
	SetApiManagementName(val *string)
	ApiManagementNameInput() *string
	ApiType() *string
	SetApiType(val *string)
	ApiTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Contact() ApiManagementApiContactOutputReference
	ContactInput() *ApiManagementApiContact
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Import() ApiManagementApiImportOutputReference
	ImportInput() *ApiManagementApiImport
	IsCurrent() cdktf.IResolvable
	IsOnline() cdktf.IResolvable
	License() ApiManagementApiLicenseOutputReference
	LicenseInput() *ApiManagementApiLicense
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Oauth2Authorization() ApiManagementApiOauth2AuthorizationOutputReference
	Oauth2AuthorizationInput() *ApiManagementApiOauth2Authorization
	OpenidAuthentication() ApiManagementApiOpenidAuthenticationOutputReference
	OpenidAuthenticationInput() *ApiManagementApiOpenidAuthentication
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Revision() *string
	SetRevision(val *string)
	RevisionDescription() *string
	SetRevisionDescription(val *string)
	RevisionDescriptionInput() *string
	RevisionInput() *string
	ServiceUrl() *string
	SetServiceUrl(val *string)
	ServiceUrlInput() *string
	SoapPassThrough() interface{}
	SetSoapPassThrough(val interface{})
	SoapPassThroughInput() interface{}
	SourceApiId() *string
	SetSourceApiId(val *string)
	SourceApiIdInput() *string
	SubscriptionKeyParameterNames() ApiManagementApiSubscriptionKeyParameterNamesOutputReference
	SubscriptionKeyParameterNamesInput() *ApiManagementApiSubscriptionKeyParameterNames
	SubscriptionRequired() interface{}
	SetSubscriptionRequired(val interface{})
	SubscriptionRequiredInput() interface{}
	TermsOfServiceUrl() *string
	SetTermsOfServiceUrl(val *string)
	TermsOfServiceUrlInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApiManagementApiTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionDescription() *string
	SetVersionDescription(val *string)
	VersionDescriptionInput() *string
	VersionInput() *string
	VersionSetId() *string
	SetVersionSetId(val *string)
	VersionSetIdInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutContact(value *ApiManagementApiContact)
	PutImport(value *ApiManagementApiImport)
	PutLicense(value *ApiManagementApiLicense)
	PutOauth2Authorization(value *ApiManagementApiOauth2Authorization)
	PutOpenidAuthentication(value *ApiManagementApiOpenidAuthentication)
	PutSubscriptionKeyParameterNames(value *ApiManagementApiSubscriptionKeyParameterNames)
	PutTimeouts(value *ApiManagementApiTimeouts)
	ResetApiType()
	ResetContact()
	ResetDescription()
	ResetDisplayName()
	ResetId()
	ResetImport()
	ResetLicense()
	ResetOauth2Authorization()
	ResetOpenidAuthentication()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPath()
	ResetProtocols()
	ResetRevisionDescription()
	ResetServiceUrl()
	ResetSoapPassThrough()
	ResetSourceApiId()
	ResetSubscriptionKeyParameterNames()
	ResetSubscriptionRequired()
	ResetTermsOfServiceUrl()
	ResetTimeouts()
	ResetVersion()
	ResetVersionDescription()
	ResetVersionSetId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiManagementApi
type jsiiProxy_ApiManagementApi struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiManagementApi) ApiManagementName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ApiManagementNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ApiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ApiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Contact() ApiManagementApiContactOutputReference {
	var returns ApiManagementApiContactOutputReference
	_jsii_.Get(
		j,
		"contact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ContactInput() *ApiManagementApiContact {
	var returns *ApiManagementApiContact
	_jsii_.Get(
		j,
		"contactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Import() ApiManagementApiImportOutputReference {
	var returns ApiManagementApiImportOutputReference
	_jsii_.Get(
		j,
		"import",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ImportInput() *ApiManagementApiImport {
	var returns *ApiManagementApiImport
	_jsii_.Get(
		j,
		"importInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) IsCurrent() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isCurrent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) IsOnline() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isOnline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) License() ApiManagementApiLicenseOutputReference {
	var returns ApiManagementApiLicenseOutputReference
	_jsii_.Get(
		j,
		"license",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) LicenseInput() *ApiManagementApiLicense {
	var returns *ApiManagementApiLicense
	_jsii_.Get(
		j,
		"licenseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Oauth2Authorization() ApiManagementApiOauth2AuthorizationOutputReference {
	var returns ApiManagementApiOauth2AuthorizationOutputReference
	_jsii_.Get(
		j,
		"oauth2Authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Oauth2AuthorizationInput() *ApiManagementApiOauth2Authorization {
	var returns *ApiManagementApiOauth2Authorization
	_jsii_.Get(
		j,
		"oauth2AuthorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) OpenidAuthentication() ApiManagementApiOpenidAuthenticationOutputReference {
	var returns ApiManagementApiOpenidAuthenticationOutputReference
	_jsii_.Get(
		j,
		"openidAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) OpenidAuthenticationInput() *ApiManagementApiOpenidAuthentication {
	var returns *ApiManagementApiOpenidAuthentication
	_jsii_.Get(
		j,
		"openidAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Revision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) RevisionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) RevisionDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) RevisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) ServiceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) SoapPassThrough() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"soapPassThrough",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) SoapPassThroughInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"soapPassThroughInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) SourceApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) SourceApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) SubscriptionKeyParameterNames() ApiManagementApiSubscriptionKeyParameterNamesOutputReference {
	var returns ApiManagementApiSubscriptionKeyParameterNamesOutputReference
	_jsii_.Get(
		j,
		"subscriptionKeyParameterNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) SubscriptionKeyParameterNamesInput() *ApiManagementApiSubscriptionKeyParameterNames {
	var returns *ApiManagementApiSubscriptionKeyParameterNames
	_jsii_.Get(
		j,
		"subscriptionKeyParameterNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) SubscriptionRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscriptionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) SubscriptionRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscriptionRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) TermsOfServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"termsOfServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) TermsOfServiceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"termsOfServiceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Timeouts() ApiManagementApiTimeoutsOutputReference {
	var returns ApiManagementApiTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) VersionDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) VersionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApi) VersionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionSetIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/api_management_api azurerm_api_management_api} Resource.
func NewApiManagementApi(scope constructs.Construct, id *string, config *ApiManagementApiConfig) ApiManagementApi {
	_init_.Initialize()

	if err := validateNewApiManagementApiParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementApi{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementApi.ApiManagementApi",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/api_management_api azurerm_api_management_api} Resource.
func NewApiManagementApi_Override(a ApiManagementApi, scope constructs.Construct, id *string, config *ApiManagementApiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementApi.ApiManagementApi",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetApiManagementName(val *string) {
	if err := j.validateSetApiManagementNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiManagementName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetApiType(val *string) {
	if err := j.validateSetApiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiType",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetRevision(val *string) {
	if err := j.validateSetRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revision",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetRevisionDescription(val *string) {
	if err := j.validateSetRevisionDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revisionDescription",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetServiceUrl(val *string) {
	if err := j.validateSetServiceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceUrl",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetSoapPassThrough(val interface{}) {
	if err := j.validateSetSoapPassThroughParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"soapPassThrough",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetSourceApiId(val *string) {
	if err := j.validateSetSourceApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceApiId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetSubscriptionRequired(val interface{}) {
	if err := j.validateSetSubscriptionRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionRequired",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetTermsOfServiceUrl(val *string) {
	if err := j.validateSetTermsOfServiceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"termsOfServiceUrl",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetVersionDescription(val *string) {
	if err := j.validateSetVersionDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionDescription",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApi)SetVersionSetId(val *string) {
	if err := j.validateSetVersionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionSetId",
		val,
	)
}

// Generates CDKTF code for importing a ApiManagementApi resource upon running "cdktf plan <stack-name>".
func ApiManagementApi_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApiManagementApi_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementApi.ApiManagementApi",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ApiManagementApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementApi.ApiManagementApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagementApi_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementApi_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementApi.ApiManagementApi",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagementApi_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementApi_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementApi.ApiManagementApi",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiManagementApi_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.apiManagementApi.ApiManagementApi",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiManagementApi) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApiManagementApi) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiManagementApi) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiManagementApi) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementApi) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiManagementApi) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiManagementApi) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiManagementApi) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiManagementApi) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiManagementApi) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiManagementApi) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiManagementApi) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementApi) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApiManagementApi) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementApi) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiManagementApi) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApiManagementApi) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiManagementApi) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiManagementApi) PutContact(value *ApiManagementApiContact) {
	if err := a.validatePutContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContact",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApi) PutImport(value *ApiManagementApiImport) {
	if err := a.validatePutImportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImport",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApi) PutLicense(value *ApiManagementApiLicense) {
	if err := a.validatePutLicenseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLicense",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApi) PutOauth2Authorization(value *ApiManagementApiOauth2Authorization) {
	if err := a.validatePutOauth2AuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauth2Authorization",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApi) PutOpenidAuthentication(value *ApiManagementApiOpenidAuthentication) {
	if err := a.validatePutOpenidAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpenidAuthentication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApi) PutSubscriptionKeyParameterNames(value *ApiManagementApiSubscriptionKeyParameterNames) {
	if err := a.validatePutSubscriptionKeyParameterNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubscriptionKeyParameterNames",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApi) PutTimeouts(value *ApiManagementApiTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetApiType() {
	_jsii_.InvokeVoid(
		a,
		"resetApiType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetContact() {
	_jsii_.InvokeVoid(
		a,
		"resetContact",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetImport() {
	_jsii_.InvokeVoid(
		a,
		"resetImport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetLicense() {
	_jsii_.InvokeVoid(
		a,
		"resetLicense",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetOauth2Authorization() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth2Authorization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetOpenidAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenidAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetProtocols() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocols",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetRevisionDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetRevisionDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetServiceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetSoapPassThrough() {
	_jsii_.InvokeVoid(
		a,
		"resetSoapPassThrough",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetSourceApiId() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceApiId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetSubscriptionKeyParameterNames() {
	_jsii_.InvokeVoid(
		a,
		"resetSubscriptionKeyParameterNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetSubscriptionRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetSubscriptionRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetTermsOfServiceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetTermsOfServiceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetVersionDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetVersionDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) ResetVersionSetId() {
	_jsii_.InvokeVoid(
		a,
		"resetVersionSetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApi) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementApi) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementApi) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementApi) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementApi) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

