package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableGoogleDirectroryMobileDevice(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_mobile_device",
		Description: "Mobile Devices defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryMobileDevices,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "customer_id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("resource_id"),
			Hydrate:    getDirectoryMobileDevice,
		},
		Columns: []*plugin.Column{
			{
				Name:        "model",
				Description: "The mobile device's model name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource_id",
				Description: "The unique ID the API service uses to identify the mobile device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "device_id",
				Description: "The serial number for a Google Sync mobile device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "brand",
				Description: "Mobile Device Brand.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "imei",
				Description: "The device's IMEI number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "adb_status",
				Description: "Adb (USB debugging) enabled or disabled on device.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "baseband_version",
				Description: "The device's baseband version.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "bootloader_version",
				Description: "Mobile Device Bootloader version.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "build_number",
				Description: "The device's operating system build number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "customer_id",
				Description: "The customer ID to retrieve all account roles.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("customer_id"),
			},
			{
				Name:        "default_language",
				Description: "The default locale used on the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "developer_options_status",
				Description: "Developer options enabled or disabled on device.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "device_compromised_status",
				Description: "The compromised device status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "device_password_status",
				Description: "The device password status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "encryption_status",
				Description: "Mobile Device Encryption Status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "etag",
				Description: "A hash of the metadata, used to ensure there were no concurrent modifications to the resource when attempting an update.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "first_sync",
				Description: "Date and time the device was first synchronized with the policy settings in the G Suite administrator control panel.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "hardware",
				Description: "Mobile Device Hardware.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "hardware_ud",
				Description: "The IMEI/MEID unique identifier for Android hardware.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "kernel_version",
				Description: "The device's kernel version.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "kind",
				Description: "The type of the API resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_sync",
				Description: "Date and time the device was last synchronized with the policy settings in the G Suite administrator control panel.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "managed_account_is_on_owner_profile",
				Description: "Boolean indicating if this account is on owner/primary profile or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "manufacturer",
				Description: "Mobile Device manufacturer.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "meid",
				Description: "The device's MEID number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "network_operator",
				Description: "Mobile Device mobile or network operator.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "os",
				Description: "The mobile device's operating system.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "privilege",
				Description: "Mobile Device DMAgentPermission.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "release_version",
				Description: "Mobile Device release version version.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "security_patch_level",
				Description: "Mobile Device Security patch level.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "serial_number",
				Description: "The device's serial number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The device's status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "supports_work_profile",
				Description: "Work profile supported on device.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "type",
				Description: "The type of mobile device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "unknown_sources_status",
				Description: "Unknown sources enabled or disabled on device.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "user_agent",
				Description: "Gives information about the device such as `os` version.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "wifi_mac_address",
				Description: "The device's MAC address on Wi-Fi networks.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "applications",
				Description: "The list of applications installed on an Android mobile device.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "email",
				Description: "List of owner's email addresses.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "name",
				Description: "List of the owner's user names.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "other_accounts_info",
				Description: "List of accounts added on device.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryMobileDevices(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
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

	resp, err := service.Mobiledevices.List(customerID).Do()
	if err != nil {
		return nil, err
	}

	for _, mobiledevice := range resp.Mobiledevices {
		d.StreamListItem(ctx, mobiledevice)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getDirectoryMobileDevice(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryMobileDevice")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	resourceId := d.KeyColumnQuals["resource_id"].GetStringValue()

	// Return nil, if no input provided
	if resourceId == "" {
		return nil, nil
	}

	resp, err := service.Mobiledevices.Get("my_customer", resourceId).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
