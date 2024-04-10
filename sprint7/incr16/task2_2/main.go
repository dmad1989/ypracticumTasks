package main

// JSONData — интерфейс для декодирования JSON.
type JSONData interface {
	DecodeJSON() interface{}
}

// YAMLData — интерфейс для декодирования YAML.
type YAMLData interface {
	DecodeYAML() interface{}
}

type Client struct {
	Data interface{}
}

func (client *Client) Decode(input JSONData) {
	client.Data = input.DecodeJSON()
}

// добавьте тип Adapter и необходимый метод
// ...

type Adapter struct {
	yamldata YAMLData
}

func (a *Adapter) DecodeJSON() interface{} {
	return a.yamldata.DecodeYAML()
}

func Load(client Client, input YAMLData) {
	adapter := &Adapter{
		yamldata: input,
	}
	client.Decode(adapter)
}
