// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webpubsubsocketio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/webpubsubsocketio/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/web_pubsub_socketio azurerm_web_pubsub_socketio}.
type WebPubsubSocketio interface {
	cdktf.TerraformResource
	AadAuthEnabled() interface{}
	SetAadAuthEnabled(val interface{})
	AadAuthEnabledInput() interface{}
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
	ExternalIp() *string
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
	Identity() WebPubsubSocketioIdentityOutputReference
	IdentityInput() *WebPubsubSocketioIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LiveTraceConnectivityLogsEnabled() interface{}
	SetLiveTraceConnectivityLogsEnabled(val interface{})
	LiveTraceConnectivityLogsEnabledInput() interface{}
	LiveTraceEnabled() interface{}
	SetLiveTraceEnabled(val interface{})
	LiveTraceEnabledInput() interface{}
	LiveTraceHttpRequestLogsEnabled() interface{}
	SetLiveTraceHttpRequestLogsEnabled(val interface{})
	LiveTraceHttpRequestLogsEnabledInput() interface{}
	LiveTraceMessagingLogsEnabled() interface{}
	SetLiveTraceMessagingLogsEnabled(val interface{})
	LiveTraceMessagingLogsEnabledInput() interface{}
	LocalAuthEnabled() interface{}
	SetLocalAuthEnabled(val interface{})
	LocalAuthEnabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrimaryAccessKey() *string
	PrimaryConnectionString() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicNetworkAccess() *string
	SetPublicNetworkAccess(val *string)
	PublicNetworkAccessInput() *string
	PublicPort() *float64
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SecondaryAccessKey() *string
	SecondaryConnectionString() *string
	ServerPort() *float64
	ServiceMode() *string
	SetServiceMode(val *string)
	ServiceModeInput() *string
	Sku() WebPubsubSocketioSkuOutputReference
	SkuInput() *WebPubsubSocketioSku
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WebPubsubSocketioTimeoutsOutputReference
	TimeoutsInput() interface{}
	TlsClientCertEnabled() interface{}
	SetTlsClientCertEnabled(val interface{})
	TlsClientCertEnabledInput() interface{}
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
	PutIdentity(value *WebPubsubSocketioIdentity)
	PutSku(value *WebPubsubSocketioSku)
	PutTimeouts(value *WebPubsubSocketioTimeouts)
	ResetAadAuthEnabled()
	ResetId()
	ResetIdentity()
	ResetLiveTraceConnectivityLogsEnabled()
	ResetLiveTraceEnabled()
	ResetLiveTraceHttpRequestLogsEnabled()
	ResetLiveTraceMessagingLogsEnabled()
	ResetLocalAuthEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccess()
	ResetServiceMode()
	ResetTags()
	ResetTimeouts()
	ResetTlsClientCertEnabled()
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

// The jsii proxy struct for WebPubsubSocketio
type jsiiProxy_WebPubsubSocketio struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WebPubsubSocketio) AadAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aadAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) AadAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aadAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) ExternalIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Identity() WebPubsubSocketioIdentityOutputReference {
	var returns WebPubsubSocketioIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) IdentityInput() *WebPubsubSocketioIdentity {
	var returns *WebPubsubSocketioIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) LiveTraceConnectivityLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"liveTraceConnectivityLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) LiveTraceConnectivityLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"liveTraceConnectivityLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) LiveTraceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"liveTraceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) LiveTraceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"liveTraceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) LiveTraceHttpRequestLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"liveTraceHttpRequestLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) LiveTraceHttpRequestLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"liveTraceHttpRequestLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) LiveTraceMessagingLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"liveTraceMessagingLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) LiveTraceMessagingLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"liveTraceMessagingLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) LocalAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) LocalAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) PrimaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) PrimaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) PublicNetworkAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNetworkAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) PublicNetworkAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNetworkAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) PublicPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"publicPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) SecondaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) SecondaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) ServerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) ServiceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) ServiceModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Sku() WebPubsubSocketioSkuOutputReference {
	var returns WebPubsubSocketioSkuOutputReference
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) SkuInput() *WebPubsubSocketioSku {
	var returns *WebPubsubSocketioSku
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) Timeouts() WebPubsubSocketioTimeoutsOutputReference {
	var returns WebPubsubSocketioTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) TlsClientCertEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsClientCertEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebPubsubSocketio) TlsClientCertEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsClientCertEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/web_pubsub_socketio azurerm_web_pubsub_socketio} Resource.
