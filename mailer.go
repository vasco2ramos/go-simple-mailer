package mailer

import (
      //"encoding/json"
	    //"io/ioutil"
	    "net/http"
      "gopkg.in/gomail.v2"
      "github.com/go-martini/martini"
)

type client struct {
        Id      int
        Email   string
        Name    string
}

type candidate struct{
        Id      int
        Email   string
        Name    string
}

type email struct{
      Candidate     candidate
      Client        client
      Report        string
}



func NewEmail(cl client, ca candidate, rep string) *email {
  return &email{
      Client: cl,
      Candidate: ca,
      Report: rep,
  }
}

// WebPost implements webservice.WebPost.
func (e *email) WebPost(params martini.Params,
  req *http.Request) (int, string) {
	// Make sure Body is closed when we are done.
	defer req.Body.Close()

	/*
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusInternalServerError, "internal error"
	}

	if len(params) != 0 {
		// No keys in params. This is not supported.
		return http.StatusMethodNotAllowed, "method not allowed"
	}

	// Get Post Request here
	var guestBookEntry GuestBookEntry
	err = json.Unmarshal(requestBody, &guestBookEntry)
	if err != nil {
		// Could not unmarshal entry.
		return http.StatusBadRequest, "invalid JSON data"
	}
  */

  // SEND EMAIL HERE
  //mailReport(g.Client, g.Report)
  //mailCandidate(g.Client, g.Candidate)

	// Everything is fine.
	return http.StatusOK, "Email Sent"
}

func mailReport(){

}

func mailCandidate(){

}



/*
func main() {
  m := gomail.NewMessage()
  m.SetHeader("From", "alex@example.com")
  m.SetHeader("To", "vasco@tyba.com")
  m.SetHeader("Subject", "Hello!")
  m.SetBody("text/html", "Hello <b>Me</b>!")
  d := gomail.NewPlainDialer("smtp.gmail.com", 587, "vascoasramos@gmail.com", "pai121natal")

  // Send the email to Bob, Cora and Dan.
  if err := d.DialAndSend(m); err != nil {
      panic(err)
  }

}
*/
