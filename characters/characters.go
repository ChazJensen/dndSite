package characters

import (
	"errors"
	"html/template"
	"io"
	"server/csv"
	"os"
	"strconv"
	"strings"
)

// logging
func log(s string) {
	println("chrsht: " + s)
}

// error checking
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetChar(usrName string, charName string) Char {
	log("retrieving " + charName)

	// Get "databases"
	usrFile, err := os.Open("./data/users.csv")
	check(err)
	charFile, err := os.Open("./data/characters.csv")
	check(err)

	// get user's character's IDs
	users, err := csv.Lookup(usrFile, "usr", usrName)
	check(err)

	usrRecords := users[0]

	var charsOwned []string = strings.Split(usrRecords[2], "|")
	var charNames []string = strings.Split(usrRecords[1], "-")
	var charID string

	if len(charsOwned) != len(charNames) {
		panic(errors.New("there was an internal error originiating from your user"))
	}

	for i := 0; i < len(charsOwned); i++ {
		if charsOwned[i] == charName {
			charID = charNames[i]
		}
	}

	if charID == "" {
		panic(errors.New("you do not own a character with this name"))
	}

	// find charName in "database"
	charsFound, err := csv.Lookup(charFile, "uniq-id", charID)
	check(err)

	if len(charsFound) > 1 {
		panic(errors.New("there was a problem with the character lookup!"))
	}

	charData := charsFound[0]

	// parse data from "database"
	name, abilities, desc := charData[1], ParseAbilities(charData[2]), charData[3]

	// get character sheet template file and load it into tmpl
	templ, err := template.ParseFiles("./templates/character_sheet.tmp")
	check(err)


	c := Char{name, abilities, desc, *templ}

	return c
}

func ParseAbilities(abilityString string) []int {
	var a []int

	abilities := strings.Split(abilityString, "-")

	for _, ability := range abilities {
		i, err := strconv.Atoi(ability)
		check(err)

		a = append(a, i)
	}

	return a
}



type Char struct {
	Name	string
	Stats	[]int
	Desc	string
	Sheet	template.Template // call Execute(Writer, data)
}

func (c *Char) ExecuteTemplate(w io.Writer) {
	c.Sheet.Execute(w, c)
}
