package dataprotectionbackuppolicypostgresql

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/dataprotectionbackuppolicypostgresql/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql azurerm_data_protection_backup_policy_postgresql}.
type DataProtectionBackupPolicyPostgresql interface {
	cdktf.TerraformResource
	BackupRepeatingTimeIntervals() *[]*string
	SetBackupRepeatingTimeIntervals(val *[]*string)
	BackupRepeatingTimeIntervalsInput() *[]*string
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
	DefaultRetentionDuration() *string
	SetDefaultRetentionDuration(val *string)
	DefaultRetentionDurationInput() *string
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
	RetentionRule() DataProtectionBackupPolicyPostgresqlRetentionRuleList
	RetentionRuleInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference
	TimeoutsInput() interface{}
	VaultName() *string
	SetVaultName(val *string)
	VaultNameInput() *string
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
	PutRetentionRule(value interface{})
	PutTimeouts(value *DataProtectionBackupPolicyPostgresqlTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetentionRule()
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

// The jsii proxy struct for DataProtectionBackupPolicyPostgresql
type jsiiProxy_DataProtectionBackupPolicyPostgresql struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) BackupRepeatingTimeIntervals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backupRepeatingTimeIntervals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) BackupRepeatingTimeIntervalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backupRepeatingTimeIntervalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) DefaultRetentionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRetentionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) DefaultRetentionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRetentionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) RetentionRule() DataProtectionBackupPolicyPostgresqlRetentionRuleList {
	var returns DataProtectionBackupPolicyPostgresqlRetentionRuleList
	_jsii_.Get(
		j,
		"retentionRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) RetentionRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retentionRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) Timeouts() DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference {
	var returns DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) VaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) VaultNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql azurerm_data_protection_backup_policy_postgresql} Resource.
