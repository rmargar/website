{{ define "partial/info.tpl" }}
<div class="clearfix">
    <p>
        {{ if gt (len .Tags) 0 }}
        on
        {{ range .Tags }}
        <a href="/blog/tag/{{ . }}" class="label secondary">{{ . }}</a>
        {{ end }}
        <a href="/blog/{{ .URLPath }}#disqus_thread" data-disqus-identifier="post-{{ .ID }}"> Comments</a>
    {{ end }}
    </p>
</div>
{{ end }}