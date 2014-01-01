// The little-book command is a translation of the examples in The Little MongoDB Book (http://openmymind.net/2011/3/28/The-Little-MongoDB-Book) to Go-Mongo.
package common_layer

import (
	"github.com/garyburd/go-mongo/mongo"
	"log"
	"os"
	"time"
)

var chapter1SampleData = mongo.A{
	mongo.M{"name": "Horny", "dob": dateTime(1992, 2, 13, 7, 47), "loves": mongo.A{"carrot", "papaya"}, "weight": 600, "gender": "m", "vampires": 63},
	mongo.M{"name": "Aurora", "dob": dateTime(1991, 0, 24, 13, 0), "loves": mongo.A{"carrot", "grape"}, "weight": 450, "gender": "f", "vampires": 43},
	mongo.M{"name": "Unicrom", "dob": dateTime(1973, 1, 9, 22, 10), "loves": mongo.A{"energon", "redbull"}, "weight": 984, "gender": "m", "vampires": 182},
	mongo.M{"name": "Roooooodles", "dob": dateTime(1979, 7, 18, 18, 44), "loves": mongo.A{"apple"}, "weight": 575, "gender": "m", "vampires": 99},
	mongo.M{"name": "Solnara", "dob": dateTime(1985, 6, 4, 2, 1), "loves": mongo.A{"apple", "carrot", "chocolate"}, "weight": 550, "gender": "f", "vampires": 80},
	mongo.M{"name": "Ayna", "dob": dateTime(1998, 2, 7, 8, 30), "loves": mongo.A{"strawberry", "lemon"}, "weight": 733, "gender": "f", "vampires": 40},
	mongo.M{"name": "Kenny", "dob": dateTime(1997, 6, 1, 10, 42), "loves": mongo.A{"grape", "lemon"}, "weight": 690, "gender": "m", "vampires": 39},
	mongo.M{"name": "Raleigh", "dob": dateTime(2005, 4, 3, 0, 57), "loves": mongo.A{"apple", "sugar"}, "weight": 421, "gender": "m", "vampires": 2},
	mongo.M{"name": "Leia", "dob": dateTime(2001, 9, 8, 14, 53), "loves": mongo.A{"apple", "watermelon"}, "weight": 601, "gender": "f", "vampires": 33},
	mongo.M{"name": "Pilot", "dob": dateTime(1997, 2, 1, 5, 3), "loves": mongo.A{"apple", "watermelon"}, "weight": 650, "gender": "m", "vampires": 54},
	mongo.M{"name": "Nimue", "dob": dateTime(1999, 11, 20, 16, 15), "loves": mongo.A{"grape", "carrot"}, "weight": 540, "gender": "f"},
	mongo.M{"name": "Dunx", "dob": dateTime(1976, 6, 18, 18, 18), "loves": mongo.A{"grape", "watermelon"}, "weight": 704, "gender": "m", "vampires": 165},
}

