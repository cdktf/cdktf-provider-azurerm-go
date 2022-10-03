package botserviceazurebot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/botserviceazurebot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot azurerm_bot_service_azure_bot}.
type BotServiceAzureBot interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeveloperAppInsightsApiKey() *string
	SetDeveloperAppInsightsApiKey(val *string)
	DeveloperAppInsightsApiKeyInput() *string
	DeveloperAppInsightsApplicationId() *string
	SetDeveloperAppInsightsApplicationId(val *string)
	DeveloperAppInsightsApplicationIdInput() *string
	DeveloperAppInsightsKey() *string
	SetDeveloperAppInsightsKey(val *string)
	DeveloperAppInsightsKeyInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
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
	LuisAppIds() *[]*string
	SetLuisAppIds(val *[]*string)
	LuisAppIdsInput() *[]*string
	LuisKey() *string
	SetLuisKey(val *string)
	LuisKeyInput() *string
	MicrosoftAppId() *string
	SetMicrosoftAppId(val *string)
	MicrosoftAppIdInput() *string
	MicrosoftAppMsiId() *string
	SetMicrosoftAppMsiId(val *string)
	MicrosoftAppMsiIdInput() *string
	MicrosoftAppTenantId() *string
	SetMicrosoftAppTenantId(val *string)
	MicrosoftAppTenantIdInput() *string
	MicrosoftAppType() *string
	SetMicrosoftAppType(val *string)
	MicrosoftAppTypeInput() *string
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
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	StreamingEndpointEnabled() interface{}
	SetStreamingEndpointEnabled(val interface{})
	StreamingEndpointEnabledInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BotServiceAzureBotTimeoutsOutputReference
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
	PutTimeouts(value *BotServiceAzureBotTimeouts)
	ResetDeveloperAppInsightsApiKey()
	ResetDeveloperAppInsightsApplicationId()
	ResetDeveloperAppInsightsKey()
	ResetDisplayName()
	ResetEndpoint()
	ResetId()
	ResetLuisAppIds()
	ResetLuisKey()
	ResetMicrosoftAppMsiId()
	ResetMicrosoftAppTenantId()
	ResetMicrosoftAppType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStreamingEndpointEnabled()
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

// The jsii proxy struct for BotServiceAzureBot
type jsiiProxy_BotServiceAzureBot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BotServiceAzureBot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) LuisAppIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"luisAppIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) LuisAppIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"luisAppIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) LuisKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"luisKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) LuisKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"luisKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) MicrosoftAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) MicrosoftAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) MicrosoftAppMsiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppMsiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) MicrosoftAppMsiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppMsiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) MicrosoftAppTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) MicrosoftAppTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) MicrosoftAppType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) MicrosoftAppTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) StreamingEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) StreamingEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Timeouts() BotServiceAzureBotTimeoutsOutputReference {
	var returns BotServiceAzureBotTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot azurerm_bot_service_azure_bot} Resource.
