package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	clasp "github.com/synesissoftware/CLASP.Go"
)

const (
	ProgramVersion = "0.1.0"
)

type CancelSearchError struct{}

func (e *CancelSearchError) Error() string {
	return "search cancelled"
}

type SearchFlag uint64

const (
	FindFiles                    = 0x1
	FindDirectories              = 0x2
	_FindLinks                   = 0x4
	_FindDevices                 = 0x8
	FindTypeMask                 = 0xF
	StopOnAccessFailure          = 0x2000
	Recursive                    = 0x10000
	DoNotFollowLinks             = 0x20000
	MarkDirectories              = 0x200000
	RecogniseTildeOnSearchRoot   = 0x4000000
	_DirectorySizeIsNumFiles     = 0x2000000
	IgnoreHiddenEntriesOnWindows = 0x8000000
)

type ProgressResult int

const (
	ProgressCancel   ProgressResult = 0
	ProgressContinue                = 1
)

// type PatternType int

// const (
// 	PatternTypeWildcards
// )

// type Pattern struct {
// 	pattern      string
// 	pattern_type PatternType
// }

/*
"#<Recls::Entry:0x00007faa40158f78

    @file_stat=#<Recls::Ximpl::FileStat
                  dev=0x1000004,
                  ino=143952663,
                  mode=040755,
                  nlink=16,
                  uid=504,
                  gid=20,
                  rdev=0x0,
                  size=512,
                  blksize=4096,
                  blocks=0,
                  atime=2025-02-22 05:06:00.192616436 +1100,
                  mtime=2025-02-22 05:05:58.11502761 +1100,
                  ctime=2025-02-22 05:05:58.11502761 +1100,
                  birthtime=2025-02-13 10:30:46.053860134 +1100>,
    @path=\"/Users/mwsis/dev/mwsis/forks/synesissoftware/recls/recls.Go\",
    @short_path=nil,
    @compare_path=\"/Users/mwsis/dev/mwsis/forks/synesissoftware/recls/recls.Go\",
    @hash=4526718436412439158,
    @drive=nil,
    @directory_path=\"/Users/mwsis/dev/mwsis/forks/synesissoftware/recls/\",
    @directory=\"/Users/mwsis/dev/mwsis/forks/synesissoftware/recls/\",
    @directory_parts=[\"/\", \"Users/\", \"mwsis/\", \"dev/\", \"mwsis/\", \"forks/\", \"synesissoftware/\", \"recls/\"],
    @file_full_name=\"recls.Go\",
    @file_short_name=nil,
    @file_name_only=\"recls\",
    @file_extension=\".Go\",
    @search_directory=nil,
    @search_relative_path=\"/Users/mwsis/dev/mwsis/forks/synesissoftware/recls/recls.Go\",
    @search_relative_directory_path=\"/Users/mwsis/dev/mwsis/forks/synesissoftware/recls/\",
    @search_relative_directory=\"/Users/mwsis/dev/mwsis/forks/synesissoftware/recls/\",
    @search_relative_directory_parts=[\"/\", \"Users/\", \"mwsis/\", \"dev/\", \"mwsis/\", \"forks/\", \"synesissoftware/\", \"recls/\"],
    @dev=16777220,
    @ino=143952663,
    @nlink=16
  >"
*/

/*
#<main.ReclsEntry
    @de=#<*os.unixDirent d recls/>,
    @StatInfo=#<main.ReclsStatInfo {}>,
    @FullPath=#<string /Users/mwsis/dev/mwsis/forks/synesissoftware/recls/recls.Go/recls>,
    @ShortPath=#<string >,
    @Volume=#<string >,
    @DirectoryPath=#<string >,
    @Directory=#<string >,
    @DirectoryParts=#<[]string []>,
    @FileName=#<string >,
    @FileStem=#<string >,
    @FileExtension=#<string >,
    @FileShortName=#<string >,
    @SearchDirectory=#<string /Users/mwsis/dev/mwsis/forks/synesissoftware/recls>,
    @SearchRelativeDirectory=#<string >,
    @SearchRelativeDirectoryParts=#<[]string []>,
    @SearchRelativeDirectoryPath=#<string >,
    @SearchRelativePath=#<string >,
  >
*/

type ReclsStatInfo struct {
	fi os.FileInfo
}

// type reclsEntryInfo struct {

// }

type ReclsEntry struct {
	de                           os.DirEntry
	StatInfo                     ReclsStatInfo
	FullPath                     string
	ShortPath                    string
	Volume                       string
	DirectoryPath                string
	Directory                    string
	DirectoryParts               []string
	FileName                     string
	FileStem                     string
	FileExtension                string
	FileShortName                string
	SearchDirectory              string
	SearchRelativeDirectory      string
	SearchRelativeDirectoryParts []string
	SearchRelativeDirectoryPath  string
	SearchRelativePath           string
}

func (re ReclsEntry) String() string {
	return re.FullPath
}

// TODO: put in different module

type InspectFlags uint64

const (
	None InspectFlags = 0
)

const (
	Alternate InspectFlags = 1 << iota
	HideMinus
	ShowPlus
)

