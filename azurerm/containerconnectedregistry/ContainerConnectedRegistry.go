// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerconnectedregistry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/containerconnectedregistry/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/container_connected_registry azurerm_container_connected_registry}.
type ContainerConnectedRegistry interface {
	cdktf.TerraformResource
	AuditLogEnabled() interface{}
	SetAuditLogEnabled(val interface{})
	AuditLogEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientTokenIds() *[]*string
	SetClientTokenIds(val *[]*string)
	ClientTokenIdsInput() *[]*string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerRegistryId() *string
	SetContainerRegistryId(val *string)
	ContainerRegistryIdInput() *string
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
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Notification() ContainerConnectedRegistryNotificationList
	NotificationInput() interface{}
	ParentRegistryId() *string
	SetParentRegistryId(val *string)
	ParentRegistryIdInput() *string
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
	SyncMessageTtl() *string
	SetSyncMessageTtl(val *string)
	SyncMessageTtlInput() *string
	SyncSchedule() *string
	SetSyncSchedule(val *string)
	SyncScheduleInput() *string
	SyncTokenId() *string
	SetSyncTokenId(val *string)
	SyncTokenIdInput() *string
	SyncWindow() *string
	SetSyncWindow(val *string)
	SyncWindowInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ContainerConnectedRegistryTimeoutsOutputReference
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
	PutNotification(value interface{})
	PutTimeouts(value *ContainerConnectedRegistryTimeouts)
	ResetAuditLogEnabled()
	ResetClientTokenIds()
	ResetId()
	ResetLogLevel()
	ResetMode()
	ResetNotification()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentRegistryId()
	ResetSyncMessageTtl()
	ResetSyncSchedule()
	ResetSyncWindow()
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

// The jsii proxy struct for ContainerConnectedRegistry
type jsiiProxy_ContainerConnectedRegistry struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ContainerConnectedRegistry) AuditLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) AuditLogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) ClientTokenIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientTokenIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) ClientTokenIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientTokenIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) ContainerRegistryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) ContainerRegistryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Notification() ContainerConnectedRegistryNotificationList {
	var returns ContainerConnectedRegistryNotificationList
	_jsii_.Get(
		j,
		"notification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) NotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) ParentRegistryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentRegistryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) ParentRegistryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentRegistryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) SyncMessageTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncMessageTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) SyncMessageTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncMessageTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) SyncSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) SyncScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) SyncTokenId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncTokenId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) SyncTokenIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncTokenIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) SyncWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) SyncWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) Timeouts() ContainerConnectedRegistryTimeoutsOutputReference {
	var returns ContainerConnectedRegistryTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerConnectedRegistry) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/container_connected_registry azurerm_container_connected_registry} Resource.
func NewContainerConnectedRegistry(scope constructs.Construct, id *string, config *ContainerConnectedRegistryConfig) ContainerConnectedRegistry {
	_init_.Initialize()

	if err := validateNewContainerConnectedRegistryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerConnectedRegistry{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerConnectedRegistry.ContainerConnectedRegistry",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/container_connected_registry azurerm_container_connected_registry} Resource.
func NewContainerConnectedRegistry_Override(c ContainerConnectedRegistry, scope constructs.Construct, id *string, config *ContainerConnectedRegistryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerConnectedRegistry.ContainerConnectedRegistry",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetAuditLogEnabled(val interface{}) {
	if err := j.validateSetAuditLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetClientTokenIds(val *[]*string) {
	if err := j.validateSetClientTokenIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTokenIds",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetContainerRegistryId(val *string) {
	if err := j.validateSetContainerRegistryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryId",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetParentRegistryId(val *string) {
	if err := j.validateSetParentRegistryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentRegistryId",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetSyncMessageTtl(val *string) {
	if err := j.validateSetSyncMessageTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncMessageTtl",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetSyncSchedule(val *string) {
	if err := j.validateSetSyncScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncSchedule",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetSyncTokenId(val *string) {
	if err := j.validateSetSyncTokenIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncTokenId",
		val,
	)
}

func (j *jsiiProxy_ContainerConnectedRegistry)SetSyncWindow(val *string) {
	if err := j.validateSetSyncWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncWindow",
		val,
	)
}

// Generates CDKTF code for importing a ContainerConnectedRegistry resource upon running "cdktf plan <stack-name>".
func ContainerConnectedRegistry_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateContainerConnectedRegistry_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerConnectedRegistry.ContainerConnectedRegistry",
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
func ContainerConnectedRegistry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerConnectedRegistry_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerConnectedRegistry.ContainerConnectedRegistry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerConnectedRegistry_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerConnectedRegistry_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerConnectedRegistry.ContainerConnectedRegistry",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerConnectedRegistry_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerConnectedRegistry_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerConnectedRegistry.ContainerConnectedRegistry",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContainerConnectedRegistry_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.containerConnectedRegistry.ContainerConnectedRegistry",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) PutNotification(value interface{}) {
	if err := c.validatePutNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNotification",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) PutTimeouts(value *ContainerConnectedRegistryTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetAuditLogEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAuditLogEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetClientTokenIds() {
	_jsii_.InvokeVoid(
		c,
		"resetClientTokenIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetLogLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetMode() {
	_jsii_.InvokeVoid(
		c,
		"resetMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetNotification() {
	_jsii_.InvokeVoid(
		c,
		"resetNotification",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetParentRegistryId() {
	_jsii_.InvokeVoid(
		c,
		"resetParentRegistryId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetSyncMessageTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetSyncMessageTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetSyncSchedule() {
	_jsii_.InvokeVoid(
		c,
		"resetSyncSchedule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetSyncWindow() {
	_jsii_.InvokeVoid(
		c,
		"resetSyncWindow",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerConnectedRegistry) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerConnectedRegistry) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

