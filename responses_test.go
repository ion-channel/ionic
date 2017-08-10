package ionic

import (
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestResponses(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Error Response", func() {
		g.It("should return a new error response", func() {
			msg := "foo error"
			fields := []string{"bar"}
			status := http.StatusUnauthorized

			b, s := NewErrorResponse(msg, fields, status)
			Expect(string(b)).To(ContainSubstring(msg))
			Expect(string(b)).To(ContainSubstring(fields[0]))
			Expect(string(b)).To(ContainSubstring(string(status)))
			Expect(s).To(Equal(status))
		})
	})
}
