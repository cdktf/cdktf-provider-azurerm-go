package kustoattacheddatabaseconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/kustoattacheddatabaseconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration azurerm_kusto_attached_database_configuration}.
type KustoAttachedDatabaseConfiguration interface {
	cdktf.TerraformResource
	AttachedDatabaseNames() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	ClusterResourceId() *string
	SetClusterResourceId(val *string)
	ClusterResourceIdInput() *string
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DefaultPrincipalModificationKind() *string
	SetDefaultPrincipalModificationKind(val *string)
	DefaultPrincipalModificationKindInput() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Sharing() KustoAttachedDatabaseConfigurationSharingOutputReference
	SharingInput() *KustoAttachedDatabaseConfigurationSharing
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KustoAttachedDatabaseConfigurationTimeoutsOutputReference
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
	PutSharing(value *KustoAttachedDatabaseConfigurationSharing)
	PutTimeouts(value *KustoAttachedDatabaseConfigurationTimeouts)
	ResetDefaultPrincipalModificationKind()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSharing()
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

// The jsii proxy struct for KustoAttachedDatabaseConfiguration
type jsiiProxy_KustoAttachedDatabaseConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) AttachedDatabaseNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attachedDatabaseNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) ClusterResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) ClusterResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) DefaultPrincipalModificationKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPrincipalModificationKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) DefaultPrincipalModificationKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPrincipalModificationKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Sharing() KustoAttachedDatabaseConfigurationSharingOutputReference {
	var returns KustoAttachedDatabaseConfigurationSharingOutputReference
	_jsii_.Get(
		j,
		"sharing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SharingInput() *KustoAttachedDatabaseConfigurationSharing {
	var returns *KustoAttachedDatabaseConfigurationSharing
	_jsii_.Get(
		j,
		"sharingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) Timeouts() KustoAttachedDatabaseConfigurationTimeoutsOutputReference {
	var returns KustoAttachedDatabaseConfigurationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration azurerm_kusto_attached_database_configuration} Resource.
func NewKustoAttachedDatabaseConfiguration(scope constructs.Construct, id *string, config *KustoAttachedDatabaseConfigurationConfig) KustoAttachedDatabaseConfiguration {
	_init_.Initialize()

	j := jsiiProxy_KustoAttachedDatabaseConfiguration{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoAttachedDatabaseConfiguration.KustoAttachedDatabaseConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration azurerm_kusto_attached_database_configuration} Resource.
func NewKustoAttachedDatabaseConfiguration_Override(k KustoAttachedDatabaseConfiguration, scope constructs.Construct, id *string, config *KustoAttachedDatabaseConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoAttachedDatabaseConfiguration.KustoAttachedDatabaseConfiguration",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetClusterResourceId(val *string) {
	_jsii_.Set(
		j,
		"clusterResourceId",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetDefaultPrincipalModificationKind(val *string) {
	_jsii_.Set(
		j,
		"defaultPrincipalModificationKind",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfiguration) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
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
func KustoAttachedDatabaseConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoAttachedDatabaseConfiguration.KustoAttachedDatabaseConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KustoAttachedDatabaseConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.kustoAttachedDatabaseConfiguration.KustoAttachedDatabaseConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) PutSharing(value *KustoAttachedDatabaseConfigurationSharing) {
	_jsii_.InvokeVoid(
		k,
		"putSharing",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) PutTimeouts(value *KustoAttachedDatabaseConfigurationTimeouts) {
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) ResetDefaultPrincipalModificationKind() {
	_jsii_.InvokeVoid(
		k,
		"resetDefaultPrincipalModificationKind",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) ResetSharing() {
	_jsii_.InvokeVoid(
		k,
		"resetSharing",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KustoAttachedDatabaseConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#cluster_name KustoAttachedDatabaseConfiguration#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#cluster_resource_id KustoAttachedDatabaseConfiguration#cluster_resource_id}.
	ClusterResourceId *string `field:"required" json:"clusterResourceId" yaml:"clusterResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#database_name KustoAttachedDatabaseConfiguration#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#location KustoAttachedDatabaseConfiguration#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#name KustoAttachedDatabaseConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#resource_group_name KustoAttachedDatabaseConfiguration#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#default_principal_modification_kind KustoAttachedDatabaseConfiguration#default_principal_modification_kind}.
	DefaultPrincipalModificationKind *string `field:"optional" json:"defaultPrincipalModificationKind" yaml:"defaultPrincipalModificationKind"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#id KustoAttachedDatabaseConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// sharing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#sharing KustoAttachedDatabaseConfiguration#sharing}
	Sharing *KustoAttachedDatabaseConfigurationSharing `field:"optional" json:"sharing" yaml:"sharing"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#timeouts KustoAttachedDatabaseConfiguration#timeouts}
	Timeouts *KustoAttachedDatabaseConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type KustoAttachedDatabaseConfigurationSharing struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#external_tables_to_exclude KustoAttachedDatabaseConfiguration#external_tables_to_exclude}.
	ExternalTablesToExclude *[]*string `field:"optional" json:"externalTablesToExclude" yaml:"externalTablesToExclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#external_tables_to_include KustoAttachedDatabaseConfiguration#external_tables_to_include}.
	ExternalTablesToInclude *[]*string `field:"optional" json:"externalTablesToInclude" yaml:"externalTablesToInclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#materialized_views_to_exclude KustoAttachedDatabaseConfiguration#materialized_views_to_exclude}.
	MaterializedViewsToExclude *[]*string `field:"optional" json:"materializedViewsToExclude" yaml:"materializedViewsToExclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#materialized_views_to_include KustoAttachedDatabaseConfiguration#materialized_views_to_include}.
	MaterializedViewsToInclude *[]*string `field:"optional" json:"materializedViewsToInclude" yaml:"materializedViewsToInclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#tables_to_exclude KustoAttachedDatabaseConfiguration#tables_to_exclude}.
	TablesToExclude *[]*string `field:"optional" json:"tablesToExclude" yaml:"tablesToExclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#tables_to_include KustoAttachedDatabaseConfiguration#tables_to_include}.
	TablesToInclude *[]*string `field:"optional" json:"tablesToInclude" yaml:"tablesToInclude"`
}

type KustoAttachedDatabaseConfigurationSharingOutputReference interface {
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
	ExternalTablesToExclude() *[]*string
	SetExternalTablesToExclude(val *[]*string)
	ExternalTablesToExcludeInput() *[]*string
	ExternalTablesToInclude() *[]*string
	SetExternalTablesToInclude(val *[]*string)
	ExternalTablesToIncludeInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KustoAttachedDatabaseConfigurationSharing
	SetInternalValue(val *KustoAttachedDatabaseConfigurationSharing)
	MaterializedViewsToExclude() *[]*string
	SetMaterializedViewsToExclude(val *[]*string)
	MaterializedViewsToExcludeInput() *[]*string
	MaterializedViewsToInclude() *[]*string
	SetMaterializedViewsToInclude(val *[]*string)
	MaterializedViewsToIncludeInput() *[]*string
	TablesToExclude() *[]*string
	SetTablesToExclude(val *[]*string)
	TablesToExcludeInput() *[]*string
	TablesToInclude() *[]*string
	SetTablesToInclude(val *[]*string)
	TablesToIncludeInput() *[]*string
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
	ResetExternalTablesToExclude()
	ResetExternalTablesToInclude()
	ResetMaterializedViewsToExclude()
	ResetMaterializedViewsToInclude()
	ResetTablesToExclude()
	ResetTablesToInclude()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KustoAttachedDatabaseConfigurationSharingOutputReference
type jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ExternalTablesToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalTablesToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ExternalTablesToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalTablesToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ExternalTablesToInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalTablesToInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ExternalTablesToIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalTablesToIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) InternalValue() *KustoAttachedDatabaseConfigurationSharing {
	var returns *KustoAttachedDatabaseConfigurationSharing
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) MaterializedViewsToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"materializedViewsToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) MaterializedViewsToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"materializedViewsToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) MaterializedViewsToInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"materializedViewsToInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) MaterializedViewsToIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"materializedViewsToIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TablesToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tablesToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TablesToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tablesToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TablesToInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tablesToInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TablesToIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tablesToIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKustoAttachedDatabaseConfigurationSharingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KustoAttachedDatabaseConfigurationSharingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoAttachedDatabaseConfiguration.KustoAttachedDatabaseConfigurationSharingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKustoAttachedDatabaseConfigurationSharingOutputReference_Override(k KustoAttachedDatabaseConfigurationSharingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoAttachedDatabaseConfiguration.KustoAttachedDatabaseConfigurationSharingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) SetExternalTablesToExclude(val *[]*string) {
	_jsii_.Set(
		j,
		"externalTablesToExclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) SetExternalTablesToInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"externalTablesToInclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) SetInternalValue(val *KustoAttachedDatabaseConfigurationSharing) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) SetMaterializedViewsToExclude(val *[]*string) {
	_jsii_.Set(
		j,
		"materializedViewsToExclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) SetMaterializedViewsToInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"materializedViewsToInclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) SetTablesToExclude(val *[]*string) {
	_jsii_.Set(
		j,
		"tablesToExclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) SetTablesToInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"tablesToInclude",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetExternalTablesToExclude() {
	_jsii_.InvokeVoid(
		k,
		"resetExternalTablesToExclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetExternalTablesToInclude() {
	_jsii_.InvokeVoid(
		k,
		"resetExternalTablesToInclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetMaterializedViewsToExclude() {
	_jsii_.InvokeVoid(
		k,
		"resetMaterializedViewsToExclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetMaterializedViewsToInclude() {
	_jsii_.InvokeVoid(
		k,
		"resetMaterializedViewsToInclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetTablesToExclude() {
	_jsii_.InvokeVoid(
		k,
		"resetTablesToExclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ResetTablesToInclude() {
	_jsii_.InvokeVoid(
		k,
		"resetTablesToInclude",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationSharingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KustoAttachedDatabaseConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#create KustoAttachedDatabaseConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#delete KustoAttachedDatabaseConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#read KustoAttachedDatabaseConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#update KustoAttachedDatabaseConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type KustoAttachedDatabaseConfigurationTimeoutsOutputReference interface {
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

// The jsii proxy struct for KustoAttachedDatabaseConfigurationTimeoutsOutputReference
type jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewKustoAttachedDatabaseConfigurationTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KustoAttachedDatabaseConfigurationTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoAttachedDatabaseConfiguration.KustoAttachedDatabaseConfigurationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKustoAttachedDatabaseConfigurationTimeoutsOutputReference_Override(k KustoAttachedDatabaseConfigurationTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoAttachedDatabaseConfiguration.KustoAttachedDatabaseConfigurationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		k,
		"resetCreate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		k,
		"resetDelete",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		k,
		"resetRead",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		k,
		"resetUpdate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoAttachedDatabaseConfigurationTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

