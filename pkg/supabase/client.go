package supabase

import (
	"errors"
	"os"

	supabase "github.com/supabase-community/supabase-go"
)

func NewSupabaseClient() (*supabase.Client, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	if supabaseURL == "" {
		return nil, errors.New("SUPABASE_URL is not set")
	}

	supabaseKey := os.Getenv("SUPABASE_KEY")
	if supabaseKey == "" {
		return nil, errors.New("SUPABASE_KEY is not set")
	}

	client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
