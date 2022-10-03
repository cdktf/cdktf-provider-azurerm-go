package mssqlvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/mssqlvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine azurerm_mssql_virtual_machine}.
type MssqlVirtualMachine interface {
	cdktf.TerraformResource
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	SqlLicenseType() *string
	SetSqlLicenseType(val *string)
	SqlLicenseTypeInput() *string
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
	PutAutoBackup(value *MssqlVirtualMachineAutoBackup)
	PutAutoPatching(value *MssqlVirtualMachineAutoPatching)
	PutKeyVaultCredential(value *MssqlVirtualMachineKeyVaultCredential)
	PutStorageConfiguration(value *MssqlVirtualMachineStorageConfiguration)
	PutTimeouts(value *MssqlVirtualMachineTimeouts)
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
	ResetStorageConfiguration()
	ResetTags()
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

// The jsii proxy struct for MssqlVirtualMachine
type jsiiProxy_MssqlVirtualMachine struct {
	internal.Type__cdktfTerraformResource
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

func (j *jsiiProxy_MssqlVirtualMachine) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine azurerm_mssql_virtual_machine} Resource.
func NewMssqlVirtualMachine(scope constructs.Construct, id *string, config *MssqlVirtualMachineConfig) MssqlVirtualMachine {
	_init_.Initialize()

	j := jsiiProxy_MssqlVirtualMachine{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine azurerm_mssql_virtual_machine} Resource.
func NewMssqlVirtualMachine_Override(m MssqlVirtualMachine, scope constructs.Construct, id *string, config *MssqlVirtualMachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachine",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetRServicesEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"rServicesEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetSqlConnectivityPort(val *float64) {
	_jsii_.Set(
		j,
		"sqlConnectivityPort",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetSqlConnectivityType(val *string) {
	_jsii_.Set(
		j,
		"sqlConnectivityType",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetSqlConnectivityUpdatePassword(val *string) {
	_jsii_.Set(
		j,
		"sqlConnectivityUpdatePassword",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetSqlConnectivityUpdateUsername(val *string) {
	_jsii_.Set(
		j,
		"sqlConnectivityUpdateUsername",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetSqlLicenseType(val *string) {
	_jsii_.Set(
		j,
		"sqlLicenseType",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachine) SetVirtualMachineId(val *string) {
	_jsii_.Set(
		j,
		"virtualMachineId",
		val,
	)
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachine",
		"isConstruct",
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

func (m *jsiiProxy_MssqlVirtualMachine) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachine) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutAutoBackup(value *MssqlVirtualMachineAutoBackup) {
	_jsii_.InvokeVoid(
		m,
		"putAutoBackup",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutAutoPatching(value *MssqlVirtualMachineAutoPatching) {
	_jsii_.InvokeVoid(
		m,
		"putAutoPatching",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutKeyVaultCredential(value *MssqlVirtualMachineKeyVaultCredential) {
	_jsii_.InvokeVoid(
		m,
		"putKeyVaultCredential",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutStorageConfiguration(value *MssqlVirtualMachineStorageConfiguration) {
	_jsii_.InvokeVoid(
		m,
		"putStorageConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachine) PutTimeouts(value *MssqlVirtualMachineTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
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

type MssqlVirtualMachineAutoBackup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#retention_period_in_days MssqlVirtualMachine#retention_period_in_days}.
	RetentionPeriodInDays *float64 `field:"required" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#storage_account_access_key MssqlVirtualMachine#storage_account_access_key}.
	StorageAccountAccessKey *string `field:"required" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#storage_blob_endpoint MssqlVirtualMachine#storage_blob_endpoint}.
	StorageBlobEndpoint *string `field:"required" json:"storageBlobEndpoint" yaml:"storageBlobEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#encryption_enabled MssqlVirtualMachine#encryption_enabled}.
	EncryptionEnabled interface{} `field:"optional" json:"encryptionEnabled" yaml:"encryptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#encryption_password MssqlVirtualMachine#encryption_password}.
	EncryptionPassword *string `field:"optional" json:"encryptionPassword" yaml:"encryptionPassword"`
	// manual_schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#manual_schedule MssqlVirtualMachine#manual_schedule}
	ManualSchedule *MssqlVirtualMachineAutoBackupManualSchedule `field:"optional" json:"manualSchedule" yaml:"manualSchedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#system_databases_backup_enabled MssqlVirtualMachine#system_databases_backup_enabled}.
	SystemDatabasesBackupEnabled interface{} `field:"optional" json:"systemDatabasesBackupEnabled" yaml:"systemDatabasesBackupEnabled"`
}

type MssqlVirtualMachineAutoBackupManualSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#full_backup_frequency MssqlVirtualMachine#full_backup_frequency}.
	FullBackupFrequency *string `field:"required" json:"fullBackupFrequency" yaml:"fullBackupFrequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#full_backup_start_hour MssqlVirtualMachine#full_backup_start_hour}.
	FullBackupStartHour *float64 `field:"required" json:"fullBackupStartHour" yaml:"fullBackupStartHour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#full_backup_window_in_hours MssqlVirtualMachine#full_backup_window_in_hours}.
	FullBackupWindowInHours *float64 `field:"required" json:"fullBackupWindowInHours" yaml:"fullBackupWindowInHours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#log_backup_frequency_in_minutes MssqlVirtualMachine#log_backup_frequency_in_minutes}.
	LogBackupFrequencyInMinutes *float64 `field:"required" json:"logBackupFrequencyInMinutes" yaml:"logBackupFrequencyInMinutes"`
}

type MssqlVirtualMachineAutoBackupManualScheduleOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	FullBackupFrequency() *string
	SetFullBackupFrequency(val *string)
	FullBackupFrequencyInput() *string
	FullBackupStartHour() *float64
	SetFullBackupStartHour(val *float64)
	FullBackupStartHourInput() *float64
	FullBackupWindowInHours() *float64
	SetFullBackupWindowInHours(val *float64)
	FullBackupWindowInHoursInput() *float64
	InternalValue() *MssqlVirtualMachineAutoBackupManualSchedule
	SetInternalValue(val *MssqlVirtualMachineAutoBackupManualSchedule)
	LogBackupFrequencyInMinutes() *float64
	SetLogBackupFrequencyInMinutes(val *float64)
	LogBackupFrequencyInMinutesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineAutoBackupManualScheduleOutputReference
type jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullBackupFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullBackupFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupStartHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullBackupStartHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupStartHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullBackupStartHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupWindowInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullBackupWindowInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupWindowInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullBackupWindowInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) InternalValue() *MssqlVirtualMachineAutoBackupManualSchedule {
	var returns *MssqlVirtualMachineAutoBackupManualSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) LogBackupFrequencyInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logBackupFrequencyInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) LogBackupFrequencyInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logBackupFrequencyInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineAutoBackupManualScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineAutoBackupManualScheduleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoBackupManualScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineAutoBackupManualScheduleOutputReference_Override(m MssqlVirtualMachineAutoBackupManualScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoBackupManualScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) SetFullBackupFrequency(val *string) {
	_jsii_.Set(
		j,
		"fullBackupFrequency",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) SetFullBackupStartHour(val *float64) {
	_jsii_.Set(
		j,
		"fullBackupStartHour",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) SetFullBackupWindowInHours(val *float64) {
	_jsii_.Set(
		j,
		"fullBackupWindowInHours",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) SetInternalValue(val *MssqlVirtualMachineAutoBackupManualSchedule) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) SetLogBackupFrequencyInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"logBackupFrequencyInMinutes",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MssqlVirtualMachineAutoBackupOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EncryptionEnabled() interface{}
	SetEncryptionEnabled(val interface{})
	EncryptionEnabledInput() interface{}
	EncryptionPassword() *string
	SetEncryptionPassword(val *string)
	EncryptionPasswordInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineAutoBackup
	SetInternalValue(val *MssqlVirtualMachineAutoBackup)
	ManualSchedule() MssqlVirtualMachineAutoBackupManualScheduleOutputReference
	ManualScheduleInput() *MssqlVirtualMachineAutoBackupManualSchedule
	RetentionPeriodInDays() *float64
	SetRetentionPeriodInDays(val *float64)
	RetentionPeriodInDaysInput() *float64
	StorageAccountAccessKey() *string
	SetStorageAccountAccessKey(val *string)
	StorageAccountAccessKeyInput() *string
	StorageBlobEndpoint() *string
	SetStorageBlobEndpoint(val *string)
	StorageBlobEndpointInput() *string
	SystemDatabasesBackupEnabled() interface{}
	SetSystemDatabasesBackupEnabled(val interface{})
	SystemDatabasesBackupEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutManualSchedule(value *MssqlVirtualMachineAutoBackupManualSchedule)
	ResetEncryptionEnabled()
	ResetEncryptionPassword()
	ResetManualSchedule()
	ResetSystemDatabasesBackupEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineAutoBackupOutputReference
type jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) EncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) EncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) EncryptionPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) EncryptionPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) InternalValue() *MssqlVirtualMachineAutoBackup {
	var returns *MssqlVirtualMachineAutoBackup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ManualSchedule() MssqlVirtualMachineAutoBackupManualScheduleOutputReference {
	var returns MssqlVirtualMachineAutoBackupManualScheduleOutputReference
	_jsii_.Get(
		j,
		"manualSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ManualScheduleInput() *MssqlVirtualMachineAutoBackupManualSchedule {
	var returns *MssqlVirtualMachineAutoBackupManualSchedule
	_jsii_.Get(
		j,
		"manualScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) RetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) RetentionPeriodInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) StorageAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) StorageBlobEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBlobEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) StorageBlobEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBlobEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SystemDatabasesBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemDatabasesBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SystemDatabasesBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemDatabasesBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineAutoBackupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineAutoBackupOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineAutoBackupOutputReference_Override(m MssqlVirtualMachineAutoBackupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SetEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SetEncryptionPassword(val *string) {
	_jsii_.Set(
		j,
		"encryptionPassword",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SetInternalValue(val *MssqlVirtualMachineAutoBackup) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SetRetentionPeriodInDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionPeriodInDays",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SetStorageAccountAccessKey(val *string) {
	_jsii_.Set(
		j,
		"storageAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SetStorageBlobEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageBlobEndpoint",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SetSystemDatabasesBackupEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"systemDatabasesBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) PutManualSchedule(value *MssqlVirtualMachineAutoBackupManualSchedule) {
	_jsii_.InvokeVoid(
		m,
		"putManualSchedule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ResetEncryptionEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ResetEncryptionPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ResetManualSchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetManualSchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ResetSystemDatabasesBackupEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetSystemDatabasesBackupEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MssqlVirtualMachineAutoPatching struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#day_of_week MssqlVirtualMachine#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#maintenance_window_duration_in_minutes MssqlVirtualMachine#maintenance_window_duration_in_minutes}.
	MaintenanceWindowDurationInMinutes *float64 `field:"required" json:"maintenanceWindowDurationInMinutes" yaml:"maintenanceWindowDurationInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#maintenance_window_starting_hour MssqlVirtualMachine#maintenance_window_starting_hour}.
	MaintenanceWindowStartingHour *float64 `field:"required" json:"maintenanceWindowStartingHour" yaml:"maintenanceWindowStartingHour"`
}

type MssqlVirtualMachineAutoPatchingOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DayOfWeek() *string
	SetDayOfWeek(val *string)
	DayOfWeekInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineAutoPatching
	SetInternalValue(val *MssqlVirtualMachineAutoPatching)
	MaintenanceWindowDurationInMinutes() *float64
	SetMaintenanceWindowDurationInMinutes(val *float64)
	MaintenanceWindowDurationInMinutesInput() *float64
	MaintenanceWindowStartingHour() *float64
	SetMaintenanceWindowStartingHour(val *float64)
	MaintenanceWindowStartingHourInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineAutoPatchingOutputReference
type jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) DayOfWeek() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) DayOfWeekInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) InternalValue() *MssqlVirtualMachineAutoPatching {
	var returns *MssqlVirtualMachineAutoPatching
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) MaintenanceWindowDurationInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceWindowDurationInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) MaintenanceWindowDurationInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceWindowDurationInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) MaintenanceWindowStartingHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceWindowStartingHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) MaintenanceWindowStartingHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceWindowStartingHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineAutoPatchingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineAutoPatchingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoPatchingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineAutoPatchingOutputReference_Override(m MssqlVirtualMachineAutoPatchingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoPatchingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) SetDayOfWeek(val *string) {
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) SetInternalValue(val *MssqlVirtualMachineAutoPatching) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) SetMaintenanceWindowDurationInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"maintenanceWindowDurationInMinutes",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) SetMaintenanceWindowStartingHour(val *float64) {
	_jsii_.Set(
		j,
		"maintenanceWindowStartingHour",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoPatchingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MssqlVirtualMachineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#sql_license_type MssqlVirtualMachine#sql_license_type}.
	SqlLicenseType *string `field:"required" json:"sqlLicenseType" yaml:"sqlLicenseType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#virtual_machine_id MssqlVirtualMachine#virtual_machine_id}.
	VirtualMachineId *string `field:"required" json:"virtualMachineId" yaml:"virtualMachineId"`
	// auto_backup block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#auto_backup MssqlVirtualMachine#auto_backup}
	AutoBackup *MssqlVirtualMachineAutoBackup `field:"optional" json:"autoBackup" yaml:"autoBackup"`
	// auto_patching block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#auto_patching MssqlVirtualMachine#auto_patching}
	AutoPatching *MssqlVirtualMachineAutoPatching `field:"optional" json:"autoPatching" yaml:"autoPatching"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#id MssqlVirtualMachine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// key_vault_credential block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#key_vault_credential MssqlVirtualMachine#key_vault_credential}
	KeyVaultCredential *MssqlVirtualMachineKeyVaultCredential `field:"optional" json:"keyVaultCredential" yaml:"keyVaultCredential"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#r_services_enabled MssqlVirtualMachine#r_services_enabled}.
	RServicesEnabled interface{} `field:"optional" json:"rServicesEnabled" yaml:"rServicesEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#sql_connectivity_port MssqlVirtualMachine#sql_connectivity_port}.
	SqlConnectivityPort *float64 `field:"optional" json:"sqlConnectivityPort" yaml:"sqlConnectivityPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#sql_connectivity_type MssqlVirtualMachine#sql_connectivity_type}.
	SqlConnectivityType *string `field:"optional" json:"sqlConnectivityType" yaml:"sqlConnectivityType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#sql_connectivity_update_password MssqlVirtualMachine#sql_connectivity_update_password}.
	SqlConnectivityUpdatePassword *string `field:"optional" json:"sqlConnectivityUpdatePassword" yaml:"sqlConnectivityUpdatePassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#sql_connectivity_update_username MssqlVirtualMachine#sql_connectivity_update_username}.
	SqlConnectivityUpdateUsername *string `field:"optional" json:"sqlConnectivityUpdateUsername" yaml:"sqlConnectivityUpdateUsername"`
	// storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#storage_configuration MssqlVirtualMachine#storage_configuration}
	StorageConfiguration *MssqlVirtualMachineStorageConfiguration `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#tags MssqlVirtualMachine#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#timeouts MssqlVirtualMachine#timeouts}
	Timeouts *MssqlVirtualMachineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type MssqlVirtualMachineKeyVaultCredential struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#key_vault_url MssqlVirtualMachine#key_vault_url}.
	KeyVaultUrl *string `field:"required" json:"keyVaultUrl" yaml:"keyVaultUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#name MssqlVirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#service_principal_name MssqlVirtualMachine#service_principal_name}.
	ServicePrincipalName *string `field:"required" json:"servicePrincipalName" yaml:"servicePrincipalName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#service_principal_secret MssqlVirtualMachine#service_principal_secret}.
	ServicePrincipalSecret *string `field:"required" json:"servicePrincipalSecret" yaml:"servicePrincipalSecret"`
}

type MssqlVirtualMachineKeyVaultCredentialOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineKeyVaultCredential
	SetInternalValue(val *MssqlVirtualMachineKeyVaultCredential)
	KeyVaultUrl() *string
	SetKeyVaultUrl(val *string)
	KeyVaultUrlInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	ServicePrincipalName() *string
	SetServicePrincipalName(val *string)
	ServicePrincipalNameInput() *string
	ServicePrincipalSecret() *string
	SetServicePrincipalSecret(val *string)
	ServicePrincipalSecretInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineKeyVaultCredentialOutputReference
type jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) InternalValue() *MssqlVirtualMachineKeyVaultCredential {
	var returns *MssqlVirtualMachineKeyVaultCredential
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) KeyVaultUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) KeyVaultUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) ServicePrincipalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) ServicePrincipalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) ServicePrincipalSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) ServicePrincipalSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineKeyVaultCredentialOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineKeyVaultCredentialOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineKeyVaultCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineKeyVaultCredentialOutputReference_Override(m MssqlVirtualMachineKeyVaultCredentialOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineKeyVaultCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) SetInternalValue(val *MssqlVirtualMachineKeyVaultCredential) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) SetKeyVaultUrl(val *string) {
	_jsii_.Set(
		j,
		"keyVaultUrl",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) SetServicePrincipalName(val *string) {
	_jsii_.Set(
		j,
		"servicePrincipalName",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) SetServicePrincipalSecret(val *string) {
	_jsii_.Set(
		j,
		"servicePrincipalSecret",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineKeyVaultCredentialOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MssqlVirtualMachineStorageConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#disk_type MssqlVirtualMachine#disk_type}.
	DiskType *string `field:"required" json:"diskType" yaml:"diskType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#storage_workload_type MssqlVirtualMachine#storage_workload_type}.
	StorageWorkloadType *string `field:"required" json:"storageWorkloadType" yaml:"storageWorkloadType"`
	// data_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#data_settings MssqlVirtualMachine#data_settings}
	DataSettings *MssqlVirtualMachineStorageConfigurationDataSettings `field:"optional" json:"dataSettings" yaml:"dataSettings"`
	// log_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#log_settings MssqlVirtualMachine#log_settings}
	LogSettings *MssqlVirtualMachineStorageConfigurationLogSettings `field:"optional" json:"logSettings" yaml:"logSettings"`
	// temp_db_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#temp_db_settings MssqlVirtualMachine#temp_db_settings}
	TempDbSettings *MssqlVirtualMachineStorageConfigurationTempDbSettings `field:"optional" json:"tempDbSettings" yaml:"tempDbSettings"`
}

type MssqlVirtualMachineStorageConfigurationDataSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#default_file_path MssqlVirtualMachine#default_file_path}.
	DefaultFilePath *string `field:"required" json:"defaultFilePath" yaml:"defaultFilePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#luns MssqlVirtualMachine#luns}.
	Luns *[]*float64 `field:"required" json:"luns" yaml:"luns"`
}

type MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultFilePath() *string
	SetDefaultFilePath(val *string)
	DefaultFilePathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineStorageConfigurationDataSettings
	SetInternalValue(val *MssqlVirtualMachineStorageConfigurationDataSettings)
	Luns() *[]*float64
	SetLuns(val *[]*float64)
	LunsInput() *[]*float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference
type jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) DefaultFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) DefaultFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) InternalValue() *MssqlVirtualMachineStorageConfigurationDataSettings {
	var returns *MssqlVirtualMachineStorageConfigurationDataSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) Luns() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"luns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) LunsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"lunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineStorageConfigurationDataSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineStorageConfigurationDataSettingsOutputReference_Override(m MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) SetDefaultFilePath(val *string) {
	_jsii_.Set(
		j,
		"defaultFilePath",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) SetInternalValue(val *MssqlVirtualMachineStorageConfigurationDataSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) SetLuns(val *[]*float64) {
	_jsii_.Set(
		j,
		"luns",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MssqlVirtualMachineStorageConfigurationLogSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#default_file_path MssqlVirtualMachine#default_file_path}.
	DefaultFilePath *string `field:"required" json:"defaultFilePath" yaml:"defaultFilePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#luns MssqlVirtualMachine#luns}.
	Luns *[]*float64 `field:"required" json:"luns" yaml:"luns"`
}

type MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultFilePath() *string
	SetDefaultFilePath(val *string)
	DefaultFilePathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineStorageConfigurationLogSettings
	SetInternalValue(val *MssqlVirtualMachineStorageConfigurationLogSettings)
	Luns() *[]*float64
	SetLuns(val *[]*float64)
	LunsInput() *[]*float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference
type jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) DefaultFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) DefaultFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) InternalValue() *MssqlVirtualMachineStorageConfigurationLogSettings {
	var returns *MssqlVirtualMachineStorageConfigurationLogSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) Luns() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"luns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) LunsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"lunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineStorageConfigurationLogSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineStorageConfigurationLogSettingsOutputReference_Override(m MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) SetDefaultFilePath(val *string) {
	_jsii_.Set(
		j,
		"defaultFilePath",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) SetInternalValue(val *MssqlVirtualMachineStorageConfigurationLogSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) SetLuns(val *[]*float64) {
	_jsii_.Set(
		j,
		"luns",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MssqlVirtualMachineStorageConfigurationOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSettings() MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference
	DataSettingsInput() *MssqlVirtualMachineStorageConfigurationDataSettings
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineStorageConfiguration
	SetInternalValue(val *MssqlVirtualMachineStorageConfiguration)
	LogSettings() MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference
	LogSettingsInput() *MssqlVirtualMachineStorageConfigurationLogSettings
	StorageWorkloadType() *string
	SetStorageWorkloadType(val *string)
	StorageWorkloadTypeInput() *string
	TempDbSettings() MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference
	TempDbSettingsInput() *MssqlVirtualMachineStorageConfigurationTempDbSettings
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutDataSettings(value *MssqlVirtualMachineStorageConfigurationDataSettings)
	PutLogSettings(value *MssqlVirtualMachineStorageConfigurationLogSettings)
	PutTempDbSettings(value *MssqlVirtualMachineStorageConfigurationTempDbSettings)
	ResetDataSettings()
	ResetLogSettings()
	ResetTempDbSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineStorageConfigurationOutputReference
type jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) DataSettings() MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference {
	var returns MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference
	_jsii_.Get(
		j,
		"dataSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) DataSettingsInput() *MssqlVirtualMachineStorageConfigurationDataSettings {
	var returns *MssqlVirtualMachineStorageConfigurationDataSettings
	_jsii_.Get(
		j,
		"dataSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) InternalValue() *MssqlVirtualMachineStorageConfiguration {
	var returns *MssqlVirtualMachineStorageConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) LogSettings() MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference {
	var returns MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference
	_jsii_.Get(
		j,
		"logSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) LogSettingsInput() *MssqlVirtualMachineStorageConfigurationLogSettings {
	var returns *MssqlVirtualMachineStorageConfigurationLogSettings
	_jsii_.Get(
		j,
		"logSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) StorageWorkloadType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageWorkloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) StorageWorkloadTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageWorkloadTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) TempDbSettings() MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference {
	var returns MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference
	_jsii_.Get(
		j,
		"tempDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) TempDbSettingsInput() *MssqlVirtualMachineStorageConfigurationTempDbSettings {
	var returns *MssqlVirtualMachineStorageConfigurationTempDbSettings
	_jsii_.Get(
		j,
		"tempDbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineStorageConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineStorageConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineStorageConfigurationOutputReference_Override(m MssqlVirtualMachineStorageConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) SetDiskType(val *string) {
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) SetInternalValue(val *MssqlVirtualMachineStorageConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) SetStorageWorkloadType(val *string) {
	_jsii_.Set(
		j,
		"storageWorkloadType",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) PutDataSettings(value *MssqlVirtualMachineStorageConfigurationDataSettings) {
	_jsii_.InvokeVoid(
		m,
		"putDataSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) PutLogSettings(value *MssqlVirtualMachineStorageConfigurationLogSettings) {
	_jsii_.InvokeVoid(
		m,
		"putLogSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) PutTempDbSettings(value *MssqlVirtualMachineStorageConfigurationTempDbSettings) {
	_jsii_.InvokeVoid(
		m,
		"putTempDbSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ResetDataSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetDataSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ResetLogSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetLogSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ResetTempDbSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetTempDbSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MssqlVirtualMachineStorageConfigurationTempDbSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#default_file_path MssqlVirtualMachine#default_file_path}.
	DefaultFilePath *string `field:"required" json:"defaultFilePath" yaml:"defaultFilePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#luns MssqlVirtualMachine#luns}.
	Luns *[]*float64 `field:"required" json:"luns" yaml:"luns"`
}

type MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultFilePath() *string
	SetDefaultFilePath(val *string)
	DefaultFilePathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineStorageConfigurationTempDbSettings
	SetInternalValue(val *MssqlVirtualMachineStorageConfigurationTempDbSettings)
	Luns() *[]*float64
	SetLuns(val *[]*float64)
	LunsInput() *[]*float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference
type jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) DefaultFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) DefaultFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) InternalValue() *MssqlVirtualMachineStorageConfigurationTempDbSettings {
	var returns *MssqlVirtualMachineStorageConfigurationTempDbSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) Luns() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"luns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) LunsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"lunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference_Override(m MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) SetDefaultFilePath(val *string) {
	_jsii_.Set(
		j,
		"defaultFilePath",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) SetInternalValue(val *MssqlVirtualMachineStorageConfigurationTempDbSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) SetLuns(val *[]*float64) {
	_jsii_.Set(
		j,
		"luns",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MssqlVirtualMachineTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#create MssqlVirtualMachine#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#delete MssqlVirtualMachine#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#read MssqlVirtualMachine#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#update MssqlVirtualMachine#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type MssqlVirtualMachineTimeoutsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineTimeoutsOutputReference
type jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineTimeoutsOutputReference_Override(m MssqlVirtualMachineTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		m,
		"resetRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

