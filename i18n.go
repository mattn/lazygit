package main

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// the function to setup the localizer
func getlocalizer() *i18n.Localizer {

	// TODO: currently the system language issn't detected
	// I'm not sure how to detect it
	var i18nObject = &i18n.Bundle{DefaultLanguage: language.Dutch}
	i18nObject.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// Dutch translation for some words
	i18nObject.AddMessages(language.Dutch,
		&i18n.Message{
			ID:    "FilesTitle",
			Other: "Bestanden",
		}, &i18n.Message{
			ID:    "BranchesTitle",
			Other: "Branches",
		}, &i18n.Message{
			ID:    "CommitsTitle",
			Other: "Commits",
		}, &i18n.Message{
			ID:    "StashTitle",
			Other: "Stash",
		}, &i18n.Message{
			ID:    "CommitMessage",
			Other: "Commit Bericht",
		}, &i18n.Message{
			ID:    "StatusTitle",
			Other: "Status",
		},
	)

	return i18n.NewLocalizer(i18nObject)
}

// setup the localizer for later use
var localizer = getlocalizer()

// MustLocalize handels the translations
// expects i18n.LocalizeConfig as input: https://godoc.org/github.com/nicksnyder/go-i18n/v2/i18n#Localizer.MustLocalize
// output: translated string
func MustLocalize(config *i18n.LocalizeConfig) string {
	return localizer.MustLocalize(config)
}

// ShortLocalize is for 1 line localizations
// ID: The id that is used in the .toml translation files
// Other: the default message it needs to return if there is no translation found or the system is english
func ShortLocalize(ID string, Other string) string {
	return MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    ID,
			Other: Other,
		},
	})
}