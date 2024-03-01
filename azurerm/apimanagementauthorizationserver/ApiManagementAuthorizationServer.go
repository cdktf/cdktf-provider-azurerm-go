// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementauthorizationserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/apimanagementauthorizationserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/api_management_authorization_server azurerm_api_management_authorization_server}.
type ApiManagementAuthorizationServer interface {
	cdktf.TerraformResource
	ApiManagementName() *string
	SetApiManagementName(val *string)
	ApiManagementNameInput() *string
	AuthorizationEndpoint() *string
	SetAuthorizationEndpoint(val *string)
	AuthorizationEndpointInput() *string
	AuthorizationMethods() *[]*string
	SetAuthorizationMethods(val *[]*string)
	AuthorizationMethodsInput() *[]*string
	BearerTokenSendingMethods() *[]*string
	SetBearerTokenSendingMethods(val *[]*string)
	BearerTokenSendingMethodsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientAuthenticationMethod() *[]*string
	SetClientAuthenticationMethod(val *[]*string)
	ClientAuthenticationMethodInput() *[]*string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientRegistrationEndpoint() *string
	SetClientRegistrationEndpoint(val *string)
	ClientRegistrationEndpointInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultScope() *string
	SetDefaultScope(val *string)
	DefaultScopeInput() *string
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
	GrantTypes() *[]*string
	SetGrantTypes(val *[]*string)
	GrantTypesInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	ResourceOwnerPassword() *string
	SetResourceOwnerPassword(val *string)
	ResourceOwnerPasswordInput() *string
	ResourceOwnerUsername() *string
	SetResourceOwnerUsername(val *string)
	ResourceOwnerUsernameInput() *string
	SupportState() interface{}
	SetSupportState(val interface{})
	SupportStateInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApiManagementAuthorizationServerTimeoutsOutputReference
	TimeoutsInput() interface{}
	TokenBodyParameter() ApiManagementAuthorizationServerTokenBodyParameterList
	TokenBodyParameterInput() interface{}
	TokenEndpoint() *string
	SetTokenEndpoint(val *string)
	TokenEndpointInput() *string
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
	PutTimeouts(value *ApiManagementAuthorizationServerTimeouts)
	PutTokenBodyParameter(value interface{})
	ResetBearerTokenSendingMethods()
	ResetClientAuthenticationMethod()
	ResetClientSecret()
	ResetDefaultScope()
	ResetDescription()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceOwnerPassword()
	ResetResourceOwnerUsername()
	ResetSupportState()
	ResetTimeouts()
	ResetTokenBodyParameter()
	ResetTokenEndpoint()
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

// The jsii proxy struct for ApiManagementAuthorizationServer
type jsiiProxy_ApiManagementAuthorizationServer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ApiManagementName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ApiManagementNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) AuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) AuthorizationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) AuthorizationMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) AuthorizationMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) BearerTokenSendingMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bearerTokenSendingMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) BearerTokenSendingMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bearerTokenSendingMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ClientAuthenticationMethod() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ClientAuthenticationMethodInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientAuthenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ClientRegistrationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientRegistrationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ClientRegistrationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientRegistrationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) DefaultScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) DefaultScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) GrantTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) GrantTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ResourceOwnerPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceOwnerPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ResourceOwnerPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceOwnerPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ResourceOwnerUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceOwnerUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) ResourceOwnerUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceOwnerUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) SupportState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) SupportStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) Timeouts() ApiManagementAuthorizationServerTimeoutsOutputReference {
	var returns ApiManagementAuthorizationServerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) TokenBodyParameter() ApiManagementAuthorizationServerTokenBodyParameterList {
	var returns ApiManagementAuthorizationServerTokenBodyParameterList
	_jsii_.Get(
		j,
		"tokenBodyParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) TokenBodyParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenBodyParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) TokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementAuthorizationServer) TokenEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/api_management_authorization_server azurerm_api_management_authorization_server} Resource.
func NewApiManagementAuthorizationServer(scope constructs.Construct, id *string, config *ApiManagementAuthorizationServerConfig) ApiManagementAuthorizationServer {
	_init_.Initialize()

	if err := validateNewApiManagementAuthorizationServerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementAuthorizationServer{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementAuthorizationServer.ApiManagementAuthorizationServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/api_management_authorization_server azurerm_api_management_authorization_server} Resource.
func NewApiManagementAuthorizationServer_Override(a ApiManagementAuthorizationServer, scope constructs.Construct, id *string, config *ApiManagementAuthorizationServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementAuthorizationServer.ApiManagementAuthorizationServer",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetApiManagementName(val *string) {
	if err := j.validateSetApiManagementNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiManagementName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetAuthorizationEndpoint(val *string) {
	if err := j.validateSetAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetAuthorizationMethods(val *[]*string) {
	if err := j.validateSetAuthorizationMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationMethods",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetBearerTokenSendingMethods(val *[]*string) {
	if err := j.validateSetBearerTokenSendingMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bearerTokenSendingMethods",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetClientAuthenticationMethod(val *[]*string) {
	if err := j.validateSetClientAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetClientRegistrationEndpoint(val *string) {
	if err := j.validateSetClientRegistrationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientRegistrationEndpoint",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetDefaultScope(val *string) {
	if err := j.validateSetDefaultScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultScope",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetGrantTypes(val *[]*string) {
	if err := j.validateSetGrantTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantTypes",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetResourceOwnerPassword(val *string) {
	if err := j.validateSetResourceOwnerPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceOwnerPassword",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetResourceOwnerUsername(val *string) {
	if err := j.validateSetResourceOwnerUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceOwnerUsername",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetSupportState(val interface{}) {
	if err := j.validateSetSupportStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportState",
		val,
	)
}

func (j *jsiiProxy_ApiManagementAuthorizationServer)SetTokenEndpoint(val *string) {
	if err := j.validateSetTokenEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpoint",
		val,
	)
}

// Generates CDKTF code for importing a ApiManagementAuthorizationServer resource upon running "cdktf plan <stack-name>".
func ApiManagementAuthorizationServer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApiManagementAuthorizationServer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementAuthorizationServer.ApiManagementAuthorizationServer",
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
func ApiManagementAuthorizationServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementAuthorizationServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementAuthorizationServer.ApiManagementAuthorizationServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagementAuthorizationServer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementAuthorizationServer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementAuthorizationServer.ApiManagementAuthorizationServer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagementAuthorizationServer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementAuthorizationServer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementAuthorizationServer.ApiManagementAuthorizationServer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiManagementAuthorizationServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.apiManagementAuthorizationServer.ApiManagementAuthorizationServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiManagementAuthorizationServer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementAuthorizationServer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiManagementAuthorizationServer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiManagementAuthorizationServer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiManagementAuthorizationServer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiManagementAuthorizationServer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiManagementAuthorizationServer) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiManagementAuthorizationServer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiManagementAuthorizationServer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementAuthorizationServer) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) PutTimeouts(value *ApiManagementAuthorizationServerTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) PutTokenBodyParameter(value interface{}) {
	if err := a.validatePutTokenBodyParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTokenBodyParameter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetBearerTokenSendingMethods() {
	_jsii_.InvokeVoid(
		a,
		"resetBearerTokenSendingMethods",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetClientAuthenticationMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetClientAuthenticationMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetDefaultScope() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetResourceOwnerPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceOwnerPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetResourceOwnerUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceOwnerUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetSupportState() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetTokenBodyParameter() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenBodyParameter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ResetTokenEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementAuthorizationServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

