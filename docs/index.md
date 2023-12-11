---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/googledirectory.svg"
brand_color: "#1967D2"
display_name: "Google Directory"
short_name: "googledirectory"
description: "Steampipe plugin for querying users, groups, org units and more from your Google Workspace directory."
og_description: "Query Google Workspace directory with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/googledirectory-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Google Directory + Steampipe

A [Google Directory](https://developers.google.com/admin-sdk/directory) contains the users, groups, domains and other organizational features of a Google Workspace. [Google Workspace](https://workspace.google.com) is a collection of cloud computing, productivity and collaboration tools, software and products developed and marketed by Google.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  full_name,
  primary_email
from
  googledirectory_user;
```

```
+----------------+----------------------------+
| full_name      | primary_email              |
+----------------+----------------------------+
| Dwight Schrute | dschrute@dundermifflin.com |
| Michael Scott  | mscott@dundermifflin.com   |
| Pam Beesly     | pbeesly@dundermifflin.com  |
+----------------+----------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/googledirectory/tables)**

## Get started

### Install

Download and install the latest Google Directory plugin:

```bash
steampipe plugin install googledirectory
```

### Credentials

| Item        | Description |
| :---------- | :-----------|
| Credentials | 1. To use **domain-wide delegation**, generate your [service account and credentials](https://developers.google.com/admin-sdk/directory/v1/guides/delegation#create_the_service_account_and_credentials) and [delegate domain-wide authority to your service account](https://developers.google.com/admin-sdk/directory/v1/guides/delegation#delegate_domain-wide_authority_to_your_service_account). Enter the following OAuth 2.0 scopes for the services that the service account can access:<br />`https://www.googleapis.com/auth/admin.directory.domain.readonly`<br />`https://www.googleapis.com/auth/admin.directory.group.readonly`<br />`https://www.googleapis.com/auth/admin.directory.orgunit.readonly`<br />`https://www.googleapis.com/auth/admin.directory.rolemanagement.readonly`<br />`https://www.googleapis.com/auth/admin.directory.user.readonly`<br />2. To use **OAuth client**, configure your [credentials](#authenticate-using-oauth-client). |
| Radius      | Each connection represents a single Google Workspace account. |
| Resolution  | 1. Credentials from the JSON file specified by the `credentials` parameter in your Steampipe config.<br />2. Credentials from the JSON file specified by the `token_path` parameter in your Steampipe config.<br />3. Credentials from the default json file location (`~/.config/gcloud/application_default_credentials.json`). |

### Configuration

Installing the latest googledirectory plugin will create a config file (`~/.steampipe/config/googledirectory.spc`) with a single connection named `googledirectory`:

```hcl
connection "googledirectory" {
  plugin = "googledirectory"

  # You may connect to Google Workspace using more than one option:
  # 1. To authenticate using domain-wide delegation, specify a service account credential file and the user email for impersonation
  # `credentials` - Either the path to a JSON credential file that contains Google application credentials,
  # or the contents of a service account key file in JSON format. If `credentials` is not specified in a connection,
  # credentials will be loaded from:
  #   - The path specified in the `GOOGLE_APPLICATION_CREDENTIALS` environment variable, if set; otherwise
  #   - The standard location (`~/.config/gcloud/application_default_credentials.json`)
  #   - The path specified for the credentials.json file ("/path/to/my/creds.json")
  # credentials = "~/.config/gcloud/application_default_credentials.json"

  # `impersonated_user_email` - The email (string) of the user which should be impersonated. Needs permissions to access the Admin APIs.
  # `impersonated_user_email` must be set, since the service account needs to impersonate a user with Admin API permissions to access the directory.
  # impersonated_user_email = "username@domain.com"

  # 2. To authenticate using OAuth 2.0, specify a client secret file
  # `token_path` - The path to a JSON credential file that contains Google application credentials.
  # If `token_path` is not specified in a connection, credentials will be loaded from:
  #  - The path specified in the `GOOGLE_APPLICATION_CREDENTIALS` environment variable, if set; otherwise
  #  - The standard location (`~/.config/gcloud/application_default_credentials.json`)
  # token_path = "~/.config/gcloud/application_default_credentials.json"
}
```

## Advanced configuration options

### Authenticate using OAuth client

You can use client secret credentials to protect the user's data by only granting tokens to authorized requestors. Use following steps to configure credentials:

- [Configure the OAuth consent screen](https://developers.google.com/workspace/guides/configure-oauth-consent).
- [Create an OAuth client ID credential](https://developers.google.com/workspace/guides/create-credentials#desktop-app) with the application type `Desktop app`, and download the client secret JSON file.
- Wherever you have the [Google Cloud SDK](https://cloud.google.com/sdk/docs/install) installed, run the following command with the correct client secret JSON file parameters:

  ```sh
  gcloud auth application-default login \
    --client-id-file=client_secret.json \
    --scopes="\
  https://www.googleapis.com/auth/admin.directory.domain.readonly,\
  https://www.googleapis.com/auth/admin.directory.group.readonly,\
  https://www.googleapis.com/auth/admin.directory.orgunit.readonly,\
  https://www.googleapis.com/auth/admin.directory.rolemanagement.readonly,\
  https://www.googleapis.com/auth/admin.directory.user.readonly"
  ```

- In the browser window that just opened, authenticate as the user you would like to make the API calls through.
- Review the output for the location of the **Application Default Credentials** file, which usually appears following the text `Credentials saved to file:`.
- Set the **Application Default Credentials** filepath in the Steampipe config `token_path` or in the `GOOGLE_APPLICATION_CREDENTIALS` environment variable.
