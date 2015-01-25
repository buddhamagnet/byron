// usage: byron e805bd8086eebade820ac8368ec2fbf4
package main

import "encoding/json"
import "fmt"
import "io"
import "io/ioutil"
import "log"
import "net/http"
import "os"
import "os/exec"

var script = "agent-install.sh"

type Agent struct {
	Id              string `json:"_id"`
	State           string
	CredentialsName string
	CreationMethod  string
	PrivateIPs      []string
	ImageId         string
	CloudType       string
	SizePreset      string
	AccountId       string
	PrivateDNS      string
	Name            string
	Zone            string
	PublicDNS       string
	Region          string
	ProviderId      string
	KeyPair         string
	Provider        string
	Type            string
	AgentKey        string
	StateUpdatedAt  StateUpdatedAt
	Group           string
	PublicIPs       []string
	UpdatedAt       string
	CreatedAt       string
	Deleted         bool
}

type StateUpdatedAt struct {
	Sec  int
	Usec int
}

var agents []Agent
var agentKey string

func main() {

	if len(os.Args) != 2 {
		fmt.Println("usage: byron <token>")
		os.Exit(1)
	}

	token := os.Args[1]

	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	id := string(body)
	fmt.Println("ID:", id)

	resp, err = http.Get("https://api.serverdensity.io/inventory/devices?token=" + token)
	if err != nil {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &agents)
	if err != nil {
		log.Fatal(err)
	}

	for _, agent := range agents {
		if agent.ProviderId == id {
			agentKey = agent.AgentKey
			break
		}
	}

	out, err := os.Create(script)
	defer out.Close()

	if err != nil {
		log.Fatal(err)
	}

	resp, err = http.Get("https://www.serverdensity.com/downloads/" + script)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = os.Chmod(script, 0755)
	err = exec.Command("sed", "-i", "s/www.example.com/localhost/g", script).Run()

	if err != nil {
		log.Fatal(err)
	}

	output, err := exec.Command("./"+script, "-a", "https://economist.serverdensity.io", "-g", "Drupal-Web", "-k", agentKey).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
