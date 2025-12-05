package config

import "testing"

func TestFromEnv(t *testing.T) {
	tests := []struct {
		name          string
		envVars       map[string]string
		wantHost      string
		wantPort      uint16
		wantServeDocs bool
		wantErr       bool
	}{
		{
			name: "Success production config",
			envVars: map[string]string{
				"HOST": "localhost",
				"PORT": "8080",
				"ENV":  "production",
			},
			wantHost:      "localhost",
			wantPort:      8080,
			wantServeDocs: false,
			wantErr:       false,
		},
		{
			name: "Success dev config (ServeDocs true)",
			envVars: map[string]string{
				"HOST": "localhost",
				"PORT": "3000",
				"ENV":  "dev",
			},
			wantHost:      "localhost",
			wantPort:      3000,
			wantServeDocs: true,
			wantErr:       false,
		},
		{
			name: "Success staging config (ServeDocs true)",
			envVars: map[string]string{
				"HOST": "localhost",
				"PORT": "8080",
				"ENV":  "staging",
			},
			wantHost:      "localhost",
			wantPort:      8080,
			wantServeDocs: true,
			wantErr:       false,
		},
		{
			name: "Success empty ENV (ServeDocs true)",
			envVars: map[string]string{
				"HOST": "localhost",
				"PORT": "8080",
				"ENV":  "",
			},
			wantHost:      "localhost",
			wantPort:      8080,
			wantServeDocs: true,
			wantErr:       false,
		},
		{
			name: "Missing HOST",
			envVars: map[string]string{
				"PORT": "8080",
			},
			wantErr: true,
		},
		{
			name: "Missing PORT",
			envVars: map[string]string{
				"HOST": "localhost",
			},
			wantErr: true,
		},
		{
			name: "Invalid PORT",
			envVars: map[string]string{
				"HOST": "localhost",
				"PORT": "not-a-number",
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			for k, v := range tc.envVars {
				t.Setenv(k, v)
			}

			// Act
			cfg, err := FromEnv()

			// Assert
			if (err != nil) != tc.wantErr {
				t.Errorf("FromEnv() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if tc.wantErr {
				return
			}

			if cfg.Host() != tc.wantHost {
				t.Errorf("FromEnv() Host = %v, want %v", cfg.Host(), tc.wantHost)
			}
			if cfg.Port() != tc.wantPort {
				t.Errorf("FromEnv() Port = %v, want %v", cfg.Port(), tc.wantPort)
			}
			if cfg.WillServeDocs() != tc.wantServeDocs {
				t.Errorf("FromEnv() ServeDocs = %v, want %v", cfg.WillServeDocs(), tc.wantServeDocs)
			}
		})
	}
}
