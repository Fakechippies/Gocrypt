package baseconversion

import (
	"github.com/liut/baseconv"
)

func Parser(fromMethod string, toMethod string, input string) string {
	if fromMethod == "base2" && toMethod == "base10" {
		return base10conv(input, fromMethod)
	} else if fromMethod == "base2" && toMethod == "base16" {
		return base16conv(input, fromMethod)
	} else if fromMethod == "base10" && toMethod == "base2" {
		return base2conv(input, fromMethod)
	} else if fromMethod == "base10" && toMethod == "base16" {
		return base16conv(input, fromMethod)
	} else if fromMethod == "base16" && toMethod == "base2" {
		return base2conv(input, fromMethod)
	} else if fromMethod == "base16" && toMethod == "base10" {
		return base10conv(input, fromMethod)
	}
	return "Error in encodings provided"
}

func base10conv(input string, from string) string {
	result := "Error in base encoding provided"
	if from == "base2" {
		result_ini, err := baseconv.Convert(input, 2, 10)
		if err != nil {
			return "Error in input provided"
		}
		result = result_ini
	} else if from == "base16" {
		result_ini, err := baseconv.Convert(input, 16, 10)
		if err != nil {
			return "Error in input provided"
		}
		result = result_ini
	}
	return result
}

func base2conv(input string, from string) string {
	result := "Error in base encodings provided"
	if from == "base10" {
		result_ini, err := baseconv.Convert(input, 10, 2)
		if err != nil {
			return "Error in input provided"
		}
		result = result_ini
	} else if from == "base16" {
		result_ini, err := baseconv.Convert(input, 16, 2)
		if err != nil {
			return "Error in input provided"
		}
		result = result_ini
	}
	return result
}

func base16conv(input string, from string) string {
	result := "Error in base encodings provided"
	if from == "base2" {
		result_ini, err := baseconv.Convert(input, 2, 16)
		if err != nil {
			return "Error in input provided"
		}
		result = result_ini
	} else if from == "base10" {
		result_ini, err := baseconv.Convert(input, 10, 16)
		if err != nil {
			return "Error in input provided"
		}
		result = result_ini
	}
	return result
}
