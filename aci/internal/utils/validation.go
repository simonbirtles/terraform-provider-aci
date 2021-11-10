package utils

import ( 
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"regexp"
	"strings"
	"fmt"
)

func ValidateList(options []string) schema.SchemaValidateFunc {
		s := strings.Join(options, "|")
		r := fmt.Sprintf("^%s$", s)
		return validation.StringMatch(regexp.MustCompile(r), fmt.Sprintf("Valid options are: %s", s) )
}
