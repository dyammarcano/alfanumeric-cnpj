package cnpjalfanumerico

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var (
	ErroDVInvalido    = errors.New("não é possível calcular o DV pois o CNPJ fornecido é inválido")
	regexCNPJSemDV    = regexp.MustCompile(`^[A-Z\d]{12}$`)
	regexCNPJ         = regexp.MustCompile(`^[A-Z\d]{12}\d{2}$`)
	regexMascara      = regexp.MustCompile(`[./-]`)
	regexNaoPermitido = regexp.MustCompile(`[^A-Z\d./-]`)
	cnpjZerado        = "00000000000000"
	pesosDV           = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

func isValidCharSet(value string) bool {
	return !regexNaoPermitido.MatchString(value)
}

func removeMascaraCNPJ(value string) string {
	return strings.ToUpper(regexMascara.ReplaceAllString(value, ""))
}

func CalculateDV(cnpj string) (string, error) {
	if !isValidCharSet(cnpj) {
		return "", ErroDVInvalido
	}

	semMascara := removeMascaraCNPJ(cnpj)
	if !regexCNPJSemDV.MatchString(semMascara) || semMascara == cnpjZerado[:12] {
		return "", ErroDVInvalido
	}

	somaDV1 := 0
	somaDV2 := 0
	for i := 0; i < 12; i++ {
		digito := int(semMascara[i] - '0')
		if digito < 0 || digito > 35 {
			if semMascara[i] >= 'A' && semMascara[i] <= 'Z' {
				digito = int(semMascara[i] - 'A' + 10)
			} else {
				return "", ErroDVInvalido
			}
		}
		somaDV1 += digito * pesosDV[i+1]
		somaDV2 += digito * pesosDV[i]
	}

	dv1 := somaDV1 % 11
	if dv1 < 2 {
		dv1 = 0
	} else {
		dv1 = 11 - dv1
	}

	somaDV2 += dv1 * pesosDV[12]
	dv2 := somaDV2 % 11
	if dv2 < 2 {
		dv2 = 0
	} else {
		dv2 = 11 - dv2
	}

	return fmt.Sprintf("%d%d", dv1, dv2), nil
}

func IsValid(cnpj string) bool {
	if !isValidCharSet(cnpj) {
		return false
	}

	semMascara := removeMascaraCNPJ(cnpj)
	if !regexCNPJ.MatchString(semMascara) || semMascara == cnpjZerado {
		return false
	}

	dvCalculado, err := CalculateDV(semMascara[:12])
	if err != nil {
		return false
	}
	return semMascara[12:] == dvCalculado
}

func FormatCNPJ(value string) string {
	value = removeMascaraCNPJ(value)
	if len(value) != 14 {
		return "CNPJ inválido"
	}

	mask := "##.###.###/####-##"
	runMask := make([]rune, len(mask))
	idx := 0
	for i, r := range mask {
		if r == '#' {
			runMask[i] = rune(value[idx])
			idx++
		} else {
			runMask[i] = r
		}
	}
	return string(runMask)
}

func UnformattedCNPJ(value string) string {
	return strings.ToUpper(regexp.MustCompile(`[^0-9A-Z]`).ReplaceAllString(value, ""))
}

func GenerateCNPJ() string {
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 12; i++ {
		sb.WriteByte(alphabet[rand.Intn(len(alphabet))])
	}
	base := sb.String()
	dv, err := CalculateDV(base)
	if err != nil {
		return ""
	}
	return base + dv
}
