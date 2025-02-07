// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccountcustomermanagedkey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/storageaccountcustomermanagedkey/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/storage_account_customer_managed_key azurerm_storage_account_customer_managed_key}.
type StorageAccountCustomerManagedKeyA interface {
	cdktf.TerraformResource
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
	FederatedIdentityClientId() *string
	SetFederatedIdentityClientId(val *string)
	FederatedIdentityClientIdInput() *string
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
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
	KeyVaultUri() *string
	SetKeyVaultUri(val *string)
	KeyVaultUriInput() *string
	KeyVersion() *string
	SetKeyVersion(val *string)
	KeyVersionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManagedHsmKeyId() *string
	SetManagedHsmKeyId(val *string)
	ManagedHsmKeyIdInput() *string
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
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StorageAccountCustomerManagedKeyTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserAssignedIdentityId() *string
	SetUserAssignedIdentityId(val *string)
	UserAssignedIdentityIdInput() *string
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
	PutTimeouts(value *StorageAccountCustomerManagedKeyTimeouts)
	ResetFederatedIdentityClientId()
	ResetId()
	ResetKeyVaultId()
	ResetKeyVaultUri()
	ResetKeyVersion()
	ResetManagedHsmKeyId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetUserAssignedIdentityId()
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

// The jsii proxy struct for StorageAccountCustomerManagedKeyA
type jsiiProxy_StorageAccountCustomerManagedKeyA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) FederatedIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federatedIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) FederatedIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federatedIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) KeyVaultUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) KeyVaultUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) KeyVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) KeyVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) ManagedHsmKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedHsmKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) ManagedHsmKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedHsmKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) Timeouts() StorageAccountCustomerManagedKeyTimeoutsOutputReference {
	var returns StorageAccountCustomerManagedKeyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) UserAssignedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA) UserAssignedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/storage_account_customer_managed_key azurerm_storage_account_customer_managed_key} Resource.
func NewStorageAccountCustomerManagedKeyA(scope constructs.Construct, id *string, config *StorageAccountCustomerManagedKeyAConfig) StorageAccountCustomerManagedKeyA {
	_init_.Initialize()

	if err := validateNewStorageAccountCustomerManagedKeyAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageAccountCustomerManagedKeyA{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccountCustomerManagedKey.StorageAccountCustomerManagedKeyA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/storage_account_customer_managed_key azurerm_storage_account_customer_managed_key} Resource.
func NewStorageAccountCustomerManagedKeyA_Override(s StorageAccountCustomerManagedKeyA, scope constructs.Construct, id *string, config *StorageAccountCustomerManagedKeyAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageAccountCustomerManagedKey.StorageAccountCustomerManagedKeyA",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetFederatedIdentityClientId(val *string) {
	if err := j.validateSetFederatedIdentityClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"federatedIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetKeyVaultId(val *string) {
	if err := j.validateSetKeyVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetKeyVaultUri(val *string) {
	if err := j.validateSetKeyVaultUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultUri",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetKeyVersion(val *string) {
	if err := j.validateSetKeyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVersion",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetManagedHsmKeyId(val *string) {
	if err := j.validateSetManagedHsmKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedHsmKeyId",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetStorageAccountId(val *string) {
	if err := j.validateSetStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_StorageAccountCustomerManagedKeyA)SetUserAssignedIdentityId(val *string) {
	if err := j.validateSetUserAssignedIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userAssignedIdentityId",
		val,
	)
}

// Generates CDKTF code for importing a StorageAccountCustomerManagedKeyA resource upon running "cdktf plan <stack-name>".
func StorageAccountCustomerManagedKeyA_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStorageAccountCustomerManagedKeyA_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageAccountCustomerManagedKey.StorageAccountCustomerManagedKeyA",
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
func StorageAccountCustomerManagedKeyA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageAccountCustomerManagedKeyA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageAccountCustomerManagedKey.StorageAccountCustomerManagedKeyA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageAccountCustomerManagedKeyA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageAccountCustomerManagedKeyA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageAccountCustomerManagedKey.StorageAccountCustomerManagedKeyA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageAccountCustomerManagedKeyA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageAccountCustomerManagedKeyA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageAccountCustomerManagedKey.StorageAccountCustomerManagedKeyA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StorageAccountCustomerManagedKeyA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.storageAccountCustomerManagedKey.StorageAccountCustomerManagedKeyA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) PutTimeouts(value *StorageAccountCustomerManagedKeyTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ResetFederatedIdentityClientId() {
	_jsii_.InvokeVoid(
		s,
		"resetFederatedIdentityClientId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ResetKeyVaultId() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyVaultId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ResetKeyVaultUri() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyVaultUri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ResetKeyVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ResetManagedHsmKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetManagedHsmKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ResetUserAssignedIdentityId() {
	_jsii_.InvokeVoid(
		s,
		"resetUserAssignedIdentityId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountCustomerManagedKeyA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

