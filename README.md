<p align="center">
	<img src="./assets/icon.png" width="120px" />
</p>
 
<p align="center">
	a zero dependency performant validation library
</p>

<p align="center">
	<a href="https://opensource.org/licenses/MIT" target="_blank" alt="License">
		<img src="https://img.shields.io/badge/License-MIT-blue.svg" />
	</a>
	<a href="https://pkg.go.dev/github.com/aacebo/owl" target="_blank" alt="Go Reference">
		<img src="https://pkg.go.dev/badge/github.com/aacebo/owl.svg" />
	</a>
	<a href="https://goreportcard.com/report/github.com/aacebo/owl" target="_blank" alt="Go Report Card">
		<img src="https://goreportcard.com/badge/github.com/aacebo/owl" />
	</a>
	<a href="https://github.com/aacebo/owl/actions/workflows/ci.yml" target="_blank" alt="Build">
		<img src="https://github.com/aacebo/owl/actions/workflows/ci.yml/badge.svg?branch=main" />
	</a>
	<a href="https://codecov.io/gh/owl/jsonschema" target="_blank" alt="codecov">
		<img src="https://codecov.io/gh/owl/jsonschema/graph/badge.svg?token=ZFJMM1BZVM" />
	</a>
</p>

## Rules

| Name			| Description											| Status |
|---------------|-------------------------------------------------------|--------|
| required		| not nil or zero value									| ✅		|
| default		| default value											| ⌛		|
| enum			| one of a set of options								| ⌛		|

### String

| Name			| Description											| Status |
|---------------|-------------------------------------------------------|--------|
| pattern		| match regular expression								| ✅		|
| format		| match format											| ✅		|
| min			| min length											| ✅		|
| max			| max length											| ✅		|

### Numeric

| Name			| Description											| Status |
|---------------|-------------------------------------------------------|--------|
| min			| minimum												| ✅		|
| max			| maximum												| ✅		|

## Formats

| Name			| Status |
|---------------|--------|
| date_time		| ✅		|
| email			| ✅		|
| ipv4			| ✅		|
| ipv6			| ✅		|
| uri			| ✅		|
| uuid			| ✅		|
