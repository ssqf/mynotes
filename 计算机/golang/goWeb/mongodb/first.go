package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Document 文档
type Document struct {
	ID        bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	LikeNum   uint            `json:"like" bson:"like"`
	SickNum   uint            `json:"sick" bson:"sick"`
	UserID    bson.ObjectId   `json:"userID" bson:"userID"`
	Content   string          `json:"content" bson:"Content"`
	CommentID []bson.ObjectId `json:"commentID" bson:"commentID"`
	Date      time.Time       `json:"date" bson:"date"`
}

// Comment 评论的结构
type Comment struct {
	Floor     uint
	ActicleID string
}

// Acticle 文章
type Acticle struct {
	Title     string
	Topic     string
	keyWork   []string
	topPicURL []string
	Pageview  uint
}

// User 用户信息
type User struct {
	ID            string
	Nickname      string
	Email         string
	HeadPortraits string
	Bio           string
	Like          string
	Passwd        string
}

var (
	dbName = "MyApp"
)

func main() {
	session, err := mgo.Dial("119.27.177.240:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	session.SetSocketTimeout(0)
	appDB := session.DB(dbName)
	docCollection := appDB.C("doc")

	doc := Document{}
	id := bson.NewObjectId()
	doc.ID = id
	doc.LikeNum = 0
	doc.SickNum = 0
	doc.UserID = bson.NewObjectId()
	doc.Date = bson.Now()
	doc.Content = "我就是一个测试例子258"
	doc.CommentID = []bson.ObjectId{bson.ObjectIdHex("565656565656565612521246"), bson.ObjectIdHex("565963265656565612521246")}

	err = docCollection.Insert(doc)

	if err != nil {
		log.Fatal(err)
	}
	var doc1 Document
	err = docCollection.Find(bson.M{"_id": id}).One(&doc1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("doc1:%v\n", doc1)

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"like": -1}},
		ReturnNew: true,
	}

	var doc3 Document
	change1 := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"like": 1}},
		ReturnNew: true,
	}
	for i := 0; i < 10000; i++ {
		var err error
		if rand.Int()%2 == 0 {
			_, err = docCollection.Find(bson.M{"_id": id}).Apply(change1, &doc3)
		} else {
			_, err = docCollection.Find(bson.M{"_id": id}).Apply(change, &doc3)
		}

		if err != nil {
			log.Fatal(err)
		}
	}
	// var doc2 Document
	// change := mgo.Change{
	// 	Update:    bson.M{"$inc": bson.M{"like": -1}},
	// 	ReturnNew: true,
	// }

	// increamMax, decrementMax := 2500, 1000
	// var w sync.WaitGroup
	// w.Add(increamMax + decrementMax + 2)
	// go func() {
	// 	defer w.Done()
	// 	for i := 0; i < decrementMax; i++ {
	// 		go func() {
	// 			defer w.Done()
	// 			_, err := docCollection.Find(bson.M{"_id": id}).Apply(change, &doc2)
	// 			//time.Sleep(3 * time.Millisecond)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 		}()
	// 	}
	// }()

	// var doc3 Document
	// change1 := mgo.Change{
	// 	Update:    bson.M{"$inc": bson.M{"like": 1}},
	// 	ReturnNew: true,
	// }
	// go func() {
	// 	defer w.Done()
	// 	for i := 0; i < increamMax; i++ {
	// 		go func() {
	// 			defer w.Done()
	// 			_, err := docCollection.Find(bson.M{"_id": id}).Apply(change1, &doc3)
	// 			//time.Sleep(3 * time.Millisecond)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 		}()
	// 	}

	// }()

	// w.Wait()
	// user := User{}
	// comment := Comment{}
	// appDB := session.DB(dbName)
	// UserCollection := appDB.C("User")
	// CommentCollection := appDB.C("Comment")
	// err = UserCollection.Insert(&user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = CommentCollection.Insert(&comment)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// user1 := User{}
	// err = UserCollection.Find(bson.M{"nickname": "tako"}).One(&user1)
	// if err != nil {
	// 	log.Fatal("user:", err)
	// }

	// comment1 := Comment{}
	// err = CommentCollection.Find(bson.M{}).One(&comment1)
	// if err != nil {
	// 	log.Fatal("comment:", err)
	// }

	// fmt.Printf("user:%v\ncomment:%v\n", user1, comment1)
}
