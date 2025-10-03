// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package signalrservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/signalrservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/signalr_service azurerm_signalr_service}.
type SignalrService interface {
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
	ConnectivityLogsEnabled() interface{}
	SetConnectivityLogsEnabled(val interface{})
	ConnectivityLogsEnabledInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Cors() SignalrServiceCorsList
	CorsInput() interface{}
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
	Hostname() *string
	HttpRequestLogsEnabled() interface{}
	SetHttpRequestLogsEnabled(val interface{})
	HttpRequestLogsEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() SignalrServiceIdentityOutputReference
	IdentityInput() *SignalrServiceIdentity
	IdInput() *string
	IpAddress() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LiveTrace() SignalrServiceLiveTraceOutputReference
	LiveTraceEnabled() interface{}
	SetLiveTraceEnabled(val interface{})
	LiveTraceEnabledInput() interface{}
	LiveTraceInput() *SignalrServiceLiveTrace
	LocalAuthEnabled() interface{}
	SetLocalAuthEnabled(val interface{})
	LocalAuthEnabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MessagingLogsEnabled() interface{}
	SetMessagingLogsEnabled(val interface{})
	MessagingLogsEnabledInput() interface{}
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
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	PublicPort() *float64
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SecondaryAccessKey() *string
	SecondaryConnectionString() *string
	ServerlessConnectionTimeoutInSeconds() *float64
	SetServerlessConnectionTimeoutInSeconds(val *float64)
	ServerlessConnectionTimeoutInSecondsInput() *float64
	ServerPort() *float64
	ServiceMode() *string
	SetServiceMode(val *string)
	ServiceModeInput() *string
	Sku() SignalrServiceSkuOutputReference
	SkuInput() *SignalrServiceSku
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SignalrServiceTimeoutsOutputReference
	TimeoutsInput() interface{}
	TlsClientCertEnabled() interface{}
	SetTlsClientCertEnabled(val interface{})
	TlsClientCertEnabledInput() interface{}
	UpstreamEndpoint() SignalrServiceUpstreamEndpointList
	UpstreamEndpointInput() interface{}
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
	PutCors(value interface{})
	PutIdentity(value *SignalrServiceIdentity)
	PutLiveTrace(value *SignalrServiceLiveTrace)
	PutSku(value *SignalrServiceSku)
	PutTimeouts(value *SignalrServiceTimeouts)
	PutUpstreamEndpoint(value interface{})
	ResetAadAuthEnabled()
	ResetConnectivityLogsEnabled()
	ResetCors()
	ResetHttpRequestLogsEnabled()
	ResetId()
	ResetIdentity()
	ResetLiveTrace()
	ResetLiveTraceEnabled()
	ResetLocalAuthEnabled()
	ResetMessagingLogsEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetServerlessConnectionTimeoutInSeconds()
	ResetServiceMode()
	ResetTags()
	ResetTimeouts()
	ResetTlsClientCertEnabled()
	ResetUpstreamEndpoint()
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

// The jsii proxy struct for SignalrService
type jsiiProxy_SignalrService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SignalrService) AadAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aadAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) AadAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aadAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) ConnectivityLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectivityLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) ConnectivityLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectivityLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Cors() SignalrServiceCorsList {
	var returns SignalrServiceCorsList
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) CorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) HttpRequestLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpRequestLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) HttpRequestLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpRequestLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Identity() SignalrServiceIdentityOutputReference {
	var returns SignalrServiceIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) IdentityInput() *SignalrServiceIdentity {
	var returns *SignalrServiceIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) LiveTrace() SignalrServiceLiveTraceOutputReference {
	var returns SignalrServiceLiveTraceOutputReference
	_jsii_.Get(
		j,
		"liveTrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) LiveTraceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"liveTraceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) LiveTraceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"liveTraceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) LiveTraceInput() *SignalrServiceLiveTrace {
	var returns *SignalrServiceLiveTrace
	_jsii_.Get(
		j,
		"liveTraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) LocalAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) LocalAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) MessagingLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messagingLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) MessagingLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messagingLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) PrimaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) PrimaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) PublicPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"publicPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) SecondaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) SecondaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) ServerlessConnectionTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverlessConnectionTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) ServerlessConnectionTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverlessConnectionTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) ServerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) ServiceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) ServiceModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Sku() SignalrServiceSkuOutputReference {
	var returns SignalrServiceSkuOutputReference
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) SkuInput() *SignalrServiceSku {
	var returns *SignalrServiceSku
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) Timeouts() SignalrServiceTimeoutsOutputReference {
	var returns SignalrServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) TlsClientCertEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsClientCertEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) TlsClientCertEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsClientCertEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) UpstreamEndpoint() SignalrServiceUpstreamEndpointList {
	var returns SignalrServiceUpstreamEndpointList
	_jsii_.Get(
		j,
		"upstreamEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignalrService) UpstreamEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upstreamEndpointInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/signalr_service azurerm_signalr_service} Resource.
func NewSignalrService(scope constructs.Construct, id *string, config *SignalrServiceConfig) SignalrService {
	_init_.Initialize()

	if err := validateNewSignalrServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SignalrService{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.signalrService.SignalrService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/signalr_service azurerm_signalr_service} Resource.
func NewSignalrService_Override(s SignalrService, scope constructs.Construct, id *string, config *SignalrServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.signalrService.SignalrService",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SignalrService)SetAadAuthEnabled(val interface{}) {
	if err := j.validateSetAadAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aadAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetConnectivityLogsEnabled(val interface{}) {
	if err := j.validateSetConnectivityLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectivityLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetHttpRequestLogsEnabled(val interface{}) {
	if err := j.validateSetHttpRequestLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRequestLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetLiveTraceEnabled(val interface{}) {
	if err := j.validateSetLiveTraceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveTraceEnabled",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetLocalAuthEnabled(val interface{}) {
	if err := j.validateSetLocalAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetMessagingLogsEnabled(val interface{}) {
	if err := j.validateSetMessagingLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messagingLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetServerlessConnectionTimeoutInSeconds(val *float64) {
	if err := j.validateSetServerlessConnectionTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverlessConnectionTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetServiceMode(val *string) {
	if err := j.validateSetServiceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceMode",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SignalrService)SetTlsClientCertEnabled(val interface{}) {
	if err := j.validateSetTlsClientCertEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsClientCertEnabled",
		val,
	)
}

// Generates CDKTF code for importing a SignalrService resource upon running "cdktf plan <stack-name>".
func SignalrService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSignalrService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.signalrService.SignalrService",
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
func SignalrService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSignalrService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.signalrService.SignalrService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SignalrService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSignalrService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.signalrService.SignalrService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SignalrService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSignalrService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.signalrService.SignalrService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SignalrService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.signalrService.SignalrService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SignalrService) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SignalrService) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SignalrService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SignalrService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SignalrService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SignalrService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SignalrService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SignalrService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SignalrService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SignalrService) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SignalrService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SignalrService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrService) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SignalrService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SignalrService) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SignalrService) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SignalrService) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SignalrService) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SignalrService) PutCors(value interface{}) {
	if err := s.validatePutCorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCors",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SignalrService) PutIdentity(value *SignalrServiceIdentity) {
	if err := s.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIdentity",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SignalrService) PutLiveTrace(value *SignalrServiceLiveTrace) {
	if err := s.validatePutLiveTraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLiveTrace",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SignalrService) PutSku(value *SignalrServiceSku) {
	if err := s.validatePutSkuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSku",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SignalrService) PutTimeouts(value *SignalrServiceTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SignalrService) PutUpstreamEndpoint(value interface{}) {
	if err := s.validatePutUpstreamEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUpstreamEndpoint",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SignalrService) ResetAadAuthEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetAadAuthEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetConnectivityLogsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetConnectivityLogsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetCors() {
	_jsii_.InvokeVoid(
		s,
		"resetCors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetHttpRequestLogsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpRequestLogsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetIdentity() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetLiveTrace() {
	_jsii_.InvokeVoid(
		s,
		"resetLiveTrace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetLiveTraceEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetLiveTraceEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetLocalAuthEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetLocalAuthEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetMessagingLogsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetMessagingLogsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetServerlessConnectionTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetServerlessConnectionTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetServiceMode() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetTlsClientCertEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetTlsClientCertEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) ResetUpstreamEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetUpstreamEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignalrService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignalrService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

