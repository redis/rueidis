package rueidis

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestScanner_Iter(t *testing.T) {
	tests := []struct {
		name     string
		entries  []ScanEntry
		err      error
		expected []string
		wantErr  bool
	}{
		{
			name: "single page",
			entries: []ScanEntry{
				{Elements: []string{"key1", "key2", "key3"}, Cursor: 0},
			},
			expected: []string{"key1", "key2", "key3"},
		},
		{
			name: "multiple pages",
			entries: []ScanEntry{
				{Elements: []string{"key1", "key2"}, Cursor: 10},
				{Elements: []string{"key3", "key4"}, Cursor: 0},
			},
			expected: []string{"key1", "key2", "key3", "key4"},
		},
		{
			name: "empty result",
			entries: []ScanEntry{
				{Elements: []string{}, Cursor: 0},
			},
			expected: []string{},
		},
		{
			name:    "error case",
			err:     errors.New("scan error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			scanner := NewScanner(func(cursor uint64) (ScanEntry, error) {
				if tt.err != nil {
					return ScanEntry{}, tt.err
				}
				if callCount >= len(tt.entries) {
					return ScanEntry{}, errors.New("unexpected call")
				}
				entry := tt.entries[callCount]
				callCount++
				return entry, nil
			})

			var result []string
			for element := range scanner.Iter() {
				result = append(result, element)
			}

			if tt.wantErr {
				if scanner.Err() == nil {
					t.Error("expected error but got none")
				}
			} else {
				if scanner.Err() != nil {
					t.Errorf("unexpected error: %v", scanner.Err())
				}
				if (len(result) != 0 || len(tt.expected) != 0) && !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("got %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestScanner_Iter2(t *testing.T) {
	tests := []struct {
		name         string
		entries      []ScanEntry
		err          error
		expectedKeys []string
		expectedVals []string
		wantErr      bool
	}{
		{
			name: "single page pairs",
			entries: []ScanEntry{
				{Elements: []string{"field1", "value1", "field2", "value2"}, Cursor: 0},
			},
			expectedKeys: []string{"field1", "field2"},
			expectedVals: []string{"value1", "value2"},
		},
		{
			name: "multiple pages pairs",
			entries: []ScanEntry{
				{Elements: []string{"field1", "value1"}, Cursor: 10},
				{Elements: []string{"field2", "value2", "field3", "value3"}, Cursor: 0},
			},
			expectedKeys: []string{"field1", "field2", "field3"},
			expectedVals: []string{"value1", "value2", "value3"},
		},
		{
			name: "odd number of elements",
			entries: []ScanEntry{
				{Elements: []string{"field1", "value1", "field2"}, Cursor: 0},
			},
			expectedKeys: []string{"field1"},
			expectedVals: []string{"value1"},
		},
		{
			name: "empty result",
			entries: []ScanEntry{
				{Elements: []string{}, Cursor: 0},
			},
			expectedKeys: []string{},
			expectedVals: []string{},
		},
		{
			name:    "error case",
			err:     errors.New("scan error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			scanner := NewScanner(func(cursor uint64) (ScanEntry, error) {
				if tt.err != nil {
					return ScanEntry{}, tt.err
				}
				if callCount >= len(tt.entries) {
					return ScanEntry{}, errors.New("unexpected call")
				}
				entry := tt.entries[callCount]
				callCount++
				fmt.Printf("callCount: %d, cursor: %d, elements: %v\n", callCount, entry.Cursor, entry.Elements)
				return entry, nil
			})

			var resultKeys, resultVals []string
			for key, val := range scanner.Iter2() {
				resultKeys = append(resultKeys, key)
				resultVals = append(resultVals, val)
			}

			if tt.wantErr {
				if scanner.Err() == nil {
					t.Error("expected error but got none")
				}
			} else {
				if scanner.Err() != nil {
					t.Errorf("unexpected error: %v", scanner.Err())
				}
				if (len(resultKeys) != 0 || len(tt.expectedKeys) != 0) && !reflect.DeepEqual(resultKeys, tt.expectedKeys) {
					t.Errorf("keys: got %v, want %v", resultKeys, tt.expectedKeys)
				}
				if (len(resultVals) != 0 || len(tt.expectedVals) != 0) && !reflect.DeepEqual(resultVals, tt.expectedVals) {
					t.Errorf("values: got %v, want %v", resultVals, tt.expectedVals)
				}
			}
		})
	}
}