func NewDataProtectionBackupPolicyPostgresql(scope constructs.Construct, id *string, config *DataProtectionBackupPolicyPostgresqlConfig) DataProtectionBackupPolicyPostgresql {
	_init_.Initialize()

	j := jsiiProxy_DataProtectionBackupPolicyPostgresql{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresql",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql azurerm_data_protection_backup_policy_postgresql} Resource.
func NewDataProtectionBackupPolicyPostgresql_Override(d DataProtectionBackupPolicyPostgresql, scope constructs.Construct, id *string, config *DataProtectionBackupPolicyPostgresqlConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresql",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetBackupRepeatingTimeIntervals(val *[]*string) {
	_jsii_.Set(
		j,
		"backupRepeatingTimeIntervals",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetDefaultRetentionDuration(val *string) {
	_jsii_.Set(
		j,
		"defaultRetentionDuration",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresql) SetVaultName(val *string) {
	_jsii_.Set(
		j,
		"vaultName",
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
func DataProtectionBackupPolicyPostgresql_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresql",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataProtectionBackupPolicyPostgresql_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresql",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) PutRetentionRule(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putRetentionRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) PutTimeouts(value *DataProtectionBackupPolicyPostgresqlTimeouts) {
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) ResetRetentionRule() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresql) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataProtectionBackupPolicyPostgresqlConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#backup_repeating_time_intervals DataProtectionBackupPolicyPostgresql#backup_repeating_time_intervals}.
	BackupRepeatingTimeIntervals *[]*string `field:"required" json:"backupRepeatingTimeIntervals" yaml:"backupRepeatingTimeIntervals"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#default_retention_duration DataProtectionBackupPolicyPostgresql#default_retention_duration}.
	DefaultRetentionDuration *string `field:"required" json:"defaultRetentionDuration" yaml:"defaultRetentionDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#name DataProtectionBackupPolicyPostgresql#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#resource_group_name DataProtectionBackupPolicyPostgresql#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#vault_name DataProtectionBackupPolicyPostgresql#vault_name}.
	VaultName *string `field:"required" json:"vaultName" yaml:"vaultName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#id DataProtectionBackupPolicyPostgresql#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// retention_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#retention_rule DataProtectionBackupPolicyPostgresql#retention_rule}
	RetentionRule interface{} `field:"optional" json:"retentionRule" yaml:"retentionRule"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#timeouts DataProtectionBackupPolicyPostgresql#timeouts}
	Timeouts *DataProtectionBackupPolicyPostgresqlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type DataProtectionBackupPolicyPostgresqlRetentionRule struct {
	// criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#criteria DataProtectionBackupPolicyPostgresql#criteria}
	Criteria *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria `field:"required" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#duration DataProtectionBackupPolicyPostgresql#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#name DataProtectionBackupPolicyPostgresql#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#priority DataProtectionBackupPolicyPostgresql#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

type DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#absolute_criteria DataProtectionBackupPolicyPostgresql#absolute_criteria}.
	AbsoluteCriteria *string `field:"optional" json:"absoluteCriteria" yaml:"absoluteCriteria"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#days_of_week DataProtectionBackupPolicyPostgresql#days_of_week}.
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#months_of_year DataProtectionBackupPolicyPostgresql#months_of_year}.
	MonthsOfYear *[]*string `field:"optional" json:"monthsOfYear" yaml:"monthsOfYear"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#scheduled_backup_times DataProtectionBackupPolicyPostgresql#scheduled_backup_times}.
	ScheduledBackupTimes *[]*string `field:"optional" json:"scheduledBackupTimes" yaml:"scheduledBackupTimes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#weeks_of_month DataProtectionBackupPolicyPostgresql#weeks_of_month}.
	WeeksOfMonth *[]*string `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}

type DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference interface {
	cdktf.ComplexObject
	AbsoluteCriteria() *string
	SetAbsoluteCriteria(val *string)
	AbsoluteCriteriaInput() *string
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
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria
	SetInternalValue(val *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria)
	MonthsOfYear() *[]*string
	SetMonthsOfYear(val *[]*string)
	MonthsOfYearInput() *[]*string
	ScheduledBackupTimes() *[]*string
	SetScheduledBackupTimes(val *[]*string)
	ScheduledBackupTimesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeksOfMonth() *[]*string
	SetWeeksOfMonth(val *[]*string)
	WeeksOfMonthInput() *[]*string
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
	ResetAbsoluteCriteria()
	ResetDaysOfWeek()
	ResetMonthsOfYear()
	ResetScheduledBackupTimes()
	ResetWeeksOfMonth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference
type jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) AbsoluteCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absoluteCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) AbsoluteCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absoluteCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) InternalValue() *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria {
	var returns *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) MonthsOfYear() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsOfYear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) MonthsOfYearInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsOfYearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) ScheduledBackupTimes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scheduledBackupTimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) ScheduledBackupTimesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scheduledBackupTimesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) WeeksOfMonth() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) WeeksOfMonthInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksOfMonthInput",
		&returns,
	)
	return returns
}


func NewDataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference_Override(d DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) SetAbsoluteCriteria(val *string) {
	_jsii_.Set(
		j,
		"absoluteCriteria",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) SetDaysOfWeek(val *[]*string) {
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) SetInternalValue(val *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) SetMonthsOfYear(val *[]*string) {
	_jsii_.Set(
		j,
		"monthsOfYear",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) SetScheduledBackupTimes(val *[]*string) {
	_jsii_.Set(
		j,
		"scheduledBackupTimes",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) SetWeeksOfMonth(val *[]*string) {
	_jsii_.Set(
		j,
		"weeksOfMonth",
		val,
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) ResetAbsoluteCriteria() {
	_jsii_.InvokeVoid(
		d,
		"resetAbsoluteCriteria",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		d,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) ResetMonthsOfYear() {
	_jsii_.InvokeVoid(
		d,
		"resetMonthsOfYear",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) ResetScheduledBackupTimes() {
	_jsii_.InvokeVoid(
		d,
		"resetScheduledBackupTimes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) ResetWeeksOfMonth() {
	_jsii_.InvokeVoid(
		d,
		"resetWeeksOfMonth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataProtectionBackupPolicyPostgresqlRetentionRuleList interface {
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
	Get(index *float64) DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataProtectionBackupPolicyPostgresqlRetentionRuleList
type jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataProtectionBackupPolicyPostgresqlRetentionRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataProtectionBackupPolicyPostgresqlRetentionRuleList {
	_init_.Initialize()

	j := jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresqlRetentionRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataProtectionBackupPolicyPostgresqlRetentionRuleList_Override(d DataProtectionBackupPolicyPostgresqlRetentionRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresqlRetentionRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) Get(index *float64) DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference {
	var returns DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference interface {
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
	Criteria() DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference
	CriteriaInput() *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	PutCriteria(value *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference
type jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) Criteria() DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference {
	var returns DataProtectionBackupPolicyPostgresqlRetentionRuleCriteriaOutputReference
	_jsii_.Get(
		j,
		"criteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) CriteriaInput() *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria {
	var returns *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria
	_jsii_.Get(
		j,
		"criteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference_Override(d DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) SetDuration(val *string) {
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) PutCriteria(value *DataProtectionBackupPolicyPostgresqlRetentionRuleCriteria) {
	_jsii_.InvokeVoid(
		d,
		"putCriteria",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlRetentionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataProtectionBackupPolicyPostgresqlTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#create DataProtectionBackupPolicyPostgresql#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#delete DataProtectionBackupPolicyPostgresql#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#read DataProtectionBackupPolicyPostgresql#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_protection_backup_policy_postgresql#update DataProtectionBackupPolicyPostgresql#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference interface {
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

// The jsii proxy struct for DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference
type jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewDataProtectionBackupPolicyPostgresqlTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataProtectionBackupPolicyPostgresqlTimeoutsOutputReference_Override(d DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupPolicyPostgresql.DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		d,
		"resetCreate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		d,
		"resetRead",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyPostgresqlTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

