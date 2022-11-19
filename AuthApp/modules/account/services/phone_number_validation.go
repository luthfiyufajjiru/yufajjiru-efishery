package services

import "regexp"

func PhoneNumberValidation(phone_number string) (res bool) {
	pattern := `\(?\+[0-9]{1,3}\)? ?-?[0-9]{1,3} ?-?[0-9]{3,5} ?-?[0-9]{4}( ?-?[0-9]{3})? ?(\w{1,10}\s?\d{1,6})?`
	res, _ = regexp.Match(pattern, []byte(phone_number))
	return
}
