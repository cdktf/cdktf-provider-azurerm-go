// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automanageconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/automanageconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/automanage_configuration azurerm_automanage_configuration}.
type AutomanageConfiguration interface {
	cdktf.TerraformResource
	Antimalware() AutomanageConfigurationAntimalwareOutputReference
	AntimalwareInput() *AutomanageConfigurationAntimalware
	AutomationAccountEnabled() interface{}
	SetAutomationAccountEnabled(val interface{})
	AutomationAccountEnabledInput() interface{}
	AzureSecurityBaseline() AutomanageConfigurationAzureSecurityBaselineOutputReference
	AzureSecurityBaselineInput() *AutomanageConfigurationAzureSecurityBaseline
	Backup() AutomanageConfigurationBackupOutputReference
	BackupInput() *AutomanageConfigurationBackup
	BootDiagnosticsEnabled() interface{}
	SetBootDiagnosticsEnabled(val interface{})
	BootDiagnosticsEnabledInput() interface{}
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
	DefenderForCloudEnabled() interface{}
	SetDefenderForCloudEnabled(val interface{})
	DefenderForCloudEnabledInput() interface{}
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
	GuestConfigurationEnabled() interface{}
	SetGuestConfigurationEnabled(val interface{})
	GuestConfigurationEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LogAnalyticsEnabled() interface{}
	SetLogAnalyticsEnabled(val interface{})
	LogAnalyticsEnabledInput() interface{}
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
	StatusChangeAlertEnabled() interface{}
	SetStatusChangeAlertEnabled(val interface{})
	StatusChangeAlertEnabledInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AutomanageConfigurationTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutAntimalware(value *AutomanageConfigurationAntimalware)
	PutAzureSecurityBaseline(value *AutomanageConfigurationAzureSecurityBaseline)
	PutBackup(value *AutomanageConfigurationBackup)
	PutTimeouts(value *AutomanageConfigurationTimeouts)
	ResetAntimalware()
	ResetAutomationAccountEnabled()
	ResetAzureSecurityBaseline()
	ResetBackup()
	ResetBootDiagnosticsEnabled()
	ResetDefenderForCloudEnabled()
	ResetGuestConfigurationEnabled()
	ResetId()
	ResetLogAnalyticsEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStatusChangeAlertEnabled()
	ResetTags()
	ResetTimeouts()
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

// The jsii proxy struct for AutomanageConfiguration
type jsiiProxy_AutomanageConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AutomanageConfiguration) Antimalware() AutomanageConfigurationAntimalwareOutputReference {
	var returns AutomanageConfigurationAntimalwareOutputReference
	_jsii_.Get(
		j,
		"antimalware",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) AntimalwareInput() *AutomanageConfigurationAntimalware {
	var returns *AutomanageConfigurationAntimalware
	_jsii_.Get(
		j,
		"antimalwareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) AutomationAccountEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automationAccountEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) AutomationAccountEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automationAccountEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) AzureSecurityBaseline() AutomanageConfigurationAzureSecurityBaselineOutputReference {
	var returns AutomanageConfigurationAzureSecurityBaselineOutputReference
	_jsii_.Get(
		j,
		"azureSecurityBaseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) AzureSecurityBaselineInput() *AutomanageConfigurationAzureSecurityBaseline {
	var returns *AutomanageConfigurationAzureSecurityBaseline
	_jsii_.Get(
		j,
		"azureSecurityBaselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Backup() AutomanageConfigurationBackupOutputReference {
	var returns AutomanageConfigurationBackupOutputReference
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) BackupInput() *AutomanageConfigurationBackup {
	var returns *AutomanageConfigurationBackup
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) BootDiagnosticsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootDiagnosticsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) BootDiagnosticsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootDiagnosticsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) DefenderForCloudEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defenderForCloudEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) DefenderForCloudEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defenderForCloudEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) GuestConfigurationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestConfigurationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) GuestConfigurationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestConfigurationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) LogAnalyticsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logAnalyticsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) LogAnalyticsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logAnalyticsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) StatusChangeAlertEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statusChangeAlertEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) StatusChangeAlertEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statusChangeAlertEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) Timeouts() AutomanageConfigurationTimeoutsOutputReference {
	var returns AutomanageConfigurationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfiguration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/automanage_configuration azurerm_automanage_configuration} Resource.
func NewAutomanageConfiguration(scope constructs.Construct, id *string, config *AutomanageConfigurationConfig) AutomanageConfiguration {
	_init_.Initialize()

	if err := validateNewAutomanageConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomanageConfiguration{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/automanage_configuration azurerm_automanage_configuration} Resource.
func NewAutomanageConfiguration_Override(a AutomanageConfiguration, scope constructs.Construct, id *string, config *AutomanageConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfiguration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetAutomationAccountEnabled(val interface{}) {
	if err := j.validateSetAutomationAccountEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automationAccountEnabled",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetBootDiagnosticsEnabled(val interface{}) {
	if err := j.validateSetBootDiagnosticsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiagnosticsEnabled",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetDefenderForCloudEnabled(val interface{}) {
	if err := j.validateSetDefenderForCloudEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defenderForCloudEnabled",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetGuestConfigurationEnabled(val interface{}) {
	if err := j.validateSetGuestConfigurationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestConfigurationEnabled",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetLogAnalyticsEnabled(val interface{}) {
	if err := j.validateSetLogAnalyticsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAnalyticsEnabled",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetStatusChangeAlertEnabled(val interface{}) {
	if err := j.validateSetStatusChangeAlertEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statusChangeAlertEnabled",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfiguration)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a AutomanageConfiguration resource upon running "cdktf plan <stack-name>".
func AutomanageConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAutomanageConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfiguration",
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
func AutomanageConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutomanageConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutomanageConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutomanageConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutomanageConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutomanageConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AutomanageConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AutomanageConfiguration) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AutomanageConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AutomanageConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AutomanageConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AutomanageConfiguration) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AutomanageConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AutomanageConfiguration) PutAntimalware(value *AutomanageConfigurationAntimalware) {
	if err := a.validatePutAntimalwareParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAntimalware",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomanageConfiguration) PutAzureSecurityBaseline(value *AutomanageConfigurationAzureSecurityBaseline) {
	if err := a.validatePutAzureSecurityBaselineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAzureSecurityBaseline",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomanageConfiguration) PutBackup(value *AutomanageConfigurationBackup) {
	if err := a.validatePutBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomanageConfiguration) PutTimeouts(value *AutomanageConfigurationTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetAntimalware() {
	_jsii_.InvokeVoid(
		a,
		"resetAntimalware",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetAutomationAccountEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAutomationAccountEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetAzureSecurityBaseline() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureSecurityBaseline",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetBackup() {
	_jsii_.InvokeVoid(
		a,
		"resetBackup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetBootDiagnosticsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetBootDiagnosticsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetDefenderForCloudEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDefenderForCloudEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetGuestConfigurationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetGuestConfigurationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetLogAnalyticsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetLogAnalyticsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetStatusChangeAlertEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetStatusChangeAlertEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

