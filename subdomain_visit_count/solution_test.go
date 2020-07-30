package subdomain_visit_count

import (
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	input := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	t.Logf("%v", subdomainVisits(input))
}