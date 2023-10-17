// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/keyvault/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/key_vault azurerm_key_vault}.
type KeyVault interface {
	cdktf.TerraformResource
	AccessPolicy() KeyVaultAccessPolicyList
	AccessPolicyInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Contact() KeyVaultContactList
	ContactInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnabledForDeployment() interface{}
	SetEnabledForDeployment(val interface{})
	EnabledForDeploymentInput() interface{}
	EnabledForDiskEncryption() interface{}
	SetEnabledForDiskEncryption(val interface{})
	EnabledForDiskEncryptionInput() interface{}
	EnabledForTemplateDeployment() interface{}
	SetEnabledForTemplateDeployment(val interface{})
	EnabledForTemplateDeploymentInput() interface{}
	EnableRbacAuthorization() interface{}
	SetEnableRbacAuthorization(val interface{})
	EnableRbacAuthorizationInput() interface{}
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
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkAcls() KeyVaultNetworkAclsOutputReference
	NetworkAclsInput() *KeyVaultNetworkAcls
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
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	PurgeProtectionEnabled() interface{}
	SetPurgeProtectionEnabled(val interface{})
	PurgeProtectionEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	SoftDeleteRetentionDays() *float64
	SetSoftDeleteRetentionDays(val *float64)
	SoftDeleteRetentionDaysInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KeyVaultTimeoutsOutputReference
	TimeoutsInput() interface{}
	VaultUri() *string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAccessPolicy(value interface{})
	PutContact(value interface{})
	PutNetworkAcls(value *KeyVaultNetworkAcls)
	PutTimeouts(value *KeyVaultTimeouts)
	ResetAccessPolicy()
	ResetContact()
	ResetEnabledForDeployment()
	ResetEnabledForDiskEncryption()
	ResetEnabledForTemplateDeployment()
	ResetEnableRbacAuthorization()
	ResetId()
	ResetNetworkAcls()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetPurgeProtectionEnabled()
	ResetSoftDeleteRetentionDays()
	ResetTags()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for KeyVault
type jsiiProxy_KeyVault struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KeyVault) AccessPolicy() KeyVaultAccessPolicyList {
	var returns KeyVaultAccessPolicyList
	_jsii_.Get(
		j,
		"accessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) AccessPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Contact() KeyVaultContactList {
	var returns KeyVaultContactList
	_jsii_.Get(
		j,
		"contact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) ContactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) EnabledForDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledForDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) EnabledForDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledForDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) EnabledForDiskEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledForDiskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) EnabledForDiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledForDiskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) EnabledForTemplateDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledForTemplateDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) EnabledForTemplateDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledForTemplateDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) EnableRbacAuthorization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRbacAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) EnableRbacAuthorizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRbacAuthorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) NetworkAcls() KeyVaultNetworkAclsOutputReference {
	var returns KeyVaultNetworkAclsOutputReference
	_jsii_.Get(
		j,
		"networkAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) NetworkAclsInput() *KeyVaultNetworkAcls {
	var returns *KeyVaultNetworkAcls
	_jsii_.Get(
		j,
		"networkAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) PurgeProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purgeProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) PurgeProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purgeProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) SoftDeleteRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"softDeleteRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) SoftDeleteRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"softDeleteRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) Timeouts() KeyVaultTimeoutsOutputReference {
	var returns KeyVaultTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVault) VaultUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultUri",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/key_vault azurerm_key_vault} Resource.
func NewKeyVault(scope constructs.Construct, id *string, config *KeyVaultConfig) KeyVault {
	_init_.Initialize()

	if err := validateNewKeyVaultParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyVault{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVault.KeyVault",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/key_vault azurerm_key_vault} Resource.
func NewKeyVault_Override(k KeyVault, scope constructs.Construct, id *string, config *KeyVaultConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVault.KeyVault",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KeyVault)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetEnabledForDeployment(val interface{}) {
	if err := j.validateSetEnabledForDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledForDeployment",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetEnabledForDiskEncryption(val interface{}) {
	if err := j.validateSetEnabledForDiskEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledForDiskEncryption",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetEnabledForTemplateDeployment(val interface{}) {
	if err := j.validateSetEnabledForTemplateDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledForTemplateDeployment",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetEnableRbacAuthorization(val interface{}) {
	if err := j.validateSetEnableRbacAuthorizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRbacAuthorization",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetPurgeProtectionEnabled(val interface{}) {
	if err := j.validateSetPurgeProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purgeProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetSoftDeleteRetentionDays(val *float64) {
	if err := j.validateSetSoftDeleteRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"softDeleteRetentionDays",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KeyVault)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

// Generates CDKTF code for importing a KeyVault resource upon running "cdktf plan <stack-name>".
func KeyVault_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKeyVault_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.keyVault.KeyVault",
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
func KeyVault_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyVault_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.keyVault.KeyVault",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KeyVault_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyVault_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.keyVault.KeyVault",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KeyVault_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyVault_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.keyVault.KeyVault",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KeyVault_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.keyVault.KeyVault",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KeyVault) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KeyVault) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KeyVault) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KeyVault) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KeyVault) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KeyVault) PutAccessPolicy(value interface{}) {
	if err := k.validatePutAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAccessPolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVault) PutContact(value interface{}) {
	if err := k.validatePutContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putContact",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVault) PutNetworkAcls(value *KeyVaultNetworkAcls) {
	if err := k.validatePutNetworkAclsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putNetworkAcls",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVault) PutTimeouts(value *KeyVaultTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVault) ResetAccessPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetAccessPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetContact() {
	_jsii_.InvokeVoid(
		k,
		"resetContact",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetEnabledForDeployment() {
	_jsii_.InvokeVoid(
		k,
		"resetEnabledForDeployment",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetEnabledForDiskEncryption() {
	_jsii_.InvokeVoid(
		k,
		"resetEnabledForDiskEncryption",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetEnabledForTemplateDeployment() {
	_jsii_.InvokeVoid(
		k,
		"resetEnabledForTemplateDeployment",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetEnableRbacAuthorization() {
	_jsii_.InvokeVoid(
		k,
		"resetEnableRbacAuthorization",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetNetworkAcls() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkAcls",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetPurgeProtectionEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPurgeProtectionEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetSoftDeleteRetentionDays() {
	_jsii_.InvokeVoid(
		k,
		"resetSoftDeleteRetentionDays",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVault) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVault) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

