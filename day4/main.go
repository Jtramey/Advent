package main

import (
	readAOC "../utils"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
	PassportId     string
	CountryId      string
}

func (p *Passport) isValid() bool {
	validBirthYear := validateYear(p.BirthYear, 1920, 2002)
	validIssueYear := validateYear(p.IssueYear, 2010, 2020)
	validExpirationYear := validateYear(p.ExpirationYear, 2020, 2030)
	validHeight := validateHeight(p.Height)
	validHairColor := validateHairColor(p.HairColor)
	validEyeColor := validateEyeColor(p.EyeColor)
	validPassportId := validatePassportId(p.PassportId)

	if validBirthYear == false {
		return false
	}
	if validIssueYear == false {
		return false
	}
	if validExpirationYear == false {
		return false
	}
	if validHeight == false {
		return false
	}
	if validHairColor == false {
		return false
	}
	if validEyeColor == false {
		return false
	}
	if validPassportId == false {
		return false
	}

	return true
}

func validateYear(year, minYear, maxYear int) bool {
	if year == 0 {
		return false
	}
	strYear := strconv.Itoa(year)
	if len(strYear) != 4 {
		return false
	}
	if (year < minYear) || (year > maxYear) {
		return false
	}

	return true
}

func validateHeight(hgt string) bool {
	if hgt == "" {
		return false
	}
	if !strings.Contains(hgt, "in") {
		if !strings.Contains(hgt, "cm") {
			return false
		}
	}
	unit := hgt[len(hgt)-2:]
	reg := regexp.MustCompile("[0-9]+")
	value, _ := strconv.Atoi(reg.FindAllString(hgt, -1)[0])
	if strings.EqualFold(unit, "cm") {
		if (value < 150) || (value > 193) {
			return false
		}
	}
	if strings.EqualFold(unit, "in") {
		if (value < 59) || (value > 76) {
			return false
		}
	}

	return true
}

func validateHairColor(hcl string) bool {
	if hcl == "" {
		return false
	}
	if hcl[0] != '#' {
		return false
	}
	_, err := strconv.ParseUint(strings.Replace(hcl, "#", "", 1), 16, 64)
	if err != nil {
		return false
	}

	return true
}

func validateEyeColor(ecl string) bool {
	validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	return readAOC.Contains(ecl, validColors)
}

func validatePassportId(pid string) bool {
	if len(pid) != 9 {
		return false
	}

	matched, err := regexp.Match(`\d{9}`, []byte(pid))
	if err != nil {
		return false
	}
	return matched
}

func main() {
	lines := readAOC.ReadInput("")
	var passports []Passport
	first := Passport{}
	passports = append(passports, first)
	blanks := 0
	for _, line := range lines {
		if line != "" {
			passports[blanks] = parseKV(line, passports[blanks])
		} else {
			blanks++
			passports = append(passports, Passport{})
			continue
		}
	}
	var validCount int
	for _, passport := range passports {
		if passport.isValid() {
			validCount++
		}
	}
	print(validCount)
}

func parseKV(line string, p Passport) Passport {
	if strings.Contains(line, " ") {
		fields := strings.Split(line, " ")
		for _, f := range fields {
			kv := strings.Split(f, ":")
			k := kv[0]
			v := kv[1]
			p = keySwitch(k, v, p)
		}
	} else {
		kv := strings.Split(line, ":")
		k := kv[0]
		v := kv[1]
		p = keySwitch(k, v, p)
	}
	return p
}

func keySwitch(k, v string, p Passport) Passport {
	switch k {
	case "byr":
		year, _ := strconv.Atoi(v)
		p.BirthYear = year
	case "iyr":
		year, _ := strconv.Atoi(v)
		p.IssueYear = year
	case "eyr":
		year, _ := strconv.Atoi(v)
		p.ExpirationYear = year
	case "hgt":
		p.Height = v
	case "hcl":
		p.HairColor = v
	case "ecl":
		p.EyeColor = v
	case "pid":
		p.PassportId = v
	case "cio":
		p.CountryId = v
	}
	return p
}