func chapter1(conn mongo.Conn) {

	log.Println("\n== CHAPTER 1 ==")

	// Create a database object.
	db := mongo.Database{conn, "learn", mongo.DefaultLastErrorCmd}

	// Create a collection object object for the "unicorns" collection.
	unicorns := db.C("unicorns")

	log.Print("\n== Add first unicorn. ==\n\n")

	err := unicorns.Insert(mongo.M{"name": "Aurora", "gender": "f", "weight": 450})
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Find all unicorns. ==\n\n")

	cursor, err := unicorns.Find(nil).Cursor()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Iterate over the documents in the result. ==\n\n")

	for cursor.HasNext() {
		var m mongo.M
		err := cursor.Next(&m)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Print("\n== Show index created on _id. ==\n\n")

	si := db.C("system.indexes")
	cursor, err = si.Find(nil).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	expectFieldValues(cursor, "name", "_id_")

	log.Print("\n== Insert a totally different document in unicorns. ==\n\n")

	err = unicorns.Insert(mongo.M{"name": "Leto", "gender": "m", "home": "Arrakeen", "worm": false})
	if err != nil {
		log.Fatal(err)
	}

	cursor, err = unicorns.Find(nil).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	expectFieldValues(cursor, "name", "Aurora", "Leto")

	log.Print("\n== Remove what we added to the database so far. ==\n\n")

	err = unicorns.Remove(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Add sample data for selector examples. ==\n\n")

	err = unicorns.Insert(chapter1SampleData...)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Find all male unicorns that weight 700 pounds. ==\n\n")

	cursor, err = unicorns.Find(mongo.M{"gender": "m", "weight": mongo.M{"$gt": 700}}).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	expectFieldValues(cursor, "name", "Unicrom", "Dunx")

	log.Print("\n== A similar query for demonstration purposes. ==\n\n")

	cursor, err = unicorns.Find(mongo.M{"gender": mongo.M{"$ne": "f"}, "weight": mongo.M{"$gte": 701}}).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	expectFieldValues(cursor, "name", "Unicrom", "Dunx")

	log.Print("\n== Find unicorns without the vampires field. ==\n\n")

	cursor, err = unicorns.Find(mongo.M{"vampires": mongo.M{"$exists": false}}).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	expectFieldValues(cursor, "name", "Nimue")

	log.Print("\n== Find female unicorns which either love apples or oranges or weigh less than 500 lbs. ==\n\n")

	cursor, err = unicorns.Find(mongo.M{
		"gender": "f",
		"$or": mongo.A{
			mongo.M{"loves": "apple"},
			mongo.M{"loves": "orange"}}}).Cursor()
	if err != nil {
		log.Fatal(err)
	}

	expectFieldValues(cursor, "name", "Solnara", "Leia")
}

func chapter2(conn mongo.Conn) {

	log.Println("\n== CHAPTER 2 ==")

	db := mongo.Database{conn, "learn", mongo.DefaultLastErrorCmd}
	unicorns := db.C("unicorns")
	hits := db.C("hits")

	log.Print("\n== Change Roooooodles' weight. ==\n\n")

	err := unicorns.Update(mongo.M{"name": "Roooooodles"}, mongo.M{"weight": 590})
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Update replaced the document. ==\n\n")

	var m mongo.M
	err = unicorns.Find(mongo.M{"name": "Roooooodles"}).One(&m)
	if err != nil && err != mongo.Done {
		log.Fatal(err)
	}

	log.Print("\n== Reset the lost fields using the set operator. ==\n\n")

	err = unicorns.Update(mongo.M{"weight": 590}, mongo.M{"$set": mongo.M{
		"name":     "Roooooodles",
		"dob":      dateTime(1979, 7, 18, 18, 44),
		"loves":    mongo.A{"apple"},
		"gender":   "m",
		"vampires": 99}})
	if err != nil {
		log.Fatal(err)
	}

	m = nil
	err = unicorns.Find(mongo.M{"name": "Roooooodles"}).One(&m)
	if err != nil {
		log.Fatal(err)
	}

	if m["name"] != "Roooooodles" {
		log.Fatal("Did not find Roooooodles")
	}

	log.Print("\n== Update weight the correct way. ==\n\n")

	err = unicorns.Update(mongo.M{"name": "Roooooodles"}, mongo.M{"$set": mongo.M{"weight": 590}})
	if err != nil {
		log.Fatal(err)
	}

	m = nil
	err = unicorns.Find(mongo.M{"name": "Roooooodles"}).One(&m)
	if err != nil {
		log.Fatal(err)
	}

	if m["weight"] != 590 {
		log.Fatalf("Expected Roooooodles' weight=590, got %d\n", m["weight"])
	}

	log.Print("\n== Correct the kill count for Pilot. ==\n\n")

	err = unicorns.Update(mongo.M{"name": "Pilot"}, mongo.M{"$inc": mongo.M{"vampires": -2}})
	if err != nil {
		log.Fatal(err)
	}

	m = nil
	err = unicorns.Find(mongo.M{"name": "Pilot"}).One(&m)
	if err != nil {
		log.Fatal(err)
	}

	if m["vampires"] != 52 {
		log.Fatalf("Expected Pilot's vampires=52, got %d\n", m["vampires"])
	}

	log.Print("\n== Aurora loves sugar. ==\n\n")

	err = unicorns.Update(mongo.M{"name": "Aurora"}, mongo.M{"$push": mongo.M{"loves": "sugar"}})
	if err != nil {
		log.Fatal(err)
	}

	m = nil
	err = unicorns.Find(mongo.M{"name": "Aurora"}).One(&m)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Updating a missing document does nothing. ==\n\n")

	err = hits.Update(mongo.M{"page": "unicorns"}, mongo.M{"$inc": mongo.M{"hits": 1}})
	if err == nil || err != mongo.ErrNotFound {
		log.Fatal(err)
	}

	err = hits.Find(mongo.M{"page": "unicorns"}).One(&m)
	if err != nil && err != mongo.Done {
		log.Fatal(err)
	}

	log.Print("\n== Upsert inserts the document if missing. ==\n\n")

	err = hits.Upsert(mongo.M{"page": "unicorns"}, mongo.M{"$inc": mongo.M{"hits": 1}})
	if err != nil {
		log.Fatal(err)
	}

	err = hits.Find(mongo.M{"page": "unicorns"}).One(&m)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Upsert updates the document if already present. ==\n\n")

	err = hits.Upsert(mongo.M{"page": "unicorns"}, mongo.M{"$inc": mongo.M{"hits": 1}})
	if err != nil {
		log.Fatal(err)
	}

	err = hits.Find(mongo.M{"page": "unicorns"}).One(&m)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Update updates a single document. ==\n\n")

	err = unicorns.Update(nil, mongo.M{"$set": mongo.M{"vaccinated": true}})
	if err != nil {
		log.Fatal(err)
	}

	cursor, err := unicorns.Find(mongo.M{"vaccinated": true}).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	expectCount(cursor, 1)

	log.Print("\n== UpdateAll updates all documents. ==\n\n")

	err = unicorns.UpdateAll(nil, mongo.M{"$set": mongo.M{"vaccinated": true}})
	if err != nil {
		log.Fatal(err)
	}

	cursor, err = unicorns.Find(mongo.M{"vaccinated": true}).Cursor()
	if err != nil {
		log.Fatal(err)
	}

	expectCount(cursor, 12)
}

func chapter3(conn mongo.Conn) {

	log.Println("\n== CHAPTER 3 ==")

	db := mongo.Database{conn, "learn", mongo.DefaultLastErrorCmd}
	unicorns := db.C("unicorns")

	log.Print("\n== Find names of all unicorns. ==\n\n")

	cursor, err := unicorns.Find(nil).Fields(mongo.M{"name": 1}).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	expectCount(cursor, 12)

	log.Print("\n== Find all unicorns ordered by decreasing weight. ==\n\n")

	cursor, err = unicorns.Find(nil).Sort(mongo.D{{"weight", -1}}).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	expectCount(cursor, 12)

	log.Print("\n== Find all unicorns ordered by name and then vampire kills. ==\n\n")

	cursor, err = unicorns.Find(nil).Sort(mongo.D{{"name", 1}, {"vampires", -1}}).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	expectCount(cursor, 12)

	log.Print("\n== Find the 2nd and 3rd heaviest unicorns. ==\n\n")

	cursor, err = unicorns.Find(nil).Sort(mongo.D{{"weight", -1}}).Limit(2).Skip(1).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	expectCount(cursor, 2)

	log.Print("\n== Count unicorns with more than 50 kills. ==\n\n")

	n, err := unicorns.Find(mongo.M{"vampires": mongo.M{"$gt": 50}}).Count()
	if err != nil {
		log.Fatal(err)
	}

	if n != 6 {
		log.Fatalf("Got count=%d, want 6", n)
	}
}

func chapter7(conn mongo.Conn) {

	log.Println("\n== CHAPTER 7 ==")

	db := mongo.Database{conn, "learn", mongo.DefaultLastErrorCmd}
	unicorns := db.C("unicorns")

	log.Print("\n== Create index on name. ==\n\n")

	err := unicorns.CreateIndex(mongo.D{{"name", 1}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Drop index on name. ==\n\n")

	err = db.Run(mongo.D{
		{"dropIndexes", unicorns.Name()},
		{"index", mongo.IndexName(mongo.D{{"name", 1}})},
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Create unique index on name. ==\n\n")

	err = unicorns.CreateIndex(mongo.D{{"name", 1}}, &mongo.IndexOptions{Unique: true})
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Create compound index on name ascending and kills descending. ==\n\n")

	err = unicorns.CreateIndex(mongo.D{{"name", 1}, {"vampires", -1}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Explain query. ==\n\n")

	var m mongo.M
	err = unicorns.Find(nil).Explain(&m)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Explain query on name. ==\n\n")

	m = nil
	err = unicorns.Find(mongo.M{"name": "Pilot"}).Explain(&m)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Enable profiling. ==\n\n")

	err = db.Run(mongo.D{{"profile", 2}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n== Get profile data for query. ==\n\n")

	cursor, err := unicorns.Find(mongo.M{"weight": mongo.M{"$gt": 600}}).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	expectFieldValues(cursor, "name", "Unicrom", "Ayna", "Kenny", "Leia", "Pilot", "Dunx")

	cursor, err = db.C("system.profile").Find(nil).Cursor()
	if err != nil {
		log.Fatal(err)
	}

	for cursor.HasNext() {
		var m mongo.M
		err := cursor.Next(&m)
		if err != nil {
			log.Fatal(err)
		}
	}
	cursor.Close()

	log.Print("\n== Profile queries that take longer than 100 ms. ==\n\n")

	err = db.Run(mongo.D{{"profile", 2}, {"slowms", 100}}, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// expectCount iterates through the cursor results and logs a fatal error if
// the number of documents is not equal to n.
func expectCount(cursor mongo.Cursor, n int) {
	defer cursor.Close()
	i := 0
	for cursor.HasNext() {
		i += 1
		var m mongo.M
		err := cursor.Next(&m)
		if err != nil {
			log.Fatal(err)
		}
	}
	if i != n {
		log.Fatalf("Got result count=%d, want %d", i, n)
	}
}

// expectFieldValues iterates through the cursor and logs a fatal error if
// the a document does not have field in values or if a value in values was not
// found in a document.
func expectFieldValues(cursor mongo.Cursor, field string, values ...interface{}) {
	defer cursor.Close()
	found := map[interface{}]bool{}
	for cursor.HasNext() {
		var m mongo.M
		err := cursor.Next(&m)
		if err != nil {
			log.Fatal(err)
		}
		found[m["name"]] = true
	}
	for _, value := range values {
		if !found[value] {
			log.Fatalf("Expected result %v not found\n", value)
		} else {
			delete(found, value)
		}
	}
	for value, _ := range found {
		log.Fatalf("Unexpected result %v found\n", value)
	}
}

// dateTime converts year, month, day hour and seconds to a time.Time
func dateTime(year, month, day, hour, minutes int) time.Time {
	return time.Date(year, time.Month(month), day, hour, minutes, 0, 0, time.Local)
}

// reset cleans up after previous runs of this applications.
func reset(conn mongo.Conn) {
	log.Print("\n== Clear documents and indexes created by previous run. ==\n\n")
	db := mongo.Database{conn, "learn", mongo.DefaultLastErrorCmd}
	db.Run(mongo.D{{"profile", 0}}, nil)
	db.C("unicorns").Remove(nil)
	db.C("hits").Remove(nil)
	db.Run(mongo.D{{"dropIndexes", "unicorns"}, {"index", "*"}}, nil)
	db.Run(mongo.D{{"dropIndexes", "hits"}, {"index", "*"}}, nil)
}

func Mongo_main() {

	// Connect to server.
	conn, err := mongo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Wrap connection with logger so that we can view the traffic to and from
	// the server.
	conn = mongo.NewLoggingConn(conn, log.New(os.Stdout, "", 0), "")

	// Clear the log prefix for more readable output.
	log.SetFlags(0)

	reset(conn)
	chapter1(conn)
	chapter2(conn)
	chapter3(conn)
	chapter7(conn)
}
