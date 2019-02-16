package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dylankilkenny/watch-cash/server/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	InitDB()
	r = router.SetupRouter(db, false)
	code := m.Run()
	os.Exit(code)
}

func TestLoginRoute(t *testing.T) {

	credentials := map[string]string{"email": "email@email.com", "password": "password"}

	makeRequest := func(t *testing.T, credentials map[string]string) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(credentials)
		req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		return w.Code
	}

	assertCorrectStatusCode := func(t *testing.T, want, got int) {
		t.Helper()
		ok := assert.Equal(t, want, got)
		if !ok {
			t.Errorf("Status Code is not %v. Got %v", want, got)
		}
	}

	t.Run("Successful login", func(t *testing.T) {
		_ = CreateUser()
		want := 200
		got := makeRequest(t, credentials)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Email or Password missing", func(t *testing.T) {
		emptyCredentials := map[string]string{"email": "", "password": ""}
		want := 400
		got := makeRequest(t, emptyCredentials)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("Email Not Found", func(t *testing.T) {
		want := 404
		got := makeRequest(t, credentials)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("Wrong password", func(t *testing.T) {
		credentials["password"] = "pass"
		_ = CreateUser()
		want := 401
		got := makeRequest(t, credentials)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}

func TestSignUpRoute(t *testing.T) {

	signup := map[string]string{"firstname": "dylan", "lastname": "kilkenny", "email": "email@email.com", "password": "password"}

	makeRequest := func(t *testing.T, credentials map[string]string) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(credentials)
		req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		return w.Code
	}

	assertCorrectStatusCode := func(t *testing.T, want, got int) {
		t.Helper()
		ok := assert.Equal(t, want, got)
		if !ok {
			t.Errorf("Status Code is not %v. Got %v", want, got)
		}
	}

	t.Run("Successful Sign Up", func(t *testing.T) {
		want := 200
		got := makeRequest(t, signup)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("Invalid form", func(t *testing.T) {
		signup = map[string]string{"email": "", "password": ""}
		want := 400
		got := makeRequest(t, signup)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("Email already exists", func(t *testing.T) {
		_ = CreateUser()
		want := 400
		got := makeRequest(t, signup)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}

func TestSubscribeToAddressRoute(t *testing.T) {

	address := map[string]string{"address": "qrwd7tucj2l6rjcgv5cr2n4t8ws83ghsjqpar98qpt"}

	makeRequest := func(t *testing.T, credentials map[string]string, token string) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(credentials)
		req, _ := http.NewRequest("POST", "/api/private/address", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)
		return w.Code
	}

	assertCorrectStatusCode := func(t *testing.T, want, got int) {
		t.Helper()
		ok := assert.Equal(t, want, got)
		if !ok {
			t.Errorf("Status Code is not %v. Got %v", want, got)
		}
	}

	t.Run("Successful subscribe to address", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		want := 200
		got := makeRequest(t, address, token)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid token", func(t *testing.T) {
		token := "invalid"
		want := 401
		got := makeRequest(t, address, token)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("Already subscribed", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		AddAddress(user)
		want := 400
		got := makeRequest(t, address, token)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

}

func TestGetSubscribedAddressesRoute(t *testing.T) {

	makeRequest := func(t *testing.T, token string) int {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/private/address", nil)
		req.Header.Add("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)
		return w.Code
	}

	assertCorrectStatusCode := func(t *testing.T, want, got int) {
		t.Helper()
		ok := assert.Equal(t, want, got)
		if !ok {
			t.Errorf("Status Code is not %v. Got %v", want, got)
		}
	}

	t.Run("Successful fetch of subsribed addresses", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		AddAddress(user)
		want := 200
		got := makeRequest(t, token)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid token", func(t *testing.T) {
		token := "invalid"
		want := 401
		got := makeRequest(t, token)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("No subscribed addresses found", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		want := 401
		got := makeRequest(t, token)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}

func TestRemoveSubscribedAddressRoute(t *testing.T) {

	address := map[string]string{"address": "qrwd7tucj2l6rjcgv5cr2n4t8ws83ghsjqpar98qpt"}

	makeRequest := func(t *testing.T, credentials map[string]string, token string) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(credentials)
		req, _ := http.NewRequest("POST", "/api/private/remove", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)
		return w.Code
	}

	assertCorrectStatusCode := func(t *testing.T, want, got int) {
		t.Helper()
		ok := assert.Equal(t, want, got)
		if !ok {
			t.Errorf("Status Code is not %v. Got %v", want, got)
		}
	}

	t.Run("Successful removal of subsribed addresses", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		AddAddress(user)
		want := 200
		got := makeRequest(t, address, token)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid token", func(t *testing.T) {
		token := "invalid"
		want := 401
		got := makeRequest(t, address, token)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("Address not found", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		want := 400
		got := makeRequest(t, address, token)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}

func TestValidateTokenRoute(t *testing.T) {

	makeRequest := func(t *testing.T, credentials map[string]string) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(credentials)
		req, _ := http.NewRequest("POST", "/api/validate", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		return w.Code
	}

	assertCorrectStatusCode := func(t *testing.T, want, got int) {
		t.Helper()
		ok := assert.Equal(t, want, got)
		if !ok {
			t.Errorf("Status Code is not %v. Got %v", want, got)
		}
	}

	t.Run("Token is valid", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		body := map[string]string{"token": token}
		want := 200
		got := makeRequest(t, body)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid token", func(t *testing.T) {
		token := "invalid"
		want := 401
		body := map[string]string{"token": token}
		got := makeRequest(t, body)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("Token is missing", func(t *testing.T) {
		want := 400
		body := map[string]string{"": ""}
		got := makeRequest(t, body)
		assertCorrectStatusCode(t, want, got)
	})
}
