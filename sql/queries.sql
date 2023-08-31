-- name: CreateEmployee :one
INSERT INTO employees (
    email, user_name, tenant_id, department_id
) VALUES ( $1, $2, $3, $3)
RETURNING *;

-- name: CreateTenant :one
INSERT INTO tenants (
    name, status, code 
) VALUES ( $1, $2, $3)
RETURNING *;

-- name: CreateDepartment :one
INSERT INTO departments (
    name, code, tenant_id
) VALUES ( $1, $2, $3)
RETURNING *;

-- name: CreateTenants :copyfrom
INSERT INTO tenants 
(name, status, code) VALUES ( $1, $2, $3);

-- name: CreateEmployees :copyfrom
INSERT INTO employees 
(email, user_name, tenant_id, department_id) VALUES ( $1, $2, $3, $4);

-- name: GetTenantIDFromName :one
SELECT id FROM tenants
WHERE name = $1;

-- name: GetTentantID :many
SELECT id FROM tenants;

-- name: UpdateEmployeFromEmail :one
UPDATE employees
SET
    user_name = COALESCE(sqlc.narg('user_name'),user_name),
    manager_id = COALESCE(sqlc.narg('manager_id'),manager_id),
    department_id = COALESCE(sqlc.narg('department_id'),department_id)
WHERE email = sqlc.arg('email')
RETURNING *;

-- name: GetDepartmentID :many
SELECT id FROM departments;

-- name: GetEmployeeWithGivenDeptCodeInTenant :many
SELECT
    employees.id,
    employees.email,
    employees.user_name,
    employees.tenant_id,
    employees.manager_id,
    employees.department_id,
    -- departments.id, -- we don't need to get the dept id 
    departments.name,
    departments.code
FROM employees
INNER JOIN departments ON departments.id = employees.department_id
WHERE 
employees.tenant_id = $1 AND
departments.code = $2;

-- name: GetCountEmployeeWithGivenDeptCodeInTenant :one
SELECT COUNT (*)
FROM employees
INNER JOIN departments ON departments.id = employees.department_id
WHERE 
employees.tenant_id = $1 AND
departments.code = $2;
    
