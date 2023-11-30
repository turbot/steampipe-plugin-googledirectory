---
title: "Steampipe Table: googledirectory_role_assignment - Query Google Directory Role Assignments using SQL"
description: "Allows users to query Role Assignments in Google Directory, providing insights into role assignments and their details."
---

# Table: googledirectory_role_assignment - Query Google Directory Role Assignments using SQL

Google Directory is a service within Google Workspace that helps manage organizational structure and browse people in your organization. It allows you to manage users, devices, and apps, and it's an essential tool for IT and system administrators. Role Assignments in Google Directory are used to assign roles to users or groups, which define what actions they can perform.

## Table Usage Guide

The `googledirectory_role_assignment` table provides insights into Role Assignments within Google Directory. As an IT or system administrator, explore role assignment-specific details through this table, including the assigned user or group, the role ID, and the assignment ID. Utilize it to uncover information about role assignments, such as the permissions associated with each role, the users or groups assigned to each role, and the scope of each assignment.

## Examples

### Basic info
Explore the allocation of roles within your Google Directory setup. This query will help you understand who holds what role and where, enhancing your security management by identifying potential misassignments or gaps.

```sql
select
  role_assignment_id,
  role_id,
  assigned_to,
  scope_type
from
  googledirectory_role_assignment;
```

### Get role assignments by role ID
Explore which roles have been assigned to different users within a specific Google Directory role. This can be useful in managing access and permissions in your organization.

```sql
select
  role_assignment_id,
  role_id,
  assigned_to,
  scope_type
from
  googledirectory_role_assignment
where
  role_id = '522363132560015';
```

### Get role assignments by user
Explore which roles have been assigned to each user in the Google Directory. This can be useful to understand the permissions and access each user has within the organization.

```sql
select
  assigned_role.role_assignment_id as role_assignment_id,
  r.role_name as role_name,
  u.full_name as user_name
from
  googledirectory_role_assignment as assigned_role,
  googledirectory_user as u,
  googledirectory_role as r
where
  assigned_role.user_key = u.id
  and assigned_role.role_id = r.role_id;
```