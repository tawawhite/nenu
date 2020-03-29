package config

var (
	// ConfigPath config.yml file path
	ConfigPath string
	// TemplatePath path to web template to use for the blog
	TemplatePath string
	// ArticlesPath path to markdown (.md) articles
	ArticlesPath string
	// OutputPath path to where the site contents will be written
	OutputPath string
	// TempPath temporary path used during content generation
	TempPath string
	// YMLConfig configuration YAML file
	YMLConfig *YML
)
