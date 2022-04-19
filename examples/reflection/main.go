package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

type AppManifest struct {
	Name     string        `yaml:"name"`
	Services []interface{} `yaml:"services"`
}

type Manifest struct {
	Applications []AppManifest `yaml:"applications"`
}

func main() {
	manifest, _ := LoadManifest("./deployment-manifest.yml")
	fmt.Printf("manifest: %#v\n\n", manifest)
	for _, app := range manifest.Applications {
		processAppManifest(app)
	}
}

func processAppManifest(manifest AppManifest) {
	fmt.Printf("appManifest: %#v\n", manifest)
	for _, svc := range manifest.Services {
		processService(svc, manifest)
	}
}

func processService(service interface{}, manifest AppManifest) {

	fmt.Println("---------------------------------------------")
	fmt.Printf("svc: %#v\n", service)
	svcStructType := reflect.TypeOf(service)
	fmt.Printf("svc type: %v\n", svcStructType)
	svcStructKind := reflect.ValueOf(service).Kind()
	fmt.Printf("svc kind: %v\n", svcStructKind)

	if svcStructKind == reflect.String {
		fmt.Println("--> string")
		svcType := service.(string)
		svcName := manifest.Name
		fmt.Printf("service type: %s name: %s\n", svcType, svcName)

	}
	if svcStructKind == reflect.Map {
		fmt.Println("--> map")
		svcMap := service.(map[interface{}]interface{})
		fmt.Printf("svcMap: %v\n", svcMap)
		svcType, present := svcMap["type"]
		if !present {
			panic("type!!!")
		}
		svcName, present := svcMap["name"]
		if !present {
			panic("name!!!")
		}
		fmt.Printf("service type: %s name: %s\n", svcType, svcName)
	}

}

func LoadManifest(manifestPath string) (*Manifest, error) {
	manifestFileContent, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return nil, err
	}
	manifest := Manifest{}
	err = yaml.Unmarshal(manifestFileContent, &manifest)
	if err != nil {
		return nil, err
	}
	return &manifest, nil
}
