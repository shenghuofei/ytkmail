package main 

import (
    "fmt"
    "net/mail"
    "net/smtp"
    "github.com/scorredoira/email"
    "strings"
    "flag"
    "os"
)

var (
    mail_host = "smtp.exmail.qq.com"
    mail_user = "from-user@xxx.com"
    mail_pwd  = "from-user-passwd"

    tos = flag.String("to","","Mail recipients separated by a semicolon(;)")
    cc = flag.String("Cc","","Mail Cc separated by a semicolon(;)")
    subject = flag.String("subject","","Mail title")
    body = flag.String("body","","Mail body")
    attach = flag.String("attach","","Mail attachments separated by commas(,)")
    help = flag.Bool("h",false,"help message")
)

func usage() {
    fmt.Println("Usage: ytkmail -to to-addr -subject subject -body body [-Cc cc-addr] [-attach attach]")
    os.Exit(1)
}

func main() {
    //flag.Usage = usage
    flag.Parse()
    if *help {
        flag.Usage()
        os.Exit(0)
    }
    if *tos == "" || *subject == "" {
	fmt.Println("[ERROR] to-addr and subject is must")
        usage()
    }

    //m := email.NewMessage(*subject, *body)
    m := email.NewHTMLMessage(*subject, *body)
    m.To = strings.Split(*tos,";")
    if *cc != "" {
        m.Cc = strings.Split(*cc,";")
    }
    m.From = mail.Address{Name: "from-user-name", Address: mail_user}
    
    if *attach != "" {
        for _,f := range strings.Split(*attach,",") {
            // add attachments
            if err := m.Attach(f); err != nil {
                fmt.Println(err)
                os.Exit(2)
            }
        }
    }

    // send it
    auth := smtp.PlainAuth("", mail_user, mail_pwd, mail_host)
    if err := email.Send(mail_host+":25", auth, m); err != nil {
        fmt.Println(err)
        os.Exit(3)
    } else {
        fmt.Println("success")
    }
}
