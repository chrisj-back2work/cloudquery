// Code generated by codegen using template {{.TemplateFilename}}; DO NOT EDIT.

package {{.AWSService | ToLower}}

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"{{.TypesImport}}"
{{range .Imports}}	{{.}}
{{end}}
)

func {{.TableFuncName}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

func {{.Table.Resolver}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
  return errors.WithStack(client.ListAndDetailResolver(ctx, meta, res, list{{.AWSSubService}}, list{{.AWSSubService}}Detail))
}

func list{{.AWSSubService}}(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().{{.AWSService | ToCamel}}

{{template "resolve_parent_defs.go.tpl" .}}
	input := {{.AWSService | ToLower}}.{{.ListMethod}}Input{
{{range .CustomInputs}}{{.}}
{{end}}
{{template "resolve_parent_vars.go.tpl" .}}
	}

	for {
		response, err := svc.{{.ListMethod}}(ctx, &input)
		if err != nil {
			return errors.WithStack(err)
		}
		for _, item := range response.{{.PaginatorListName}} {
			detailChan <- item
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func list{{.AWSSubService}}Detail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	cl := meta.(*client.Client)
	item := listInfo.({{.PaginatorListType}})
	svc := cl.Services().{{.AWSService | ToCamel}}
	response, err := svc.{{.GetMethod}}(ctx, &{{.AWSService | ToLower}}.{{.GetMethod}}Input{
{{range $v := .GetAndListOrder}}	{{$v}}: {{index $.MatchedGetAndListFields $v}},
{{end}}
	})
	if err != nil {
		{{.CustomErrorBlock}}
		if cl.IsNotFoundError(err) {
			return
		}
		errorChan <- errors.WithStack(err)
		return
	}
	resultsChan <- *response.{{.ItemName}}
}

{{if .HasTags}}
func resolve{{.AWSService}}{{.AWSSubService}}Tags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().{{.AWSService | ToCamel}}
	item := resource.Item.({{.PaginatorListType}})
	params := {{.AWSService | ToLower}}.ListTagsForResourceInput{
		ResourceARN: {{.CustomTagField | Coalesce "item.ARN"}},
	}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return errors.WithStack(err)
		}
		client.TagsIntoMap(result.Tags, tags)
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return errors.WithStack(resource.Set(c.Name, tags))
}
{{end}}
