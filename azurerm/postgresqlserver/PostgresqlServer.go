// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/postgresqlserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/postgresql_server azurerm_postgresql_server}.
type PostgresqlServer interface {
	cdktf.TerraformResource
	AdministratorLogin() *string
	SetAdministratorLogin(val *string)
	AdministratorLoginInput() *string
	AdministratorLoginPassword() *string
	SetAdministratorLoginPassword(val *string)
	AdministratorLoginPasswordInput() *string
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
	CreationSourceServerId() *string
	SetCreationSourceServerId(val *string)
	CreationSourceServerIdInput() *string
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
	Id() *string
	SetId(val *string)
	Identity() PostgresqlServerIdentityOutputReference
	IdentityInput() *PostgresqlServerIdentity
	IdInput() *string
	InfrastructureEncryptionEnabled() interface{}
	SetInfrastructureEncryptionEnabled(val interface{})
	InfrastructureEncryptionEnabledInput() interface{}
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
	RestorePointInTime() *string
	SetRestorePointInTime(val *string)
	RestorePointInTimeInput() *string
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	SslEnforcementEnabled() interface{}
	SetSslEnforcementEnabled(val interface{})
	SslEnforcementEnabledInput() interface{}
	SslMinimalTlsVersionEnforced() *string
	SetSslMinimalTlsVersionEnforced(val *string)
	SslMinimalTlsVersionEnforcedInput() *string
	StorageMb() *float64
	SetStorageMb(val *float64)
	StorageMbInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThreatDetectionPolicy() PostgresqlServerThreatDetectionPolicyOutputReference
	ThreatDetectionPolicyInput() *PostgresqlServerThreatDetectionPolicy
	Timeouts() PostgresqlServerTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutIdentity(value *PostgresqlServerIdentity)
	PutThreatDetectionPolicy(value *PostgresqlServerThreatDetectionPolicy)
	PutTimeouts(value *PostgresqlServerTimeouts)
	ResetAdministratorLogin()
	ResetAdministratorLoginPassword()
	ResetAutoGrowEnabled()
	ResetBackupRetentionDays()
	ResetCreateMode()
	ResetCreationSourceServerId()
	ResetGeoRedundantBackupEnabled()
	ResetId()
	ResetIdentity()
	ResetInfrastructureEncryptionEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetRestorePointInTime()
	ResetSslMinimalTlsVersionEnforced()
	ResetStorageMb()
	ResetTags()
	ResetThreatDetectionPolicy()
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

// The jsii proxy struct for PostgresqlServer
type jsiiProxy_PostgresqlServer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PostgresqlServer) AdministratorLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) AdministratorLoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) AdministratorLoginPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) AdministratorLoginPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) AutoGrowEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoGrowEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) AutoGrowEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoGrowEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) BackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) BackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) CreateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) CreateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) CreationSourceServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationSourceServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) CreationSourceServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationSourceServerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) GeoRedundantBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoRedundantBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) GeoRedundantBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoRedundantBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Identity() PostgresqlServerIdentityOutputReference {
	var returns PostgresqlServerIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) IdentityInput() *PostgresqlServerIdentity {
	var returns *PostgresqlServerIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) InfrastructureEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infrastructureEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) InfrastructureEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infrastructureEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) RestorePointInTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restorePointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) RestorePointInTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restorePointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) SslEnforcementEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnforcementEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) SslEnforcementEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslEnforcementEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) SslMinimalTlsVersionEnforced() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMinimalTlsVersionEnforced",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) SslMinimalTlsVersionEnforcedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMinimalTlsVersionEnforcedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) StorageMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) StorageMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) ThreatDetectionPolicy() PostgresqlServerThreatDetectionPolicyOutputReference {
	var returns PostgresqlServerThreatDetectionPolicyOutputReference
	_jsii_.Get(
		j,
		"threatDetectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) ThreatDetectionPolicyInput() *PostgresqlServerThreatDetectionPolicy {
	var returns *PostgresqlServerThreatDetectionPolicy
	_jsii_.Get(
		j,
		"threatDetectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Timeouts() PostgresqlServerTimeoutsOutputReference {
	var returns PostgresqlServerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServer) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/postgresql_server azurerm_postgresql_server} Resource.
