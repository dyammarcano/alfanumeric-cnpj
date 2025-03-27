package cnpj

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
	pesosDV           = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjZerado        = "00000000000000"
)

func isValidCharSet(value string) bool {
	return !regexNaoPermitido.MatchString(value)
}

func removeMascaraCNPJ(value string) string {
	return strings.ToUpper(regexMascara.ReplaceAllString(value, ""))
}

func CalculateDV(value string) (string, error) {
	if !isValidCharSet(value) {
		return "", ErroDVInvalido
	}

	semMascara := removeMascaraCNPJ(value)
	if len(semMascara) == 14 {
		semMascara = semMascara[:12]
	}

	if !regexCNPJSemDV.MatchString(semMascara) {
		return "", ErroDVInvalido
	}

	if !regexCNPJSemDV.MatchString(semMascara) || semMascara == cnpjZerado[:12] {
		return "", ErroDVInvalido
	}

	//// Verifica se há ao menos um caractere alfabético (A–Z)
	//hasAlpha := false
	//for _, r := range semMascara {
	//	if r >= 'A' && r <= 'Z' {
	//		hasAlpha = true
	//		break
	//	}
	//}
	//if !hasAlpha {
	//	return "", fmt.Errorf("o CNPJ deve conter pelo menos um caractere alfabético (A-Z)")
	//}

	somaDV1, somaDV2, j := 0, 0, 0

	for i := 0; i < 12; i++ {
		somaDV1 += int(rune(semMascara[i])-48) * pesosDV[j+1]
		somaDV2 += int(rune(semMascara[i])-48) * pesosDV[j]
		j = (j + 1) % len(pesosDV)
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

func IsValid(value string) bool {
	if !isValidCharSet(value) {
		return false
	}

	var dv string

	semMascara := removeMascaraCNPJ(value)
	if regexCNPJ.MatchString(semMascara) {
		dv = semMascara[12:]
		semMascara = semMascara[:12]
	}

	if !regexCNPJSemDV.MatchString(semMascara) {
		return false
	}

	dvCalculado, err := CalculateDV(semMascara)
	if err != nil {
		return false
	}
	return dv == dvCalculado
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
