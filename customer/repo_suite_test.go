package customer

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
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

func (suite *CustomerRepoTestSuite) createBranchName(suiteName, testName string) string {
	return fmt.Sprintf("%s_%s", suiteName, testName)
}

func (suite *CustomerRepoTestSuite) AfterTest(suiteName, testName string) {
	if suite.db != nil {
		branchName := suite.createBranchName(suiteName, testName)
		_, err := suite.db.ExecContext(suite.ctx, "CALL DOLT_COMMIT('-Am', ?)", fmt.Sprintf("Finished testing on %s", branchName))
		if err != nil {
			if !strings.Contains(err.Error(), "nothing to commit") {
				log.Fatal(err)
			}
		}
		err = suite.db.Close()
		if err != nil {
			log.Fatal(err)
		}
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

	newBranch := suite.createBranchName(suiteName, testName)

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
