#!/bin/sh
sed -i 's/_\/home\/donovan\/Development\/project-limitless\/Limitless.Badger\/src\/badger/badger/g' coverage.out
sed -i 's/_\/home\/travis\/gopath\/src\/github.com\/ProjectLimitless\/Badger\/src\/badger/badger/g' coverage.out
