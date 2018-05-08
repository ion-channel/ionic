package penname

import (
	"fmt"
	"io"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestPenname(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Closing", func() {
		var p *PenName

		g.BeforeEach(func() {
			p = &PenName{}
		})

		g.It("should record when it is closed", func() {
			Expect(p.Closed).To(BeFalse())
			err := p.Close()
			Expect(err).NotTo(HaveOccurred())
			Expect(p.Closed).To(BeTrue())
		})

		g.It("should return an error when told to", func() {
			errText := "I wasn't able to close"

			p.ReturnError(fmt.Errorf(errText))
			err := p.Close()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(errText))
		})

		g.It("should function as a closer", func() {
			err := func(c io.Closer) error {
				return c.Close()
			}(p)
			Expect(err).NotTo(HaveOccurred())

			p.ReturnError(fmt.Errorf("Some error"))
			err = func(c io.Closer) error {
				return c.Close()
			}(p)
			Expect(err).To(HaveOccurred())
		})

		g.It("should reset its closed state when told to", func() {
			Expect(p.Closed).ToNot(BeTrue())
			p.Close()
			Expect(p.Closed).To(BeTrue())
			p.Reset()
			Expect(p.Closed).ToNot(BeTrue())
		})
	})

	g.Describe("Writing", func() {
		var p *PenName

		g.BeforeEach(func() {
			p = &PenName{}
		})

		g.It("should record what is written to it", func() {
			t := "Nothing to see here move along"

			n, err := p.Write([]byte(t))
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(len(t)))
			Expect(p.Written).To(Equal([]byte(t)))
		})

		g.It("should return an error when told to", func() {
			errText := "Something went wrong"
			p.ReturnError(fmt.Errorf(errText))

			_, err := p.Write([]byte("I'm trying to do something"))
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(errText))
		})

		g.It("should function as a writer", func() {
			t := "Nothing to see here move along"

			fmt.Fprint(p, t)
			Expect(p.Written).To(Equal([]byte(t)))
		})

		g.It("should reset what was written when told to", func() {
			t := "Nothing to see here move along"

			n, err := p.Write([]byte(t))
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(len(t)))
			Expect(p.Written).To(Equal([]byte(t)))

			p.Reset()
			Expect(p.Written).To(Equal([]byte("")))
		})
	})
}
