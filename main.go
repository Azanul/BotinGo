package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

type Gopher struct {
	Name string `json:"name"`
}

var Hyes = []string{"hi", "hello", "hallo", "aloha", "hey", "bonjour", "konnichiwa", "ciao",
	"how you doin'?", "whatcha cookin' good lookin'"}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	for _, b := range Hyes {
		if strings.HasPrefix(strings.ToLower(m.Content), b) {
			fmt.Println(strings.Index(strings.ToLower(m.Content), b))
			rand_greet := Hyes[rand.Intn(len(Hyes))]
			caser := cases.Title(language.AmericanEnglish)
			_, err := s.ChannelMessageSend(m.ChannelID, caser.String(rand_greet)+" "+m.Author.Mention())
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	if m.Content == "!gopher" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Baap ko yaad kiya beta ji ne.")
		if err != nil {
			fmt.Println(err)
		}
	}

	if strings.ToLower(m.Content) == "suna" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Aati kya khandala.?")
		if err != nil {
			fmt.Println(err)
		}
	}

	if m.Content == "$inspire" {
		_, err := s.ChannelMessageSend(m.ChannelID, "This link might be of help "+m.Author.Mention()+"https://en.wikipedia.org/wiki/Suicide_methods#List")
		if err != nil {
			fmt.Println(err)
		}
	}

	// if m.Content == "!random" {

	// 	//Call the KuteGo API and retrieve a random Gopher
	// 	response, err := http.Get(KuteGoAPIURL + "/gopher/random/")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	defer response.Body.Close()

	// 	if response.StatusCode == 200 {
	// 		_, err = s.ChannelFileSend(m.ChannelID, "random-gopher.png", response.Body)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	} else {
	// 		fmt.Println("Error: Can't get random Gopher! :-(")
	// 	}
	// }

	// if m.Content == "!gophers" {

	// 	//Call the KuteGo API and display the list of available Gophers
	// 	response, err := http.Get(KuteGoAPIURL + "/gophers/")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	defer response.Body.Close()

	// 	if response.StatusCode == 200 {
	// 		// Transform our response to a []byte
	// 		body, err := ioutil.ReadAll(response.Body)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}

	// 		// Put only needed informations of the JSON document in our array of Gopher
	// 		var data []Gopher
	// 		err = json.Unmarshal(body, &data)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}

	// 		// Create a string with all of the Gopher's name and a blank line as separator
	// 		var gophers strings.Builder
	// 		for _, gopher := range data {
	// 			gophers.WriteString(gopher.Name + "\n")
	// 		}

	// 		// Send a text message with the list of Gophers
	// 		_, err = s.ChannelMessageSend(m.ChannelID, gophers.String())
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	} else {
	// 		fmt.Println("Error: Can't get list of Gophers! :-(")
	// 	}
	// }
}
