package restore

import (
	"testing"

	pv "github.com/bufbuild/protovalidate-go"
	"github.com/stretchr/testify/require"
)

func TestValidateRecreateRestoreName(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// Valid cases
		{
			name:    "test", // allowed only alphabets
			wantErr: false,
		},
		{
			name:    "test123", // allowed alphanumeric
			wantErr: false,
		},
		{
			name:    "test-123", // allowed hyphen
			wantErr: false,
		},
		{
			name:    "Test_123", // allowed underscore
			wantErr: false,
		},
		{
			name:    "Test", // case sensitive
			wantErr: false,
		},
		// Invalid cases
		{
			name:    "123", // required alphabet
			wantErr: true,
		},
		{
			name:    "_-_", // required alphabet
			wantErr: true,
		},
		{
			name:    "a", // required min length
			wantErr: true,
		},
		{
			name:    "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmno", // required < max length
			wantErr: true,
		},
	}

	v, err := pv.New()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &RecreateRestoreRequest{
				Name: tt.name,
			}

			err := v.Validate(m)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
