package global

{{- if .HasGlobal }}

import "slot-server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}