package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hcl string
	hgt string
	ecl string
	pid string
	cid string
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	passportSlice := make([]Passport, 0)

	var passports int
	var currentPassport Passport
	for scanner.Scan() {

		line := scanner.Text()

		data := strings.Split(line, " ")
		if data[0] == "" {
			passportSlice = append(passportSlice, currentPassport)
			currentPassport = Passport{}
			passports++
		} else {
			//fmt.Println(line)
			for _, a := range data {
				switch a[0:3] {
				case "byr":
					currentPassport.byr = a[4:]
				case "iyr":
					currentPassport.iyr = a[4:]
				case "eyr":
					currentPassport.eyr = a[4:]
				case "hcl":
					currentPassport.hcl = a[4:]
				case "hgt":
					currentPassport.hgt = a[4:]
				case "ecl":
					currentPassport.ecl = a[4:]
				case "pid":
					currentPassport.pid = a[4:]
				case "cid":
					currentPassport.cid = a[4:]
				default:
					fmt.Println(a[0:3])
				}
			}
		}
	}
	// Add last passport
	passportSlice = append(passportSlice, currentPassport)

	valid := 0
	notvalid := 0
	for _, pass := range passportSlice {
		if validatePassport(pass) {
			valid = valid + 1
		} else {
			notvalid = notvalid + 1
		}
	}
	fmt.Println("Valid Passports", valid)
	fmt.Println("Not valid Passports", notvalid)

}

func validatePassport(p Passport) (valid bool) {
	valid = true

	valid = validateBYR(p.byr) && validateIYR(p.iyr) && validateEYR(p.eyr) &&
		validateHCL(p.hcl) && validateHGT(p.hgt) &&
		validateECL(p.ecl) && validatePID(p.pid)

	return
}

func validateBYR(byrr string) (valid bool) {
	valid = true
	byr, _ := strconv.Atoi(byrr)
	if byr != 0 {
		if !(byr < 2003 && byr > 1919) {
			valid = false
		}
	} else {
		valid = false
	}
	return
}
func validateIYR(iyrr string) (valid bool) {
	valid = true
	iyr, _ := strconv.Atoi(iyrr)
	if iyr != 0 {
		if !(iyr < 2021 && iyr > 2009) {
			valid = false
		}
	} else {
		valid = false
	}
	return
}
func validateEYR(eyrr string) (valid bool) {
	valid = true
	eyr, _ := strconv.Atoi(eyrr)
	if eyr != 0 {
		if !(eyr < 2031 && eyr > 2019) {
			valid = false
		}
	} else {
		valid = false
	}
	return
}

//OK
func validateHCL(hcl string) (valid bool) {
	valid = true
	if hcl != "" {
		if !(hcl[0:1] == "#" && len(hcl) == 7) {
			valid = false
		}
	} else {
		valid = false
	}
	return
}
func validateHGT(hgt string) (valid bool) {
	valid = true
	if hgt != "" {
		unit := hgt[len(hgt)-2:]
		val, _ := strconv.Atoi(hgt[0 : len(hgt)-2])

		if val != 0 {
			if unit == "cm" {
				if !(val < 194 && val > 149) {
					valid = false
				}
			} else if unit == "in" {
				if !(val < 77 && val > 58) {
					valid = false

				}
			} else {
				valid = false
			}
		} else {
			valid = false
		}
	} else {
		valid = false
	}

	return
}
func validateECL(ecl string) (valid bool) {
	valid = true
	if ecl != "" {
		if !(ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth") {
			valid = false
		}
	} else {
		valid = false
	}

	return
}
func validatePID(pidd string) (valid bool) {
	valid = true
	pid, _ := strconv.Atoi(pidd)
	if pid != 0 {
		if len(pidd) != 9 {
			valid = false
		}
	} else {
		valid = false
	}
	fmt.Println("Valid ? ", valid, "\t", pid)

	return
}
