// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermoracleautonomousdatabaseclonefromdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/dataazurermoracleautonomousdatabaseclonefromdatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/data-sources/oracle_autonomous_database_clone_from_database azurerm_oracle_autonomous_database_clone_from_database}.
type DataAzurermOracleAutonomousDatabaseCloneFromDatabase interface {
	cdktf.TerraformDataSource
	ActualUsedDataStorageSizeInTb() *float64
	AllocatedStorageSizeInTb() *float64
	AllowedIpAddresses() *[]*string
	AutoScalingEnabled() cdktf.IResolvable
	AutoScalingForStorageEnabled() cdktf.IResolvable
	AvailableUpgradeVersions() *[]*string
	BackupRetentionPeriodInDays() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CharacterSet() *string
	ComputeCount() *float64
	ComputeModel() *string
	ConnectionStrings() *[]*string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuCoreCount() *float64
	CustomerContacts() *[]*string
	DatabaseVersion() *string
	DatabaseWorkload() *string
	DataStorageSizeInGb() *float64
	DataStorageSizeInTb() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	FailedDataRecoveryInSeconds() *float64
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
	InMemoryAreaInGb() *float64
	LicenseModel() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecycleDetails() *string
	LifecycleState() *string
	LocalAdgAutoFailoverMaxDataLossLimitInSeconds() *float64
	LocalDataGuardEnabled() cdktf.IResolvable
	Location() *string
	LongTermBackupSchedule() DataAzurermOracleAutonomousDatabaseCloneFromDatabaseLongTermBackupScheduleList
	MemoryPerOracleComputeUnitInGb() *float64
	MtlsConnectionRequired() cdktf.IResolvable
	Name() *string
	SetName(val *string)
	NameInput() *string
	NationalCharacterSet() *string
	NextLongTermBackupTimestamp() *string
	// The tree node.
	Node() constructs.Node
	Ocid() *string
	OciUrl() *string
	PeerDatabaseIds() *[]*string
	Preview() cdktf.IResolvable
	PreviewVersionWithServiceTermsAccepted() cdktf.IResolvable
	PrivateEndpointIp() *string
	PrivateEndpointLabel() *string
	PrivateEndpointUrl() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProvisionableCpus() *[]*float64
	// Experimental.
	RawOverrides() interface{}
	ReconnectCloneEnabled() cdktf.IResolvable
	RefreshableClone() cdktf.IResolvable
	RefreshableStatus() *string
	RemoteDataGuardEnabled() cdktf.IResolvable
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ServiceConsoleUrl() *string
	SourceAutonomousDatabaseId() *string
	SqlWebDeveloperUrl() *string
	SubnetId() *string
	SupportedRegionsToCloneTo() *[]*string
	Tags() cdktf.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeCreatedInUtc() *string
	TimeDataGuardRoleChangedInUtc() *string
	TimeDeletionOfFreeAutonomousDatabaseInUtc() *string
	TimeLocalDataGuardEnabledInUtc() *string
	TimeMaintenanceBeginInUtc() *string
	TimeMaintenanceEndInUtc() *string
	TimeOfLastFailoverInUtc() *string
	TimeOfLastRefreshInUtc() *string
	TimeOfLastRefreshPointInUtc() *string
	TimeOfLastSwitchoverInUtc() *string
	Timeouts() DataAzurermOracleAutonomousDatabaseCloneFromDatabaseTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeReclamationOfFreeAutonomousDatabaseInUtc() *string
	TimeUntilReconnectInUtc() *string
	UsedDataStorageSizeInGb() *float64
	UsedDataStorageSizeInTb() *float64
	VirtualNetworkId() *string
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
	PutTimeouts(value *DataAzurermOracleAutonomousDatabaseCloneFromDatabaseTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAzurermOracleAutonomousDatabaseCloneFromDatabase
type jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ActualUsedDataStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actualUsedDataStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) AllocatedStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) AllowedIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) AutoScalingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) AutoScalingForStorageEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingForStorageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) AvailableUpgradeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableUpgradeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) BackupRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ConnectionStrings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) CustomerContacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customerContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) DatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) DatabaseWorkload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseWorkload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) DataStorageSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) DataStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) FailedDataRecoveryInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failedDataRecoveryInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) InMemoryAreaInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inMemoryAreaInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) LocalAdgAutoFailoverMaxDataLossLimitInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localAdgAutoFailoverMaxDataLossLimitInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) LocalDataGuardEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"localDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) LongTermBackupSchedule() DataAzurermOracleAutonomousDatabaseCloneFromDatabaseLongTermBackupScheduleList {
	var returns DataAzurermOracleAutonomousDatabaseCloneFromDatabaseLongTermBackupScheduleList
	_jsii_.Get(
		j,
		"longTermBackupSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) MemoryPerOracleComputeUnitInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) MtlsConnectionRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mtlsConnectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) NationalCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nationalCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) NextLongTermBackupTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextLongTermBackupTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) PeerDatabaseIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerDatabaseIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Preview() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"preview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) PreviewVersionWithServiceTermsAccepted() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"previewVersionWithServiceTermsAccepted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) PrivateEndpointIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) PrivateEndpointLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) PrivateEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ProvisionableCpus() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"provisionableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ReconnectCloneEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconnectCloneEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) RefreshableClone() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"refreshableClone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) RefreshableStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) RemoteDataGuardEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"remoteDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ServiceConsoleUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceConsoleUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) SourceAutonomousDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAutonomousDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) SqlWebDeveloperUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlWebDeveloperUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) SupportedRegionsToCloneTo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedRegionsToCloneTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeCreatedInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeCreatedInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeDataGuardRoleChangedInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeDataGuardRoleChangedInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeDeletionOfFreeAutonomousDatabaseInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeDeletionOfFreeAutonomousDatabaseInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeLocalDataGuardEnabledInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeLocalDataGuardEnabledInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeMaintenanceBeginInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceBeginInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeMaintenanceEndInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceEndInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeOfLastFailoverInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastFailoverInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeOfLastRefreshInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastRefreshInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeOfLastRefreshPointInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastRefreshPointInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeOfLastSwitchoverInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastSwitchoverInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) Timeouts() DataAzurermOracleAutonomousDatabaseCloneFromDatabaseTimeoutsOutputReference {
	var returns DataAzurermOracleAutonomousDatabaseCloneFromDatabaseTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeReclamationOfFreeAutonomousDatabaseInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeReclamationOfFreeAutonomousDatabaseInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) TimeUntilReconnectInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUntilReconnectInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) UsedDataStorageSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) UsedDataStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/data-sources/oracle_autonomous_database_clone_from_database azurerm_oracle_autonomous_database_clone_from_database} Data Source.
