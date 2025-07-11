// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbpostgresqlcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/cosmosdbpostgresqlcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/cosmosdb_postgresql_cluster azurerm_cosmosdb_postgresql_cluster}.
type CosmosdbPostgresqlCluster interface {
	cdktf.TerraformResource
	AdministratorLoginPassword() *string
	SetAdministratorLoginPassword(val *string)
	AdministratorLoginPasswordInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CitusVersion() *string
	SetCitusVersion(val *string)
	CitusVersionInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoordinatorPublicIpAccessEnabled() interface{}
	SetCoordinatorPublicIpAccessEnabled(val interface{})
	CoordinatorPublicIpAccessEnabledInput() interface{}
	CoordinatorServerEdition() *string
	SetCoordinatorServerEdition(val *string)
	CoordinatorServerEditionInput() *string
	CoordinatorStorageQuotaInMb() *float64
	SetCoordinatorStorageQuotaInMb(val *float64)
	CoordinatorStorageQuotaInMbInput() *float64
	CoordinatorVcoreCount() *float64
	SetCoordinatorVcoreCount(val *float64)
	CoordinatorVcoreCountInput() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EarliestRestoreTime() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HaEnabled() interface{}
	SetHaEnabled(val interface{})
	HaEnabledInput() interface{}
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
	MaintenanceWindow() CosmosdbPostgresqlClusterMaintenanceWindowOutputReference
	MaintenanceWindowInput() *CosmosdbPostgresqlClusterMaintenanceWindow
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	NodePublicIpAccessEnabled() interface{}
	SetNodePublicIpAccessEnabled(val interface{})
	NodePublicIpAccessEnabledInput() interface{}
	NodeServerEdition() *string
	SetNodeServerEdition(val *string)
	NodeServerEditionInput() *string
	NodeStorageQuotaInMb() *float64
	SetNodeStorageQuotaInMb(val *float64)
	NodeStorageQuotaInMbInput() *float64
	NodeVcores() *float64
	SetNodeVcores(val *float64)
	NodeVcoresInput() *float64
	PointInTimeInUtc() *string
	SetPointInTimeInUtc(val *string)
	PointInTimeInUtcInput() *string
	PreferredPrimaryZone() *string
	SetPreferredPrimaryZone(val *string)
	PreferredPrimaryZoneInput() *string
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
	Servers() CosmosdbPostgresqlClusterServersList
	ShardsOnCoordinatorEnabled() interface{}
	SetShardsOnCoordinatorEnabled(val interface{})
	ShardsOnCoordinatorEnabledInput() interface{}
	SourceLocation() *string
	SetSourceLocation(val *string)
	SourceLocationInput() *string
	SourceResourceId() *string
	SetSourceResourceId(val *string)
	SourceResourceIdInput() *string
	SqlVersion() *string
	SetSqlVersion(val *string)
	SqlVersionInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CosmosdbPostgresqlClusterTimeoutsOutputReference
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
	PutMaintenanceWindow(value *CosmosdbPostgresqlClusterMaintenanceWindow)
	PutTimeouts(value *CosmosdbPostgresqlClusterTimeouts)
	ResetAdministratorLoginPassword()
	ResetCitusVersion()
	ResetCoordinatorPublicIpAccessEnabled()
	ResetCoordinatorServerEdition()
	ResetCoordinatorStorageQuotaInMb()
	ResetCoordinatorVcoreCount()
	ResetHaEnabled()
	ResetId()
	ResetMaintenanceWindow()
	ResetNodePublicIpAccessEnabled()
	ResetNodeServerEdition()
	ResetNodeStorageQuotaInMb()
	ResetNodeVcores()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeInUtc()
	ResetPreferredPrimaryZone()
	ResetShardsOnCoordinatorEnabled()
	ResetSourceLocation()
	ResetSourceResourceId()
	ResetSqlVersion()
	ResetTags()
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

// The jsii proxy struct for CosmosdbPostgresqlCluster
type jsiiProxy_CosmosdbPostgresqlCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) AdministratorLoginPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) AdministratorLoginPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) CitusVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"citusVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) CitusVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"citusVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) CoordinatorPublicIpAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coordinatorPublicIpAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) CoordinatorPublicIpAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coordinatorPublicIpAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) CoordinatorServerEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coordinatorServerEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) CoordinatorServerEditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coordinatorServerEditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) CoordinatorStorageQuotaInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coordinatorStorageQuotaInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) CoordinatorStorageQuotaInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coordinatorStorageQuotaInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) CoordinatorVcoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coordinatorVcoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) CoordinatorVcoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coordinatorVcoreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) EarliestRestoreTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"earliestRestoreTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) HaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) HaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"haEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) MaintenanceWindow() CosmosdbPostgresqlClusterMaintenanceWindowOutputReference {
	var returns CosmosdbPostgresqlClusterMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) MaintenanceWindowInput() *CosmosdbPostgresqlClusterMaintenanceWindow {
	var returns *CosmosdbPostgresqlClusterMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) NodePublicIpAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodePublicIpAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) NodePublicIpAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodePublicIpAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) NodeServerEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeServerEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) NodeServerEditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeServerEditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) NodeStorageQuotaInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeStorageQuotaInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) NodeStorageQuotaInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeStorageQuotaInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) NodeVcores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeVcores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) NodeVcoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeVcoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) PointInTimeInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) PointInTimeInUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeInUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) PreferredPrimaryZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredPrimaryZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) PreferredPrimaryZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredPrimaryZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Servers() CosmosdbPostgresqlClusterServersList {
	var returns CosmosdbPostgresqlClusterServersList
	_jsii_.Get(
		j,
		"servers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) ShardsOnCoordinatorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shardsOnCoordinatorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) ShardsOnCoordinatorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shardsOnCoordinatorEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) SourceLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) SourceLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) SourceResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) SourceResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) SqlVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) SqlVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) Timeouts() CosmosdbPostgresqlClusterTimeoutsOutputReference {
	var returns CosmosdbPostgresqlClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/cosmosdb_postgresql_cluster azurerm_cosmosdb_postgresql_cluster} Resource.
func NewCosmosdbPostgresqlCluster(scope constructs.Construct, id *string, config *CosmosdbPostgresqlClusterConfig) CosmosdbPostgresqlCluster {
	_init_.Initialize()

	if err := validateNewCosmosdbPostgresqlClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CosmosdbPostgresqlCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbPostgresqlCluster.CosmosdbPostgresqlCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/cosmosdb_postgresql_cluster azurerm_cosmosdb_postgresql_cluster} Resource.
func NewCosmosdbPostgresqlCluster_Override(c CosmosdbPostgresqlCluster, scope constructs.Construct, id *string, config *CosmosdbPostgresqlClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbPostgresqlCluster.CosmosdbPostgresqlCluster",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetAdministratorLoginPassword(val *string) {
	if err := j.validateSetAdministratorLoginPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorLoginPassword",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetCitusVersion(val *string) {
	if err := j.validateSetCitusVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"citusVersion",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetCoordinatorPublicIpAccessEnabled(val interface{}) {
	if err := j.validateSetCoordinatorPublicIpAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coordinatorPublicIpAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetCoordinatorServerEdition(val *string) {
	if err := j.validateSetCoordinatorServerEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coordinatorServerEdition",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetCoordinatorStorageQuotaInMb(val *float64) {
	if err := j.validateSetCoordinatorStorageQuotaInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coordinatorStorageQuotaInMb",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetCoordinatorVcoreCount(val *float64) {
	if err := j.validateSetCoordinatorVcoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coordinatorVcoreCount",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetHaEnabled(val interface{}) {
	if err := j.validateSetHaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetNodePublicIpAccessEnabled(val interface{}) {
	if err := j.validateSetNodePublicIpAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodePublicIpAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetNodeServerEdition(val *string) {
	if err := j.validateSetNodeServerEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeServerEdition",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetNodeStorageQuotaInMb(val *float64) {
	if err := j.validateSetNodeStorageQuotaInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeStorageQuotaInMb",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetNodeVcores(val *float64) {
	if err := j.validateSetNodeVcoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeVcores",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetPointInTimeInUtc(val *string) {
	if err := j.validateSetPointInTimeInUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTimeInUtc",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetPreferredPrimaryZone(val *string) {
	if err := j.validateSetPreferredPrimaryZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredPrimaryZone",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetShardsOnCoordinatorEnabled(val interface{}) {
	if err := j.validateSetShardsOnCoordinatorEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shardsOnCoordinatorEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetSourceLocation(val *string) {
	if err := j.validateSetSourceLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceLocation",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetSourceResourceId(val *string) {
	if err := j.validateSetSourceResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceResourceId",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetSqlVersion(val *string) {
	if err := j.validateSetSqlVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlVersion",
		val,
	)
}

func (j *jsiiProxy_CosmosdbPostgresqlCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a CosmosdbPostgresqlCluster resource upon running "cdktf plan <stack-name>".
func CosmosdbPostgresqlCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCosmosdbPostgresqlCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbPostgresqlCluster.CosmosdbPostgresqlCluster",
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
func CosmosdbPostgresqlCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbPostgresqlCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbPostgresqlCluster.CosmosdbPostgresqlCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CosmosdbPostgresqlCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbPostgresqlCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbPostgresqlCluster.CosmosdbPostgresqlCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CosmosdbPostgresqlCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbPostgresqlCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbPostgresqlCluster.CosmosdbPostgresqlCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CosmosdbPostgresqlCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cosmosdbPostgresqlCluster.CosmosdbPostgresqlCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CosmosdbPostgresqlCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CosmosdbPostgresqlCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CosmosdbPostgresqlCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CosmosdbPostgresqlCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CosmosdbPostgresqlCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CosmosdbPostgresqlCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CosmosdbPostgresqlCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CosmosdbPostgresqlCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CosmosdbPostgresqlCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CosmosdbPostgresqlCluster) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) PutMaintenanceWindow(value *CosmosdbPostgresqlClusterMaintenanceWindow) {
	if err := c.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) PutTimeouts(value *CosmosdbPostgresqlClusterTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetAdministratorLoginPassword() {
	_jsii_.InvokeVoid(
		c,
		"resetAdministratorLoginPassword",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetCitusVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetCitusVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetCoordinatorPublicIpAccessEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetCoordinatorPublicIpAccessEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetCoordinatorServerEdition() {
	_jsii_.InvokeVoid(
		c,
		"resetCoordinatorServerEdition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetCoordinatorStorageQuotaInMb() {
	_jsii_.InvokeVoid(
		c,
		"resetCoordinatorStorageQuotaInMb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetCoordinatorVcoreCount() {
	_jsii_.InvokeVoid(
		c,
		"resetCoordinatorVcoreCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetHaEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetHaEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		c,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetNodePublicIpAccessEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetNodePublicIpAccessEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetNodeServerEdition() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeServerEdition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetNodeStorageQuotaInMb() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeStorageQuotaInMb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetNodeVcores() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeVcores",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetPointInTimeInUtc() {
	_jsii_.InvokeVoid(
		c,
		"resetPointInTimeInUtc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetPreferredPrimaryZone() {
	_jsii_.InvokeVoid(
		c,
		"resetPreferredPrimaryZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetShardsOnCoordinatorEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetShardsOnCoordinatorEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetSourceLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetSourceResourceId() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceResourceId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetSqlVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetSqlVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbPostgresqlCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

