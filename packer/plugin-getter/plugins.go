package plugingetter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/packer/hcl2template/addrs"
	"github.com/hashicorp/packer/packer-plugin-sdk/tmp"
)

type Requirements []*Requirement

// Requirement describes a required plugin and how it is installed. Usually a list
// of required plugins is generated from a config file. From it we check what
// is actually installed and what needs to happen to get in the desired state.
type Requirement struct {
	// How the user named the plugin
	Accessor string

	// Something like github.com/hashicorp/packer-plugin-amazon
	Identifier *addrs.Plugin

	// VersionConstraints as defined by user. Empty ( to be avoided ) means
	// highest found version.
	VersionConstraints version.Constraints
}

type BinaryInstallationOptions struct {
	// Usually "_x4" for the 4th API version protocol
	// Should be "_x4.exe" on windows.
	// Extension is just a file suffix really.
	Extension string
	// OS and ARCH usually should be runtime.GOOS and runtime.ARCH, they allow
	// to pick the correct binary.
	OS, ARCH string

	Checksummers []Checksummer
}

type ListInstallationsOptions struct {
	// FromFolders where plugins could be installed. Paths should be absolute for
	// safety but can also be relative.
	FromFolders []string

	BinaryInstallationOptions
}

func (pr Requirement) filenamePrefix() string {
	return "packer-plugin-" + pr.Identifier.Type + "_"
}

func (opts BinaryInstallationOptions) filenameSuffix() string {
	return "_" + opts.OS + "_" + opts.ARCH + opts.Extension
}

// ListInstallations lists unique installed versions of plugin Requirement pr
// with opts as a filter.
//
// Installations are sorted by version and one binary per version is returned.
// Last binary detected takes precedence: in the order 'FromFolders' option.
//
// You must pass at least one option to Checksumers for a binary to be even
// consider.
func (pr Requirement) ListInstallations(opts ListInstallationsOptions) (InstallList, error) {
	res := InstallList{}
	filenamePrefix := pr.filenamePrefix()
	filenameSuffix := opts.filenameSuffix()
	log.Printf("[TRACE] listing potential installations for %q that match %q. %#v", pr.Identifier.ForDisplay(), pr.VersionConstraints, opts)
	for _, knownFolder := range opts.FromFolders {
		glob := filepath.Join(knownFolder, pr.Identifier.Hostname, pr.Identifier.Namespace, pr.Identifier.Type, filenamePrefix+"*"+filenameSuffix)

		matches, err := filepath.Glob(glob)
		if err != nil {
			return nil, fmt.Errorf("ListInstallations: %q failed to list binaries in folder: %v", pr.Identifier.String(), err)
		}
		for _, path := range matches {
			fname := filepath.Base(path)
			if fname == "." {
				continue
			}

			// base name could look like packer-plugin-amazon_v1.2.3_darwin_amd64.0_x4
			versionStr := strings.TrimPrefix(fname, filenamePrefix)
			versionStr = strings.TrimSuffix(versionStr, filenameSuffix)
			pv, err := version.NewVersion(versionStr)
			if err != nil {
				// could not be parsed, ignoring the file
				log.Printf("found %q with an incorrect %q version, ignoring it. %v", path, versionStr, err)
				continue
			}

			// no constraint means always pass
			if !pr.VersionConstraints.Check(pv) {
				log.Printf("[TRACE] version %q of file %q does not match constraint %q", versionStr, path, pr.VersionConstraints.String())
				continue
			}

			checksumOk := false
			for _, checksummer := range opts.Checksummers {

				cs, err := checksummer.GetChecksumOfFile(path)
				if err != nil {
					log.Printf("[TRACE] GetChecksumOfFile(%q) failed: %v", path, err)
					continue
				}

				if err := checksummer.ChecksumFile(cs, path); err != nil {
					log.Printf("[TRACE] ChecksumFile(%q) failed: %v", path, err)
					continue
				}
				checksumOk = true
				break
			}
			if !checksumOk {
				log.Printf("[TRACE] No checksum found for %q ignoring possibly unsafe binary", path)
				continue
			}

			log.Printf("found %q", path)
			res.InsertSortedUniq(&Installation{
				BinaryPath: path,
				Version:    versionStr,
			})
		}
	}
	return res, nil
}

// InstallList is a list of installs
type InstallList []*Installation

