package find

import (

	// "fmt"

	"regexp"
	"testing"

	. "github.com/franela/goblin"
)

func TestLoad(t *testing.T) {
	g := Goblin(t)
	g.Describe("Load", func() {
		g.It("Should find files correctly", func() {

			opts := FindOptions{}

			opts.RegularFilesOnly = true
			opts.SortByRecentModTime = true
			res, err := Find("testdata/", opts)
			g.Assert(err).IsNil()
			g.Assert(len(res)).Equal(3)

			// fmt.Println(res)

			opts = FindOptions{}

			opts.RegularFilesOnly = true
			opts.SortByRecentModTime = true
			opts.Recursive = true
			res, err = Find("testdata/", opts)
			g.Assert(err).IsNil()
			g.Assert(len(res)).Equal(8)

			// fmt.Println(res)

			opts = FindOptions{}

			opts.DirectoriesOnly = true
			opts.Recursive = true
			res, err = Find("testdata/", opts)
			g.Assert(err).IsNil()
			g.Assert(len(res)).Equal(4)

			// fmt.Println(res)

			re, err := regexp.Compile(`^.+\.(txt)$`)
			g.Assert(err).IsNil()
			opts = FindOptions{}

			opts.RegularFilesOnly = true
			opts.MatchRegex = re
			res, err = Find("testdata/", opts)
			g.Assert(err).IsNil()
			g.Assert(len(res)).Equal(3)

			// fmt.Println(res)

			re, err = regexp.Compile(`^.+\.(bin)$`)
			g.Assert(err).IsNil()
			opts = FindOptions{}

			opts.RegularFilesOnly = true
			opts.MatchRegex = re
			_, err = Find("testdata/", opts)
			g.Assert(err).IsNotNil()

			// fmt.Println(res)

			opts = FindOptions{}

			opts.RegularFilesOnly = true
			opts.MatchExtension = ".txt"
			res, err = Find("testdata/", opts)
			g.Assert(err).IsNil()
			g.Assert(len(res)).Equal(3)

			// fmt.Println(res)

		})
	})
}
