package find

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

type FindOptions struct {
	Recursive bool

	StopAtFirstMatch bool
	RegularFilesOnly bool
	DirectoriesOnly  bool
	MatchRegex       *regexp.Regexp
	MatchExtension   string

	// Sort newest to oldest
	SortByRecentModTime bool

	// Sort oldest to newest
	ReverseSortByRecentModTime bool

	// Add these later
	// // Sort by name
	// SortByName bool

	// // Sort by name reversed
	// ReverseSortByName bool
}

type Found struct {
	Path string
	Info os.FileInfo
}

func Find(dir string, opts FindOptions) ([]Found, error) {
	var result []Found

	if opts.Recursive {
		err := filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if opts.RegularFilesOnly {
				if !fi.Mode().IsRegular() {
					return nil
				}
			}

			if opts.DirectoriesOnly {
				if !fi.Mode().IsDir() {
					return nil
				}
			}

			if nil != opts.MatchRegex {
				if !opts.MatchRegex.MatchString(fi.Name()) {
					return nil
				}
			}

			if opts.MatchExtension != "" {
				if opts.MatchExtension != filepath.Ext(fi.Name()) {
					return nil
				}
			}

			result = append(result, Found{
				Path: path,
				Info: fi,
			})

			if opts.StopAtFirstMatch {
				return io.EOF
			}

			return nil
		})

		if err == io.EOF {
			err = nil
		}

		if err != nil {
			return result, err
		}
	} else {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return result, err
		}

		for _, fi := range files {
			if opts.RegularFilesOnly {
				if !fi.Mode().IsRegular() {
					continue
				}
			}

			if opts.DirectoriesOnly {
				if !fi.Mode().IsDir() {
					continue
				}
			}

			if nil != opts.MatchRegex {
				if !opts.MatchRegex.MatchString(fi.Name()) {
					continue
				}
			}

			if opts.MatchExtension != "" {
				if opts.MatchExtension != filepath.Ext(fi.Name()) {
					continue
				}
			}

			result = append(result, Found{
				Path: filepath.Join(dir, fi.Name()),
				Info: fi,
			})

			if opts.StopAtFirstMatch {
				break
			}
		}
	}

	if opts.SortByRecentModTime {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Info.ModTime().After(result[j].Info.ModTime())
		})
	} else if opts.ReverseSortByRecentModTime {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Info.ModTime().Before(result[j].Info.ModTime())
		})
	}

	if len(result) > 0 {
		return result, nil
	}

	return result, errors.New("no files found")
}
