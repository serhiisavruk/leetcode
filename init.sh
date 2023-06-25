#!/bin/bash

# Get folder name from user input
echo "Enter the project name: "
read project_name

# Check if directory exists
if [ -d "$project_name" ]; then
  echo "Directory $project_name already exists."
  exit 1
fi

# Create new directory
mkdir $project_name

# Change directory to the new project directory
cd $project_name

# Create go.mod file
go mod init $project_name

# Create main.go file
cat << EOF > main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(DoSomething())
}
EOF

# Create task.go file
cat << EOF > task.go
package main

func DoSomething() string {
	return "Task done!"
}
EOF

# Create task_test.go file
cat << EOF > task_test.go
package main

import "testing"

func TestDoSomething(t *testing.T) {
	expected := "Task done!"
	result := DoSomething()
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}
EOF

# End of script
echo "Go project $project_name has been successfully created."
