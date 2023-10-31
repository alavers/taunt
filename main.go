package main

import (
	"fmt"
	"net/http"
  "os"
	"os/exec"
  "path/filepath"
  "runtime"
  "strings"
  "strconv"
)

func shell() {
	// Get the directory of the currently running executable
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Dir(b)
	scriptPath := filepath.Join(dir, "taunt.sh")

	cmd := exec.Command(scriptPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
  fmt.Println(strings.Trim(string(output), "\n"))
}

func getPort() string {
  port := os.Getenv("PORT")
  if port == "" {
    return "8080"
  }

  _, err := strconv.Atoi(port)
  if err != nil {
		fmt.Printf("Invalid PORT value %s, using default 8080\n", port)
		return "8080"
	}

	return port
}

func main() {
	http.HandleFunc("/taunt", tauntHandler)
  port := getPort()
	fmt.Println("Server is running on http://localhost:"+port)
	http.ListenAndServe(":"+port, nil)
}

func tauntHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}

  shell()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 OK"))
}
