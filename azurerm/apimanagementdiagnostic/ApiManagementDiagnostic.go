// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementdiagnostic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/apimanagementdiagnostic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_diagnostic azurerm_api_management_diagnostic}.
type ApiManagementDiagnostic interface {
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
	BackendRequest() ApiManagementDiagnosticBackendRequestOutputReference
	BackendRequestInput() *ApiManagementDiagnosticBackendRequest
	BackendResponse() ApiManagementDiagnosticBackendResponseOutputReference
	BackendResponseInput() *ApiManagementDiagnosticBackendResponse
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
	FrontendRequest() ApiManagementDiagnosticFrontendRequestOutputReference
	FrontendRequestInput() *ApiManagementDiagnosticFrontendRequest
	FrontendResponse() ApiManagementDiagnosticFrontendResponseOutputReference
	FrontendResponseInput() *ApiManagementDiagnosticFrontendResponse
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
	Timeouts() ApiManagementDiagnosticTimeoutsOutputReference
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
	PutBackendRequest(value *ApiManagementDiagnosticBackendRequest)
	PutBackendResponse(value *ApiManagementDiagnosticBackendResponse)
	PutFrontendRequest(value *ApiManagementDiagnosticFrontendRequest)
	PutFrontendResponse(value *ApiManagementDiagnosticFrontendResponse)
	PutTimeouts(value *ApiManagementDiagnosticTimeouts)
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

// The jsii proxy struct for ApiManagementDiagnostic
type jsiiProxy_ApiManagementDiagnostic struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiManagementDiagnostic) AlwaysLogErrors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysLogErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) AlwaysLogErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysLogErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) ApiManagementLoggerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementLoggerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) ApiManagementLoggerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementLoggerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) ApiManagementName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) ApiManagementNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) BackendRequest() ApiManagementDiagnosticBackendRequestOutputReference {
	var returns ApiManagementDiagnosticBackendRequestOutputReference
	_jsii_.Get(
		j,
		"backendRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) BackendRequestInput() *ApiManagementDiagnosticBackendRequest {
	var returns *ApiManagementDiagnosticBackendRequest
	_jsii_.Get(
		j,
		"backendRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) BackendResponse() ApiManagementDiagnosticBackendResponseOutputReference {
	var returns ApiManagementDiagnosticBackendResponseOutputReference
	_jsii_.Get(
		j,
		"backendResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) BackendResponseInput() *ApiManagementDiagnosticBackendResponse {
	var returns *ApiManagementDiagnosticBackendResponse
	_jsii_.Get(
		j,
		"backendResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) FrontendRequest() ApiManagementDiagnosticFrontendRequestOutputReference {
	var returns ApiManagementDiagnosticFrontendRequestOutputReference
	_jsii_.Get(
		j,
		"frontendRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) FrontendRequestInput() *ApiManagementDiagnosticFrontendRequest {
	var returns *ApiManagementDiagnosticFrontendRequest
	_jsii_.Get(
		j,
		"frontendRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) FrontendResponse() ApiManagementDiagnosticFrontendResponseOutputReference {
	var returns ApiManagementDiagnosticFrontendResponseOutputReference
	_jsii_.Get(
		j,
		"frontendResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) FrontendResponseInput() *ApiManagementDiagnosticFrontendResponse {
	var returns *ApiManagementDiagnosticFrontendResponse
	_jsii_.Get(
		j,
		"frontendResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) HttpCorrelationProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorrelationProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) HttpCorrelationProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCorrelationProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) LogClientIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logClientIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) LogClientIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logClientIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) OperationNameFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) OperationNameFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) SamplingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) SamplingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) Timeouts() ApiManagementDiagnosticTimeoutsOutputReference {
	var returns ApiManagementDiagnosticTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) Verbosity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verbosity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnostic) VerbosityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verbosityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_diagnostic azurerm_api_management_diagnostic} Resource.
func NewApiManagementDiagnostic(scope constructs.Construct, id *string, config *ApiManagementDiagnosticConfig) ApiManagementDiagnostic {
	_init_.Initialize()

	if err := validateNewApiManagementDiagnosticParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementDiagnostic{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementDiagnostic.ApiManagementDiagnostic",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_diagnostic azurerm_api_management_diagnostic} Resource.
func NewApiManagementDiagnostic_Override(a ApiManagementDiagnostic, scope constructs.Construct, id *string, config *ApiManagementDiagnosticConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementDiagnostic.ApiManagementDiagnostic",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetAlwaysLogErrors(val interface{}) {
	if err := j.validateSetAlwaysLogErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysLogErrors",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetApiManagementLoggerId(val *string) {
	if err := j.validateSetApiManagementLoggerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiManagementLoggerId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetApiManagementName(val *string) {
	if err := j.validateSetApiManagementNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiManagementName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetHttpCorrelationProtocol(val *string) {
	if err := j.validateSetHttpCorrelationProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCorrelationProtocol",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetIdentifier(val *string) {
	if err := j.validateSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetLogClientIp(val interface{}) {
	if err := j.validateSetLogClientIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logClientIp",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetOperationNameFormat(val *string) {
	if err := j.validateSetOperationNameFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationNameFormat",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetSamplingPercentage(val *float64) {
	if err := j.validateSetSamplingPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplingPercentage",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnostic)SetVerbosity(val *string) {
	if err := j.validateSetVerbosityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verbosity",
		val,
	)
}

// Generates CDKTF code for importing a ApiManagementDiagnostic resource upon running "cdktf plan <stack-name>".
func ApiManagementDiagnostic_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApiManagementDiagnostic_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementDiagnostic.ApiManagementDiagnostic",
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
func ApiManagementDiagnostic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementDiagnostic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementDiagnostic.ApiManagementDiagnostic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagementDiagnostic_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementDiagnostic_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementDiagnostic.ApiManagementDiagnostic",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagementDiagnostic_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagementDiagnostic_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagementDiagnostic.ApiManagementDiagnostic",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiManagementDiagnostic_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.apiManagementDiagnostic.ApiManagementDiagnostic",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiManagementDiagnostic) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiManagementDiagnostic) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementDiagnostic) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiManagementDiagnostic) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiManagementDiagnostic) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiManagementDiagnostic) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiManagementDiagnostic) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiManagementDiagnostic) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiManagementDiagnostic) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiManagementDiagnostic) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnostic) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementDiagnostic) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) PutBackendRequest(value *ApiManagementDiagnosticBackendRequest) {
	if err := a.validatePutBackendRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackendRequest",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) PutBackendResponse(value *ApiManagementDiagnosticBackendResponse) {
	if err := a.validatePutBackendResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackendResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) PutFrontendRequest(value *ApiManagementDiagnosticFrontendRequest) {
	if err := a.validatePutFrontendRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrontendRequest",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) PutFrontendResponse(value *ApiManagementDiagnosticFrontendResponse) {
	if err := a.validatePutFrontendResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrontendResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) PutTimeouts(value *ApiManagementDiagnosticTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetAlwaysLogErrors() {
	_jsii_.InvokeVoid(
		a,
		"resetAlwaysLogErrors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetBackendRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetBackendResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetFrontendRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetFrontendRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetFrontendResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetFrontendResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetHttpCorrelationProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpCorrelationProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetLogClientIp() {
	_jsii_.InvokeVoid(
		a,
		"resetLogClientIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetOperationNameFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetOperationNameFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetSamplingPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetSamplingPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) ResetVerbosity() {
	_jsii_.InvokeVoid(
		a,
		"resetVerbosity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnostic) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnostic) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnostic) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnostic) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnostic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnostic) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