type Debug interface {
	Inspect(flags InspectFlags, depth int) string
}

func do_Inspect(i interface{}, flags InspectFlags) string {
	return fmt.Sprintf("#<%[1]T %[1]s>", i)
}

func do_Inspect_os_DirEntry(de os.DirEntry, flags InspectFlags) string {
	return fmt.Sprintf("#<%[1]T %[1]s>", de)
}

func do_Inspect_os_FileInfo(fi os.FileInfo, flags InspectFlags) string {
	return fmt.Sprintf("#<%[1]T %+[1]v>", fi)

	// type FileInfo interface {
	// 	Name() string       // base name of the file
	// 	Size() int64        // length in bytes for regular files; system-dependent for others
	// 	Mode() FileMode     // file mode bits
	// 	ModTime() time.Time // modification time
	// 	IsDir() bool        // abbreviation for Mode().IsDir()
	// 	Sys() any           // underlying data source (can return nil)
	// }

}

func do_Inspect_String(s string, flags InspectFlags) string {
	return fmt.Sprintf("#<%[1]T %[1]s>", s)
}

func do_Inspect_arr_String(strings []string, flags InspectFlags) string {
	return fmt.Sprintf("#<%[1]T %[1]s>", strings)
}

// re.StatInfo.Inspect(),
// (re.FullPath),
// do_Inspect_String(re.ShortPath),
// do_Inspect_String(re.Volume),
// do_Inspect_String(re.DirectoryPath),
// do_Inspect_String(re.Directory),
// do_Inspect_arr_String(

func (si ReclsStatInfo) Inspect(flags InspectFlags, depth int) string {
	sep0 := " "
	sep := ", "
	if 0 != (Alternate & flags) {
		spacer := strings.Repeat("  ", 1+depth)
		sep0 = "\n" + spacer
		sep = ",\n" + spacer
	}

	return fmt.Sprintf(
		"#<%T%s"+
			"@fi=%s%s"+
			">",
		si, sep0,
		do_Inspect_os_FileInfo(si.fi, flags), sep,
	)
}

func (re ReclsEntry) Inspect(flags InspectFlags, depth int) string {
	sep0 := " "
	sep := ", "
	if 0 != (Alternate & flags) {
		spacer := strings.Repeat("  ", 1+depth)
		sep0 = "\n" + spacer
		sep = ",\n" + spacer
	}

	return fmt.Sprintf(
		"#<%T%s"+
			"@de=%s%s"+
			"@StatInfo=%s%s"+
			"@FullPath=%s%s"+
			"@ShortPath=%s%s"+
			"@Volume=%s%s"+
			"@DirectoryPath=%s%s"+
			"@Directory=%s%s"+
			"@DirectoryParts=%s%s"+
			"@FileName=%s%s"+
			"@FileStem=%s%s"+
			"@FileExtension=%s%s"+
			"@FileShortName=%s%s"+
			"@SearchDirectory=%s%s"+
			"@SearchRelativeDirectory=%s%s"+
			"@SearchRelativeDirectoryParts=%s%s"+
			"@SearchRelativeDirectoryPath=%s%s"+
			"@SearchRelativePath=%s%s"+
			">",
		re, sep0,
		do_Inspect_os_DirEntry(re.de, flags), sep,
		re.StatInfo.Inspect(flags, 1+depth), sep,
		do_Inspect_String(re.FullPath, flags), sep,
		do_Inspect_String(re.ShortPath, flags), sep,
		do_Inspect_String(re.Volume, flags), sep,
		do_Inspect_String(re.DirectoryPath, flags), sep,
		do_Inspect_String(re.Directory, flags), sep,
		do_Inspect_arr_String(re.DirectoryParts, flags), sep,
		do_Inspect_String(re.FileName, flags), sep,
		do_Inspect_String(re.FileStem, flags), sep,
		do_Inspect_String(re.FileExtension, flags), sep,
		do_Inspect_String(re.FileShortName, flags), sep,
		do_Inspect_String(re.SearchDirectory, flags), sep,
		do_Inspect_String(re.SearchRelativeDirectory, flags), sep,
		do_Inspect_arr_String(re.SearchRelativeDirectoryParts, flags), sep,
		do_Inspect_String(re.SearchRelativeDirectoryPath, flags), sep,
		do_Inspect_String(re.SearchRelativePath, flags), sep,
	)
}

func calcDirectoryParts(directory string) []string {
	directory_parts := strings.Split(directory, "/")

	for i, directory_part := range directory_parts {
		directory_parts[i] = "/" + directory_part
	}

	return directory_parts
}

