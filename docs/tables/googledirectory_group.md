---
title: "Steampipe Table: googledirectory_group - Query Google Directory Groups using SQL"
description: "Allows users to query Google Directory Groups, specifically the group details and members, providing insights into the structure and membership of groups within the Google Workspace."
---

# Table: googledirectory_group - Query Google Directory Groups using SQL

Google Directory is a service within Google Workspace that allows you to manage, create, and view groups and their members. It provides a centralized way to set up and manage groups for various Google Workspace resources, including users, emails, and more. Google Directory helps you stay informed about the organization and membership of your Google Workspace resources.

## Table Usage Guide

The `googledirectory_group` table provides insights into groups within Google Workspace. As a system administrator, explore group-specific details through this table, including group names, emails, and associated metadata. Utilize it to uncover information about groups, such as those with certain members, the hierarchy of groups, and the verification of group properties.

## Examples

### Basic info
Explore the basic information of Google Directory groups to gain insights into group names, IDs, associated emails, and creation details. This can be useful for managing and auditing group settings and memberships.

```sql+postgres
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group;
```

```sql+sqlite
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group;
```

### Get group by ID
Discover the details of a specific group in your Google Directory by using its unique ID. This can be useful for gaining insights into group information such as its name, email, and administrative creation data.

```sql+postgres
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group
where
  id = '02ce457p6conzyd';
```

```sql+sqlite
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group
where
  id = '02ce457p6conzyd';
```

### Get group by email
Determine the areas in which a specific email address is associated with a group, allowing you to understand the context and scope of that group's administration. This can be particularly useful for managing and auditing access permissions in a large organization.

```sql+postgres
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group
where
  email = 'scranton@dundermifflin.com';
```

```sql+sqlite
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group
where
  email = 'scranton@dundermifflin.com';
```

### List top 5 groups by member count
Explore the five most populated groups within your Google Directory. This could be useful for understanding which groups are most active or require the most resources.

```sql+postgres
select
  name,
  direct_members_count
from
  googledirectory_group
order by
  direct_members_count desc
limit 5;
```

```sql+sqlite
select
  name,
  direct_members_count
from
  googledirectory_group
order by
  direct_members_count desc
limit 5;
```

### List all groups and their members
Explore which members belong to specific groups within your Google Directory. This allows you to assess the composition of each group, aiding in tasks like group management and access control.

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
  googledirectory_group_member as m
on
  g.id = m.group_id
order by
  g.name,
  m.email;
```

### List groups using the [query filter](https://developers.google.com/admin-sdk/directory/v1/guides/search-groups)
Explore which groups have been created by admins within the Google Directory, specifically focusing on those associated with an email containing 'steampipe'. This can be beneficial in understanding the extent of 'steampipe' usage across different groups.

```sql+postgres
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group
where
  query = 'email:steampipe*';
```

```sql+sqlite
select
  name,
  id,
  email,
  admin_created
from
  googledirectory_group
where
  query = 'email:steampipe*';
```