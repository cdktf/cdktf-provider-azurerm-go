package mediatransform


type MediaTransformOutputVideoAnalyzerPreset struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_transform#audio_analysis_mode MediaTransform#audio_analysis_mode}.
	AudioAnalysisMode *string `field:"optional" json:"audioAnalysisMode" yaml:"audioAnalysisMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_transform#audio_language MediaTransform#audio_language}.
	AudioLanguage *string `field:"optional" json:"audioLanguage" yaml:"audioLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_transform#insights_type MediaTransform#insights_type}.
	InsightsType *string `field:"optional" json:"insightsType" yaml:"insightsType"`
}

