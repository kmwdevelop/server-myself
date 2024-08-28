package db

import "testing"

func TestInsert(t *testing.T) {
	InitMongoDB("mongodb+srv://test:1111@cluster0.rhxlati.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	InsertOne("test")
}
