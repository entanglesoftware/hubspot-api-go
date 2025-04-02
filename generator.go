package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Resource struct {
	Name    string `yaml:"name"`    // Name of the resource (e.g., contacts, companies)
	Enabled bool   `yaml:"enabled"` // Whether to generate code for this resource
}

type GeneratorConfig struct {
	Generator string     `yaml:"generator"` // Generator version (e.g., 1.0.0)
	Package   string     `yaml:"package"`   // Default package name for all resources
	Resources []Resource `yaml:"resources"` // List of resources to generate
}

// Dynamic config template
const configTemplate = `
package: {{.PackageName}}
output: {{.OutputPath}}
generate:
  {{- if .ModelsOnly }}
  models: true
  chi-server: false
  client: false
  {{- else if .ClientOnly }}
  client: true
  models: false
  chi-server: false
  {{- else }}
  models: false
  client: false
  chi-server: false
  {{- end }}
output-options:
  skip-prune: true
import-mapping:
  components.yaml: "-"
`

func main() {
	baseDir := "./openapi" // Base directory to scan for generator.yaml files
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Base(path) == "generator.yaml" {
			processGeneratorFile(path)
		}
		if filepath.Base(path) == "commerce_code_generator.yaml" {
			processGeneratorFile(path)
		}
		if filepath.Base(path) == "association_code_generator.yaml" {
			processGeneratorFile(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error scanning directories: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Code generation complete.")
}

func processGeneratorFile(generatorFile string) {
	fmt.Printf("Processing generator file: %s\n", generatorFile)

	var config GeneratorConfig
	file, err := os.Open(generatorFile)
	if err != nil {
		fmt.Printf("Error opening generator file %s: %v\n", generatorFile, err)
		return
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		fmt.Printf("Error decoding generator file %s: %v\n", generatorFile, err)
		return
	}

	if config.Generator != "1.0.0" {
		fmt.Printf("Skipping file %s (unsupported generator version: %s).\n", generatorFile, config.Generator)
		return
	}

	for _, resource := range config.Resources {
		if !resource.Enabled {
			fmt.Printf("Skipping generation for %s (disabled).\n", resource.Name)
			continue
		}

		resourcePath := filepath.Join(filepath.Dir(generatorFile), resource.Name)
		cleanedPath := strings.Replace(filepath.Dir(generatorFile), "openapi", "", 1)
		outputDir := filepath.Join("codegen", cleanedPath, resource.Name)

		// Generate `models` file first using `components.yaml` relative to the resource path
		componentsFile := filepath.Join(resourcePath, "components.yaml")
		if _, err := os.Stat(componentsFile); err == nil {
			modelsOutputPath := filepath.Join(outputDir, "models.gen.go")
			generateFile(resource.Name, componentsFile, modelsOutputPath, true, false, false)
		} else {
			fmt.Printf("No components.yaml found for %s, skipping models generation.\n", resource.Name)
		}

		// Process resource file for client and orchestration
		if isDirectory(resourcePath) {
			mainYaml := findOpenAPISpec(resourcePath)
			if mainYaml == "" {
				fmt.Printf("Skipping %s (no valid OpenAPI 3.0.0 file found).\n", resource.Name)
				continue
			}
			generateFiles(resource.Name, mainYaml, outputDir)
		} else {
			generateFiles(resource.Name, resourcePath, outputDir)
		}
	}
}

func findOpenAPISpec(directory string) string {
	var mainYaml string
	_ = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".yaml" {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer file.Close()

		var doc map[string]interface{}
		if err := yaml.NewDecoder(file).Decode(&doc); err != nil {
			return nil
		}

		if version, ok := doc["openapi"]; ok && version == "3.0.0" {
			mainYaml = path
			return filepath.SkipDir
		}
		return nil
	})
	return mainYaml
}

func generateFiles(resourceName, resourcePath, outputDir string) {
	fmt.Printf("Generating files for resource: %s\n", resourceName)

	// Generate client file
	clientOutputPath := filepath.Join(outputDir, "client.gen.go")
	generateFile(resourceName, resourcePath, clientOutputPath, false, true, true)

	// Generate orchestration/resource-specific file
	resourceOutputPath := filepath.Join(outputDir, fmt.Sprintf("%s.gen.go", resourceName))
	generateFile(resourceName, resourcePath, resourceOutputPath, false, false, false)
}

func generateFile(resourceName, resourcePath, outputPath string, modelsOnly, clientOnly, chiServer bool) {
	fmt.Printf("Generating file: %s\n", outputPath)

	// Ensure the output directory exists
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating directories for %s: %v\n", outputPath, err)
		return
	}

	// Generate dynamic config content
	tmpl, err := template.New("config").Parse(configTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	data := struct {
		PackageName string
		OutputPath  string
		ModelsOnly  bool
		ClientOnly  bool
		ChiServer   bool
	}{
		PackageName: resourceName, // Use resource name as package
		OutputPath:  outputPath,
		ModelsOnly:  modelsOnly,
		ClientOnly:  clientOnly,
		ChiServer:   chiServer,
	}

	var configBuffer bytes.Buffer
	if err := tmpl.Execute(&configBuffer, data); err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	// Write the dynamic config file
	dynamicConfigPath := filepath.Join(outputDir, "config.dynamic.yaml")
	if err := os.WriteFile(dynamicConfigPath, configBuffer.Bytes(), 0644); err != nil {
		fmt.Printf("Error writing config file: %v\n", err)
		return
	}

	// Run `oapi-codegen` with the dynamic configuration
	cmd := exec.Command(
		"go", "run", "-modfile=./tools/go.mod",
		"github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen",
		"--config="+dynamicConfigPath,
		resourcePath,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error generating file %s: %v\n", outputPath, err)
	}

	// Remove the dynamic config file
	if err := os.Remove(dynamicConfigPath); err != nil {
		fmt.Printf("Error removing config file %s: %v\n", dynamicConfigPath, err)
	}

	fmt.Printf("File generated at: %s\n", outputPath)
}

func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