func NewWebPubsubSocketio(scope constructs.Construct, id *string, config *WebPubsubSocketioConfig) WebPubsubSocketio {
	_init_.Initialize()

	if err := validateNewWebPubsubSocketioParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebPubsubSocketio{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.webPubsubSocketio.WebPubsubSocketio",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/web_pubsub_socketio azurerm_web_pubsub_socketio} Resource.
func NewWebPubsubSocketio_Override(w WebPubsubSocketio, scope constructs.Construct, id *string, config *WebPubsubSocketioConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.webPubsubSocketio.WebPubsubSocketio",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetAadAuthEnabled(val interface{}) {
	if err := j.validateSetAadAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aadAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetLiveTraceConnectivityLogsEnabled(val interface{}) {
	if err := j.validateSetLiveTraceConnectivityLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveTraceConnectivityLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetLiveTraceEnabled(val interface{}) {
	if err := j.validateSetLiveTraceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveTraceEnabled",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetLiveTraceHttpRequestLogsEnabled(val interface{}) {
	if err := j.validateSetLiveTraceHttpRequestLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveTraceHttpRequestLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetLiveTraceMessagingLogsEnabled(val interface{}) {
	if err := j.validateSetLiveTraceMessagingLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveTraceMessagingLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetLocalAuthEnabled(val interface{}) {
	if err := j.validateSetLocalAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetPublicNetworkAccess(val *string) {
	if err := j.validateSetPublicNetworkAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccess",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetServiceMode(val *string) {
	if err := j.validateSetServiceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceMode",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WebPubsubSocketio)SetTlsClientCertEnabled(val interface{}) {
	if err := j.validateSetTlsClientCertEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsClientCertEnabled",
		val,
	)
}

// Generates CDKTF code for importing a WebPubsubSocketio resource upon running "cdktf plan <stack-name>".
func WebPubsubSocketio_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWebPubsubSocketio_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.webPubsubSocketio.WebPubsubSocketio",
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
func WebPubsubSocketio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWebPubsubSocketio_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.webPubsubSocketio.WebPubsubSocketio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WebPubsubSocketio_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWebPubsubSocketio_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.webPubsubSocketio.WebPubsubSocketio",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WebPubsubSocketio_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWebPubsubSocketio_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.webPubsubSocketio.WebPubsubSocketio",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WebPubsubSocketio_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.webPubsubSocketio.WebPubsubSocketio",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WebPubsubSocketio) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WebPubsubSocketio) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WebPubsubSocketio) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WebPubsubSocketio) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WebPubsubSocketio) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WebPubsubSocketio) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WebPubsubSocketio) PutIdentity(value *WebPubsubSocketioIdentity) {
	if err := w.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIdentity",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WebPubsubSocketio) PutSku(value *WebPubsubSocketioSku) {
	if err := w.validatePutSkuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSku",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WebPubsubSocketio) PutTimeouts(value *WebPubsubSocketioTimeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetAadAuthEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetAadAuthEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetIdentity() {
	_jsii_.InvokeVoid(
		w,
		"resetIdentity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetLiveTraceConnectivityLogsEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetLiveTraceConnectivityLogsEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetLiveTraceEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetLiveTraceEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetLiveTraceHttpRequestLogsEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetLiveTraceHttpRequestLogsEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetLiveTraceMessagingLogsEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetLiveTraceMessagingLogsEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetLocalAuthEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetLocalAuthEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetPublicNetworkAccess() {
	_jsii_.InvokeVoid(
		w,
		"resetPublicNetworkAccess",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetServiceMode() {
	_jsii_.InvokeVoid(
		w,
		"resetServiceMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) ResetTlsClientCertEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetTlsClientCertEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebPubsubSocketio) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebPubsubSocketio) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

