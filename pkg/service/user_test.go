package service

import (
	"context"
	"fmt"
	"ginserver/adaters/redis"
	"ginserver/pkg"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestUser_CreateUser(t *testing.T) {
	endpoint, err := prepareContainer(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := redis.NewRedigoRepository(endpoint)
	uc := NewUser(repo)

	t.Run("test add and read Jose", func(t *testing.T) {
		ctx := context.Background()
		user1 := pkg.UserCreateInput{
			Name: "Jose",
		}
		_, err := uc.CreateUser(ctx, user1)
		if err != nil {
			t.Fatal(err)
		}
		all, err := uc.GetAll(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if len(all) != 1 {
			t.Fatal("expected one user")
		}
		if all[0].Name != "Jose" {
			t.Fatal("expected user Jose")
		}
	})
}

func prepareContainer(ctx context.Context) (string, error) {

	// Configuración del contenedor de Redis
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	// Crear el contenedor
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return "", err
	}

	// Obtener la IP y el puerto del contenedor
	ip, err := redisC.Host(ctx)
	if err != nil {
		return "", err
	}
	port, err := redisC.MappedPort(ctx, nat.Port("6379"))
	if err != nil {
		return "", err
	}

	// Formar la URL de Redis
	redisURL := fmt.Sprintf("%s:%s", ip, port.Port())

	return redisURL, nil
}
