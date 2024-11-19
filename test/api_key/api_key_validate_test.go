package api_key_test

import (
	"RuRu/internal/api/guard/api_key"
	"RuRu/internal/logger"
	"testing"
)

func TestValidateApiKey(t *testing.T) {
	t.Setenv("ENV", "test")
	t.Setenv("API_KEY", "a9d1a556-4850-459d-8a33-0b66099e9684")
	t.Setenv("API_KEY_SALT", "G-STAR^_^")
	log := logger.SetupLogger()

	t.Log("Test of validation apiKey")
	{
		testId := 1
		t.Logf("Test id: %d;\t Validate When apiKey OK", testId)
		{
			OkKey, _ := api_key.Generate(log)
			_, err := api_key.Validate(OkKey, log)
			if err != nil {
				t.Fatal(err.Error())
			}
		}

		testId++
		t.Logf("Test id: %d;\t Validate When apiKey NOT OK", testId)
		{
			OkKey, _ := api_key.GenerateFrom("BEBE", log)
			ok, _ := api_key.Validate(OkKey, log)
			if ok {
				t.Fatal("Validation not ok!")
			}
		}
	}
}
