package main

import (
	"fmt"
	"os"

	"github.com/darkcrux/go-bintray/bintray"
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	defaultSubject    = os.Getenv("BINTRAY_SUBJECT")
	defaultApiKey     = os.Getenv("BINTRAY_API_KEY")
	defaultRepository = os.Getenv("BINTRAY_REPOSITORY")
	defaultPackage    = os.Getenv("BINTRAY_PACKAGE")

	bintrayCmd = kingpin.New("bintray", "A command line interface to bintray.")

	subject = bintrayCmd.Flag("subject", "the username").Default(defaultSubject).String()
	apiKey  = bintrayCmd.Flag("api-key", "the api key").Default(defaultApiKey).String()
	repo    = bintrayCmd.Flag("repository", "the repository").Default(defaultRepository).String()
	pkg     = bintrayCmd.Flag("package", "the package").Default(defaultPackage).String()

	pkgExists = bintrayCmd.Command("package-exists", "check if package exists")

	listVersions = bintrayCmd.Command("list-versions", "list all versions of a package")

	createVersion         = bintrayCmd.Command("create-version", "create a new version")
	createrVersionVersion = createVersion.Arg("version", "version name").Required().String()

	uploadFile          = bintrayCmd.Command("upload-file", "upload a file to a version")
	uploadFileVersion   = uploadFile.Arg("version", "the version").Required().String()
	uploadFilePath      = uploadFile.Arg("file", "the file to upload").Required().String()
	uploadFilePrjId     = uploadFile.Flag("project-id", "the Project ID").String()
	uploadFilePrjName   = uploadFile.Flag("project-name", "the Project Name").String()
	uploadFileMavenRepo = uploadFile.Flag("maven-repo", "the Maven Repository").Bool()

	publish        = bintrayCmd.Command("publish", "publish version")
	publishVersion = publish.Arg("version", "the version").String()
)

type bintrayClient interface {
	PackageExists(subject, repo, pkg string) (bool, error)
	GetVersions(subject, repo, pkg string) ([]string, error)
	CreateVersion(subject, repo, pkg, version string) error
	UploadFile(subject, repo, pkg, version, prjId, prjName, filePath string, mavenRepo bool) error
	Publish(subject, repo, pkg, version string) error
}

type bintrayPackage struct {
	subject string
	repo    string
	pkg     string
	bintrayClient
}

func (bp *bintrayPackage) exists() {
	exists, err := bp.PackageExists(bp.subject, bp.repo, bp.pkg)
	if err != nil {
		handleError(err)
	}
	fmt.Println(exists)
}

func (bp *bintrayPackage) listVersions() {
	versions, err := bp.GetVersions(bp.subject, bp.repo, bp.pkg)
	if err != nil {
		handleError(err)
	}
	fmt.Printf("%s/%s:\n", bp.repo, bp.pkg)
	for _, version := range versions {
		fmt.Printf(" - %s\n", version)
	}
}

func (bp *bintrayPackage) createVersion(version string) {
	err := bp.CreateVersion(bp.subject, bp.repo, bp.pkg, version)
	if err != nil {
		handleError(err)
	}
	fmt.Printf("%s/%s %s created.\n", bp.repo, bp.pkg, version)
}

func (bp *bintrayPackage) uploadFile(version, prjId, prjName, filePath string, mavenRepo bool) {
	err := bp.UploadFile(bp.subject, bp.repo, bp.pkg, version, prjId, prjName, filePath, mavenRepo)
	if err != nil {
		handleError(err)
	}
	fmt.Printf("Successfully uploaded %s to version %s of %s in %s.\n", filePath, version, bp.pkg, bp.repo)
}

func (bp *bintrayPackage) publish(version string) {
	err := bp.Publish(bp.subject, bp.repo, bp.pkg, version)
	if err != nil {
		handleError(err)
	}
	fmt.Println("Successfully published %s/%s %s.\n", bp.repo, bp.pkg, version)

}

func main() {
	bintrayCmd.Version("0.1.0")
	command := kingpin.MustParse(bintrayCmd.Parse(os.Args[1:]))

	btpkg := &bintrayPackage{
		subject:       *subject,
		repo:          *repo,
		pkg:           *pkg,
		bintrayClient: bintray.NewClient(nil, *subject, *apiKey),
	}

	switch command {
	case pkgExists.FullCommand():
		btpkg.exists()
	case listVersions.FullCommand():
		btpkg.listVersions()
	case createVersion.FullCommand():
		btpkg.createVersion(*createrVersionVersion)
	case uploadFile.FullCommand():
		btpkg.uploadFile(*uploadFileVersion, *uploadFilePrjId, *uploadFilePrjName, *uploadFilePath, *uploadFileMavenRepo)
	case publish.FullCommand():
		btpkg.publish(*publishVersion)
	}
}

func handleError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
