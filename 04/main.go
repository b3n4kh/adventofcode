package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Passport struct
type Passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func validEyeColor(color string) bool {
	validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, allowedColor := range validEyeColors {
		if color == allowedColor {
			return true
		}
	}
	return false
}

func setField(v *Passport, field string, value string) bool {
	r := reflect.ValueOf(v)
	exportedField := strings.Title(field)
	structField := reflect.Indirect(r).FieldByName(exportedField)
	if !structField.IsValid() || !structField.CanSet() {
		//return false
	}
	val := reflect.ValueOf(value)
	structField.Set(val)
	return true
}

func hasNoEmptyField(v *Passport) bool {
	r := reflect.ValueOf(v).Elem()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		if f.String() == "" {
			return false
		}
	}
	return true
}

func parsepassport(input string) (pass Passport) {
	fields := strings.Fields(input)

	for _, field := range fields {
		fieldSlice := strings.Split(field, ":")
		_ = setField(&pass, fieldSlice[0], fieldSlice[1])
	}
	return pass
}

func isPassValid(pass Passport) bool {
	if pass.Cid == "" {
		pass.Cid = "1337"
	}
	if !hasNoEmptyField(&pass) {
		return false
	}

	byr, err := strconv.Atoi(pass.Byr)
	if (byr < 1920) || (byr > 2002) || err != nil {
		return false
	}

	iyr, err := strconv.Atoi(pass.Iyr)
	if (iyr < 2010) || (iyr > 2020) || err != nil {
		return false
	}

	eyr, err := strconv.Atoi(pass.Eyr)
	if (eyr < 2020) || (eyr > 2030) || err != nil {
		return false
	}

	hgt, err := strconv.Atoi(pass.Hgt[0 : len(pass.Hgt)-2])
	hgtUnit := pass.Hgt[len(pass.Hgt)-2:]
	if hgtUnit == "cm" {
		if (hgt < 150) || (hgt > 193) || err != nil {
			return false
		}
	}

	if hgtUnit == "in" {
		if (hgt < 59) || (hgt > 76) || err != nil {
			return false
		}
	}
	_, err = strconv.ParseUint(pass.Hcl[1:], 16, 64)
	if (pass.Hcl[0] != '#') || (len(pass.Hcl) != 7) || err != nil {
		return false
	}

	if !validEyeColor(pass.Ecl) {
		return false
	}

	_, err = strconv.Atoi(pass.Pid)
	if err != nil || (len(pass.Pid) != 9) {
		return false
	}
	return true
}

func main() {
	dat, err := os.Open("input.txt")
	check(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line, passport string
	var validPassport, scannedPassport int
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			scannedPassport++
			if isPassValid(parsepassport(passport)) {
				validPassport++
			} else {
				// fmt.Printf("INVALID: %+v\n\n", passport)
			}
			passport = ""
		}
		passport += line + " "
	}
	fmt.Printf("\n\nScanned Passports: %v\n", scannedPassport)

	fmt.Printf("Valid Passports: %v", validPassport)
}
