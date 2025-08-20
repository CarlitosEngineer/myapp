package main

import (
	"fmt"
	// "io/fs"
	"os"
	"path/filepath"
	"strings"
)

var ignoredDirs = map[string]bool{
	".git":   true,
	"target": true,
	"temp":   true,
}

func generateFolderStructure(dirPath, prefix string) (string, error) {
	var result strings.Builder

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return "", fmt.Errorf("error leyendo directorio: %w", err)
	}

	for i, entry := range entries {
		name := entry.Name()
		if ignoredDirs[name] {
			continue
		}

		isLast := i == len(entries)-1
		connector := "├── "
		if isLast {
			connector = "└── "
		}

		result.WriteString(fmt.Sprintf("%s%s%s\n", prefix, connector, name))

		if entry.IsDir() {
			subPrefix := prefix
			if isLast {
				subPrefix += "    "
			} else {
				subPrefix += "│   "
			}
			subStructure, err := generateFolderStructure(filepath.Join(dirPath, name), subPrefix)
			if err != nil {
				return "", err
			}
			result.WriteString(subStructure)
		}
	}

	return result.String(), nil
}

func writeStructureToFile(dirPath, outputFile string) error {
	structure, err := generateFolderStructure(dirPath, "")
	if err != nil {
		return fmt.Errorf("error generando estructura: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(outputFile), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creando carpeta de salida: %w", err)
	}

	return os.WriteFile(outputFile, []byte(structure), 0644)
}

func main() {
	fmt.Print("Ingresa la ruta de la carpeta original: ")
	var inputPath string
	fmt.Scanln(&inputPath)

	outputPath := filepath.Join("temp", "folder_structure.txt")
	err := writeStructureToFile(inputPath, outputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Estructura de carpetas guardada en: %s\n", outputPath)
}

// go run scripts/map_project_files.go
// /home/user_carlos01/documents/myapp
