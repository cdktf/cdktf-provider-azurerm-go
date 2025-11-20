// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagesyncserverendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/storagesyncserverendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/storage_sync_server_endpoint azurerm_storage_sync_server_endpoint}.
type StorageSyncServerEndpoint interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudTieringEnabled() interface{}
	SetCloudTieringEnabled(val interface{})
	CloudTieringEnabledInput() interface{}
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
	InitialDownloadPolicy() *string
	SetInitialDownloadPolicy(val *string)
	InitialDownloadPolicyInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalCacheMode() *string
	SetLocalCacheMode(val *string)
	LocalCacheModeInput() *string
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
	RegisteredServerId() *string
	SetRegisteredServerId(val *string)
	RegisteredServerIdInput() *string
	ServerLocalPath() *string
	SetServerLocalPath(val *string)
	ServerLocalPathInput() *string
	StorageSyncGroupId() *string
	SetStorageSyncGroupId(val *string)
	StorageSyncGroupIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TierFilesOlderThanDays() *float64
	SetTierFilesOlderThanDays(val *float64)
	TierFilesOlderThanDaysInput() *float64
	Timeouts() StorageSyncServerEndpointTimeoutsOutputReference
	TimeoutsInput() interface{}
	VolumeFreeSpacePercent() *float64
	SetVolumeFreeSpacePercent(val *float64)
	VolumeFreeSpacePercentInput() *float64
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
	PutTimeouts(value *StorageSyncServerEndpointTimeouts)
	ResetCloudTieringEnabled()
	ResetId()
	ResetInitialDownloadPolicy()
	ResetLocalCacheMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTierFilesOlderThanDays()
	ResetTimeouts()
	ResetVolumeFreeSpacePercent()
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

// The jsii proxy struct for StorageSyncServerEndpoint
type jsiiProxy_StorageSyncServerEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StorageSyncServerEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) CloudTieringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudTieringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) CloudTieringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudTieringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) InitialDownloadPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialDownloadPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) InitialDownloadPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialDownloadPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) LocalCacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localCacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) LocalCacheModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localCacheModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) RegisteredServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) RegisteredServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredServerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) ServerLocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverLocalPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) ServerLocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverLocalPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) StorageSyncGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSyncGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) StorageSyncGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSyncGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) TierFilesOlderThanDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierFilesOlderThanDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) TierFilesOlderThanDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierFilesOlderThanDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) Timeouts() StorageSyncServerEndpointTimeoutsOutputReference {
	var returns StorageSyncServerEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) VolumeFreeSpacePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeFreeSpacePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageSyncServerEndpoint) VolumeFreeSpacePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeFreeSpacePercentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/storage_sync_server_endpoint azurerm_storage_sync_server_endpoint} Resource.
func NewStorageSyncServerEndpoint(scope constructs.Construct, id *string, config *StorageSyncServerEndpointConfig) StorageSyncServerEndpoint {
	_init_.Initialize()

	if err := validateNewStorageSyncServerEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageSyncServerEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageSyncServerEndpoint.StorageSyncServerEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/storage_sync_server_endpoint azurerm_storage_sync_server_endpoint} Resource.
func NewStorageSyncServerEndpoint_Override(s StorageSyncServerEndpoint, scope constructs.Construct, id *string, config *StorageSyncServerEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageSyncServerEndpoint.StorageSyncServerEndpoint",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetCloudTieringEnabled(val interface{}) {
	if err := j.validateSetCloudTieringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudTieringEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetInitialDownloadPolicy(val *string) {
	if err := j.validateSetInitialDownloadPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDownloadPolicy",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetLocalCacheMode(val *string) {
	if err := j.validateSetLocalCacheModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localCacheMode",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetRegisteredServerId(val *string) {
	if err := j.validateSetRegisteredServerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registeredServerId",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetServerLocalPath(val *string) {
	if err := j.validateSetServerLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverLocalPath",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetStorageSyncGroupId(val *string) {
	if err := j.validateSetStorageSyncGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSyncGroupId",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetTierFilesOlderThanDays(val *float64) {
	if err := j.validateSetTierFilesOlderThanDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tierFilesOlderThanDays",
		val,
	)
}

func (j *jsiiProxy_StorageSyncServerEndpoint)SetVolumeFreeSpacePercent(val *float64) {
	if err := j.validateSetVolumeFreeSpacePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeFreeSpacePercent",
		val,
	)
}

// Generates CDKTF code for importing a StorageSyncServerEndpoint resource upon running "cdktf plan <stack-name>".
func StorageSyncServerEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStorageSyncServerEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageSyncServerEndpoint.StorageSyncServerEndpoint",
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
func StorageSyncServerEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageSyncServerEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageSyncServerEndpoint.StorageSyncServerEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageSyncServerEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageSyncServerEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageSyncServerEndpoint.StorageSyncServerEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageSyncServerEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageSyncServerEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageSyncServerEndpoint.StorageSyncServerEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StorageSyncServerEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.storageSyncServerEndpoint.StorageSyncServerEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageSyncServerEndpoint) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageSyncServerEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageSyncServerEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageSyncServerEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageSyncServerEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageSyncServerEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageSyncServerEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageSyncServerEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageSyncServerEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageSyncServerEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageSyncServerEndpoint) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) PutTimeouts(value *StorageSyncServerEndpointTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ResetCloudTieringEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudTieringEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ResetInitialDownloadPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetInitialDownloadPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ResetLocalCacheMode() {
	_jsii_.InvokeVoid(
		s,
		"resetLocalCacheMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ResetTierFilesOlderThanDays() {
	_jsii_.InvokeVoid(
		s,
		"resetTierFilesOlderThanDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ResetVolumeFreeSpacePercent() {
	_jsii_.InvokeVoid(
		s,
		"resetVolumeFreeSpacePercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageSyncServerEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageSyncServerEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageSyncServerEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

