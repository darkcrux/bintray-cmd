# bintray-cmd
Command line version for go-bintray

[![Build Status](https://travis-ci.org/darkcrux/bintray-cmd.png)](https://travis-ci.org/darkcrux/bintray-cmd)

#### Binary Release

Latest release:
- [darwin-amd64](https://bintray.com/artifact/download/darkcrux/generic/bintray-latest-darwin-amd64.tar.gz)
- [linux-386](https://bintray.com/artifact/download/darkcrux/generic/bintray-latest-linux-386.tar.gz)
- [linux-amd64](https://bintray.com/artifact/download/darkcrux/generic/bintray-latest-linux-amd64.tar.gz)

#### Bintray Credentials

Subject, API Key, Repository, and Package name can be supplied using environment variables or by the command flags:

| Flags       | env                 |
|-------------|---------------------|
|--subject    | $BINTRAY_SUBJECT    |
|--api-key    | $BINTRAY_API_KEY    |
|--repository | $BINTRAY_REPOSITORY |
|--package    | $BINTRAY_PACKAGE    |

#### Available Commands

- `bintray package-exists`
- `bintray list-versions`
- `bintray create-version`
- `bintray upload-file`
- `bintray publish`

#### More Info

`bintray --help`