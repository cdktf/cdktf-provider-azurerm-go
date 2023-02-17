package mediatransform


type MediaTransformOutput struct {
	// audio_analyzer_preset block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_transform#audio_analyzer_preset MediaTransform#audio_analyzer_preset}
	AudioAnalyzerPreset *MediaTransformOutputAudioAnalyzerPreset `field:"optional" json:"audioAnalyzerPreset" yaml:"audioAnalyzerPreset"`
	// builtin_preset block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_transform#builtin_preset MediaTransform#builtin_preset}
	BuiltinPreset *MediaTransformOutputBuiltinPreset `field:"optional" json:"builtinPreset" yaml:"builtinPreset"`
	// face_detector_preset block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_transform#face_detector_preset MediaTransform#face_detector_preset}
	FaceDetectorPreset *MediaTransformOutputFaceDetectorPreset `field:"optional" json:"faceDetectorPreset" yaml:"faceDetectorPreset"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_transform#on_error_action MediaTransform#on_error_action}.
	OnErrorAction *string `field:"optional" json:"onErrorAction" yaml:"onErrorAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_transform#relative_priority MediaTransform#relative_priority}.
	RelativePriority *string `field:"optional" json:"relativePriority" yaml:"relativePriority"`
	// video_analyzer_preset block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_transform#video_analyzer_preset MediaTransform#video_analyzer_preset}
	VideoAnalyzerPreset *MediaTransformOutputVideoAnalyzerPreset `field:"optional" json:"videoAnalyzerPreset" yaml:"videoAnalyzerPreset"`
}

