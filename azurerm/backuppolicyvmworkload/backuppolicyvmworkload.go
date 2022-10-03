package backuppolicyvmworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/backuppolicyvmworkload/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload azurerm_backup_policy_vm_workload}.
type BackupPolicyVmWorkload interface {
	cdktf.TerraformResource
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProtectionPolicy() BackupPolicyVmWorkloadProtectionPolicyList
	ProtectionPolicyInput() interface{}
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
	RecoveryVaultName() *string
	SetRecoveryVaultName(val *string)
	RecoveryVaultNameInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Settings() BackupPolicyVmWorkloadSettingsOutputReference
	SettingsInput() *BackupPolicyVmWorkloadSettings
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BackupPolicyVmWorkloadTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkloadType() *string
	SetWorkloadType(val *string)
	WorkloadTypeInput() *string
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
	PutProtectionPolicy(value interface{})
	PutSettings(value *BackupPolicyVmWorkloadSettings)
	PutTimeouts(value *BackupPolicyVmWorkloadTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for BackupPolicyVmWorkload
type jsiiProxy_BackupPolicyVmWorkload struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BackupPolicyVmWorkload) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) ProtectionPolicy() BackupPolicyVmWorkloadProtectionPolicyList {
	var returns BackupPolicyVmWorkloadProtectionPolicyList
	_jsii_.Get(
		j,
		"protectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) ProtectionPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) RecoveryVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) RecoveryVaultNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) Settings() BackupPolicyVmWorkloadSettingsOutputReference {
	var returns BackupPolicyVmWorkloadSettingsOutputReference
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SettingsInput() *BackupPolicyVmWorkloadSettings {
	var returns *BackupPolicyVmWorkloadSettings
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) Timeouts() BackupPolicyVmWorkloadTimeoutsOutputReference {
	var returns BackupPolicyVmWorkloadTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) WorkloadType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkload) WorkloadTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload azurerm_backup_policy_vm_workload} Resource.
