---
title: "Steampipe Table: googledirectory_domain - Query Google Directory Domains using SQL"
description: "Allows users to query Google Directory Domains, providing detailed information about the domain and its associated settings and configurations."
---

# Table: googledirectory_domain - Query Google Directory Domains using SQL

Google Directory Domains is a resource within Google Workspace that allows you to manage your organization's domains. It provides a centralized way to set up and manage domains, including domain verification, alias management, and more. Google Directory Domains helps you stay informed about the status and settings of your domains and take appropriate actions when needed.

## Table Usage Guide

The `googledirectory_domain` table provides insights into domains within Google Workspace Directory. As a system administrator, explore domain-specific details through this table, including domain name, whether the domain is verified, and associated metadata. Utilize it to uncover information about domains, such as their verification status, and to manage domain aliases.

## Examples

### Basic info
Explore which domains within your Google Directory are primary and when they were created. This can be beneficial for assessing domain configurations and understanding their establishment timeline.

```sql+postgres
select
  domain_name,
  creation_time,
  is_primary
from
  googledirectory_domain;
```

```sql+sqlite
select
  domain_name,
  creation_time,
  is_primary
from
  googledirectory_domain;
```

### List unverified domains
Discover the segments that include unverified domains in your Google Directory. This can help you identify potential security risks and take necessary actions to verify these domains.

```sql+postgres
select
  domain_name,
  creation_time,
  verified
from
  googledirectory_domain
where
  not verified;
```

```sql+sqlite
select
  domain_name,
  creation_time,
  verified
from
  googledirectory_domain
where
  not verified;
```