// InsertSortedUniq inserts the installation in the right spot in the list by
// comparing the version lexicographically.
// A Duplicate version will replace any already present version.
func (l *InstallList) InsertSortedUniq(install *Installation) {
	pos := sort.Search(len(*l), func(i int) bool { return (*l)[i].Version >= install.Version })
	if len(*l) > pos && (*l)[pos].Version == install.Version {
		(*l)[pos] = install
		return
	}
	(*l) = append((*l), nil)
	copy((*l)[pos+1:], (*l)[pos:])
	(*l)[pos] = install
}

// Installation describes a plugin installation
type Installation struct {
	// path to where binary is installed, if installed.
	// Ex: /usr/azr/.packer.d/plugins/github.com/hashicorp/packer-plugin-amazon/packer-plugin-amazon_v1.2.3_darwin_amd64
	BinaryPath string

	// Version of this plugin, if installed and versionned. Ex:
	//  * v1.2.3 for packer-plugin-amazon_v1.2.3_darwin_.0_x5
	//  * empty  for packer-plugin-amazon
	Version string
}

// InstallOptions describes the possible options for installing the plugin that
// fits the plugin Requirement.
type InstallOptions struct {
	// Different means to get releases, sha256 and binary files.
	Getters []Getter

	// Any downloaded binary and checksum file will be put in this folder.
	InFolders []string

	// If empty then we will try to fetch it.
	Version string

	BinaryInstallationOptions
}

type GetOptions struct {
	PluginRequirement *Requirement

	// If empty then we will try to fetch it.
	Version string

	BinaryInstallationOptions
}

func (gp *GetOptions) ExpectedFilename() string {
	return gp.PluginRequirement.filenamePrefix() + gp.Version + gp.BinaryInstallationOptions.filenameSuffix()
}

// A Getter helps get the appropriate files to download a binary.
type Getter interface {
	// Get:
	//  * 'releases'
	//  * 'sha256'
	//  * 'binary'
	Get(what string, opts GetOptions) (io.ReadCloser, error)
}

type Release struct {
	Version string `json:"version"`
}

func ParseReleases(f io.ReadCloser) ([]Release, error) {
	var releases []Release
	defer f.Close()
	return releases, json.NewDecoder(f).Decode(&releases)
}

