package main

import (

	// importing the packages
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Config struct {
	Cluster_name   string
	One_click_apps string
	Num_nodes      string
	Doks_region    string
	Droplet_size   string
	Vpc_uuid       string
	Access_token   string
	Api_url        string
	Kbench_tests   string
}

func main() {

	fmt.Println("Checking if OS is Windows!!!")

	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	}

	fmt.Println("OS is not Windows!!!")

	fmt.Println("Reading the contents of the configuration file!")

	absPath, _ := filepath.Abs("doks_config.json")

	jsonFile, err := ioutil.ReadFile(absPath)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened the configuration file")

	var payload Config
	err = json.Unmarshal(jsonFile, &payload)

	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	} else {
		fmt.Println("Read the Configuration file successfully!")
	}

	// Fetching the absolute path of the shell scripts.

	absPath_1, _ := filepath.Abs("init_k8s.sh")
	absPath_2, _ := filepath.Abs("create_k8s.sh")
	absPath_3, _ := filepath.Abs("save_kubeconfig.sh")
	absPath_4, _ := filepath.Abs("create_k8s_vpc.sh")
	kbench_path, _ := filepath.Abs("run.sh")
	k_bench_install_path, _ := filepath.Abs("install.sh")

	// Making the shell scripts as executable

	_, err = exec.Command("chmod", "+x", absPath_1).Output()
	if err != nil {

		fmt.Println("Error during changing permissions for init_k8s.sh file: ", err)
	} else {
		fmt.Println("Change permissions for init_k8s.sh successful!")
	}

	_, err = exec.Command("chmod", "+x", absPath_2).Output()
	if err != nil {

		fmt.Println("Error during changing permissions for create_k8s.sh file: ", err)
	} else {
		fmt.Println("Change permissions for create_k8s.sh successful!")
	}

	_, err = exec.Command("chmod", "+x", absPath_3).Output()
	if err != nil {

		fmt.Println("Error during changing permissions for save_kubeconfig.sh file: ", err)
	} else {
		fmt.Println("Change permissions for save_kubeconfig.sh successful!")
	}

	_, err = exec.Command("chmod", "+x", absPath_4).Output()
	if err != nil {

		fmt.Println("Error during changing permissions for create_k8s_vpc.sh file: ", err)
	} else {
		fmt.Println("Change permissions for create_k8s_vpc.sh successful!")
	}

	_, err = exec.Command("chmod", "+x", kbench_path).Output()
	if err != nil {

		fmt.Println("Error during changing permissions for KBench Runner file: ", err)
	} else {
		fmt.Println("Change permissions for KBench Runner successful!")
	}

	_, err = exec.Command("chmod", "+x", k_bench_install_path).Output()
	if err != nil {

		fmt.Println("Error during changing permissions for KBench Installer file: ", err)
	} else {
		fmt.Println("Change permissions for KBench Installer successful!")
	}

	// K-Bench Install code will be written later!!!

	fmt.Println("Initializing the DOCTL Authorization!")

	_, err = exec.Command(absPath_1, payload.Access_token, payload.Api_url).Output()
	if err != nil {

		fmt.Println("Error during Initialization: ", err)
	} else {
		fmt.Println("DOCTL Initialization successful!")
	}

	fmt.Println("Creating the K8S cluster using DOCTL!")

	_, err = exec.Command(absPath_2, payload.Cluster_name, payload.One_click_apps, payload.Num_nodes, payload.Doks_region, payload.Droplet_size, payload.Access_token, payload.Api_url).Output()
	if err != nil {

		fmt.Println("Error during creating the k8s cluster: ", err)
	} else {
		fmt.Println("K8S Cluster Creation successful! - ", payload.Cluster_name)
	}

	fmt.Println("Saving the KUBECONFIG to the local!")

	_, err = exec.Command(absPath_3, payload.Cluster_name, payload.Access_token, payload.Api_url).Output()
	if err != nil {

		fmt.Println("Error during saving the Kubeconfig file: ", err)
	} else {
		fmt.Println("Kubeconfig save successful!")
	}

	// Running the KBench Tests
	fmt.Println("Running the KBench Tests on the cluster!")

	_, err = exec.Command(kbench_path, "-t", payload.Kbench_tests).Output()
	if err != nil {

		fmt.Println("Error while running K-Bench Tests: ", err)
	} else {
		fmt.Println("KBench Tests Ran Successfully on the cluster - ", payload.Cluster_name)
	}

}
