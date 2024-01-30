package testhelpers

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/dolt"
	"path/filepath"
)

type DoltContainer struct {
	*dolt.DoltContainer
	ConnectionString string
}

func CreateDoltContainer(ctx context.Context) (*DoltContainer, error) {
	doltContainer, err := dolt.RunContainer(ctx,
		testcontainers.WithImage("dolthub/dolt-sql-server:latest"),
		dolt.WithScripts(filepath.Join("..", "testdata", "init-db.sql")),
		dolt.WithDatabase("test_db"),
		dolt.WithUsername("tester"),
		dolt.WithPassword("testing"),
	)
	if err != nil {
		return nil, err
	}

	connStr, err := doltContainer.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}

	return &DoltContainer{
		DoltContainer:    doltContainer,
		ConnectionString: connStr,
	}, nil
}

func CreateDoltContainerFromClone(ctx context.Context) (*DoltContainer, error) {
	doltContainer, err := dolt.RunContainer(ctx,
		testcontainers.WithImage("dolthub/dolt-sql-server:latest"),
		dolt.WithScripts(filepath.Join("..", "testdata", "clone-db.sh")),
		dolt.WithDatabase("test_db"),
		dolt.WithUsername("tester"),
		dolt.WithPassword("testing"),
		dolt.WithDoltCloneRemoteUrl("https://doltremoteapi.dolthub.com/dolthub/testcontainers-go-demo"),
	)
	if err != nil {
		return nil, err
	}

	connStr, err := doltContainer.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}

	return &DoltContainer{
		DoltContainer:    doltContainer,
		ConnectionString: connStr,
	}, nil
}