func NewPostgresqlServer(scope constructs.Construct, id *string, config *PostgresqlServerConfig) PostgresqlServer {
	_init_.Initialize()

	if err := validateNewPostgresqlServerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresqlServer{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlServer.PostgresqlServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/postgresql_server azurerm_postgresql_server} Resource.
func NewPostgresqlServer_Override(p PostgresqlServer, scope constructs.Construct, id *string, config *PostgresqlServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.postgresqlServer.PostgresqlServer",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetAdministratorLogin(val *string) {
	if err := j.validateSetAdministratorLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorLogin",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetAdministratorLoginPassword(val *string) {
	if err := j.validateSetAdministratorLoginPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorLoginPassword",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetAutoGrowEnabled(val interface{}) {
	if err := j.validateSetAutoGrowEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoGrowEnabled",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetBackupRetentionDays(val *float64) {
	if err := j.validateSetBackupRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetCreateMode(val *string) {
	if err := j.validateSetCreateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createMode",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetCreationSourceServerId(val *string) {
	if err := j.validateSetCreationSourceServerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationSourceServerId",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetGeoRedundantBackupEnabled(val interface{}) {
	if err := j.validateSetGeoRedundantBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoRedundantBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetInfrastructureEncryptionEnabled(val interface{}) {
	if err := j.validateSetInfrastructureEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetRestorePointInTime(val *string) {
	if err := j.validateSetRestorePointInTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restorePointInTime",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetSslEnforcementEnabled(val interface{}) {
	if err := j.validateSetSslEnforcementEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslEnforcementEnabled",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetSslMinimalTlsVersionEnforced(val *string) {
	if err := j.validateSetSslMinimalTlsVersionEnforcedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslMinimalTlsVersionEnforced",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetStorageMb(val *float64) {
	if err := j.validateSetStorageMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageMb",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServer)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a PostgresqlServer resource upon running "cdktf plan <stack-name>".
func PostgresqlServer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePostgresqlServer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.postgresqlServer.PostgresqlServer",
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
func PostgresqlServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePostgresqlServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.postgresqlServer.PostgresqlServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PostgresqlServer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePostgresqlServer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.postgresqlServer.PostgresqlServer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PostgresqlServer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePostgresqlServer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.postgresqlServer.PostgresqlServer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PostgresqlServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.postgresqlServer.PostgresqlServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PostgresqlServer) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PostgresqlServer) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PostgresqlServer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PostgresqlServer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PostgresqlServer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PostgresqlServer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PostgresqlServer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PostgresqlServer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PostgresqlServer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PostgresqlServer) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PostgresqlServer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PostgresqlServer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PostgresqlServer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PostgresqlServer) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PostgresqlServer) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PostgresqlServer) PutIdentity(value *PostgresqlServerIdentity) {
	if err := p.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIdentity",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlServer) PutThreatDetectionPolicy(value *PostgresqlServerThreatDetectionPolicy) {
	if err := p.validatePutThreatDetectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putThreatDetectionPolicy",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlServer) PutTimeouts(value *PostgresqlServerTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetAdministratorLogin() {
	_jsii_.InvokeVoid(
		p,
		"resetAdministratorLogin",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetAdministratorLoginPassword() {
	_jsii_.InvokeVoid(
		p,
		"resetAdministratorLoginPassword",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetAutoGrowEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoGrowEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetBackupRetentionDays() {
	_jsii_.InvokeVoid(
		p,
		"resetBackupRetentionDays",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetCreateMode() {
	_jsii_.InvokeVoid(
		p,
		"resetCreateMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetCreationSourceServerId() {
	_jsii_.InvokeVoid(
		p,
		"resetCreationSourceServerId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetGeoRedundantBackupEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetGeoRedundantBackupEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetIdentity() {
	_jsii_.InvokeVoid(
		p,
		"resetIdentity",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetInfrastructureEncryptionEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetInfrastructureEncryptionEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetRestorePointInTime() {
	_jsii_.InvokeVoid(
		p,
		"resetRestorePointInTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetSslMinimalTlsVersionEnforced() {
	_jsii_.InvokeVoid(
		p,
		"resetSslMinimalTlsVersionEnforced",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetStorageMb() {
	_jsii_.InvokeVoid(
		p,
		"resetStorageMb",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetThreatDetectionPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetThreatDetectionPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

