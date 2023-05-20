package databricksworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/databricksworkspace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/databricks_workspace azurerm_databricks_workspace}.
type DatabricksWorkspace interface {
	cdktf.TerraformResource
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
	CustomerManagedKeyEnabled() interface{}
	SetCustomerManagedKeyEnabled(val interface{})
	CustomerManagedKeyEnabledInput() interface{}
	CustomParameters() DatabricksWorkspaceCustomParametersOutputReference
	CustomParametersInput() *DatabricksWorkspaceCustomParameters
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskEncryptionSetId() *string
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
	InfrastructureEncryptionEnabled() interface{}
	SetInfrastructureEncryptionEnabled(val interface{})
	InfrastructureEncryptionEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancerBackendAddressPoolId() *string
	SetLoadBalancerBackendAddressPoolId(val *string)
	LoadBalancerBackendAddressPoolIdInput() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagedDiskCmkKeyVaultKeyId() *string
	SetManagedDiskCmkKeyVaultKeyId(val *string)
	ManagedDiskCmkKeyVaultKeyIdInput() *string
	ManagedDiskCmkRotationToLatestVersionEnabled() interface{}
	SetManagedDiskCmkRotationToLatestVersionEnabled(val interface{})
	ManagedDiskCmkRotationToLatestVersionEnabledInput() interface{}
	ManagedDiskIdentity() DatabricksWorkspaceManagedDiskIdentityList
	ManagedResourceGroupId() *string
	ManagedResourceGroupName() *string
	SetManagedResourceGroupName(val *string)
	ManagedResourceGroupNameInput() *string
	ManagedServicesCmkKeyVaultKeyId() *string
	SetManagedServicesCmkKeyVaultKeyId(val *string)
	ManagedServicesCmkKeyVaultKeyIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkSecurityGroupRulesRequired() *string
	SetNetworkSecurityGroupRulesRequired(val *string)
	NetworkSecurityGroupRulesRequiredInput() *string
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
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	StorageAccountIdentity() DatabricksWorkspaceStorageAccountIdentityList
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DatabricksWorkspaceTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkspaceId() *string
	WorkspaceUrl() *string
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
	PutCustomParameters(value *DatabricksWorkspaceCustomParameters)
	PutTimeouts(value *DatabricksWorkspaceTimeouts)
	ResetCustomerManagedKeyEnabled()
	ResetCustomParameters()
	ResetId()
	ResetInfrastructureEncryptionEnabled()
	ResetLoadBalancerBackendAddressPoolId()
	ResetManagedDiskCmkKeyVaultKeyId()
	ResetManagedDiskCmkRotationToLatestVersionEnabled()
	ResetManagedResourceGroupName()
	ResetManagedServicesCmkKeyVaultKeyId()
	ResetNetworkSecurityGroupRulesRequired()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetTags()
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

// The jsii proxy struct for DatabricksWorkspace
type jsiiProxy_DatabricksWorkspace struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabricksWorkspace) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) CustomerManagedKeyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerManagedKeyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) CustomerManagedKeyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerManagedKeyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) CustomParameters() DatabricksWorkspaceCustomParametersOutputReference {
	var returns DatabricksWorkspaceCustomParametersOutputReference
	_jsii_.Get(
		j,
		"customParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) CustomParametersInput() *DatabricksWorkspaceCustomParameters {
	var returns *DatabricksWorkspaceCustomParameters
	_jsii_.Get(
		j,
		"customParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) DiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) InfrastructureEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infrastructureEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) InfrastructureEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infrastructureEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) LoadBalancerBackendAddressPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerBackendAddressPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) LoadBalancerBackendAddressPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerBackendAddressPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ManagedDiskCmkKeyVaultKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDiskCmkKeyVaultKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ManagedDiskCmkKeyVaultKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDiskCmkKeyVaultKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ManagedDiskCmkRotationToLatestVersionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedDiskCmkRotationToLatestVersionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ManagedDiskCmkRotationToLatestVersionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedDiskCmkRotationToLatestVersionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ManagedDiskIdentity() DatabricksWorkspaceManagedDiskIdentityList {
	var returns DatabricksWorkspaceManagedDiskIdentityList
	_jsii_.Get(
		j,
		"managedDiskIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ManagedResourceGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ManagedResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ManagedResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ManagedServicesCmkKeyVaultKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedServicesCmkKeyVaultKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ManagedServicesCmkKeyVaultKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedServicesCmkKeyVaultKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) NetworkSecurityGroupRulesRequired() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupRulesRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) NetworkSecurityGroupRulesRequiredInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupRulesRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) StorageAccountIdentity() DatabricksWorkspaceStorageAccountIdentityList {
	var returns DatabricksWorkspaceStorageAccountIdentityList
	_jsii_.Get(
		j,
		"storageAccountIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) Timeouts() DatabricksWorkspaceTimeoutsOutputReference {
	var returns DatabricksWorkspaceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspace) WorkspaceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceUrl",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/databricks_workspace azurerm_databricks_workspace} Resource.
func NewDatabricksWorkspace(scope constructs.Construct, id *string, config *DatabricksWorkspaceConfig) DatabricksWorkspace {
	_init_.Initialize()

	if err := validateNewDatabricksWorkspaceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabricksWorkspace{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/databricks_workspace azurerm_databricks_workspace} Resource.
func NewDatabricksWorkspace_Override(d DatabricksWorkspace, scope constructs.Construct, id *string, config *DatabricksWorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspace",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetCustomerManagedKeyEnabled(val interface{}) {
	if err := j.validateSetCustomerManagedKeyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerManagedKeyEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetInfrastructureEncryptionEnabled(val interface{}) {
	if err := j.validateSetInfrastructureEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetLoadBalancerBackendAddressPoolId(val *string) {
	if err := j.validateSetLoadBalancerBackendAddressPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerBackendAddressPoolId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetManagedDiskCmkKeyVaultKeyId(val *string) {
	if err := j.validateSetManagedDiskCmkKeyVaultKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedDiskCmkKeyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetManagedDiskCmkRotationToLatestVersionEnabled(val interface{}) {
	if err := j.validateSetManagedDiskCmkRotationToLatestVersionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedDiskCmkRotationToLatestVersionEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetManagedResourceGroupName(val *string) {
	if err := j.validateSetManagedResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetManagedServicesCmkKeyVaultKeyId(val *string) {
	if err := j.validateSetManagedServicesCmkKeyVaultKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedServicesCmkKeyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetNetworkSecurityGroupRulesRequired(val *string) {
	if err := j.validateSetNetworkSecurityGroupRulesRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkSecurityGroupRulesRequired",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
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
func DatabricksWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabricksWorkspace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabricksWorkspace_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabricksWorkspace_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspace",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabricksWorkspace_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabricksWorkspace_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspace",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabricksWorkspace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspace",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabricksWorkspace) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabricksWorkspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabricksWorkspace) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabricksWorkspace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabricksWorkspace) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabricksWorkspace) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabricksWorkspace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabricksWorkspace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabricksWorkspace) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabricksWorkspace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabricksWorkspace) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabricksWorkspace) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabricksWorkspace) PutCustomParameters(value *DatabricksWorkspaceCustomParameters) {
	if err := d.validatePutCustomParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomParameters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabricksWorkspace) PutTimeouts(value *DatabricksWorkspaceTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetCustomerManagedKeyEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomerManagedKeyEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetCustomParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetInfrastructureEncryptionEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetInfrastructureEncryptionEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetLoadBalancerBackendAddressPoolId() {
	_jsii_.InvokeVoid(
		d,
		"resetLoadBalancerBackendAddressPoolId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetManagedDiskCmkKeyVaultKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedDiskCmkKeyVaultKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetManagedDiskCmkRotationToLatestVersionEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedDiskCmkRotationToLatestVersionEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetManagedResourceGroupName() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedResourceGroupName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetManagedServicesCmkKeyVaultKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedServicesCmkKeyVaultKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetNetworkSecurityGroupRulesRequired() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkSecurityGroupRulesRequired",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspace) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspace) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

