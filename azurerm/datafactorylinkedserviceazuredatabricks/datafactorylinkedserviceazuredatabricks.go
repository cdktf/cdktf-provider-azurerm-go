package datafactorylinkedserviceazuredatabricks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/datafactorylinkedserviceazuredatabricks/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks azurerm_data_factory_linked_service_azure_databricks}.
type DataFactoryLinkedServiceAzureDatabricks interface {
	cdktf.TerraformResource
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	AdbDomain() *string
	SetAdbDomain(val *string)
	AdbDomainInput() *string
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
	ExistingClusterId() *string
	SetExistingClusterId(val *string)
	ExistingClusterIdInput() *string
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
	InstancePool() DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference
	InstancePoolInput() *DataFactoryLinkedServiceAzureDatabricksInstancePool
	IntegrationRuntimeName() *string
	SetIntegrationRuntimeName(val *string)
	IntegrationRuntimeNameInput() *string
	KeyVaultPassword() DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference
	KeyVaultPasswordInput() *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MsiWorkSpaceResourceId() *string
	SetMsiWorkSpaceResourceId(val *string)
	MsiWorkSpaceResourceIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NewClusterConfig() DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference
	NewClusterConfigInput() *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference
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
	PutInstancePool(value *DataFactoryLinkedServiceAzureDatabricksInstancePool)
	PutKeyVaultPassword(value *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword)
	PutNewClusterConfig(value *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig)
	PutTimeouts(value *DataFactoryLinkedServiceAzureDatabricksTimeouts)
	ResetAccessToken()
	ResetAdditionalProperties()
	ResetAnnotations()
	ResetDescription()
	ResetExistingClusterId()
	ResetId()
	ResetInstancePool()
	ResetIntegrationRuntimeName()
	ResetKeyVaultPassword()
	ResetMsiWorkSpaceResourceId()
	ResetNewClusterConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
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

// The jsii proxy struct for DataFactoryLinkedServiceAzureDatabricks
type jsiiProxy_DataFactoryLinkedServiceAzureDatabricks struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) AdbDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adbDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) AdbDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adbDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) AdditionalProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) AdditionalPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Annotations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) AnnotationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) DataFactoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) DataFactoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ExistingClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ExistingClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) InstancePool() DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference {
	var returns DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference
	_jsii_.Get(
		j,
		"instancePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) InstancePoolInput() *DataFactoryLinkedServiceAzureDatabricksInstancePool {
	var returns *DataFactoryLinkedServiceAzureDatabricksInstancePool
	_jsii_.Get(
		j,
		"instancePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) IntegrationRuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationRuntimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) IntegrationRuntimeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationRuntimeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) KeyVaultPassword() DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference {
	var returns DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference
	_jsii_.Get(
		j,
		"keyVaultPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) KeyVaultPasswordInput() *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword {
	var returns *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword
	_jsii_.Get(
		j,
		"keyVaultPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) MsiWorkSpaceResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiWorkSpaceResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) MsiWorkSpaceResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiWorkSpaceResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) NewClusterConfig() DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference {
	var returns DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference
	_jsii_.Get(
		j,
		"newClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) NewClusterConfigInput() *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig {
	var returns *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig
	_jsii_.Get(
		j,
		"newClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) Timeouts() DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference {
	var returns DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks azurerm_data_factory_linked_service_azure_databricks} Resource.
func NewDataFactoryLinkedServiceAzureDatabricks(scope constructs.Construct, id *string, config *DataFactoryLinkedServiceAzureDatabricksConfig) DataFactoryLinkedServiceAzureDatabricks {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryLinkedServiceAzureDatabricks{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricks",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks azurerm_data_factory_linked_service_azure_databricks} Resource.
func NewDataFactoryLinkedServiceAzureDatabricks_Override(d DataFactoryLinkedServiceAzureDatabricks, scope constructs.Construct, id *string, config *DataFactoryLinkedServiceAzureDatabricksConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricks",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetAccessToken(val *string) {
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetAdbDomain(val *string) {
	_jsii_.Set(
		j,
		"adbDomain",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetAdditionalProperties(val *map[string]*string) {
	_jsii_.Set(
		j,
		"additionalProperties",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetAnnotations(val *[]*string) {
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetDataFactoryId(val *string) {
	_jsii_.Set(
		j,
		"dataFactoryId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetExistingClusterId(val *string) {
	_jsii_.Set(
		j,
		"existingClusterId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetIntegrationRuntimeName(val *string) {
	_jsii_.Set(
		j,
		"integrationRuntimeName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetMsiWorkSpaceResourceId(val *string) {
	_jsii_.Set(
		j,
		"msiWorkSpaceResourceId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
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
func DataFactoryLinkedServiceAzureDatabricks_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricks",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataFactoryLinkedServiceAzureDatabricks_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricks",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) PutInstancePool(value *DataFactoryLinkedServiceAzureDatabricksInstancePool) {
	_jsii_.InvokeVoid(
		d,
		"putInstancePool",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) PutKeyVaultPassword(value *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword) {
	_jsii_.InvokeVoid(
		d,
		"putKeyVaultPassword",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) PutNewClusterConfig(value *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig) {
	_jsii_.InvokeVoid(
		d,
		"putNewClusterConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) PutTimeouts(value *DataFactoryLinkedServiceAzureDatabricksTimeouts) {
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetAccessToken() {
	_jsii_.InvokeVoid(
		d,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetAdditionalProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetAnnotations() {
	_jsii_.InvokeVoid(
		d,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetExistingClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetExistingClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetInstancePool() {
	_jsii_.InvokeVoid(
		d,
		"resetInstancePool",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetIntegrationRuntimeName() {
	_jsii_.InvokeVoid(
		d,
		"resetIntegrationRuntimeName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetKeyVaultPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyVaultPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetMsiWorkSpaceResourceId() {
	_jsii_.InvokeVoid(
		d,
		"resetMsiWorkSpaceResourceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetNewClusterConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetNewClusterConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricks) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryLinkedServiceAzureDatabricksConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#adb_domain DataFactoryLinkedServiceAzureDatabricks#adb_domain}.
	AdbDomain *string `field:"required" json:"adbDomain" yaml:"adbDomain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#data_factory_id DataFactoryLinkedServiceAzureDatabricks#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#name DataFactoryLinkedServiceAzureDatabricks#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#access_token DataFactoryLinkedServiceAzureDatabricks#access_token}.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#additional_properties DataFactoryLinkedServiceAzureDatabricks#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#annotations DataFactoryLinkedServiceAzureDatabricks#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#description DataFactoryLinkedServiceAzureDatabricks#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#existing_cluster_id DataFactoryLinkedServiceAzureDatabricks#existing_cluster_id}.
	ExistingClusterId *string `field:"optional" json:"existingClusterId" yaml:"existingClusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#id DataFactoryLinkedServiceAzureDatabricks#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instance_pool block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#instance_pool DataFactoryLinkedServiceAzureDatabricks#instance_pool}
	InstancePool *DataFactoryLinkedServiceAzureDatabricksInstancePool `field:"optional" json:"instancePool" yaml:"instancePool"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#integration_runtime_name DataFactoryLinkedServiceAzureDatabricks#integration_runtime_name}.
	IntegrationRuntimeName *string `field:"optional" json:"integrationRuntimeName" yaml:"integrationRuntimeName"`
	// key_vault_password block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#key_vault_password DataFactoryLinkedServiceAzureDatabricks#key_vault_password}
	KeyVaultPassword *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword `field:"optional" json:"keyVaultPassword" yaml:"keyVaultPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#msi_work_space_resource_id DataFactoryLinkedServiceAzureDatabricks#msi_work_space_resource_id}.
	MsiWorkSpaceResourceId *string `field:"optional" json:"msiWorkSpaceResourceId" yaml:"msiWorkSpaceResourceId"`
	// new_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#new_cluster_config DataFactoryLinkedServiceAzureDatabricks#new_cluster_config}
	NewClusterConfig *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig `field:"optional" json:"newClusterConfig" yaml:"newClusterConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#parameters DataFactoryLinkedServiceAzureDatabricks#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#timeouts DataFactoryLinkedServiceAzureDatabricks#timeouts}
	Timeouts *DataFactoryLinkedServiceAzureDatabricksTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type DataFactoryLinkedServiceAzureDatabricksInstancePool struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#cluster_version DataFactoryLinkedServiceAzureDatabricks#cluster_version}.
	ClusterVersion *string `field:"required" json:"clusterVersion" yaml:"clusterVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#instance_pool_id DataFactoryLinkedServiceAzureDatabricks#instance_pool_id}.
	InstancePoolId *string `field:"required" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#max_number_of_workers DataFactoryLinkedServiceAzureDatabricks#max_number_of_workers}.
	MaxNumberOfWorkers *float64 `field:"optional" json:"maxNumberOfWorkers" yaml:"maxNumberOfWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#min_number_of_workers DataFactoryLinkedServiceAzureDatabricks#min_number_of_workers}.
	MinNumberOfWorkers *float64 `field:"optional" json:"minNumberOfWorkers" yaml:"minNumberOfWorkers"`
}

type DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference interface {
	cdktf.ComplexObject
	ClusterVersion() *string
	SetClusterVersion(val *string)
	ClusterVersionInput() *string
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
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	InternalValue() *DataFactoryLinkedServiceAzureDatabricksInstancePool
	SetInternalValue(val *DataFactoryLinkedServiceAzureDatabricksInstancePool)
	MaxNumberOfWorkers() *float64
	SetMaxNumberOfWorkers(val *float64)
	MaxNumberOfWorkersInput() *float64
	MinNumberOfWorkers() *float64
	SetMinNumberOfWorkers(val *float64)
	MinNumberOfWorkersInput() *float64
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
	ResetMaxNumberOfWorkers()
	ResetMinNumberOfWorkers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference
type jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) ClusterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) InternalValue() *DataFactoryLinkedServiceAzureDatabricksInstancePool {
	var returns *DataFactoryLinkedServiceAzureDatabricksInstancePool
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) MaxNumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) MaxNumberOfWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumberOfWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) MinNumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) MinNumberOfWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumberOfWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference_Override(d DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) SetClusterVersion(val *string) {
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) SetInstancePoolId(val *string) {
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) SetInternalValue(val *DataFactoryLinkedServiceAzureDatabricksInstancePool) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) SetMaxNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"maxNumberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) SetMinNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"minNumberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) ResetMaxNumberOfWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxNumberOfWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) ResetMinNumberOfWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetMinNumberOfWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksInstancePoolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#linked_service_name DataFactoryLinkedServiceAzureDatabricks#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#secret_name DataFactoryLinkedServiceAzureDatabricks#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

type DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference interface {
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
	InternalValue() *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword
	SetInternalValue(val *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword)
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

// The jsii proxy struct for DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference
type jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) InternalValue() *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword {
	var returns *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) LinkedServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) LinkedServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) SecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference_Override(d DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) SetInternalValue(val *DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) SetLinkedServiceName(val *string) {
	_jsii_.Set(
		j,
		"linkedServiceName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) SetSecretName(val *string) {
	_jsii_.Set(
		j,
		"secretName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksKeyVaultPasswordOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryLinkedServiceAzureDatabricksNewClusterConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#cluster_version DataFactoryLinkedServiceAzureDatabricks#cluster_version}.
	ClusterVersion *string `field:"required" json:"clusterVersion" yaml:"clusterVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#node_type DataFactoryLinkedServiceAzureDatabricks#node_type}.
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#custom_tags DataFactoryLinkedServiceAzureDatabricks#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#driver_node_type DataFactoryLinkedServiceAzureDatabricks#driver_node_type}.
	DriverNodeType *string `field:"optional" json:"driverNodeType" yaml:"driverNodeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#init_scripts DataFactoryLinkedServiceAzureDatabricks#init_scripts}.
	InitScripts *[]*string `field:"optional" json:"initScripts" yaml:"initScripts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#log_destination DataFactoryLinkedServiceAzureDatabricks#log_destination}.
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#max_number_of_workers DataFactoryLinkedServiceAzureDatabricks#max_number_of_workers}.
	MaxNumberOfWorkers *float64 `field:"optional" json:"maxNumberOfWorkers" yaml:"maxNumberOfWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#min_number_of_workers DataFactoryLinkedServiceAzureDatabricks#min_number_of_workers}.
	MinNumberOfWorkers *float64 `field:"optional" json:"minNumberOfWorkers" yaml:"minNumberOfWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#spark_config DataFactoryLinkedServiceAzureDatabricks#spark_config}.
	SparkConfig *map[string]*string `field:"optional" json:"sparkConfig" yaml:"sparkConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#spark_environment_variables DataFactoryLinkedServiceAzureDatabricks#spark_environment_variables}.
	SparkEnvironmentVariables *map[string]*string `field:"optional" json:"sparkEnvironmentVariables" yaml:"sparkEnvironmentVariables"`
}

type DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference interface {
	cdktf.ComplexObject
	ClusterVersion() *string
	SetClusterVersion(val *string)
	ClusterVersionInput() *string
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
	CustomTags() *map[string]*string
	SetCustomTags(val *map[string]*string)
	CustomTagsInput() *map[string]*string
	DriverNodeType() *string
	SetDriverNodeType(val *string)
	DriverNodeTypeInput() *string
	// Experimental.
	Fqn() *string
	InitScripts() *[]*string
	SetInitScripts(val *[]*string)
	InitScriptsInput() *[]*string
	InternalValue() *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig
	SetInternalValue(val *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig)
	LogDestination() *string
	SetLogDestination(val *string)
	LogDestinationInput() *string
	MaxNumberOfWorkers() *float64
	SetMaxNumberOfWorkers(val *float64)
	MaxNumberOfWorkersInput() *float64
	MinNumberOfWorkers() *float64
	SetMinNumberOfWorkers(val *float64)
	MinNumberOfWorkersInput() *float64
	NodeType() *string
	SetNodeType(val *string)
	NodeTypeInput() *string
	SparkConfig() *map[string]*string
	SetSparkConfig(val *map[string]*string)
	SparkConfigInput() *map[string]*string
	SparkEnvironmentVariables() *map[string]*string
	SetSparkEnvironmentVariables(val *map[string]*string)
	SparkEnvironmentVariablesInput() *map[string]*string
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
	ResetCustomTags()
	ResetDriverNodeType()
	ResetInitScripts()
	ResetLogDestination()
	ResetMaxNumberOfWorkers()
	ResetMinNumberOfWorkers()
	ResetSparkConfig()
	ResetSparkEnvironmentVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference
type jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ClusterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) DriverNodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) DriverNodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) InitScripts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"initScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) InitScriptsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"initScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) InternalValue() *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig {
	var returns *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) LogDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) LogDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) MaxNumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) MaxNumberOfWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumberOfWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) MinNumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) MinNumberOfWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumberOfWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SparkConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SparkConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SparkEnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SparkEnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference_Override(d DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetClusterVersion(val *string) {
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetCustomTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetDriverNodeType(val *string) {
	_jsii_.Set(
		j,
		"driverNodeType",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetInitScripts(val *[]*string) {
	_jsii_.Set(
		j,
		"initScripts",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetInternalValue(val *DataFactoryLinkedServiceAzureDatabricksNewClusterConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetLogDestination(val *string) {
	_jsii_.Set(
		j,
		"logDestination",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetMaxNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"maxNumberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetMinNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"minNumberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetNodeType(val *string) {
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetSparkConfig(val *map[string]*string) {
	_jsii_.Set(
		j,
		"sparkConfig",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetSparkEnvironmentVariables(val *map[string]*string) {
	_jsii_.Set(
		j,
		"sparkEnvironmentVariables",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ResetDriverNodeType() {
	_jsii_.InvokeVoid(
		d,
		"resetDriverNodeType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ResetInitScripts() {
	_jsii_.InvokeVoid(
		d,
		"resetInitScripts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ResetLogDestination() {
	_jsii_.InvokeVoid(
		d,
		"resetLogDestination",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ResetMaxNumberOfWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxNumberOfWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ResetMinNumberOfWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetMinNumberOfWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ResetSparkConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ResetSparkEnvironmentVariables() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkEnvironmentVariables",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksNewClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryLinkedServiceAzureDatabricksTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#create DataFactoryLinkedServiceAzureDatabricks#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#delete DataFactoryLinkedServiceAzureDatabricks#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#read DataFactoryLinkedServiceAzureDatabricks#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_databricks#update DataFactoryLinkedServiceAzureDatabricks#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference interface {
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

// The jsii proxy struct for DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference
type jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewDataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference_Override(d DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureDatabricks.DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		d,
		"resetCreate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		d,
		"resetRead",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureDatabricksTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