func NewDataAzurermOracleAutonomousDatabaseCloneFromDatabase(scope constructs.Construct, id *string, config *DataAzurermOracleAutonomousDatabaseCloneFromDatabaseConfig) DataAzurermOracleAutonomousDatabaseCloneFromDatabase {
	_init_.Initialize()

	if err := validateNewDataAzurermOracleAutonomousDatabaseCloneFromDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/data-sources/oracle_autonomous_database_clone_from_database azurerm_oracle_autonomous_database_clone_from_database} Data Source.
func NewDataAzurermOracleAutonomousDatabaseCloneFromDatabase_Override(d DataAzurermOracleAutonomousDatabaseCloneFromDatabase, scope constructs.Construct, id *string, config *DataAzurermOracleAutonomousDatabaseCloneFromDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabase",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

// Generates CDKTF code for importing a DataAzurermOracleAutonomousDatabaseCloneFromDatabase resource upon running "cdktf plan <stack-name>".
func DataAzurermOracleAutonomousDatabaseCloneFromDatabase_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCloneFromDatabase_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabase",
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
func DataAzurermOracleAutonomousDatabaseCloneFromDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCloneFromDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermOracleAutonomousDatabaseCloneFromDatabase_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCloneFromDatabase_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabase",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermOracleAutonomousDatabaseCloneFromDatabase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCloneFromDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermOracleAutonomousDatabaseCloneFromDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromDatabase.DataAzurermOracleAutonomousDatabaseCloneFromDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) PutTimeouts(value *DataAzurermOracleAutonomousDatabaseCloneFromDatabaseTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

