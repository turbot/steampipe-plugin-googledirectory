package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

//// TABLE DEFINITION

func tableGoogleDirectroryPrivilege(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_privilege",
		Description: "Privileges defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryPrivileges,
		},
		Columns: []*plugin.Column{
			{
				Name:        "privilege_name",
				Description: "The name of the privilege.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "service_name",
				Description: "The name of the service this privilege is for.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "service_id",
				Description: "The obfuscated ID of the service this privilege is for.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_ou_scopable",
				Description: "Indicates if the privilege can be restricted to an organization unit.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "etag",
				Description: "Specifies the etag of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "kind",
				Description: "The type of the API resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "child_privileges",
				Description: "A list of child privileges. Privileges for a service form a tree. Each privilege can have a list of child privileges; this list is empty for a leaf privilege.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryPrivileges(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	resp, err := service.Privileges.List("my_customer").Do()
	if err != nil {
		return nil, err
	}

	for _, role := range resp.Items {
		d.StreamListItem(ctx, role)
	}

	return nil, err
}
