module github.com/rollbar/terraform-provider-rollbar

go 1.15

// Until https://github.com/rs/zerolog/pull/266 or https://github.com/rs/zerolog/pull/267
// is included in the next release
replace github.com/rs/zerolog => github.com/jmcvetta/zerolog v1.20.1-0.20201102133610-4cc56b8f3f5a

require (
	github.com/brianvoe/gofakeit/v5 v5.11.2
	github.com/dnaeon/go-vcr v1.1.0
	github.com/go-resty/resty/v2 v2.4.0
	github.com/jarcoal/httpmock v1.0.8
	github.com/rs/zerolog v1.20.0
	github.com/stretchr/testify v1.7.0
)
