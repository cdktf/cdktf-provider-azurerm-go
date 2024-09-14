// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppAuthSettingsV2 struct {
	// login block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#login WindowsWebApp#login}
	Login *WindowsWebAppAuthSettingsV2Login `field:"required" json:"login" yaml:"login"`
	// active_directory_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#active_directory_v2 WindowsWebApp#active_directory_v2}
	ActiveDirectoryV2 *WindowsWebAppAuthSettingsV2ActiveDirectoryV2 `field:"optional" json:"activeDirectoryV2" yaml:"activeDirectoryV2"`
	// apple_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#apple_v2 WindowsWebApp#apple_v2}
	AppleV2 *WindowsWebAppAuthSettingsV2AppleV2 `field:"optional" json:"appleV2" yaml:"appleV2"`
	// Should the AuthV2 Settings be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#auth_enabled WindowsWebApp#auth_enabled}
	AuthEnabled interface{} `field:"optional" json:"authEnabled" yaml:"authEnabled"`
	// azure_static_web_app_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#azure_static_web_app_v2 WindowsWebApp#azure_static_web_app_v2}
	AzureStaticWebAppV2 *WindowsWebAppAuthSettingsV2AzureStaticWebAppV2 `field:"optional" json:"azureStaticWebAppV2" yaml:"azureStaticWebAppV2"`
	// The path to the App Auth settings. **Note:** Relative Paths are evaluated from the Site Root directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#config_file_path WindowsWebApp#config_file_path}
	ConfigFilePath *string `field:"optional" json:"configFilePath" yaml:"configFilePath"`
	// custom_oidc_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#custom_oidc_v2 WindowsWebApp#custom_oidc_v2}
	CustomOidcV2 interface{} `field:"optional" json:"customOidcV2" yaml:"customOidcV2"`
	// The Default Authentication Provider to use when the `unauthenticated_action` is set to `RedirectToLoginPage`.
	//
	// Possible values include: `apple`, `azureactivedirectory`, `facebook`, `github`, `google`, `twitter` and the `name` of your `custom_oidc_v2` provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#default_provider WindowsWebApp#default_provider}
	DefaultProvider *string `field:"optional" json:"defaultProvider" yaml:"defaultProvider"`
	// The paths which should be excluded from the `unauthenticated_action` when it is set to `RedirectToLoginPage`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#excluded_paths WindowsWebApp#excluded_paths}
	ExcludedPaths *[]*string `field:"optional" json:"excludedPaths" yaml:"excludedPaths"`
	// facebook_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#facebook_v2 WindowsWebApp#facebook_v2}
	FacebookV2 *WindowsWebAppAuthSettingsV2FacebookV2 `field:"optional" json:"facebookV2" yaml:"facebookV2"`
	// The convention used to determine the url of the request made.
	//
	// Possible values include `ForwardProxyConventionNoProxy`, `ForwardProxyConventionStandard`, `ForwardProxyConventionCustom`. Defaults to `ForwardProxyConventionNoProxy`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#forward_proxy_convention WindowsWebApp#forward_proxy_convention}
	ForwardProxyConvention *string `field:"optional" json:"forwardProxyConvention" yaml:"forwardProxyConvention"`
	// The name of the header containing the host of the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#forward_proxy_custom_host_header_name WindowsWebApp#forward_proxy_custom_host_header_name}
	ForwardProxyCustomHostHeaderName *string `field:"optional" json:"forwardProxyCustomHostHeaderName" yaml:"forwardProxyCustomHostHeaderName"`
	// The name of the header containing the scheme of the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#forward_proxy_custom_scheme_header_name WindowsWebApp#forward_proxy_custom_scheme_header_name}
	ForwardProxyCustomSchemeHeaderName *string `field:"optional" json:"forwardProxyCustomSchemeHeaderName" yaml:"forwardProxyCustomSchemeHeaderName"`
	// github_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#github_v2 WindowsWebApp#github_v2}
	GithubV2 *WindowsWebAppAuthSettingsV2GithubV2 `field:"optional" json:"githubV2" yaml:"githubV2"`
	// google_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#google_v2 WindowsWebApp#google_v2}
	GoogleV2 *WindowsWebAppAuthSettingsV2GoogleV2 `field:"optional" json:"googleV2" yaml:"googleV2"`
	// The prefix that should precede all the authentication and authorisation paths. Defaults to `/.auth`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#http_route_api_prefix WindowsWebApp#http_route_api_prefix}
	HttpRouteApiPrefix *string `field:"optional" json:"httpRouteApiPrefix" yaml:"httpRouteApiPrefix"`
	// microsoft_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#microsoft_v2 WindowsWebApp#microsoft_v2}
	MicrosoftV2 *WindowsWebAppAuthSettingsV2MicrosoftV2 `field:"optional" json:"microsoftV2" yaml:"microsoftV2"`
	// Should the authentication flow be used for all requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#require_authentication WindowsWebApp#require_authentication}
	RequireAuthentication interface{} `field:"optional" json:"requireAuthentication" yaml:"requireAuthentication"`
	// Should HTTPS be required on connections? Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#require_https WindowsWebApp#require_https}
	RequireHttps interface{} `field:"optional" json:"requireHttps" yaml:"requireHttps"`
	// The Runtime Version of the Authentication and Authorisation feature of this App. Defaults to `~1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#runtime_version WindowsWebApp#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// twitter_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#twitter_v2 WindowsWebApp#twitter_v2}
	TwitterV2 *WindowsWebAppAuthSettingsV2TwitterV2 `field:"optional" json:"twitterV2" yaml:"twitterV2"`
	// The action to take for requests made without authentication.
	//
	// Possible values include `RedirectToLoginPage`, `AllowAnonymous`, `Return401`, and `Return403`. Defaults to `RedirectToLoginPage`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app#unauthenticated_action WindowsWebApp#unauthenticated_action}
	UnauthenticatedAction *string `field:"optional" json:"unauthenticatedAction" yaml:"unauthenticatedAction"`
}

