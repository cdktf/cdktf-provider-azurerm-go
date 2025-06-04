// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticsworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/loganalyticsworkspace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/log_analytics_workspace azurerm_log_analytics_workspace}.
type LogAnalyticsWorkspace interface {
	cdktf.TerraformResource
	AllowResourceOnlyPermissions() interface{}
	SetAllowResourceOnlyPermissions(val interface{})
	AllowResourceOnlyPermissionsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CmkForQueryForced() interface{}
	SetCmkForQueryForced(val interface{})
	CmkForQueryForcedInput() interface{}
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
	DailyQuotaGb() *float64
	SetDailyQuotaGb(val *float64)
	DailyQuotaGbInput() *float64
	DataCollectionRuleId() *string
	SetDataCollectionRuleId(val *string)
	DataCollectionRuleIdInput() *string
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
	Identity() LogAnalyticsWorkspaceIdentityOutputReference
	IdentityInput() *LogAnalyticsWorkspaceIdentity
	IdInput() *string
	ImmediateDataPurgeOn30DaysEnabled() interface{}
	SetImmediateDataPurgeOn30DaysEnabled(val interface{})
	ImmediateDataPurgeOn30DaysEnabledInput() interface{}
	InternetIngestionEnabled() interface{}
	SetInternetIngestionEnabled(val interface{})
	InternetIngestionEnabledInput() interface{}
	InternetQueryEnabled() interface{}
	SetInternetQueryEnabled(val interface{})
	InternetQueryEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalAuthenticationDisabled() interface{}
	SetLocalAuthenticationDisabled(val interface{})
	LocalAuthenticationDisabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrimarySharedKey() *string
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
	ReservationCapacityInGbPerDay() *float64
	SetReservationCapacityInGbPerDay(val *float64)
	ReservationCapacityInGbPerDayInput() *float64
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	SecondarySharedKey() *string
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LogAnalyticsWorkspaceTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkspaceId() *string
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
	PutIdentity(value *LogAnalyticsWorkspaceIdentity)
	PutTimeouts(value *LogAnalyticsWorkspaceTimeouts)
	ResetAllowResourceOnlyPermissions()
	ResetCmkForQueryForced()
	ResetDailyQuotaGb()
	ResetDataCollectionRuleId()
	ResetId()
	ResetIdentity()
	ResetImmediateDataPurgeOn30DaysEnabled()
	ResetInternetIngestionEnabled()
	ResetInternetQueryEnabled()
	ResetLocalAuthenticationDisabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReservationCapacityInGbPerDay()
	ResetRetentionInDays()
	ResetSku()
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

// The jsii proxy struct for LogAnalyticsWorkspace
type jsiiProxy_LogAnalyticsWorkspace struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LogAnalyticsWorkspace) AllowResourceOnlyPermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowResourceOnlyPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) AllowResourceOnlyPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowResourceOnlyPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) CmkForQueryForced() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmkForQueryForced",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) CmkForQueryForcedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cmkForQueryForcedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) DailyQuotaGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyQuotaGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) DailyQuotaGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyQuotaGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) DataCollectionRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCollectionRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) DataCollectionRuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCollectionRuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Identity() LogAnalyticsWorkspaceIdentityOutputReference {
	var returns LogAnalyticsWorkspaceIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) IdentityInput() *LogAnalyticsWorkspaceIdentity {
	var returns *LogAnalyticsWorkspaceIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) ImmediateDataPurgeOn30DaysEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"immediateDataPurgeOn30DaysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) ImmediateDataPurgeOn30DaysEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"immediateDataPurgeOn30DaysEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) InternetIngestionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetIngestionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) InternetIngestionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetIngestionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) InternetQueryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetQueryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) InternetQueryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetQueryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) LocalAuthenticationDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthenticationDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) LocalAuthenticationDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthenticationDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) PrimarySharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primarySharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) ReservationCapacityInGbPerDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservationCapacityInGbPerDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) ReservationCapacityInGbPerDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservationCapacityInGbPerDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) SecondarySharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondarySharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) Timeouts() LogAnalyticsWorkspaceTimeoutsOutputReference {
	var returns LogAnalyticsWorkspaceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspace) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/log_analytics_workspace azurerm_log_analytics_workspace} Resource.
