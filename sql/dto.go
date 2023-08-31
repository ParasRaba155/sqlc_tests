package queries

import "fmt"

func (c GetEmployeeWithGivenDeptCodeInTenantRow) String() string {
	return fmt.Sprintf(
		"id = %s, email = %s, user_name = %s tenant_id = %s department_id = %d code = %s dept_name = %s",
		c.ID.Bytes,
		c.Email,
		c.UserName,
		c.TenantID.Bytes,
		c.DepartmentID.Int16,
		c.Code,
		c.Name,
	)
}
