package task

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestAddSingle(t *testing.T) {
	emptyContentError := errors.New("content has empty value")
	tests := []struct {
		testName       string
		inputContent   string
		inputJira      string
		expectedOutput TaskSlice
		expectedErr    error
	}{
		{"adding valid values", "fixed a bug", "CONFIG-9501", TaskSlice{{Content: "fixed a bug", JiraTicket: "CONFIG-9501", dateCreated: time.Now()}}, nil},
		{"adding empty content", "", "CONFIG-9501", nil, emptyContentError},
		{"adding empty jira", "fixed a bug", "", TaskSlice{{Content: "fixed a bug", JiraTicket: "", dateCreated: time.Now()}}, nil},
		{"adding empty everything", "", "", nil, emptyContentError},
	}
	fmt.Println("TestAdd")
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			actual := TaskSlice{}
			err := actual.Add(tt.inputContent, tt.inputJira)
			if tt.expectedErr != nil {
				if err == nil {
					t.Errorf("expected error %v, but no error returned", err)
				} else if err.Error() != tt.expectedErr.Error() {
					t.Errorf("got error .%v., wanted error .%v.\n", err, tt.expectedErr)
				}
			} else if tt.expectedErr == nil {
				if err != nil {
					t.Errorf("unexpected error .%v. returned\n", err)
				}
				if !isEqual(actual, tt.expectedOutput) {
					t.Errorf("got .%v., wanted .%v.\n", actual, tt.expectedOutput)
				}
			}
		})

	}
}
func TestAddMultiple(t *testing.T) {
	emptyContentError := errors.New("content has empty value")
	tests := []struct {
		testName       string
		inputContent   string
		inputJira      string
		inputSlice     TaskSlice
		expectedOutput TaskSlice
		expectedErr    error
	}{
		{"adding value to non-empty slice", "fixed another issue", "CONFIG-9502", TaskSlice{{Content: "fixed a bug", JiraTicket: "CONFIG-9501", dateCreated: time.Now()}}, TaskSlice{{Content: "fixed a bug", JiraTicket: "CONFIG-9501", dateCreated: time.Now()}, {Content: "fixed another issue", JiraTicket: "CONFIG-9502", dateCreated: time.Now()}}, nil},
		{"adding missing content to non-empty slice", "", "CONFIG-9502", TaskSlice{{Content: "fixed a bug", JiraTicket: "CONFIG-9501", dateCreated: time.Now()}}, nil, emptyContentError},
	}
	fmt.Println("TestAddMultiple")
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.inputSlice.Add(tt.inputContent, tt.inputJira)
			if tt.expectedErr != nil {
				if err == nil {
					t.Errorf("expected error %v, but no error returned", err)
				} else if err.Error() != tt.expectedErr.Error() {
					t.Errorf("got error .%v., wanted error .%v.\n", err, tt.expectedErr)
				}
			} else if tt.expectedErr == nil {
				if err != nil {
					t.Errorf("unexpected error .%v. returned\n", err)
				}
				if !isEqual(tt.inputSlice, tt.expectedOutput) {
					t.Errorf("got .%v., wanted .%v.\n", tt.inputSlice, tt.expectedOutput)
				}
			}
		})

	}

}

// func TestGetFromFile(t *testing.T) {}

// Created a comparison bc reflection would always be false due to time.now
func isEqual(a TaskSlice, b TaskSlice) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i].Content != b[i].Content {
			return false
		}
		if a[i].JiraTicket != b[i].JiraTicket {
			return false
		}
	}
	return true

}
