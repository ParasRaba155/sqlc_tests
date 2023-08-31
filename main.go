package main

import (
	"context"
	"fmt"
	"log"

	queries "github.com/ParasRaba155/sqlc_tests/sql"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

const db_url = `postgres://sqlc-test:pass@localhost:5432/sqlc-test?connect_timeout=5`

func main() {
	pool, err := Connect()
	if err != nil {
		log.Fatalf("error in connecting: %s", err)
	}
	querier := queries.New(pool)

	/////// the below function must only be called first time
	/////// for preparing/seeding the database and then it shouldn't
	/////// be called
	// // queries.CreateTenanatsOnce(querier)
	// // queries.CreateDepartmentsOnce(querier)
	// // queries.CreateEmployeesOnce(querier)
	// // queries.UpdateEmployeesDepartmentOnce(querier)
	//////// comment ends here

	botreeID := uuid.MustParse("55aff191-650a-4a54-a5e6-52368bb4a95b")

	// here starts join experimentation
	emp, err := querier.GetEmployeeWithGivenDeptCodeInTenant(
		context.TODO(),
		queries.GetEmployeeWithGivenDeptCodeInTenantParams{
			TenantID: pgtype.UUID{
				Bytes: botreeID,
				Valid: true,
			},
			Code: "SL",
		})
	if err != nil {
		log.Fatalf("error in botree employee getting: %s", err)
	}

	log.Printf("length from first query: %d", len(emp))

	count, err := querier.GetCountEmployeeWithGivenDeptCodeInTenant(
		context.TODO(),
		queries.GetCountEmployeeWithGivenDeptCodeInTenantParams{
			TenantID: pgtype.UUID{
				Bytes: botreeID,
				Valid: true,
			},
			Code: "SL",
		})

    if err != nil {
		log.Fatalf("error in botree employee getting count: %s", err)
    }

	log.Printf("length from 2nd query: %d", count)
}

func Connect() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.TODO(), db_url)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error connecting to db : %w", err)
	}
	log.Printf("database connected successfully")
	return pool, nil
}
