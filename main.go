package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"sync"
)

const (
	destination  = "./"
	downloadPath = "https://community.opengroup.org/osdu/data/data-definitions/-/archive/master/data-definitions-master.tar.gz?path=Generated"
)

var includeFolders = [4]string{
	"master-data",
	"reference-data",
	"work-product",
	"work-product-component",
}

type language struct {
	Name          string
	FileExtension string
	FolderPath    string
	ExtraArgs     []string
}

var languages = []language{
	{Name: "go", FolderPath: "./", FileExtension: "go"},
}

func main() {
	os.Mkdir(destination, 0755)
	removeOldFolders()

	filename, err := downloadSchemas(downloadPath)
	check(err)
	defer os.RemoveAll(filename)

	untared, err := untar(filename)
	check(err)
	defer os.RemoveAll(untared)

	generatedFolder := untared + "/data-definitions-master-Generated/Generated/"

	var wg sync.WaitGroup

	for _, item := range includeFolders {
		wg.Add(1)
		go processFolder(generatedFolder, item, &wg)
	}

	wg.Wait()
}

func untar(tarFile string) (string, error) {
	file, err := os.Open(tarFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	tempDir, err := os.MkdirTemp("", "untarred-")
	if err != nil {
		return "", err
	}

	// Extract files from the tar archive
	for {
		header, err := tarReader.Next()

		switch {
		case err == io.EOF:
			return tempDir, nil

		case err != nil:
			return "", err

		case header == nil:
			continue

		}

		// Create the file or directory
		targetPath := filepath.Join(tempDir, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return "", err
			}
		case tar.TypeReg:
			outFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return "", err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return "", err
			}
			outFile.Close()
		}

	}

}

func downloadSchemas(url string) (string, error) {
	tempFile, err := os.CreateTemp("", "downloaded-*.tar.gz")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download file: %s", resp.Status)
	}

	if _, err = io.Copy(tempFile, resp.Body); err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func processFolder(source, item string, wg *sync.WaitGroup) {
	defer wg.Done()

	files, err := os.ReadDir(source + "/" + item)
	check(err)

	schemaFiles := make([]string, len(files))
	i := 0
	for _, v := range files {

		SchemaFile := source + "/" + item + "/" + v.Name()
		schemaFiles[i] = SchemaFile
		i++
	}

	cmds := make([]*exec.Cmd, len(languages))

	i = 0
	for _, lang := range languages {
		os.MkdirAll(lang.FolderPath+item, os.ModePerm)
		cmd := exec.Command("quicktype", "--package", strings.ReplaceAll(item, "-", ""), "-s", "schema", "-o", lang.FolderPath+item+"/"+item+"."+lang.FileExtension, "--lang", lang.Name)
		cmd.Args = append(cmd.Args, lang.ExtraArgs...)

		cmd.Args = append(cmd.Args, schemaFiles...)
		cmds[i] = cmd
		i++
	}

	var wg2 sync.WaitGroup
	for _, v := range cmds {
		wg2.Add(1)
		go func() {
			_, err = v.Output()
			check(err)
			defer wg2.Done()
		}()
		wg2.Wait()
	}

	check(err)
}

func removeOldFolders() {
	dir, err := os.ReadDir(destination)
	check(err)

	keys := make([]string, 0, len(languages))
	for _, k := range languages {
		keys = append(keys, k.Name)
	}

	for _, item := range dir {
		if !item.IsDir() || strings.HasPrefix(item.Name(), ".") {
			continue
		}

		if !slices.Contains(keys, item.Name()) {
			os.RemoveAll(item.Name())
		}

	}
}