func (pr *Requirement) InstallLatest(opts InstallOptions) (*Installation, error) {

	getters := opts.Getters

	getOpts := GetOptions{
		pr,
		opts.Version,
		opts.BinaryInstallationOptions,
	}

	if getOpts.Version == "" {
		log.Printf("[TRACE] getting available versions for the the %s plugin", pr.Identifier.ForDisplay())
		for _, getter := range getters {

			releasesFile, err := getter.Get("releases", getOpts)
			if err != nil {
				err := fmt.Errorf("%q getter could not get release: %w", getter, err)
				log.Printf("[TRACE] %s", err.Error())
				continue
			}

			releases, err := ParseReleases(releasesFile)
			if err != nil {
				err := fmt.Errorf("could not parse release: %w", err)
				log.Printf("[TRACE] %s", err.Error())
				continue
			}
			if len(releases) == 0 {
				err := fmt.Errorf("no release found")
				log.Printf("[TRACE] %s", err.Error())
				continue
			}
			versions := version.Collection{}
			for _, release := range releases {
				v, err := version.NewVersion(release.Version)
				if err != nil {
					panic(err)
				}
				if pr.VersionConstraints.Check(v) {
					versions = append(versions, v)
				}
			}
			if len(versions) == 0 {
				err := fmt.Errorf("no matching version found in releases. In %v", releases)
				log.Printf("[TRACE] %s", err.Error())
				continue
			}
			sort.Sort(sort.Reverse(versions))
			log.Printf("[DEBUG] found %s", versions)
			getOpts.Version = "v" + versions[0].String()
			break
		}
	}

	if getOpts.Version == "" {
		err := fmt.Errorf("no release version found for the %s plugin", pr.Identifier.ForDisplay())
		return nil, err
	}

	outputFile := filepath.Join(
		// Pick last folder as it's the one with the highest priority
		opts.InFolders[len(opts.InFolders)-1],
		// add expected full path
		filepath.Join(pr.Identifier.Parts()...),
		// Get expected file name
		getOpts.ExpectedFilename(),
	)

	// create directories if need be
	if err := os.MkdirAll(filepath.Dir(outputFile), 0755); err != nil {
		err := fmt.Errorf("could not create plugin folder %q: %w", filepath.Dir(outputFile), err)
		log.Printf("[TRACE] %s", err.Error())
		return nil, err
	}

	log.Printf("[TRACE] selecting the %q version to install the %s plugin in %q...", getOpts.Version, pr.Identifier.ForDisplay(), outputFile)

	var checksum *FileChecksum
	for _, checksummer := range opts.Checksummers {
		// First check if checksum file is already here in the expected
		// download folder. Here we want to download a binary so we only check
		// for an existing checksum file from the folder we want to download
		// into.
		cs, err := checksummer.GetChecksumOfFile(outputFile)
		if err == nil && len(cs) > 0 {
			checksum = &FileChecksum{
				Expected:    cs,
				Checksummer: checksummer,
			}
			log.Printf("[TRACE] found a pre-exising %q checksum file", checksummer.Type)
			break
		}
	}

	if checksum == nil {
		log.Printf("[TRACE] no checksum file found locally, getting one")
		for _, getter := range getters {
			for _, checksummer := range opts.Checksummers {

				checksumFile, err := getter.Get(checksummer.Type, getOpts)
				if err != nil {
					err := fmt.Errorf("could not get checksum file for %s version %s. Is the file present on the release and correctly named ? %s", pr.Identifier.ForDisplay(), getOpts.Version, err)
					log.Printf("[TRACE] %s", err.Error())
					return nil, err
				}
				cs, err := checksummer.ParseChecksum(checksumFile)
				_ = checksumFile.Close()
				if err != nil {
					log.Printf("[TRACE] could not parse %s checksum: %v. Make sure the checksum file contains the checksum and only the checksum.", checksummer.Type, err)
					continue
				}

				if err := ioutil.WriteFile(outputFile+checksummer.FileExt(), []byte(cs.String()), 0666); err != nil {
					err := fmt.Errorf("Could not write checksum file %w", err)
					log.Printf("[TRACE] %s", err.Error())
					return nil, err
				}
				log.Printf("[TRACE] wrote %q file", outputFile+checksummer.FileExt())
				checksum = &FileChecksum{
					Expected:    cs,
					Checksummer: checksummer,
				}
			}
		}
	}

	if checksum == nil {
		return nil, fmt.Errorf("Could not find a valid checksum for %s.", outputFile)
	}

	// if outputFile is there and matches the checksum: do nothing
	if err := checksum.ChecksumFile(checksum.Expected, outputFile); err == nil {
		log.Printf("[TRACE] %s %s is already correctly installed in %q", pr.Identifier.ForDisplay(), getOpts.Version, outputFile)
		return nil, nil
	}

	for _, getter := range getters {
		// create temporary file that will receive a temporary binary
		tmpFile, err := tmp.File(getOpts.ExpectedFilename())
		if err != nil {
			return nil, fmt.Errorf("could not create temporary file to dowload plugin: %w", err)
		}

		// start fetching binary
		binary, err := getter.Get("binary", getOpts)
		if err != nil {
			err := fmt.Errorf("could not get binary for %s version %s. Is the file present on the release and correctly named ? %s", pr.Identifier.ForDisplay(), getOpts.Version, err)
			log.Printf("[TRACE] %v", err)
			continue
		}
		defer binary.Close()

		// write binary to tmp file
		_, err = io.Copy(tmpFile, binary)
		if err != nil {
			err := fmt.Errorf("Error getting plugin: %w", err)
			log.Printf("[TRACE] %v, trying another getter", err)
			continue
		}

		if _, err := tmpFile.Seek(0, 0); err != nil {
			err := fmt.Errorf("Error seeking begining of temporary file for checksumming: %w", err)
			log.Printf("[TRACE] %v, continuing", err)
			continue
		}

		// verify that the checksum for the file is what we expect.
		if err := checksum.Checksum(checksum.Expected, tmpFile); err != nil {
			err := fmt.Errorf("%w. Is the checksum file correct ? Is the binary file correct ?", err)
			log.Printf("%s", err)
			log.Printf("removing temporary plugin binary.")
			if err := os.Remove(tmpFile.Name()); err != nil {
				log.Printf("[TRACE] %v, continuing", err)
			}
			continue
		}

		if err := tmpFile.Close(); err != nil {
			err := fmt.Errorf("Failed to close tmp file %w", err)
			log.Printf("[TRACE] %v, continuing", err)
		}

		if err := os.Rename(tmpFile.Name(), outputFile); err != nil {
			err := fmt.Errorf("Failed to rename tmp file to correct location %w", err)
			return nil, err
		}

		// Success !!
		return &Installation{
			BinaryPath: outputFile,
			Version:    getOpts.Version,
		}, nil
	}

	return nil, nil
}
