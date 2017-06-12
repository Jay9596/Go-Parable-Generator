package main

import (
	"fmt"
	"math/rand"
	"time"

	"strings"

	"github.com/spf13/viper"
)

type config struct {
	starter           []string
	cannot            []string
	cannotPresent     []string
	cannotPlural      []string
	can               []string
	canPresent        []string
	action            []string
	presentParticiple []string
	foodAdjective     []string
	food              []string
	thing             []string
	adjective         []string
}

var cf config

func setupViper() {
	viper.SetConfigType("json")
	//viper.AddConfigPath("src/github.com/jay9596/Go_Parable_Generator")
	viper.AddConfigPath(".")
}
func main() {
	setupViper()
	readFile()
	num := getNumber()
	generatePhrases(num)
}

func readFile() {
	viper.SetConfigName("reassuring")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	} else {
		cf.starter = viper.GetStringSlice("starter")
		cf.cannot = viper.GetStringSlice("cannot")
		cf.cannotPresent = viper.GetStringSlice("cannot_present")
		cf.cannotPlural = viper.GetStringSlice("cannot_plural")
		cf.can = viper.GetStringSlice("can")
		cf.canPresent = viper.GetStringSlice("can_present")
		cf.action = viper.GetStringSlice("action")
		cf.presentParticiple = viper.GetStringSlice("present_participle")
		cf.foodAdjective = viper.GetStringSlice("food_adjective")
		cf.food = viper.GetStringSlice("food")
		cf.thing = viper.GetStringSlice("thing")
		cf.adjective = viper.GetStringSlice("adjective")
	}
}

func getNumber() int {
	var num int
	fmt.Print("Enter number of string to generate: ")
	fmt.Scanf("%d", &num)
	return num
}

func generatePhrases(count int) {
	curr := time.Now()
	for i := 0; i < count; i++ {
		fmt.Printf("%d. %s\n", i+1, genPhrase())
	}
	timeDiff := time.Now().Sub(curr)
	fmt.Println("Generated in: ", timeDiff)
}

func genPhrase() string {
	phrase := getRandom(cf.starter)
	return evaluate(phrase)
}

func evaluate(phrase string) string {
	lt := strings.IndexRune(phrase, '{')
	if lt == -1 {
		return phrase
	}
	rt := strings.IndexRune(phrase, '}')
	key := phrase[lt+1 : rt]
	val := getValue(key)
	phrase = newStr(phrase, key, val)
	phrase = evaluate(phrase)
	return phrase
}

func newStr(ph string, key string, val string) string {
	return strings.Replace(ph, "{"+key+"}", val, -1)
}

func getValue(key string) string {
	switch key {
	case "cannot":
		return getRandom(cf.cannot)
	case "cannot_present":
		return getRandom(cf.cannotPresent)
	case "cannot_plural":
		return getRandom(cf.cannotPlural)
	case "can":
		return getRandom(cf.can)
	case "can_present":
		return getRandom(cf.canPresent)
	case "action":
		return getRandom(cf.action)
	case "present_participle":
		return getRandom(cf.presentParticiple)
	case "food_adjective":
		return getRandom(cf.foodAdjective)
	case "food":
		return getRandom(cf.food)
	case "thing":
		return getRandom(cf.thing)
	case "adjective":
		return getRandom(cf.adjective)
	}
	return ""
}

func getRandom(arr []string) string {
	l := len(arr)
	if l > 0 {
		return arr[rand.Intn(l)]
	}
	return " "
}
