// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermoracleautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/dataazurermoracleautonomousdatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/data-sources/oracle_autonomous_database azurerm_oracle_autonomous_database}.
type DataAzurermOracleAutonomousDatabase interface {
	cdktf.TerraformDataSource
	ActualUsedDataStorageSizeInTbs() *float64
	AllocatedStorageSizeInTbs() *float64
	AllowedIps() *[]*float64
	AutonomousDatabaseId() *string
	AutoScalingEnabled() cdktf.IResolvable
	AutoScalingForStorageEnabled() cdktf.IResolvable
	AvailableUpgradeVersions() *[]*string
	BackupRetentionPeriodInDays() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CharacterSet() *string
	ComputeCount() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuCoreCount() *float64
	DataStorageSizeInGbs() *float64
	DataStorageSizeInTbs() *float64
	DbNodeStorageSizeInGbs() *float64
	DbVersion() *string
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
	InMemoryAreaInGbs() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecycleDetails() *string
	LocalAdgAutoFailoverMaxDataLossLimit() *float64
	LocalDataGuardEnabled() cdktf.IResolvable
	Location() *string
	MemoryPerOracleComputeUnitInGbs() *float64
	MtlsConnectionRequired() cdktf.IResolvable
	Name() *string
	SetName(val *string)
	NameInput() *string
	NationalCharacterSet() *string
	NextLongTermBackupTimeStamp() *string
	// The tree node.
	Node() constructs.Node
	Ocid() *string
	OciUrl() *string
	PeerDbId() *string
	PeerDbIds() *[]*string
	Preview() cdktf.IResolvable
	PreviewVersionWithServiceTermsAccepted() cdktf.IResolvable
	PrivateEndpoint() *string
	PrivateEndpointIp() *string
	PrivateEndpointLabel() *string
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
	TimeCreated() *string
	TimeDataGuardRoleChanged() *string
	TimeDeletionOfFreeAutonomousDatabase() *string
	TimeLocalDataGuardEnabledOn() *string
	TimeMaintenanceBegin() *string
	TimeMaintenanceEnd() *string
	TimeOfLastFailover() *string
	TimeOfLastRefresh() *string
	TimeOfLastRefreshPoint() *string
	TimeOfLastSwitchover() *string
	Timeouts() DataAzurermOracleAutonomousDatabaseTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeReclamationOfFreeAutonomousDatabase() *string
	UsedDataStorageSizeInGbs() *float64
	UsedDataStorageSizeInTbs() *float64
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
	PutTimeouts(value *DataAzurermOracleAutonomousDatabaseTimeouts)
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

// The jsii proxy struct for DataAzurermOracleAutonomousDatabase
type jsiiProxy_DataAzurermOracleAutonomousDatabase struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) ActualUsedDataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actualUsedDataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) AllocatedStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) AllowedIps() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) AutonomousDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autonomousDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) AutoScalingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) AutoScalingForStorageEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingForStorageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) AvailableUpgradeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableUpgradeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) BackupRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) DataStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) DataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) DbNodeStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) DbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) FailedDataRecoveryInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failedDataRecoveryInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) InMemoryAreaInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inMemoryAreaInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) LocalAdgAutoFailoverMaxDataLossLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localAdgAutoFailoverMaxDataLossLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) LocalDataGuardEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"localDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) MemoryPerOracleComputeUnitInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) MtlsConnectionRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mtlsConnectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) NationalCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nationalCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) NextLongTermBackupTimeStamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextLongTermBackupTimeStamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) PeerDbId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerDbId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) PeerDbIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerDbIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Preview() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"preview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) PreviewVersionWithServiceTermsAccepted() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"previewVersionWithServiceTermsAccepted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) PrivateEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) PrivateEndpointIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) PrivateEndpointLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) ProvisionableCpus() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"provisionableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) RemoteDataGuardEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"remoteDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) ServiceConsoleUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceConsoleUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) SqlWebDeveloperUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlWebDeveloperUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) SupportedRegionsToCloneTo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedRegionsToCloneTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeCreated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeCreated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeDataGuardRoleChanged() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeDataGuardRoleChanged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeDeletionOfFreeAutonomousDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeDeletionOfFreeAutonomousDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeLocalDataGuardEnabledOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeLocalDataGuardEnabledOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeMaintenanceBegin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceBegin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeMaintenanceEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeOfLastFailover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastFailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeOfLastRefresh() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastRefresh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeOfLastRefreshPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastRefreshPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeOfLastSwitchover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastSwitchover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) Timeouts() DataAzurermOracleAutonomousDatabaseTimeoutsOutputReference {
	var returns DataAzurermOracleAutonomousDatabaseTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) TimeReclamationOfFreeAutonomousDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeReclamationOfFreeAutonomousDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) UsedDataStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) UsedDataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/data-sources/oracle_autonomous_database azurerm_oracle_autonomous_database} Data Source.
func NewDataAzurermOracleAutonomousDatabase(scope constructs.Construct, id *string, config *DataAzurermOracleAutonomousDatabaseConfig) DataAzurermOracleAutonomousDatabase {
	_init_.Initialize()

	if err := validateNewDataAzurermOracleAutonomousDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermOracleAutonomousDatabase{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabase.DataAzurermOracleAutonomousDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/data-sources/oracle_autonomous_database azurerm_oracle_autonomous_database} Data Source.
func NewDataAzurermOracleAutonomousDatabase_Override(d DataAzurermOracleAutonomousDatabase, scope constructs.Construct, id *string, config *DataAzurermOracleAutonomousDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabase.DataAzurermOracleAutonomousDatabase",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabase)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

// Generates CDKTF code for importing a DataAzurermOracleAutonomousDatabase resource upon running "cdktf plan <stack-name>".
func DataAzurermOracleAutonomousDatabase_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabase_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabase.DataAzurermOracleAutonomousDatabase",
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
func DataAzurermOracleAutonomousDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabase.DataAzurermOracleAutonomousDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermOracleAutonomousDatabase_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabase_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabase.DataAzurermOracleAutonomousDatabase",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermOracleAutonomousDatabase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabase.DataAzurermOracleAutonomousDatabase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermOracleAutonomousDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermOracleAutonomousDatabase.DataAzurermOracleAutonomousDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) PutTimeouts(value *DataAzurermOracleAutonomousDatabaseTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

