package mailslurper

import "github.com/labstack/echo"

/*
ApplyThemeRequest is used to update and apply a theme change
*/
type ApplyThemeRequest struct {
	Theme string `json:"theme"`
}

/*
NewApplyThemeRequest creates a new struct from a POST request
*/
func NewApplyThemeRequest(ctx echo.Context) (*ApplyThemeRequest, error) {
	var err error
	var result *ApplyThemeRequest

	err = ctx.Bind(&result)
	return result, err
}
