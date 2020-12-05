//Package tokenization adds support for tokenizing emv / PCI cards in a secure fashion.

package tokenization

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func TestCard_NewToken(t *testing.T) {
	db, err := opendDb("fortest.db")
	if err != nil {
		t.Fatalf("Error in open db: %v", err)
	}

	if _, err := db.Exec(stmt); err != nil {
		t.Fatalf("Error in exec-ing sql stmt: %v", err)
	}

	testCard := Card{Token: "443", Pan: "23232", Pin: "2323", Expdate: "3232", Fingerprint: "3232", Biller: "3232", db: db}

	tests := []struct {
		name    string
		fields  Card
		wantErr bool
	}{
		{"testing successful", testCard, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := tt.fields.NewToken(); (err != nil) != tt.wantErr {
				t.Errorf("Card.NewToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCard_NewFromToken(t *testing.T) {

	db, err := opendDb("fortest.db")
	if err != nil {
		t.Fatalf("Error in open db: %v", err)
	}

	if _, err := db.Exec(stmt); err != nil {
		t.Fatalf("Error in exec-ing sql stmt: %v", err)
	}

	testCard := Card{Token: "443", Pan: "23232", Pin: "2323", Expdate: "3232", Fingerprint: "3232", Biller: "3232", db: db}

	type fields struct {
		Pan         string
		Pin         string
		Expdate     string
		Token       string
		Fingerprint string
		Biller      string
		LastPan     string
		db          *sqlx.DB
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  Card
		args    args
		want    *Card
		wantErr bool
	}{
		{"new from token - success", Card{db: db}, args{"1lFIFTPIn3pKH8ehpqgT5A5KJzM"}, &testCard, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.fields.NewFromToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Card.NewFromToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Pan, tt.want.Pan) {
				t.Errorf("Card.NewFromToken() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.Pin, tt.want.Pin) {
				t.Errorf("Card.NewFromToken() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.Expdate, tt.want.Expdate) {
				t.Errorf("Card.NewFromToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
