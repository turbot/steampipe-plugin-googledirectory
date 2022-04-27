package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"

	admin "google.golang.org/api/admin/directory/v1"
)

//// TABLE DEFINITION

func tableGoogleDirectoryRole(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_role",
		Description: "Roles defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryRoles,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "customer_id",
					Require: plugin.Optional,
				},
			},
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("role_id"),
			Hydrate:    getDirectoryRole,
		},
		Columns: []*plugin.Column{
			{
				Name:        "role_name",
				Description: "The name of the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role_id",
				Description: "The unique ID for the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_super_admin_role",
				Description: "Indicates whether the role is a super admin role, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_system_role",
				Description: "Indicates whether the role is a pre-defined system role, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "role_description",
				Description: "A short description of the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "customer_id",
				Description: "The customer ID to retrieve all account roles.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("customer_id"),
			},
			{
				Name:        "etag",
				Description: "A hash of the metadata, used to ensure there were no concurrent modifications to the resource when attempting an update.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "kind",
				Description: "The type of the API resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role_privileges",
				Description: "The set of privileges that are granted to this role.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryRoles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	// Set default value to my_customer, to represent current account
	customerID := "my_customer"
	if d.KeyColumnQuals["customer_id"] != nil {
		customerID = d.KeyColumnQuals["customer_id"].GetStringValue()
	}

	resp := service.Roles.List(customerID)
	if err := resp.Pages(ctx, func(page *admin.Roles) error {
		for _, role := range page.Items {
			d.StreamListItem(ctx, role)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if plugin.IsCancelled(ctx) {
				page.NextPageToken = ""
				break
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, err
}

//// HYDRATE FUNCTIONS

func getDirectoryRole(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryRole")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	roleID := d.KeyColumnQuals["role_id"].GetStringValue()

	// Return nil, if no input provided
	if roleID == "" {
		return nil, nil
	}

	resp, err := service.Roles.Get("my_customer", roleID).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
