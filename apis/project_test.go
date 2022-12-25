package apis_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

func TestProjectsList(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:            "unauthorized",
			Method:          http.MethodGet,
			Url:             "/api/projects",
			ExpectedStatus:  401,
			ExpectedContent: []string{`"data":{}`},
		},
		{
			Name:   "authorized as user",
			Method: http.MethodGet,
			Url:    "/api/projects",
			RequestHeaders: map[string]string{
				"Authorization": "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6IjRxMXhsY2xtZmxva3UzMyIsInR5cGUiOiJhdXRoUmVjb3JkIiwiY29sbGVjdGlvbklkIjoiX3BiX3VzZXJzX2F1dGhfIiwiZXhwIjoyMjA4OTg1MjYxfQ.UwD8JvkbQtXpymT09d7J6fdA0aP9g4FJ1GPh_ggEkzc",
			},
			ExpectedStatus:  401,
			ExpectedContent: []string{`"data":{}`},
		},
		{
			Name:   "authorized as admin",
			Method: http.MethodGet,
			Url:    "/api/projects",
			RequestHeaders: map[string]string{
				"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InN5d2JoZWNuaDQ2cmhtMCIsInR5cGUiOiJhZG1pbiIsImV4cCI6MjIwODk4NTI2MX0.M1m--VOqGyv0d23eeUc0r9xE8ZzHaYVmVFw1VZW6gT8",
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"page":1`,
				`"perPage":30`,
				`"totalItems":7`,
				`"items":[{`,
				`"id":"_pb_users_auth_"`,
				`"id":"v851q4r790rhknl"`,
				`"id":"kpv709sk2lqbqk8"`,
				`"id":"wsmn24bux7wo113"`,
				`"id":"sz5l5z67tg7gku0"`,
				`"id":"wzlqyes4orhoygb"`,
				`"id":"4d1blo5cuycfaca"`,
				`"type":"auth"`,
				`"type":"base"`,
			},
			ExpectedEvents: map[string]int{
				"OnCollectionsListRequest": 1,
			},
		},
		{
			Name:   "authorized as admin + paging and sorting",
			Method: http.MethodGet,
			Url:    "/api/projects?page=1&perPage=2&sort=-created",
			RequestHeaders: map[string]string{
				"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InN5d2JoZWNuaDQ2cmhtMCIsInR5cGUiOiJhZG1pbiIsImV4cCI6MjIwODk4NTI2MX0.M1m--VOqGyv0d23eeUc0r9xE8ZzHaYVmVFw1VZW6gT8",
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"page":2`,
				`"perPage":2`,
				`"totalItems":7`,
				`"items":[{`,
				`"id":"4d1blo5cuycfaca"`,
				`"id":"wzlqyes4orhoygb"`,
			},
			ExpectedEvents: map[string]int{
				"OnCollectionsListRequest": 1,
			},
		},
		{
			Name:   "authorized as admin + invalid filter",
			Method: http.MethodGet,
			Url:    "/api/projects?filter=invalidfield~'demo2'",
			RequestHeaders: map[string]string{
				"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InN5d2JoZWNuaDQ2cmhtMCIsInR5cGUiOiJhZG1pbiIsImV4cCI6MjIwODk4NTI2MX0.M1m--VOqGyv0d23eeUc0r9xE8ZzHaYVmVFw1VZW6gT8",
			},
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
		},
		{
			Name:   "authorized as admin + valid filter",
			Method: http.MethodGet,
			Url:    "/api/projects?filter=name~'demo'",
			RequestHeaders: map[string]string{
				"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InN5d2JoZWNuaDQ2cmhtMCIsInR5cGUiOiJhZG1pbiIsImV4cCI6MjIwODk4NTI2MX0.M1m--VOqGyv0d23eeUc0r9xE8ZzHaYVmVFw1VZW6gT8",
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"page":1`,
				`"perPage":30`,
				`"totalItems":4`,
				`"items":[{`,
				`"id":"wsmn24bux7wo113"`,
				`"id":"sz5l5z67tg7gku0"`,
				`"id":"wzlqyes4orhoygb"`,
				`"id":"4d1blo5cuycfaca"`,
			},
			ExpectedEvents: map[string]int{
				"OnCollectionsListRequest": 1,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestProjectCreate(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:            "unauthorized",
			Method:          http.MethodPost,
			Url:             "/api/projects",
			ExpectedStatus:  401,
			ExpectedContent: []string{`"data":{}`},
		},
		{
			Name:   "authorized as user",
			Method: http.MethodPost,
			Url:    "/api/projects",
			RequestHeaders: map[string]string{
				"Authorization": "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6IjRxMXhsY2xtZmxva3UzMyIsInR5cGUiOiJhdXRoUmVjb3JkIiwiY29sbGVjdGlvbklkIjoiX3BiX3VzZXJzX2F1dGhfIiwiZXhwIjoyMjA4OTg1MjYxfQ.UwD8JvkbQtXpymT09d7J6fdA0aP9g4FJ1GPh_ggEkzc",
			},
			ExpectedStatus:  401,
			ExpectedContent: []string{`"data":{}`},
		},
		{
			Name:   "authorized as admin + invalid data (eg. existing name)",
			Method: http.MethodPost,
			Url:    "/api/projects",
			Body:   strings.NewReader(`{"name":"testing"}`),
			RequestHeaders: map[string]string{
				"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InN5d2JoZWNuaDQ2cmhtMCIsInR5cGUiOiJhZG1pbiIsImV4cCI6MjIwODk4NTI2MX0.M1m--VOqGyv0d23eeUc0r9xE8ZzHaYVmVFw1VZW6gT8",
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"data":{`,
				`"name":{"code":"validation_project_name_exists"`,
			},
		},
		{
			Name:   "authorized as admin + valid data",
			Method: http.MethodPost,
			Url:    "/api/projects",
			Body:   strings.NewReader(`{"name":"test_project_go"}`),
			RequestHeaders: map[string]string{
				"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InN5d2JoZWNuaDQ2cmhtMCIsInR5cGUiOiJhZG1pbiIsImV4cCI6MjIwODk4NTI2MX0.M1m--VOqGyv0d23eeUc0r9xE8ZzHaYVmVFw1VZW6gT8",
			},
			ExpectedStatus: 200,
			// ExpectedContent: []string{
			// 	`"id":`,
			// 	`"name":"test_project_go"`,
			// },
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
		print(scenario.Body)
	}
}
