/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// Param defines api parameters from http://omdbapi.com/.
package poster

type Params map[Param]string

// Param is an api parameter definition from http://omdbapi.com/
type Param = struct {
	Name        string
	Required    bool
	Default     string
	Description string
}

var Search = Param{
	Name:        "s",
	Required:    true,
	Default:     "1",
	Description: "Movie title to search for.",
}

var Page = Param{
	Name:        "page",
	Required:    false,
	Default:     "1",
	Description: "Page number to return. (1-100)",
}

var Id = Param{
	Name:        "i",
	Required:    false,
	Default:     "",
	Description: "A valid IMDb ID (e.g. tt1285016)",
}

var Title = Param{
	Name:        "t",
	Required:    false,
	Default:     "",
	Description: "Movie title to search for.",
}

var Type = Param{
	Name:        "type",
	Required:    false,
	Default:     "",
	Description: "Type of result to return. (movie, series, episode)",
}

var Year = Param{
	Name:        "y",
	Required:    false,
	Default:     "",
	Description: "Year of release.",
}

var Plot = Param{
	Name:        "plot",
	Required:    false,
	Default:     "short",
	Description: "Return short or full plot. (short, full)",
}

var Return = Param{
	Name:        "r",
	Required:    false,
	Default:     "json",
	Description: "The data type to return.",
}

var Callback = Param{
	Name:        "callback",
	Required:    false,
	Default:     "",
	Description: "JSONP callback name.",
}

var APIVersion = Param{
	Name:        "v",
	Required:    false,
	Default:     "1",
	Description: "API version (reserved for future use).",
}

var APIKey = Param{
	Name:        "apikey",
	Required:    true,
	Default:     "",
	Description: "API Key",
}
