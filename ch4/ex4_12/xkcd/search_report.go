/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package xkcd

import "text/template"

const searchItem = `#{{.Num}} ca. {{.Month}}/{{.Year}}
  Title: {{.Title}}
  Alt: {{.Alt}}
  URL: https://xkcd.com/{{.Num}}/
---------------------------
`

var SearchReport = template.Must(template.New("searchItem").Parse(searchItem))
