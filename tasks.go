package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

var agents []Agent

func removeAgent(client http.Client, token string) {}

func installAgent(client http.Client, token string, user string) {
	script := "agent-install.sh"
	agentKey := ""

	// Retrieve instance ID from Amazon.
	resp := getInstanceId(client)
	defer resp.Body.Close()

	// Read the response and store the instance Id.
	body, err := ioutil.ReadAll(resp.Body)
	id := string(body)
	fmt.Println("ID:", id)

	// Contact the inventory with the token.
	resp = getDevices(token, client)
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	abortOnFail(err)

	// Store the JSON response into the agent struct slice.
	err = json.Unmarshal(body, &agents)
	abortOnFail(err)

	// Iterate over the slice to find the correct agent.
	for _, agent := range agents {
		if agent.ProviderId == id {
			agentKey = agent.AgentKey
			break
		}
	}

	// Open the installer file.
	out, err := os.Create(script)
	defer out.Close()
	abortOnFail(err)

	resp = getInstaller(script, client)
	defer resp.Body.Close()
	abortOnFail(err)

	// Copy the downloaded installer file.
	_, err = io.Copy(out, resp.Body)
	abortOnFail(err)

	// Make the installer executable and swap out
	// the example domains for localhost.
	err = os.Chmod(script, 0755)
	err = exec.Command("sed", "-i", "s/www.example.com/localhost/g", script).Run()
	abortOnFail(err)

	// Now execute the installer.
	output, err := exec.Command("./"+script, "-a", "https://economist.serverdensity.io", "-g", user, "-k", agentKey).Output()
	abortOnFail(err)

	fmt.Println(string(output))
}
