package backuppolicyvm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/backuppolicyvm/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm azurerm_backup_policy_vm}.
type BackupPolicyVm interface {
	cdktf.TerraformResource
	Backup() BackupPolicyVmBackupOutputReference
	BackupInput() *BackupPolicyVmBackup
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
	InstantRestoreRetentionDays() *float64
	SetInstantRestoreRetentionDays(val *float64)
	InstantRestoreRetentionDaysInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PolicyType() *string
	SetPolicyType(val *string)
	PolicyTypeInput() *string
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
	RetentionDaily() BackupPolicyVmRetentionDailyOutputReference
	RetentionDailyInput() *BackupPolicyVmRetentionDaily
	RetentionMonthly() BackupPolicyVmRetentionMonthlyOutputReference
	RetentionMonthlyInput() *BackupPolicyVmRetentionMonthly
	RetentionWeekly() BackupPolicyVmRetentionWeeklyOutputReference
	RetentionWeeklyInput() *BackupPolicyVmRetentionWeekly
	RetentionYearly() BackupPolicyVmRetentionYearlyOutputReference
	RetentionYearlyInput() *BackupPolicyVmRetentionYearly
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BackupPolicyVmTimeoutsOutputReference
	TimeoutsInput() interface{}
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
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
	PutBackup(value *BackupPolicyVmBackup)
	PutRetentionDaily(value *BackupPolicyVmRetentionDaily)
	PutRetentionMonthly(value *BackupPolicyVmRetentionMonthly)
	PutRetentionWeekly(value *BackupPolicyVmRetentionWeekly)
	PutRetentionYearly(value *BackupPolicyVmRetentionYearly)
	PutTimeouts(value *BackupPolicyVmTimeouts)
	ResetId()
	ResetInstantRestoreRetentionDays()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyType()
	ResetRetentionDaily()
	ResetRetentionMonthly()
	ResetRetentionWeekly()
	ResetRetentionYearly()
	ResetTimeouts()
	ResetTimezone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for BackupPolicyVm
type jsiiProxy_BackupPolicyVm struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BackupPolicyVm) Backup() BackupPolicyVmBackupOutputReference {
	var returns BackupPolicyVmBackupOutputReference
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) BackupInput() *BackupPolicyVmBackup {
	var returns *BackupPolicyVmBackup
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) InstantRestoreRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instantRestoreRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) InstantRestoreRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instantRestoreRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) RecoveryVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) RecoveryVaultNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) RetentionDaily() BackupPolicyVmRetentionDailyOutputReference {
	var returns BackupPolicyVmRetentionDailyOutputReference
	_jsii_.Get(
		j,
		"retentionDaily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) RetentionDailyInput() *BackupPolicyVmRetentionDaily {
	var returns *BackupPolicyVmRetentionDaily
	_jsii_.Get(
		j,
		"retentionDailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) RetentionMonthly() BackupPolicyVmRetentionMonthlyOutputReference {
	var returns BackupPolicyVmRetentionMonthlyOutputReference
	_jsii_.Get(
		j,
		"retentionMonthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) RetentionMonthlyInput() *BackupPolicyVmRetentionMonthly {
	var returns *BackupPolicyVmRetentionMonthly
	_jsii_.Get(
		j,
		"retentionMonthlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) RetentionWeekly() BackupPolicyVmRetentionWeeklyOutputReference {
	var returns BackupPolicyVmRetentionWeeklyOutputReference
	_jsii_.Get(
		j,
		"retentionWeekly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) RetentionWeeklyInput() *BackupPolicyVmRetentionWeekly {
	var returns *BackupPolicyVmRetentionWeekly
	_jsii_.Get(
		j,
		"retentionWeeklyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) RetentionYearly() BackupPolicyVmRetentionYearlyOutputReference {
	var returns BackupPolicyVmRetentionYearlyOutputReference
	_jsii_.Get(
		j,
		"retentionYearly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) RetentionYearlyInput() *BackupPolicyVmRetentionYearly {
	var returns *BackupPolicyVmRetentionYearly
	_jsii_.Get(
		j,
		"retentionYearlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) Timeouts() BackupPolicyVmTimeoutsOutputReference {
	var returns BackupPolicyVmTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVm) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm azurerm_backup_policy_vm} Resource.
func NewBackupPolicyVm(scope constructs.Construct, id *string, config *BackupPolicyVmConfig) BackupPolicyVm {
	_init_.Initialize()

	if err := validateNewBackupPolicyVmParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupPolicyVm{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVm.BackupPolicyVm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/backup_policy_vm azurerm_backup_policy_vm} Resource.
func NewBackupPolicyVm_Override(b BackupPolicyVm, scope constructs.Construct, id *string, config *BackupPolicyVmConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVm.BackupPolicyVm",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetInstantRestoreRetentionDays(val *float64) {
	if err := j.validateSetInstantRestoreRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instantRestoreRetentionDays",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetPolicyType(val *string) {
	if err := j.validateSetPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetRecoveryVaultName(val *string) {
	if err := j.validateSetRecoveryVaultNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryVaultName",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVm)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
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
func BackupPolicyVm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupPolicyVm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.backupPolicyVm.BackupPolicyVm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BackupPolicyVm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.backupPolicyVm.BackupPolicyVm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BackupPolicyVm) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BackupPolicyVm) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BackupPolicyVm) PutBackup(value *BackupPolicyVmBackup) {
	if err := b.validatePutBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putBackup",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVm) PutRetentionDaily(value *BackupPolicyVmRetentionDaily) {
	if err := b.validatePutRetentionDailyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRetentionDaily",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVm) PutRetentionMonthly(value *BackupPolicyVmRetentionMonthly) {
	if err := b.validatePutRetentionMonthlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRetentionMonthly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVm) PutRetentionWeekly(value *BackupPolicyVmRetentionWeekly) {
	if err := b.validatePutRetentionWeeklyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRetentionWeekly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVm) PutRetentionYearly(value *BackupPolicyVmRetentionYearly) {
	if err := b.validatePutRetentionYearlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRetentionYearly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVm) PutTimeouts(value *BackupPolicyVmTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVm) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVm) ResetInstantRestoreRetentionDays() {
	_jsii_.InvokeVoid(
		b,
		"resetInstantRestoreRetentionDays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVm) ResetPolicyType() {
	_jsii_.InvokeVoid(
		b,
		"resetPolicyType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVm) ResetRetentionDaily() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionDaily",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVm) ResetRetentionMonthly() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionMonthly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVm) ResetRetentionWeekly() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionWeekly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVm) ResetRetentionYearly() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionYearly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVm) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVm) ResetTimezone() {
	_jsii_.InvokeVoid(
		b,
		"resetTimezone",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVm) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

