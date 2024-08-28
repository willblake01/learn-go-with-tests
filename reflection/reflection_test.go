package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Will"},
			[]string{"Will"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Will", "Austin"},
			[]string{"Will", "Austin"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Will", 45},
			[]string{"Will"},
		},
		{
			"nested fields",
			Person{
				"Will",
				Profile{45, "Austin"},
			},
			[]string{"Will", "Austin"},
		},
		{
			"pointers to things",
			&Person{
				"Will",
				Profile{45, "Austin"},
			},
			[]string{"Will", "Austin"},
		},
		{
			"slices",
			[]Profile{
				{45, "Austin"},
				{34, "Reykjavik"},
			},
			[]string{"Austin", "Reykjavik"},
		},
		{
			"arrays",
			[2]Profile{
				{45, "Austin"},
				{34, "Reykjavik"},
			},
			[]string{"Austin", "Reykjavik"},
		},
		{
			"maps",
			map[string]string{
				"Foo":  "Bar",
				"Fizz": "Bang",
			},
			[]string{"Bar", "Bang"},
		},
	}

	assertContains := func(t testing.TB, haystack []string, needle string) {
		t.Helper()
		contains := false
		for _, x := range haystack {
			if x == needle {
				contains = true
			}
		}
		if !contains {
			t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
		}
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo":  "Bar",
			"Fizz": "Bang",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Bang")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Austin"}
			aChannel <- Profile{34, "Reykjavik"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Austin", "Reykjavik"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
