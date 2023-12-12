package googledirectory

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type googledirectoryConfig struct {
	CredentialFile        *string `hcl:"credential_file"`
	Credentials           *string `hcl:"credentials"`
	ImpersonatedUserEmail *string `hcl:"impersonated_user_email"`
	TokenPath             *string `hcl:"token_path"`
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
