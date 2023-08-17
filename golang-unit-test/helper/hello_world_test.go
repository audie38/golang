package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*	Go Unit Test Command Line:
 *	go test = run all test
 *	go test -v = to show what test functions running
 *	go test -v ./... = to run all test from root folder
 *	go test -v -run=[test functions name] = to run specific test functions
 *  go test -v -run=[test function name]/[subtest name] = to run specific test functions subtest
 */

/* Unit Test Failed Feedback Syntax:
 * t.Fail() = failed the unit test and continue
 * t.FailNow() = failed the unit test and not continue
 * t.Error() = similar to t.Fail() but can pass failed feedback arguments
 * t.Fatal() = similar to t.FailNow() but can pass failed feedback arguments
 */

/* Unit Test Syntax:
 * t.Skip() = to skip unit test
 * m *testing.M = can be use to check before and after unit test run
 * t.Run() = can be used to create sub test
 */

/* Assertion Syntax:
 * assert.equal = will run t.Fail() if the actual != expected result
 * assert.equal(t, expected result, actual result, feedback if actual != expected)
 * require.equal = will run t.FailNow() if the actual != expected result
 * require.equal(t, expected result, actual result, feedback if actual != expected)
 */

 func TestMain(m *testing.M){
	// Before
	fmt.Println("Before Unit Test")
	m.Run()
	// After
	fmt.Println("After Unit Test")
 }

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Audie")
	if result != "Hello Audie"{
		t.Error("Result must be Hello Audie")
	}
	fmt.Println("Finish Unit Test t.Error")
}
func TestHelloWorldDummy(t *testing.T){
	result := HelloWorld("Dummy")
	if result != "Hello Dummy"{
		t.Fatal("Result must be Hello Dummy")
	}
	fmt.Println("Finish Unit Test t.Fatal")
}

func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Dummy")
	require.Equal(t, "Hello Dummy", result, "Result must be Hello Dummy")
}

func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("Ichigo")
	assert.Equal(t, "Hello Ichigo", result, "Result must be Hello Ichigo")
}

func TestSkip(t *testing.T){
	if runtime.GOOS == "windows"{
		t.Skip("Cannot run on Windows")
	}

	result := HelloWorld("Test")
	assert.Equal(t, "Hello Test", result, "Result must be Hello Test")
}

func TestSubTest(t *testing.T){
	t.Run("Ichigo", func(t *testing.T){
		result := HelloWorld("Dummy")
		require.Equal(t, "Hello Ichigo", result, "Result must be Hello Ichigo")
	})
	t.Run("Kurosaki", func(t *testing.T){
		result := HelloWorld("Kurosaki")
		assert.Equal(t, "Hello Kurosaki", result, "Result must be Hello Kurosaki")
	})
}