func NewLogAnalyticsWorkspace(scope constructs.Construct, id *string, config *LogAnalyticsWorkspaceConfig) LogAnalyticsWorkspace {
	_init_.Initialize()

	if err := validateNewLogAnalyticsWorkspaceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogAnalyticsWorkspace{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.logAnalyticsWorkspace.LogAnalyticsWorkspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/log_analytics_workspace azurerm_log_analytics_workspace} Resource.
func NewLogAnalyticsWorkspace_Override(l LogAnalyticsWorkspace, scope constructs.Construct, id *string, config *LogAnalyticsWorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.logAnalyticsWorkspace.LogAnalyticsWorkspace",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetAllowResourceOnlyPermissions(val interface{}) {
	if err := j.validateSetAllowResourceOnlyPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowResourceOnlyPermissions",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetCmkForQueryForced(val interface{}) {
	if err := j.validateSetCmkForQueryForcedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmkForQueryForced",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetDailyQuotaGb(val *float64) {
	if err := j.validateSetDailyQuotaGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyQuotaGb",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetDataCollectionRuleId(val *string) {
	if err := j.validateSetDataCollectionRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCollectionRuleId",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetImmediateDataPurgeOn30DaysEnabled(val interface{}) {
	if err := j.validateSetImmediateDataPurgeOn30DaysEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"immediateDataPurgeOn30DaysEnabled",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetInternetIngestionEnabled(val interface{}) {
	if err := j.validateSetInternetIngestionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internetIngestionEnabled",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetInternetQueryEnabled(val interface{}) {
	if err := j.validateSetInternetQueryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internetQueryEnabled",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetLocalAuthenticationDisabled(val interface{}) {
	if err := j.validateSetLocalAuthenticationDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthenticationDisabled",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetReservationCapacityInGbPerDay(val *float64) {
	if err := j.validateSetReservationCapacityInGbPerDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservationCapacityInGbPerDay",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetRetentionInDays(val *float64) {
	if err := j.validateSetRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspace)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a LogAnalyticsWorkspace resource upon running "cdktf plan <stack-name>".
func LogAnalyticsWorkspace_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLogAnalyticsWorkspace_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logAnalyticsWorkspace.LogAnalyticsWorkspace",
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
func LogAnalyticsWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogAnalyticsWorkspace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logAnalyticsWorkspace.LogAnalyticsWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogAnalyticsWorkspace_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogAnalyticsWorkspace_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logAnalyticsWorkspace.LogAnalyticsWorkspace",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogAnalyticsWorkspace_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogAnalyticsWorkspace_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logAnalyticsWorkspace.LogAnalyticsWorkspace",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LogAnalyticsWorkspace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.logAnalyticsWorkspace.LogAnalyticsWorkspace",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) PutIdentity(value *LogAnalyticsWorkspaceIdentity) {
	if err := l.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIdentity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) PutTimeouts(value *LogAnalyticsWorkspaceTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetAllowResourceOnlyPermissions() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowResourceOnlyPermissions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetCmkForQueryForced() {
	_jsii_.InvokeVoid(
		l,
		"resetCmkForQueryForced",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetDailyQuotaGb() {
	_jsii_.InvokeVoid(
		l,
		"resetDailyQuotaGb",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetDataCollectionRuleId() {
	_jsii_.InvokeVoid(
		l,
		"resetDataCollectionRuleId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetIdentity() {
	_jsii_.InvokeVoid(
		l,
		"resetIdentity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetImmediateDataPurgeOn30DaysEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetImmediateDataPurgeOn30DaysEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetInternetIngestionEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetInternetIngestionEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetInternetQueryEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetInternetQueryEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetLocalAuthenticationDisabled() {
	_jsii_.InvokeVoid(
		l,
		"resetLocalAuthenticationDisabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetReservationCapacityInGbPerDay() {
	_jsii_.InvokeVoid(
		l,
		"resetReservationCapacityInGbPerDay",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetRetentionInDays() {
	_jsii_.InvokeVoid(
		l,
		"resetRetentionInDays",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetSku() {
	_jsii_.InvokeVoid(
		l,
		"resetSku",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspace) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

