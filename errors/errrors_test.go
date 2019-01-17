package errors

import (
	"fmt"
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestErrors(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Errors", func() {
		g.It("return a new ion error", func() {
			err := fmt.Errorf("json: invalid key")
			ierr := Errors("a response body", 404, "something went wrong: %v", err.Error())

			Expect(ierr.ResponseBody).To(Equal("a response body"))
			Expect(ierr.ResponseStatus).To(Equal(404))
			Expect(ierr.Error()).To(Equal("ionic: (404) something went wrong: json: invalid key"))
		})
	})
}
