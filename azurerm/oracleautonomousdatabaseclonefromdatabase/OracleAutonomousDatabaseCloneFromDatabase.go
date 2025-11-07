// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracleautonomousdatabaseclonefromdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/oracleautonomousdatabaseclonefromdatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/oracle_autonomous_database_clone_from_database azurerm_oracle_autonomous_database_clone_from_database}.
type OracleAutonomousDatabaseCloneFromDatabase interface {
	cdktf.TerraformResource
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AllowedIpAddresses() *[]*string
	SetAllowedIpAddresses(val *[]*string)
	AllowedIpAddressesInput() *[]*string
	AutoScalingEnabled() interface{}
	SetAutoScalingEnabled(val interface{})
	AutoScalingEnabledInput() interface{}
	AutoScalingForStorageEnabled() interface{}
	SetAutoScalingForStorageEnabled(val interface{})
	AutoScalingForStorageEnabledInput() interface{}
	BackupRetentionPeriodInDays() *float64
	SetBackupRetentionPeriodInDays(val *float64)
	BackupRetentionPeriodInDaysInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CharacterSet() *string
	SetCharacterSet(val *string)
	CharacterSetInput() *string
	CloneType() *string
	SetCloneType(val *string)
	CloneTypeInput() *string
	ComputeCount() *float64
	SetComputeCount(val *float64)
	ComputeCountInput() *float64
	ComputeModel() *string
	SetComputeModel(val *string)
	ComputeModelInput() *string
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
	CustomerContacts() *[]*string
	SetCustomerContacts(val *[]*string)
	CustomerContactsInput() *[]*string
	DatabaseVersion() *string
	SetDatabaseVersion(val *string)
	DatabaseVersionInput() *string
	DatabaseWorkload() *string
	SetDatabaseWorkload(val *string)
	DatabaseWorkloadInput() *string
	DataStorageSizeInTb() *float64
	SetDataStorageSizeInTb(val *float64)
	DataStorageSizeInTbInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
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
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MtlsConnectionRequired() interface{}
	SetMtlsConnectionRequired(val interface{})
	MtlsConnectionRequiredInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NationalCharacterSet() *string
	SetNationalCharacterSet(val *string)
	NationalCharacterSetInput() *string
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
	RefreshableModel() *string
	SetRefreshableModel(val *string)
	RefreshableModelInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SourceAutonomousDatabaseId() *string
	SetSourceAutonomousDatabaseId(val *string)
	SourceAutonomousDatabaseIdInput() *string
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
	Timeouts() OracleAutonomousDatabaseCloneFromDatabaseTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
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
	PutTimeouts(value *OracleAutonomousDatabaseCloneFromDatabaseTimeouts)
	ResetAllowedIpAddresses()
	ResetCustomerContacts()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRefreshableModel()
	ResetSubnetId()
	ResetTags()
	ResetTimeouts()
	ResetVirtualNetworkId()
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

// The jsii proxy struct for OracleAutonomousDatabaseCloneFromDatabase
type jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) AllowedIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) AllowedIpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) AutoScalingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) AutoScalingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) AutoScalingForStorageEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingForStorageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) AutoScalingForStorageEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingForStorageEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) BackupRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) BackupRetentionPeriodInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) CharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) CloneType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) CloneTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ComputeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ComputeModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) CustomerContacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customerContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) CustomerContactsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customerContactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) DatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) DatabaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) DatabaseWorkload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseWorkload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) DatabaseWorkloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseWorkloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) DataStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) DataStorageSizeInTbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) MtlsConnectionRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mtlsConnectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) MtlsConnectionRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mtlsConnectionRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) NationalCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nationalCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) NationalCharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nationalCharacterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) RefreshableModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) RefreshableModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) SourceAutonomousDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAutonomousDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) SourceAutonomousDatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAutonomousDatabaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) Timeouts() OracleAutonomousDatabaseCloneFromDatabaseTimeoutsOutputReference {
	var returns OracleAutonomousDatabaseCloneFromDatabaseTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/oracle_autonomous_database_clone_from_database azurerm_oracle_autonomous_database_clone_from_database} Resource.
