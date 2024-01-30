package customer

import "github.com/stretchr/testify/assert"

func (suite *RemoteRepoTestSuite) TestCreateCustomer() {
	t := suite.T()

	customer, err := suite.Repository().CreateCustomer(suite.Ctx(), &Customer{
		Name:  "Henry",
		Email: "henry@gmail.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Id)
}

func (suite *RemoteRepoTestSuite) TestGetCustomerByEmail() {
	t := suite.T()

	customer, err := suite.Repository().GetCustomerByEmail(suite.Ctx(), "lisa@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Lisa", customer.Name)
	assert.Equal(t, "lisa@gmail.com", customer.Email)
}

func (suite *RemoteRepoTestSuite) TestUpdateCustomer() {
	t := suite.T()

	customer, err := suite.Repository().GetCustomerByEmail(suite.Ctx(), "vicki@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Vicki", customer.Name)
	assert.Equal(t, "vicki@gmail.com", customer.Email)

	id := customer.Id
	customer.Name = "Samantha"
	customer.Email = "samantha@gmail.com"
	err = suite.Repository().UpdateCustomer(suite.Ctx(), customer)
	assert.NoError(t, err)

	customer, err = suite.Repository().GetCustomerByEmail(suite.Ctx(), "samantha@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Samantha", customer.Name)
	assert.Equal(t, "samantha@gmail.com", customer.Email)
	assert.Equal(t, id, customer.Id)
}

func (suite *RemoteRepoTestSuite) TestDeleteCustomer() {
	t := suite.T()

	customer, err := suite.Repository().GetCustomerByEmail(suite.Ctx(), "megan@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Megan", customer.Name)
	assert.Equal(t, "megan@gmail.com", customer.Email)

	err = suite.Repository().DeleteCustomer(suite.Ctx(), customer)
	assert.NoError(t, err)

	customer, err = suite.Repository().GetCustomerByEmail(suite.Ctx(), "megan@gmail.com")
	assert.NoError(t, err)
	assert.Nil(t, customer)
}
