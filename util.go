package hugo

import (
	"regexp"

	"gopkg.in/devalecs/tgo.v1"
)

type Util struct{}

func (u *Util) MessageTextMatch(update *tgo.Update, pattern string) bool {
	if update.Message == nil {
		return false
	}

	match, _ := regexp.MatchString(pattern, update.Message.Text)

	return match
}

func (u *Util) InlineQueryMatch(update *tgo.Update, pattern string) bool {
	if update.InlineQuery == nil {
		return false
	}

	match, _ := regexp.MatchString(pattern, update.InlineQuery.Query)

	return match
}

func (u *Util) MessageTextMatchParams(update *tgo.Update, pattern string) (bool, []interface{}) {
	var submatch []interface{}

	if update.Message == nil {
		return false, submatch
	}

	r := regexp.MustCompile(pattern)
	if !r.MatchString(update.Message.Text) {
		return false, submatch
	}

	match := r.FindStringSubmatch(update.Message.Text)

	return true, u.stringToInterfaceSlice(match[1:])
}

func (u *Util) InlineQueryMatchParams(update *tgo.Update, pattern string) (bool, []interface{}) {
	var submatch []interface{}

	if update.InlineQuery == nil {
		return false, submatch
	}

	r := regexp.MustCompile(pattern)
	if !r.MatchString(update.InlineQuery.Query) {
		return false, submatch
	}

	match := r.FindStringSubmatch(update.InlineQuery.Query)

	return true, u.stringToInterfaceSlice(match[1:])
}

func (u *Util) stringToInterfaceSlice(stringSlices []string) []interface{} {
	var interfaceSlices []interface{}

	for _, stringSlice := range stringSlices {
		interfaceSlices = append(interfaceSlices, interface{}(stringSlice))
	}

	return interfaceSlices
}