func NewOracleAutonomousDatabaseCloneFromDatabase(scope constructs.Construct, id *string, config *OracleAutonomousDatabaseCloneFromDatabaseConfig) OracleAutonomousDatabaseCloneFromDatabase {
	_init_.Initialize()

	if err := validateNewOracleAutonomousDatabaseCloneFromDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.oracleAutonomousDatabaseCloneFromDatabase.OracleAutonomousDatabaseCloneFromDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/oracle_autonomous_database_clone_from_database azurerm_oracle_autonomous_database_clone_from_database} Resource.
func NewOracleAutonomousDatabaseCloneFromDatabase_Override(o OracleAutonomousDatabaseCloneFromDatabase, scope constructs.Construct, id *string, config *OracleAutonomousDatabaseCloneFromDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.oracleAutonomousDatabaseCloneFromDatabase.OracleAutonomousDatabaseCloneFromDatabase",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetAllowedIpAddresses(val *[]*string) {
	if err := j.validateSetAllowedIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIpAddresses",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetAutoScalingEnabled(val interface{}) {
	if err := j.validateSetAutoScalingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetAutoScalingForStorageEnabled(val interface{}) {
	if err := j.validateSetAutoScalingForStorageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingForStorageEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetBackupRetentionPeriodInDays(val *float64) {
	if err := j.validateSetBackupRetentionPeriodInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriodInDays",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetCharacterSet(val *string) {
	if err := j.validateSetCharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSet",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetCloneType(val *string) {
	if err := j.validateSetCloneTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloneType",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetComputeCount(val *float64) {
	if err := j.validateSetComputeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeCount",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetComputeModel(val *string) {
	if err := j.validateSetComputeModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeModel",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetCustomerContacts(val *[]*string) {
	if err := j.validateSetCustomerContactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerContacts",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetDatabaseVersion(val *string) {
	if err := j.validateSetDatabaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseVersion",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetDatabaseWorkload(val *string) {
	if err := j.validateSetDatabaseWorkloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseWorkload",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetDataStorageSizeInTb(val *float64) {
	if err := j.validateSetDataStorageSizeInTbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeInTb",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetMtlsConnectionRequired(val interface{}) {
	if err := j.validateSetMtlsConnectionRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtlsConnectionRequired",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetNationalCharacterSet(val *string) {
	if err := j.validateSetNationalCharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nationalCharacterSet",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetRefreshableModel(val *string) {
	if err := j.validateSetRefreshableModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshableModel",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetSourceAutonomousDatabaseId(val *string) {
	if err := j.validateSetSourceAutonomousDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAutonomousDatabaseId",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase)SetVirtualNetworkId(val *string) {
	if err := j.validateSetVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

// Generates CDKTF code for importing a OracleAutonomousDatabaseCloneFromDatabase resource upon running "cdktf plan <stack-name>".
func OracleAutonomousDatabaseCloneFromDatabase_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOracleAutonomousDatabaseCloneFromDatabase_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.oracleAutonomousDatabaseCloneFromDatabase.OracleAutonomousDatabaseCloneFromDatabase",
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
func OracleAutonomousDatabaseCloneFromDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOracleAutonomousDatabaseCloneFromDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.oracleAutonomousDatabaseCloneFromDatabase.OracleAutonomousDatabaseCloneFromDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OracleAutonomousDatabaseCloneFromDatabase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOracleAutonomousDatabaseCloneFromDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.oracleAutonomousDatabaseCloneFromDatabase.OracleAutonomousDatabaseCloneFromDatabase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OracleAutonomousDatabaseCloneFromDatabase_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOracleAutonomousDatabaseCloneFromDatabase_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.oracleAutonomousDatabaseCloneFromDatabase.OracleAutonomousDatabaseCloneFromDatabase",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OracleAutonomousDatabaseCloneFromDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.oracleAutonomousDatabaseCloneFromDatabase.OracleAutonomousDatabaseCloneFromDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) PutTimeouts(value *OracleAutonomousDatabaseCloneFromDatabaseTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ResetAllowedIpAddresses() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowedIpAddresses",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ResetCustomerContacts() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomerContacts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ResetRefreshableModel() {
	_jsii_.InvokeVoid(
		o,
		"resetRefreshableModel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ResetSubnetId() {
	_jsii_.InvokeVoid(
		o,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		o,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleAutonomousDatabaseCloneFromDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

