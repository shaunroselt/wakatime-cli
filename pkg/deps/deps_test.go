package deps_test

import (
	"regexp"
	"testing"

	"github.com/wakatime/wakatime-cli/pkg/deps"
	"github.com/wakatime/wakatime-cli/pkg/heartbeat"
	"github.com/wakatime/wakatime-cli/pkg/regex"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithDetection(t *testing.T) {
	opt := deps.WithDetection(deps.Config{})

	h := opt(func(hh []heartbeat.Heartbeat) ([]heartbeat.Result, error) {
		assert.Equal(t, []heartbeat.Heartbeat{
			{
				Dependencies: []string{
					"os",
					"github.com/wakatime/wakatime-cli/pkg/heartbeat",
				},
				Entity:     "testdata/golang_minimal.go",
				EntityType: heartbeat.FileType,
				Language:   heartbeat.LanguagePtr(heartbeat.LanguageGo),
			},
		}, hh)

		return []heartbeat.Result{
			{
				Status: 201,
			},
		}, nil
	})

	result, err := h([]heartbeat.Heartbeat{{
		Entity:     "testdata/golang_minimal.go",
		EntityType: heartbeat.FileType,
		Language:   heartbeat.LanguagePtr(heartbeat.LanguageGo),
	}})
	require.NoError(t, err)

	assert.Equal(t, []heartbeat.Result{
		{
			Status: 201,
		},
	}, result)
}

func TestWithDetection_SkipSanitized(t *testing.T) {
	opt := deps.WithDetection(deps.Config{
		FilePatterns: []regex.Regex{regexp.MustCompile(".*")},
	})

	h := opt(func(hh []heartbeat.Heartbeat) ([]heartbeat.Result, error) {
		assert.Len(t, hh[0].Dependencies, 0)

		return []heartbeat.Result{
			{
				Status: 201,
			},
		}, nil
	})

	result, err := h([]heartbeat.Heartbeat{{
		Entity:     "testdata/golang.go",
		EntityType: heartbeat.FileType,
		Language:   heartbeat.LanguagePtr(heartbeat.LanguageGo),
	}})
	require.NoError(t, err)

	assert.Equal(t, []heartbeat.Result{
		{
			Status: 201,
		},
	}, result)
}

func TestWithDetection_LocalFile(t *testing.T) {
	opt := deps.WithDetection(deps.Config{})

	h := opt(func(hh []heartbeat.Heartbeat) ([]heartbeat.Result, error) {
		assert.Equal(t, []heartbeat.Heartbeat{
			{
				Dependencies: []string{
					"os",
					"github.com/wakatime/wakatime-cli/pkg/heartbeat",
				},
				Entity:     "testdata/golang.go",
				EntityType: heartbeat.FileType,
				Language:   heartbeat.LanguagePtr(heartbeat.LanguageGo),
				LocalFile:  "testdata/golang_minimal.go",
			},
		}, hh)

		return []heartbeat.Result{
			{
				Status: 201,
			},
		}, nil
	})

	result, err := h([]heartbeat.Heartbeat{{
		Entity:     "testdata/golang.go",
		EntityType: heartbeat.FileType,
		Language:   heartbeat.LanguagePtr(heartbeat.LanguageGo),
		LocalFile:  "testdata/golang_minimal.go",
	}})
	require.NoError(t, err)

	assert.Equal(t, []heartbeat.Result{
		{
			Status: 201,
		},
	}, result)
}

func TestWithDetection_NonFileType(t *testing.T) {
	opt := deps.WithDetection(deps.Config{})

	h := opt(func(hh []heartbeat.Heartbeat) ([]heartbeat.Result, error) {
		assert.Equal(t, []heartbeat.Heartbeat{
			{
				Entity:     "testdata/codefiles/golang.go",
				EntityType: heartbeat.AppType,
			},
		}, hh)

		return []heartbeat.Result{
			{
				Status: 201,
			},
		}, nil
	})

	result, err := h([]heartbeat.Heartbeat{{
		Entity:     "testdata/codefiles/golang.go",
		EntityType: heartbeat.AppType,
	}})
	require.NoError(t, err)

	assert.Equal(t, []heartbeat.Result{
		{
			Status: 201,
		},
	}, result)
}

func TestDetect(t *testing.T) {
	tests := map[string]struct {
		Filepath     string
		Language     heartbeat.Language
		Dependencies []string
	}{
		"c": {
			Filepath:     "testdata/c_minimal.c",
			Language:     heartbeat.LanguageC,
			Dependencies: []string{"openssl"},
		},
		"csharp": {
			Filepath:     "testdata/csharp_minimal.cs",
			Language:     heartbeat.LanguageCSharp,
			Dependencies: []string{"WakaTime"},
		},
		"elm": {
			Filepath:     "testdata/elm_minimal.elm",
			Language:     heartbeat.LanguageElm,
			Dependencies: []string{"Html"},
		},
		"java": {
			Filepath:     "testdata/java_minimal.java",
			Language:     heartbeat.LanguageJava,
			Dependencies: []string{"foobar"},
		},
		"golang": {
			Filepath: "testdata/golang_minimal.go",
			Language: heartbeat.LanguageGo,
			Dependencies: []string{
				"os",
				"github.com/wakatime/wakatime-cli/pkg/heartbeat",
			},
		},
		"haskell": {
			Filepath:     "testdata/haskell_minimal.hs",
			Language:     heartbeat.LanguageHaskell,
			Dependencies: []string{"Control"},
		},
		"haxe": {
			Filepath:     "testdata/haxe_minimal.hx",
			Language:     heartbeat.LanguageHaxe,
			Dependencies: []string{"alpha"},
		},
		"javascript": {
			Filepath:     "testdata/es6_minimal.js",
			Language:     heartbeat.LanguageJavaScript,
			Dependencies: []string{"bravo"},
		},
		"kotlin": {
			Filepath:     "testdata/kotlin_minimal.kt",
			Language:     heartbeat.LanguageKotlin,
			Dependencies: []string{"alpha.time"},
		},
		"php": {
			Filepath:     "testdata/php_minimal.php",
			Language:     heartbeat.LanguagePHP,
			Dependencies: []string{"Interop", "FooBarOne"},
		},
		"python": {
			Filepath:     "testdata/python_minimal.py",
			Language:     heartbeat.LanguagePython,
			Dependencies: []string{"flask", "simplejson"},
		},
		"rust": {
			Filepath:     "testdata/rust_minimal.rs",
			Language:     heartbeat.LanguageRust,
			Dependencies: []string{"syn"},
		},
		"scala": {
			Filepath:     "testdata/scala_minimal.scala",
			Language:     heartbeat.LanguageScala,
			Dependencies: []string{"com.alpha.SomeClass"},
		},
		"swift": {
			Filepath:     "testdata/swift_minimal.swift",
			Language:     heartbeat.LanguageSwift,
			Dependencies: []string{"Swift"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			deps, err := deps.Detect(test.Filepath, test.Language)
			require.NoError(t, err)

			assert.Equal(t, test.Dependencies, deps)
		})
	}
}

func TestDetect_DuplicatesRemoved(t *testing.T) {
	deps, err := deps.Detect(
		"testdata/golang_duplicate.go",
		heartbeat.LanguageGo,
	)
	require.NoError(t, err)

	assert.Equal(t, []string{
		"os",
	}, deps)
}

func TestDetect_LongDependenciesRemoved(t *testing.T) {
	deps, err := deps.Detect(
		"testdata/python_with_long_import.py",
		heartbeat.LanguagePython,
	)
	require.NoError(t, err)

	assert.Equal(t, []string{
		"django",
		"flask",
		// nolint:lll
		"notlongenoughnotlongenoughnotlongenoughnotlongenoughnotlongenoughnotlongenoughnotlongenoughnotlongenoughnotlongenoughnotlongenoughnotlongenoughnotlongenoughnotlongenoughnotlongenoughnotlongenoughnotlo",
	}, deps)
}
