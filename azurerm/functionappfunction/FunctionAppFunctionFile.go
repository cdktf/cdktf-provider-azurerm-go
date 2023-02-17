package functionappfunction


type FunctionAppFunctionFile struct {
	// The content of the file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/function_app_function#content FunctionAppFunction#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The filename of the file to be uploaded.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/function_app_function#name FunctionAppFunction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

