#!/bin/bash

go test ./*day*/*p* | grep -v '\[no test files\]'