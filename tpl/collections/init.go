// Copyright 2017 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collections

import (
	"github.com/spf13/hugo/deps"
	"github.com/spf13/hugo/tpl/internal"
)

const name = "collections"

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		ctx := New(d)

		examples := [][2]string{
			{`delimit: {{ delimit (slice "A" "B" "C") ", " " and " }}`, `delimit: A, B and C`},
			{`echoParam: {{ echoParam .Params "langCode" }}`, `echoParam: en`},
			{`in: {{ if in "this string contains a substring" "substring" }}Substring found!{{ end }}`, `in: Substring found!`},
			{
				`querify 1: {{ (querify "foo" 1 "bar" 2 "baz" "with spaces" "qux" "this&that=those") | safeHTML }}`,
				`querify 1: bar=2&baz=with+spaces&foo=1&qux=this%26that%3Dthose`},
			{
				`querify 2: <a href="https://www.google.com?{{ (querify "q" "test" "page" 3) | safeURL }}">Search</a>`,
				`querify 2: <a href="https://www.google.com?page=3&amp;q=test">Search</a>`},
			{`sort: {{ slice "B" "C" "A" | sort }}`, `sort: [A B C]`},
			{`seq: {{ seq 3 }}`, `seq: [1 2 3]`},
			{`union: {{ union (slice 1 2 3) (slice 3 4 5) }}`, `union: [1 2 3 4 5]`},
		}

		return &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func() interface{} { return ctx },
			Aliases: map[string]interface{}{
				"after":     ctx.After,
				"apply":     ctx.Apply,
				"delimit":   ctx.Delimit,
				"dict":      ctx.Dictionary,
				"echoParam": ctx.EchoParam,
				"first":     ctx.First,
				"in":        ctx.In,
				"index":     ctx.Index,
				"intersect": ctx.Intersect,
				"isSet":     ctx.IsSet,
				"isset":     ctx.IsSet,
				"last":      ctx.Last,
				"querify":   ctx.Querify,
				"shuffle":   ctx.Shuffle,
				"slice":     ctx.Slice,
				"sort":      ctx.Sort,
				"union":     ctx.Union,
				"where":     ctx.Where,
				"seq":       ctx.Seq,
			},
			Examples: examples,
		}

	}

	internal.AddTemplateFuncsNamespace(f)
}