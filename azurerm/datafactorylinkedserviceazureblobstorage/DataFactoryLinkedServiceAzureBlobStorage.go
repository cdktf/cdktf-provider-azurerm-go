package datafactorylinkedserviceazureblobstorage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/datafactorylinkedserviceazureblobstorage/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/data_factory_linked_service_azure_blob_storage azurerm_data_factory_linked_service_azure_blob_storage}.
type DataFactoryLinkedServiceAzureBlobStorage interface {
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
	ConnectionStringInsecure() *string
	SetConnectionStringInsecure(val *string)
	ConnectionStringInsecureInput() *string
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
	KeyVaultSasToken() DataFactoryLinkedServiceAzureBlobStorageKeyVaultSasTokenOutputReference
	KeyVaultSasTokenInput() *DataFactoryLinkedServiceAzureBlobStorageKeyVaultSasToken
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
	SasUri() *string
	SetSasUri(val *string)
	SasUriInput() *string
	ServiceEndpoint() *string
	SetServiceEndpoint(val *string)
	ServiceEndpointInput() *string
	ServicePrincipalId() *string
	SetServicePrincipalId(val *string)
	ServicePrincipalIdInput() *string
	ServicePrincipalKey() *string
	SetServicePrincipalKey(val *string)
	ServicePrincipalKeyInput() *string
	ServicePrincipalLinkedKeyVaultKey() DataFactoryLinkedServiceAzureBlobStorageServicePrincipalLinkedKeyVaultKeyOutputReference
	ServicePrincipalLinkedKeyVaultKeyInput() *DataFactoryLinkedServiceAzureBlobStorageServicePrincipalLinkedKeyVaultKey
	StorageKind() *string
	SetStorageKind(val *string)
	StorageKindInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataFactoryLinkedServiceAzureBlobStorageTimeoutsOutputReference
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
	PutKeyVaultSasToken(value *DataFactoryLinkedServiceAzureBlobStorageKeyVaultSasToken)
	PutServicePrincipalLinkedKeyVaultKey(value *DataFactoryLinkedServiceAzureBlobStorageServicePrincipalLinkedKeyVaultKey)
	PutTimeouts(value *DataFactoryLinkedServiceAzureBlobStorageTimeouts)
	ResetAdditionalProperties()
	ResetAnnotations()
	ResetConnectionString()
	ResetConnectionStringInsecure()
	ResetDescription()
	ResetId()
	ResetIntegrationRuntimeName()
	ResetKeyVaultSasToken()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetSasUri()
	ResetServiceEndpoint()
	ResetServicePrincipalId()
	ResetServicePrincipalKey()
	ResetServicePrincipalLinkedKeyVaultKey()
	ResetStorageKind()
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

// The jsii proxy struct for DataFactoryLinkedServiceAzureBlobStorage
type jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) AdditionalProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) AdditionalPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Annotations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) AnnotationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ConnectionStringInsecure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStringInsecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ConnectionStringInsecureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStringInsecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) DataFactoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) DataFactoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) IntegrationRuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationRuntimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) IntegrationRuntimeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationRuntimeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) KeyVaultSasToken() DataFactoryLinkedServiceAzureBlobStorageKeyVaultSasTokenOutputReference {
	var returns DataFactoryLinkedServiceAzureBlobStorageKeyVaultSasTokenOutputReference
	_jsii_.Get(
		j,
		"keyVaultSasToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) KeyVaultSasTokenInput() *DataFactoryLinkedServiceAzureBlobStorageKeyVaultSasToken {
	var returns *DataFactoryLinkedServiceAzureBlobStorageKeyVaultSasToken
	_jsii_.Get(
		j,
		"keyVaultSasTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) SasUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) SasUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ServiceEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ServiceEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ServicePrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ServicePrincipalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ServicePrincipalKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ServicePrincipalKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ServicePrincipalLinkedKeyVaultKey() DataFactoryLinkedServiceAzureBlobStorageServicePrincipalLinkedKeyVaultKeyOutputReference {
	var returns DataFactoryLinkedServiceAzureBlobStorageServicePrincipalLinkedKeyVaultKeyOutputReference
	_jsii_.Get(
		j,
		"servicePrincipalLinkedKeyVaultKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ServicePrincipalLinkedKeyVaultKeyInput() *DataFactoryLinkedServiceAzureBlobStorageServicePrincipalLinkedKeyVaultKey {
	var returns *DataFactoryLinkedServiceAzureBlobStorageServicePrincipalLinkedKeyVaultKey
	_jsii_.Get(
		j,
		"servicePrincipalLinkedKeyVaultKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) StorageKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) StorageKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) Timeouts() DataFactoryLinkedServiceAzureBlobStorageTimeoutsOutputReference {
	var returns DataFactoryLinkedServiceAzureBlobStorageTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) UseManagedIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) UseManagedIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useManagedIdentityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/data_factory_linked_service_azure_blob_storage azurerm_data_factory_linked_service_azure_blob_storage} Resource.
func NewDataFactoryLinkedServiceAzureBlobStorage(scope constructs.Construct, id *string, config *DataFactoryLinkedServiceAzureBlobStorageConfig) DataFactoryLinkedServiceAzureBlobStorage {
	_init_.Initialize()

	if err := validateNewDataFactoryLinkedServiceAzureBlobStorageParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureBlobStorage.DataFactoryLinkedServiceAzureBlobStorage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/data_factory_linked_service_azure_blob_storage azurerm_data_factory_linked_service_azure_blob_storage} Resource.
func NewDataFactoryLinkedServiceAzureBlobStorage_Override(d DataFactoryLinkedServiceAzureBlobStorage, scope constructs.Construct, id *string, config *DataFactoryLinkedServiceAzureBlobStorageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureBlobStorage.DataFactoryLinkedServiceAzureBlobStorage",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetAdditionalProperties(val *map[string]*string) {
	if err := j.validateSetAdditionalPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalProperties",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetAnnotations(val *[]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetConnectionString(val *string) {
	if err := j.validateSetConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionString",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetConnectionStringInsecure(val *string) {
	if err := j.validateSetConnectionStringInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionStringInsecure",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetDataFactoryId(val *string) {
	if err := j.validateSetDataFactoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFactoryId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetIntegrationRuntimeName(val *string) {
	if err := j.validateSetIntegrationRuntimeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationRuntimeName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetSasUri(val *string) {
	if err := j.validateSetSasUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sasUri",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetServiceEndpoint(val *string) {
	if err := j.validateSetServiceEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceEndpoint",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetServicePrincipalId(val *string) {
	if err := j.validateSetServicePrincipalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetServicePrincipalKey(val *string) {
	if err := j.validateSetServicePrincipalKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalKey",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetStorageKind(val *string) {
	if err := j.validateSetStorageKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageKind",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage)SetUseManagedIdentity(val interface{}) {
	if err := j.validateSetUseManagedIdentityParameters(val); err != nil {
		panic(err)
	}
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
func DataFactoryLinkedServiceAzureBlobStorage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryLinkedServiceAzureBlobStorage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureBlobStorage.DataFactoryLinkedServiceAzureBlobStorage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFactoryLinkedServiceAzureBlobStorage_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryLinkedServiceAzureBlobStorage_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureBlobStorage.DataFactoryLinkedServiceAzureBlobStorage",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFactoryLinkedServiceAzureBlobStorage_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryLinkedServiceAzureBlobStorage_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureBlobStorage.DataFactoryLinkedServiceAzureBlobStorage",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataFactoryLinkedServiceAzureBlobStorage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataFactoryLinkedServiceAzureBlobStorage.DataFactoryLinkedServiceAzureBlobStorage",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) PutKeyVaultSasToken(value *DataFactoryLinkedServiceAzureBlobStorageKeyVaultSasToken) {
	if err := d.validatePutKeyVaultSasTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKeyVaultSasToken",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) PutServicePrincipalLinkedKeyVaultKey(value *DataFactoryLinkedServiceAzureBlobStorageServicePrincipalLinkedKeyVaultKey) {
	if err := d.validatePutServicePrincipalLinkedKeyVaultKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServicePrincipalLinkedKeyVaultKey",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) PutTimeouts(value *DataFactoryLinkedServiceAzureBlobStorageTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetAdditionalProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetAnnotations() {
	_jsii_.InvokeVoid(
		d,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetConnectionString() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionString",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetConnectionStringInsecure() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionStringInsecure",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetIntegrationRuntimeName() {
	_jsii_.InvokeVoid(
		d,
		"resetIntegrationRuntimeName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetKeyVaultSasToken() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyVaultSasToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetSasUri() {
	_jsii_.InvokeVoid(
		d,
		"resetSasUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetServiceEndpoint() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceEndpoint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetServicePrincipalId() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetServicePrincipalKey() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetServicePrincipalLinkedKeyVaultKey() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalLinkedKeyVaultKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetStorageKind() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageKind",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetTenantId() {
	_jsii_.InvokeVoid(
		d,
		"resetTenantId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ResetUseManagedIdentity() {
	_jsii_.InvokeVoid(
		d,
		"resetUseManagedIdentity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryLinkedServiceAzureBlobStorage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

