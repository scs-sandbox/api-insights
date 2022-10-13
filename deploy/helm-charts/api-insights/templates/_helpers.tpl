{{/* vim: set filetype=mustache: */}}
{{/*
Name of the chart.
*/}}
{{- define "api-insights.name" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name -}}
{{- end -}}

{{/*
Helm labels.
*/}}
{{- define "api-insights.labels" -}}
    app.kubernetes.io/name: {{ include "api-insights.name" . }}
    app.kubernetes.io/managed-by: {{ .Release.Service }}
    app.kubernetes.io/instance: {{ .Release.Name }}
    helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version }}
{{- end -}}
