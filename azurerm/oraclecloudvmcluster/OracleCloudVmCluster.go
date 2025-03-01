// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oraclecloudvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/oraclecloudvmcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/oracle_cloud_vm_cluster azurerm_oracle_cloud_vm_cluster}.
type OracleCloudVmCluster interface {
	cdktf.TerraformResource
	BackupSubnetCidr() *string
	SetBackupSubnetCidr(val *string)
	BackupSubnetCidrInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudExadataInfrastructureId() *string
	SetCloudExadataInfrastructureId(val *string)
	CloudExadataInfrastructureIdInput() *string
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
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
	CpuCoreCount() *float64
	SetCpuCoreCount(val *float64)
	CpuCoreCountInput() *float64
	DataCollectionOptions() OracleCloudVmClusterDataCollectionOptionsOutputReference
	DataCollectionOptionsInput() *OracleCloudVmClusterDataCollectionOptions
	DataStoragePercentage() *float64
	SetDataStoragePercentage(val *float64)
	DataStoragePercentageInput() *float64
	DataStorageSizeInTbs() *float64
	SetDataStorageSizeInTbs(val *float64)
	DataStorageSizeInTbsInput() *float64
	DbNodeStorageSizeInGbs() *float64
	SetDbNodeStorageSizeInGbs(val *float64)
	DbNodeStorageSizeInGbsInput() *float64
	DbServers() *[]*string
	SetDbServers(val *[]*string)
	DbServersInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GiVersion() *string
	SetGiVersion(val *string)
	GiVersionInput() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameActual() *string
	HostnameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalBackupEnabled() interface{}
	SetLocalBackupEnabled(val interface{})
	LocalBackupEnabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MemorySizeInGbs() *float64
	SetMemorySizeInGbs(val *float64)
	MemorySizeInGbsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Ocid() *string
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
	ScanListenerPortTcp() *float64
	SetScanListenerPortTcp(val *float64)
	ScanListenerPortTcpInput() *float64
	ScanListenerPortTcpSsl() *float64
	SetScanListenerPortTcpSsl(val *float64)
	ScanListenerPortTcpSslInput() *float64
	SparseDiskgroupEnabled() interface{}
	SetSparseDiskgroupEnabled(val interface{})
	SparseDiskgroupEnabledInput() interface{}
	SshPublicKeys() *[]*string
	SetSshPublicKeys(val *[]*string)
	SshPublicKeysInput() *[]*string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() OracleCloudVmClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	PutDataCollectionOptions(value *OracleCloudVmClusterDataCollectionOptions)
	PutTimeouts(value *OracleCloudVmClusterTimeouts)
	ResetBackupSubnetCidr()
	ResetClusterName()
	ResetDataCollectionOptions()
	ResetDataStoragePercentage()
	ResetDataStorageSizeInTbs()
	ResetDbNodeStorageSizeInGbs()
	ResetDomain()
	ResetId()
	ResetLocalBackupEnabled()
	ResetMemorySizeInGbs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScanListenerPortTcp()
	ResetScanListenerPortTcpSsl()
	ResetSparseDiskgroupEnabled()
	ResetTags()
	ResetTimeouts()
	ResetTimeZone()
	ResetZoneId()
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

// The jsii proxy struct for OracleCloudVmCluster
type jsiiProxy_OracleCloudVmCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OracleCloudVmCluster) BackupSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) BackupSubnetCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupSubnetCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) CloudExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) CloudExadataInfrastructureIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) CpuCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DataCollectionOptions() OracleCloudVmClusterDataCollectionOptionsOutputReference {
	var returns OracleCloudVmClusterDataCollectionOptionsOutputReference
	_jsii_.Get(
		j,
		"dataCollectionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DataCollectionOptionsInput() *OracleCloudVmClusterDataCollectionOptions {
	var returns *OracleCloudVmClusterDataCollectionOptions
	_jsii_.Get(
		j,
		"dataCollectionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DataStoragePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStoragePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DataStoragePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStoragePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DataStorageSizeInTbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DbNodeStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DbNodeStorageSizeInGbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DbServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DbServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) GiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) GiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) HostnameActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) LocalBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) LocalBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) MemorySizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) MemorySizeInGbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ScanListenerPortTcp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ScanListenerPortTcpInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ScanListenerPortTcpSsl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcpSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ScanListenerPortTcpSslInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcpSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) SparseDiskgroupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparseDiskgroupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) SparseDiskgroupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparseDiskgroupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) Timeouts() OracleCloudVmClusterTimeoutsOutputReference {
	var returns OracleCloudVmClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleCloudVmCluster) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/oracle_cloud_vm_cluster azurerm_oracle_cloud_vm_cluster} Resource.
func NewOracleCloudVmCluster(scope constructs.Construct, id *string, config *OracleCloudVmClusterConfig) OracleCloudVmCluster {
	_init_.Initialize()

	if err := validateNewOracleCloudVmClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleCloudVmCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.oracleCloudVmCluster.OracleCloudVmCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/oracle_cloud_vm_cluster azurerm_oracle_cloud_vm_cluster} Resource.
func NewOracleCloudVmCluster_Override(o OracleCloudVmCluster, scope constructs.Construct, id *string, config *OracleCloudVmClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.oracleCloudVmCluster.OracleCloudVmCluster",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetBackupSubnetCidr(val *string) {
	if err := j.validateSetBackupSubnetCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupSubnetCidr",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetCloudExadataInfrastructureId(val *string) {
	if err := j.validateSetCloudExadataInfrastructureIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudExadataInfrastructureId",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetCpuCoreCount(val *float64) {
	if err := j.validateSetCpuCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCoreCount",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetDataStoragePercentage(val *float64) {
	if err := j.validateSetDataStoragePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStoragePercentage",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetDataStorageSizeInTbs(val *float64) {
	if err := j.validateSetDataStorageSizeInTbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeInTbs",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetDbNodeStorageSizeInGbs(val *float64) {
	if err := j.validateSetDbNodeStorageSizeInGbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbNodeStorageSizeInGbs",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetDbServers(val *[]*string) {
	if err := j.validateSetDbServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbServers",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetGiVersion(val *string) {
	if err := j.validateSetGiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"giVersion",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetLocalBackupEnabled(val interface{}) {
	if err := j.validateSetLocalBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetMemorySizeInGbs(val *float64) {
	if err := j.validateSetMemorySizeInGbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySizeInGbs",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetScanListenerPortTcp(val *float64) {
	if err := j.validateSetScanListenerPortTcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanListenerPortTcp",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetScanListenerPortTcpSsl(val *float64) {
	if err := j.validateSetScanListenerPortTcpSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanListenerPortTcpSsl",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetSparseDiskgroupEnabled(val interface{}) {
	if err := j.validateSetSparseDiskgroupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparseDiskgroupEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetVirtualNetworkId(val *string) {
	if err := j.validateSetVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_OracleCloudVmCluster)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

// Generates CDKTF code for importing a OracleCloudVmCluster resource upon running "cdktf plan <stack-name>".
func OracleCloudVmCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOracleCloudVmCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.oracleCloudVmCluster.OracleCloudVmCluster",
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
func OracleCloudVmCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOracleCloudVmCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.oracleCloudVmCluster.OracleCloudVmCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OracleCloudVmCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOracleCloudVmCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.oracleCloudVmCluster.OracleCloudVmCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OracleCloudVmCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOracleCloudVmCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.oracleCloudVmCluster.OracleCloudVmCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OracleCloudVmCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.oracleCloudVmCluster.OracleCloudVmCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) PutDataCollectionOptions(value *OracleCloudVmClusterDataCollectionOptions) {
	if err := o.validatePutDataCollectionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDataCollectionOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) PutTimeouts(value *OracleCloudVmClusterTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetBackupSubnetCidr() {
	_jsii_.InvokeVoid(
		o,
		"resetBackupSubnetCidr",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetClusterName() {
	_jsii_.InvokeVoid(
		o,
		"resetClusterName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetDataCollectionOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetDataCollectionOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetDataStoragePercentage() {
	_jsii_.InvokeVoid(
		o,
		"resetDataStoragePercentage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetDataStorageSizeInTbs() {
	_jsii_.InvokeVoid(
		o,
		"resetDataStorageSizeInTbs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetDbNodeStorageSizeInGbs() {
	_jsii_.InvokeVoid(
		o,
		"resetDbNodeStorageSizeInGbs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetDomain() {
	_jsii_.InvokeVoid(
		o,
		"resetDomain",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetLocalBackupEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetLocalBackupEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetMemorySizeInGbs() {
	_jsii_.InvokeVoid(
		o,
		"resetMemorySizeInGbs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetScanListenerPortTcp() {
	_jsii_.InvokeVoid(
		o,
		"resetScanListenerPortTcp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetScanListenerPortTcpSsl() {
	_jsii_.InvokeVoid(
		o,
		"resetScanListenerPortTcpSsl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetSparseDiskgroupEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetSparseDiskgroupEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetTimeZone() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) ResetZoneId() {
	_jsii_.InvokeVoid(
		o,
		"resetZoneId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleCloudVmCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleCloudVmCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

