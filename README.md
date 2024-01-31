# testcontainers-go-demo

This repository demonstrates how you can use `testcontainers-go` with Dolt's built in module, [currently in review](https://github.com/testcontainers/testcontainers-go/pull/2177).

[Dolt](https://www.doltdb.com) is a SQL database that you can fork, clone, branch, merge, push and pull just like a Git repository.

This demo repository is modeled from the [Getting Started with TestContainers for Go](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/) document which
walks users through writing some tests using Golang and the PostgreSQL `testcontainer` API.

Here, we've used the same repository structure as defined in that document, but have used the Dolt `testcontainer` module API in place of PostgreSQL.

Additionally, this repository leverages two unique features of Dolt that makes it an excellent relational database for testing.

In `customer/repo_suite_test.go` the `BeforeTest()` method ensures each test is executed on a distinct database branch, guaranteeing that tests start with the same database state and
do not impact any other tests running against the Dolt container.

Similarly, in `customer/clone_suite_test.go`, the `SetupSuite()` shows how a Dolt container can use data hosted on a Dolt remote, like [DoltHub](https://www.dolthub.com), by cloning that data into the Dolt container before running tests.

## Running Tests

You can run all tests by cloning this repository then running:
```bash
go test ./...
```


