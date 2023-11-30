---
title: "Steampipe Table: googledirectory_domain_alias - Query Google Workspace Domain Aliases using SQL"
description: "Allows users to query Domain Aliases in Google Workspace, specifically providing insights into the aliases associated with a Google Workspace domain."
---

# Table: googledirectory_domain_alias - Query Google Workspace Domain Aliases using SQL

A Google Workspace Domain Alias is an alternative name for a Google Workspace domain, which allows users to log in to their accounts and services using different domain names. Domain aliases are particularly useful for organizations that operate under multiple brand names or have different domains for different departments. They are managed through the Google Admin console and can be used with all Google Workspace services.

## Table Usage Guide

The `googledirectory_domain_alias` table provides insights into domain aliases within Google Workspace. As a Google Workspace administrator, explore alias-specific details through this table, including the parent domain name, creation time, and whether the alias is verified. Utilize it to manage and monitor your organization's domain aliases, ensuring that all aliases are correctly set up and verified.

## Examples

### Basic info
Explore which domain aliases in your Google Directory have been verified and when they were created. This can be used to maintain a secure and organized domain structure.

```sql
select
  domain_alias_name,
  creation_time,
  verified
from
  googledirectory_domain_alias;
```

### List unverified domain aliases
Discover the segments that consist of unverified domain aliases, enabling you to identify potential areas of risk and take appropriate action to verify them.

```sql
select
  domain_alias_name,
  creation_time,
  verified
from
  googledirectory_domain_alias
where
  not verified;
```

### List domain aliases by parent domain
Explore the different domain aliases associated with a specific parent domain. This can be useful for understanding the structure and organization of your domain aliases, as well as for verifying their creation times and statuses.

```sql
select
  domain_alias_name,
  parent_domain_name,
  creation_time,
  verified
from
  googledirectory_domain_alias
where
  parent_domain_name = 'domain.com';
```