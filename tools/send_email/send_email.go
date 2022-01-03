package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func main() {
	var (
		tf  string
		err error
		mb  *Mailbox
	)

	if len(os.Args) == 1 {
		log.Fatalln("please provide emial config files")
	}

	for _, tf = range os.Args[1:] {
		if mb, err = sendEmail(tf); err != nil {
			log.Fatalf("send %q: %v\n", tf, err)
		} else {
			log.Printf("successed send %q: %q\n", tf, mb.Title)
		}
	}
}

func sendEmail(tf string) (mb *Mailbox, err error) {
	var conf *viper.Viper

	mb = new(Mailbox)
	conf = viper.New()
	conf.SetConfigName("project config")
	switch {
	case strings.HasSuffix(tf, ".toml"):
		conf.SetConfigType("toml")
	case strings.HasSuffix(tf, ".yaml") || strings.HasSuffix(tf, ".yml"):
		conf.SetConfigType("yaml")
	default:
		return nil, fmt.Errorf("support toml(.toml) or yaml(.yaml, .yml) files only")
	}

	conf.SetConfigFile(tf)
	if err = conf.ReadInConfig(); err != nil {
		return nil, err
	}

	if err = conf.UnmarshalKey("send", mb); err != nil {
		return nil, err
	}

	err = mb.Send()
	return mb, err
}

type Mailbox struct {
	SMTPAddr string `mapstructure:"smtpAddr"`
	SMTPPort int    `mapstructure:"smtpPort"`
	Sender   string `mapstructure:"sender"`
	Password string `mapstructure:"password"`

	Recipients []string `mapstructure:"recipients"`
	Title      string   `mapstructure:"title"`
	Body       string   `mapstructure:"body"`
	Attachs    []string `mapstructure:"attachs"`
}

func (mb *Mailbox) Send() (err error) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", mb.Sender)
	msg.SetHeader("To", mb.Recipients...)
	msg.SetHeader("Subject", mb.Title)
	msg.SetBody("text/html", mb.Body)

	for i := range mb.Attachs {
		msg.Attach(mb.Attachs[i])
	}

	return gomail.NewDialer(mb.SMTPAddr, mb.SMTPPort, mb.Sender, mb.Password).DialAndSend(msg)
}
