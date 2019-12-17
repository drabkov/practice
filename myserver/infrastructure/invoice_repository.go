package infrastructure

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"myserver/domain"
)

//Repository
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "mystore"

// DOCNAME the name of the document
const DOCNAME = "invoices"

func (r Repository) GetInvoices() domain.Invoices {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := domain.Invoices{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

func (r Repository) AddInvoice(invoice domain.Invoice) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	invoice.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(invoice)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (r Repository) UpdateInvoice(invoice domain.Invoice) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	//invoice.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).UpdateId(invoice.ID, invoice)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (r Repository) DeleteInvoice(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Verify id is ObjectId, otherwise fail
	if !bson.IsObjectIdHex(id) {
		return "404"
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "500"
	}

	// Write status
	return "200"
}
