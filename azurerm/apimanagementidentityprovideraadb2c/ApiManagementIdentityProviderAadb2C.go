// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementidentityprovideraadb2c

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/apimanagementidentityprovideraadb2c/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/api_management_identity_provider_aadb2c azurerm_api_management_identity_provider_aadb2c}.
type ApiManagementIdentityProviderAadb2C interface {
	cdktf.TerraformResource
	AllowedTenant() *string
	SetAllowedTenant(val *string)
	AllowedTenantInput() *string
	ApiManagementName() *string
	SetApiManagementName(val *string)
	ApiManagementNameInput() *string
	Authority() *string
	SetAuthority(val *string)
	AuthorityInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientLibrary() *string
	SetClientLibrary(val *string)
	ClientLibraryInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PasswordResetPolicy() *string
	SetPasswordResetPolicy(val *string)
	PasswordResetPolicyInput() *string
	ProfileEditingPolicy() *string
	SetProfileEditingPolicy(val *string)
	ProfileEditingPolicyInput() *string
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
	SigninPolicy() *string
	SetSigninPolicy(val *string)
	SigninPolicyInput() *string
	SigninTenant() *string
	SetSigninTenant(val *string)
	SigninTenantInput() *string
	SignupPolicy() *string
	SetSignupPolicy(val *string)
	SignupPolicyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApiManagementIdentityProviderAadb2CTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *ApiManagementIdentityProviderAadb2CTimeouts)
	ResetClientLibrary()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordResetPolicy()
	ResetProfileEditingPolicy()
	ResetTimeouts()
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

// The jsii proxy struct for ApiManagementIdentityProviderAadb2C
type jsiiProxy_ApiManagementIdentityProviderAadb2C struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) AllowedTenant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedTenant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) AllowedTenantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedTenantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ApiManagementName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ApiManagementNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) Authority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) AuthorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ClientLibrary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientLibrary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ClientLibraryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientLibraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) PasswordResetPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordResetPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) PasswordResetPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordResetPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ProfileEditingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileEditingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ProfileEditingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileEditingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) SigninPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signinPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) SigninPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signinPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) SigninTenant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signinTenant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) SigninTenantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signinTenantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) SignupPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) SignupPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signupPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) Timeouts() ApiManagementIdentityProviderAadb2CTimeoutsOutputReference {
	var returns ApiManagementIdentityProviderAadb2CTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/api_management_identity_provider_aadb2c azurerm_api_management_identity_provider_aadb2c} Resource.
func NewApiManagementIdentityProviderAadb2C(scope constructs.Construct, id *string, config *ApiManagementIdentityProviderAadb2CConfig) ApiManagementIdentityProviderAadb2C {
	_init_.Initialize()

	if err := validateNewApiManagementIdentityProviderAadb2CParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementIdentityProviderAadb2C{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementIdentityProviderAadb2C.ApiManagementIdentityProviderAadb2C",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/api_management_identity_provider_aadb2c azurerm_api_management_identity_provider_aadb2c} Resource.
func NewApiManagementIdentityProviderAadb2C_Override(a ApiManagementIdentityProviderAadb2C, scope constructs.Construct, id *string, config *ApiManagementIdentityProviderAadb2CConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementIdentityProviderAadb2C.ApiManagementIdentityProviderAadb2C",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetAllowedTenant(val *string) {
	if err := j.validateSetAllowedTenantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedTenant",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetApiManagementName(val *string) {
	if err := j.validateSetApiManagementNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiManagementName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetAuthority(val *string) {
	if err := j.validateSetAuthorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authority",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetClientLibrary(val *string) {
	if err := j.validateSetClientLibraryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientLibrary",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetPasswordResetPolicy(val *string) {
	if err := j.validateSetPasswordResetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordResetPolicy",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetProfileEditingPolicy(val *string) {
	if err := j.validateSetProfileEditingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileEditingPolicy",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetSigninPolicy(val *string) {
	if err := j.validateSetSigninPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signinPolicy",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetSigninTenant(val *string) {
	if err := j.validateSetSigninTenantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signinTenant",
		val,
	)
}

func (j *jsiiProxy_ApiManagementIdentityProviderAadb2C)SetSignupPolicy(val *string) {
	if err := j.validateSetSignupPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signupPolicy",
		val,
	)
}

// Generates CDKTF code for importing a ApiManagementIdentityProviderAadb2C resource upon running "cdktf plan <stack-name>".
func ApiManagementIdentityProviderAadb2C_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApiManagementIdentityProviderAadb2C_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementIdentityProviderAadb2C.ApiManagementIdentityProviderAadb2C",
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
func ApiManagementIdentityProviderAadb2C_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementIdentityProviderAadb2C_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementIdentityProviderAadb2C.ApiManagementIdentityProviderAadb2C",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagementIdentityProviderAadb2C_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementIdentityProviderAadb2C_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementIdentityProviderAadb2C.ApiManagementIdentityProviderAadb2C",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagementIdentityProviderAadb2C_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementIdentityProviderAadb2C_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementIdentityProviderAadb2C.ApiManagementIdentityProviderAadb2C",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiManagementIdentityProviderAadb2C_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.apiManagementIdentityProviderAadb2C.ApiManagementIdentityProviderAadb2C",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) PutTimeouts(value *ApiManagementIdentityProviderAadb2CTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) ResetClientLibrary() {
	_jsii_.InvokeVoid(
		a,
		"resetClientLibrary",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) ResetPasswordResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordResetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) ResetProfileEditingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetProfileEditingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementIdentityProviderAadb2C) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

