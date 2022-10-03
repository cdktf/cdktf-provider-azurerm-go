package containerregistrytask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/containerregistrytask/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task azurerm_container_registry_task}.
type ContainerRegistryTask interface {
	cdktf.TerraformResource
	AgentPoolName() *string
	SetAgentPoolName(val *string)
	AgentPoolNameInput() *string
	AgentSetting() ContainerRegistryTaskAgentSettingOutputReference
	AgentSettingInput() *ContainerRegistryTaskAgentSetting
	BaseImageTrigger() ContainerRegistryTaskBaseImageTriggerOutputReference
	BaseImageTriggerInput() *ContainerRegistryTaskBaseImageTrigger
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerRegistryId() *string
	SetContainerRegistryId(val *string)
	ContainerRegistryIdInput() *string
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DockerStep() ContainerRegistryTaskDockerStepOutputReference
	DockerStepInput() *ContainerRegistryTaskDockerStep
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EncodedStep() ContainerRegistryTaskEncodedStepOutputReference
	EncodedStepInput() *ContainerRegistryTaskEncodedStep
	FileStep() ContainerRegistryTaskFileStepOutputReference
	FileStepInput() *ContainerRegistryTaskFileStep
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
	Identity() ContainerRegistryTaskIdentityOutputReference
	IdentityInput() *ContainerRegistryTaskIdentity
	IdInput() *string
	IsSystemTask() interface{}
	SetIsSystemTask(val interface{})
	IsSystemTaskInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogTemplate() *string
	SetLogTemplate(val *string)
	LogTemplateInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Platform() ContainerRegistryTaskPlatformOutputReference
	PlatformInput() *ContainerRegistryTaskPlatform
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
	RegistryCredential() ContainerRegistryTaskRegistryCredentialOutputReference
	RegistryCredentialInput() *ContainerRegistryTaskRegistryCredential
	SourceTrigger() ContainerRegistryTaskSourceTriggerList
	SourceTriggerInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeoutInSeconds() *float64
	SetTimeoutInSeconds(val *float64)
	TimeoutInSecondsInput() *float64
	Timeouts() ContainerRegistryTaskTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimerTrigger() ContainerRegistryTaskTimerTriggerList
	TimerTriggerInput() interface{}
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
	PutAgentSetting(value *ContainerRegistryTaskAgentSetting)
	PutBaseImageTrigger(value *ContainerRegistryTaskBaseImageTrigger)
	PutDockerStep(value *ContainerRegistryTaskDockerStep)
	PutEncodedStep(value *ContainerRegistryTaskEncodedStep)
	PutFileStep(value *ContainerRegistryTaskFileStep)
	PutIdentity(value *ContainerRegistryTaskIdentity)
	PutPlatform(value *ContainerRegistryTaskPlatform)
	PutRegistryCredential(value *ContainerRegistryTaskRegistryCredential)
	PutSourceTrigger(value interface{})
	PutTimeouts(value *ContainerRegistryTaskTimeouts)
	PutTimerTrigger(value interface{})
	ResetAgentPoolName()
	ResetAgentSetting()
	ResetBaseImageTrigger()
	ResetDockerStep()
	ResetEnabled()
	ResetEncodedStep()
	ResetFileStep()
	ResetId()
	ResetIdentity()
	ResetIsSystemTask()
	ResetLogTemplate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatform()
	ResetRegistryCredential()
	ResetSourceTrigger()
	ResetTags()
	ResetTimeoutInSeconds()
	ResetTimeouts()
	ResetTimerTrigger()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ContainerRegistryTask
type jsiiProxy_ContainerRegistryTask struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ContainerRegistryTask) AgentPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) AgentPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) AgentSetting() ContainerRegistryTaskAgentSettingOutputReference {
	var returns ContainerRegistryTaskAgentSettingOutputReference
	_jsii_.Get(
		j,
		"agentSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) AgentSettingInput() *ContainerRegistryTaskAgentSetting {
	var returns *ContainerRegistryTaskAgentSetting
	_jsii_.Get(
		j,
		"agentSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) BaseImageTrigger() ContainerRegistryTaskBaseImageTriggerOutputReference {
	var returns ContainerRegistryTaskBaseImageTriggerOutputReference
	_jsii_.Get(
		j,
		"baseImageTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) BaseImageTriggerInput() *ContainerRegistryTaskBaseImageTrigger {
	var returns *ContainerRegistryTaskBaseImageTrigger
	_jsii_.Get(
		j,
		"baseImageTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) ContainerRegistryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) ContainerRegistryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) DockerStep() ContainerRegistryTaskDockerStepOutputReference {
	var returns ContainerRegistryTaskDockerStepOutputReference
	_jsii_.Get(
		j,
		"dockerStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) DockerStepInput() *ContainerRegistryTaskDockerStep {
	var returns *ContainerRegistryTaskDockerStep
	_jsii_.Get(
		j,
		"dockerStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) EncodedStep() ContainerRegistryTaskEncodedStepOutputReference {
	var returns ContainerRegistryTaskEncodedStepOutputReference
	_jsii_.Get(
		j,
		"encodedStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) EncodedStepInput() *ContainerRegistryTaskEncodedStep {
	var returns *ContainerRegistryTaskEncodedStep
	_jsii_.Get(
		j,
		"encodedStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) FileStep() ContainerRegistryTaskFileStepOutputReference {
	var returns ContainerRegistryTaskFileStepOutputReference
	_jsii_.Get(
		j,
		"fileStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) FileStepInput() *ContainerRegistryTaskFileStep {
	var returns *ContainerRegistryTaskFileStep
	_jsii_.Get(
		j,
		"fileStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Identity() ContainerRegistryTaskIdentityOutputReference {
	var returns ContainerRegistryTaskIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) IdentityInput() *ContainerRegistryTaskIdentity {
	var returns *ContainerRegistryTaskIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) IsSystemTask() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSystemTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) IsSystemTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSystemTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) LogTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) LogTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Platform() ContainerRegistryTaskPlatformOutputReference {
	var returns ContainerRegistryTaskPlatformOutputReference
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) PlatformInput() *ContainerRegistryTaskPlatform {
	var returns *ContainerRegistryTaskPlatform
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) RegistryCredential() ContainerRegistryTaskRegistryCredentialOutputReference {
	var returns ContainerRegistryTaskRegistryCredentialOutputReference
	_jsii_.Get(
		j,
		"registryCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) RegistryCredentialInput() *ContainerRegistryTaskRegistryCredential {
	var returns *ContainerRegistryTaskRegistryCredential
	_jsii_.Get(
		j,
		"registryCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) SourceTrigger() ContainerRegistryTaskSourceTriggerList {
	var returns ContainerRegistryTaskSourceTriggerList
	_jsii_.Get(
		j,
		"sourceTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) SourceTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Timeouts() ContainerRegistryTaskTimeoutsOutputReference {
	var returns ContainerRegistryTaskTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TimerTrigger() ContainerRegistryTaskTimerTriggerList {
	var returns ContainerRegistryTaskTimerTriggerList
	_jsii_.Get(
		j,
		"timerTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TimerTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timerTriggerInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task azurerm_container_registry_task} Resource.
func NewContainerRegistryTask(scope constructs.Construct, id *string, config *ContainerRegistryTaskConfig) ContainerRegistryTask {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTask{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTask",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task azurerm_container_registry_task} Resource.
func NewContainerRegistryTask_Override(c ContainerRegistryTask, scope constructs.Construct, id *string, config *ContainerRegistryTaskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTask",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetAgentPoolName(val *string) {
	_jsii_.Set(
		j,
		"agentPoolName",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetContainerRegistryId(val *string) {
	_jsii_.Set(
		j,
		"containerRegistryId",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetIsSystemTask(val interface{}) {
	_jsii_.Set(
		j,
		"isSystemTask",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetLogTemplate(val *string) {
	_jsii_.Set(
		j,
		"logTemplate",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask) SetTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInSeconds",
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
func ContainerRegistryTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContainerRegistryTask_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTask",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutAgentSetting(value *ContainerRegistryTaskAgentSetting) {
	_jsii_.InvokeVoid(
		c,
		"putAgentSetting",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutBaseImageTrigger(value *ContainerRegistryTaskBaseImageTrigger) {
	_jsii_.InvokeVoid(
		c,
		"putBaseImageTrigger",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutDockerStep(value *ContainerRegistryTaskDockerStep) {
	_jsii_.InvokeVoid(
		c,
		"putDockerStep",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutEncodedStep(value *ContainerRegistryTaskEncodedStep) {
	_jsii_.InvokeVoid(
		c,
		"putEncodedStep",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutFileStep(value *ContainerRegistryTaskFileStep) {
	_jsii_.InvokeVoid(
		c,
		"putFileStep",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutIdentity(value *ContainerRegistryTaskIdentity) {
	_jsii_.InvokeVoid(
		c,
		"putIdentity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutPlatform(value *ContainerRegistryTaskPlatform) {
	_jsii_.InvokeVoid(
		c,
		"putPlatform",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutRegistryCredential(value *ContainerRegistryTaskRegistryCredential) {
	_jsii_.InvokeVoid(
		c,
		"putRegistryCredential",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutSourceTrigger(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putSourceTrigger",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutTimeouts(value *ContainerRegistryTaskTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutTimerTrigger(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putTimerTrigger",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetAgentPoolName() {
	_jsii_.InvokeVoid(
		c,
		"resetAgentPoolName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetAgentSetting() {
	_jsii_.InvokeVoid(
		c,
		"resetAgentSetting",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetBaseImageTrigger() {
	_jsii_.InvokeVoid(
		c,
		"resetBaseImageTrigger",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetDockerStep() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerStep",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetEncodedStep() {
	_jsii_.InvokeVoid(
		c,
		"resetEncodedStep",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetFileStep() {
	_jsii_.InvokeVoid(
		c,
		"resetFileStep",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetIdentity() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetIsSystemTask() {
	_jsii_.InvokeVoid(
		c,
		"resetIsSystemTask",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetLogTemplate() {
	_jsii_.InvokeVoid(
		c,
		"resetLogTemplate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetPlatform() {
	_jsii_.InvokeVoid(
		c,
		"resetPlatform",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetRegistryCredential() {
	_jsii_.InvokeVoid(
		c,
		"resetRegistryCredential",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetSourceTrigger() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceTrigger",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetTimerTrigger() {
	_jsii_.InvokeVoid(
		c,
		"resetTimerTrigger",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskAgentSetting struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#cpu ContainerRegistryTask#cpu}.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
}

type ContainerRegistryTaskAgentSettingOutputReference interface {
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
	Cpu() *float64
	SetCpu(val *float64)
	CpuInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerRegistryTaskAgentSetting
	SetInternalValue(val *ContainerRegistryTaskAgentSetting)
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

// The jsii proxy struct for ContainerRegistryTaskAgentSettingOutputReference
type jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) CpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) InternalValue() *ContainerRegistryTaskAgentSetting {
	var returns *ContainerRegistryTaskAgentSetting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskAgentSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskAgentSettingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskAgentSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskAgentSettingOutputReference_Override(c ContainerRegistryTaskAgentSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskAgentSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) SetCpu(val *float64) {
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) SetInternalValue(val *ContainerRegistryTaskAgentSetting) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskAgentSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskBaseImageTrigger struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#name ContainerRegistryTask#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#type ContainerRegistryTask#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#enabled ContainerRegistryTask#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#update_trigger_endpoint ContainerRegistryTask#update_trigger_endpoint}.
	UpdateTriggerEndpoint *string `field:"optional" json:"updateTriggerEndpoint" yaml:"updateTriggerEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#update_trigger_payload_type ContainerRegistryTask#update_trigger_payload_type}.
	UpdateTriggerPayloadType *string `field:"optional" json:"updateTriggerPayloadType" yaml:"updateTriggerPayloadType"`
}

type ContainerRegistryTaskBaseImageTriggerOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerRegistryTaskBaseImageTrigger
	SetInternalValue(val *ContainerRegistryTaskBaseImageTrigger)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UpdateTriggerEndpoint() *string
	SetUpdateTriggerEndpoint(val *string)
	UpdateTriggerEndpointInput() *string
	UpdateTriggerPayloadType() *string
	SetUpdateTriggerPayloadType(val *string)
	UpdateTriggerPayloadTypeInput() *string
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
	ResetEnabled()
	ResetUpdateTriggerEndpoint()
	ResetUpdateTriggerPayloadType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskBaseImageTriggerOutputReference
type jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) InternalValue() *ContainerRegistryTaskBaseImageTrigger {
	var returns *ContainerRegistryTaskBaseImageTrigger
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) UpdateTriggerEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTriggerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) UpdateTriggerEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTriggerEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) UpdateTriggerPayloadType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTriggerPayloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) UpdateTriggerPayloadTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTriggerPayloadTypeInput",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskBaseImageTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskBaseImageTriggerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskBaseImageTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskBaseImageTriggerOutputReference_Override(c ContainerRegistryTaskBaseImageTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskBaseImageTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) SetInternalValue(val *ContainerRegistryTaskBaseImageTrigger) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) SetUpdateTriggerEndpoint(val *string) {
	_jsii_.Set(
		j,
		"updateTriggerEndpoint",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) SetUpdateTriggerPayloadType(val *string) {
	_jsii_.Set(
		j,
		"updateTriggerPayloadType",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) ResetUpdateTriggerEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdateTriggerEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) ResetUpdateTriggerPayloadType() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdateTriggerPayloadType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskBaseImageTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#container_registry_id ContainerRegistryTask#container_registry_id}.
	ContainerRegistryId *string `field:"required" json:"containerRegistryId" yaml:"containerRegistryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#name ContainerRegistryTask#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#agent_pool_name ContainerRegistryTask#agent_pool_name}.
	AgentPoolName *string `field:"optional" json:"agentPoolName" yaml:"agentPoolName"`
	// agent_setting block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#agent_setting ContainerRegistryTask#agent_setting}
	AgentSetting *ContainerRegistryTaskAgentSetting `field:"optional" json:"agentSetting" yaml:"agentSetting"`
	// base_image_trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#base_image_trigger ContainerRegistryTask#base_image_trigger}
	BaseImageTrigger *ContainerRegistryTaskBaseImageTrigger `field:"optional" json:"baseImageTrigger" yaml:"baseImageTrigger"`
	// docker_step block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#docker_step ContainerRegistryTask#docker_step}
	DockerStep *ContainerRegistryTaskDockerStep `field:"optional" json:"dockerStep" yaml:"dockerStep"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#enabled ContainerRegistryTask#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encoded_step block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#encoded_step ContainerRegistryTask#encoded_step}
	EncodedStep *ContainerRegistryTaskEncodedStep `field:"optional" json:"encodedStep" yaml:"encodedStep"`
	// file_step block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#file_step ContainerRegistryTask#file_step}
	FileStep *ContainerRegistryTaskFileStep `field:"optional" json:"fileStep" yaml:"fileStep"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#id ContainerRegistryTask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#identity ContainerRegistryTask#identity}
	Identity *ContainerRegistryTaskIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#is_system_task ContainerRegistryTask#is_system_task}.
	IsSystemTask interface{} `field:"optional" json:"isSystemTask" yaml:"isSystemTask"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#log_template ContainerRegistryTask#log_template}.
	LogTemplate *string `field:"optional" json:"logTemplate" yaml:"logTemplate"`
	// platform block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#platform ContainerRegistryTask#platform}
	Platform *ContainerRegistryTaskPlatform `field:"optional" json:"platform" yaml:"platform"`
	// registry_credential block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#registry_credential ContainerRegistryTask#registry_credential}
	RegistryCredential *ContainerRegistryTaskRegistryCredential `field:"optional" json:"registryCredential" yaml:"registryCredential"`
	// source_trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#source_trigger ContainerRegistryTask#source_trigger}
	SourceTrigger interface{} `field:"optional" json:"sourceTrigger" yaml:"sourceTrigger"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#tags ContainerRegistryTask#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#timeout_in_seconds ContainerRegistryTask#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#timeouts ContainerRegistryTask#timeouts}
	Timeouts *ContainerRegistryTaskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// timer_trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#timer_trigger ContainerRegistryTask#timer_trigger}
	TimerTrigger interface{} `field:"optional" json:"timerTrigger" yaml:"timerTrigger"`
}

type ContainerRegistryTaskDockerStep struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#context_access_token ContainerRegistryTask#context_access_token}.
	ContextAccessToken *string `field:"required" json:"contextAccessToken" yaml:"contextAccessToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#context_path ContainerRegistryTask#context_path}.
	ContextPath *string `field:"required" json:"contextPath" yaml:"contextPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#dockerfile_path ContainerRegistryTask#dockerfile_path}.
	DockerfilePath *string `field:"required" json:"dockerfilePath" yaml:"dockerfilePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#arguments ContainerRegistryTask#arguments}.
	Arguments *map[string]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#cache_enabled ContainerRegistryTask#cache_enabled}.
	CacheEnabled interface{} `field:"optional" json:"cacheEnabled" yaml:"cacheEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#image_names ContainerRegistryTask#image_names}.
	ImageNames *[]*string `field:"optional" json:"imageNames" yaml:"imageNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#push_enabled ContainerRegistryTask#push_enabled}.
	PushEnabled interface{} `field:"optional" json:"pushEnabled" yaml:"pushEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#secret_arguments ContainerRegistryTask#secret_arguments}.
	SecretArguments *map[string]*string `field:"optional" json:"secretArguments" yaml:"secretArguments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#target ContainerRegistryTask#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

type ContainerRegistryTaskDockerStepOutputReference interface {
	cdktf.ComplexObject
	Arguments() *map[string]*string
	SetArguments(val *map[string]*string)
	ArgumentsInput() *map[string]*string
	CacheEnabled() interface{}
	SetCacheEnabled(val interface{})
	CacheEnabledInput() interface{}
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
	ContextAccessToken() *string
	SetContextAccessToken(val *string)
	ContextAccessTokenInput() *string
	ContextPath() *string
	SetContextPath(val *string)
	ContextPathInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DockerfilePath() *string
	SetDockerfilePath(val *string)
	DockerfilePathInput() *string
	// Experimental.
	Fqn() *string
	ImageNames() *[]*string
	SetImageNames(val *[]*string)
	ImageNamesInput() *[]*string
	InternalValue() *ContainerRegistryTaskDockerStep
	SetInternalValue(val *ContainerRegistryTaskDockerStep)
	PushEnabled() interface{}
	SetPushEnabled(val interface{})
	PushEnabledInput() interface{}
	SecretArguments() *map[string]*string
	SetSecretArguments(val *map[string]*string)
	SecretArgumentsInput() *map[string]*string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
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
	ResetArguments()
	ResetCacheEnabled()
	ResetImageNames()
	ResetPushEnabled()
	ResetSecretArguments()
	ResetTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskDockerStepOutputReference
type jsiiProxy_ContainerRegistryTaskDockerStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) Arguments() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ArgumentsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) CacheEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) CacheEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ContextAccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ContextAccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ContextPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ContextPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) DockerfilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) DockerfilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ImageNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ImageNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imageNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) InternalValue() *ContainerRegistryTaskDockerStep {
	var returns *ContainerRegistryTaskDockerStep
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) PushEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) PushEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SecretArguments() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SecretArgumentsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskDockerStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskDockerStepOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskDockerStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskDockerStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskDockerStepOutputReference_Override(c ContainerRegistryTaskDockerStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskDockerStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetArguments(val *map[string]*string) {
	_jsii_.Set(
		j,
		"arguments",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetCacheEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cacheEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetContextAccessToken(val *string) {
	_jsii_.Set(
		j,
		"contextAccessToken",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetContextPath(val *string) {
	_jsii_.Set(
		j,
		"contextPath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetDockerfilePath(val *string) {
	_jsii_.Set(
		j,
		"dockerfilePath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetImageNames(val *[]*string) {
	_jsii_.Set(
		j,
		"imageNames",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetInternalValue(val *ContainerRegistryTaskDockerStep) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetPushEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"pushEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetSecretArguments(val *map[string]*string) {
	_jsii_.Set(
		j,
		"secretArguments",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetTarget(val *string) {
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetArguments() {
	_jsii_.InvokeVoid(
		c,
		"resetArguments",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetCacheEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetImageNames() {
	_jsii_.InvokeVoid(
		c,
		"resetImageNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetPushEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetPushEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetSecretArguments() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretArguments",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskDockerStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskEncodedStep struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#task_content ContainerRegistryTask#task_content}.
	TaskContent *string `field:"required" json:"taskContent" yaml:"taskContent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#context_access_token ContainerRegistryTask#context_access_token}.
	ContextAccessToken *string `field:"optional" json:"contextAccessToken" yaml:"contextAccessToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#context_path ContainerRegistryTask#context_path}.
	ContextPath *string `field:"optional" json:"contextPath" yaml:"contextPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#secret_values ContainerRegistryTask#secret_values}.
	SecretValues *map[string]*string `field:"optional" json:"secretValues" yaml:"secretValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#value_content ContainerRegistryTask#value_content}.
	ValueContent *string `field:"optional" json:"valueContent" yaml:"valueContent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#values ContainerRegistryTask#values}.
	Values *map[string]*string `field:"optional" json:"values" yaml:"values"`
}

type ContainerRegistryTaskEncodedStepOutputReference interface {
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
	ContextAccessToken() *string
	SetContextAccessToken(val *string)
	ContextAccessTokenInput() *string
	ContextPath() *string
	SetContextPath(val *string)
	ContextPathInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerRegistryTaskEncodedStep
	SetInternalValue(val *ContainerRegistryTaskEncodedStep)
	SecretValues() *map[string]*string
	SetSecretValues(val *map[string]*string)
	SecretValuesInput() *map[string]*string
	TaskContent() *string
	SetTaskContent(val *string)
	TaskContentInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValueContent() *string
	SetValueContent(val *string)
	ValueContentInput() *string
	Values() *map[string]*string
	SetValues(val *map[string]*string)
	ValuesInput() *map[string]*string
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
	ResetContextAccessToken()
	ResetContextPath()
	ResetSecretValues()
	ResetValueContent()
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskEncodedStepOutputReference
type jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ContextAccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ContextAccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ContextPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ContextPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) InternalValue() *ContainerRegistryTaskEncodedStep {
	var returns *ContainerRegistryTaskEncodedStep
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SecretValues() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SecretValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) TaskContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) TaskContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ValueContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ValueContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) Values() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskEncodedStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskEncodedStepOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskEncodedStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskEncodedStepOutputReference_Override(c ContainerRegistryTaskEncodedStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskEncodedStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SetContextAccessToken(val *string) {
	_jsii_.Set(
		j,
		"contextAccessToken",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SetContextPath(val *string) {
	_jsii_.Set(
		j,
		"contextPath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SetInternalValue(val *ContainerRegistryTaskEncodedStep) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SetSecretValues(val *map[string]*string) {
	_jsii_.Set(
		j,
		"secretValues",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SetTaskContent(val *string) {
	_jsii_.Set(
		j,
		"taskContent",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SetValueContent(val *string) {
	_jsii_.Set(
		j,
		"valueContent",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) SetValues(val *map[string]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ResetContextAccessToken() {
	_jsii_.InvokeVoid(
		c,
		"resetContextAccessToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ResetContextPath() {
	_jsii_.InvokeVoid(
		c,
		"resetContextPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ResetSecretValues() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ResetValueContent() {
	_jsii_.InvokeVoid(
		c,
		"resetValueContent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		c,
		"resetValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskEncodedStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskFileStep struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#task_file_path ContainerRegistryTask#task_file_path}.
	TaskFilePath *string `field:"required" json:"taskFilePath" yaml:"taskFilePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#context_access_token ContainerRegistryTask#context_access_token}.
	ContextAccessToken *string `field:"optional" json:"contextAccessToken" yaml:"contextAccessToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#context_path ContainerRegistryTask#context_path}.
	ContextPath *string `field:"optional" json:"contextPath" yaml:"contextPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#secret_values ContainerRegistryTask#secret_values}.
	SecretValues *map[string]*string `field:"optional" json:"secretValues" yaml:"secretValues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#value_file_path ContainerRegistryTask#value_file_path}.
	ValueFilePath *string `field:"optional" json:"valueFilePath" yaml:"valueFilePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#values ContainerRegistryTask#values}.
	Values *map[string]*string `field:"optional" json:"values" yaml:"values"`
}

type ContainerRegistryTaskFileStepOutputReference interface {
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
	ContextAccessToken() *string
	SetContextAccessToken(val *string)
	ContextAccessTokenInput() *string
	ContextPath() *string
	SetContextPath(val *string)
	ContextPathInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerRegistryTaskFileStep
	SetInternalValue(val *ContainerRegistryTaskFileStep)
	SecretValues() *map[string]*string
	SetSecretValues(val *map[string]*string)
	SecretValuesInput() *map[string]*string
	TaskFilePath() *string
	SetTaskFilePath(val *string)
	TaskFilePathInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValueFilePath() *string
	SetValueFilePath(val *string)
	ValueFilePathInput() *string
	Values() *map[string]*string
	SetValues(val *map[string]*string)
	ValuesInput() *map[string]*string
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
	ResetContextAccessToken()
	ResetContextPath()
	ResetSecretValues()
	ResetValueFilePath()
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskFileStepOutputReference
type jsiiProxy_ContainerRegistryTaskFileStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ContextAccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ContextAccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextAccessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ContextPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ContextPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) InternalValue() *ContainerRegistryTaskFileStep {
	var returns *ContainerRegistryTaskFileStep
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SecretValues() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SecretValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"secretValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) TaskFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) TaskFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ValueFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ValueFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) Values() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskFileStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskFileStepOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskFileStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskFileStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskFileStepOutputReference_Override(c ContainerRegistryTaskFileStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskFileStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SetContextAccessToken(val *string) {
	_jsii_.Set(
		j,
		"contextAccessToken",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SetContextPath(val *string) {
	_jsii_.Set(
		j,
		"contextPath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SetInternalValue(val *ContainerRegistryTaskFileStep) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SetSecretValues(val *map[string]*string) {
	_jsii_.Set(
		j,
		"secretValues",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SetTaskFilePath(val *string) {
	_jsii_.Set(
		j,
		"taskFilePath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SetValueFilePath(val *string) {
	_jsii_.Set(
		j,
		"valueFilePath",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) SetValues(val *map[string]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ResetContextAccessToken() {
	_jsii_.InvokeVoid(
		c,
		"resetContextAccessToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ResetContextPath() {
	_jsii_.InvokeVoid(
		c,
		"resetContextPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ResetSecretValues() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ResetValueFilePath() {
	_jsii_.InvokeVoid(
		c,
		"resetValueFilePath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		c,
		"resetValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskFileStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#type ContainerRegistryTask#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#identity_ids ContainerRegistryTask#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

type ContainerRegistryTaskIdentityOutputReference interface {
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
	IdentityIds() *[]*string
	SetIdentityIds(val *[]*string)
	IdentityIdsInput() *[]*string
	InternalValue() *ContainerRegistryTaskIdentity
	SetInternalValue(val *ContainerRegistryTaskIdentity)
	PrincipalId() *string
	TenantId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetIdentityIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskIdentityOutputReference
type jsiiProxy_ContainerRegistryTaskIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) InternalValue() *ContainerRegistryTaskIdentity {
	var returns *ContainerRegistryTaskIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskIdentityOutputReference_Override(c ContainerRegistryTaskIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) SetInternalValue(val *ContainerRegistryTaskIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) ResetIdentityIds() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentityIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskPlatform struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#os ContainerRegistryTask#os}.
	Os *string `field:"required" json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#architecture ContainerRegistryTask#architecture}.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#variant ContainerRegistryTask#variant}.
	Variant *string `field:"optional" json:"variant" yaml:"variant"`
}

type ContainerRegistryTaskPlatformOutputReference interface {
	cdktf.ComplexObject
	Architecture() *string
	SetArchitecture(val *string)
	ArchitectureInput() *string
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
	InternalValue() *ContainerRegistryTaskPlatform
	SetInternalValue(val *ContainerRegistryTaskPlatform)
	Os() *string
	SetOs(val *string)
	OsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Variant() *string
	SetVariant(val *string)
	VariantInput() *string
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
	ResetArchitecture()
	ResetVariant()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskPlatformOutputReference
type jsiiProxy_ContainerRegistryTaskPlatformOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) InternalValue() *ContainerRegistryTaskPlatform {
	var returns *ContainerRegistryTaskPlatform
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) OsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) Variant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) VariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantInput",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskPlatformOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskPlatformOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskPlatformOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskPlatformOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskPlatformOutputReference_Override(c ContainerRegistryTaskPlatformOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskPlatformOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) SetArchitecture(val *string) {
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) SetInternalValue(val *ContainerRegistryTaskPlatform) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) SetOs(val *string) {
	_jsii_.Set(
		j,
		"os",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) SetVariant(val *string) {
	_jsii_.Set(
		j,
		"variant",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) ResetArchitecture() {
	_jsii_.InvokeVoid(
		c,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) ResetVariant() {
	_jsii_.InvokeVoid(
		c,
		"resetVariant",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskPlatformOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskRegistryCredential struct {
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#custom ContainerRegistryTask#custom}
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#source ContainerRegistryTask#source}
	Source *ContainerRegistryTaskRegistryCredentialSource `field:"optional" json:"source" yaml:"source"`
}

type ContainerRegistryTaskRegistryCredentialCustom struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#login_server ContainerRegistryTask#login_server}.
	LoginServer *string `field:"required" json:"loginServer" yaml:"loginServer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#identity ContainerRegistryTask#identity}.
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#password ContainerRegistryTask#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#username ContainerRegistryTask#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

type ContainerRegistryTaskRegistryCredentialCustomList interface {
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
	Get(index *float64) ContainerRegistryTaskRegistryCredentialCustomOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskRegistryCredentialCustomList
type jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskRegistryCredentialCustomList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ContainerRegistryTaskRegistryCredentialCustomList {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskRegistryCredentialCustomList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskRegistryCredentialCustomList_Override(c ContainerRegistryTaskRegistryCredentialCustomList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskRegistryCredentialCustomList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) Get(index *float64) ContainerRegistryTaskRegistryCredentialCustomOutputReference {
	var returns ContainerRegistryTaskRegistryCredentialCustomOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskRegistryCredentialCustomOutputReference interface {
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
	Identity() *string
	SetIdentity(val *string)
	IdentityInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoginServer() *string
	SetLoginServer(val *string)
	LoginServerInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetIdentity()
	ResetPassword()
	ResetUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskRegistryCredentialCustomOutputReference
type jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) Identity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) IdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) LoginServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) LoginServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskRegistryCredentialCustomOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ContainerRegistryTaskRegistryCredentialCustomOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskRegistryCredentialCustomOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskRegistryCredentialCustomOutputReference_Override(c ContainerRegistryTaskRegistryCredentialCustomOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskRegistryCredentialCustomOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) SetIdentity(val *string) {
	_jsii_.Set(
		j,
		"identity",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) SetLoginServer(val *string) {
	_jsii_.Set(
		j,
		"loginServer",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) ResetIdentity() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		c,
		"resetPassword",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		c,
		"resetUsername",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialCustomOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskRegistryCredentialOutputReference interface {
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
	Custom() ContainerRegistryTaskRegistryCredentialCustomList
	CustomInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerRegistryTaskRegistryCredential
	SetInternalValue(val *ContainerRegistryTaskRegistryCredential)
	Source() ContainerRegistryTaskRegistryCredentialSourceOutputReference
	SourceInput() *ContainerRegistryTaskRegistryCredentialSource
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
	PutCustom(value interface{})
	PutSource(value *ContainerRegistryTaskRegistryCredentialSource)
	ResetCustom()
	ResetSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskRegistryCredentialOutputReference
type jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) Custom() ContainerRegistryTaskRegistryCredentialCustomList {
	var returns ContainerRegistryTaskRegistryCredentialCustomList
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) CustomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) InternalValue() *ContainerRegistryTaskRegistryCredential {
	var returns *ContainerRegistryTaskRegistryCredential
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) Source() ContainerRegistryTaskRegistryCredentialSourceOutputReference {
	var returns ContainerRegistryTaskRegistryCredentialSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) SourceInput() *ContainerRegistryTaskRegistryCredentialSource {
	var returns *ContainerRegistryTaskRegistryCredentialSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskRegistryCredentialOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskRegistryCredentialOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskRegistryCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskRegistryCredentialOutputReference_Override(c ContainerRegistryTaskRegistryCredentialOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskRegistryCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) SetInternalValue(val *ContainerRegistryTaskRegistryCredential) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) PutCustom(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putCustom",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) PutSource(value *ContainerRegistryTaskRegistryCredentialSource) {
	_jsii_.InvokeVoid(
		c,
		"putSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		c,
		"resetCustom",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		c,
		"resetSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskRegistryCredentialSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#login_mode ContainerRegistryTask#login_mode}.
	LoginMode *string `field:"required" json:"loginMode" yaml:"loginMode"`
}

type ContainerRegistryTaskRegistryCredentialSourceOutputReference interface {
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
	InternalValue() *ContainerRegistryTaskRegistryCredentialSource
	SetInternalValue(val *ContainerRegistryTaskRegistryCredentialSource)
	LoginMode() *string
	SetLoginMode(val *string)
	LoginModeInput() *string
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

// The jsii proxy struct for ContainerRegistryTaskRegistryCredentialSourceOutputReference
type jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) InternalValue() *ContainerRegistryTaskRegistryCredentialSource {
	var returns *ContainerRegistryTaskRegistryCredentialSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) LoginMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) LoginModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskRegistryCredentialSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskRegistryCredentialSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskRegistryCredentialSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskRegistryCredentialSourceOutputReference_Override(c ContainerRegistryTaskRegistryCredentialSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskRegistryCredentialSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) SetInternalValue(val *ContainerRegistryTaskRegistryCredentialSource) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) SetLoginMode(val *string) {
	_jsii_.Set(
		j,
		"loginMode",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskRegistryCredentialSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskSourceTrigger struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#events ContainerRegistryTask#events}.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#name ContainerRegistryTask#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#repository_url ContainerRegistryTask#repository_url}.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#source_type ContainerRegistryTask#source_type}.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#authentication ContainerRegistryTask#authentication}
	Authentication *ContainerRegistryTaskSourceTriggerAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#branch ContainerRegistryTask#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#enabled ContainerRegistryTask#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

type ContainerRegistryTaskSourceTriggerAuthentication struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#token ContainerRegistryTask#token}.
	Token *string `field:"required" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#token_type ContainerRegistryTask#token_type}.
	TokenType *string `field:"required" json:"tokenType" yaml:"tokenType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#expire_in_seconds ContainerRegistryTask#expire_in_seconds}.
	ExpireInSeconds *float64 `field:"optional" json:"expireInSeconds" yaml:"expireInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#refresh_token ContainerRegistryTask#refresh_token}.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#scope ContainerRegistryTask#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

type ContainerRegistryTaskSourceTriggerAuthenticationOutputReference interface {
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
	ExpireInSeconds() *float64
	SetExpireInSeconds(val *float64)
	ExpireInSecondsInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerRegistryTaskSourceTriggerAuthentication
	SetInternalValue(val *ContainerRegistryTaskSourceTriggerAuthentication)
	RefreshToken() *string
	SetRefreshToken(val *string)
	RefreshTokenInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	TokenType() *string
	SetTokenType(val *string)
	TokenTypeInput() *string
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
	ResetExpireInSeconds()
	ResetRefreshToken()
	ResetScope()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskSourceTriggerAuthenticationOutputReference
type jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) ExpireInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expireInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) ExpireInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expireInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) InternalValue() *ContainerRegistryTaskSourceTriggerAuthentication {
	var returns *ContainerRegistryTaskSourceTriggerAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) RefreshToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) RefreshTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskSourceTriggerAuthenticationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskSourceTriggerAuthenticationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskSourceTriggerAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskSourceTriggerAuthenticationOutputReference_Override(c ContainerRegistryTaskSourceTriggerAuthenticationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskSourceTriggerAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) SetExpireInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"expireInSeconds",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) SetInternalValue(val *ContainerRegistryTaskSourceTriggerAuthentication) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) SetRefreshToken(val *string) {
	_jsii_.Set(
		j,
		"refreshToken",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) SetTokenType(val *string) {
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) ResetExpireInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetExpireInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) ResetRefreshToken() {
	_jsii_.InvokeVoid(
		c,
		"resetRefreshToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		c,
		"resetScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskSourceTriggerList interface {
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
	Get(index *float64) ContainerRegistryTaskSourceTriggerOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskSourceTriggerList
type jsiiProxy_ContainerRegistryTaskSourceTriggerList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskSourceTriggerList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ContainerRegistryTaskSourceTriggerList {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskSourceTriggerList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskSourceTriggerList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskSourceTriggerList_Override(c ContainerRegistryTaskSourceTriggerList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskSourceTriggerList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerList) Get(index *float64) ContainerRegistryTaskSourceTriggerOutputReference {
	var returns ContainerRegistryTaskSourceTriggerOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskSourceTriggerOutputReference interface {
	cdktf.ComplexObject
	Authentication() ContainerRegistryTaskSourceTriggerAuthenticationOutputReference
	AuthenticationInput() *ContainerRegistryTaskSourceTriggerAuthentication
	Branch() *string
	SetBranch(val *string)
	BranchInput() *string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Events() *[]*string
	SetEvents(val *[]*string)
	EventsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	RepositoryUrl() *string
	SetRepositoryUrl(val *string)
	RepositoryUrlInput() *string
	SourceType() *string
	SetSourceType(val *string)
	SourceTypeInput() *string
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
	PutAuthentication(value *ContainerRegistryTaskSourceTriggerAuthentication)
	ResetAuthentication()
	ResetBranch()
	ResetEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskSourceTriggerOutputReference
type jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) Authentication() ContainerRegistryTaskSourceTriggerAuthenticationOutputReference {
	var returns ContainerRegistryTaskSourceTriggerAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) AuthenticationInput() *ContainerRegistryTaskSourceTriggerAuthentication {
	var returns *ContainerRegistryTaskSourceTriggerAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) Events() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) EventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) RepositoryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) RepositoryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskSourceTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ContainerRegistryTaskSourceTriggerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskSourceTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskSourceTriggerOutputReference_Override(c ContainerRegistryTaskSourceTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskSourceTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SetBranch(val *string) {
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SetEvents(val *[]*string) {
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SetRepositoryUrl(val *string) {
	_jsii_.Set(
		j,
		"repositoryUrl",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SetSourceType(val *string) {
	_jsii_.Set(
		j,
		"sourceType",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) PutAuthentication(value *ContainerRegistryTaskSourceTriggerAuthentication) {
	_jsii_.InvokeVoid(
		c,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) ResetAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) ResetBranch() {
	_jsii_.InvokeVoid(
		c,
		"resetBranch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskSourceTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#create ContainerRegistryTask#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#delete ContainerRegistryTask#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#read ContainerRegistryTask#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#update ContainerRegistryTask#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ContainerRegistryTaskTimeoutsOutputReference interface {
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

// The jsii proxy struct for ContainerRegistryTaskTimeoutsOutputReference
type jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerRegistryTaskTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskTimeoutsOutputReference_Override(c ContainerRegistryTaskTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		c,
		"resetCreate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		c,
		"resetRead",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskTimerTrigger struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#name ContainerRegistryTask#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#schedule ContainerRegistryTask#schedule}.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_registry_task#enabled ContainerRegistryTask#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

type ContainerRegistryTaskTimerTriggerList interface {
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
	Get(index *float64) ContainerRegistryTaskTimerTriggerOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskTimerTriggerList
type jsiiProxy_ContainerRegistryTaskTimerTriggerList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskTimerTriggerList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ContainerRegistryTaskTimerTriggerList {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskTimerTriggerList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskTimerTriggerList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskTimerTriggerList_Override(c ContainerRegistryTaskTimerTriggerList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskTimerTriggerList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerList) Get(index *float64) ContainerRegistryTaskTimerTriggerOutputReference {
	var returns ContainerRegistryTaskTimerTriggerOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ContainerRegistryTaskTimerTriggerOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
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
	ResetEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerRegistryTaskTimerTriggerOutputReference
type jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerRegistryTaskTimerTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ContainerRegistryTaskTimerTriggerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskTimerTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewContainerRegistryTaskTimerTriggerOutputReference_Override(c ContainerRegistryTaskTimerTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistryTask.ContainerRegistryTaskTimerTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTaskTimerTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

