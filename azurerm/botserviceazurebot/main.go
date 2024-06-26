// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package botserviceazurebot

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.botServiceAzureBot.BotServiceAzureBot",
		reflect.TypeOf((*BotServiceAzureBot)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cmkKeyVaultKeyUrl", GoGetter: "CmkKeyVaultKeyUrl"},
			_jsii_.MemberProperty{JsiiProperty: "cmkKeyVaultKeyUrlInput", GoGetter: "CmkKeyVaultKeyUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "developerAppInsightsApiKey", GoGetter: "DeveloperAppInsightsApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "developerAppInsightsApiKeyInput", GoGetter: "DeveloperAppInsightsApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "developerAppInsightsApplicationId", GoGetter: "DeveloperAppInsightsApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "developerAppInsightsApplicationIdInput", GoGetter: "DeveloperAppInsightsApplicationIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "developerAppInsightsKey", GoGetter: "DeveloperAppInsightsKey"},
			_jsii_.MemberProperty{JsiiProperty: "developerAppInsightsKeyInput", GoGetter: "DeveloperAppInsightsKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "endpointInput", GoGetter: "EndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "iconUrl", GoGetter: "IconUrl"},
			_jsii_.MemberProperty{JsiiProperty: "iconUrlInput", GoGetter: "IconUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "localAuthenticationEnabled", GoGetter: "LocalAuthenticationEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "localAuthenticationEnabledInput", GoGetter: "LocalAuthenticationEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "luisAppIds", GoGetter: "LuisAppIds"},
			_jsii_.MemberProperty{JsiiProperty: "luisAppIdsInput", GoGetter: "LuisAppIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "luisKey", GoGetter: "LuisKey"},
			_jsii_.MemberProperty{JsiiProperty: "luisKeyInput", GoGetter: "LuisKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftAppId", GoGetter: "MicrosoftAppId"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftAppIdInput", GoGetter: "MicrosoftAppIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftAppMsiId", GoGetter: "MicrosoftAppMsiId"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftAppMsiIdInput", GoGetter: "MicrosoftAppMsiIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftAppTenantId", GoGetter: "MicrosoftAppTenantId"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftAppTenantIdInput", GoGetter: "MicrosoftAppTenantIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftAppType", GoGetter: "MicrosoftAppType"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftAppTypeInput", GoGetter: "MicrosoftAppTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "publicNetworkAccessEnabled", GoGetter: "PublicNetworkAccessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "publicNetworkAccessEnabledInput", GoGetter: "PublicNetworkAccessEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCmkKeyVaultKeyUrl", GoMethod: "ResetCmkKeyVaultKeyUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeveloperAppInsightsApiKey", GoMethod: "ResetDeveloperAppInsightsApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeveloperAppInsightsApplicationId", GoMethod: "ResetDeveloperAppInsightsApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeveloperAppInsightsKey", GoMethod: "ResetDeveloperAppInsightsKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpoint", GoMethod: "ResetEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetIconUrl", GoMethod: "ResetIconUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalAuthenticationEnabled", GoMethod: "ResetLocalAuthenticationEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLuisAppIds", GoMethod: "ResetLuisAppIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetLuisKey", GoMethod: "ResetLuisKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetMicrosoftAppMsiId", GoMethod: "ResetMicrosoftAppMsiId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMicrosoftAppTenantId", GoMethod: "ResetMicrosoftAppTenantId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMicrosoftAppType", GoMethod: "ResetMicrosoftAppType"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicNetworkAccessEnabled", GoMethod: "ResetPublicNetworkAccessEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetStreamingEndpointEnabled", GoMethod: "ResetStreamingEndpointEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sku", GoGetter: "Sku"},
			_jsii_.MemberProperty{JsiiProperty: "skuInput", GoGetter: "SkuInput"},
			_jsii_.MemberProperty{JsiiProperty: "streamingEndpointEnabled", GoGetter: "StreamingEndpointEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "streamingEndpointEnabledInput", GoGetter: "StreamingEndpointEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_BotServiceAzureBot{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.botServiceAzureBot.BotServiceAzureBotConfig",
		reflect.TypeOf((*BotServiceAzureBotConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.botServiceAzureBot.BotServiceAzureBotTimeouts",
		reflect.TypeOf((*BotServiceAzureBotTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.botServiceAzureBot.BotServiceAzureBotTimeoutsOutputReference",
		reflect.TypeOf((*BotServiceAzureBotTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_BotServiceAzureBotTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
