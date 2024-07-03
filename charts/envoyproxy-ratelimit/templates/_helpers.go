{{/*
Expand the name of the chart.
*/}}
{{- define "ratelimit.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Format pdb.
*/}}
{{- define "pdb.formatValue" }}
{{- $valueToFormat := . | toString }}
{{- if contains "%" $valueToFormat }}
{{- print $valueToFormat | quote }}
{{- else }}
{{- print $valueToFormat }}
{{- end }}
{{- end }}


{{/*
Kubernetes version
*/}}
{{- define "ratelimit.capabilities.kubeVersion" -}}
{{- default .Capabilities.KubeVersion.Version .Values.kubeVersion -}}
{{- end -}}
