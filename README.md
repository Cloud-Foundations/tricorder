# Library for publishing metrics which can be pulled or viewed via HTTP
[![Build Status](https://travis-ci.org/Cloud-Foundations/tricorder.svg?branch=master)](https://travis-ci.org/Cloud-Foundations/tricorder)
[![Coverage Status](https://coveralls.io/repos/github/Cloud-Foundations/tricorder/badge.svg?branch=master)](https://coveralls.io/github/Cloud-Foundations/tricorder?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Cloud-Foundations/tricorder)](https://goreportcard.com/report/github.com/Cloud-Foundations/tricorder)

This is a library which can be used by applications to publish metrics.
The metrics are made available via HTTP, using the same port as the webserver
for the application. There is a browsable, human-friendly interface as well
as machine-friendly interfaces.

Please see the
[design document](https://docs.google.com/document/d/142Llj30LplgxWhOLOprqH59hS01EJ9iC1THV3no5oy0/pub)
and the
[public Go API documentation]
(https://godoc.org/github.com/Cloud-Foundations/tricorder/go/tricorder)
for more information.

## Contributions

All contributions must be unencumbered. It is the responsibility of
the contributor to ensure compliance with all laws, copyrights,
patents and contracts.

## LICENSE

Copyright 2015 Symantec Corporation.
Copyright 2019 cloud-foundations.org

Licensed under the Apache License, Version 2.0 (the “License”); you
may not use this file except in compliance with the License.

You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0 Unless required by
applicable law or agreed to in writing, software distributed under the
License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for
the specific language governing permissions and limitations under the
License.