func NewBotServiceAzureBot(scope constructs.Construct, id *string, config *BotServiceAzureBotConfig) BotServiceAzureBot {
	_init_.Initialize()

	j := jsiiProxy_BotServiceAzureBot{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.botServiceAzureBot.BotServiceAzureBot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot azurerm_bot_service_azure_bot} Resource.
func NewBotServiceAzureBot_Override(b BotServiceAzureBot, scope constructs.Construct, id *string, config *BotServiceAzureBotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.botServiceAzureBot.BotServiceAzureBot",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetDeveloperAppInsightsApiKey(val *string) {
	_jsii_.Set(
		j,
		"developerAppInsightsApiKey",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetDeveloperAppInsightsApplicationId(val *string) {
	_jsii_.Set(
		j,
		"developerAppInsightsApplicationId",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetDeveloperAppInsightsKey(val *string) {
	_jsii_.Set(
		j,
		"developerAppInsightsKey",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetEndpoint(val *string) {
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetLuisAppIds(val *[]*string) {
	_jsii_.Set(
		j,
		"luisAppIds",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetLuisKey(val *string) {
	_jsii_.Set(
		j,
		"luisKey",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetMicrosoftAppId(val *string) {
	_jsii_.Set(
		j,
		"microsoftAppId",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetMicrosoftAppMsiId(val *string) {
	_jsii_.Set(
		j,
		"microsoftAppMsiId",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetMicrosoftAppTenantId(val *string) {
	_jsii_.Set(
		j,
		"microsoftAppTenantId",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetMicrosoftAppType(val *string) {
	_jsii_.Set(
		j,
		"microsoftAppType",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetSku(val *string) {
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetStreamingEndpointEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"streamingEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot) SetTags(val *map[string]*string) {
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
func BotServiceAzureBot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.botServiceAzureBot.BotServiceAzureBot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BotServiceAzureBot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.botServiceAzureBot.BotServiceAzureBot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BotServiceAzureBot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BotServiceAzureBot) PutTimeouts(value *BotServiceAzureBotTimeouts) {
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetDeveloperAppInsightsApiKey() {
	_jsii_.InvokeVoid(
		b,
		"resetDeveloperAppInsightsApiKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetDeveloperAppInsightsApplicationId() {
	_jsii_.InvokeVoid(
		b,
		"resetDeveloperAppInsightsApplicationId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetDeveloperAppInsightsKey() {
	_jsii_.InvokeVoid(
		b,
		"resetDeveloperAppInsightsKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetDisplayName() {
	_jsii_.InvokeVoid(
		b,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetEndpoint() {
	_jsii_.InvokeVoid(
		b,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetLuisAppIds() {
	_jsii_.InvokeVoid(
		b,
		"resetLuisAppIds",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetLuisKey() {
	_jsii_.InvokeVoid(
		b,
		"resetLuisKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetMicrosoftAppMsiId() {
	_jsii_.InvokeVoid(
		b,
		"resetMicrosoftAppMsiId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetMicrosoftAppTenantId() {
	_jsii_.InvokeVoid(
		b,
		"resetMicrosoftAppTenantId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetMicrosoftAppType() {
	_jsii_.InvokeVoid(
		b,
		"resetMicrosoftAppType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetStreamingEndpointEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetStreamingEndpointEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BotServiceAzureBotConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#location BotServiceAzureBot#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#microsoft_app_id BotServiceAzureBot#microsoft_app_id}.
	MicrosoftAppId *string `field:"required" json:"microsoftAppId" yaml:"microsoftAppId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#name BotServiceAzureBot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#resource_group_name BotServiceAzureBot#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#sku BotServiceAzureBot#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#developer_app_insights_api_key BotServiceAzureBot#developer_app_insights_api_key}.
	DeveloperAppInsightsApiKey *string `field:"optional" json:"developerAppInsightsApiKey" yaml:"developerAppInsightsApiKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#developer_app_insights_application_id BotServiceAzureBot#developer_app_insights_application_id}.
	DeveloperAppInsightsApplicationId *string `field:"optional" json:"developerAppInsightsApplicationId" yaml:"developerAppInsightsApplicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#developer_app_insights_key BotServiceAzureBot#developer_app_insights_key}.
	DeveloperAppInsightsKey *string `field:"optional" json:"developerAppInsightsKey" yaml:"developerAppInsightsKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#display_name BotServiceAzureBot#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#endpoint BotServiceAzureBot#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#id BotServiceAzureBot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#luis_app_ids BotServiceAzureBot#luis_app_ids}.
	LuisAppIds *[]*string `field:"optional" json:"luisAppIds" yaml:"luisAppIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#luis_key BotServiceAzureBot#luis_key}.
	LuisKey *string `field:"optional" json:"luisKey" yaml:"luisKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#microsoft_app_msi_id BotServiceAzureBot#microsoft_app_msi_id}.
	MicrosoftAppMsiId *string `field:"optional" json:"microsoftAppMsiId" yaml:"microsoftAppMsiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#microsoft_app_tenant_id BotServiceAzureBot#microsoft_app_tenant_id}.
	MicrosoftAppTenantId *string `field:"optional" json:"microsoftAppTenantId" yaml:"microsoftAppTenantId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#microsoft_app_type BotServiceAzureBot#microsoft_app_type}.
	MicrosoftAppType *string `field:"optional" json:"microsoftAppType" yaml:"microsoftAppType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#streaming_endpoint_enabled BotServiceAzureBot#streaming_endpoint_enabled}.
	StreamingEndpointEnabled interface{} `field:"optional" json:"streamingEndpointEnabled" yaml:"streamingEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#tags BotServiceAzureBot#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#timeouts BotServiceAzureBot#timeouts}
	Timeouts *BotServiceAzureBotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type BotServiceAzureBotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#create BotServiceAzureBot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#delete BotServiceAzureBot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#read BotServiceAzureBot#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bot_service_azure_bot#update BotServiceAzureBot#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type BotServiceAzureBotTimeoutsOutputReference interface {
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

// The jsii proxy struct for BotServiceAzureBotTimeoutsOutputReference
type jsiiProxy_BotServiceAzureBotTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewBotServiceAzureBotTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BotServiceAzureBotTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BotServiceAzureBotTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.botServiceAzureBot.BotServiceAzureBotTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBotServiceAzureBotTimeoutsOutputReference_Override(b BotServiceAzureBotTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.botServiceAzureBot.BotServiceAzureBotTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		b,
		"resetCreate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		b,
		"resetDelete",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		b,
		"resetRead",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		b,
		"resetUpdate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBotTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

