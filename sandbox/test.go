package main

import (
	"fmt"
	"time"
)

func main(){
	//layout := "2006-01-02 15:04:05"
	//str := "2018-03-01 10:15:53"
	//t, err := time.Parse(layout, str)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(reflect.TypeOf(t))

	//type M map[string]interface{}
	////M{"j": M{"$ne": 3}, "k": M{"$gt": 10}}
	//m := M{"j": M{"$ne": 3}, "k": M{"$gt": 10}}
	//fmt.Println(m)
	//

	//fromDate := time.Date(2014, time.November, 4, 0, 0, 0, 0, time.UTC)
	//toDate := time.Date(2014, time.November, 5, 0, 0, 0, 0, time.UTC)
	//
	//var sales_his []Sale
	//err = c.Find(
	//	bson.M{
	//		"sale_date": bson.M{
	//			"$gt": fromDate,
	//			"$lt": toDate,
	//		},
	//	}).All(&sales_his)
	now := time.Now()
	fmt.Println(ISODate("2010-04-29T00:00:00.000Z")	now)
}


