---
title: "Steampipe Table: googledirectory_privilege - Query Google Directory Privileges using SQL"
description: "Allows users to query Google Directory Privileges, specifically providing insights into the various rights granted to administrative roles."
---

# Table: googledirectory_privilege - Query Google Directory Privileges using SQL

Google Directory Privileges is a resource within Google Workspace Admin SDK that manages and provides information about the various rights granted to administrative roles. It allows administrators to create, update, and delete roles that contain one or more privileges. It is a key component in managing access control within Google Workspace.

## Table Usage Guide

The `googledirectory_privilege` table provides insights into the privileges within Google Workspace Admin SDK. As an administrator, explore privilege-specific details through this table, including service IDs, privilege names, and associated metadata. Utilize it to uncover information about privileges, such as those associated with specific roles, and manage access control effectively within your Google Workspace environment.

## Examples

### Basic info
Explore which privileges within the Google Directory service are applicable to Organizational Units. This can aid in understanding the scope of access control and managing permissions effectively.

```sql
select
  privilege_name,
  service_name,
  service_id,
  is_ou_scopable
from
  googledirectory_privilege;
```

### List privileges by service
Explore the distribution of privileges across different services. This can help in assessing the security posture by identifying services with a high count of privileges.

```sql
select
  service_name,
  count(*)
from
  googledirectory_privilege
group by
  service_name
order by
  count desc;
```

### List privileges for each role
This example allows you to examine the specific permissions associated with each role within your Google Directory. It's useful for ensuring that roles are correctly configured and that each role has the appropriate level of access, enhancing your overall security posture.

```sql
select
  r.role_name as role_name,
  p.service_name as service_name,
  p.privilege_name as privilege_name
from
  googledirectory_role as r,
  jsonb_array_elements(r.role_privileges) as rp,
  googledirectory_privilege as p
where
  rp ->> 'serviceId' = p.service_id
  and rp ->> 'privilegeName' = p.privilege_name
order by
  role_name,
  service_name,
  privilege_name;
```