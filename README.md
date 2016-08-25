# Badger

`Linux` `Windows`

A server that combines multiple continuous integration tools' statuses into a single badge with templating.

## Why use Badger

Some projects are built for different platforms using different continuous integration tools, each of them have their own build status badges that look different. This tool creates a URL to embed that generates a single status using each platform you specify.

## Sample Badge

The badge generated in the sample shows a simple build status check for the [ioRPC Project](https://github.com/ProjectLimitless/ioRPC) which uses AppVeyor Windows builds and Travis CI Linux builds.

Clicking on the badge will take you to the generated status page.

[![Badger](https://www.projectlimitless.io/badger/sample/badge)](https://www.projectlimitless.io/badger/sample)

## Getting Started

### Linux
    > git clone https://github.com/ProjectLimitless/Badger.git Limitless.Badger
    > cd Limitless.Badger/
    > make
    > ./bin/badger

Open [http://127.0.0.1:8000/sample](http://127.0.0.1:8000/sample) in your browser for the status page or [http://127.0.0.1:8000/sample/badge](http://127.0.0.1:8000/sample/badge) for the badge

### Windows
    > git clone https://github.com/ProjectLimitless/Badger.git Limitless.Badger
    > cd Limitless.Badger
    > Make.cmd
    > bin\badger.exe

_Note: The Windows 'Make.cmd' script **temporarily** changes your GOPATH to the vendor directory. It is reverted after compiling._

Open [http://127.0.0.1:8000/sample](http://127.0.0.1:8000/sample) in your browser for the status page or [http://127.0.0.1:8000/sample/badge](http://127.0.0.1:8000/sample/badge) for the badge

## Tests

The project makes use of Go's built-in testing.

TODO: Add tests

---
*A part of Project Limitless*

[![Project Limitless](https://www.donovansolms.com/downloads/projectlimitless.jpg)](https://www.projectlimitless.io)

[https://www.projectlimitless.io](https://www.projectlimitless.io)
