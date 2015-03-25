package main

import (
	"testing"
)

type fixture struct {
	URL      string
	expected []string
}

func TestURLOptions(t *testing.T) {
	fixtures := []fixture{
		fixture{"/service-manual/digital-by-default.html",
			[]string{"./service-manual/digital-by-default.md"},
		},
		fixture{"/service-manual/assets/images/digital-by-default.png",
			[]string{"./service-manual/assets/images/digital-by-default.png"},
		},
		fixture{"/service-manual/digital-by-default/",
			[]string{"./service-manual/digital-by-default/index.md",
				"./service-manual/digital-by-default/index.html"},
		},
		fixture{"/service-manual/digital-by-default",
			[]string{"./service-manual/digital-by-default/index.md",
				"./service-manual/digital-by-default/index.html",
				"./service-manual/digital-by-default.md"},
		},
	}
	for _, f := range fixtures {
		expectURLs(t, f, getURLOptions(f.URL))
	}
}

func expectURLs(t *testing.T, f fixture, options []string) {
	if len(f.expected) != len(options) {
		t.Errorf("Expected: %v but got %v", len(f.expected), options)
	}
	for i, _ := range f.expected {
		if f.expected[i] != options[i] {
			t.Errorf("Expected <%v> but got <%v>", f.expected[i], options[i])
		}
	}
}
