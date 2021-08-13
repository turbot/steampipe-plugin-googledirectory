package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

//// TABLE DEFINITION

func tableGoogleDirectoryGroupMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_group_member",
		Description: "Group members defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate:    listDirectoryGroupMembers,
			KeyColumns: plugin.SingleColumn("group_id"),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"group_id", "id"}),
			Hydrate:    getDirectoryGroupMember,
		},
		Columns: []*plugin.Column{
			{
				Name:        "group_id",
				Description: "Specifies the ID of the group, the user belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("group_id"),
			},
			{
				Name:        "id",
				Description: "The unique ID of the group member.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "Specifies the member's email address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role",
				Description: "Specifies the role of the member in a group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "Specifies the status of the member.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "delivery_settings",
				Description: "Defines mail delivery preferences of member.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getDirectoryGroupMember,
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
				Name:        "type",
				Description: "The type of group member.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryGroupMembers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}
	groupID := d.KeyColumnQuals["group_id"].GetStringValue()

	resp := service.Members.List(groupID)
	if err := resp.Pages(ctx, func(page *admin.Members) error {
		for _, member := range page.Members {
			d.StreamListItem(ctx, member)
		}
		return nil
	}); err != nil {
		// Return nil, if given group is not present
		if err.(*googleapi.Error).Code == 404 {
			return nil, nil
		}
		return nil, err
	}

	return nil, err
}

//// HYDRATE FUNCTIONS

func getDirectoryGroupMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryGroupMember")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	var groupID, memberID string
	if h.Item != nil {
		data := h.Item.(*admin.Member)
		groupID = d.KeyColumnQuals["group_id"].GetStringValue()
		memberID = data.Id
	} else {
		groupID = d.KeyColumnQuals["group_id"].GetStringValue()
		memberID = d.KeyColumnQuals["id"].GetStringValue()
	}

	// Return nil, if no input provided
	if groupID == "" || memberID == "" {
		return nil, nil
	}

	resp, err := service.Members.Get(groupID, memberID).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
