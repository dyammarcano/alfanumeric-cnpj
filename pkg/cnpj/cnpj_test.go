package cnpj

import (
	"errors"
	"testing"
)

func TestCalculateDV(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"TK.10B.O3I/H1GA", "13"},
		{"PF.0YG.0F8/C4WB", "92"},
		{"1Q.7ZW.SVW/QR9O", "87"},
		{"TK10BO3IH1GA", "13"},
		{"PF0YG0F8C4WB", "92"},
		{"1Q7ZWSVWQR9O", "87"},
	}

	for _, tt := range tests {
		dv, err := CalculateDV(tt.input)
		if err != nil {
			t.Errorf("CalculateDV(%s) = %s, expected %s, error: %v", tt.input, dv, tt.expected, err)
			return
		}

		if dv != tt.expected {
			t.Errorf("CalculateDV(%s) = %s, expected %s", tt.input, dv, tt.expected)
		}
	}
}

// TestCalculateDV_InvalidCases tests error cases for calculateDV
func TestCalculateDV_InvalidCases(t *testing.T) {
	invalidInputs := []string{
		"",               // Empty
		"'!@#$%&*-_=+^~", // Only invalid characters
		"$0123456789A",   // Invalid character at the beginning
		"012345?6789A",   // Invalid character in the middle
		"0123456789A#",   // Invalid character at the end
		"12ABc34501DE",   // Contains lowercase letters
		"00000000000",    // Too few digits
		"0000000000191",  // Too many digits
	}

	for _, input := range invalidInputs {
		_, err := CalculateDV(input)
		if !errors.Is(err, ErroDVInvalido) {
			t.Error("expected no error")
			continue
		}
	}
}

// TestValidateCNPJ tests the validity of different CNPJs
func TestValidateCNPJ(t *testing.T) {
	validCNPJs := []string{
		"12.ABC.345/01DE-35",
		"90.021.382/0001-22",
		"90.024.778/0001-23",
		"90.025.108/0001-21",
		"90.025.255/0001-00",
		"90.024.420/0001-09",
		"90.024.781/0001-47",
		"04.740.714/0001-97",
		"44.108.058/0001-29",
		"90.024.780/0001-00",
		"90.024.779/0001-78",
		"00000000000191",
		"ABCDEFGHIJKL80",
	}

	for _, v := range validCNPJs {
		if !IsValid(v) {
			t.Errorf("ValidateCNPJ(%s) should return true, but returned false", v)
		}
	}
}

// TestValidateCNPJ_Invalid tests invalid CNPJs
func TestValidateCNPJ_Invalid(t *testing.T) {
	invalidCNPJs := []string{
		"",                   // Empty
		"'!@#$%&*-_=+^~",     // Only invalid characters
		"$0123456789ABC",     // Invalid character at the beginning
		"0123456?789ABC",     // Invalid character in the middle
		"0123456789ABC#",     // Invalid character at the end
		"12.ABc.345/01DE-35", // Contains lowercase letters
		"0000000000019",      // Too few digits
		"000000000001911",    // Too many digits
		"0000000000019L",     // Letter instead of second DV
		"000000000001P1",     // Letter instead of first DV
		"00000000000192",     // Invalid check digit
		"ABCDEFGHIJKL81",     // Invalid check digit
		"00000000000000",     // All zeroes
		"00.000.000/0000-00", // All zeroes with mask
	}

	for _, v := range invalidCNPJs {
		if IsValid(v) {
			t.Errorf("ValidateCNPJ(%s) should return false, but returned true", v)
		}
	}
}
