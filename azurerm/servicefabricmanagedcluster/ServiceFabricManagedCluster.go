// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabricmanagedcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/servicefabricmanagedcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/service_fabric_managed_cluster azurerm_service_fabric_managed_cluster}.
type ServiceFabricManagedCluster interface {
	cdktf.TerraformResource
	Authentication() ServiceFabricManagedClusterAuthenticationOutputReference
	AuthenticationInput() *ServiceFabricManagedClusterAuthentication
	BackupServiceEnabled() interface{}
	SetBackupServiceEnabled(val interface{})
	BackupServiceEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientConnectionPort() *float64
	SetClientConnectionPort(val *float64)
	ClientConnectionPortInput() *float64
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
	CustomFabricSetting() ServiceFabricManagedClusterCustomFabricSettingList
	CustomFabricSettingInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DnsName() *string
	SetDnsName(val *string)
	DnsNameInput() *string
	DnsServiceEnabled() interface{}
	SetDnsServiceEnabled(val interface{})
	DnsServiceEnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpGatewayPort() *float64
	SetHttpGatewayPort(val *float64)
	HttpGatewayPortInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	LbRule() ServiceFabricManagedClusterLbRuleList
	LbRuleInput() interface{}
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
	NodeType() ServiceFabricManagedClusterNodeTypeList
	NodeTypeInput() interface{}
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
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
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ServiceFabricManagedClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpgradeWave() *string
	SetUpgradeWave(val *string)
	UpgradeWaveInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutAuthentication(value *ServiceFabricManagedClusterAuthentication)
	PutCustomFabricSetting(value interface{})
	PutLbRule(value interface{})
	PutNodeType(value interface{})
	PutTimeouts(value *ServiceFabricManagedClusterTimeouts)
	ResetAuthentication()
	ResetBackupServiceEnabled()
	ResetCustomFabricSetting()
	ResetDnsName()
	ResetDnsServiceEnabled()
	ResetId()
	ResetNodeType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetSku()
	ResetTags()
	ResetTimeouts()
	ResetUpgradeWave()
	ResetUsername()
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

// The jsii proxy struct for ServiceFabricManagedCluster
type jsiiProxy_ServiceFabricManagedCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Authentication() ServiceFabricManagedClusterAuthenticationOutputReference {
	var returns ServiceFabricManagedClusterAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) AuthenticationInput() *ServiceFabricManagedClusterAuthentication {
	var returns *ServiceFabricManagedClusterAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) BackupServiceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupServiceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) BackupServiceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupServiceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ClientConnectionPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientConnectionPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ClientConnectionPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientConnectionPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) CustomFabricSetting() ServiceFabricManagedClusterCustomFabricSettingList {
	var returns ServiceFabricManagedClusterCustomFabricSettingList
	_jsii_.Get(
		j,
		"customFabricSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) CustomFabricSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFabricSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) DnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) DnsServiceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsServiceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) DnsServiceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsServiceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) HttpGatewayPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpGatewayPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) HttpGatewayPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpGatewayPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) LbRule() ServiceFabricManagedClusterLbRuleList {
	var returns ServiceFabricManagedClusterLbRuleList
	_jsii_.Get(
		j,
		"lbRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) LbRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) NodeType() ServiceFabricManagedClusterNodeTypeList {
	var returns ServiceFabricManagedClusterNodeTypeList
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) NodeTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Timeouts() ServiceFabricManagedClusterTimeoutsOutputReference {
	var returns ServiceFabricManagedClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) UpgradeWave() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeWave",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) UpgradeWaveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeWaveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/service_fabric_managed_cluster azurerm_service_fabric_managed_cluster} Resource.
func NewServiceFabricManagedCluster(scope constructs.Construct, id *string, config *ServiceFabricManagedClusterConfig) ServiceFabricManagedCluster {
	_init_.Initialize()

	if err := validateNewServiceFabricManagedClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceFabricManagedCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/service_fabric_managed_cluster azurerm_service_fabric_managed_cluster} Resource.
func NewServiceFabricManagedCluster_Override(s ServiceFabricManagedCluster, scope constructs.Construct, id *string, config *ServiceFabricManagedClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedCluster",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetBackupServiceEnabled(val interface{}) {
	if err := j.validateSetBackupServiceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupServiceEnabled",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetClientConnectionPort(val *float64) {
	if err := j.validateSetClientConnectionPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientConnectionPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetDnsName(val *string) {
	if err := j.validateSetDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetDnsServiceEnabled(val interface{}) {
	if err := j.validateSetDnsServiceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServiceEnabled",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetHttpGatewayPort(val *float64) {
	if err := j.validateSetHttpGatewayPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpGatewayPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetUpgradeWave(val *string) {
	if err := j.validateSetUpgradeWaveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeWave",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a ServiceFabricManagedCluster resource upon running "cdktf plan <stack-name>".
func ServiceFabricManagedCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServiceFabricManagedCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedCluster",
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
func ServiceFabricManagedCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceFabricManagedCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceFabricManagedCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceFabricManagedCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceFabricManagedCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceFabricManagedCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServiceFabricManagedCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServiceFabricManagedCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceFabricManagedCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServiceFabricManagedCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServiceFabricManagedCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServiceFabricManagedCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServiceFabricManagedCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServiceFabricManagedCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServiceFabricManagedCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServiceFabricManagedCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceFabricManagedCluster) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) PutAuthentication(value *ServiceFabricManagedClusterAuthentication) {
	if err := s.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) PutCustomFabricSetting(value interface{}) {
	if err := s.validatePutCustomFabricSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomFabricSetting",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) PutLbRule(value interface{}) {
	if err := s.validatePutLbRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLbRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) PutNodeType(value interface{}) {
	if err := s.validatePutNodeTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNodeType",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) PutTimeouts(value *ServiceFabricManagedClusterTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetAuthentication() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetBackupServiceEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetBackupServiceEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetCustomFabricSetting() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomFabricSetting",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetDnsName() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetDnsServiceEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsServiceEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetNodeType() {
	_jsii_.InvokeVoid(
		s,
		"resetNodeType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetSku() {
	_jsii_.InvokeVoid(
		s,
		"resetSku",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetUpgradeWave() {
	_jsii_.InvokeVoid(
		s,
		"resetUpgradeWave",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetUsername() {
	_jsii_.InvokeVoid(
		s,
		"resetUsername",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

