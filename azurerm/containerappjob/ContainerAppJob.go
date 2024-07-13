// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/containerappjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/container_app_job azurerm_container_app_job}.
type ContainerAppJob interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerAppEnvironmentId() *string
	SetContainerAppEnvironmentId(val *string)
	ContainerAppEnvironmentIdInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EventStreamEndpoint() *string
	EventTriggerConfig() ContainerAppJobEventTriggerConfigOutputReference
	EventTriggerConfigInput() *ContainerAppJobEventTriggerConfig
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
	Identity() ContainerAppJobIdentityOutputReference
	IdentityInput() *ContainerAppJobIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManualTriggerConfig() ContainerAppJobManualTriggerConfigOutputReference
	ManualTriggerConfigInput() *ContainerAppJobManualTriggerConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutboundIpAddresses() *[]*string
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
	Registries() ContainerAppJobRegistriesList
	RegistriesInput() interface{}
	Registry() ContainerAppJobRegistryList
	RegistryInput() interface{}
	ReplicaRetryLimit() *float64
	SetReplicaRetryLimit(val *float64)
	ReplicaRetryLimitInput() *float64
	ReplicaTimeoutInSeconds() *float64
	SetReplicaTimeoutInSeconds(val *float64)
	ReplicaTimeoutInSecondsInput() *float64
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ScheduleTriggerConfig() ContainerAppJobScheduleTriggerConfigOutputReference
	ScheduleTriggerConfigInput() *ContainerAppJobScheduleTriggerConfig
	Secret() ContainerAppJobSecretList
	SecretInput() interface{}
	Secrets() ContainerAppJobSecretsList
	SecretsInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	Template() ContainerAppJobTemplateOutputReference
	TemplateInput() *ContainerAppJobTemplate
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ContainerAppJobTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkloadProfileName() *string
	SetWorkloadProfileName(val *string)
	WorkloadProfileNameInput() *string
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
	PutEventTriggerConfig(value *ContainerAppJobEventTriggerConfig)
	PutIdentity(value *ContainerAppJobIdentity)
	PutManualTriggerConfig(value *ContainerAppJobManualTriggerConfig)
	PutRegistries(value interface{})
	PutRegistry(value interface{})
	PutScheduleTriggerConfig(value *ContainerAppJobScheduleTriggerConfig)
	PutSecret(value interface{})
	PutSecrets(value interface{})
	PutTemplate(value *ContainerAppJobTemplate)
	PutTimeouts(value *ContainerAppJobTimeouts)
	ResetEventTriggerConfig()
	ResetId()
	ResetIdentity()
	ResetManualTriggerConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegistries()
	ResetRegistry()
	ResetReplicaRetryLimit()
	ResetScheduleTriggerConfig()
	ResetSecret()
	ResetSecrets()
	ResetTags()
	ResetTimeouts()
	ResetWorkloadProfileName()
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

