// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractSeparatorsToIndex(opts ...interface{}) *opt.SeparatorsToIndexOption {
	for _, o := range opts {
		if v, ok := o.(opt.SeparatorsToIndexOption); ok {
			return &v
		}
	}
	return nil
}