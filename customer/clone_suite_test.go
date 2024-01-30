package customer

import (
	"context"
	"log"
	"testing"

	"github.com/dolthub/testcontainers-go-demo/testhelpers"
	"github.com/stretchr/testify/suite"
)

type RemoteRepoTestSuite struct {
	c *CustomerRepoTestSuite
}

func (suite *RemoteRepoTestSuite) T() *testing.T {
	return suite.c.T()
}

func (suite *RemoteRepoTestSuite) SetT(t *testing.T) {
	suite.c.SetT(t)
}

func (suite *RemoteRepoTestSuite) SetS(s suite.TestingSuite) {
	suite.c.SetS(s)
}

func (suite *RemoteRepoTestSuite) Ctx() context.Context {
	return suite.c.ctx
}

func (suite *RemoteRepoTestSuite) Repository() *Repository {
	return suite.c.repository
}

func (suite *RemoteRepoTestSuite) SetupSuite() {
	suite.c.ctx = context.Background()
	doltContainer, err := testhelpers.CreateDoltContainerFromClone(suite.c.ctx)
	if err != nil {
		log.Fatal(err)
	}

	suite.c.doltContainer = doltContainer
	if suite.c.startRef == "" {
		suite.c.startRef = "main"
	}
}

func (suite *RemoteRepoTestSuite) TearDownSuite() {
	suite.c.TearDownSuite()
}

func (suite *RemoteRepoTestSuite) AfterTest(suiteName, testName string) {
	suite.c.AfterTest(suiteName, testName)
}

func (suite *RemoteRepoTestSuite) BeforeTest(suiteName, testName string) {
	suite.c.BeforeTest(suiteName, testName)
}

func TestRemoteRepoTestSuite(t *testing.T) {
	s := &RemoteRepoTestSuite{c: new(CustomerRepoTestSuite)}
	suite.Run(t, s)
}
