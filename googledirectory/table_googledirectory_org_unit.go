package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableGoogleDirectoryOrgUnit(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_org_unit",
		Description: "OrgUnits defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryOrgUnits,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "customer_id",
					Require: plugin.Optional,
				},
			},
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"org_unit_id", "org_unit_path"}),
			Hydrate:    getDirectoryOrgUnit,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The organizational unit's path name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "org_unit_id",
				Description: "The unique ID of the organizational unit.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "org_unit_path",
				Description: "The full path to the organizational unit.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "block_inheritance",
				Description: "Determines if a sub-organizational unit can inherit the settings of the parent organization.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "customer_id",
				Description: "The customer ID to retrieve all account roles.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("customer_id"),
			},
			{
				Name:        "description",
				Description: "A short description of the organizational unit.",
				Type:        proto.ColumnType_STRING,
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
				Name:        "parent_org_unit_id",
				Description: "The unique ID of the parent organizational unit.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "parent_org_unit_path",
				Description: "The organizational unit's parent path.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryOrgUnits(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	// Set default value to my_customer, to represent current account
	customerID := "my_customer"
	if d.EqualsQuals["customer_id"] != nil {
		customerID = d.EqualsQuals["customer_id"].GetStringValue()
	}

	resp, err := service.Orgunits.List(customerID).Do()
	if err != nil {
		return nil, err
	}

	for _, orgUnit := range resp.OrganizationUnits {
		d.StreamListItem(ctx, orgUnit)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if plugin.IsCancelled(ctx) {
			break
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getDirectoryOrgUnit(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryOrgUnit")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	orgUnitID := d.EqualsQuals["org_unit_id"].GetStringValue()
	orgUnitPath := d.EqualsQuals["org_unit_path"].GetStringValue()

	// Return nil, if no input provided
	if orgUnitID == "" && orgUnitPath == "" {
		return nil, nil
	}

	var inputStr string
	if orgUnitID == "" {
		inputStr = orgUnitPath
	} else {
		inputStr = orgUnitID
	}

	resp, err := service.Orgunits.Get("my_customer", inputStr).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
