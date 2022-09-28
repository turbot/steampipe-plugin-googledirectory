package main

import (
	"github.com/turbot/steampipe-plugin-googledirectory/googledirectory"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: googledirectory.Plugin})
}
