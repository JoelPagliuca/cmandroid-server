# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
* More documentation

### Fixed
* Test coverage

## [1.0.0] - 2018-10-25
### Added
* API endpoints
	* `GET /listdevices`
	* `POST /tap`
* CLI interface
	* `-adb` Path to adb (default "adb")
	* `-h` Display help text
	* `-p` Port to run the server on (default 8080)