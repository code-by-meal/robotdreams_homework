package main

import (
	"fmt"
	documentstore "rb_homework/lesson_03/document_store"
)

func main() {
	// --- СЦЕНАРІЙ 1: Put new document  ---
	doc1 := &documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key":   {Type: documentstore.DocumentFieldTypeString, Value: "Document 1"},
			"name":  {Type: documentstore.DocumentFieldTypeString, Value: "Robot"},
			"age":   {Type: documentstore.DocumentFieldTypeNumber, Value: 30},
			"admin": {Type: documentstore.DocumentFieldTypeBool, Value: true},
		},
	}

	documentstore.Put(doc1)

	fmt.Println("Puted `Document 1` to store")

	// --- СЦЕНАРІЙ 2: Get document from store ---
	if doc, ok := documentstore.Get("Document 1"); ok {
		fmt.Println("Document 1:", doc.Fields)
	} else {
		fmt.Println("`Document 1` not found")
	}

	// --- СЦЕНАРІЙ 3: Get unexists document from store ---
	if _, ok := documentstore.Get("Document 2"); !ok {
		fmt.Println("`Document 2` not found (expected)")
	}

	// --- СЦЕНАРІЙ 4: Put second new document ---
	doc2 := &documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key":  {Type: documentstore.DocumentFieldTypeString, Value: "Document 2"},
			"name": {Type: documentstore.DocumentFieldTypeString, Value: "Dreams"},
		},
	}
	documentstore.Put(doc2)

	fmt.Println("Put `Document 2` to store")

	// --- СЦЕНАРІЙ 5: List all documents ---
	fmt.Println("List of all documents:", documentstore.List())

	for _, d := range documentstore.List() {
		fmt.Println(d.Fields)
	}

	// --- СЦЕНАРІЙ 6: Видалення документа ---
	if documentstore.Delete("Document 1") {
		fmt.Println("`Document 1` deleted")
	} else {
		fmt.Println("`Document 1` not found to delete")
	}

	// --- СЦЕНАРІЙ 7: Перевірка що user1 видалено ---
	if _, ok := documentstore.Get("Document 1"); !ok {
		fmt.Println("`Document 1` no longer exists (expected)")
	}

	// --- СЦЕНАРІЙ 8: Остаточний список документів ---
	fmt.Println("Final documents:")
	for _, d := range documentstore.List() {
		fmt.Println(d.Fields)
	}
}
