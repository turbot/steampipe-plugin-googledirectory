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

[Google Workspace](https://workspace.google.com) is a collection of cloud computing, productivity and collaboration tools, software and products developed and marketed by Google. Google Directory is the users, groups, domains and other organizational features of Google Workspace.

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

| Item        | Description                                                                                                                         |
| :---------- | :-----------------------------------------------------------------------------------------------------------------------------------|
| Credentials | Google Directory requires a [Private key](https://developers.google.com/admin-sdk/directory/v1/guides/delegation) for all requests. |
| Radius      | Each connection represents a single Google Workspace account.                                                                       |

### Configuration

Installing the latest googledirectory plugin will create a config file (`~/.steampipe/config/googledirectory.spc`) with a single connection named `googledirectory`:

```hcl
connection "googledirectory" {
  plugin           = "googledirectory"
  credential_file  = "~/my-service-account-creds.json"
  impersonate_user = "username@dmain.com"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-googledirectory
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
