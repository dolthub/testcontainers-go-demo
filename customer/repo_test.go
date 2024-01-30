package customer

import "github.com/stretchr/testify/assert"

func (suite *CustomerRepoTestSuite) TestCreateCustomer() {
	t := suite.T()

	customer, err := suite.repository.CreateCustomer(suite.ctx, &Customer{
		Name:  "Henry",
		Email: "henry@gmail.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Id)
}

func (suite *CustomerRepoTestSuite) TestGetCustomerByEmail() {
	t := suite.T()

	customer, err := suite.repository.GetCustomerByEmail(suite.ctx, "john@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "John", customer.Name)
	assert.Equal(t, "john@gmail.com", customer.Email)
}

func (suite *CustomerRepoTestSuite) TestUpdateCustomer() {
	t := suite.T()

	customer, err := suite.repository.GetCustomerByEmail(suite.ctx, "john@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "John", customer.Name)
	assert.Equal(t, "john@gmail.com", customer.Email)

	id := customer.Id
	customer.Name = "JohnAlt"
	customer.Email = "john-alt@gmail.com"
	err = suite.repository.UpdateCustomer(suite.ctx, customer)
	assert.NoError(t, err)

	customer, err = suite.repository.GetCustomerByEmail(suite.ctx, "john-alt@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "JohnAlt", customer.Name)
	assert.Equal(t, "john-alt@gmail.com", customer.Email)
	assert.Equal(t, id, customer.Id)
}

func (suite *CustomerRepoTestSuite) TestDeleteCustomer() {
	t := suite.T()

	customer, err := suite.repository.GetCustomerByEmail(suite.ctx, "john@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "John", customer.Name)
	assert.Equal(t, "john@gmail.com", customer.Email)

	err = suite.repository.DeleteCustomer(suite.ctx, customer)
	assert.NoError(t, err)

	customer, err = suite.repository.GetCustomerByEmail(suite.ctx, "john@gmail.com")
	assert.NoError(t, err)
	assert.Nil(t, customer)
}
