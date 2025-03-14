// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlflexibleserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/postgresqlflexibleserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/postgresql_flexible_server azurerm_postgresql_flexible_server}.
type PostgresqlFlexibleServer interface {
	cdktf.TerraformResource
	AdministratorLogin() *string
	SetAdministratorLogin(val *string)
	AdministratorLoginInput() *string
	AdministratorPassword() *string
	SetAdministratorPassword(val *string)
	AdministratorPasswordInput() *string
	AdministratorPasswordWo() *string
	SetAdministratorPasswordWo(val *string)
	AdministratorPasswordWoInput() *string
	AdministratorPasswordWoVersion() *float64
	SetAdministratorPasswordWoVersion(val *float64)
	AdministratorPasswordWoVersionInput() *float64
	Authentication() PostgresqlFlexibleServerAuthenticationOutputReference
	AuthenticationInput() *PostgresqlFlexibleServerAuthentication
	AutoGrowEnabled() interface{}
	SetAutoGrowEnabled(val interface{})
	AutoGrowEnabledInput() interface{}
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
	CustomerManagedKey() PostgresqlFlexibleServerCustomerManagedKeyOutputReference
	CustomerManagedKeyInput() *PostgresqlFlexibleServerCustomerManagedKey
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
	HighAvailability() PostgresqlFlexibleServerHighAvailabilityOutputReference
	HighAvailabilityInput() *PostgresqlFlexibleServerHighAvailability
	Id() *string
	SetId(val *string)
	Identity() PostgresqlFlexibleServerIdentityOutputReference
	IdentityInput() *PostgresqlFlexibleServerIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceWindow() PostgresqlFlexibleServerMaintenanceWindowOutputReference
	MaintenanceWindowInput() *PostgresqlFlexibleServerMaintenanceWindow
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
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
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
	StorageMb() *float64
	SetStorageMb(val *float64)
	StorageMbInput() *float64
	StorageTier() *string
	SetStorageTier(val *string)
	StorageTierInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() PostgresqlFlexibleServerTimeoutsOutputReference
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
	PutAuthentication(value *PostgresqlFlexibleServerAuthentication)
	PutCustomerManagedKey(value *PostgresqlFlexibleServerCustomerManagedKey)
	PutHighAvailability(value *PostgresqlFlexibleServerHighAvailability)
	PutIdentity(value *PostgresqlFlexibleServerIdentity)
	PutMaintenanceWindow(value *PostgresqlFlexibleServerMaintenanceWindow)
	PutTimeouts(value *PostgresqlFlexibleServerTimeouts)
	ResetAdministratorLogin()
	ResetAdministratorPassword()
	ResetAdministratorPasswordWo()
	ResetAdministratorPasswordWoVersion()
	ResetAuthentication()
	ResetAutoGrowEnabled()
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
	ResetPublicNetworkAccessEnabled()
	ResetReplicationRole()
	ResetSkuName()
	ResetSourceServerId()
	ResetStorageMb()
	ResetStorageTier()
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

// The jsii proxy struct for PostgresqlFlexibleServer
type jsiiProxy_PostgresqlFlexibleServer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorLoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorPasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPasswordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorPasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPasswordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorPasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"administratorPasswordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AdministratorPasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"administratorPasswordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Authentication() PostgresqlFlexibleServerAuthenticationOutputReference {
	var returns PostgresqlFlexibleServerAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AuthenticationInput() *PostgresqlFlexibleServerAuthentication {
	var returns *PostgresqlFlexibleServerAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AutoGrowEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoGrowEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) AutoGrowEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoGrowEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) BackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) BackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) CreateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) CreateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) CustomerManagedKey() PostgresqlFlexibleServerCustomerManagedKeyOutputReference {
	var returns PostgresqlFlexibleServerCustomerManagedKeyOutputReference
	_jsii_.Get(
		j,
		"customerManagedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) CustomerManagedKeyInput() *PostgresqlFlexibleServerCustomerManagedKey {
	var returns *PostgresqlFlexibleServerCustomerManagedKey
	_jsii_.Get(
		j,
		"customerManagedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) DelegatedSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) DelegatedSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) GeoRedundantBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoRedundantBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) GeoRedundantBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoRedundantBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) HighAvailability() PostgresqlFlexibleServerHighAvailabilityOutputReference {
	var returns PostgresqlFlexibleServerHighAvailabilityOutputReference
	_jsii_.Get(
		j,
		"highAvailability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) HighAvailabilityInput() *PostgresqlFlexibleServerHighAvailability {
	var returns *PostgresqlFlexibleServerHighAvailability
	_jsii_.Get(
		j,
		"highAvailabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Identity() PostgresqlFlexibleServerIdentityOutputReference {
	var returns PostgresqlFlexibleServerIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) IdentityInput() *PostgresqlFlexibleServerIdentity {
	var returns *PostgresqlFlexibleServerIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) MaintenanceWindow() PostgresqlFlexibleServerMaintenanceWindowOutputReference {
	var returns PostgresqlFlexibleServerMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) MaintenanceWindowInput() *PostgresqlFlexibleServerMaintenanceWindow {
	var returns *PostgresqlFlexibleServerMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) PointInTimeRestoreTimeInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeRestoreTimeInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) PointInTimeRestoreTimeInUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeRestoreTimeInUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) PrivateDnsZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) PrivateDnsZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ReplicationRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ReplicationRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SourceServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) SourceServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceServerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) StorageMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) StorageMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) StorageTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) StorageTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Timeouts() PostgresqlFlexibleServerTimeoutsOutputReference {
	var returns PostgresqlFlexibleServerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlFlexibleServer) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/postgresql_flexible_server azurerm_postgresql_flexible_server} Resource.
