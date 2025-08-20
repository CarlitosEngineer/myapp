package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var ignoredDirs = map[string]bool{
	".git":    true,
	"target":  true,
	"temp":    true,
	".vscode": true,
}

func generateFolderStructure(dirPath, prefix string) (string, error) {
	var result strings.Builder

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return "", fmt.Errorf("error leyendo directorio: %w", err)
	}

	// Filtrar SOLO directorios y omitir los ignorados
	dirs := make([]os.DirEntry, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() && !ignoredDirs[e.Name()] {
			dirs = append(dirs, e)
		}
	}

	// Orden estable por nombre (opcional pero recomendado)
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })

	for i, entry := range dirs {
		name := entry.Name()
		isLast := i == len(dirs)-1

		connector := "├── "
		if isLast {
			connector = "└── "
		}
		result.WriteString(fmt.Sprintf("%s%s%s\n", prefix, connector, name))

		// Recursión solo en carpetas
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

	return result.String(), nil
}

func writeStructureToFile(dirPath, outputFile string) error {
	structure, err := generateFolderStructure(dirPath, "")
	if err != nil {
		return fmt.Errorf("error generando estructura: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(outputFile), os.ModePerm); err != nil {
		return fmt.Errorf("error creando carpeta de salida: %w", err)
	}

	return os.WriteFile(outputFile, []byte(structure), 0644)
}

func main() {
	fmt.Print("Ingresa la ruta de la carpeta original: ")
	var inputPath string
	fmt.Scanln(&inputPath)

	outputPath := filepath.Join("temp", "map_project_folders.txt")
	if err := writeStructureToFile(inputPath, outputPath); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Estructura de carpetas guardada en: %s\n", outputPath)
}

// Ejemplo:
// go run scripts/map_project_folders.go
// /home/user_carlos01/documents/myapp
