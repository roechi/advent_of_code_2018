module day02/app

go 1.15

replace github.com/roechi/utils => ../../utils

replace github.com/roechi/idscanner => ../idscanner

require (
	github.com/roechi/idscanner v0.0.0-00010101000000-000000000000
	github.com/roechi/utils v0.0.0-00010101000000-000000000000
)
