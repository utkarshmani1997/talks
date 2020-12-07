package bootstrap

import (
	"fmt"
	"os/exec"
	"time"
)

func Start(name string, port string) {
	fmt.Printf("Starting PostgreSQL container named %s at port %s...\n", name, port)
	_, err := exec.LookPath("docker")
	if err != nil {
		panic("Docker not installed")
	}
	cmd := exec.Command("docker", "run", "--name", name,
		"-d", "-p", port+":5432", "-e", "POSTGRES_DB=test", "-e", "POSTGRES_USER=test",
		"-e", "POSTGRES_PASSWORD=test", "docker.bintray.io/postgres:9.5.2")
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output[:]))
	time.Sleep(time.Second * 10)
	if err != nil {
		panic(fmt.Sprintf("Failed to run docker command: %s.\n", err))
	}
}

func Stop(name string) {
	fmt.Println("Deleting PostgreSQL container...")
	cmd := exec.Command("docker", "rm", "-f", name)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Failed to run docker command: %s", err))
	}
}
