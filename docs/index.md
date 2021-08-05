---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/googledirectory.svg"
brand_color: "#ea4335"
display_name: "Google Directory"
short_name: "googledirectory"
description: "Steampipe plugin for querying users, groups, org units and more from your Google Workspace directory."
og_description: "Query Google Workspace directory with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/googledirectory-social-graphic.png"
---

# Google Directory + Steampipe

[Google Directory](https://developers.google.com/admin-sdk/directory) is the users, groups, domains and other organizational features of Google Workspace. [Google Workspace](https://workspace.google.com) is a collection of cloud computing, productivity and collaboration tools, software and products developed and marketed by Google.

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
| Credentials | Generate your [service account and credentials](https://developers.google.com/admin-sdk/directory/v1/guides/delegation#create_the_service_account_and_credentials) and configure to [Delegate domain-wide authority to your service account](https://developers.google.com/admin-sdk/directory/v1/guides/delegation#delegate_domain-wide_authority_to_your_service_account). Enter the following OAuth 2.0 scopes for the services that the service account can access:<br />`https://www.googleapis.com/auth/admin.directory.user.readonly`<br />`https://www.googleapis.com/auth/admin.directory.domain.readonly`<br />`https://www.googleapis.com/auth/admin.directory.group.readonly`<br />`https://www.googleapis.com/auth/admin.directory.orgunit.readonly`<br />`https://www.googleapis.com/auth/admin.directory.rolemanagement.readonly` |
| Radius      | Each connection represents a single Google Workspace account. |
| Resolution  | 1. Credentials from the json file specified by the `credential_file` parameter in your steampipe config.<br />2. Credentials from the json file specified by the `GOOGLE_APPLICATION_CREDENTIALS` environment variable. |

### Configuration

Installing the latest googledirectory plugin will create a config file (`~/.steampipe/config/googledirectory.spc`) with a single connection named `googledirectory`:

```hcl
connection "googledirectory" {
  plugin = "googledirectory"

  # `credential_file` (required) - The path to a JSON credential file that contains service account credentials.
  # If `credential_file` is not specified in a connection, credentials will be loaded from:
  #  - The path specified in the `GOOGLE_APPLICATION_CREDENTIALS` environment variable, if set
  #credential_file = "/path/to/<public_key_fingerprint>-privatekey.json"

  # `impersonated_user_email` (required) - The email (string) of the user which should be impersonated. Needs permissions to access the Admin APIs.
  # `impersonated_user_email` must be set, since the service account needs to impersonate one of those users to access the directory.
  #impersonated_user_email = "username@domain.com"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-googledirectory
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)

## Advanced configuration options

By default, the plugin uses the credential file path provided in the connection config. You can also specify static credentials using environment variables, for example:

```sh
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/<public_key_fingerprint>-privatekey.json
```
