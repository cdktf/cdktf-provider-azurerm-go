// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/mssqlvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/mssql_virtual_machine azurerm_mssql_virtual_machine}.
type MssqlVirtualMachine interface {
	cdktf.TerraformResource
	Assessment() MssqlVirtualMachineAssessmentOutputReference
	AssessmentInput() *MssqlVirtualMachineAssessment
	AutoBackup() MssqlVirtualMachineAutoBackupOutputReference
	AutoBackupInput() *MssqlVirtualMachineAutoBackup
	AutoPatching() MssqlVirtualMachineAutoPatchingOutputReference
	AutoPatchingInput() *MssqlVirtualMachineAutoPatching
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
	KeyVaultCredential() MssqlVirtualMachineKeyVaultCredentialOutputReference
	KeyVaultCredentialInput() *MssqlVirtualMachineKeyVaultCredential
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	RServicesEnabled() interface{}
	SetRServicesEnabled(val interface{})
	RServicesEnabledInput() interface{}
	SqlConnectivityPort() *float64
	SetSqlConnectivityPort(val *float64)
	SqlConnectivityPortInput() *float64
	SqlConnectivityType() *string
	SetSqlConnectivityType(val *string)
	SqlConnectivityTypeInput() *string
	SqlConnectivityUpdatePassword() *string
	SetSqlConnectivityUpdatePassword(val *string)
	SqlConnectivityUpdatePasswordInput() *string
	SqlConnectivityUpdateUsername() *string
	SetSqlConnectivityUpdateUsername(val *string)
	SqlConnectivityUpdateUsernameInput() *string
	SqlInstance() MssqlVirtualMachineSqlInstanceOutputReference
	SqlInstanceInput() *MssqlVirtualMachineSqlInstance
	SqlLicenseType() *string
	SetSqlLicenseType(val *string)
	SqlLicenseTypeInput() *string
	SqlVirtualMachineGroupId() *string
	SetSqlVirtualMachineGroupId(val *string)
	SqlVirtualMachineGroupIdInput() *string
	StorageConfiguration() MssqlVirtualMachineStorageConfigurationOutputReference
	StorageConfigurationInput() *MssqlVirtualMachineStorageConfiguration
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MssqlVirtualMachineTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualMachineId() *string
	SetVirtualMachineId(val *string)
	VirtualMachineIdInput() *string
	WsfcDomainCredential() MssqlVirtualMachineWsfcDomainCredentialOutputReference
	WsfcDomainCredentialInput() *MssqlVirtualMachineWsfcDomainCredential
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
	PutAssessment(value *MssqlVirtualMachineAssessment)
	PutAutoBackup(value *MssqlVirtualMachineAutoBackup)
	PutAutoPatching(value *MssqlVirtualMachineAutoPatching)
	PutKeyVaultCredential(value *MssqlVirtualMachineKeyVaultCredential)
	PutSqlInstance(value *MssqlVirtualMachineSqlInstance)
	PutStorageConfiguration(value *MssqlVirtualMachineStorageConfiguration)
	PutTimeouts(value *MssqlVirtualMachineTimeouts)
	PutWsfcDomainCredential(value *MssqlVirtualMachineWsfcDomainCredential)
	ResetAssessment()
	ResetAutoBackup()
	ResetAutoPatching()
	ResetId()
	ResetKeyVaultCredential()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRServicesEnabled()
	ResetSqlConnectivityPort()
	ResetSqlConnectivityType()
	ResetSqlConnectivityUpdatePassword()
	ResetSqlConnectivityUpdateUsername()
	ResetSqlInstance()
	ResetSqlLicenseType()
	ResetSqlVirtualMachineGroupId()
	ResetStorageConfiguration()
	ResetTags()
	ResetTimeouts()
	ResetWsfcDomainCredential()
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

// The jsii proxy struct for MssqlVirtualMachine
type jsiiProxy_MssqlVirtualMachine struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MssqlVirtualMachine) Assessment() MssqlVirtualMachineAssessmentOutputReference {
	var returns MssqlVirtualMachineAssessmentOutputReference
	_jsii_.Get(
		j,
		"assessment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) AssessmentInput() *MssqlVirtualMachineAssessment {
	var returns *MssqlVirtualMachineAssessment
	_jsii_.Get(
		j,
		"assessmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) AutoBackup() MssqlVirtualMachineAutoBackupOutputReference {
	var returns MssqlVirtualMachineAutoBackupOutputReference
	_jsii_.Get(
		j,
		"autoBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) AutoBackupInput() *MssqlVirtualMachineAutoBackup {
	var returns *MssqlVirtualMachineAutoBackup
	_jsii_.Get(
		j,
		"autoBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) AutoPatching() MssqlVirtualMachineAutoPatchingOutputReference {
	var returns MssqlVirtualMachineAutoPatchingOutputReference
	_jsii_.Get(
		j,
		"autoPatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) AutoPatchingInput() *MssqlVirtualMachineAutoPatching {
	var returns *MssqlVirtualMachineAutoPatching
	_jsii_.Get(
		j,
		"autoPatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) KeyVaultCredential() MssqlVirtualMachineKeyVaultCredentialOutputReference {
	var returns MssqlVirtualMachineKeyVaultCredentialOutputReference
	_jsii_.Get(
		j,
		"keyVaultCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) KeyVaultCredentialInput() *MssqlVirtualMachineKeyVaultCredential {
	var returns *MssqlVirtualMachineKeyVaultCredential
	_jsii_.Get(
		j,
		"keyVaultCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) RServicesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rServicesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) RServicesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rServicesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlConnectivityPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqlConnectivityPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlConnectivityPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqlConnectivityPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlConnectivityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlConnectivityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlConnectivityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlConnectivityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlConnectivityUpdatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlConnectivityUpdatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlConnectivityUpdatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlConnectivityUpdatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlConnectivityUpdateUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlConnectivityUpdateUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlConnectivityUpdateUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlConnectivityUpdateUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlInstance() MssqlVirtualMachineSqlInstanceOutputReference {
	var returns MssqlVirtualMachineSqlInstanceOutputReference
	_jsii_.Get(
		j,
		"sqlInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlInstanceInput() *MssqlVirtualMachineSqlInstance {
	var returns *MssqlVirtualMachineSqlInstance
	_jsii_.Get(
		j,
		"sqlInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlLicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlLicenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlLicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlLicenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlVirtualMachineGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVirtualMachineGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) SqlVirtualMachineGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVirtualMachineGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) StorageConfiguration() MssqlVirtualMachineStorageConfigurationOutputReference {
	var returns MssqlVirtualMachineStorageConfigurationOutputReference
	_jsii_.Get(
		j,
		"storageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) StorageConfigurationInput() *MssqlVirtualMachineStorageConfiguration {
	var returns *MssqlVirtualMachineStorageConfiguration
	_jsii_.Get(
		j,
		"storageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) Timeouts() MssqlVirtualMachineTimeoutsOutputReference {
	var returns MssqlVirtualMachineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) VirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) VirtualMachineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) WsfcDomainCredential() MssqlVirtualMachineWsfcDomainCredentialOutputReference {
	var returns MssqlVirtualMachineWsfcDomainCredentialOutputReference
	_jsii_.Get(
		j,
		"wsfcDomainCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachine) WsfcDomainCredentialInput() *MssqlVirtualMachineWsfcDomainCredential {
	var returns *MssqlVirtualMachineWsfcDomainCredential
	_jsii_.Get(
		j,
		"wsfcDomainCredentialInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/mssql_virtual_machine azurerm_mssql_virtual_machine} Resource.
func NewMssqlVirtualMachine(scope constructs.Construct, id *string, config *MssqlVirtualMachineConfig) MssqlVirtualMachine {
	_init_.Initialize()

	if err := validateNewMssqlVirtualMachineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlVirtualMachine{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/mssql_virtual_machine azurerm_mssql_virtual_machine} Resource.
func NewMssqlVirtualMachine_Override(m MssqlVirtualMachine, scope constructs.Construct, id *string, config *MssqlVirtualMachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachine",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetRServicesEnabled(val interface{}) {
	if err := j.validateSetRServicesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rServicesEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetSqlConnectivityPort(val *float64) {
	if err := j.validateSetSqlConnectivityPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlConnectivityPort",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetSqlConnectivityType(val *string) {
	if err := j.validateSetSqlConnectivityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlConnectivityType",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetSqlConnectivityUpdatePassword(val *string) {
	if err := j.validateSetSqlConnectivityUpdatePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlConnectivityUpdatePassword",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetSqlConnectivityUpdateUsername(val *string) {
	if err := j.validateSetSqlConnectivityUpdateUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlConnectivityUpdateUsername",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetSqlLicenseType(val *string) {
	if err := j.validateSetSqlLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlLicenseType",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetSqlVirtualMachineGroupId(val *string) {
	if err := j.validateSetSqlVirtualMachineGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlVirtualMachineGroupId",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine)SetVirtualMachineId(val *string) {
	if err := j.validateSetVirtualMachineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualMachineId",
		val,
	)
}

// Generates CDKTF code for importing a MssqlVirtualMachine resource upon running "cdktf plan <stack-name>".
func MssqlVirtualMachine_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMssqlVirtualMachine_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachine",
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
func MssqlVirtualMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlVirtualMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MssqlVirtualMachine_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlVirtualMachine_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachine",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MssqlVirtualMachine_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlVirtualMachine_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachine",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MssqlVirtualMachine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachine",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MssqlVirtualMachine) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachine) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlVirtualMachine) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlVirtualMachine) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlVirtualMachine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachine) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlVirtualMachine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlVirtualMachine) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachine) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachine) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutAssessment(value *MssqlVirtualMachineAssessment) {
	if err := m.validatePutAssessmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAssessment",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutAutoBackup(value *MssqlVirtualMachineAutoBackup) {
	if err := m.validatePutAutoBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAutoBackup",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutAutoPatching(value *MssqlVirtualMachineAutoPatching) {
	if err := m.validatePutAutoPatchingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAutoPatching",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutKeyVaultCredential(value *MssqlVirtualMachineKeyVaultCredential) {
	if err := m.validatePutKeyVaultCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putKeyVaultCredential",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutSqlInstance(value *MssqlVirtualMachineSqlInstance) {
	if err := m.validatePutSqlInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSqlInstance",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutStorageConfiguration(value *MssqlVirtualMachineStorageConfiguration) {
	if err := m.validatePutStorageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStorageConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutTimeouts(value *MssqlVirtualMachineTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutWsfcDomainCredential(value *MssqlVirtualMachineWsfcDomainCredential) {
	if err := m.validatePutWsfcDomainCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWsfcDomainCredential",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetAssessment() {
	_jsii_.InvokeVoid(
		m,
		"resetAssessment",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetAutoBackup() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoBackup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetAutoPatching() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoPatching",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetKeyVaultCredential() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyVaultCredential",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetRServicesEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetRServicesEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetSqlConnectivityPort() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlConnectivityPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetSqlConnectivityType() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlConnectivityType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetSqlConnectivityUpdatePassword() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlConnectivityUpdatePassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetSqlConnectivityUpdateUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlConnectivityUpdateUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetSqlInstance() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlInstance",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetSqlLicenseType() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlLicenseType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetSqlVirtualMachineGroupId() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlVirtualMachineGroupId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetStorageConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) ResetWsfcDomainCredential() {
	_jsii_.InvokeVoid(
		m,
		"resetWsfcDomainCredential",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachine) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachine) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

