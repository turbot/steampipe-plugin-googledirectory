---
title: "Steampipe Table: googledirectory_group_member - Query Google Directory Group Members using SQL"
description: "Allows users to query Google Directory Group Members, specifically providing details about each member of a group, their roles, and type."
---

# Table: googledirectory_group_member - Query Google Directory Group Members using SQL

Google Directory is a service within Google Workspace that provides a centralized way to manage and organize users, groups, and devices in an organization. It allows administrators to manage access to services and delegate administrative tasks. Google Directory Group Member represents a member of a group within the Google Directory.

## Table Usage Guide

The `googledirectory_group_member` table provides insights into each member of a group within Google Directory. As an IT administrator, explore member-specific details through this table, including roles, type, and associated metadata. Utilize it to uncover information about group members, such as their roles within the group, the type of member (user, group, or service account), and other relevant details.

**Important Notes**
- You can specify the `group_id` optionally in the `where` clause to get the members associated with a group.

## Examples

### Basic info
Explore which roles are assigned to different members of a specific Google Directory group. This is useful for managing access permissions and ensuring the right individuals have the appropriate roles.

```sql+postgres
select
  group_id,
  id,
  email,
  role
from
  googledirectory_group_member
where
  group_id = '01ksv4uv1gexk1h';
```

```sql+sqlite
select
  group_id,
  id,
  email,
  role
from
  googledirectory_group_member
where
  group_id = '01ksv4uv1gexk1h';
```

### List all owners of a group
Discover the segments that have a specific ownership within a group. This can be useful for managing group permissions and understanding the distribution of roles within a group.

```sql+postgres
select
  group_id,
  id,
  email,
  role
from
  googledirectory_group_member
where
  group_id = '01ksv4uv1gexk1h'
  and role = 'OWNER';
```

```sql+sqlite
select
  group_id,
  id,
  email,
  role
from
  googledirectory_group_member
where
  group_id = '01ksv4uv1gexk1h'
  and role = 'OWNER';
```

### List role counts for a group
Explore which roles within a specific group have the highest membership count. This can help in understanding the distribution of roles within the group, allowing for better management and organization.

```sql+postgres
select
  role,
  count(*)
from
  googledirectory_group_member
where
  group_id = '01ksv4uv1gexk1h'
group by role
order by
  count desc;
```

```sql+sqlite
select
  role,
  count(*)
from
  googledirectory_group_member
where
  group_id = '01ksv4uv1gexk1h'
group by role
order by
  count(*) desc;
```

### List all groups and their members
Explore the relationships between various groups and their respective members to understand the structure and dynamics within your organization. This can be particularly useful for managing access permissions, coordinating team activities, or identifying communication patterns.

```sql+postgres
select
  g.id as group_id,
  g.name as group_name,
  m.email as member_email
from
  googledirectory_group as g,
  googledirectory_group_member as m
where
  g.id = m.group_id
order by
  g.name,
  m.email;
```

```sql+sqlite
select
  g.id as group_id,
  g.name as group_name,
  m.email as member_email
from
  googledirectory_group as g
join
  googledirectory_group_member as m on g.id = m.group_id
order by
  g.name,
  m.email;
```