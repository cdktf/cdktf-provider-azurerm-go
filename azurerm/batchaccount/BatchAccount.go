package batchaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/batchaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/batch_account azurerm_batch_account}.
type BatchAccount interface {
	cdktf.TerraformResource
	AccountEndpoint() *string
	AllowedAuthenticationModes() *[]*string
	SetAllowedAuthenticationModes(val *[]*string)
	AllowedAuthenticationModesInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Encryption() BatchAccountEncryptionList
	EncryptionInput() interface{}
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
	Identity() BatchAccountIdentityOutputReference
	IdentityInput() *BatchAccountIdentity
	IdInput() *string
	KeyVaultReference() BatchAccountKeyVaultReferenceOutputReference
	KeyVaultReferenceInput() *BatchAccountKeyVaultReference
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
	// The tree node.
	Node() constructs.Node
	PoolAllocationMode() *string
	SetPoolAllocationMode(val *string)
	PoolAllocationModeInput() *string
	PrimaryAccessKey() *string
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
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SecondaryAccessKey() *string
	StorageAccountAuthenticationMode() *string
	SetStorageAccountAuthenticationMode(val *string)
	StorageAccountAuthenticationModeInput() *string
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	StorageAccountNodeIdentity() *string
	SetStorageAccountNodeIdentity(val *string)
	StorageAccountNodeIdentityInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BatchAccountTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutEncryption(value interface{})
	PutIdentity(value *BatchAccountIdentity)
	PutKeyVaultReference(value *BatchAccountKeyVaultReference)
	PutTimeouts(value *BatchAccountTimeouts)
	ResetAllowedAuthenticationModes()
	ResetEncryption()
	ResetId()
	ResetIdentity()
	ResetKeyVaultReference()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPoolAllocationMode()
	ResetPublicNetworkAccessEnabled()
	ResetStorageAccountAuthenticationMode()
	ResetStorageAccountId()
	ResetStorageAccountNodeIdentity()
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

// The jsii proxy struct for BatchAccount
type jsiiProxy_BatchAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BatchAccount) AccountEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) AllowedAuthenticationModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAuthenticationModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) AllowedAuthenticationModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAuthenticationModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Encryption() BatchAccountEncryptionList {
	var returns BatchAccountEncryptionList
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) EncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Identity() BatchAccountIdentityOutputReference {
	var returns BatchAccountIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) IdentityInput() *BatchAccountIdentity {
	var returns *BatchAccountIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) KeyVaultReference() BatchAccountKeyVaultReferenceOutputReference {
	var returns BatchAccountKeyVaultReferenceOutputReference
	_jsii_.Get(
		j,
		"keyVaultReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) KeyVaultReferenceInput() *BatchAccountKeyVaultReference {
	var returns *BatchAccountKeyVaultReference
	_jsii_.Get(
		j,
		"keyVaultReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) PoolAllocationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolAllocationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) PoolAllocationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolAllocationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) PrimaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) SecondaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) StorageAccountAuthenticationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAuthenticationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) StorageAccountAuthenticationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAuthenticationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) StorageAccountNodeIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountNodeIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) StorageAccountNodeIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountNodeIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) Timeouts() BatchAccountTimeoutsOutputReference {
	var returns BatchAccountTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchAccount) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/batch_account azurerm_batch_account} Resource.
func NewBatchAccount(scope constructs.Construct, id *string, config *BatchAccountConfig) BatchAccount {
	_init_.Initialize()

	if err := validateNewBatchAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchAccount{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchAccount.BatchAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/batch_account azurerm_batch_account} Resource.
func NewBatchAccount_Override(b BatchAccount, scope constructs.Construct, id *string, config *BatchAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.batchAccount.BatchAccount",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BatchAccount)SetAllowedAuthenticationModes(val *[]*string) {
	if err := j.validateSetAllowedAuthenticationModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAuthenticationModes",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetPoolAllocationMode(val *string) {
	if err := j.validateSetPoolAllocationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolAllocationMode",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetStorageAccountAuthenticationMode(val *string) {
	if err := j.validateSetStorageAccountAuthenticationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountAuthenticationMode",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetStorageAccountId(val *string) {
	if err := j.validateSetStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetStorageAccountNodeIdentity(val *string) {
	if err := j.validateSetStorageAccountNodeIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountNodeIdentity",
		val,
	)
}

func (j *jsiiProxy_BatchAccount)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
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
func BatchAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBatchAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.batchAccount.BatchAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BatchAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBatchAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.batchAccount.BatchAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BatchAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBatchAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.batchAccount.BatchAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BatchAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.batchAccount.BatchAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BatchAccount) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BatchAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BatchAccount) PutEncryption(value interface{}) {
	if err := b.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEncryption",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchAccount) PutIdentity(value *BatchAccountIdentity) {
	if err := b.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putIdentity",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchAccount) PutKeyVaultReference(value *BatchAccountKeyVaultReference) {
	if err := b.validatePutKeyVaultReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putKeyVaultReference",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchAccount) PutTimeouts(value *BatchAccountTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchAccount) ResetAllowedAuthenticationModes() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedAuthenticationModes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetEncryption() {
	_jsii_.InvokeVoid(
		b,
		"resetEncryption",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetIdentity() {
	_jsii_.InvokeVoid(
		b,
		"resetIdentity",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetKeyVaultReference() {
	_jsii_.InvokeVoid(
		b,
		"resetKeyVaultReference",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetPoolAllocationMode() {
	_jsii_.InvokeVoid(
		b,
		"resetPoolAllocationMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetStorageAccountAuthenticationMode() {
	_jsii_.InvokeVoid(
		b,
		"resetStorageAccountAuthenticationMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetStorageAccountId() {
	_jsii_.InvokeVoid(
		b,
		"resetStorageAccountId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetStorageAccountNodeIdentity() {
	_jsii_.InvokeVoid(
		b,
		"resetStorageAccountNodeIdentity",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

