---
title: "Steampipe Table: googledirectory_org_unit - Query Google Directory Org Units using SQL"
description: "Allows users to query Google Directory Org Units, providing detailed information about organizational units within Google Workspace."
---

# Table: googledirectory_org_unit - Query Google Directory Org Units using SQL

Google Directory is a service within Google Workspace that manages and organizes information about users, groups, and devices. It provides a centralized way to manage organizational units, users, groups, and devices in a Google Workspace account. Google Directory helps you stay informed about the structure and organization of your Google Workspace resources.

## Table Usage Guide

The `googledirectory_org_unit` table provides insights into organizational units within Google Directory. As a system administrator, explore unit-specific details through this table, including names, descriptions, parent organizational units, and associated metadata. Utilize it to uncover information about the hierarchy and structure of your organization within Google Workspace.

## Examples

### Basic info
Explore the organization structure within Google Directory to understand its hierarchy and descriptions. This can be beneficial for managing resources and permissions within your organization.

```sql+postgres
select
  name,
  org_unit_id,
  org_unit_path,
  description
from
  googledirectory_org_unit;
```

```sql+sqlite
select
  name,
  org_unit_id,
  org_unit_path,
  description
from
  googledirectory_org_unit;
```

### Get org unit by ID
Explore the specific organizational unit within Google Directory by using its unique ID. This assists in obtaining detailed information about the unit, such as its name, path, and description, which can be useful for managing and understanding the structure of your organization.

```sql+postgres
select
  name,
  org_unit_id,
  org_unit_path,
  description
from
  googledirectory_org_unit
where
  org_unit_id = 'id:03pk8a4z4t34g1w';
```

```sql+sqlite
select
  name,
  org_unit_id,
  org_unit_path,
  description
from
  googledirectory_org_unit
where
  org_unit_id = 'id:03pk8a4z4t34g1w';
```

### Get org unit by path
Explore the specific organizational unit within your Google Directory by its unique path. This allows you to obtain crucial details about the unit, such as its name and description, which can be beneficial for managing your organizational structure.

```sql+postgres
select
  name,
  org_unit_id,
  org_unit_path,
  description
from
  googledirectory_org_unit
where
  org_unit_path = '/DM';
```

```sql+sqlite
select
  name,
  org_unit_id,
  org_unit_path,
  description
from
  googledirectory_org_unit
where
  org_unit_path = '/DM';
```