func NewBackupPolicyVmWorkload(scope constructs.Construct, id *string, config *BackupPolicyVmWorkloadConfig) BackupPolicyVmWorkload {
	_init_.Initialize()

	j := jsiiProxy_BackupPolicyVmWorkload{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkload",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload azurerm_backup_policy_vm_workload} Resource.
func NewBackupPolicyVmWorkload_Override(b BackupPolicyVmWorkload, scope constructs.Construct, id *string, config *BackupPolicyVmWorkloadConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkload",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetRecoveryVaultName(val *string) {
	_jsii_.Set(
		j,
		"recoveryVaultName",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkload) SetWorkloadType(val *string) {
	_jsii_.Set(
		j,
		"workloadType",
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
func BackupPolicyVmWorkload_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkload",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BackupPolicyVmWorkload_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkload",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkload) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkload) PutProtectionPolicy(value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"putProtectionPolicy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkload) PutSettings(value *BackupPolicyVmWorkloadSettings) {
	_jsii_.InvokeVoid(
		b,
		"putSettings",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkload) PutTimeouts(value *BackupPolicyVmWorkloadTimeouts) {
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkload) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkload) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkload) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkload) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkload) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BackupPolicyVmWorkloadConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#name BackupPolicyVmWorkload#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// protection_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#protection_policy BackupPolicyVmWorkload#protection_policy}
	ProtectionPolicy interface{} `field:"required" json:"protectionPolicy" yaml:"protectionPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#recovery_vault_name BackupPolicyVmWorkload#recovery_vault_name}.
	RecoveryVaultName *string `field:"required" json:"recoveryVaultName" yaml:"recoveryVaultName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#resource_group_name BackupPolicyVmWorkload#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#settings BackupPolicyVmWorkload#settings}
	Settings *BackupPolicyVmWorkloadSettings `field:"required" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#workload_type BackupPolicyVmWorkload#workload_type}.
	WorkloadType *string `field:"required" json:"workloadType" yaml:"workloadType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#id BackupPolicyVmWorkload#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#timeouts BackupPolicyVmWorkload#timeouts}
	Timeouts *BackupPolicyVmWorkloadTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type BackupPolicyVmWorkloadProtectionPolicy struct {
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#backup BackupPolicyVmWorkload#backup}
	Backup *BackupPolicyVmWorkloadProtectionPolicyBackup `field:"required" json:"backup" yaml:"backup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#policy_type BackupPolicyVmWorkload#policy_type}.
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// retention_daily block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#retention_daily BackupPolicyVmWorkload#retention_daily}
	RetentionDaily *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily `field:"optional" json:"retentionDaily" yaml:"retentionDaily"`
	// retention_monthly block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#retention_monthly BackupPolicyVmWorkload#retention_monthly}
	RetentionMonthly *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly `field:"optional" json:"retentionMonthly" yaml:"retentionMonthly"`
	// retention_weekly block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#retention_weekly BackupPolicyVmWorkload#retention_weekly}
	RetentionWeekly *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly `field:"optional" json:"retentionWeekly" yaml:"retentionWeekly"`
	// retention_yearly block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#retention_yearly BackupPolicyVmWorkload#retention_yearly}
	RetentionYearly *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly `field:"optional" json:"retentionYearly" yaml:"retentionYearly"`
	// simple_retention block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#simple_retention BackupPolicyVmWorkload#simple_retention}
	SimpleRetention *BackupPolicyVmWorkloadProtectionPolicySimpleRetention `field:"optional" json:"simpleRetention" yaml:"simpleRetention"`
}

type BackupPolicyVmWorkloadProtectionPolicyBackup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#frequency BackupPolicyVmWorkload#frequency}.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#frequency_in_minutes BackupPolicyVmWorkload#frequency_in_minutes}.
	FrequencyInMinutes *float64 `field:"optional" json:"frequencyInMinutes" yaml:"frequencyInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#time BackupPolicyVmWorkload#time}.
	Time *string `field:"optional" json:"time" yaml:"time"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#weekdays BackupPolicyVmWorkload#weekdays}.
	Weekdays *[]*string `field:"optional" json:"weekdays" yaml:"weekdays"`
}

type BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference interface {
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
	Frequency() *string
	SetFrequency(val *string)
	FrequencyInMinutes() *float64
	SetFrequencyInMinutes(val *float64)
	FrequencyInMinutesInput() *float64
	FrequencyInput() *string
	InternalValue() *BackupPolicyVmWorkloadProtectionPolicyBackup
	SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicyBackup)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Time() *string
	SetTime(val *string)
	TimeInput() *string
	Weekdays() *[]*string
	SetWeekdays(val *[]*string)
	WeekdaysInput() *[]*string
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
	ResetFrequency()
	ResetFrequencyInMinutes()
	ResetTime()
	ResetWeekdays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference
type jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) FrequencyInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) FrequencyInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) InternalValue() *BackupPolicyVmWorkloadProtectionPolicyBackup {
	var returns *BackupPolicyVmWorkloadProtectionPolicyBackup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) Time() *string {
	var returns *string
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) TimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) Weekdays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) WeekdaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdaysInput",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmWorkloadProtectionPolicyBackupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupPolicyVmWorkloadProtectionPolicyBackupOutputReference_Override(b BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) SetFrequency(val *string) {
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) SetFrequencyInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"frequencyInMinutes",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicyBackup) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) SetTime(val *string) {
	_jsii_.Set(
		j,
		"time",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) SetWeekdays(val *[]*string) {
	_jsii_.Set(
		j,
		"weekdays",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) ResetFrequency() {
	_jsii_.InvokeVoid(
		b,
		"resetFrequency",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) ResetFrequencyInMinutes() {
	_jsii_.InvokeVoid(
		b,
		"resetFrequencyInMinutes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) ResetTime() {
	_jsii_.InvokeVoid(
		b,
		"resetTime",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) ResetWeekdays() {
	_jsii_.InvokeVoid(
		b,
		"resetWeekdays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BackupPolicyVmWorkloadProtectionPolicyList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BackupPolicyVmWorkloadProtectionPolicyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupPolicyVmWorkloadProtectionPolicyList
type jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmWorkloadProtectionPolicyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BackupPolicyVmWorkloadProtectionPolicyList {
	_init_.Initialize()

	j := jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBackupPolicyVmWorkloadProtectionPolicyList_Override(b BackupPolicyVmWorkloadProtectionPolicyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) Get(index *float64) BackupPolicyVmWorkloadProtectionPolicyOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicyOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BackupPolicyVmWorkloadProtectionPolicyOutputReference interface {
	cdktf.ComplexObject
	Backup() BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference
	BackupInput() *BackupPolicyVmWorkloadProtectionPolicyBackup
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PolicyType() *string
	SetPolicyType(val *string)
	PolicyTypeInput() *string
	RetentionDaily() BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference
	RetentionDailyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily
	RetentionMonthly() BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference
	RetentionMonthlyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly
	RetentionWeekly() BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference
	RetentionWeeklyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly
	RetentionYearly() BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference
	RetentionYearlyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly
	SimpleRetention() BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference
	SimpleRetentionInput() *BackupPolicyVmWorkloadProtectionPolicySimpleRetention
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
	PutBackup(value *BackupPolicyVmWorkloadProtectionPolicyBackup)
	PutRetentionDaily(value *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily)
	PutRetentionMonthly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly)
	PutRetentionWeekly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly)
	PutRetentionYearly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly)
	PutSimpleRetention(value *BackupPolicyVmWorkloadProtectionPolicySimpleRetention)
	ResetRetentionDaily()
	ResetRetentionMonthly()
	ResetRetentionWeekly()
	ResetRetentionYearly()
	ResetSimpleRetention()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupPolicyVmWorkloadProtectionPolicyOutputReference
type jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) Backup() BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) BackupInput() *BackupPolicyVmWorkloadProtectionPolicyBackup {
	var returns *BackupPolicyVmWorkloadProtectionPolicyBackup
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionDaily() BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference
	_jsii_.Get(
		j,
		"retentionDaily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionDailyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily
	_jsii_.Get(
		j,
		"retentionDailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionMonthly() BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference
	_jsii_.Get(
		j,
		"retentionMonthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionMonthlyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly
	_jsii_.Get(
		j,
		"retentionMonthlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionWeekly() BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference
	_jsii_.Get(
		j,
		"retentionWeekly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionWeeklyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly
	_jsii_.Get(
		j,
		"retentionWeeklyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionYearly() BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference
	_jsii_.Get(
		j,
		"retentionYearly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionYearlyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly
	_jsii_.Get(
		j,
		"retentionYearlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) SimpleRetention() BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference
	_jsii_.Get(
		j,
		"simpleRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) SimpleRetentionInput() *BackupPolicyVmWorkloadProtectionPolicySimpleRetention {
	var returns *BackupPolicyVmWorkloadProtectionPolicySimpleRetention
	_jsii_.Get(
		j,
		"simpleRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmWorkloadProtectionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BackupPolicyVmWorkloadProtectionPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBackupPolicyVmWorkloadProtectionPolicyOutputReference_Override(b BackupPolicyVmWorkloadProtectionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) SetPolicyType(val *string) {
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutBackup(value *BackupPolicyVmWorkloadProtectionPolicyBackup) {
	_jsii_.InvokeVoid(
		b,
		"putBackup",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutRetentionDaily(value *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily) {
	_jsii_.InvokeVoid(
		b,
		"putRetentionDaily",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutRetentionMonthly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly) {
	_jsii_.InvokeVoid(
		b,
		"putRetentionMonthly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutRetentionWeekly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly) {
	_jsii_.InvokeVoid(
		b,
		"putRetentionWeekly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutRetentionYearly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly) {
	_jsii_.InvokeVoid(
		b,
		"putRetentionYearly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutSimpleRetention(value *BackupPolicyVmWorkloadProtectionPolicySimpleRetention) {
	_jsii_.InvokeVoid(
		b,
		"putSimpleRetention",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ResetRetentionDaily() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionDaily",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ResetRetentionMonthly() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionMonthly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ResetRetentionWeekly() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionWeekly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ResetRetentionYearly() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionYearly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ResetSimpleRetention() {
	_jsii_.InvokeVoid(
		b,
		"resetSimpleRetention",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BackupPolicyVmWorkloadProtectionPolicyRetentionDaily struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#count BackupPolicyVmWorkload#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
}

type BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference interface {
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily
	SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily)
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

// The jsii proxy struct for BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference
type jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) InternalValue() *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference_Override(b BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#count BackupPolicyVmWorkload#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#format_type BackupPolicyVmWorkload#format_type}.
	FormatType *string `field:"required" json:"formatType" yaml:"formatType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#monthdays BackupPolicyVmWorkload#monthdays}.
	Monthdays *[]*float64 `field:"optional" json:"monthdays" yaml:"monthdays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#weekdays BackupPolicyVmWorkload#weekdays}.
	Weekdays *[]*string `field:"optional" json:"weekdays" yaml:"weekdays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#weeks BackupPolicyVmWorkload#weeks}.
	Weeks *[]*string `field:"optional" json:"weeks" yaml:"weeks"`
}

type BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference interface {
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FormatType() *string
	SetFormatType(val *string)
	FormatTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly
	SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly)
	Monthdays() *[]*float64
	SetMonthdays(val *[]*float64)
	MonthdaysInput() *[]*float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Weekdays() *[]*string
	SetWeekdays(val *[]*string)
	WeekdaysInput() *[]*string
	Weeks() *[]*string
	SetWeeks(val *[]*string)
	WeeksInput() *[]*string
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
	ResetMonthdays()
	ResetWeekdays()
	ResetWeeks()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference
type jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) FormatType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) FormatTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) InternalValue() *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) Monthdays() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"monthdays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) MonthdaysInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"monthdaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) Weekdays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) WeekdaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) Weeks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) WeeksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksInput",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference_Override(b BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) SetFormatType(val *string) {
	_jsii_.Set(
		j,
		"formatType",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) SetMonthdays(val *[]*float64) {
	_jsii_.Set(
		j,
		"monthdays",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) SetWeekdays(val *[]*string) {
	_jsii_.Set(
		j,
		"weekdays",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) SetWeeks(val *[]*string) {
	_jsii_.Set(
		j,
		"weeks",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) ResetMonthdays() {
	_jsii_.InvokeVoid(
		b,
		"resetMonthdays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) ResetWeekdays() {
	_jsii_.InvokeVoid(
		b,
		"resetWeekdays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) ResetWeeks() {
	_jsii_.InvokeVoid(
		b,
		"resetWeeks",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#count BackupPolicyVmWorkload#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#weekdays BackupPolicyVmWorkload#weekdays}.
	Weekdays *[]*string `field:"required" json:"weekdays" yaml:"weekdays"`
}

type BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference interface {
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly
	SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Weekdays() *[]*string
	SetWeekdays(val *[]*string)
	WeekdaysInput() *[]*string
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

// The jsii proxy struct for BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference
type jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) InternalValue() *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) Weekdays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) WeekdaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdaysInput",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference_Override(b BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) SetWeekdays(val *[]*string) {
	_jsii_.Set(
		j,
		"weekdays",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BackupPolicyVmWorkloadProtectionPolicyRetentionYearly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#count BackupPolicyVmWorkload#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#format_type BackupPolicyVmWorkload#format_type}.
	FormatType *string `field:"required" json:"formatType" yaml:"formatType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#months BackupPolicyVmWorkload#months}.
	Months *[]*string `field:"required" json:"months" yaml:"months"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#monthdays BackupPolicyVmWorkload#monthdays}.
	Monthdays *[]*float64 `field:"optional" json:"monthdays" yaml:"monthdays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#weekdays BackupPolicyVmWorkload#weekdays}.
	Weekdays *[]*string `field:"optional" json:"weekdays" yaml:"weekdays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#weeks BackupPolicyVmWorkload#weeks}.
	Weeks *[]*string `field:"optional" json:"weeks" yaml:"weeks"`
}

type BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference interface {
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FormatType() *string
	SetFormatType(val *string)
	FormatTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly
	SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly)
	Monthdays() *[]*float64
	SetMonthdays(val *[]*float64)
	MonthdaysInput() *[]*float64
	Months() *[]*string
	SetMonths(val *[]*string)
	MonthsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Weekdays() *[]*string
	SetWeekdays(val *[]*string)
	WeekdaysInput() *[]*string
	Weeks() *[]*string
	SetWeeks(val *[]*string)
	WeeksInput() *[]*string
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
	ResetMonthdays()
	ResetWeekdays()
	ResetWeeks()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference
type jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) FormatType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) FormatTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) InternalValue() *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) Monthdays() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"monthdays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) MonthdaysInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"monthdaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) Months() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) MonthsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) Weekdays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) WeekdaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) Weeks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) WeeksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksInput",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference_Override(b BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) SetFormatType(val *string) {
	_jsii_.Set(
		j,
		"formatType",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) SetMonthdays(val *[]*float64) {
	_jsii_.Set(
		j,
		"monthdays",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) SetMonths(val *[]*string) {
	_jsii_.Set(
		j,
		"months",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) SetWeekdays(val *[]*string) {
	_jsii_.Set(
		j,
		"weekdays",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) SetWeeks(val *[]*string) {
	_jsii_.Set(
		j,
		"weeks",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) ResetMonthdays() {
	_jsii_.InvokeVoid(
		b,
		"resetMonthdays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) ResetWeekdays() {
	_jsii_.InvokeVoid(
		b,
		"resetWeekdays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) ResetWeeks() {
	_jsii_.InvokeVoid(
		b,
		"resetWeeks",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BackupPolicyVmWorkloadProtectionPolicySimpleRetention struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#count BackupPolicyVmWorkload#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
}

type BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference interface {
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *BackupPolicyVmWorkloadProtectionPolicySimpleRetention
	SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicySimpleRetention)
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

// The jsii proxy struct for BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference
type jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) InternalValue() *BackupPolicyVmWorkloadProtectionPolicySimpleRetention {
	var returns *BackupPolicyVmWorkloadProtectionPolicySimpleRetention
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference_Override(b BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) SetInternalValue(val *BackupPolicyVmWorkloadProtectionPolicySimpleRetention) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BackupPolicyVmWorkloadSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#time_zone BackupPolicyVmWorkload#time_zone}.
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#compression_enabled BackupPolicyVmWorkload#compression_enabled}.
	CompressionEnabled interface{} `field:"optional" json:"compressionEnabled" yaml:"compressionEnabled"`
}

type BackupPolicyVmWorkloadSettingsOutputReference interface {
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
	CompressionEnabled() interface{}
	SetCompressionEnabled(val interface{})
	CompressionEnabledInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *BackupPolicyVmWorkloadSettings
	SetInternalValue(val *BackupPolicyVmWorkloadSettings)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	ResetCompressionEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupPolicyVmWorkloadSettingsOutputReference
type jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) CompressionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) CompressionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) InternalValue() *BackupPolicyVmWorkloadSettings {
	var returns *BackupPolicyVmWorkloadSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmWorkloadSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupPolicyVmWorkloadSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupPolicyVmWorkloadSettingsOutputReference_Override(b BackupPolicyVmWorkloadSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) SetCompressionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"compressionEnabled",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) SetInternalValue(val *BackupPolicyVmWorkloadSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) SetTimeZone(val *string) {
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) ResetCompressionEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetCompressionEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BackupPolicyVmWorkloadTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#create BackupPolicyVmWorkload#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#delete BackupPolicyVmWorkload#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#read BackupPolicyVmWorkload#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm_workload#update BackupPolicyVmWorkload#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type BackupPolicyVmWorkloadTimeoutsOutputReference interface {
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

// The jsii proxy struct for BackupPolicyVmWorkloadTimeoutsOutputReference
type jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmWorkloadTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupPolicyVmWorkloadTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupPolicyVmWorkloadTimeoutsOutputReference_Override(b BackupPolicyVmWorkloadTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		b,
		"resetCreate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		b,
		"resetDelete",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		b,
		"resetRead",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		b,
		"resetUpdate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