func createEntry(search_dir_full_path string, entry_full_path string, search_relative_path string, de os.DirEntry, fi os.FileInfo) ReclsEntry {

	volume := filepath.VolumeName(entry_full_path)
	directory_path := filepath.Dir(entry_full_path)
	file_base := filepath.Base(entry_full_path)
	file_ext := filepath.Ext(entry_full_path)
	file_stem := file_base[0:(len(file_base) - len(file_ext))]
	file_short_name := ""
	if 0 != len(volume) {
		// TODO: calculate short-name
	}

	directory_parts := calcDirectoryParts(directory_path)

	return ReclsEntry{
		de:                      de,
		StatInfo:                ReclsStatInfo{fi},
		FullPath:                entry_full_path,
		ShortPath:               "",
		Volume:                  volume,
		DirectoryPath:           directory_path,
		Directory:               directory_path,
		DirectoryParts:          directory_parts,
		FileName:                file_base,
		FileStem:                file_stem,
		FileExtension:           file_ext,
		FileShortName:           file_short_name,
		SearchDirectory:         search_dir_full_path,
		SearchRelativeDirectory: "",
		// SearchRelativeDirectoryParts: []string,
		SearchRelativeDirectoryPath: "",
		SearchRelativePath:          search_relative_path,
	}
}

type ReclsSearchFunc func(re ReclsEntry, err error) error

type ReclsProgressFn func(directory string) ProgressResult

/*
type SearchParams struct {
	root_dir string
	// patterns     []Pattern // zero or more patterns, against any one of which an entry can be matched
	patterns     string
	search_flags SearchFlag
}

func (search_params SearchParams) SearchFlags(args ...interface{}) SearchParams {

	var search_flags uint64 = 0

	for _, arg := range args {
		switch arg.(type) {
		case SearchFlag:

			search_flags |= arg.(uint64)
		default:

			panic(fmt.Sprintf("function `SearchFlags` given an argument of type `%T`, but should be of type `SearchFlag`", arg))
		}
	}

	return search_params
}
*/

// T.B.C.
//
// # Returns:
//   - `num_found` is the number of entries that were encountered that match
//     the given search parameters;
//   - `num_skipped` is the number of entries that were skipped during
//     processing;
//   - `err` indicates the level of success in the function. A value of
//     `nil` indicates that an uninterrupted search has run to completion. A
//     value of `CancelSearchError` indicates that the callback function (`f`)
func RSearch(search_dir string, patterns string, flags uint64, f ReclsSearchFunc) (num_found uint64, num_skipped uint64, err error) {
	fmt.Printf("RSearch(search_dir{%[1]T}='%[1]v', patterns{%[2]T}='%[2]v', flags{%[3]T}=0x%[3]016x, f{%[4]T}=%[4]v)\n", search_dir, patterns, flags, f)

	num_found = 0
	num_skipped = 0

	search_dir_full_path, err := filepath.Abs(search_dir)

	if err != nil {
		return 0, 0, err
	}

	fmt.Printf("search_dir_full_path=%v\n", search_dir_full_path)

	entries, err := os.ReadDir(search_dir_full_path)
	if err != nil {
		return 0, 0, err
	}

	for _, de := range entries {

		entry_full_path, err := filepath.Abs(filepath.Join(search_dir_full_path, de.Name()))

		if err != nil {
			return 0, 0, err
		}

		search_relative_path, err := filepath.Rel(search_dir_full_path, entry_full_path)

		if err != nil {
			return 0, 0, err
		}

		fs, err := os.Stat(entry_full_path)

		if err != nil {
			return 0, 0, err
		}

		re := createEntry(search_dir_full_path, entry_full_path, search_relative_path, de, fs)

		f(re, nil)

		num_found += 1
	}

	return num_found, num_skipped, nil
}

/*
func SearchByParams(search_params SearchParams, sf ReclsSearchFunc, pf ReclsProgressFn) (num_found uint64, num_skipped uint64, err error) {

	var _ = search_params

	num_found = 0
	num_skipped = 0

	return num_found, num_skipped, nil
}
*/

func main() {
	fmt.Printf("InspectFlags.None=%d\n", None)
	fmt.Printf("InspectFlags.Alternate=%d\n", Alternate)
	fmt.Printf("InspectFlags.HideMinus=%d\n", HideMinus)
	fmt.Printf("InspectFlags.ShowPlus=%d\n", ShowPlus)

	args := clasp.Parse(os.Args, clasp.ParseParams{})

	if args.FlagIsSpecified(clasp.HelpFlag()) {

		clasp.ShowUsage(nil, clasp.UsageParams{

			Version:   ProgramVersion,
			InfoLines: []string{"CLASP.Go Examples", "", ":version:", ""},
		})
	}

	if args.FlagIsSpecified(clasp.VersionFlag()) {

		clasp.ShowVersion(nil, clasp.UsageParams{Version: ProgramVersion})
	}

	var root_dir string

	if 0 == len(args.Values) {
		root_dir = "."
	} else {
		root_dir = args.Values[0].Value
	}

	fmt.Printf("root_dir=%v\n", root_dir)

	num_found, num_skipped, err := RSearch(root_dir, "*.*", uint64(0), func(re ReclsEntry, err error) error {

		// fmt.Printf("\t\tentry=%v (%T): %s %+v\n", re, re, re.Inspect(Alternate, 2), re)
		fmt.Printf("\tentry=%v (%T): %s\n\n", re, re, re.Inspect(Alternate, 8))

		return nil
	})

	if err != nil {

		fmt.Printf("failed to search: %v\n", err)
	} else {

		fmt.Printf("\tfound %v; skipped %v\n", num_found, num_skipped)
	}
}
