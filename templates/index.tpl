{{ define "content" }}

    {{ template "partial/heading.tpl" . }}
    {{ template "partial/posts.tpl" .}}
    {{ template "partial/footer.tpl" .}}

{{ end }}
