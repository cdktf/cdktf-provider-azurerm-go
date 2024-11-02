// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consumptionbudgetmanagementgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/consumptionbudgetmanagementgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/consumption_budget_management_group azurerm_consumption_budget_management_group}.
type ConsumptionBudgetManagementGroup interface {
	cdktf.TerraformResource
	Amount() *float64
	SetAmount(val *float64)
	AmountInput() *float64
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
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
	Filter() ConsumptionBudgetManagementGroupFilterOutputReference
	FilterInput() *ConsumptionBudgetManagementGroupFilter
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
	ManagementGroupId() *string
	SetManagementGroupId(val *string)
	ManagementGroupIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Notification() ConsumptionBudgetManagementGroupNotificationList
	NotificationInput() interface{}
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeGrain() *string
	SetTimeGrain(val *string)
	TimeGrainInput() *string
	Timeouts() ConsumptionBudgetManagementGroupTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimePeriod() ConsumptionBudgetManagementGroupTimePeriodOutputReference
	TimePeriodInput() *ConsumptionBudgetManagementGroupTimePeriod
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
	PutFilter(value *ConsumptionBudgetManagementGroupFilter)
	PutNotification(value interface{})
	PutTimeouts(value *ConsumptionBudgetManagementGroupTimeouts)
	PutTimePeriod(value *ConsumptionBudgetManagementGroupTimePeriod)
	ResetEtag()
	ResetFilter()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeGrain()
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

// The jsii proxy struct for ConsumptionBudgetManagementGroup
type jsiiProxy_ConsumptionBudgetManagementGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Amount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"amount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) AmountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"amountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Filter() ConsumptionBudgetManagementGroupFilterOutputReference {
	var returns ConsumptionBudgetManagementGroupFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) FilterInput() *ConsumptionBudgetManagementGroupFilter {
	var returns *ConsumptionBudgetManagementGroupFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) ManagementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) ManagementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Notification() ConsumptionBudgetManagementGroupNotificationList {
	var returns ConsumptionBudgetManagementGroupNotificationList
	_jsii_.Get(
		j,
		"notification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) NotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) TimeGrain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeGrain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) TimeGrainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeGrainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) Timeouts() ConsumptionBudgetManagementGroupTimeoutsOutputReference {
	var returns ConsumptionBudgetManagementGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) TimePeriod() ConsumptionBudgetManagementGroupTimePeriodOutputReference {
	var returns ConsumptionBudgetManagementGroupTimePeriodOutputReference
	_jsii_.Get(
		j,
		"timePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup) TimePeriodInput() *ConsumptionBudgetManagementGroupTimePeriod {
	var returns *ConsumptionBudgetManagementGroupTimePeriod
	_jsii_.Get(
		j,
		"timePeriodInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/consumption_budget_management_group azurerm_consumption_budget_management_group} Resource.
func NewConsumptionBudgetManagementGroup(scope constructs.Construct, id *string, config *ConsumptionBudgetManagementGroupConfig) ConsumptionBudgetManagementGroup {
	_init_.Initialize()

	if err := validateNewConsumptionBudgetManagementGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConsumptionBudgetManagementGroup{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.consumptionBudgetManagementGroup.ConsumptionBudgetManagementGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/consumption_budget_management_group azurerm_consumption_budget_management_group} Resource.
func NewConsumptionBudgetManagementGroup_Override(c ConsumptionBudgetManagementGroup, scope constructs.Construct, id *string, config *ConsumptionBudgetManagementGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.consumptionBudgetManagementGroup.ConsumptionBudgetManagementGroup",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetAmount(val *float64) {
	if err := j.validateSetAmountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amount",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetManagementGroupId(val *string) {
	if err := j.validateSetManagementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managementGroupId",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ConsumptionBudgetManagementGroup)SetTimeGrain(val *string) {
	if err := j.validateSetTimeGrainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeGrain",
		val,
	)
}

// Generates CDKTF code for importing a ConsumptionBudgetManagementGroup resource upon running "cdktf plan <stack-name>".
func ConsumptionBudgetManagementGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateConsumptionBudgetManagementGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.consumptionBudgetManagementGroup.ConsumptionBudgetManagementGroup",
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
func ConsumptionBudgetManagementGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsumptionBudgetManagementGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.consumptionBudgetManagementGroup.ConsumptionBudgetManagementGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConsumptionBudgetManagementGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsumptionBudgetManagementGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.consumptionBudgetManagementGroup.ConsumptionBudgetManagementGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConsumptionBudgetManagementGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsumptionBudgetManagementGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.consumptionBudgetManagementGroup.ConsumptionBudgetManagementGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConsumptionBudgetManagementGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.consumptionBudgetManagementGroup.ConsumptionBudgetManagementGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) PutFilter(value *ConsumptionBudgetManagementGroupFilter) {
	if err := c.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFilter",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) PutNotification(value interface{}) {
	if err := c.validatePutNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNotification",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) PutTimeouts(value *ConsumptionBudgetManagementGroupTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) PutTimePeriod(value *ConsumptionBudgetManagementGroupTimePeriod) {
	if err := c.validatePutTimePeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimePeriod",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) ResetEtag() {
	_jsii_.InvokeVoid(
		c,
		"resetEtag",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) ResetFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) ResetTimeGrain() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeGrain",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumptionBudgetManagementGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

