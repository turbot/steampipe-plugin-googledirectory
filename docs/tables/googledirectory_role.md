---
title: "Steampipe Table: googledirectory_role - Query Google Directory Roles using SQL"
description: "Allows users to query Google Directory Roles, specifically the details about roles within Google Workspace. This includes role ID, role name, role description, and associated privileges."
---

# Table: googledirectory_role - Query Google Directory Roles using SQL

Google Directory is a service within Google Cloud that allows you to manage your organization's users, groups, and devices. It provides a centralized way to set up and manage roles for various Google Workspace resources. Google Directory helps you stay informed about the roles and their associated privileges within your Google Workspace.

## Table Usage Guide

The `googledirectory_role` table provides insights into roles within Google Workspace. As a Google Workspace administrator, explore role-specific details through this table, including role ID, role name, role description, and associated privileges. Utilize it to uncover information about roles, such as their privileges and the details associated with each role.

## Examples

### Basic info
Analyze the settings to understand the roles within your Google Directory, specifically identifying which roles have super admin or system privileges. This can be useful for auditing access rights and maintaining security within your organization.

```sql+postgres
select
  role_name,
  role_id,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role;
```

```sql+sqlite
select
  role_name,
  role_id,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role;
```

### Get role by ID
Explore which Google Directory roles possess certain identifiers, enabling you to pinpoint specific roles for administrative or system purposes. This is useful in managing user access and permissions in your Google Directory.

```sql+postgres
select
  role_name,
  role_id,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role
where
  role_id = '02ce457p6conzyd';
```

```sql+sqlite
select
  role_name,
  role_id,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role
where
  role_id = '02ce457p6conzyd';
```

### List super admin roles
Explore which roles hold super admin privileges in your Google Directory, to manage permissions and secure your system effectively. This query helps you identify those roles, providing valuable information for system administration and security.

```sql+postgres
select
  role_id,
  role_name,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role
where
  is_super_admin_role;
```

```sql+sqlite
select
  role_id,
  role_name,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role
where
  is_super_admin_role = 1;
```

### List system roles
Discover the segments that identify all system roles in the Google Directory, providing a way to assess which roles have super admin privileges. This can be beneficial for auditing purposes or to manage user permissions effectively.

```sql+postgres
select
  role_id,
  role_name,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role
where
  is_system_role;
```

```sql+sqlite
select
  role_id,
  role_name,
  is_super_admin_role,
  is_system_role
from
  googledirectory_role
where
  is_system_role = 1;
```

### List privileges by role
Explore which privileges are associated with each role in Google Directory. This can be useful in managing access control and ensuring that each role has the correct privileges for its intended function.

```sql+postgres
select
  role_name,
  p ->> 'serviceId' as service_id,
  p ->> 'privilegeName' as privilege
from
  googledirectory_role as r,
  jsonb_array_elements(r.role_privileges) as p
order by
  role_name,
  service_id,
  privilege;
```

```sql+sqlite
select
  role_name,
  json_extract(p.value, '$.serviceId') as service_id,
  json_extract(p.value, '$.privilegeName') as privilege
from
  googledirectory_role as r,
  json_each(r.role_privileges) as p
order by
  role_name,
  service_id,
  privilege;
```