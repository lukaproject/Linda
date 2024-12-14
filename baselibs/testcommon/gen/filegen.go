package gen

import (
	"Linda/baselibs/abstractions/xos"
	"errors"
	"math/rand"
	"os"
	"path"
)

type FileGenerateRoles struct {
	// Must be a absolute path
	RootDir string

	MaxDirDepth int
	MaxNameLen  int
	// total file count generated.
	MaxCount int
}

// FileGenerate
// this is a function to generate a bunch of files, in many directories.
// You can generate from you input roles.
func FileGenerate(roles FileGenerateRoles) (err error) {
	if !xos.PathExists(roles.RootDir) {
		err = errors.New("RootDir is not existed")
		return
	}
	roles.MaxDirDepth = max(roles.MaxDirDepth, 0)
	roles.MaxNameLen = max(roles.MaxNameLen, 1)
	dirs, err := dirOnlyGenerate(dirOnlyGenerateRoles{
		RootDir:       roles.RootDir,
		MaxDirDepth:   roles.MaxDirDepth,
		MaxDirNameLen: roles.MaxNameLen,
		MaxDirsCount:  roles.MaxCount,
	})
	if err != nil {
		return err
	}
	dirs = append(dirs, roles.RootDir)
	rest := roles.MaxCount
	for _, dir := range dirs {
		cur := rand.Intn(rest) + 1
		n, err := fileOnlyGenerate(fileOnlyGenerateRoles{
			RootDir:        dir,
			MaxFileNameLen: roles.MaxNameLen,
			MinCount:       1,
			MaxCount:       cur,
		})
		if err != nil {
			// fmt.Printf("dir=%s, not created\n", dir)
			return err
		}
		rest -= n
		if rest <= 0 {
			return nil
		}
	}
	return
}

type fileOnlyGenerateRoles struct {
	RootDir        string
	MaxFileNameLen int
	MaxCount       int
	MinCount       int
}

// fileOnlyGenerate
// generate files in a root dir.
// return n is how many files generated.
func fileOnlyGenerate(roles fileOnlyGenerateRoles) (n int, err error) {
	if !xos.PathExists(roles.RootDir) {
		err = errors.New("RootDir is not existed")
		return
	}
	n = roles.MinCount + rand.Intn(roles.MaxCount-roles.MinCount+1)
	fnameList := make([]string, 0, n)
	for i := 0; i < n; i++ {
		var fname string
		fname, err = StrGenerate(CharsetLowerCase+CharsetUpperCase, 1, roles.MaxCount)
		if err != nil {
			return
		}
		fnameList = append(fnameList, path.Join(roles.RootDir, fname))
	}

	for _, fullFileName := range fnameList {
		if err = xos.Touch(fullFileName); err != nil {
			return
		}
	}

	return
}

type dirOnlyGenerateRoles struct {
	RootDir       string
	MaxDirDepth   int
	MaxDirNameLen int
	MaxDirsCount  int
}

// dirOnlyGenerate generate directories follow the roles you input.
func dirOnlyGenerate(roles dirOnlyGenerateRoles) (dirs []string, err error) {
	dirs = make([]string, 0)
	if roles.MaxDirDepth <= 0 {
		// no more depth, create directory.
		// fmt.Printf("dir=%s, created\n", roles.RootDir)
		xos.MkdirAll(roles.RootDir, os.ModePerm)
		dirs = append(dirs, roles.RootDir)
		return
	}
	if roles.MaxDirsCount <= 0 {
		return
	}
	rest := roles.MaxDirsCount
	cur := rand.Intn(roles.MaxDirsCount) + 1
	rest -= cur
	for i := 0; i < cur; i++ {
		var dirname string
		dirname, err = StrGenerate(CharsetLowerCase+CharsetUpperCase, 1, roles.MaxDirNameLen)
		if err != nil {
			return nil, err
		}
		fullDirName := path.Join(roles.RootDir, dirname)
		var subDirs []string
		subDirs, err = dirOnlyGenerate(dirOnlyGenerateRoles{
			RootDir:       fullDirName,
			MaxDirDepth:   roles.MaxDirDepth - 1,
			MaxDirNameLen: roles.MaxDirNameLen,
			MaxDirsCount:  rest,
		})
		if err != nil {
			return nil, err
		}
		dirs = append(dirs, subDirs...)
		rest -= len(subDirs)
		if rest == 0 {
			return
		}
	}
	return
}
