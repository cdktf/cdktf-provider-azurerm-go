package datafactorylinkedserviceazuresqldatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/datafactorylinkedserviceazuresqldatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database azurerm_data_factory_linked_service_azure_sql_database}.
type DataFactoryLinkedServiceAzureSqlDatabase interface {
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
	ConnectionString() *string
	SetConnectionString(val *string)
	ConnectionStringInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	KeyVaultConnectionString() DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference
	KeyVaultConnectionStringInput() *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString
	KeyVaultPassword() DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference
	KeyVaultPasswordInput() *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword
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
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference
	TimeoutsInput() interface{}
	UseManagedIdentity() interface{}
	SetUseManagedIdentity(val interface{})
	UseManagedIdentityInput() interface{}
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
	PutKeyVaultConnectionString(value *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString)
	PutKeyVaultPassword(value *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword)
	PutTimeouts(value *DataFactoryLinkedServiceAzureSqlDatabaseTimeouts)
	ResetAdditionalProperties()
	ResetAnnotations()
	ResetConnectionString()
	ResetDescription()
	ResetId()
	ResetIntegrationRuntimeName()
	ResetKeyVaultConnectionString()
	ResetKeyVaultPassword()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetServicePrincipalId()
	ResetServicePrincipalKey()
	ResetTenantId()
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

// The jsii proxy struct for DataFactoryLinkedServiceAzureSqlDatabase
type jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) AdditionalProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) AdditionalPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Annotations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) AnnotationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) DataFactoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) DataFactoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) IntegrationRuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationRuntimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) IntegrationRuntimeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationRuntimeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) KeyVaultConnectionString() DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference {
	var returns DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference
	_jsii_.Get(
		j,
		"keyVaultConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) KeyVaultConnectionStringInput() *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString {
	var returns *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString
	_jsii_.Get(
		j,
		"keyVaultConnectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) KeyVaultPassword() DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference {
	var returns DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference
	_jsii_.Get(
		j,
		"keyVaultPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) KeyVaultPasswordInput() *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword {
	var returns *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword
	_jsii_.Get(
		j,
		"keyVaultPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ServicePrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ServicePrincipalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ServicePrincipalKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ServicePrincipalKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) Timeouts() DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference {
	var returns DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) UseManagedIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) UseManagedIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useManagedIdentityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database azurerm_data_factory_linked_service_azure_sql_database} Resource.
func NewDataFactoryLinkedServiceAzureSqlDatabase(scope constructs.Construct, id *string, config *DataFactoryLinkedServiceAzureSqlDatabaseConfig) DataFactoryLinkedServiceAzureSqlDatabase {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureSqlDatabase.DataFactoryLinkedServiceAzureSqlDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database azurerm_data_factory_linked_service_azure_sql_database} Resource.
func NewDataFactoryLinkedServiceAzureSqlDatabase_Override(d DataFactoryLinkedServiceAzureSqlDatabase, scope constructs.Construct, id *string, config *DataFactoryLinkedServiceAzureSqlDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureSqlDatabase.DataFactoryLinkedServiceAzureSqlDatabase",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetAdditionalProperties(val *map[string]*string) {
	_jsii_.Set(
		j,
		"additionalProperties",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetAnnotations(val *[]*string) {
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetConnectionString(val *string) {
	_jsii_.Set(
		j,
		"connectionString",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetDataFactoryId(val *string) {
	_jsii_.Set(
		j,
		"dataFactoryId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetIntegrationRuntimeName(val *string) {
	_jsii_.Set(
		j,
		"integrationRuntimeName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetServicePrincipalId(val *string) {
	_jsii_.Set(
		j,
		"servicePrincipalId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetServicePrincipalKey(val *string) {
	_jsii_.Set(
		j,
		"servicePrincipalKey",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SetUseManagedIdentity(val interface{}) {
	_jsii_.Set(
		j,
		"useManagedIdentity",
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
func DataFactoryLinkedServiceAzureSqlDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureSqlDatabase.DataFactoryLinkedServiceAzureSqlDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataFactoryLinkedServiceAzureSqlDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureSqlDatabase.DataFactoryLinkedServiceAzureSqlDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) PutKeyVaultConnectionString(value *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString) {
	_jsii_.InvokeVoid(
		d,
		"putKeyVaultConnectionString",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) PutKeyVaultPassword(value *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword) {
	_jsii_.InvokeVoid(
		d,
		"putKeyVaultPassword",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) PutTimeouts(value *DataFactoryLinkedServiceAzureSqlDatabaseTimeouts) {
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetAdditionalProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetAnnotations() {
	_jsii_.InvokeVoid(
		d,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetConnectionString() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetIntegrationRuntimeName() {
	_jsii_.InvokeVoid(
		d,
		"resetIntegrationRuntimeName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetKeyVaultConnectionString() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyVaultConnectionString",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetKeyVaultPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyVaultPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetServicePrincipalId() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetServicePrincipalKey() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetTenantId() {
	_jsii_.InvokeVoid(
		d,
		"resetTenantId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ResetUseManagedIdentity() {
	_jsii_.InvokeVoid(
		d,
		"resetUseManagedIdentity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryLinkedServiceAzureSqlDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#data_factory_id DataFactoryLinkedServiceAzureSqlDatabase#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#name DataFactoryLinkedServiceAzureSqlDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#additional_properties DataFactoryLinkedServiceAzureSqlDatabase#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#annotations DataFactoryLinkedServiceAzureSqlDatabase#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#connection_string DataFactoryLinkedServiceAzureSqlDatabase#connection_string}.
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#description DataFactoryLinkedServiceAzureSqlDatabase#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#id DataFactoryLinkedServiceAzureSqlDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#integration_runtime_name DataFactoryLinkedServiceAzureSqlDatabase#integration_runtime_name}.
	IntegrationRuntimeName *string `field:"optional" json:"integrationRuntimeName" yaml:"integrationRuntimeName"`
	// key_vault_connection_string block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#key_vault_connection_string DataFactoryLinkedServiceAzureSqlDatabase#key_vault_connection_string}
	KeyVaultConnectionString *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString `field:"optional" json:"keyVaultConnectionString" yaml:"keyVaultConnectionString"`
	// key_vault_password block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#key_vault_password DataFactoryLinkedServiceAzureSqlDatabase#key_vault_password}
	KeyVaultPassword *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword `field:"optional" json:"keyVaultPassword" yaml:"keyVaultPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#parameters DataFactoryLinkedServiceAzureSqlDatabase#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#service_principal_id DataFactoryLinkedServiceAzureSqlDatabase#service_principal_id}.
	ServicePrincipalId *string `field:"optional" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#service_principal_key DataFactoryLinkedServiceAzureSqlDatabase#service_principal_key}.
	ServicePrincipalKey *string `field:"optional" json:"servicePrincipalKey" yaml:"servicePrincipalKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#tenant_id DataFactoryLinkedServiceAzureSqlDatabase#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#timeouts DataFactoryLinkedServiceAzureSqlDatabase#timeouts}
	Timeouts *DataFactoryLinkedServiceAzureSqlDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#use_managed_identity DataFactoryLinkedServiceAzureSqlDatabase#use_managed_identity}.
	UseManagedIdentity interface{} `field:"optional" json:"useManagedIdentity" yaml:"useManagedIdentity"`
}

type DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#linked_service_name DataFactoryLinkedServiceAzureSqlDatabase#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#secret_name DataFactoryLinkedServiceAzureSqlDatabase#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

type DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference interface {
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
	InternalValue() *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString
	SetInternalValue(val *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString)
	LinkedServiceName() *string
	SetLinkedServiceName(val *string)
	LinkedServiceNameInput() *string
	SecretName() *string
	SetSecretName(val *string)
	SecretNameInput() *string
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

// The jsii proxy struct for DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference
type jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) InternalValue() *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString {
	var returns *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) LinkedServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) LinkedServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) SecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureSqlDatabase.DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference_Override(d DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureSqlDatabase.DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) SetInternalValue(val *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionString) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) SetLinkedServiceName(val *string) {
	_jsii_.Set(
		j,
		"linkedServiceName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) SetSecretName(val *string) {
	_jsii_.Set(
		j,
		"secretName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultConnectionStringOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#linked_service_name DataFactoryLinkedServiceAzureSqlDatabase#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#secret_name DataFactoryLinkedServiceAzureSqlDatabase#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

type DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference interface {
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
	InternalValue() *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword
	SetInternalValue(val *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword)
	LinkedServiceName() *string
	SetLinkedServiceName(val *string)
	LinkedServiceNameInput() *string
	SecretName() *string
	SetSecretName(val *string)
	SecretNameInput() *string
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

// The jsii proxy struct for DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference
type jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) InternalValue() *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword {
	var returns *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) LinkedServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) LinkedServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) SecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureSqlDatabase.DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference_Override(d DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureSqlDatabase.DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) SetInternalValue(val *DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) SetLinkedServiceName(val *string) {
	_jsii_.Set(
		j,
		"linkedServiceName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) SetSecretName(val *string) {
	_jsii_.Set(
		j,
		"secretName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPasswordOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryLinkedServiceAzureSqlDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#create DataFactoryLinkedServiceAzureSqlDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#delete DataFactoryLinkedServiceAzureSqlDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#read DataFactoryLinkedServiceAzureSqlDatabase#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_sql_database#update DataFactoryLinkedServiceAzureSqlDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference interface {
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

// The jsii proxy struct for DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference
type jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewDataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureSqlDatabase.DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference_Override(d DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureSqlDatabase.DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		d,
		"resetCreate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		d,
		"resetRead",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureSqlDatabaseTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

