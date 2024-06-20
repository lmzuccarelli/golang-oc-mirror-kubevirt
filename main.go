package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/microlib/simple"
	"gopkg.in/yaml.v2"
)

type InstallerBootableImages struct {
	Stream        string        `json:"stream"`
	Metadata      Metadata      `json:"metadata"`
	Architectures Architectures `json:"architectures"`
}

type Metadata struct {
	LastModified time.Time `json:"last-modified"`
	Generator    string    `json:"generator"`
}

type Artifacts struct {
	Kubevirt Kubevirt `json:"kubevirt"`
}

type Kubevirt struct {
	Release   string `json:"release"`
	Image     string `json:"image"`
	DigestRef string `json:"digest-ref"`
}

type Images struct {
	Kubevirt Kubevirt `json:"kubevirt"`
}

type X8664 struct {
	Artifacts Artifacts `json:"artifacts"`
	Images    Images    `json:"images"`
}

type Architectures struct {
	X8664 X8664 `json:"x86_64"`
}

type InstallerConfigMap struct {
	Kind     string `yaml:"kind"`
	Metadata struct {
		Annotations struct {
			IncludeReleaseOpenshiftIoIbmCloudManaged             string `yaml:"include.release.openshift.io/ibm-cloud-managed"`
			IncludeReleaseOpenshiftIoSelfManagedHighAvailability string `yaml:"include.release.openshift.io/self-managed-high-availability"`
			IncludeReleaseOpenshiftIoSingleNodeDeveloper         string `yaml:"include.release.openshift.io/single-node-developer"`
		} `yaml:"annotations"`
		CreationTimestamp interface{} `yaml:"creationTimestamp"`
		Name              string      `yaml:"name"`
		Namespace         string      `yaml:"namespace"`
	} `yaml:"metadata"`
	APIVersion string `yaml:"apiVersion"`
	Data       struct {
		ReleaseVersion string `yaml:"releaseVersion"`
		Stream         string `yaml:"stream"`
	} `yaml:"data"`
}

func main() {

	logger := &simple.Logger{Level: "info"}
	var ibi InstallerBootableImages
	var icm InstallerConfigMap

	// parse the main yaml file
	file, _ := os.ReadFile("0000_50_installer_coreos-bootimages.yaml")
	errs := yaml.Unmarshal(file, &icm)
	if errs != nil {
		logger.Error(fmt.Sprintf("reading yaml file %v", errs))
		os.Exit(-1)
	}

	logger.Trace(fmt.Sprintf("data %v", icm.Data.Stream))
	// now parse the json section
	errs = json.Unmarshal([]byte(icm.Data.Stream), &ibi)
	if errs != nil {
		logger.Error(fmt.Sprintf("parsing json from configmap data %v", errs))
		os.Exit(-1)
	}
	logger.Info(fmt.Sprintf("kubevirt %v", ibi.Architectures.X8664.Images.Kubevirt.DigestRef))

}
