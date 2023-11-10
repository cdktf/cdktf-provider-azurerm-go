// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapidiagnostic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/apimanagementapidiagnostic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/api_management_api_diagnostic azurerm_api_management_api_diagnostic}.
type ApiManagementApiDiagnostic interface {
	cdktf.TerraformResource
	AlwaysLogErrors() interface{}
	SetAlwaysLogErrors(val interface{})
	AlwaysLogErrorsInput() interface{}
	ApiManagementLoggerId() *string
	SetApiManagementLoggerId(val *string)
	ApiManagementLoggerIdInput() *string
	ApiManagementName() *string
	SetApiManagementName(val *string)
	ApiManagementNameInput() *string
	ApiName() *string
	SetApiName(val *string)
	ApiNameInput() *string
	BackendRequest() ApiManagementApiDiagnosticBackendRequestOutputReference
	BackendRequestInput() *ApiManagementApiDiagnosticBackendRequest
	BackendResponse() ApiManagementApiDiagnosticBackendResponseOutputReference
	BackendResponseInput() *ApiManagementApiDiagnosticBackendResponse
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
	FrontendRequest() ApiManagementApiDiagnosticFrontendRequestOutputReference
	FrontendRequestInput() *ApiManagementApiDiagnosticFrontendRequest
	FrontendResponse() ApiManagementApiDiagnosticFrontendResponseOutputReference
	FrontendResponseInput() *ApiManagementApiDiagnosticFrontendResponse
	HttpCorrelationProtocol() *string
	SetHttpCorrelationProtocol(val *string)
	HttpCorrelationProtocolInput() *string
	Id() *string
	SetId(val *string)
	Identifier() *string
	SetIdentifier(val *string)
	IdentifierInput() *string
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogClientIp() interface{}
	SetLogClientIp(val interface{})
	LogClientIpInput() interface{}
	// The tree node.
	Node() constructs.Node
	OperationNameFormat() *string
	SetOperationNameFormat(val *string)
	OperationNameFormatInput() *string
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
	SamplingPercentage() *float64
	SetSamplingPercentage(val *float64)
	SamplingPercentageInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApiManagementApiDiagnosticTimeoutsOutputReference
	TimeoutsInput() interface{}
	Verbosity() *string
	SetVerbosity(val *string)
	VerbosityInput() *string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBackendRequest(value *ApiManagementApiDiagnosticBackendRequest)
	PutBackendResponse(value *ApiManagementApiDiagnosticBackendResponse)
	PutFrontendRequest(value *ApiManagementApiDiagnosticFrontendRequest)
	PutFrontendResponse(value *ApiManagementApiDiagnosticFrontendResponse)
	PutTimeouts(value *ApiManagementApiDiagnosticTimeouts)
	ResetAlwaysLogErrors()
	ResetBackendRequest()
	ResetBackendResponse()
	ResetFrontendRequest()
	ResetFrontendResponse()
	ResetHttpCorrelationProtocol()
	ResetId()
	ResetLogClientIp()
	ResetOperationNameFormat()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSamplingPercentage()
	ResetTimeouts()
	ResetVerbosity()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiManagementApiDiagnostic
type jsiiProxy_ApiManagementApiDiagnostic struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) AlwaysLogErrors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysLogErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) AlwaysLogErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysLogErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) ApiManagementLoggerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementLoggerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) ApiManagementLoggerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementLoggerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) ApiManagementName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) ApiManagementNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) ApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) ApiNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) BackendRequest() ApiManagementApiDiagnosticBackendRequestOutputReference {
	var returns ApiManagementApiDiagnosticBackendRequestOutputReference
	_jsii_.Get(
		j,
		"backendRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) BackendRequestInput() *ApiManagementApiDiagnosticBackendRequest {
	var returns *ApiManagementApiDiagnosticBackendRequest
	_jsii_.Get(
		j,
		"backendRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) BackendResponse() ApiManagementApiDiagnosticBackendResponseOutputReference {
	var returns ApiManagementApiDiagnosticBackendResponseOutputReference
	_jsii_.Get(
		j,
		"backendResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) BackendResponseInput() *ApiManagementApiDiagnosticBackendResponse {
	var returns *ApiManagementApiDiagnosticBackendResponse
	_jsii_.Get(
		j,
		"backendResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) FrontendRequest() ApiManagementApiDiagnosticFrontendRequestOutputReference {
	var returns ApiManagementApiDiagnosticFrontendRequestOutputReference
	_jsii_.Get(
		j,
		"frontendRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) FrontendRequestInput() *ApiManagementApiDiagnosticFrontendRequest {
	var returns *ApiManagementApiDiagnosticFrontendRequest
	_jsii_.Get(
		j,
		"frontendRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) FrontendResponse() ApiManagementApiDiagnosticFrontendResponseOutputReference {
	var returns ApiManagementApiDiagnosticFrontendResponseOutputReference
	_jsii_.Get(
		j,
		"frontendResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) FrontendResponseInput() *ApiManagementApiDiagnosticFrontendResponse {
	var returns *ApiManagementApiDiagnosticFrontendResponse
	_jsii_.Get(
		j,
		"frontendResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) HttpCorrelationProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorrelationProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) HttpCorrelationProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorrelationProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) LogClientIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logClientIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) LogClientIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logClientIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) OperationNameFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) OperationNameFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) SamplingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) SamplingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) Timeouts() ApiManagementApiDiagnosticTimeoutsOutputReference {
	var returns ApiManagementApiDiagnosticTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) Verbosity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verbosity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnostic) VerbosityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verbosityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/api_management_api_diagnostic azurerm_api_management_api_diagnostic} Resource.
func NewApiManagementApiDiagnostic(scope constructs.Construct, id *string, config *ApiManagementApiDiagnosticConfig) ApiManagementApiDiagnostic {
	_init_.Initialize()

	if err := validateNewApiManagementApiDiagnosticParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementApiDiagnostic{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementApiDiagnostic.ApiManagementApiDiagnostic",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/api_management_api_diagnostic azurerm_api_management_api_diagnostic} Resource.
func NewApiManagementApiDiagnostic_Override(a ApiManagementApiDiagnostic, scope constructs.Construct, id *string, config *ApiManagementApiDiagnosticConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementApiDiagnostic.ApiManagementApiDiagnostic",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetAlwaysLogErrors(val interface{}) {
	if err := j.validateSetAlwaysLogErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysLogErrors",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetApiManagementLoggerId(val *string) {
	if err := j.validateSetApiManagementLoggerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiManagementLoggerId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetApiManagementName(val *string) {
	if err := j.validateSetApiManagementNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiManagementName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetApiName(val *string) {
	if err := j.validateSetApiNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetHttpCorrelationProtocol(val *string) {
	if err := j.validateSetHttpCorrelationProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCorrelationProtocol",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetIdentifier(val *string) {
	if err := j.validateSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetLogClientIp(val interface{}) {
	if err := j.validateSetLogClientIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logClientIp",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetOperationNameFormat(val *string) {
	if err := j.validateSetOperationNameFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationNameFormat",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetSamplingPercentage(val *float64) {
	if err := j.validateSetSamplingPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplingPercentage",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnostic)SetVerbosity(val *string) {
	if err := j.validateSetVerbosityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verbosity",
		val,
	)
}

// Generates CDKTF code for importing a ApiManagementApiDiagnostic resource upon running "cdktf plan <stack-name>".
func ApiManagementApiDiagnostic_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApiManagementApiDiagnostic_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementApiDiagnostic.ApiManagementApiDiagnostic",
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
func ApiManagementApiDiagnostic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementApiDiagnostic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementApiDiagnostic.ApiManagementApiDiagnostic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagementApiDiagnostic_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementApiDiagnostic_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementApiDiagnostic.ApiManagementApiDiagnostic",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagementApiDiagnostic_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementApiDiagnostic_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementApiDiagnostic.ApiManagementApiDiagnostic",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiManagementApiDiagnostic_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.apiManagementApiDiagnostic.ApiManagementApiDiagnostic",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiManagementApiDiagnostic) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementApiDiagnostic) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiManagementApiDiagnostic) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiManagementApiDiagnostic) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiManagementApiDiagnostic) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiManagementApiDiagnostic) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiManagementApiDiagnostic) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiManagementApiDiagnostic) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiManagementApiDiagnostic) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementApiDiagnostic) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) PutBackendRequest(value *ApiManagementApiDiagnosticBackendRequest) {
	if err := a.validatePutBackendRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackendRequest",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) PutBackendResponse(value *ApiManagementApiDiagnosticBackendResponse) {
	if err := a.validatePutBackendResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackendResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) PutFrontendRequest(value *ApiManagementApiDiagnosticFrontendRequest) {
	if err := a.validatePutFrontendRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrontendRequest",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) PutFrontendResponse(value *ApiManagementApiDiagnosticFrontendResponse) {
	if err := a.validatePutFrontendResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrontendResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) PutTimeouts(value *ApiManagementApiDiagnosticTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetAlwaysLogErrors() {
	_jsii_.InvokeVoid(
		a,
		"resetAlwaysLogErrors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetBackendRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetBackendResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetFrontendRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetFrontendRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetFrontendResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetFrontendResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetHttpCorrelationProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpCorrelationProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetLogClientIp() {
	_jsii_.InvokeVoid(
		a,
		"resetLogClientIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetOperationNameFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetOperationNameFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetSamplingPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetSamplingPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ResetVerbosity() {
	_jsii_.InvokeVoid(
		a,
		"resetVerbosity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementApiDiagnostic) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

