package function

/*
 * @Description:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2022-03-14 19:55:10
 * @LastEditors: shahao
 * @LastEditTime: 2022-03-14 19:57:29
 */

import (
	"net/url"
	"strings"
)

type UrlBuilder struct {
	u     *url.URL
	query url.Values
}

func ParseUrl(rawUrl string) *UrlBuilder {
	ub := &UrlBuilder{}
	ub.u, _ = url.Parse(rawUrl)
	ub.query = ub.u.Query()
	return ub
}

func (builder *UrlBuilder) AddQuery(name, value string) *UrlBuilder {
	builder.query.Add(name, value)
	return builder
}

func (builder *UrlBuilder) AddQueries(queries map[string]string) *UrlBuilder {
	for name, value := range queries {
		builder.AddQuery(name, value)
	}
	return builder
}

func (builder *UrlBuilder) GetQuery() url.Values {
	return builder.query
}

func (builder *UrlBuilder) GetURL() *url.URL {
	return builder.u
}

func (builder *UrlBuilder) Build() *url.URL {
	builder.u.RawQuery = builder.query.Encode()
	return builder.u
}

func (builder *UrlBuilder) BuildStr() string {
	return builder.Build().String()
}

func UrlJoin(parts ...string) string {
	sep := "/"
	var ss []string
	for i, part := range parts {
		part = strings.TrimSpace(part)
		var (
			from = 0
			to   = len(part)
		)
		if strings.Index(part, sep) == 0 {
			from = 1
		}
		if strings.LastIndex(part, sep) == len(part)-1 {
			to = len(part) - 1
		}
		part = part[from:to]

		ss = append(ss, part)
		if i != len(parts)-1 {
			ss = append(ss, sep)
		}
	}
	return strings.Join(ss, "")
}
