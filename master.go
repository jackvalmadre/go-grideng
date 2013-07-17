package grideng

import (
	"fmt"
	"os"
	"os/exec"
)

// Executes all tasks and returns a list of output files.
func Master(inputs []Input, resources string, cmdArgs []string) ([]string, error) {
	// Serialize inputs.
	if err := saveAllInputs(inputs); err != nil {
		return nil, err
	}

	var args []string
	// Submitting a binary job.
	args = append(args, "-b", "y")
	// Wait for jobs to finish.
	args = append(args, "-sync", "y")
	// Use current working directory.
	args = append(args, "-cwd")
	// Use same environment variables.
	args = append(args, "-V")
	// Use same environment variables.
	args = append(args, "-t", fmt.Sprintf("1-%d", len(inputs)))
	// Set resources.
	if len(resources) > 0 {
		args = append(args, "-l", resources)
	}
	// Redirect stdout.
	args = append(args, "-o", `stdout-$TASK_ID`)
	// Redirect stderr.
	args = append(args, "-e", `stderr-$TASK_ID`)

	// Name of executable to run.
	args = append(args, os.Args[0])
	args = append(args, cmdArgs...)

	// Submit.
	cmd := exec.Command("qsub", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Print("qsub")
	for _, arg := range args {
		fmt.Print(" ", arg)
	}
	fmt.Println()

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	// Check success/failure of tasks.
	files := make([]string, len(inputs))
	for i, input := range inputs {
		name := input.Name()
		outfile := outputFile(name)
		// Check if output file exists.
		if _, err := os.Stat(outfile); err == nil {
			files[i] = outfile
		}
	}
	return files, nil
}

// Attempts to save each task input to a file.
func saveAllInputs(inputs []Input) error {
	for i, input := range inputs {
		// Grid Engine task ID (one-indexed not zero-indexed).
		num := i + 1
		// Save input to file.
		file := inputFile(num)
		if err := saveInput(input, file); err != nil {
			return err
		}
	}
	return nil
}

// Attempts to save the input for a task to a file.
// Open files are closed on return.
func saveInput(input Input, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := input.Write(file); err != nil {
		return err
	}
	return nil
}
