//Package tokenization adds support for tokenizing emv / PCI cards in a secure fashion.

package tokenization

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // sqlite3 specific entry
	"github.com/segmentio/ksuid"
)

var stmt = `

CREATE TABLE IF NOT EXISTS Cards (
	token text primary key,
	pan text unique not null,
	pin text not null,
	expdate text not null,
	fingerprint text unique not null,
	biller_id integer not null
)
`
var db *sqlx.DB

func init() {
	var err error
	db, err = opendDb("tokenization.db")
	if err != nil {
		log.Fatalf("Error in connecting to DB: %v", err)
	}
	if _, err := db.Exec(stmt); err != nil {
		log.Printf("Error in exec-ing sql stmt: %v", err)
	}
}

//Card is payment card to be tokenized
type Card struct {
	Pan         string   `json:"pan,omitempty" db:"pan"`
	Pin         string   `json:"pin,omitempty" db:"pin"`
	Expdate     string   `json:"expdate,omitempty" db:"expdate"`
	Token       string   `json:"token,omitempty" db:"token"`
	Fingerprint string   `json:"fingerprint,omitempty" db:"fingerprint"`
	Biller      string   `json:"biller,omitempty" db:"biller_id"`
	LastPan     string   `json:"last_pan,omitempty"`
	db          *sqlx.DB `json:"db,omitempty"`
}

//NewCard creates a new card to be used by this package consumers
func NewCard() (*Card, error) {
	var c Card
	c.db = db
	if err := c.db.Ping(); err != nil {
		return nil, err
	}
	return &c, nil
}

//NewToken generate a new ksuid compatible token
func (c *Card) NewToken() error {
	id, err := ksuid.NewRandom()
	if err != nil {
		log.Printf("Error in ksuid: %v", err)
		return err
	}
	c.Token = id.String()
	if err := c.write(); err != nil {
		log.Printf("Error in db: %v", err)
		return err
	}
	return nil
}

func (c *Card) GetTokenized() Card {
	c.LastPan = c.Pan[len(c.Pan)-4:]
	c.Pan = ""
	return *c
}

func (c *Card) write() error {
	if _, err := c.db.NamedExec("INSERT INTO CARDS(token, pan, pin, expdate, fingerprint, biller_id) VALUES(:token, :pan, :pin, :expdate, :fingerprint, :biller_id)", c); err != nil {
		log.Printf("Error in c.write() db: %v", err)
		return err
	}
	return nil
}

func (c *Card) read(token string) error {
	if err := c.db.Get(c, "SELECT * FROM CARDS WHERE TOKEN = ?", token); err != nil {
		return err
	}
	return nil
}

//NewFromToken generates a new Card item from a token
func (c Card) NewFromToken(token string) (*Card, error) {
	if err := c.read(token); err != nil {
		return nil, err
	}
	return &c, nil
}
