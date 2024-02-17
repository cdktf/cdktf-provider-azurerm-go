// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryreplicationrecoveryplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/siterecoveryreplicationrecoveryplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/site_recovery_replication_recovery_plan azurerm_site_recovery_replication_recovery_plan}.
type SiteRecoveryReplicationRecoveryPlan interface {
	cdktf.TerraformResource
	AzureToAzureSettings() SiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsOutputReference
	AzureToAzureSettingsInput() *SiteRecoveryReplicationRecoveryPlanAzureToAzureSettings
	BootRecoveryGroup() SiteRecoveryReplicationRecoveryPlanBootRecoveryGroupList
	BootRecoveryGroupInput() interface{}
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
	FailoverRecoveryGroup() SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupOutputReference
	FailoverRecoveryGroupInput() *SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroup
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
	RecoveryGroup() SiteRecoveryReplicationRecoveryPlanRecoveryGroupList
	RecoveryGroupInput() interface{}
	RecoveryVaultId() *string
	SetRecoveryVaultId(val *string)
	RecoveryVaultIdInput() *string
	ShutdownRecoveryGroup() SiteRecoveryReplicationRecoveryPlanShutdownRecoveryGroupOutputReference
	ShutdownRecoveryGroupInput() *SiteRecoveryReplicationRecoveryPlanShutdownRecoveryGroup
	SourceRecoveryFabricId() *string
	SetSourceRecoveryFabricId(val *string)
	SourceRecoveryFabricIdInput() *string
	TargetRecoveryFabricId() *string
	SetTargetRecoveryFabricId(val *string)
	TargetRecoveryFabricIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SiteRecoveryReplicationRecoveryPlanTimeoutsOutputReference
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
	PutAzureToAzureSettings(value *SiteRecoveryReplicationRecoveryPlanAzureToAzureSettings)
	PutBootRecoveryGroup(value interface{})
	PutFailoverRecoveryGroup(value *SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroup)
	PutRecoveryGroup(value interface{})
	PutShutdownRecoveryGroup(value *SiteRecoveryReplicationRecoveryPlanShutdownRecoveryGroup)
	PutTimeouts(value *SiteRecoveryReplicationRecoveryPlanTimeouts)
	ResetAzureToAzureSettings()
	ResetBootRecoveryGroup()
	ResetFailoverRecoveryGroup()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRecoveryGroup()
	ResetShutdownRecoveryGroup()
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

// The jsii proxy struct for SiteRecoveryReplicationRecoveryPlan
type jsiiProxy_SiteRecoveryReplicationRecoveryPlan struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) AzureToAzureSettings() SiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsOutputReference {
	var returns SiteRecoveryReplicationRecoveryPlanAzureToAzureSettingsOutputReference
	_jsii_.Get(
		j,
		"azureToAzureSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) AzureToAzureSettingsInput() *SiteRecoveryReplicationRecoveryPlanAzureToAzureSettings {
	var returns *SiteRecoveryReplicationRecoveryPlanAzureToAzureSettings
	_jsii_.Get(
		j,
		"azureToAzureSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) BootRecoveryGroup() SiteRecoveryReplicationRecoveryPlanBootRecoveryGroupList {
	var returns SiteRecoveryReplicationRecoveryPlanBootRecoveryGroupList
	_jsii_.Get(
		j,
		"bootRecoveryGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) BootRecoveryGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootRecoveryGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) FailoverRecoveryGroup() SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupOutputReference {
	var returns SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroupOutputReference
	_jsii_.Get(
		j,
		"failoverRecoveryGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) FailoverRecoveryGroupInput() *SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroup {
	var returns *SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroup
	_jsii_.Get(
		j,
		"failoverRecoveryGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) RecoveryGroup() SiteRecoveryReplicationRecoveryPlanRecoveryGroupList {
	var returns SiteRecoveryReplicationRecoveryPlanRecoveryGroupList
	_jsii_.Get(
		j,
		"recoveryGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) RecoveryGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recoveryGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) RecoveryVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) RecoveryVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ShutdownRecoveryGroup() SiteRecoveryReplicationRecoveryPlanShutdownRecoveryGroupOutputReference {
	var returns SiteRecoveryReplicationRecoveryPlanShutdownRecoveryGroupOutputReference
	_jsii_.Get(
		j,
		"shutdownRecoveryGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ShutdownRecoveryGroupInput() *SiteRecoveryReplicationRecoveryPlanShutdownRecoveryGroup {
	var returns *SiteRecoveryReplicationRecoveryPlanShutdownRecoveryGroup
	_jsii_.Get(
		j,
		"shutdownRecoveryGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) SourceRecoveryFabricId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRecoveryFabricId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) SourceRecoveryFabricIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRecoveryFabricIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) TargetRecoveryFabricId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRecoveryFabricId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) TargetRecoveryFabricIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRecoveryFabricIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) Timeouts() SiteRecoveryReplicationRecoveryPlanTimeoutsOutputReference {
	var returns SiteRecoveryReplicationRecoveryPlanTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/site_recovery_replication_recovery_plan azurerm_site_recovery_replication_recovery_plan} Resource.
func NewSiteRecoveryReplicationRecoveryPlan(scope constructs.Construct, id *string, config *SiteRecoveryReplicationRecoveryPlanConfig) SiteRecoveryReplicationRecoveryPlan {
	_init_.Initialize()

	if err := validateNewSiteRecoveryReplicationRecoveryPlanParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SiteRecoveryReplicationRecoveryPlan{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryReplicationRecoveryPlan.SiteRecoveryReplicationRecoveryPlan",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/site_recovery_replication_recovery_plan azurerm_site_recovery_replication_recovery_plan} Resource.
func NewSiteRecoveryReplicationRecoveryPlan_Override(s SiteRecoveryReplicationRecoveryPlan, scope constructs.Construct, id *string, config *SiteRecoveryReplicationRecoveryPlanConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryReplicationRecoveryPlan.SiteRecoveryReplicationRecoveryPlan",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetRecoveryVaultId(val *string) {
	if err := j.validateSetRecoveryVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryVaultId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetSourceRecoveryFabricId(val *string) {
	if err := j.validateSetSourceRecoveryFabricIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRecoveryFabricId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicationRecoveryPlan)SetTargetRecoveryFabricId(val *string) {
	if err := j.validateSetTargetRecoveryFabricIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRecoveryFabricId",
		val,
	)
}

// Generates CDKTF code for importing a SiteRecoveryReplicationRecoveryPlan resource upon running "cdktf plan <stack-name>".
func SiteRecoveryReplicationRecoveryPlan_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSiteRecoveryReplicationRecoveryPlan_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryReplicationRecoveryPlan.SiteRecoveryReplicationRecoveryPlan",
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
func SiteRecoveryReplicationRecoveryPlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryReplicationRecoveryPlan_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryReplicationRecoveryPlan.SiteRecoveryReplicationRecoveryPlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SiteRecoveryReplicationRecoveryPlan_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryReplicationRecoveryPlan_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryReplicationRecoveryPlan.SiteRecoveryReplicationRecoveryPlan",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SiteRecoveryReplicationRecoveryPlan_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryReplicationRecoveryPlan_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryReplicationRecoveryPlan.SiteRecoveryReplicationRecoveryPlan",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SiteRecoveryReplicationRecoveryPlan_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.siteRecoveryReplicationRecoveryPlan.SiteRecoveryReplicationRecoveryPlan",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) PutAzureToAzureSettings(value *SiteRecoveryReplicationRecoveryPlanAzureToAzureSettings) {
	if err := s.validatePutAzureToAzureSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureToAzureSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) PutBootRecoveryGroup(value interface{}) {
	if err := s.validatePutBootRecoveryGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBootRecoveryGroup",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) PutFailoverRecoveryGroup(value *SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroup) {
	if err := s.validatePutFailoverRecoveryGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFailoverRecoveryGroup",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) PutRecoveryGroup(value interface{}) {
	if err := s.validatePutRecoveryGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRecoveryGroup",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) PutShutdownRecoveryGroup(value *SiteRecoveryReplicationRecoveryPlanShutdownRecoveryGroup) {
	if err := s.validatePutShutdownRecoveryGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putShutdownRecoveryGroup",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) PutTimeouts(value *SiteRecoveryReplicationRecoveryPlanTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ResetAzureToAzureSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureToAzureSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ResetBootRecoveryGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetBootRecoveryGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ResetFailoverRecoveryGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetFailoverRecoveryGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ResetRecoveryGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetRecoveryGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ResetShutdownRecoveryGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetShutdownRecoveryGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicationRecoveryPlan) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

