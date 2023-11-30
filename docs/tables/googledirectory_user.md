---
title: "Steampipe Table: googledirectory_user - Query Google Directory Users using SQL"
description: "Allows users to query Google Directory Users, specifically retrieving detailed information about user accounts within the Google Workspace domain."
---

# Table: googledirectory_user - Query Google Directory Users using SQL

Google Directory is a service within Google Workspace that provides a centralized way to manage and access user account information. It allows administrators to manage users, groups, and devices, as well as to configure security settings for the domain. Google Directory helps to maintain the integrity of the domain's data by providing a structured way to manage user account information.

## Table Usage Guide

The `googledirectory_user` table provides insights into user accounts within Google Workspace. As an IT administrator, explore user-specific details through this table, including email addresses, names, and administrative status. Utilize it to uncover information about users, such as their last login time, whether their account is suspended, and the organizational units to which they belong.

## Examples

### Basic info
Explore which users have administrative privileges in your Google Directory and when they were created. This can be useful for auditing purposes and ensuring that only authorized individuals have admin access.

```sql
select
  full_name,
  id,
  primary_email,
  creation_time,
  is_delegated_admin,
  customer_id
from
  googledirectory_user;
```

### Get user by ID
Discover the details of a specific user in the Google Directory, such as their full name, primary email, and creation time. This can be useful for administrators who need to verify user information or investigate account activity.

```sql
select
  full_name,
  id,
  primary_email,
  creation_time,
  is_delegated_admin,
  customer_id
from
  googledirectory_user
where
  id = '119982672925259996273';
```

### Get user by primary email
Discover the details of a specific user by using their primary email. This can be particularly useful for gaining insights into user's profile details, creation time, and customer ID in a business context.

```sql
select
  full_name,
  id,
  primary_email,
  creation_time,
  is_delegated_admin,
  customer_id
from
  googledirectory_user
where
  primary_email = 'mscott@dundermifflin.com';
```

### List administrators
Discover the users who hold administrative or delegated administrative roles in your Google Directory. This can be useful for auditing access control and ensuring only authorized individuals have elevated permissions.

```sql
select
  id,
  full_name,
  primary_email,
  is_admin,
  is_delegated_admin
from
  googledirectory_user
where
  is_admin
  or is_delegated_admin;
```

### List users without two-step verification
Discover the segments that have users who haven't enabled two-step verification. This can be beneficial for enhancing the security measures within your organization.

```sql
select
  id,
  full_name,
  primary_email,
  is_enrolled_in_2sv,
  is_enforced_in_2sv
from
  googledirectory_user
where
  not is_enrolled_in_2sv
  or not is_enforced_in_2sv;
```

### List users who have not logged in for more than 30 days
The query is used to identify users who have been inactive for over a month. This can be useful for IT administrators to manage user accounts and security, by potentially flagging these accounts for follow-up or deactivation.

```sql
select
  id,
  full_name,
  primary_email,
  last_login_time
from
  googledirectory_user
where
  last_login_time < current_timestamp - interval '30 days';
```

### List users using the [query filter](https://developers.google.com/admin-sdk/directory/v1/guides/search-users)
Discover the segments that include users with a specific attribute in their name. This is useful in scenarios where you need to identify and group users based on shared characteristics for targeted communication or management.

```sql
select
  id,
  full_name,
  primary_email,
  last_login_time
from
  googledirectory_user
where
  query = 'givenName:steampipe*';
```