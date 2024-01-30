package customer

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/dolthub/testcontainers-go-demo/testhelpers"
	"github.com/stretchr/testify/suite"
)

type CustomerRepoTestSuite struct {
	suite.Suite
	doltContainer *testhelpers.DoltContainer
	repository    *Repository
	db            *sql.DB
	ctx           context.Context
	startRef      string
}

func (suite *CustomerRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	doltContainer, err := testhelpers.CreateDoltContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}

	suite.doltContainer = doltContainer
	if suite.startRef == "" {
		suite.startRef = "main"
	}
}

func (suite *CustomerRepoTestSuite) TearDownSuite() {
	if err := suite.doltContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating dolt container: %s", err)
	}
}

func (suite *CustomerRepoTestSuite) AfterTest(suiteName, testName string) {
	if suite.db != nil {
		suite.db.Close()
		suite.db = nil
		suite.repository = nil
	}
}

func (suite *CustomerRepoTestSuite) BeforeTest(suiteName, testName string) {
	// connect to the dolt server
	db, err := sql.Open("mysql", suite.doltContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	newBranch := fmt.Sprintf("%s_%s", suiteName, testName)

	// checkout new branch for test from the designated starting point
	_, err = db.ExecContext(suite.ctx, "CALL DOLT_CHECKOUT(?, '-b', ?);", suite.startRef, newBranch)
	if err != nil {
		log.Fatal(err)
	}

	suite.db = db

	repository, err := NewRepository(suite.ctx, suite.db)
	if err != nil {
		log.Fatal(err)
	}

	suite.repository = repository
}

func TestCustomerRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepoTestSuite))
}
