package googledirectory

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type googledirectoryConfig struct {
	CredentialFile        *string `cty:"credential_file"`
	ImpersonatedUserEmail *string `cty:"impersonated_user_email"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"credential_file": {
		Type: schema.TypeString,
	},
	"impersonated_user_email": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &googledirectoryConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) googledirectoryConfig {
	if connection == nil || connection.Config == nil {
		return googledirectoryConfig{}
	}
	config, _ := connection.Config.(googledirectoryConfig)
	return config
}
