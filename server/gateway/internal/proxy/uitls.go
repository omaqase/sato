package proxy

import "strings"

func GetRequiredRoles(path string) []string {
	if strings.HasPrefix(path, "/admin") {
		return []string{"admin", "super_admin"}
	}

	if strings.HasPrefix(path, "/seller") {
		return []string{"seller", "admin", "super-admin"}
	}

	return nil
}
