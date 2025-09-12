package supabase

import (
	"errors"

	"github.com/absendulu-project/backend/pkg/config"
	supabase "github.com/supabase-community/supabase-go"
)

func NewSupabaseClient() (*supabase.Client, error) {
	cfg := config.GetConfig()

	supabaseURL := cfg.Supabase.SupabaseURL
	if supabaseURL == "" {
		return nil, errors.New("SUPABASE_URL is not set")
	}

	supabaseKey := cfg.Supabase.SupabaseKey
	if supabaseKey == "" {
		return nil, errors.New("SUPABASE_KEY is not set")
	}

	client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
