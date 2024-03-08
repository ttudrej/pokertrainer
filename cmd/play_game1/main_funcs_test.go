/*
File: main_funcs_test.go, version history:
v0.1	2018/08/15	Tomasz Tudrej

*/

// package main

import "testing"

// TestFirstFuncToTestWith, doc line for the first test case
func TestFirstFuncToTestWith(t *testing.T) {

	cases := []struct {
		inputProvided, outputExpected string
	}{
		{"a", "aa"},
		{"blah", "blahblah"},
	}

	for _, c := range cases {

		twicethestring := firstFuncToTestWith(c.inputProvided)

		// fmt.Println("s is      ", s)
		// fmt.Println("twice is: ", twicethestring)
		// fmt.Println("s+s is:   ", s+s)

		if twicethestring != c.outputExpected {
			t.Errorf("expected: %v, got %v", c.outputExpected, twicethestring)
		}
	}
}

/*
// TestReadingConfigFiles doc line
func TestReadingConfigFiles(t *testing.T) {

	cases := []struct {
		fileName, filePath string
	}{
		{"config", "./"},
	}

	for _, c := range cases {

		viper.SetConfigName(c.fileName) // no need to include file extension
		viper.AddConfigPath(c.filePath) // set the path of your config file

		err := viper.ReadInConfig()

		if err != nil {
			fmt.Println("Config file not found...")
		} else {

			if viper.IsSet("testCredentials.testLogin") == false {
				t.Error("login not set in ", c.filePath, "/", c.fileName)
			}

			if viper.IsSet("testCredentials.testPassword") == false {
				t.Error("password not set in ", c.filePath, "/", c.fileName)
			}
		}
	}
}

*/

// EOF
