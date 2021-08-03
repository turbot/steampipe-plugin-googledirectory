package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"

	admin "google.golang.org/api/admin/directory/v1"
)

//// TABLE DEFINITION

func tableGoogleDirectroryGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_group",
		Description: "Groups defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryGroups,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "email"}),
			Hydrate:    getDirectoryGroup,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The group's display name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The unique ID of a group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "Specifies the group's email address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "admin_created",
				Description: "Indicates whether the group is created by an administrator, or by an user.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "description",
				Description: "An extended description to help users determine the purpose of a group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "direct_members_count",
				Description: "The number of users that are direct members of the group.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "etag",
				Description: "Specifies the ETag of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "kind",
				Description: "The type of the API resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "aliases",
				Description: "A list of the group's alias email addresses.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "non_editable_aliases",
				Description: "A list of the group's non-editable alias email addresses that are outside of the account's primary domain or subdomains.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	resp := service.Groups.List().Customer("my_customer")
	if err := resp.Pages(ctx, func(page *admin.Groups) error {
		for _, group := range page.Groups {
			d.StreamListItem(ctx, group)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, err
}

//// HYDRATE FUNCTIONS

func getDirectoryGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryGroup")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetStringValue()
	email := d.KeyColumnQuals["email"].GetStringValue()

	// Return nil, if no input provided
	if id == "" && email == "" {
		return nil, nil
	}

	var inputStr string
	if id == "" {
		inputStr = email
	} else {
		inputStr = id
	}

	resp, err := service.Groups.Get(inputStr).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
