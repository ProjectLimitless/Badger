#!/bin/sh
# This script remaps the coverage paths to the actual Travis CI path
sed -i 's/_\/home\/donovan\/Development\/project-limitless\/Limitless.Badger\/src\/badger/badger/g' coverage.out
sed -i 's/_\/home\/travis\/gopath\/src\/github.com\/ProjectLimitless\/Badger\/src/github.com\/ProjectLimitless\/Badger\/src\//g' coverage.out
