{{- define "operator.name" -}}
{{- printf "%s-operator" (include "common.names.fullname" .) | trunc 63 | trimSuffix "-"  -}}
{{- end -}}
