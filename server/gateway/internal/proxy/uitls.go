package proxy

import "strings"

func GetRequiredRoles(path string) []string {
	if strings.HasPrefix(path, "/api/") {
		return []string{"user", "admin", "super_admin", "seller"}
	}

	if strings.HasPrefix(path, "/admin") {
		return []string{"admin", "super_admin", "seller"}
	}

	if strings.HasPrefix(path, "/seller") {
		return []string{"seller", "admin", "super_admin"}
	}

	return nil
}
