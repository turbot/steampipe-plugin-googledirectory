# Table: googledirectory_privilege

Query information about privileges defined in the Google Workspace.

## Examples

### List all privileges

```sql
select
  privilege_name,
  service_name,
  service_id,
  is_ou_scopable
from
  googledirectory_privilege;
```

### Privilege count by service

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

### Privileges by role and service

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

### Roles by privilege service

```sql
select
  p.service_name as service_name,
  p.privilege_name as privilege_name,
  r.role_name as role_name
from
  googledirectory_role as r,
  jsonb_array_elements(r.role_privileges) as rp,
  googledirectory_privilege as p
where
  rp ->> 'serviceId' = p.service_id
  and rp ->> 'privilegeName' = p.privilege_name
order by
  service_name,
  privilege_name,
  role_name;
```
