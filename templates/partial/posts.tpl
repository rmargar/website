{{ define "partial/posts.tpl" }}
<div class="row medium-8 large-7 columns">
    {{ range .Posts }}
    <h2 class="blog-header"> <a href="/blog/{{ .URLPath }}"> {{ .Title }} </a> <small class="blog-header">
            {{ .Added | format_date }}</small></h2>
    <div class="blog-post">
        <div class="clearfix">
            {{ template "partial/info.tpl" . }}
        </div>
        <div class="post-image-container"><img src="/static/assets/jpeg/rmargar.jpeg"></div>
    </div>

    <p class="post-summary">
        {{- if .Summary -}}
        {{ .Summary | striptags }} <a href="blog/{{ .URLPath }}">[Read more]</a>
        {{- else -}}
        {{ .Content | striptags | truncatechars 255 }} <a href="blog/{{ .URLPath }}">[Read more]</a>
        {{- end -}}
    </p>
    <hr>

    {{ end }}
    <script id="dsq-count-scr" src="https://rmargar-net.disqus.com/count.js" async></script>
</div>

{{ end }}