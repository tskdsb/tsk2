package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/go-openapi/spec"

	"github.com/tskdsb/tsk2/pkg/api"
)

// The location of the parameter. Possible values are "query", "header", "path", "formData" or "body".
const (
	ParameterInTypeQuery    string = "query"
	ParameterInTypeHeader   string = "header"
	ParameterInTypePath     string = "path"
	ParameterInTypeFormData string = "formData"
	ParameterInTypeBody     string = "body"
)

// The type of the object. The value MUST be one of "string", "number", "integer", "boolean", or "array".
func GoType2APIType(goType reflect.Kind) string {
	if goType >= reflect.Int && goType <= reflect.Uint64 {
		return "integer"
	}

	switch goType {
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "boolean"
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return "array"
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		return "number"
	}

	return ""
}

func ParseField(field reflect.StructField) (schema spec.Schema) {

	fieldTag := field.Tag.Get(StructTagAPI)

	tagMap := make(map[string]string)
	tags := strings.Split(fieldTag, ",")
	for _, value := range tags {
		kv := strings.Split(value, "=")
		tagMap[kv[0]] = kv[1]
	}

	schemaProps := spec.SchemaProps{}

	schemaProps.Type = []string{GoType2APIType(field.Type.Kind())}
	schemaProps.Format = field.Type.Name()

	schema.SchemaProps = schemaProps
	return schema
}

const StructTagAPI = "api"

// Generate properties for Properties in Definitions
func GenProp(resource interface{}) map[string]spec.Schema {
	properties := make(map[string]spec.Schema)

	reflectResource := reflect.ValueOf(resource)
	for i := 0; i < reflectResource.NumField(); i++ {
		properties[strings.ToLower(reflectResource.Type().Field(i).Name)] = ParseField(reflectResource.Type().Field(i))
	}

	return properties
}

func NewAPI(resource interface{}) *spec.SwaggerProps {
	return &spec.SwaggerProps{
		Swagger:  "2.0",
		Host:     "192.168.22.160",
		BasePath: "/v1beta1",
		Info: &spec.Info{
			InfoProps: spec.InfoProps{
				Description: "Clever API 1.4.0 Pre",
				Title:       "Clever API",
				Version:     "v1beta1",
			},
		},

		Definitions: spec.Definitions{
			"tsk": spec.Schema{
				SchemaProps: spec.SchemaProps{
					Type:       []string{"object"},
					Properties: GenProp(resource),
				},
			},
		},

		Paths: &spec.Paths{
			Paths: map[string]spec.PathItem{
				"/tsk": spec.PathItem{
					PathItemProps: spec.PathItemProps{
						Get: &spec.Operation{
							OperationProps: spec.OperationProps{
								Description: "Get Tsk",
								Summary:     "Summary",
								Parameters: []spec.Parameter{
									{
										ParamProps: spec.ParamProps{
											In: ParameterInTypeBody,
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Ref: spec.MustCreateRef("#/definitions/tsk"),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func run() {
	tsk := NewAPI(api.Tsk{})
	data, _ := json.MarshalIndent(tsk, "", "  ")
	fileName := "xx.json"
	file, _ := os.Create(fileName)
	file.Write(data)
}

type Tsk struct {
	Name string
	Age  int
}

func test1() {

}

func main() {
	fmt.Printf("%x\n",time.Now().Nanosecond())
	fmt.Printf("%x\n",time.Now().Nanosecond())
	fmt.Printf("%x\n",time.Now().Nanosecond())
	fmt.Printf("%x\n",time.Now().Nanosecond())
}
