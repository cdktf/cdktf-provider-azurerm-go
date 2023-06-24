package apimanagement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/apimanagement/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/api_management azurerm_api_management}.
type ApiManagement interface {
	cdktf.TerraformResource
	AdditionalLocation() ApiManagementAdditionalLocationList
	AdditionalLocationInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() ApiManagementCertificateList
	CertificateInput() interface{}
	ClientCertificateEnabled() interface{}
	SetClientCertificateEnabled(val interface{})
	ClientCertificateEnabledInput() interface{}
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
	Delegation() ApiManagementDelegationOutputReference
	DelegationInput() *ApiManagementDelegation
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeveloperPortalUrl() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayDisabled() interface{}
	SetGatewayDisabled(val interface{})
	GatewayDisabledInput() interface{}
	GatewayRegionalUrl() *string
	GatewayUrl() *string
	HostnameConfiguration() ApiManagementHostnameConfigurationOutputReference
	HostnameConfigurationInput() *ApiManagementHostnameConfiguration
	Id() *string
	SetId(val *string)
	Identity() ApiManagementIdentityOutputReference
	IdentityInput() *ApiManagementIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagementApiUrl() *string
	MinApiVersion() *string
	SetMinApiVersion(val *string)
	MinApiVersionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationSenderEmail() *string
	SetNotificationSenderEmail(val *string)
	NotificationSenderEmailInput() *string
	Policy() ApiManagementPolicyList
	PolicyInput() interface{}
	PortalUrl() *string
	PrivateIpAddresses() *[]*string
	Protocols() ApiManagementProtocolsOutputReference
	ProtocolsInput() *ApiManagementProtocols
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicIpAddresses() *[]*string
	PublicIpAddressId() *string
	SetPublicIpAddressId(val *string)
	PublicIpAddressIdInput() *string
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	PublisherEmail() *string
	SetPublisherEmail(val *string)
	PublisherEmailInput() *string
	PublisherName() *string
	SetPublisherName(val *string)
	PublisherNameInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ScmUrl() *string
	Security() ApiManagementSecurityOutputReference
	SecurityInput() *ApiManagementSecurity
	SignIn() ApiManagementSignInOutputReference
	SignInInput() *ApiManagementSignIn
	SignUp() ApiManagementSignUpOutputReference
	SignUpInput() *ApiManagementSignUp
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TenantAccess() ApiManagementTenantAccessOutputReference
	TenantAccessInput() *ApiManagementTenantAccess
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApiManagementTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualNetworkConfiguration() ApiManagementVirtualNetworkConfigurationOutputReference
	VirtualNetworkConfigurationInput() *ApiManagementVirtualNetworkConfiguration
	VirtualNetworkType() *string
	SetVirtualNetworkType(val *string)
	VirtualNetworkTypeInput() *string
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
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
	PutAdditionalLocation(value interface{})
	PutCertificate(value interface{})
	PutDelegation(value *ApiManagementDelegation)
	PutHostnameConfiguration(value *ApiManagementHostnameConfiguration)
	PutIdentity(value *ApiManagementIdentity)
	PutPolicy(value interface{})
	PutProtocols(value *ApiManagementProtocols)
	PutSecurity(value *ApiManagementSecurity)
	PutSignIn(value *ApiManagementSignIn)
	PutSignUp(value *ApiManagementSignUp)
	PutTenantAccess(value *ApiManagementTenantAccess)
	PutTimeouts(value *ApiManagementTimeouts)
	PutVirtualNetworkConfiguration(value *ApiManagementVirtualNetworkConfiguration)
	ResetAdditionalLocation()
	ResetCertificate()
	ResetClientCertificateEnabled()
	ResetDelegation()
	ResetGatewayDisabled()
	ResetHostnameConfiguration()
	ResetId()
	ResetIdentity()
	ResetMinApiVersion()
	ResetNotificationSenderEmail()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicy()
	ResetProtocols()
	ResetPublicIpAddressId()
	ResetPublicNetworkAccessEnabled()
	ResetSecurity()
	ResetSignIn()
	ResetSignUp()
	ResetTags()
	ResetTenantAccess()
	ResetTimeouts()
	ResetVirtualNetworkConfiguration()
	ResetVirtualNetworkType()
	ResetZones()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiManagement
type jsiiProxy_ApiManagement struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiManagement) AdditionalLocation() ApiManagementAdditionalLocationList {
	var returns ApiManagementAdditionalLocationList
	_jsii_.Get(
		j,
		"additionalLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) AdditionalLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Certificate() ApiManagementCertificateList {
	var returns ApiManagementCertificateList
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ClientCertificateEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ClientCertificateEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Delegation() ApiManagementDelegationOutputReference {
	var returns ApiManagementDelegationOutputReference
	_jsii_.Get(
		j,
		"delegation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) DelegationInput() *ApiManagementDelegation {
	var returns *ApiManagementDelegation
	_jsii_.Get(
		j,
		"delegationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) DeveloperPortalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerPortalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) GatewayDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) GatewayDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) GatewayRegionalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRegionalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) GatewayUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) HostnameConfiguration() ApiManagementHostnameConfigurationOutputReference {
	var returns ApiManagementHostnameConfigurationOutputReference
	_jsii_.Get(
		j,
		"hostnameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) HostnameConfigurationInput() *ApiManagementHostnameConfiguration {
	var returns *ApiManagementHostnameConfiguration
	_jsii_.Get(
		j,
		"hostnameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Identity() ApiManagementIdentityOutputReference {
	var returns ApiManagementIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) IdentityInput() *ApiManagementIdentity {
	var returns *ApiManagementIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ManagementApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementApiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) MinApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) MinApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minApiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) NotificationSenderEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationSenderEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) NotificationSenderEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationSenderEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Policy() ApiManagementPolicyList {
	var returns ApiManagementPolicyList
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PortalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Protocols() ApiManagementProtocolsOutputReference {
	var returns ApiManagementProtocolsOutputReference
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ProtocolsInput() *ApiManagementProtocols {
	var returns *ApiManagementProtocols
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublicIpAddressId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublicIpAddressIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublisherEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublisherEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublisherName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) PublisherNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ScmUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Security() ApiManagementSecurityOutputReference {
	var returns ApiManagementSecurityOutputReference
	_jsii_.Get(
		j,
		"security",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SecurityInput() *ApiManagementSecurity {
	var returns *ApiManagementSecurity
	_jsii_.Get(
		j,
		"securityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SignIn() ApiManagementSignInOutputReference {
	var returns ApiManagementSignInOutputReference
	_jsii_.Get(
		j,
		"signIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SignInInput() *ApiManagementSignIn {
	var returns *ApiManagementSignIn
	_jsii_.Get(
		j,
		"signInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SignUp() ApiManagementSignUpOutputReference {
	var returns ApiManagementSignUpOutputReference
	_jsii_.Get(
		j,
		"signUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SignUpInput() *ApiManagementSignUp {
	var returns *ApiManagementSignUp
	_jsii_.Get(
		j,
		"signUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TenantAccess() ApiManagementTenantAccessOutputReference {
	var returns ApiManagementTenantAccessOutputReference
	_jsii_.Get(
		j,
		"tenantAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TenantAccessInput() *ApiManagementTenantAccess {
	var returns *ApiManagementTenantAccess
	_jsii_.Get(
		j,
		"tenantAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Timeouts() ApiManagementTimeoutsOutputReference {
	var returns ApiManagementTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) VirtualNetworkConfiguration() ApiManagementVirtualNetworkConfigurationOutputReference {
	var returns ApiManagementVirtualNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"virtualNetworkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) VirtualNetworkConfigurationInput() *ApiManagementVirtualNetworkConfiguration {
	var returns *ApiManagementVirtualNetworkConfiguration
	_jsii_.Get(
		j,
		"virtualNetworkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) VirtualNetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) VirtualNetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagement) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/api_management azurerm_api_management} Resource.
func NewApiManagement(scope constructs.Construct, id *string, config *ApiManagementConfig) ApiManagement {
	_init_.Initialize()

	if err := validateNewApiManagementParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagement{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagement",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/api_management azurerm_api_management} Resource.
func NewApiManagement_Override(a ApiManagement, scope constructs.Construct, id *string, config *ApiManagementConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagement",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiManagement)SetClientCertificateEnabled(val interface{}) {
	if err := j.validateSetClientCertificateEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetGatewayDisabled(val interface{}) {
	if err := j.validateSetGatewayDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayDisabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetMinApiVersion(val *string) {
	if err := j.validateSetMinApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minApiVersion",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetNotificationSenderEmail(val *string) {
	if err := j.validateSetNotificationSenderEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationSenderEmail",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetPublicIpAddressId(val *string) {
	if err := j.validateSetPublicIpAddressIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpAddressId",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetPublisherEmail(val *string) {
	if err := j.validateSetPublisherEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publisherEmail",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetPublisherName(val *string) {
	if err := j.validateSetPublisherNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publisherName",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetVirtualNetworkType(val *string) {
	if err := j.validateSetVirtualNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkType",
		val,
	)
}

func (j *jsiiProxy_ApiManagement)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
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
func ApiManagement_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagement_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagement.ApiManagement",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagement_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagement_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagement.ApiManagement",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiManagement_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiManagement_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.apiManagement.ApiManagement",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiManagement_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.apiManagement.ApiManagement",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiManagement) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiManagement) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiManagement) PutAdditionalLocation(value interface{}) {
	if err := a.validatePutAdditionalLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdditionalLocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutCertificate(value interface{}) {
	if err := a.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutDelegation(value *ApiManagementDelegation) {
	if err := a.validatePutDelegationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDelegation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutHostnameConfiguration(value *ApiManagementHostnameConfiguration) {
	if err := a.validatePutHostnameConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHostnameConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutIdentity(value *ApiManagementIdentity) {
	if err := a.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdentity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutPolicy(value interface{}) {
	if err := a.validatePutPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutProtocols(value *ApiManagementProtocols) {
	if err := a.validatePutProtocolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProtocols",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutSecurity(value *ApiManagementSecurity) {
	if err := a.validatePutSecurityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecurity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutSignIn(value *ApiManagementSignIn) {
	if err := a.validatePutSignInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSignIn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutSignUp(value *ApiManagementSignUp) {
	if err := a.validatePutSignUpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSignUp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutTenantAccess(value *ApiManagementTenantAccess) {
	if err := a.validatePutTenantAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTenantAccess",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutTimeouts(value *ApiManagementTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) PutVirtualNetworkConfiguration(value *ApiManagementVirtualNetworkConfiguration) {
	if err := a.validatePutVirtualNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVirtualNetworkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagement) ResetAdditionalLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetClientCertificateEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetDelegation() {
	_jsii_.InvokeVoid(
		a,
		"resetDelegation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetGatewayDisabled() {
	_jsii_.InvokeVoid(
		a,
		"resetGatewayDisabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetHostnameConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetHostnameConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetMinApiVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetMinApiVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetNotificationSenderEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationSenderEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetProtocols() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocols",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetPublicIpAddressId() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicIpAddressId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetSecurity() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetSignIn() {
	_jsii_.InvokeVoid(
		a,
		"resetSignIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetSignUp() {
	_jsii_.InvokeVoid(
		a,
		"resetSignUp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetTenantAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetTenantAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetVirtualNetworkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualNetworkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetVirtualNetworkType() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualNetworkType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) ResetZones() {
	_jsii_.InvokeVoid(
		a,
		"resetZones",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagement) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagement) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

