package testutils

import (
	"net"
	"strconv"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func getFreePort() (port int, err error) {
	var a *net.TCPAddr
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return l.Addr().(*net.TCPAddr).Port, nil
		}
	}
	return
}

var (
	pool     *dockertest.Pool
	resource *dockertest.Resource
)

func SetupDatabase(t *testing.T) int {
	var err error
	pool, err = dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}

	port, err := getFreePort()
	if err != nil {
		t.Fatalf("Could not to get free port: %s", err)
	}

	resource, err = pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "postgres",
		Tag:          "latest",
		Env:          []string{"POSTGRES_PASSWORD=secret", "POSTGRES_DB=testdb"},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: strconv.Itoa(port)},
			},
		},
	})
	if err != nil {
		t.Fatalf("Could not start resource: %s", err)
	}

	return port
}

func TeardownDatabase(t *testing.T) {
	if err := pool.Purge(resource); err != nil {
		t.Fatalf("Could not purge resource: %s", err)
	}
}
