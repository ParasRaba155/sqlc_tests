package queries

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

// should use this once to create tenanats
func CreateTenanatsOnce(q Querier) {
	tenants := []CreateTenantsParams{
		{
			Name:   "botree technologies",
			Status: true,
			Code: pgtype.Text{
				String: "bot",
				Valid:  true,
			},
		},
		{
			Name:   "insync operations",
			Status: true,
			Code: pgtype.Text{
				String: "ins",
				Valid:  true,
			},
		},
		{
			Name:   "tntra techonologies",
			Status: true,
			Code: pgtype.Text{
				String: "tix",
				Valid:  true,
			},
		},
	}
	rows, err := q.CreateTenants(context.TODO(), tenants)
	if err != nil {
		log.Printf("could not create tenants: %s", err)
		return
	}
	log.Printf("number of tenants created are: %d", rows)
}

func CreateEmployeesOnce(q Querier) {
	arr := generateRandomEmployee(q)
	rows, err := q.CreateEmployees(context.TODO(), arr[:])
	if err != nil {
		log.Printf("could not create employees: %s", err)
		return
	}
	log.Printf("number of employee created are: %d", rows)
}

func CreateDepartmentsOnce(q Querier) {
	depts := []CreateDepartmentParams{}
	ids, err := q.GetTentantID(context.TODO())
	if err != nil {
		log.Printf("could not get tenant: %s", err)
		return
	}

	for i := range ids {
		d := CreateDepartmentParams{
			Name:     "sales",
			Code:     "SL",
			TenantID: ids[i],
		}
		depts = append(depts, d)
	}

	for i := range ids {
		d := CreateDepartmentParams{
			Name:     "marketing",
			Code:     "MK",
			TenantID: ids[i],
		}
		depts = append(depts, d)
	}

	for i := range depts {
		_, err := q.CreateDepartment(context.TODO(), depts[i])
		if err != nil {
			log.Printf("could not create department: %s", err)
			return
		}
	}
}

func UpdateEmployeesDepartmentOnce(q Querier) {
	emails := readFile("./emails.txt")
	deptsID, err := q.GetDepartmentID(context.TODO())
	if err != nil {
		log.Printf("could not get all departments: %s", err)
	}
	for i := range emails {
		arg := UpdateEmployeFromEmailParams{
			DepartmentID: pgtype.Int2{
				Int16: getRandomElelmentFromList(deptsID),
				Valid: true,
			},
			Email: emails[i],
		}
		_, err := q.UpdateEmployeFromEmail(context.TODO(), arg)
		if err != nil {
			log.Printf("could not get all departments: %s", err)
		}
	}
}
