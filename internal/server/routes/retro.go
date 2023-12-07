package routes

import "github.com/labstack/echo/v4"


func RetroPage(c echo.Context) error {
    // display a form for every category in given retro
    return nil
}

func RetroItemCreate(c echo.Context) error {
    // handle data from an items on retro box page
    // get name of the category and content of new entry
    return nil
}

func RetroNewForm(c echo.Context) error {
    // not sure that this is a form lol
    return nil
}

func RetroCreate(c echo.Context) error {
    // handle a data from a retro form (or request)
    // it should use template and current user, and add a date of creation
    return nil
}

func Templates(c echo.Context) error {
    // render a list of templates for some user
    return nil
}

func TemplatesNew(c echo.Context) error {
    // render a form for new tempalates
    return nil
}

func TempalatesCreate(c echo.Context) error {
    // should handle data to create new template
    return nil
}
