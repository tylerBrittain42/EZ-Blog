package basicArticle

import (
	"testing"
)

// getTitle() string
// getContent() string
func TestGetTitle(t *testing.T) {

	tests := []struct {
		description    string
		input          string
		expectedOutput string
		shouldError    bool
		errorMessage   string
	}{
		{
			description:    "strip md extension",
			input:          "testArticle.md",
			expectedOutput: "testArticle",
			shouldError:    false,
			errorMessage:   "",
		}, {
			description:    "strip txt extension",
			input:          "testArticle.txt",
			expectedOutput: "testArticle",
			shouldError:    false,
			errorMessage:   "",
		}, {
			description:    "too many dots",
			input:          "test.Article.txt",
			expectedOutput: "",
			shouldError:    true,
			errorMessage:   "filename must contain a single dot",
		}, {
			description:    "no dots",
			input:          "testArticletxt",
			expectedOutput: "",
			shouldError:    true,
			errorMessage:   "filename must contain a single dot",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actualOutput, err := GetTitle(tt.input)
			if tt.shouldError {
				if err == nil {
					t.Errorf("GetTitle() did not recieve an error when expected")
				}
				if err.Error() != tt.errorMessage {
					t.Errorf("GetTitle() recieved error '%v', but wanted error '%v'", err, tt.errorMessage)
				}
				if tt.expectedOutput != actualOutput {
					t.Errorf("GetTitle() recieved response %v when an error was expected", actualOutput)
				}
			} else {
				if err != nil {
					t.Errorf("GetTitle() recieved an unexpected error: %v", err)
				}
				if tt.expectedOutput != actualOutput {
					t.Errorf("GetTitle() expects %v, got %v", tt.expectedOutput, actualOutput)
				}
			}
		})
	}
}
