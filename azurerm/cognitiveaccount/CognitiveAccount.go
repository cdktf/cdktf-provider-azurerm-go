// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitiveaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/cognitiveaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/cognitive_account azurerm_cognitive_account}.
type CognitiveAccount interface {
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
	CustomerManagedKey() CognitiveAccountCustomerManagedKeyOutputReference
	CustomerManagedKeyInput() *CognitiveAccountCustomerManagedKey
	CustomQuestionAnsweringSearchServiceId() *string
	SetCustomQuestionAnsweringSearchServiceId(val *string)
	CustomQuestionAnsweringSearchServiceIdInput() *string
	CustomQuestionAnsweringSearchServiceKey() *string
	SetCustomQuestionAnsweringSearchServiceKey(val *string)
	CustomQuestionAnsweringSearchServiceKeyInput() *string
	CustomSubdomainName() *string
	SetCustomSubdomainName(val *string)
	CustomSubdomainNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DynamicThrottlingEnabled() interface{}
	SetDynamicThrottlingEnabled(val interface{})
	DynamicThrottlingEnabledInput() interface{}
	Endpoint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdns() *[]*string
	SetFqdns(val *[]*string)
	FqdnsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	Identity() CognitiveAccountIdentityOutputReference
	IdentityInput() *CognitiveAccountIdentity
	IdInput() *string
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalAuthEnabled() interface{}
	SetLocalAuthEnabled(val interface{})
	LocalAuthEnabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MetricsAdvisorAadClientId() *string
	SetMetricsAdvisorAadClientId(val *string)
	MetricsAdvisorAadClientIdInput() *string
	MetricsAdvisorAadTenantId() *string
	SetMetricsAdvisorAadTenantId(val *string)
	MetricsAdvisorAadTenantIdInput() *string
	MetricsAdvisorSuperUserName() *string
	SetMetricsAdvisorSuperUserName(val *string)
	MetricsAdvisorSuperUserNameInput() *string
	MetricsAdvisorWebsiteName() *string
	SetMetricsAdvisorWebsiteName(val *string)
	MetricsAdvisorWebsiteNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkAcls() CognitiveAccountNetworkAclsOutputReference
	NetworkAclsInput() *CognitiveAccountNetworkAcls
	// The tree node.
	Node() constructs.Node
	OutboundNetworkAccessRestricted() interface{}
	SetOutboundNetworkAccessRestricted(val interface{})
	OutboundNetworkAccessRestrictedInput() interface{}
	PrimaryAccessKey() *string
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
	QnaRuntimeEndpoint() *string
	SetQnaRuntimeEndpoint(val *string)
	QnaRuntimeEndpointInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SecondaryAccessKey() *string
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	Storage() CognitiveAccountStorageList
	StorageInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CognitiveAccountTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutCustomerManagedKey(value *CognitiveAccountCustomerManagedKey)
	PutIdentity(value *CognitiveAccountIdentity)
	PutNetworkAcls(value *CognitiveAccountNetworkAcls)
	PutStorage(value interface{})
	PutTimeouts(value *CognitiveAccountTimeouts)
	ResetCustomerManagedKey()
	ResetCustomQuestionAnsweringSearchServiceId()
	ResetCustomQuestionAnsweringSearchServiceKey()
	ResetCustomSubdomainName()
	ResetDynamicThrottlingEnabled()
	ResetFqdns()
	ResetId()
	ResetIdentity()
	ResetLocalAuthEnabled()
	ResetMetricsAdvisorAadClientId()
	ResetMetricsAdvisorAadTenantId()
	ResetMetricsAdvisorSuperUserName()
	ResetMetricsAdvisorWebsiteName()
	ResetNetworkAcls()
	ResetOutboundNetworkAccessRestricted()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetQnaRuntimeEndpoint()
	ResetStorage()
	ResetTags()
	ResetTimeouts()
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

// The jsii proxy struct for CognitiveAccount
type jsiiProxy_CognitiveAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitiveAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) CustomerManagedKey() CognitiveAccountCustomerManagedKeyOutputReference {
	var returns CognitiveAccountCustomerManagedKeyOutputReference
	_jsii_.Get(
		j,
		"customerManagedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) CustomerManagedKeyInput() *CognitiveAccountCustomerManagedKey {
	var returns *CognitiveAccountCustomerManagedKey
	_jsii_.Get(
		j,
		"customerManagedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) CustomQuestionAnsweringSearchServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customQuestionAnsweringSearchServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) CustomQuestionAnsweringSearchServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customQuestionAnsweringSearchServiceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) CustomQuestionAnsweringSearchServiceKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customQuestionAnsweringSearchServiceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) CustomQuestionAnsweringSearchServiceKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customQuestionAnsweringSearchServiceKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) CustomSubdomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSubdomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) CustomSubdomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSubdomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) DynamicThrottlingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicThrottlingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) DynamicThrottlingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicThrottlingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Fqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) FqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Identity() CognitiveAccountIdentityOutputReference {
	var returns CognitiveAccountIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) IdentityInput() *CognitiveAccountIdentity {
	var returns *CognitiveAccountIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) LocalAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) LocalAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) MetricsAdvisorAadClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsAdvisorAadClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) MetricsAdvisorAadClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsAdvisorAadClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) MetricsAdvisorAadTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsAdvisorAadTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) MetricsAdvisorAadTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsAdvisorAadTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) MetricsAdvisorSuperUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsAdvisorSuperUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) MetricsAdvisorSuperUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsAdvisorSuperUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) MetricsAdvisorWebsiteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsAdvisorWebsiteName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) MetricsAdvisorWebsiteNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsAdvisorWebsiteNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) NetworkAcls() CognitiveAccountNetworkAclsOutputReference {
	var returns CognitiveAccountNetworkAclsOutputReference
	_jsii_.Get(
		j,
		"networkAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) NetworkAclsInput() *CognitiveAccountNetworkAcls {
	var returns *CognitiveAccountNetworkAcls
	_jsii_.Get(
		j,
		"networkAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) OutboundNetworkAccessRestricted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundNetworkAccessRestricted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) OutboundNetworkAccessRestrictedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundNetworkAccessRestrictedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) PrimaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) QnaRuntimeEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qnaRuntimeEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) QnaRuntimeEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qnaRuntimeEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) SecondaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Storage() CognitiveAccountStorageList {
	var returns CognitiveAccountStorageList
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) StorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) Timeouts() CognitiveAccountTimeoutsOutputReference {
	var returns CognitiveAccountTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitiveAccount) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/cognitive_account azurerm_cognitive_account} Resource.
