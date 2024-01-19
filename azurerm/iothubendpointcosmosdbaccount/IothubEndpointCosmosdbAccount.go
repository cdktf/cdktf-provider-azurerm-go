// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubendpointcosmosdbaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/iothubendpointcosmosdbaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/iothub_endpoint_cosmosdb_account azurerm_iothub_endpoint_cosmosdb_account}.
type IothubEndpointCosmosdbAccount interface {
	cdktf.TerraformResource
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerName() *string
	SetContainerName(val *string)
	ContainerNameInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndpointUri() *string
	SetEndpointUri(val *string)
	EndpointUriInput() *string
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
	IdentityId() *string
	SetIdentityId(val *string)
	IdentityIdInput() *string
	IdInput() *string
	IothubId() *string
	SetIothubId(val *string)
	IothubIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PartitionKeyName() *string
	SetPartitionKeyName(val *string)
	PartitionKeyNameInput() *string
	PartitionKeyTemplate() *string
	SetPartitionKeyTemplate(val *string)
	PartitionKeyTemplateInput() *string
	PrimaryKey() *string
	SetPrimaryKey(val *string)
	PrimaryKeyInput() *string
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
	SecondaryKey() *string
	SetSecondaryKey(val *string)
	SecondaryKeyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() IothubEndpointCosmosdbAccountTimeoutsOutputReference
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
	PutTimeouts(value *IothubEndpointCosmosdbAccountTimeouts)
	ResetAuthenticationType()
	ResetId()
	ResetIdentityId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartitionKeyName()
	ResetPartitionKeyTemplate()
	ResetPrimaryKey()
	ResetSecondaryKey()
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

// The jsii proxy struct for IothubEndpointCosmosdbAccount
type jsiiProxy_IothubEndpointCosmosdbAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) ContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) EndpointUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) EndpointUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) IdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) IdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) IothubId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) IothubIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) PartitionKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) PartitionKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) PartitionKeyTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKeyTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) PartitionKeyTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKeyTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) PrimaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) PrimaryKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) SecondaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) SecondaryKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) Timeouts() IothubEndpointCosmosdbAccountTimeoutsOutputReference {
	var returns IothubEndpointCosmosdbAccountTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/iothub_endpoint_cosmosdb_account azurerm_iothub_endpoint_cosmosdb_account} Resource.
func NewIothubEndpointCosmosdbAccount(scope constructs.Construct, id *string, config *IothubEndpointCosmosdbAccountConfig) IothubEndpointCosmosdbAccount {
	_init_.Initialize()

	if err := validateNewIothubEndpointCosmosdbAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IothubEndpointCosmosdbAccount{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothubEndpointCosmosdbAccount.IothubEndpointCosmosdbAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/iothub_endpoint_cosmosdb_account azurerm_iothub_endpoint_cosmosdb_account} Resource.
func NewIothubEndpointCosmosdbAccount_Override(i IothubEndpointCosmosdbAccount, scope constructs.Construct, id *string, config *IothubEndpointCosmosdbAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothubEndpointCosmosdbAccount.IothubEndpointCosmosdbAccount",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetContainerName(val *string) {
	if err := j.validateSetContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetEndpointUri(val *string) {
	if err := j.validateSetEndpointUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointUri",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetIdentityId(val *string) {
	if err := j.validateSetIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityId",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetIothubId(val *string) {
	if err := j.validateSetIothubIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iothubId",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetPartitionKeyName(val *string) {
	if err := j.validateSetPartitionKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionKeyName",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetPartitionKeyTemplate(val *string) {
	if err := j.validateSetPartitionKeyTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionKeyTemplate",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetPrimaryKey(val *string) {
	if err := j.validateSetPrimaryKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKey",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_IothubEndpointCosmosdbAccount)SetSecondaryKey(val *string) {
	if err := j.validateSetSecondaryKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryKey",
		val,
	)
}

// Generates CDKTF code for importing a IothubEndpointCosmosdbAccount resource upon running "cdktf plan <stack-name>".
func IothubEndpointCosmosdbAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIothubEndpointCosmosdbAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothubEndpointCosmosdbAccount.IothubEndpointCosmosdbAccount",
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
func IothubEndpointCosmosdbAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothubEndpointCosmosdbAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothubEndpointCosmosdbAccount.IothubEndpointCosmosdbAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IothubEndpointCosmosdbAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothubEndpointCosmosdbAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothubEndpointCosmosdbAccount.IothubEndpointCosmosdbAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IothubEndpointCosmosdbAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothubEndpointCosmosdbAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothubEndpointCosmosdbAccount.IothubEndpointCosmosdbAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IothubEndpointCosmosdbAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.iothubEndpointCosmosdbAccount.IothubEndpointCosmosdbAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) PutTimeouts(value *IothubEndpointCosmosdbAccountTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ResetIdentityId() {
	_jsii_.InvokeVoid(
		i,
		"resetIdentityId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ResetPartitionKeyName() {
	_jsii_.InvokeVoid(
		i,
		"resetPartitionKeyName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ResetPartitionKeyTemplate() {
	_jsii_.InvokeVoid(
		i,
		"resetPartitionKeyTemplate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ResetPrimaryKey() {
	_jsii_.InvokeVoid(
		i,
		"resetPrimaryKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ResetSecondaryKey() {
	_jsii_.InvokeVoid(
		i,
		"resetSecondaryKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubEndpointCosmosdbAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

