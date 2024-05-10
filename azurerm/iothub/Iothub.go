// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/iothub/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/iothub azurerm_iothub}.
type Iothub interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudToDevice() IothubCloudToDeviceOutputReference
	CloudToDeviceInput() *IothubCloudToDevice
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
	Endpoint() IothubEndpointList
	EndpointInput() interface{}
	Enrichment() IothubEnrichmentList
	EnrichmentInput() interface{}
	EventHubEventsEndpoint() *string
	EventHubEventsNamespace() *string
	EventHubEventsPath() *string
	EventHubOperationsEndpoint() *string
	EventHubOperationsPath() *string
	EventHubPartitionCount() *float64
	SetEventHubPartitionCount(val *float64)
	EventHubPartitionCountInput() *float64
	EventHubRetentionInDays() *float64
	SetEventHubRetentionInDays(val *float64)
	EventHubRetentionInDaysInput() *float64
	FallbackRoute() IothubFallbackRouteOutputReference
	FallbackRouteInput() *IothubFallbackRoute
	FileUpload() IothubFileUploadOutputReference
	FileUploadInput() *IothubFileUpload
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	Id() *string
	SetId(val *string)
	Identity() IothubIdentityOutputReference
	IdentityInput() *IothubIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalAuthenticationEnabled() interface{}
	SetLocalAuthenticationEnabled(val interface{})
	LocalAuthenticationEnabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MinTlsVersion() *string
	SetMinTlsVersion(val *string)
	MinTlsVersionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkRuleSet() IothubNetworkRuleSetList
	NetworkRuleSetInput() interface{}
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
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Route() IothubRouteList
	RouteInput() interface{}
	SharedAccessPolicy() IothubSharedAccessPolicyList
	Sku() IothubSkuOutputReference
	SkuInput() *IothubSku
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() IothubTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
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
	PutCloudToDevice(value *IothubCloudToDevice)
	PutEndpoint(value interface{})
	PutEnrichment(value interface{})
	PutFallbackRoute(value *IothubFallbackRoute)
	PutFileUpload(value *IothubFileUpload)
	PutIdentity(value *IothubIdentity)
	PutNetworkRuleSet(value interface{})
	PutRoute(value interface{})
	PutSku(value *IothubSku)
	PutTimeouts(value *IothubTimeouts)
	ResetCloudToDevice()
	ResetEndpoint()
	ResetEnrichment()
	ResetEventHubPartitionCount()
	ResetEventHubRetentionInDays()
	ResetFallbackRoute()
	ResetFileUpload()
	ResetId()
	ResetIdentity()
	ResetLocalAuthenticationEnabled()
	ResetMinTlsVersion()
	ResetNetworkRuleSet()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetRoute()
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

