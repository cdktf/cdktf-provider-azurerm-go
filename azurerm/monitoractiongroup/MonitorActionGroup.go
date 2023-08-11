package monitoractiongroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/monitoractiongroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_action_group azurerm_monitor_action_group}.
type MonitorActionGroup interface {
	cdktf.TerraformResource
	ArmRoleReceiver() MonitorActionGroupArmRoleReceiverList
	ArmRoleReceiverInput() interface{}
	AutomationRunbookReceiver() MonitorActionGroupAutomationRunbookReceiverList
	AutomationRunbookReceiverInput() interface{}
	AzureAppPushReceiver() MonitorActionGroupAzureAppPushReceiverList
	AzureAppPushReceiverInput() interface{}
	AzureFunctionReceiver() MonitorActionGroupAzureFunctionReceiverList
	AzureFunctionReceiverInput() interface{}
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
	EmailReceiver() MonitorActionGroupEmailReceiverList
	EmailReceiverInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EventHubReceiver() MonitorActionGroupEventHubReceiverList
	EventHubReceiverInput() interface{}
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
	ItsmReceiver() MonitorActionGroupItsmReceiverList
	ItsmReceiverInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LogicAppReceiver() MonitorActionGroupLogicAppReceiverList
	LogicAppReceiverInput() interface{}
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
	ShortName() *string
	SetShortName(val *string)
	ShortNameInput() *string
	SmsReceiver() MonitorActionGroupSmsReceiverList
	SmsReceiverInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MonitorActionGroupTimeoutsOutputReference
	TimeoutsInput() interface{}
	VoiceReceiver() MonitorActionGroupVoiceReceiverList
	VoiceReceiverInput() interface{}
	WebhookReceiver() MonitorActionGroupWebhookReceiverList
	WebhookReceiverInput() interface{}
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
	PutArmRoleReceiver(value interface{})
	PutAutomationRunbookReceiver(value interface{})
	PutAzureAppPushReceiver(value interface{})
	PutAzureFunctionReceiver(value interface{})
	PutEmailReceiver(value interface{})
	PutEventHubReceiver(value interface{})
	PutItsmReceiver(value interface{})
	PutLogicAppReceiver(value interface{})
	PutSmsReceiver(value interface{})
	PutTimeouts(value *MonitorActionGroupTimeouts)
	PutVoiceReceiver(value interface{})
	PutWebhookReceiver(value interface{})
	ResetArmRoleReceiver()
	ResetAutomationRunbookReceiver()
	ResetAzureAppPushReceiver()
	ResetAzureFunctionReceiver()
	ResetEmailReceiver()
	ResetEnabled()
	ResetEventHubReceiver()
	ResetId()
	ResetItsmReceiver()
	ResetLocation()
	ResetLogicAppReceiver()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSmsReceiver()
	ResetTags()
	ResetTimeouts()
	ResetVoiceReceiver()
	ResetWebhookReceiver()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MonitorActionGroup
type jsiiProxy_MonitorActionGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MonitorActionGroup) ArmRoleReceiver() MonitorActionGroupArmRoleReceiverList {
	var returns MonitorActionGroupArmRoleReceiverList
	_jsii_.Get(
		j,
		"armRoleReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) ArmRoleReceiverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"armRoleReceiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) AutomationRunbookReceiver() MonitorActionGroupAutomationRunbookReceiverList {
	var returns MonitorActionGroupAutomationRunbookReceiverList
	_jsii_.Get(
		j,
		"automationRunbookReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) AutomationRunbookReceiverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automationRunbookReceiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) AzureAppPushReceiver() MonitorActionGroupAzureAppPushReceiverList {
	var returns MonitorActionGroupAzureAppPushReceiverList
	_jsii_.Get(
		j,
		"azureAppPushReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) AzureAppPushReceiverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureAppPushReceiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) AzureFunctionReceiver() MonitorActionGroupAzureFunctionReceiverList {
	var returns MonitorActionGroupAzureFunctionReceiverList
	_jsii_.Get(
		j,
		"azureFunctionReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) AzureFunctionReceiverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureFunctionReceiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) EmailReceiver() MonitorActionGroupEmailReceiverList {
	var returns MonitorActionGroupEmailReceiverList
	_jsii_.Get(
		j,
		"emailReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) EmailReceiverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailReceiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) EventHubReceiver() MonitorActionGroupEventHubReceiverList {
	var returns MonitorActionGroupEventHubReceiverList
	_jsii_.Get(
		j,
		"eventHubReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) EventHubReceiverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventHubReceiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) ItsmReceiver() MonitorActionGroupItsmReceiverList {
	var returns MonitorActionGroupItsmReceiverList
	_jsii_.Get(
		j,
		"itsmReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) ItsmReceiverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itsmReceiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) LogicAppReceiver() MonitorActionGroupLogicAppReceiverList {
	var returns MonitorActionGroupLogicAppReceiverList
	_jsii_.Get(
		j,
		"logicAppReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) LogicAppReceiverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logicAppReceiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) ShortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) ShortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) SmsReceiver() MonitorActionGroupSmsReceiverList {
	var returns MonitorActionGroupSmsReceiverList
	_jsii_.Get(
		j,
		"smsReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) SmsReceiverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smsReceiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) Timeouts() MonitorActionGroupTimeoutsOutputReference {
	var returns MonitorActionGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) VoiceReceiver() MonitorActionGroupVoiceReceiverList {
	var returns MonitorActionGroupVoiceReceiverList
	_jsii_.Get(
		j,
		"voiceReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) VoiceReceiverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"voiceReceiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) WebhookReceiver() MonitorActionGroupWebhookReceiverList {
	var returns MonitorActionGroupWebhookReceiverList
	_jsii_.Get(
		j,
		"webhookReceiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActionGroup) WebhookReceiverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webhookReceiverInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_action_group azurerm_monitor_action_group} Resource.
func NewMonitorActionGroup(scope constructs.Construct, id *string, config *MonitorActionGroupConfig) MonitorActionGroup {
	_init_.Initialize()

	if err := validateNewMonitorActionGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorActionGroup{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActionGroup.MonitorActionGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_action_group azurerm_monitor_action_group} Resource.
func NewMonitorActionGroup_Override(m MonitorActionGroup, scope constructs.Construct, id *string, config *MonitorActionGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActionGroup.MonitorActionGroup",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetShortName(val *string) {
	if err := j.validateSetShortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shortName",
		val,
	)
}

func (j *jsiiProxy_MonitorActionGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
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
func MonitorActionGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorActionGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorActionGroup.MonitorActionGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorActionGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorActionGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorActionGroup.MonitorActionGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorActionGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorActionGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorActionGroup.MonitorActionGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MonitorActionGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.monitorActionGroup.MonitorActionGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MonitorActionGroup) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorActionGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActionGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorActionGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorActionGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorActionGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorActionGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorActionGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorActionGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorActionGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActionGroup) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutArmRoleReceiver(value interface{}) {
	if err := m.validatePutArmRoleReceiverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putArmRoleReceiver",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutAutomationRunbookReceiver(value interface{}) {
	if err := m.validatePutAutomationRunbookReceiverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAutomationRunbookReceiver",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutAzureAppPushReceiver(value interface{}) {
	if err := m.validatePutAzureAppPushReceiverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAzureAppPushReceiver",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutAzureFunctionReceiver(value interface{}) {
	if err := m.validatePutAzureFunctionReceiverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAzureFunctionReceiver",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutEmailReceiver(value interface{}) {
	if err := m.validatePutEmailReceiverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEmailReceiver",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutEventHubReceiver(value interface{}) {
	if err := m.validatePutEventHubReceiverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEventHubReceiver",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutItsmReceiver(value interface{}) {
	if err := m.validatePutItsmReceiverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putItsmReceiver",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutLogicAppReceiver(value interface{}) {
	if err := m.validatePutLogicAppReceiverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLogicAppReceiver",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutSmsReceiver(value interface{}) {
	if err := m.validatePutSmsReceiverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSmsReceiver",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutTimeouts(value *MonitorActionGroupTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutVoiceReceiver(value interface{}) {
	if err := m.validatePutVoiceReceiverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVoiceReceiver",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) PutWebhookReceiver(value interface{}) {
	if err := m.validatePutWebhookReceiverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWebhookReceiver",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetArmRoleReceiver() {
	_jsii_.InvokeVoid(
		m,
		"resetArmRoleReceiver",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetAutomationRunbookReceiver() {
	_jsii_.InvokeVoid(
		m,
		"resetAutomationRunbookReceiver",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetAzureAppPushReceiver() {
	_jsii_.InvokeVoid(
		m,
		"resetAzureAppPushReceiver",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetAzureFunctionReceiver() {
	_jsii_.InvokeVoid(
		m,
		"resetAzureFunctionReceiver",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetEmailReceiver() {
	_jsii_.InvokeVoid(
		m,
		"resetEmailReceiver",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetEventHubReceiver() {
	_jsii_.InvokeVoid(
		m,
		"resetEventHubReceiver",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetItsmReceiver() {
	_jsii_.InvokeVoid(
		m,
		"resetItsmReceiver",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetLocation() {
	_jsii_.InvokeVoid(
		m,
		"resetLocation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetLogicAppReceiver() {
	_jsii_.InvokeVoid(
		m,
		"resetLogicAppReceiver",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetSmsReceiver() {
	_jsii_.InvokeVoid(
		m,
		"resetSmsReceiver",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetVoiceReceiver() {
	_jsii_.InvokeVoid(
		m,
		"resetVoiceReceiver",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) ResetWebhookReceiver() {
	_jsii_.InvokeVoid(
		m,
		"resetWebhookReceiver",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActionGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActionGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

