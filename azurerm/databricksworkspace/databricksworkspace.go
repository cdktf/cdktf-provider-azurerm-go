package databricksworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/databricksworkspace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace azurerm_databricks_workspace}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CustomerManagedKeyEnabled() interface{}
	SetCustomerManagedKeyEnabled(val interface{})
	CustomerManagedKeyEnabledInput() interface{}
	CustomParameters() DatabricksWorkspaceCustomParametersOutputReference
	CustomParametersInput() *DatabricksWorkspaceCustomParameters
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

func (j *jsiiProxy_DatabricksWorkspace) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace azurerm_databricks_workspace} Resource.
func NewDatabricksWorkspace(scope constructs.Construct, id *string, config *DatabricksWorkspaceConfig) DatabricksWorkspace {
	_init_.Initialize()

	j := jsiiProxy_DatabricksWorkspace{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace azurerm_databricks_workspace} Resource.
func NewDatabricksWorkspace_Override(d DatabricksWorkspace, scope constructs.Construct, id *string, config *DatabricksWorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspace",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetCustomerManagedKeyEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"customerManagedKeyEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetInfrastructureEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"infrastructureEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetLoadBalancerBackendAddressPoolId(val *string) {
	_jsii_.Set(
		j,
		"loadBalancerBackendAddressPoolId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetManagedResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"managedResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetManagedServicesCmkKeyVaultKeyId(val *string) {
	_jsii_.Set(
		j,
		"managedServicesCmkKeyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetNetworkSecurityGroupRulesRequired(val *string) {
	_jsii_.Set(
		j,
		"networkSecurityGroupRulesRequired",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetPublicNetworkAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetSku(val *string) {
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspace) SetTags(val *map[string]*string) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspace",
		"isConstruct",
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
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabricksWorkspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabricksWorkspace) PutCustomParameters(value *DatabricksWorkspaceCustomParameters) {
	_jsii_.InvokeVoid(
		d,
		"putCustomParameters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabricksWorkspace) PutTimeouts(value *DatabricksWorkspaceTimeouts) {
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

type DatabricksWorkspaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#location DatabricksWorkspace#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#name DatabricksWorkspace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#resource_group_name DatabricksWorkspace#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#sku DatabricksWorkspace#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#customer_managed_key_enabled DatabricksWorkspace#customer_managed_key_enabled}.
	CustomerManagedKeyEnabled interface{} `field:"optional" json:"customerManagedKeyEnabled" yaml:"customerManagedKeyEnabled"`
	// custom_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#custom_parameters DatabricksWorkspace#custom_parameters}
	CustomParameters *DatabricksWorkspaceCustomParameters `field:"optional" json:"customParameters" yaml:"customParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#id DatabricksWorkspace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#infrastructure_encryption_enabled DatabricksWorkspace#infrastructure_encryption_enabled}.
	InfrastructureEncryptionEnabled interface{} `field:"optional" json:"infrastructureEncryptionEnabled" yaml:"infrastructureEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#load_balancer_backend_address_pool_id DatabricksWorkspace#load_balancer_backend_address_pool_id}.
	LoadBalancerBackendAddressPoolId *string `field:"optional" json:"loadBalancerBackendAddressPoolId" yaml:"loadBalancerBackendAddressPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#managed_resource_group_name DatabricksWorkspace#managed_resource_group_name}.
	ManagedResourceGroupName *string `field:"optional" json:"managedResourceGroupName" yaml:"managedResourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#managed_services_cmk_key_vault_key_id DatabricksWorkspace#managed_services_cmk_key_vault_key_id}.
	ManagedServicesCmkKeyVaultKeyId *string `field:"optional" json:"managedServicesCmkKeyVaultKeyId" yaml:"managedServicesCmkKeyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#network_security_group_rules_required DatabricksWorkspace#network_security_group_rules_required}.
	NetworkSecurityGroupRulesRequired *string `field:"optional" json:"networkSecurityGroupRulesRequired" yaml:"networkSecurityGroupRulesRequired"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#public_network_access_enabled DatabricksWorkspace#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#tags DatabricksWorkspace#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#timeouts DatabricksWorkspace#timeouts}
	Timeouts *DatabricksWorkspaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type DatabricksWorkspaceCustomParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#machine_learning_workspace_id DatabricksWorkspace#machine_learning_workspace_id}.
	MachineLearningWorkspaceId *string `field:"optional" json:"machineLearningWorkspaceId" yaml:"machineLearningWorkspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#nat_gateway_name DatabricksWorkspace#nat_gateway_name}.
	NatGatewayName *string `field:"optional" json:"natGatewayName" yaml:"natGatewayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#no_public_ip DatabricksWorkspace#no_public_ip}.
	NoPublicIp interface{} `field:"optional" json:"noPublicIp" yaml:"noPublicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#private_subnet_name DatabricksWorkspace#private_subnet_name}.
	PrivateSubnetName *string `field:"optional" json:"privateSubnetName" yaml:"privateSubnetName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#private_subnet_network_security_group_association_id DatabricksWorkspace#private_subnet_network_security_group_association_id}.
	PrivateSubnetNetworkSecurityGroupAssociationId *string `field:"optional" json:"privateSubnetNetworkSecurityGroupAssociationId" yaml:"privateSubnetNetworkSecurityGroupAssociationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#public_ip_name DatabricksWorkspace#public_ip_name}.
	PublicIpName *string `field:"optional" json:"publicIpName" yaml:"publicIpName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#public_subnet_name DatabricksWorkspace#public_subnet_name}.
	PublicSubnetName *string `field:"optional" json:"publicSubnetName" yaml:"publicSubnetName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#public_subnet_network_security_group_association_id DatabricksWorkspace#public_subnet_network_security_group_association_id}.
	PublicSubnetNetworkSecurityGroupAssociationId *string `field:"optional" json:"publicSubnetNetworkSecurityGroupAssociationId" yaml:"publicSubnetNetworkSecurityGroupAssociationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#storage_account_name DatabricksWorkspace#storage_account_name}.
	StorageAccountName *string `field:"optional" json:"storageAccountName" yaml:"storageAccountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#storage_account_sku_name DatabricksWorkspace#storage_account_sku_name}.
	StorageAccountSkuName *string `field:"optional" json:"storageAccountSkuName" yaml:"storageAccountSkuName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#virtual_network_id DatabricksWorkspace#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#vnet_address_prefix DatabricksWorkspace#vnet_address_prefix}.
	VnetAddressPrefix *string `field:"optional" json:"vnetAddressPrefix" yaml:"vnetAddressPrefix"`
}

type DatabricksWorkspaceCustomParametersOutputReference interface {
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
	InternalValue() *DatabricksWorkspaceCustomParameters
	SetInternalValue(val *DatabricksWorkspaceCustomParameters)
	MachineLearningWorkspaceId() *string
	SetMachineLearningWorkspaceId(val *string)
	MachineLearningWorkspaceIdInput() *string
	NatGatewayName() *string
	SetNatGatewayName(val *string)
	NatGatewayNameInput() *string
	NoPublicIp() interface{}
	SetNoPublicIp(val interface{})
	NoPublicIpInput() interface{}
	PrivateSubnetName() *string
	SetPrivateSubnetName(val *string)
	PrivateSubnetNameInput() *string
	PrivateSubnetNetworkSecurityGroupAssociationId() *string
	SetPrivateSubnetNetworkSecurityGroupAssociationId(val *string)
	PrivateSubnetNetworkSecurityGroupAssociationIdInput() *string
	PublicIpName() *string
	SetPublicIpName(val *string)
	PublicIpNameInput() *string
	PublicSubnetName() *string
	SetPublicSubnetName(val *string)
	PublicSubnetNameInput() *string
	PublicSubnetNetworkSecurityGroupAssociationId() *string
	SetPublicSubnetNetworkSecurityGroupAssociationId(val *string)
	PublicSubnetNetworkSecurityGroupAssociationIdInput() *string
	StorageAccountName() *string
	SetStorageAccountName(val *string)
	StorageAccountNameInput() *string
	StorageAccountSkuName() *string
	SetStorageAccountSkuName(val *string)
	StorageAccountSkuNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
	VnetAddressPrefix() *string
	SetVnetAddressPrefix(val *string)
	VnetAddressPrefixInput() *string
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
	ResetMachineLearningWorkspaceId()
	ResetNatGatewayName()
	ResetNoPublicIp()
	ResetPrivateSubnetName()
	ResetPrivateSubnetNetworkSecurityGroupAssociationId()
	ResetPublicIpName()
	ResetPublicSubnetName()
	ResetPublicSubnetNetworkSecurityGroupAssociationId()
	ResetStorageAccountName()
	ResetStorageAccountSkuName()
	ResetVirtualNetworkId()
	ResetVnetAddressPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabricksWorkspaceCustomParametersOutputReference
type jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) InternalValue() *DatabricksWorkspaceCustomParameters {
	var returns *DatabricksWorkspaceCustomParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) MachineLearningWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineLearningWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) MachineLearningWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineLearningWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) NatGatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) NatGatewayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) NoPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) NoPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PrivateSubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PrivateSubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PrivateSubnetNetworkSecurityGroupAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetNetworkSecurityGroupAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PrivateSubnetNetworkSecurityGroupAssociationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetNetworkSecurityGroupAssociationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicIpName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicIpNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicSubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicSubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicSubnetNetworkSecurityGroupAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetNetworkSecurityGroupAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicSubnetNetworkSecurityGroupAssociationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetNetworkSecurityGroupAssociationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) StorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) StorageAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) StorageAccountSkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountSkuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) StorageAccountSkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountSkuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) VnetAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) VnetAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetAddressPrefixInput",
		&returns,
	)
	return returns
}


func NewDatabricksWorkspaceCustomParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabricksWorkspaceCustomParametersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceCustomParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabricksWorkspaceCustomParametersOutputReference_Override(d DatabricksWorkspaceCustomParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceCustomParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetInternalValue(val *DatabricksWorkspaceCustomParameters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetMachineLearningWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"machineLearningWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetNatGatewayName(val *string) {
	_jsii_.Set(
		j,
		"natGatewayName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetNoPublicIp(val interface{}) {
	_jsii_.Set(
		j,
		"noPublicIp",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetPrivateSubnetName(val *string) {
	_jsii_.Set(
		j,
		"privateSubnetName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetPrivateSubnetNetworkSecurityGroupAssociationId(val *string) {
	_jsii_.Set(
		j,
		"privateSubnetNetworkSecurityGroupAssociationId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetPublicIpName(val *string) {
	_jsii_.Set(
		j,
		"publicIpName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetPublicSubnetName(val *string) {
	_jsii_.Set(
		j,
		"publicSubnetName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetPublicSubnetNetworkSecurityGroupAssociationId(val *string) {
	_jsii_.Set(
		j,
		"publicSubnetNetworkSecurityGroupAssociationId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetStorageAccountName(val *string) {
	_jsii_.Set(
		j,
		"storageAccountName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetStorageAccountSkuName(val *string) {
	_jsii_.Set(
		j,
		"storageAccountSkuName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetVirtualNetworkId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) SetVnetAddressPrefix(val *string) {
	_jsii_.Set(
		j,
		"vnetAddressPrefix",
		val,
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetMachineLearningWorkspaceId() {
	_jsii_.InvokeVoid(
		d,
		"resetMachineLearningWorkspaceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetNatGatewayName() {
	_jsii_.InvokeVoid(
		d,
		"resetNatGatewayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetNoPublicIp() {
	_jsii_.InvokeVoid(
		d,
		"resetNoPublicIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetPrivateSubnetName() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateSubnetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetPrivateSubnetNetworkSecurityGroupAssociationId() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateSubnetNetworkSecurityGroupAssociationId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetPublicIpName() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicIpName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetPublicSubnetName() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicSubnetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetPublicSubnetNetworkSecurityGroupAssociationId() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicSubnetNetworkSecurityGroupAssociationId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetStorageAccountName() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageAccountName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetStorageAccountSkuName() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageAccountSkuName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		d,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetVnetAddressPrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetVnetAddressPrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DatabricksWorkspaceStorageAccountIdentity struct {
}

type DatabricksWorkspaceStorageAccountIdentityList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DatabricksWorkspaceStorageAccountIdentityOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabricksWorkspaceStorageAccountIdentityList
type jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDatabricksWorkspaceStorageAccountIdentityList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DatabricksWorkspaceStorageAccountIdentityList {
	_init_.Initialize()

	j := jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceStorageAccountIdentityList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDatabricksWorkspaceStorageAccountIdentityList_Override(d DatabricksWorkspaceStorageAccountIdentityList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceStorageAccountIdentityList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) Get(index *float64) DatabricksWorkspaceStorageAccountIdentityOutputReference {
	var returns DatabricksWorkspaceStorageAccountIdentityOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DatabricksWorkspaceStorageAccountIdentityOutputReference interface {
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
	InternalValue() *DatabricksWorkspaceStorageAccountIdentity
	SetInternalValue(val *DatabricksWorkspaceStorageAccountIdentity)
	PrincipalId() *string
	TenantId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
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

// The jsii proxy struct for DatabricksWorkspaceStorageAccountIdentityOutputReference
type jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) InternalValue() *DatabricksWorkspaceStorageAccountIdentity {
	var returns *DatabricksWorkspaceStorageAccountIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewDatabricksWorkspaceStorageAccountIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatabricksWorkspaceStorageAccountIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceStorageAccountIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatabricksWorkspaceStorageAccountIdentityOutputReference_Override(d DatabricksWorkspaceStorageAccountIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceStorageAccountIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) SetInternalValue(val *DatabricksWorkspaceStorageAccountIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DatabricksWorkspaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#create DatabricksWorkspace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#delete DatabricksWorkspace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#read DatabricksWorkspace#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_workspace#update DatabricksWorkspace#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type DatabricksWorkspaceTimeoutsOutputReference interface {
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

// The jsii proxy struct for DatabricksWorkspaceTimeoutsOutputReference
type jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewDatabricksWorkspaceTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabricksWorkspaceTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabricksWorkspaceTimeoutsOutputReference_Override(d DatabricksWorkspaceTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		d,
		"resetCreate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		d,
		"resetRead",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

