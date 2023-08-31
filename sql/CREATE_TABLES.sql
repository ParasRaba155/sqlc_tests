CREATE TABLE IF NOT EXISTS tenants (
    id UUID DEFAULT gen_random_uuid () PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status BOOLEAN NOT NULL,
    code VARCHAR(10),
    size INTEGER,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS employees (
    id UUID DEFAULT gen_random_uuid () PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    tenant_id UUID NOT NULL,
    manager_id UUID,
    FOREIGN KEY (manager_id) REFERENCES employees(id),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id)
);


CREATE TABLE IF NOT EXISTS departments (
    id SMALLSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(100) NOT NULL,
    tenant_id UUID NOT NULL,
    FOREIGN KEY (tenant_id) REFERENCES tenants(id)
);

-- NOTE: HERE the type of department_id is set as SMALLINT
-- this is because SMALLSERIAL types are behind the scenes just
-- SMALLINT types but postgres auto increments them
-- if we were to add here the SMALLSERIAL type than we can not set it at all
ALTER TABLE employees
ADD COLUMN IF NOT EXISTS
department_id SMALLINT DEFAULT NULL REFERENCES departments(id);