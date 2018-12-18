package core

import (
	"path"

	"github.com/go-openapi/jsonpointer"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

type NodeData struct {
	Key string
	Ref string
}

type DocumentVisitor func(node interface{}, data NodeData)

func Walk(node interface{}, data NodeData, visitor DocumentVisitor) {
	visitor(node, data)

	switch v := node.(type) {
	case *loads.Document:
		Walk(v.Spec(), NodeData{}, visitor)

	case *spec.Swagger:
		for apiPath, pathItem := range v.Paths.Paths {
			Walk(&pathItem, NodeData{Key: apiPath, Ref: path.Join("#/paths", jsonpointer.Escape(apiPath))}, visitor)
		}
		Walk(&v.Definitions, NodeData{Ref: "#/definitions/"}, visitor)
		for parameterName, parameter := range v.Parameters {
			Walk(&parameter, NodeData{Key: parameterName, Ref: path.Join("#/parameters/", parameterName)}, visitor)
		}
		for responseName, response := range v.Responses {
			Walk(&response, NodeData{Key: responseName, Ref: path.Join("#/responses/", responseName)}, visitor)
		}

	case *spec.PathItem:
		if v.Get != nil {
			Walk(v.Get, NodeData{Key: "Get", Ref: path.Join(data.Ref, "get")}, visitor)
		}
		if v.Post != nil {
			Walk(v.Post, NodeData{Key: "Post", Ref: path.Join(data.Ref, "post")}, visitor)
		}
		if v.Put != nil {
			Walk(v.Put, NodeData{Key: "Put", Ref: path.Join(data.Ref, "put")}, visitor)
		}
		if v.Patch != nil {
			Walk(v.Patch, NodeData{Key: "Patch", Ref: path.Join(data.Ref, "patch")}, visitor)
		}
		if v.Delete != nil {
			Walk(v.Delete, NodeData{Key: "Delete", Ref: path.Join(data.Ref, "delete")}, visitor)
		}
		if v.Options != nil {
			Walk(v.Options, NodeData{Key: "Options", Ref: path.Join(data.Ref, "options")}, visitor)
		}
		if v.Head != nil {
			Walk(v.Head, NodeData{Key: "Head", Ref: path.Join(data.Ref, "head")}, visitor)
		}

	case *spec.Operation:
		for _, parameter := range v.Parameters {
			Walk(&parameter, NodeData{}, visitor)
		}
		for statusCode, response := range v.Responses.StatusCodeResponses {
			Walk(&response, NodeData{Key: string(statusCode)}, visitor)
		}

	case *spec.Response:
		if v.Schema != nil {
			Walk(v.Schema, NodeData{}, visitor)
		}

	case *spec.Parameter:
		if v.Schema != nil {
			Walk(v.Schema, NodeData{}, visitor)
		}

	case *spec.Definitions:
		for name, schema := range *v {
			Walk(&schema, NodeData{Key: name, Ref: path.Join(data.Ref, name)}, visitor)
		}

	case *spec.Schema:
		for name, schema := range v.Properties {
			Walk(&schema, NodeData{Key: name, Ref: path.Join(data.Ref, name)}, visitor)
		}
	}
}
