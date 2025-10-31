// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermoracleautonomousdatabaseclonefrombackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/dataazurermoracleautonomousdatabaseclonefrombackup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/data-sources/oracle_autonomous_database_clone_from_backup azurerm_oracle_autonomous_database_clone_from_backup}.
type DataAzurermOracleAutonomousDatabaseCloneFromBackup interface {
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
	LongTermBackupSchedule() DataAzurermOracleAutonomousDatabaseCloneFromBackupLongTermBackupScheduleList
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
	Timeouts() DataAzurermOracleAutonomousDatabaseCloneFromBackupTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeReclamationOfFreeAutonomousDatabaseInUtc() *string
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
	PutTimeouts(value *DataAzurermOracleAutonomousDatabaseCloneFromBackupTimeouts)
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

// The jsii proxy struct for DataAzurermOracleAutonomousDatabaseCloneFromBackup
type jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ActualUsedDataStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actualUsedDataStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) AllocatedStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) AllowedIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) AutoScalingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) AutoScalingForStorageEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingForStorageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) AvailableUpgradeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableUpgradeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) BackupRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ConnectionStrings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) CustomerContacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customerContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) DatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) DatabaseWorkload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseWorkload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) DataStorageSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) DataStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) FailedDataRecoveryInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failedDataRecoveryInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) InMemoryAreaInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inMemoryAreaInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) LocalAdgAutoFailoverMaxDataLossLimitInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localAdgAutoFailoverMaxDataLossLimitInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) LocalDataGuardEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"localDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) LongTermBackupSchedule() DataAzurermOracleAutonomousDatabaseCloneFromBackupLongTermBackupScheduleList {
	var returns DataAzurermOracleAutonomousDatabaseCloneFromBackupLongTermBackupScheduleList
	_jsii_.Get(
		j,
		"longTermBackupSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) MemoryPerOracleComputeUnitInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) MtlsConnectionRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mtlsConnectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) NationalCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nationalCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) NextLongTermBackupTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextLongTermBackupTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) PeerDatabaseIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerDatabaseIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Preview() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"preview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) PreviewVersionWithServiceTermsAccepted() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"previewVersionWithServiceTermsAccepted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) PrivateEndpointIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) PrivateEndpointLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) PrivateEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ProvisionableCpus() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"provisionableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) RemoteDataGuardEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"remoteDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ServiceConsoleUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceConsoleUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) SourceAutonomousDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAutonomousDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) SqlWebDeveloperUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlWebDeveloperUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) SupportedRegionsToCloneTo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedRegionsToCloneTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeCreatedInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeCreatedInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeDataGuardRoleChangedInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeDataGuardRoleChangedInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeDeletionOfFreeAutonomousDatabaseInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeDeletionOfFreeAutonomousDatabaseInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeLocalDataGuardEnabledInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeLocalDataGuardEnabledInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeMaintenanceBeginInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceBeginInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeMaintenanceEndInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceEndInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeOfLastFailoverInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastFailoverInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeOfLastRefreshInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastRefreshInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeOfLastRefreshPointInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastRefreshPointInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeOfLastSwitchoverInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastSwitchoverInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) Timeouts() DataAzurermOracleAutonomousDatabaseCloneFromBackupTimeoutsOutputReference {
	var returns DataAzurermOracleAutonomousDatabaseCloneFromBackupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) TimeReclamationOfFreeAutonomousDatabaseInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeReclamationOfFreeAutonomousDatabaseInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) UsedDataStorageSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) UsedDataStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/data-sources/oracle_autonomous_database_clone_from_backup azurerm_oracle_autonomous_database_clone_from_backup} Data Source.
func NewDataAzurermOracleAutonomousDatabaseCloneFromBackup(scope constructs.Construct, id *string, config *DataAzurermOracleAutonomousDatabaseCloneFromBackupConfig) DataAzurermOracleAutonomousDatabaseCloneFromBackup {
	_init_.Initialize()

	if err := validateNewDataAzurermOracleAutonomousDatabaseCloneFromBackupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromBackup.DataAzurermOracleAutonomousDatabaseCloneFromBackup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/data-sources/oracle_autonomous_database_clone_from_backup azurerm_oracle_autonomous_database_clone_from_backup} Data Source.
func NewDataAzurermOracleAutonomousDatabaseCloneFromBackup_Override(d DataAzurermOracleAutonomousDatabaseCloneFromBackup, scope constructs.Construct, id *string, config *DataAzurermOracleAutonomousDatabaseCloneFromBackupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromBackup.DataAzurermOracleAutonomousDatabaseCloneFromBackup",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

// Generates CDKTF code for importing a DataAzurermOracleAutonomousDatabaseCloneFromBackup resource upon running "cdktf plan <stack-name>".
func DataAzurermOracleAutonomousDatabaseCloneFromBackup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCloneFromBackup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromBackup.DataAzurermOracleAutonomousDatabaseCloneFromBackup",
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
func DataAzurermOracleAutonomousDatabaseCloneFromBackup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCloneFromBackup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromBackup.DataAzurermOracleAutonomousDatabaseCloneFromBackup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermOracleAutonomousDatabaseCloneFromBackup_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCloneFromBackup_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromBackup.DataAzurermOracleAutonomousDatabaseCloneFromBackup",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermOracleAutonomousDatabaseCloneFromBackup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCloneFromBackup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromBackup.DataAzurermOracleAutonomousDatabaseCloneFromBackup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermOracleAutonomousDatabaseCloneFromBackup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabaseCloneFromBackup.DataAzurermOracleAutonomousDatabaseCloneFromBackup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) PutTimeouts(value *DataAzurermOracleAutonomousDatabaseCloneFromBackupTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCloneFromBackup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

