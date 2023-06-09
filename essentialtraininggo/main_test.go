package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strconv"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func almostEqual(v1, v2 float64) bool {
// 	return Abs(v1-v2) <= 0.001
// }

// func TestSimple(t *testing.T) {
// 	val, err := Sqrt(4.0)

// 	if err != nil {
// 		t.Fatalf("error in calculation - %f", err)
// 	}

// 	if !almostEqual(val, 2) {
// 		t.Fatalf("Bad value - %f", val)
// 	}
// }

// // Table-Driven tests

// type TestCase struct {
// 	value    float64
// 	expected float64
// }

// func TestMany(tt *testing.T) {
// 	// prepare test cases!
// 	testCases := []TestCase{
// 		{0.0, 0.0},
// 		{4.0, 2.0},
// 		{9.0, 3.0},
// 	}

// 	for _, tc := range testCases {
// 		// running sub-test
// 		tt.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
// 			out, err := Sqrt(tc.value)
// 			if err != nil {
// 				t.Fatalf("error")
// 			}

// 			if !almostEqual(out, tc.expected) {
// 				t.Fatalf("%f != %f", out, tc.expected)
// 			}
// 		})
// 	}
// }

// func TestIt(t *testing.T) {
// 	val, err := Sqrt(4)
// 	require.NoError(t, err)
// 	require.InDelta(t, 2, val, 0.001)
// }

// func loadSqrtCasesFromCsv(t *testing.T) []TestCase {
// 	//open the Csv file
// 	f, err := os.Open("sqrt_cases.csv")
// 	defer f.Close()
// 	require.NoError(t, err)
// 	var cases []TestCase
// 	rdr := csv.NewReader(f)
// 	// start reading line by line
// 	for {
// 		record, err := rdr.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		require.NoError(t, err)
// 		// convert the string record to float record
// 		val, err := strconv.ParseFloat(record[0], 64)
// 		require.NoError(t, err)
// 		expected, err := strconv.ParseFloat(record[1], 64)
// 		require.NoError(t, err)
// 		tc := TestCase{val, expected}
// 		cases = append(cases, tc)
// 	}

// 	return cases
// }

// func TestCsv(t *testing.T) {
// 	for _, tc := range loadSqrtCasesFromCsv(t) {
// 		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
// 			out, err := Sqrt(tc.value)
// 			require.NoError(t, err)
// 			require.InDelta(t, tc.expected, out, 0.001)
// 		})
// 	}
// }