// The jsii proxy struct for Iothub
type jsiiProxy_Iothub struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Iothub) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) CloudToDevice() IothubCloudToDeviceOutputReference {
	var returns IothubCloudToDeviceOutputReference
	_jsii_.Get(
		j,
		"cloudToDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) CloudToDeviceInput() *IothubCloudToDevice {
	var returns *IothubCloudToDevice
	_jsii_.Get(
		j,
		"cloudToDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Endpoint() IothubEndpointList {
	var returns IothubEndpointList
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) EndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Enrichment() IothubEnrichmentList {
	var returns IothubEnrichmentList
	_jsii_.Get(
		j,
		"enrichment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) EnrichmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrichmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) EventHubEventsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubEventsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) EventHubEventsNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubEventsNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) EventHubEventsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubEventsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) EventHubOperationsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubOperationsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) EventHubOperationsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubOperationsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) EventHubPartitionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventHubPartitionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) EventHubPartitionCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventHubPartitionCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) EventHubRetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventHubRetentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) EventHubRetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventHubRetentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) FallbackRoute() IothubFallbackRouteOutputReference {
	var returns IothubFallbackRouteOutputReference
	_jsii_.Get(
		j,
		"fallbackRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) FallbackRouteInput() *IothubFallbackRoute {
	var returns *IothubFallbackRoute
	_jsii_.Get(
		j,
		"fallbackRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) FileUpload() IothubFileUploadOutputReference {
	var returns IothubFileUploadOutputReference
	_jsii_.Get(
		j,
		"fileUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) FileUploadInput() *IothubFileUpload {
	var returns *IothubFileUpload
	_jsii_.Get(
		j,
		"fileUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Identity() IothubIdentityOutputReference {
	var returns IothubIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) IdentityInput() *IothubIdentity {
	var returns *IothubIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) LocalAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) LocalAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) MinTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) MinTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) NetworkRuleSet() IothubNetworkRuleSetList {
	var returns IothubNetworkRuleSetList
	_jsii_.Get(
		j,
		"networkRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) NetworkRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Route() IothubRouteList {
	var returns IothubRouteList
	_jsii_.Get(
		j,
		"route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) RouteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) SharedAccessPolicy() IothubSharedAccessPolicyList {
	var returns IothubSharedAccessPolicyList
	_jsii_.Get(
		j,
		"sharedAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Sku() IothubSkuOutputReference {
	var returns IothubSkuOutputReference
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) SkuInput() *IothubSku {
	var returns *IothubSku
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Timeouts() IothubTimeoutsOutputReference {
	var returns IothubTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Iothub) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/iothub azurerm_iothub} Resource.
func NewIothub(scope constructs.Construct, id *string, config *IothubConfig) Iothub {
	_init_.Initialize()

	if err := validateNewIothubParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Iothub{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothub.Iothub",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/iothub azurerm_iothub} Resource.
func NewIothub_Override(i Iothub, scope constructs.Construct, id *string, config *IothubConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothub.Iothub",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_Iothub)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetEventHubPartitionCount(val *float64) {
	if err := j.validateSetEventHubPartitionCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventHubPartitionCount",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetEventHubRetentionInDays(val *float64) {
	if err := j.validateSetEventHubRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventHubRetentionInDays",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetLocalAuthenticationEnabled(val interface{}) {
	if err := j.validateSetLocalAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetMinTlsVersion(val *string) {
	if err := j.validateSetMinTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTlsVersion",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_Iothub)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a Iothub resource upon running "cdktf plan <stack-name>".
func Iothub_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIothub_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothub.Iothub",
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
func Iothub_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothub_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothub.Iothub",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Iothub_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothub_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothub.Iothub",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Iothub_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothub_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothub.Iothub",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Iothub_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.iothub.Iothub",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_Iothub) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_Iothub) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_Iothub) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_Iothub) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_Iothub) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_Iothub) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_Iothub) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_Iothub) PutCloudToDevice(value *IothubCloudToDevice) {
	if err := i.validatePutCloudToDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudToDevice",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Iothub) PutEndpoint(value interface{}) {
	if err := i.validatePutEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEndpoint",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Iothub) PutEnrichment(value interface{}) {
	if err := i.validatePutEnrichmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEnrichment",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Iothub) PutFallbackRoute(value *IothubFallbackRoute) {
	if err := i.validatePutFallbackRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFallbackRoute",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Iothub) PutFileUpload(value *IothubFileUpload) {
	if err := i.validatePutFileUploadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFileUpload",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Iothub) PutIdentity(value *IothubIdentity) {
	if err := i.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIdentity",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Iothub) PutNetworkRuleSet(value interface{}) {
	if err := i.validatePutNetworkRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNetworkRuleSet",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Iothub) PutRoute(value interface{}) {
	if err := i.validatePutRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRoute",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Iothub) PutSku(value *IothubSku) {
	if err := i.validatePutSkuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSku",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Iothub) PutTimeouts(value *IothubTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Iothub) ResetCloudToDevice() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudToDevice",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetEndpoint() {
	_jsii_.InvokeVoid(
		i,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetEnrichment() {
	_jsii_.InvokeVoid(
		i,
		"resetEnrichment",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetEventHubPartitionCount() {
	_jsii_.InvokeVoid(
		i,
		"resetEventHubPartitionCount",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetEventHubRetentionInDays() {
	_jsii_.InvokeVoid(
		i,
		"resetEventHubRetentionInDays",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetFallbackRoute() {
	_jsii_.InvokeVoid(
		i,
		"resetFallbackRoute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetFileUpload() {
	_jsii_.InvokeVoid(
		i,
		"resetFileUpload",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetIdentity() {
	_jsii_.InvokeVoid(
		i,
		"resetIdentity",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetLocalAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetLocalAuthenticationEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetMinTlsVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetMinTlsVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetNetworkRuleSet() {
	_jsii_.InvokeVoid(
		i,
		"resetNetworkRuleSet",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetRoute() {
	_jsii_.InvokeVoid(
		i,
		"resetRoute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Iothub) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Iothub) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

