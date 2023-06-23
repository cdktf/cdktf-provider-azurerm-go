package automationrunbook

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/automationrunbook/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/automation_runbook azurerm_automation_runbook}.
type AutomationRunbook interface {
	cdktf.TerraformResource
	AutomationAccountName() *string
	SetAutomationAccountName(val *string)
	AutomationAccountNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Draft() AutomationRunbookDraftOutputReference
	DraftInput() *AutomationRunbookDraft
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
	JobSchedule() AutomationRunbookJobScheduleList
	JobScheduleInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LogActivityTraceLevel() *float64
	SetLogActivityTraceLevel(val *float64)
	LogActivityTraceLevelInput() *float64
	LogProgress() interface{}
	SetLogProgress(val interface{})
	LogProgressInput() interface{}
	LogVerbose() interface{}
	SetLogVerbose(val interface{})
	LogVerboseInput() interface{}
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
	PublishContentLink() AutomationRunbookPublishContentLinkOutputReference
	PublishContentLinkInput() *AutomationRunbookPublishContentLink
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RunbookType() *string
	SetRunbookType(val *string)
	RunbookTypeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AutomationRunbookTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutDraft(value *AutomationRunbookDraft)
	PutJobSchedule(value interface{})
	PutPublishContentLink(value *AutomationRunbookPublishContentLink)
	PutTimeouts(value *AutomationRunbookTimeouts)
	ResetContent()
	ResetDescription()
	ResetDraft()
	ResetId()
	ResetJobSchedule()
	ResetLogActivityTraceLevel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublishContentLink()
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

// The jsii proxy struct for AutomationRunbook
type jsiiProxy_AutomationRunbook struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AutomationRunbook) AutomationAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) AutomationAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Draft() AutomationRunbookDraftOutputReference {
	var returns AutomationRunbookDraftOutputReference
	_jsii_.Get(
		j,
		"draft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) DraftInput() *AutomationRunbookDraft {
	var returns *AutomationRunbookDraft
	_jsii_.Get(
		j,
		"draftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) JobSchedule() AutomationRunbookJobScheduleList {
	var returns AutomationRunbookJobScheduleList
	_jsii_.Get(
		j,
		"jobSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) JobScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) LogActivityTraceLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logActivityTraceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) LogActivityTraceLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logActivityTraceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) LogProgress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) LogProgressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logProgressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) LogVerbose() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logVerbose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) LogVerboseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logVerboseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) PublishContentLink() AutomationRunbookPublishContentLinkOutputReference {
	var returns AutomationRunbookPublishContentLinkOutputReference
	_jsii_.Get(
		j,
		"publishContentLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) PublishContentLinkInput() *AutomationRunbookPublishContentLink {
	var returns *AutomationRunbookPublishContentLink
	_jsii_.Get(
		j,
		"publishContentLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) RunbookType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) RunbookTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) Timeouts() AutomationRunbookTimeoutsOutputReference {
	var returns AutomationRunbookTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationRunbook) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/automation_runbook azurerm_automation_runbook} Resource.
func NewAutomationRunbook(scope constructs.Construct, id *string, config *AutomationRunbookConfig) AutomationRunbook {
	_init_.Initialize()

	if err := validateNewAutomationRunbookParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomationRunbook{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationRunbook.AutomationRunbook",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/automation_runbook azurerm_automation_runbook} Resource.
func NewAutomationRunbook_Override(a AutomationRunbook, scope constructs.Construct, id *string, config *AutomationRunbookConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationRunbook.AutomationRunbook",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetAutomationAccountName(val *string) {
	if err := j.validateSetAutomationAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automationAccountName",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetLogActivityTraceLevel(val *float64) {
	if err := j.validateSetLogActivityTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logActivityTraceLevel",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetLogProgress(val interface{}) {
	if err := j.validateSetLogProgressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logProgress",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetLogVerbose(val interface{}) {
	if err := j.validateSetLogVerboseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logVerbose",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetRunbookType(val *string) {
	if err := j.validateSetRunbookTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runbookType",
		val,
	)
}

func (j *jsiiProxy_AutomationRunbook)SetTags(val *map[string]*string) {
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
func AutomationRunbook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutomationRunbook_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.automationRunbook.AutomationRunbook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutomationRunbook_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutomationRunbook_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.automationRunbook.AutomationRunbook",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutomationRunbook_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutomationRunbook_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.automationRunbook.AutomationRunbook",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AutomationRunbook_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.automationRunbook.AutomationRunbook",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AutomationRunbook) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AutomationRunbook) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutomationRunbook) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomationRunbook) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutomationRunbook) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutomationRunbook) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutomationRunbook) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutomationRunbook) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutomationRunbook) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutomationRunbook) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutomationRunbook) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomationRunbook) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AutomationRunbook) PutDraft(value *AutomationRunbookDraft) {
	if err := a.validatePutDraftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDraft",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationRunbook) PutJobSchedule(value interface{}) {
	if err := a.validatePutJobScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJobSchedule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationRunbook) PutPublishContentLink(value *AutomationRunbookPublishContentLink) {
	if err := a.validatePutPublishContentLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPublishContentLink",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationRunbook) PutTimeouts(value *AutomationRunbookTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationRunbook) ResetContent() {
	_jsii_.InvokeVoid(
		a,
		"resetContent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationRunbook) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationRunbook) ResetDraft() {
	_jsii_.InvokeVoid(
		a,
		"resetDraft",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationRunbook) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationRunbook) ResetJobSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetJobSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationRunbook) ResetLogActivityTraceLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetLogActivityTraceLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationRunbook) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationRunbook) ResetPublishContentLink() {
	_jsii_.InvokeVoid(
		a,
		"resetPublishContentLink",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationRunbook) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationRunbook) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationRunbook) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationRunbook) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationRunbook) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationRunbook) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