// The jsii proxy struct for ContainerAppJob
type jsiiProxy_ContainerAppJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ContainerAppJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ContainerAppEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAppEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ContainerAppEnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAppEnvironmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) EventStreamEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventStreamEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) EventTriggerConfig() ContainerAppJobEventTriggerConfigOutputReference {
	var returns ContainerAppJobEventTriggerConfigOutputReference
	_jsii_.Get(
		j,
		"eventTriggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) EventTriggerConfigInput() *ContainerAppJobEventTriggerConfig {
	var returns *ContainerAppJobEventTriggerConfig
	_jsii_.Get(
		j,
		"eventTriggerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Identity() ContainerAppJobIdentityOutputReference {
	var returns ContainerAppJobIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) IdentityInput() *ContainerAppJobIdentity {
	var returns *ContainerAppJobIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ManualTriggerConfig() ContainerAppJobManualTriggerConfigOutputReference {
	var returns ContainerAppJobManualTriggerConfigOutputReference
	_jsii_.Get(
		j,
		"manualTriggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ManualTriggerConfigInput() *ContainerAppJobManualTriggerConfig {
	var returns *ContainerAppJobManualTriggerConfig
	_jsii_.Get(
		j,
		"manualTriggerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) OutboundIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Registries() ContainerAppJobRegistriesList {
	var returns ContainerAppJobRegistriesList
	_jsii_.Get(
		j,
		"registries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) RegistriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Registry() ContainerAppJobRegistryList {
	var returns ContainerAppJobRegistryList
	_jsii_.Get(
		j,
		"registry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) RegistryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ReplicaRetryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaRetryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ReplicaRetryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaRetryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ReplicaTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ReplicaTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ScheduleTriggerConfig() ContainerAppJobScheduleTriggerConfigOutputReference {
	var returns ContainerAppJobScheduleTriggerConfigOutputReference
	_jsii_.Get(
		j,
		"scheduleTriggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) ScheduleTriggerConfigInput() *ContainerAppJobScheduleTriggerConfig {
	var returns *ContainerAppJobScheduleTriggerConfig
	_jsii_.Get(
		j,
		"scheduleTriggerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Secret() ContainerAppJobSecretList {
	var returns ContainerAppJobSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Secrets() ContainerAppJobSecretsList {
	var returns ContainerAppJobSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) SecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Template() ContainerAppJobTemplateOutputReference {
	var returns ContainerAppJobTemplateOutputReference
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) TemplateInput() *ContainerAppJobTemplate {
	var returns *ContainerAppJobTemplate
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) Timeouts() ContainerAppJobTimeoutsOutputReference {
	var returns ContainerAppJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) WorkloadProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJob) WorkloadProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadProfileNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/container_app_job azurerm_container_app_job} Resource.
func NewContainerAppJob(scope constructs.Construct, id *string, config *ContainerAppJobConfig) ContainerAppJob {
	_init_.Initialize()

	if err := validateNewContainerAppJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerAppJob{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerAppJob.ContainerAppJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/container_app_job azurerm_container_app_job} Resource.
func NewContainerAppJob_Override(c ContainerAppJob, scope constructs.Construct, id *string, config *ContainerAppJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerAppJob.ContainerAppJob",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetContainerAppEnvironmentId(val *string) {
	if err := j.validateSetContainerAppEnvironmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerAppEnvironmentId",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetReplicaRetryLimit(val *float64) {
	if err := j.validateSetReplicaRetryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaRetryLimit",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetReplicaTimeoutInSeconds(val *float64) {
	if err := j.validateSetReplicaTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJob)SetWorkloadProfileName(val *string) {
	if err := j.validateSetWorkloadProfileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadProfileName",
		val,
	)
}

// Generates CDKTF code for importing a ContainerAppJob resource upon running "cdktf plan <stack-name>".
func ContainerAppJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateContainerAppJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerAppJob.ContainerAppJob",
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
func ContainerAppJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAppJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerAppJob.ContainerAppJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerAppJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAppJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerAppJob.ContainerAppJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerAppJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAppJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerAppJob.ContainerAppJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContainerAppJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.containerAppJob.ContainerAppJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContainerAppJob) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ContainerAppJob) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ContainerAppJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerAppJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAppJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerAppJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerAppJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerAppJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerAppJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerAppJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerAppJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerAppJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppJob) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ContainerAppJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAppJob) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerAppJob) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ContainerAppJob) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerAppJob) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ContainerAppJob) PutEventTriggerConfig(value *ContainerAppJobEventTriggerConfig) {
	if err := c.validatePutEventTriggerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEventTriggerConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppJob) PutIdentity(value *ContainerAppJobIdentity) {
	if err := c.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIdentity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppJob) PutManualTriggerConfig(value *ContainerAppJobManualTriggerConfig) {
	if err := c.validatePutManualTriggerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putManualTriggerConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppJob) PutRegistries(value interface{}) {
	if err := c.validatePutRegistriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegistries",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppJob) PutRegistry(value interface{}) {
	if err := c.validatePutRegistryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegistry",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppJob) PutScheduleTriggerConfig(value *ContainerAppJobScheduleTriggerConfig) {
	if err := c.validatePutScheduleTriggerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putScheduleTriggerConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppJob) PutSecret(value interface{}) {
	if err := c.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecret",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppJob) PutSecrets(value interface{}) {
	if err := c.validatePutSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecrets",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppJob) PutTemplate(value *ContainerAppJobTemplate) {
	if err := c.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTemplate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppJob) PutTimeouts(value *ContainerAppJobTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetEventTriggerConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetEventTriggerConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetIdentity() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetManualTriggerConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetManualTriggerConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetRegistries() {
	_jsii_.InvokeVoid(
		c,
		"resetRegistries",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetRegistry() {
	_jsii_.InvokeVoid(
		c,
		"resetRegistry",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetReplicaRetryLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetReplicaRetryLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetScheduleTriggerConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetScheduleTriggerConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetSecret() {
	_jsii_.InvokeVoid(
		c,
		"resetSecret",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetSecrets() {
	_jsii_.InvokeVoid(
		c,
		"resetSecrets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) ResetWorkloadProfileName() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkloadProfileName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