func NewCognitiveAccount(scope constructs.Construct, id *string, config *CognitiveAccountConfig) CognitiveAccount {
	_init_.Initialize()

	if err := validateNewCognitiveAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitiveAccount{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cognitiveAccount.CognitiveAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/cognitive_account azurerm_cognitive_account} Resource.
func NewCognitiveAccount_Override(c CognitiveAccount, scope constructs.Construct, id *string, config *CognitiveAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cognitiveAccount.CognitiveAccount",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetCustomQuestionAnsweringSearchServiceId(val *string) {
	if err := j.validateSetCustomQuestionAnsweringSearchServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customQuestionAnsweringSearchServiceId",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetCustomQuestionAnsweringSearchServiceKey(val *string) {
	if err := j.validateSetCustomQuestionAnsweringSearchServiceKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customQuestionAnsweringSearchServiceKey",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetCustomSubdomainName(val *string) {
	if err := j.validateSetCustomSubdomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSubdomainName",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetDynamicThrottlingEnabled(val interface{}) {
	if err := j.validateSetDynamicThrottlingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicThrottlingEnabled",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetFqdns(val *[]*string) {
	if err := j.validateSetFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fqdns",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetLocalAuthEnabled(val interface{}) {
	if err := j.validateSetLocalAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetMetricsAdvisorAadClientId(val *string) {
	if err := j.validateSetMetricsAdvisorAadClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsAdvisorAadClientId",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetMetricsAdvisorAadTenantId(val *string) {
	if err := j.validateSetMetricsAdvisorAadTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsAdvisorAadTenantId",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetMetricsAdvisorSuperUserName(val *string) {
	if err := j.validateSetMetricsAdvisorSuperUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsAdvisorSuperUserName",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetMetricsAdvisorWebsiteName(val *string) {
	if err := j.validateSetMetricsAdvisorWebsiteNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsAdvisorWebsiteName",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetOutboundNetworkAccessRestricted(val interface{}) {
	if err := j.validateSetOutboundNetworkAccessRestrictedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundNetworkAccessRestricted",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetQnaRuntimeEndpoint(val *string) {
	if err := j.validateSetQnaRuntimeEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qnaRuntimeEndpoint",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_CognitiveAccount)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a CognitiveAccount resource upon running "cdktf plan <stack-name>".
func CognitiveAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCognitiveAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cognitiveAccount.CognitiveAccount",
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
func CognitiveAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitiveAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cognitiveAccount.CognitiveAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitiveAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitiveAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cognitiveAccount.CognitiveAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitiveAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitiveAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cognitiveAccount.CognitiveAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitiveAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cognitiveAccount.CognitiveAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CognitiveAccount) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CognitiveAccount) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CognitiveAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CognitiveAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CognitiveAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CognitiveAccount) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CognitiveAccount) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitiveAccount) PutCustomerManagedKey(value *CognitiveAccountCustomerManagedKey) {
	if err := c.validatePutCustomerManagedKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomerManagedKey",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitiveAccount) PutIdentity(value *CognitiveAccountIdentity) {
	if err := c.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIdentity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitiveAccount) PutNetworkAcls(value *CognitiveAccountNetworkAcls) {
	if err := c.validatePutNetworkAclsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkAcls",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitiveAccount) PutStorage(value interface{}) {
	if err := c.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStorage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitiveAccount) PutTimeouts(value *CognitiveAccountTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetCustomerManagedKey() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomerManagedKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetCustomQuestionAnsweringSearchServiceId() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomQuestionAnsweringSearchServiceId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetCustomQuestionAnsweringSearchServiceKey() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomQuestionAnsweringSearchServiceKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetCustomSubdomainName() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomSubdomainName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetDynamicThrottlingEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDynamicThrottlingEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetFqdns() {
	_jsii_.InvokeVoid(
		c,
		"resetFqdns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetIdentity() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetLocalAuthEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalAuthEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetMetricsAdvisorAadClientId() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricsAdvisorAadClientId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetMetricsAdvisorAadTenantId() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricsAdvisorAadTenantId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetMetricsAdvisorSuperUserName() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricsAdvisorSuperUserName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetMetricsAdvisorWebsiteName() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricsAdvisorWebsiteName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetNetworkAcls() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkAcls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetOutboundNetworkAccessRestricted() {
	_jsii_.InvokeVoid(
		c,
		"resetOutboundNetworkAccessRestricted",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetQnaRuntimeEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetQnaRuntimeEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitiveAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitiveAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

