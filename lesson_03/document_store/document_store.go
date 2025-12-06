package documentstore

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value any
}

type Document struct {
	Fields map[string]DocumentField
}

var documents = map[string]*Document{}

func Put(doc *Document) {
	// 1. Перевірити що документ містить в мапі поле `key` типу `string`
	// 2. Додати Document до локальної мапи з документами
	// TODO: Implement

	// 1.
	field, ok := doc.Fields["key"]

	if !ok {
		return
	}

	value, ok := field.Value.(string)

	if !ok {
		return
	}

	// 2.
	documents[value] = doc
}

func Get(key string) (*Document, bool) {
	// Потрібно повернути документ по ключу
	// Якщо документ знайдено, повертаємо `true` та поінтер на документ
	// Інакше повертаємо `false` та `nil`
	// TODO: Implement

	doc, ok := documents[key]

	if !ok {
		return nil, false
	}

	return doc, true
}

func Delete(key string) bool {
	// Видаляємо документа по ключу.
	// Повертаємо `true` якщо ми знайшли і видалили документі
	// Повертаємо `false` якщо документ не знайдено
	// TODO: Implement

	_, ok := documents[key]

	if !ok {
		return false
	}

	delete(documents, key)

	return true
}

func List() []*Document {
	// Повертаємо список усіх документів
	// TODO: Implement

	list := []*Document{}

	for _, doc := range documents {
		list = append(list, doc)
	}

	return list
}
