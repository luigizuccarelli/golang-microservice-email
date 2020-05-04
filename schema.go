package main

import (
	"github.com/globalsign/mgo/bson"
)

// ShcemaInterface - acts as an interface wrapper for our profile schema
// All the go microservices will using this schema
type SchemaInterface struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	LastUpdate int64         `json:"lastupdate,omitempty"`
	MetaInfo   string        `json:"metainfo,omitempty"`
	Count      int64         `json:"count,omitempty"`
	TotalPages int64         `json:"totalpages,omitempty"`
}

// Response schema
type Response struct {
	Code       int    `json:"code,omitempty"`
	StatusCode string `json:"statuscode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Payload    string `json:"payload"`
}
