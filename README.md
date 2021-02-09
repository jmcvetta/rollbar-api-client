Rollbar Management API Client
=============================

A client for the Rollbar API, providing access to project & team management
functionality. Used internally by the Rollbar Terraform provider.  Does **not**
provide error reporting functionality. If you want to use Rollbar to collect
errors from your application, this is the wrong client.

See [`rollbar-go`](https://github.com/rollbar/rollbar-go) for the official Go
client for Rollbar, which *does* provide support for error reporting.


Status
------

[![Tests](https://github.com/jmcvetta/rollbar-api-client/workflows/Tests/badge.svg)](https://github.com/jmcvetta/rollbar-api-client/actions)
[![Coverage Status](https://coveralls.io/repos/github/jmcvetta/rollbar-api-client/badge.svg)](https://coveralls.io/github/rollbar/terraform-provider-rollbar)
[![CodeQL](https://github.com/jmcvetta/rollbar-api-client/workflows/CodeQL/badge.svg)](https://github.com/jmcvetta/rollbar-api-client/actions?query=workflow%3ACodeQL)
[![ShiftLeft Scan](https://github.com/jmcvetta/rollbar-api-client/workflows/ShiftLeft%20Scan/badge.svg)](https://github.com/jmcvetta/rollbar-api-client/actions?query=workflow%3A%22ShiftLeft+Scan%22)
[![Maintainability](https://api.codeclimate.com/v1/badges/c5097d1a11f6f2310089/maintainability)](https://codeclimate.com/github.com/jmcvetta/rollbar-api-client/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/c5097d1a11f6f2310089/test_coverage)](https://codeclimate.com/github/jmcvetta/rollbar-api-client/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/jmcvetta/rollbar-api-client)](https://goreportcard.com/report/github.com/jmcvetta/rollbar-api-client)


Documentation
-------------

[![Documentation](https://godoc.org/github.com/jmcvetta/rollbar-api-client?status.svg)](http://godoc.org/github.com/jmcvetta/rollbar-api-client)
