package ctfvisitor

import (
	"context"
	"log"

	"github.com/tebeka/selenium"
)

type Handler func(ctx context.Context, wd selenium.WebDriver) error

// VisitHandler is a generic wrapper around Get() that also installs cookies
// first before visiting the page
func CookieHandler(path string, cookies []*selenium.Cookie) Handler {
	path := path
	cookies := cookies
	h := func(ctx context.Context, wd selenium.WebDriver) error {
		for _, cookie := range cookies {
			if err := wd.AddCookie(cookie); err != nil {
				return err
			}
		}
		log.Printf("getting: %v", path)

		return wd.Get(path)
	}

	return h
}
