// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package newrelictagrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/newrelictagrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/new_relic_tag_rule azurerm_new_relic_tag_rule}.
type NewRelicTagRule interface {
	cdktf.TerraformResource
	ActivityLogEnabled() interface{}
	SetActivityLogEnabled(val interface{})
	ActivityLogEnabledInput() interface{}
	AzureActiveDirectoryLogEnabled() interface{}
	SetAzureActiveDirectoryLogEnabled(val interface{})
	AzureActiveDirectoryLogEnabledInput() interface{}
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogTagFilter() NewRelicTagRuleLogTagFilterList
	LogTagFilterInput() interface{}
	MetricEnabled() interface{}
	SetMetricEnabled(val interface{})
	MetricEnabledInput() interface{}
	MetricTagFilter() NewRelicTagRuleMetricTagFilterList
	MetricTagFilterInput() interface{}
	MonitorId() *string
	SetMonitorId(val *string)
	MonitorIdInput() *string
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
	SubscriptionLogEnabled() interface{}
	SetSubscriptionLogEnabled(val interface{})
	SubscriptionLogEnabledInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NewRelicTagRuleTimeoutsOutputReference
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
	PutLogTagFilter(value interface{})
	PutMetricTagFilter(value interface{})
	PutTimeouts(value *NewRelicTagRuleTimeouts)
	ResetActivityLogEnabled()
	ResetAzureActiveDirectoryLogEnabled()
	ResetId()
	ResetLogTagFilter()
	ResetMetricEnabled()
	ResetMetricTagFilter()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSubscriptionLogEnabled()
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

// The jsii proxy struct for NewRelicTagRule
type jsiiProxy_NewRelicTagRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NewRelicTagRule) ActivityLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activityLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) ActivityLogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activityLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) AzureActiveDirectoryLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureActiveDirectoryLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) AzureActiveDirectoryLogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureActiveDirectoryLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) LogTagFilter() NewRelicTagRuleLogTagFilterList {
	var returns NewRelicTagRuleLogTagFilterList
	_jsii_.Get(
		j,
		"logTagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) LogTagFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logTagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) MetricEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) MetricEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) MetricTagFilter() NewRelicTagRuleMetricTagFilterList {
	var returns NewRelicTagRuleMetricTagFilterList
	_jsii_.Get(
		j,
		"metricTagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) MetricTagFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricTagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) MonitorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) MonitorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) SubscriptionLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscriptionLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) SubscriptionLogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscriptionLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) Timeouts() NewRelicTagRuleTimeoutsOutputReference {
	var returns NewRelicTagRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicTagRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/new_relic_tag_rule azurerm_new_relic_tag_rule} Resource.
func NewNewRelicTagRule(scope constructs.Construct, id *string, config *NewRelicTagRuleConfig) NewRelicTagRule {
	_init_.Initialize()

	if err := validateNewNewRelicTagRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NewRelicTagRule{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.newRelicTagRule.NewRelicTagRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/new_relic_tag_rule azurerm_new_relic_tag_rule} Resource.
func NewNewRelicTagRule_Override(n NewRelicTagRule, scope constructs.Construct, id *string, config *NewRelicTagRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.newRelicTagRule.NewRelicTagRule",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetActivityLogEnabled(val interface{}) {
	if err := j.validateSetActivityLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activityLogEnabled",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetAzureActiveDirectoryLogEnabled(val interface{}) {
	if err := j.validateSetAzureActiveDirectoryLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureActiveDirectoryLogEnabled",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetMetricEnabled(val interface{}) {
	if err := j.validateSetMetricEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricEnabled",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetMonitorId(val *string) {
	if err := j.validateSetMonitorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorId",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NewRelicTagRule)SetSubscriptionLogEnabled(val interface{}) {
	if err := j.validateSetSubscriptionLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionLogEnabled",
		val,
	)
}

// Generates CDKTF code for importing a NewRelicTagRule resource upon running "cdktf plan <stack-name>".
func NewRelicTagRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNewRelicTagRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.newRelicTagRule.NewRelicTagRule",
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
func NewRelicTagRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNewRelicTagRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.newRelicTagRule.NewRelicTagRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NewRelicTagRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNewRelicTagRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.newRelicTagRule.NewRelicTagRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NewRelicTagRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNewRelicTagRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.newRelicTagRule.NewRelicTagRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NewRelicTagRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.newRelicTagRule.NewRelicTagRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NewRelicTagRule) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NewRelicTagRule) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NewRelicTagRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NewRelicTagRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NewRelicTagRule) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NewRelicTagRule) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NewRelicTagRule) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NewRelicTagRule) PutLogTagFilter(value interface{}) {
	if err := n.validatePutLogTagFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putLogTagFilter",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NewRelicTagRule) PutMetricTagFilter(value interface{}) {
	if err := n.validatePutMetricTagFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putMetricTagFilter",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NewRelicTagRule) PutTimeouts(value *NewRelicTagRuleTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NewRelicTagRule) ResetActivityLogEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetActivityLogEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicTagRule) ResetAzureActiveDirectoryLogEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetAzureActiveDirectoryLogEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicTagRule) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicTagRule) ResetLogTagFilter() {
	_jsii_.InvokeVoid(
		n,
		"resetLogTagFilter",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicTagRule) ResetMetricEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetMetricEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicTagRule) ResetMetricTagFilter() {
	_jsii_.InvokeVoid(
		n,
		"resetMetricTagFilter",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicTagRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicTagRule) ResetSubscriptionLogEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetSubscriptionLogEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicTagRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicTagRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicTagRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

