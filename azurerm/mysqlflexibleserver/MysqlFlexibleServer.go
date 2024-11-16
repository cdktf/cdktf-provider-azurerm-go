// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mysqlflexibleserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/mysqlflexibleserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/mysql_flexible_server azurerm_mysql_flexible_server}.
type MysqlFlexibleServer interface {
	cdktf.TerraformResource
	AdministratorLogin() *string
	SetAdministratorLogin(val *string)
	AdministratorLoginInput() *string
	AdministratorPassword() *string
	SetAdministratorPassword(val *string)
	AdministratorPasswordInput() *string
	BackupRetentionDays() *float64
	SetBackupRetentionDays(val *float64)
	BackupRetentionDaysInput() *float64
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
	CreateMode() *string
	SetCreateMode(val *string)
	CreateModeInput() *string
	CustomerManagedKey() MysqlFlexibleServerCustomerManagedKeyOutputReference
	CustomerManagedKeyInput() *MysqlFlexibleServerCustomerManagedKey
	DelegatedSubnetId() *string
	SetDelegatedSubnetId(val *string)
	DelegatedSubnetIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeoRedundantBackupEnabled() interface{}
	SetGeoRedundantBackupEnabled(val interface{})
	GeoRedundantBackupEnabledInput() interface{}
	HighAvailability() MysqlFlexibleServerHighAvailabilityOutputReference
	HighAvailabilityInput() *MysqlFlexibleServerHighAvailability
	Id() *string
	SetId(val *string)
	Identity() MysqlFlexibleServerIdentityOutputReference
	IdentityInput() *MysqlFlexibleServerIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceWindow() MysqlFlexibleServerMaintenanceWindowOutputReference
	MaintenanceWindowInput() *MysqlFlexibleServerMaintenanceWindow
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PointInTimeRestoreTimeInUtc() *string
	SetPointInTimeRestoreTimeInUtc(val *string)
	PointInTimeRestoreTimeInUtcInput() *string
	PrivateDnsZoneId() *string
	SetPrivateDnsZoneId(val *string)
	PrivateDnsZoneIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicNetworkAccessEnabled() cdktf.IResolvable
	// Experimental.
	RawOverrides() interface{}
	ReplicaCapacity() *float64
	ReplicationRole() *string
	SetReplicationRole(val *string)
	ReplicationRoleInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	SourceServerId() *string
	SetSourceServerId(val *string)
	SourceServerIdInput() *string
	Storage() MysqlFlexibleServerStorageOutputReference
	StorageInput() *MysqlFlexibleServerStorage
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MysqlFlexibleServerTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutCustomerManagedKey(value *MysqlFlexibleServerCustomerManagedKey)
	PutHighAvailability(value *MysqlFlexibleServerHighAvailability)
	PutIdentity(value *MysqlFlexibleServerIdentity)
	PutMaintenanceWindow(value *MysqlFlexibleServerMaintenanceWindow)
	PutStorage(value *MysqlFlexibleServerStorage)
	PutTimeouts(value *MysqlFlexibleServerTimeouts)
	ResetAdministratorLogin()
	ResetAdministratorPassword()
	ResetBackupRetentionDays()
	ResetCreateMode()
	ResetCustomerManagedKey()
	ResetDelegatedSubnetId()
	ResetGeoRedundantBackupEnabled()
	ResetHighAvailability()
	ResetId()
	ResetIdentity()
	ResetMaintenanceWindow()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeRestoreTimeInUtc()
	ResetPrivateDnsZoneId()
	ResetReplicationRole()
	ResetSkuName()
	ResetSourceServerId()
	ResetStorage()
	ResetTags()
	ResetTimeouts()
	ResetVersion()
	ResetZone()
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

// The jsii proxy struct for MysqlFlexibleServer
type jsiiProxy_MysqlFlexibleServer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MysqlFlexibleServer) AdministratorLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) AdministratorLoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) AdministratorPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) AdministratorPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) BackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) BackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) CreateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) CreateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) CustomerManagedKey() MysqlFlexibleServerCustomerManagedKeyOutputReference {
	var returns MysqlFlexibleServerCustomerManagedKeyOutputReference
	_jsii_.Get(
		j,
		"customerManagedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) CustomerManagedKeyInput() *MysqlFlexibleServerCustomerManagedKey {
	var returns *MysqlFlexibleServerCustomerManagedKey
	_jsii_.Get(
		j,
		"customerManagedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) DelegatedSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) DelegatedSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) GeoRedundantBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoRedundantBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) GeoRedundantBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoRedundantBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) HighAvailability() MysqlFlexibleServerHighAvailabilityOutputReference {
	var returns MysqlFlexibleServerHighAvailabilityOutputReference
	_jsii_.Get(
		j,
		"highAvailability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) HighAvailabilityInput() *MysqlFlexibleServerHighAvailability {
	var returns *MysqlFlexibleServerHighAvailability
	_jsii_.Get(
		j,
		"highAvailabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Identity() MysqlFlexibleServerIdentityOutputReference {
	var returns MysqlFlexibleServerIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) IdentityInput() *MysqlFlexibleServerIdentity {
	var returns *MysqlFlexibleServerIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) MaintenanceWindow() MysqlFlexibleServerMaintenanceWindowOutputReference {
	var returns MysqlFlexibleServerMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) MaintenanceWindowInput() *MysqlFlexibleServerMaintenanceWindow {
	var returns *MysqlFlexibleServerMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) PointInTimeRestoreTimeInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeRestoreTimeInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) PointInTimeRestoreTimeInUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeRestoreTimeInUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) PrivateDnsZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) PrivateDnsZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) PublicNetworkAccessEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) ReplicaCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) ReplicationRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) ReplicationRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) SourceServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) SourceServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Storage() MysqlFlexibleServerStorageOutputReference {
	var returns MysqlFlexibleServerStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) StorageInput() *MysqlFlexibleServerStorage {
	var returns *MysqlFlexibleServerStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Timeouts() MysqlFlexibleServerTimeoutsOutputReference {
	var returns MysqlFlexibleServerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlFlexibleServer) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/mysql_flexible_server azurerm_mysql_flexible_server} Resource.
