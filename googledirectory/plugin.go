/*
Package googledirectory implements a steampipe plugin for googledirectory.

This plugin provides data that Steampipe uses to present foreign
tables that represent Google Directory resources.
*/
package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-googledirectory"

// Plugin creates this (googledirectory) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"404", "400"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"googledirectory_domain":       tableGoogleDirectroryDomain(ctx),
			"googledirectory_group":        tableGoogleDirectroryGroup(ctx),
			"googledirectory_group_member": tableGoogleDirectroryGroupMember(ctx),
			"googledirectory_org_unit":     tableGoogleDirectroryOrgUnit(ctx),
			"googledirectory_privilege":    tableGoogleDirectroryPrivilege(ctx),
			"googledirectory_role":         tableGoogleDirectroryRole(ctx),
			"googledirectory_user":         tableGoogleDirectroryUser(ctx),
		},
	}

	return p
}
