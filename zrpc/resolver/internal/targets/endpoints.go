/**
 * @Author:      leafney
 * @Date:        2022-06-08 16:26
 * @Project:     go-zero
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package targets

import (
	"google.golang.org/grpc/resolver"
	"strings"
)

const slashSeparator = "/"

// GetAuthority returns the authority of the target.
func GetAuthority(target resolver.Target) string {
	return target.URL.Host
}

// GetEndpoints returns the endpoints from the given target.
func GetEndpoints(target resolver.Target) string {
	return strings.Trim(target.URL.Path, slashSeparator)
}
