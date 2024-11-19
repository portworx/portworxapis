package common

import (
	"strings"
	"testing"

	pv "github.com/bufbuild/protovalidate-go"
	"github.com/stretchr/testify/require"
)

// TestMetaNameValidate tests the Meta.Name validation
func TestMetaNameValidate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// Valid cases
		{
			name:    "with spaces", // allowed spaces
			wantErr: false,
		},
		{
			name:    "with multiple spaces", // allowed multiple spaces
			wantErr: false,
		},
		{
			name:    "test", // allowed only alphabets
			wantErr: false,
		},
		{
			name:    "test123", // allowed alphanumeric
			wantErr: false,
		},
		{
			name:    "test-123-12", // allowed hyphen
			wantErr: false,
		},
		{
			name:    "Test_123_123", // allowed underscore
			wantErr: false,
		},
		{
			name:    "Test", // case sensitive
			wantErr: false,
		},
		{
			name:    "123", // allow only numbers
			wantErr: false,
		},
		// Invalid cases
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
		{
			name:    "test-", // must end with alphanumeric
			wantErr: true,
		},
		{
			name:    "test ", // must end with alphanumeric
			wantErr: true,
		},
		{
			name:    "-test", // must start with alphanumeric
			wantErr: true,
		},
		{
			name:    "_test", // must start with alphanumeric
			wantErr: true,
		},
		{
			name:    "asd&*()+/!@", // special characters other than '-' and '_' not allowed
			wantErr: true,
		},
	}

	v, err := pv.New()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &Meta{
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

func TestMetaDescriptionValidate(t *testing.T) {
	tests := []struct {
		description string
		wantErr     bool
	}{
		// Valid cases
		{
			description: "",
			wantErr:     false,
		},
		{
			description: "test",
			wantErr:     false,
		},
		{
			description: "This is a very long description of the resource that is used to validate the maximum length of the description field",
			wantErr:     true,
		},
	}

	v, err := pv.New()

	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			m := &Meta{
				Name:        "test",
				Description: tt.description,
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

func TestLabelsNames(t *testing.T) {
	successCases := []string{
		"simple",
		"now-with-dashes",
		"1-starts-with-num",
		"1234",
		"simple/simple",
		"now-with-dashes/simple",
		"now-with-dashes/now-with-dashes",
		"now.with.dots/simple",
		"now-with.dashes-and.dots/simple",
		"1-num.2-num/3-num",
		"1234/5678",
		"1.2.3.4/5678",
		"Uppercase_Is_OK_123",
		"example.com/Uppercase_Is_OK_123",
		"requests.storage-foo",
		strings.Repeat("a", 63),
	}

	errorCases := []string{
		"nospecialchars%^=@",
		"cantendwithadash-",
		"-cantstartwithadash-",
		"only/one/slash",
		"Example.com/abc",
		"example_com/abc",
		"example.com/",
		"/simple",
		strings.Repeat("a", 64),
	}

	v, err := pv.New()
	require.NoError(t, err)

	for _, label := range successCases {
		t.Run(label, func(t *testing.T) {
			m := &Meta{
				Name:        "test",
				Description: "test",
				Labels: map[string]string{
					label: "value",
				},
			}

			err := v.Validate(m)
			require.NoError(t, err)
		})
	}

	for _, label := range errorCases {
		t.Run(label, func(t *testing.T) {
			m := &Meta{
				Name:        "test",
				Description: "test",
				Labels: map[string]string{
					label: "value",
				},
			}

			err := v.Validate(m)
			require.Error(t, err)
		})
	}
}

func TestIsValidLabelValue(t *testing.T) {
	successCases := []string{
		"simple",
		"now-with-dashes",
		"1-starts-with-num",
		"end-with-num-1",
		"1234",                  // only num
		strings.Repeat("a", 63), // to the limit
		"",
	}

	errorCases := []string{
		"nospecialchars%^=@",
		"Tama-nui-te-rā.is.Māori.sun",
		"\\backslashes\\are\\bad",
		"-starts-with-dash",
		"ends-with-dash-",
		".starts.with.dot",
		"ends.with.dot.",
		strings.Repeat("a", 64), // over the limit
	}

	v, err := pv.New()
	require.NoError(t, err)

	for i := range successCases {
		t.Run(successCases[i], func(t *testing.T) {
			m := &Meta{
				Name:        "test",
				Description: "test",
				Labels: map[string]string{
					"key": successCases[i],
				},
			}

			err := v.Validate(m)
			require.NoError(t, err)
		})
	}

	for i := range errorCases {
		t.Run(errorCases[i], func(t *testing.T) {
			m := &Meta{
				Name:        "test",
				Description: "test",
				Labels: map[string]string{
					"key": errorCases[i],
				},
			}

			err := v.Validate(m)
			require.Error(t, err)
		})
	}
}

func TestIsValidateAnnotationKey(t *testing.T) {
	successCases := []string{
		"simple",
		"now-with-dashes",
		"1-starts-with-num",
		"1234",
		"simple/simple",
		"now-with-dashes/simple",
		"now-with-dashes/now-with-dashes",
		"now.with.dots/simple",
		"now-with.dashes-and.dots/simple",
		"1-num.2-num/3-num",
		"1234/5678",
		"1.2.3.4/5678",
		"Uppercase_Is_OK_123",
		"example.com/Uppercase_Is_OK_123",
		"requests.storage-foo",
		strings.Repeat("a", 63),
	}

	errorCases := []string{
		"nospecialchars%^=@",
		"cantendwithadash-",
		"-cantstartwithadash-",
		"only/one/slash",
		"Example.com/abc",
		"example_com/abc",
		"example.com/",
		"/simple",
		strings.Repeat("a", 64),
	}

	v, err := pv.New()
	require.NoError(t, err)

	for _, tt := range successCases {
		t.Run(tt, func(t *testing.T) {
			m := &Meta{
				Name:        "test",
				Description: "test",
				Annotations: map[string]string{
					tt: "value",
				},
			}

			err := v.Validate(m)
			require.NoError(t, err)
		})
	}

	for _, tt := range errorCases {
		t.Run(tt, func(t *testing.T) {
			m := &Meta{
				Name:        "test",
				Description: "test",
				Annotations: map[string]string{
					tt: "value",
				},
			}

			err := v.Validate(m)
			require.Error(t, err)
		})
	}
}
