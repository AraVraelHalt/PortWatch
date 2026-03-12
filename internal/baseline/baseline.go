package baseline

import (
	"encoding/json"
	"fmt"
	"os"
  "path/filepath"
	"port-scanner/internal/models"
)

// Loads saved open ports from JSON
func LoadBaseline(filename string) map[int]bool {
	baselineMap := make(map[int]bool)

	data, err := os.ReadFile(filename)
	
  if err != nil {
		fmt.Println("No baseline file found, creating new one.")
		return baselineMap
	}

	var result models.ScanResult
	
  if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Error parsing baseline:", err)
		return baselineMap
	}

	for _, p := range result.OpenPorts {
		baselineMap[p] = true
	}
  
	return baselineMap
}

// Writes open ports to JSON
func SaveBaseline(filename string, ports []int) {
  dir := filepath.Dir(filename)
	data := models.ScanResult{OpenPorts: ports}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	
  if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
  
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		fmt.Println("Error writing baseline:", err)
	}
}
