package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// Function to extract time in minutes from commit messages
func extractTime(commitMessages []string) (int, error) {
	totalMinutes := 0
	// Define regex for extracting hours and minutes
	timePattern := `⏰\s*(\d+h)?\s*(\d+m)?`

	// Iterate through each commit message
	for _, line := range commitMessages {
		re := regexp.MustCompile(timePattern)
		matches := re.FindStringSubmatch(line)

		if len(matches) > 0 {
			// Extract hours
			var hours, minutes int
			if matches[1] != "" {
				hours, _ = strconv.Atoi(strings.TrimSuffix(matches[1], "h"))
			}
			if matches[2] != "" {
				minutes, _ = strconv.Atoi(strings.TrimSuffix(matches[2], "m"))
			}

			// Convert hours to minutes and add to the total
			totalMinutes += (hours * 60) + minutes
		}
	}

	return totalMinutes, nil
}

func ensureNewline(s string) string {
	// Check if the last character is a newline
	if len(s) > 0 && s[len(s)-1] != '\n' {
		// Add newline if not present
		return s + "\n"
	}
	// Return string as is if it already ends with a newline
	return s
}

// Function to update or append the total time to README.md
func updateReadme(readmeFile string, totalTime string) error {
	content, err := os.ReadFile(readmeFile)
	if err != nil {
		return fmt.Errorf("unable to read README.md: %v", err)
	}

	readmeContent := string(content)
	sectionHeader := "## Total time spent"
	timeSpentLine := "- Total time spent: " + totalTime

	if strings.Contains(readmeContent, sectionHeader) {
		// If the section exists, replace the total time line
		re := regexp.MustCompile(`(?m)^- Total time spent:.*$`)

		// If the line exists (even if empty), replace it
		if re.MatchString(readmeContent) {
			readmeContent = re.ReplaceAllString(readmeContent, timeSpentLine)
		} else {
			// If the line is missing, add the time spent line
			readmeContent = strings.Replace(readmeContent, sectionHeader+"\n", sectionHeader+"\n\n"+timeSpentLine+"\n", 1)
		}
	} else {
		// If the section does not exist, append it
		readmeContent += "\n\n" + sectionHeader + "\n\n" + timeSpentLine
	}

	// Write the updated content back to README.md
	readmeContent = ensureNewline(readmeContent)
	err = os.WriteFile(readmeFile, []byte(readmeContent), 0644)
	if err != nil {
		return fmt.Errorf("unable to write to README.md: %v", err)
	}

	return nil
}

func main() {
	// Command to get the commit messages
	cmd := exec.Command("git", "log", "--pretty=format:%s")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error getting commit messages: %v\n", err)
		return
	}

	// Split commit messages into lines
	commitMessages := strings.Split(string(output), "\n")

	// Extract total time from commit messages
	totalMinutes, err := extractTime(commitMessages)
	if err != nil {
		fmt.Printf("Error extracting time: %v\n", err)
		return
	}

	// Convert total minutes to hours and minutes
	totalHours := totalMinutes / 60
	remainingMinutes := totalMinutes % 60
	totalTime := fmt.Sprintf("%d hours %d minutes", totalHours, remainingMinutes)

	// Path to the README file
	readmeFile := "README.md"

	// Update the README with the total time
	err = updateReadme(readmeFile, totalTime)
	if err != nil {
		fmt.Printf("Error updating README.md: %v\n", err)
		return
	}

	// Print success message
	fmt.Println("Total time updated successfully in README.md")
}
