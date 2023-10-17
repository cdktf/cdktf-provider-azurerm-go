// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicekusto

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/datafactorylinkedservicekusto/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/data_factory_linked_service_kusto azurerm_data_factory_linked_service_kusto}.
type DataFactoryLinkedServiceKusto interface {
	cdktf.TerraformResource
	AdditionalProperties() *map[string]*string
	SetAdditionalProperties(val *map[string]*string)
	AdditionalPropertiesInput() *map[string]*string
	Annotations() *[]*string
	SetAnnotations(val *[]*string)
	AnnotationsInput() *[]*string
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
	DataFactoryId() *string
	SetDataFactoryId(val *string)
	DataFactoryIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	IntegrationRuntimeName() *string
	SetIntegrationRuntimeName(val *string)
	IntegrationRuntimeNameInput() *string
	KustoDatabaseName() *string
	SetKustoDatabaseName(val *string)
	KustoDatabaseNameInput() *string
	KustoEndpoint() *string
	SetKustoEndpoint(val *string)
	KustoEndpointInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
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
	ServicePrincipalId() *string
	SetServicePrincipalId(val *string)
	ServicePrincipalIdInput() *string
	ServicePrincipalKey() *string
	SetServicePrincipalKey(val *string)
	ServicePrincipalKeyInput() *string
	Tenant() *string
	SetTenant(val *string)
	TenantInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataFactoryLinkedServiceKustoTimeoutsOutputReference
	TimeoutsInput() interface{}
	UseManagedIdentity() interface{}
	SetUseManagedIdentity(val interface{})
	UseManagedIdentityInput() interface{}
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
	PutTimeouts(value *DataFactoryLinkedServiceKustoTimeouts)
	ResetAdditionalProperties()
	ResetAnnotations()
	ResetDescription()
	ResetId()
	ResetIntegrationRuntimeName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetServicePrincipalId()
	ResetServicePrincipalKey()
	ResetTenant()
	ResetTimeouts()
	ResetUseManagedIdentity()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataFactoryLinkedServiceKusto
type jsiiProxy_DataFactoryLinkedServiceKusto struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) AdditionalProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) AdditionalPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Annotations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) AnnotationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) DataFactoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) DataFactoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) IntegrationRuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationRuntimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) IntegrationRuntimeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationRuntimeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) KustoDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) KustoDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) KustoEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) KustoEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) ServicePrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) ServicePrincipalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) ServicePrincipalKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) ServicePrincipalKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Tenant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) TenantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) Timeouts() DataFactoryLinkedServiceKustoTimeoutsOutputReference {
	var returns DataFactoryLinkedServiceKustoTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) UseManagedIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto) UseManagedIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useManagedIdentityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/data_factory_linked_service_kusto azurerm_data_factory_linked_service_kusto} Resource.
func NewDataFactoryLinkedServiceKusto(scope constructs.Construct, id *string, config *DataFactoryLinkedServiceKustoConfig) DataFactoryLinkedServiceKusto {
	_init_.Initialize()

	if err := validateNewDataFactoryLinkedServiceKustoParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryLinkedServiceKusto{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceKusto.DataFactoryLinkedServiceKusto",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/data_factory_linked_service_kusto azurerm_data_factory_linked_service_kusto} Resource.
func NewDataFactoryLinkedServiceKusto_Override(d DataFactoryLinkedServiceKusto, scope constructs.Construct, id *string, config *DataFactoryLinkedServiceKustoConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceKusto.DataFactoryLinkedServiceKusto",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetAdditionalProperties(val *map[string]*string) {
	if err := j.validateSetAdditionalPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalProperties",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetAnnotations(val *[]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetDataFactoryId(val *string) {
	if err := j.validateSetDataFactoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFactoryId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetIntegrationRuntimeName(val *string) {
	if err := j.validateSetIntegrationRuntimeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationRuntimeName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetKustoDatabaseName(val *string) {
	if err := j.validateSetKustoDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kustoDatabaseName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetKustoEndpoint(val *string) {
	if err := j.validateSetKustoEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kustoEndpoint",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetServicePrincipalId(val *string) {
	if err := j.validateSetServicePrincipalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetServicePrincipalKey(val *string) {
	if err := j.validateSetServicePrincipalKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalKey",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetTenant(val *string) {
	if err := j.validateSetTenantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenant",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceKusto)SetUseManagedIdentity(val interface{}) {
	if err := j.validateSetUseManagedIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useManagedIdentity",
		val,
	)
}

// Generates CDKTF code for importing a DataFactoryLinkedServiceKusto resource upon running "cdktf plan <stack-name>".
func DataFactoryLinkedServiceKusto_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataFactoryLinkedServiceKusto_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceKusto.DataFactoryLinkedServiceKusto",
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
func DataFactoryLinkedServiceKusto_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryLinkedServiceKusto_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceKusto.DataFactoryLinkedServiceKusto",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFactoryLinkedServiceKusto_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryLinkedServiceKusto_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceKusto.DataFactoryLinkedServiceKusto",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFactoryLinkedServiceKusto_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryLinkedServiceKusto_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceKusto.DataFactoryLinkedServiceKusto",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataFactoryLinkedServiceKusto_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceKusto.DataFactoryLinkedServiceKusto",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) PutTimeouts(value *DataFactoryLinkedServiceKustoTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetAdditionalProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetAnnotations() {
	_jsii_.InvokeVoid(
		d,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetIntegrationRuntimeName() {
	_jsii_.InvokeVoid(
		d,
		"resetIntegrationRuntimeName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetServicePrincipalId() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetServicePrincipalKey() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetTenant() {
	_jsii_.InvokeVoid(
		d,
		"resetTenant",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ResetUseManagedIdentity() {
	_jsii_.InvokeVoid(
		d,
		"resetUseManagedIdentity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceKusto) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