func NewMysqlFlexibleServer(scope constructs.Construct, id *string, config *MysqlFlexibleServerConfig) MysqlFlexibleServer {
	_init_.Initialize()

	if err := validateNewMysqlFlexibleServerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MysqlFlexibleServer{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mysqlFlexibleServer.MysqlFlexibleServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/mysql_flexible_server azurerm_mysql_flexible_server} Resource.
func NewMysqlFlexibleServer_Override(m MysqlFlexibleServer, scope constructs.Construct, id *string, config *MysqlFlexibleServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mysqlFlexibleServer.MysqlFlexibleServer",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetAdministratorLogin(val *string) {
	if err := j.validateSetAdministratorLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorLogin",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetAdministratorPassword(val *string) {
	if err := j.validateSetAdministratorPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorPassword",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetBackupRetentionDays(val *float64) {
	if err := j.validateSetBackupRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetCreateMode(val *string) {
	if err := j.validateSetCreateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createMode",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetDelegatedSubnetId(val *string) {
	if err := j.validateSetDelegatedSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delegatedSubnetId",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetGeoRedundantBackupEnabled(val interface{}) {
	if err := j.validateSetGeoRedundantBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoRedundantBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetPointInTimeRestoreTimeInUtc(val *string) {
	if err := j.validateSetPointInTimeRestoreTimeInUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTimeRestoreTimeInUtc",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetPrivateDnsZoneId(val *string) {
	if err := j.validateSetPrivateDnsZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsZoneId",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetReplicationRole(val *string) {
	if err := j.validateSetReplicationRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationRole",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetSourceServerId(val *string) {
	if err := j.validateSetSourceServerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceServerId",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_MysqlFlexibleServer)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a MysqlFlexibleServer resource upon running "cdktf plan <stack-name>".
func MysqlFlexibleServer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMysqlFlexibleServer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mysqlFlexibleServer.MysqlFlexibleServer",
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
func MysqlFlexibleServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMysqlFlexibleServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mysqlFlexibleServer.MysqlFlexibleServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MysqlFlexibleServer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMysqlFlexibleServer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mysqlFlexibleServer.MysqlFlexibleServer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MysqlFlexibleServer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMysqlFlexibleServer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mysqlFlexibleServer.MysqlFlexibleServer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MysqlFlexibleServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mysqlFlexibleServer.MysqlFlexibleServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) PutCustomerManagedKey(value *MysqlFlexibleServerCustomerManagedKey) {
	if err := m.validatePutCustomerManagedKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCustomerManagedKey",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) PutHighAvailability(value *MysqlFlexibleServerHighAvailability) {
	if err := m.validatePutHighAvailabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putHighAvailability",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) PutIdentity(value *MysqlFlexibleServerIdentity) {
	if err := m.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIdentity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) PutMaintenanceWindow(value *MysqlFlexibleServerMaintenanceWindow) {
	if err := m.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) PutStorage(value *MysqlFlexibleServerStorage) {
	if err := m.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStorage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) PutTimeouts(value *MysqlFlexibleServerTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetAdministratorLogin() {
	_jsii_.InvokeVoid(
		m,
		"resetAdministratorLogin",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetAdministratorPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetAdministratorPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetBackupRetentionDays() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupRetentionDays",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetCreateMode() {
	_jsii_.InvokeVoid(
		m,
		"resetCreateMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetCustomerManagedKey() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomerManagedKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetDelegatedSubnetId() {
	_jsii_.InvokeVoid(
		m,
		"resetDelegatedSubnetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetGeoRedundantBackupEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetGeoRedundantBackupEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetHighAvailability() {
	_jsii_.InvokeVoid(
		m,
		"resetHighAvailability",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetIdentity() {
	_jsii_.InvokeVoid(
		m,
		"resetIdentity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetPointInTimeRestoreTimeInUtc() {
	_jsii_.InvokeVoid(
		m,
		"resetPointInTimeRestoreTimeInUtc",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetPrivateDnsZoneId() {
	_jsii_.InvokeVoid(
		m,
		"resetPrivateDnsZoneId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetReplicationRole() {
	_jsii_.InvokeVoid(
		m,
		"resetReplicationRole",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetSkuName() {
	_jsii_.InvokeVoid(
		m,
		"resetSkuName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetSourceServerId() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceServerId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetStorage() {
	_jsii_.InvokeVoid(
		m,
		"resetStorage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) ResetZone() {
	_jsii_.InvokeVoid(
		m,
		"resetZone",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlFlexibleServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlFlexibleServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

