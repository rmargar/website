{{ define "title" }}{{ .Title }}{{ end }}

{{ define "head_extra" }}
<script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/highlight.min.js"></script>
<script>
    hljs.initHighlightingOnLoad();
</script>
<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/styles/atom-one-dark.min.css">
{{ end }}

{{ define "content" }}
{{ template "partial/heading.tpl" . }}


<div class="row medium-8 large-7 columns">
    <h1 class="blog-header"> {{ .Post.Title }}<small> {{ .Post.Added | format_date }}</small></h1>
    <div class="post-image-container"><img class="thumbnail" src="../static/assets/jpeg/rmargar.jpeg"></div>
    {{ template "partial/info.tpl" .Post }}

    {{ .ContentHTML | noescape }}

    <hr>
    {{ template "partial/disqus.tpl" . }}
</div>

{{ template "partial/footer.tpl" . }}

{{ end }}