func NewPostgresqlFlexibleServer(scope constructs.Construct, id *string, config *PostgresqlFlexibleServerConfig) PostgresqlFlexibleServer {
	_init_.Initialize()

	if err := validateNewPostgresqlFlexibleServerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresqlFlexibleServer{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/postgresql_flexible_server azurerm_postgresql_flexible_server} Resource.
func NewPostgresqlFlexibleServer_Override(p PostgresqlFlexibleServer, scope constructs.Construct, id *string, config *PostgresqlFlexibleServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServer",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetAdministratorLogin(val *string) {
	if err := j.validateSetAdministratorLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorLogin",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetAdministratorPassword(val *string) {
	if err := j.validateSetAdministratorPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorPassword",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetAdministratorPasswordWo(val *string) {
	if err := j.validateSetAdministratorPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorPasswordWo",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetAdministratorPasswordWoVersion(val *float64) {
	if err := j.validateSetAdministratorPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorPasswordWoVersion",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetAutoGrowEnabled(val interface{}) {
	if err := j.validateSetAutoGrowEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoGrowEnabled",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetBackupRetentionDays(val *float64) {
	if err := j.validateSetBackupRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetCreateMode(val *string) {
	if err := j.validateSetCreateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createMode",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetDelegatedSubnetId(val *string) {
	if err := j.validateSetDelegatedSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delegatedSubnetId",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetGeoRedundantBackupEnabled(val interface{}) {
	if err := j.validateSetGeoRedundantBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoRedundantBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetPointInTimeRestoreTimeInUtc(val *string) {
	if err := j.validateSetPointInTimeRestoreTimeInUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTimeRestoreTimeInUtc",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetPrivateDnsZoneId(val *string) {
	if err := j.validateSetPrivateDnsZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsZoneId",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetReplicationRole(val *string) {
	if err := j.validateSetReplicationRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationRole",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetSourceServerId(val *string) {
	if err := j.validateSetSourceServerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceServerId",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetStorageMb(val *float64) {
	if err := j.validateSetStorageMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageMb",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetStorageTier(val *string) {
	if err := j.validateSetStorageTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageTier",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_PostgresqlFlexibleServer)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a PostgresqlFlexibleServer resource upon running "cdktf plan <stack-name>".
func PostgresqlFlexibleServer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePostgresqlFlexibleServer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServer",
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
func PostgresqlFlexibleServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePostgresqlFlexibleServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PostgresqlFlexibleServer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePostgresqlFlexibleServer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PostgresqlFlexibleServer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePostgresqlFlexibleServer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PostgresqlFlexibleServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.postgresqlFlexibleServer.PostgresqlFlexibleServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) PutAuthentication(value *PostgresqlFlexibleServerAuthentication) {
	if err := p.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) PutCustomerManagedKey(value *PostgresqlFlexibleServerCustomerManagedKey) {
	if err := p.validatePutCustomerManagedKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCustomerManagedKey",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) PutHighAvailability(value *PostgresqlFlexibleServerHighAvailability) {
	if err := p.validatePutHighAvailabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHighAvailability",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) PutIdentity(value *PostgresqlFlexibleServerIdentity) {
	if err := p.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIdentity",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) PutMaintenanceWindow(value *PostgresqlFlexibleServerMaintenanceWindow) {
	if err := p.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) PutTimeouts(value *PostgresqlFlexibleServerTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetAdministratorLogin() {
	_jsii_.InvokeVoid(
		p,
		"resetAdministratorLogin",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetAdministratorPassword() {
	_jsii_.InvokeVoid(
		p,
		"resetAdministratorPassword",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetAdministratorPasswordWo() {
	_jsii_.InvokeVoid(
		p,
		"resetAdministratorPasswordWo",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetAdministratorPasswordWoVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetAdministratorPasswordWoVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetAuthentication() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetAutoGrowEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoGrowEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetBackupRetentionDays() {
	_jsii_.InvokeVoid(
		p,
		"resetBackupRetentionDays",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetCreateMode() {
	_jsii_.InvokeVoid(
		p,
		"resetCreateMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetCustomerManagedKey() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomerManagedKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetDelegatedSubnetId() {
	_jsii_.InvokeVoid(
		p,
		"resetDelegatedSubnetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetGeoRedundantBackupEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetGeoRedundantBackupEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetHighAvailability() {
	_jsii_.InvokeVoid(
		p,
		"resetHighAvailability",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetIdentity() {
	_jsii_.InvokeVoid(
		p,
		"resetIdentity",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		p,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetPointInTimeRestoreTimeInUtc() {
	_jsii_.InvokeVoid(
		p,
		"resetPointInTimeRestoreTimeInUtc",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetPrivateDnsZoneId() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateDnsZoneId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetReplicationRole() {
	_jsii_.InvokeVoid(
		p,
		"resetReplicationRole",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetSkuName() {
	_jsii_.InvokeVoid(
		p,
		"resetSkuName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetSourceServerId() {
	_jsii_.InvokeVoid(
		p,
		"resetSourceServerId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetStorageMb() {
	_jsii_.InvokeVoid(
		p,
		"resetStorageMb",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetStorageTier() {
	_jsii_.InvokeVoid(
		p,
		"resetStorageTier",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ResetZone() {
	_jsii_.InvokeVoid(
		p,
		"resetZone",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlFlexibleServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlFlexibleServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

