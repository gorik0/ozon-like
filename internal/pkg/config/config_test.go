package config

import (
	"os"
	"os/exec"
	"testing"
)

func TestMustLoad(t *testing.T) {
	tests := []struct {
		name            string
		mustExistDotEnv bool
	}{
		{"mustExistEnvFdile", true},
		{"mustNotExistEnvFile", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.mustExistDotEnv {
				println("WELL DONE!!!")
				TestNotCrasher(t)

			} else {
				TestCrasher(t)
				println("WELL DONE CRASSHHSHHSH!!!")

			}
		})
	}
}

func TestNotCrasher(t *testing.T) {
	//
	//cmd := exec.Command(os.Args[0], "-test.run=TestCrasher")
	//cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	//err := cmd.Run()
	//if e, ok := err.(*exec.ExitError); ok && !e.Success() {
	//	return
	//}
	//t.Fatalf("process ran with err %v, want exit status 1", err)

}

func TestCrasher(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		MustLoad()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCrasher")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
