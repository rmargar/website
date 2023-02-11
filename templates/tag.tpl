{{ define "title" }} Posts on "{{ .Tag }}"{{ end }}

{{ define "content" }}
    {{ template "partial/heading.tpl" . }}
    <div class="row medium-8 large-7 columns" ><h2 class="blog-header">Posts on <span class="label secondary" style="font-size: inherit;">{{ .Tag }}</span></h2></div>
    
    {{ if lt (len .Posts)  1 }}
    <div class="row medium-8 large-7 columns" ><h3 class="blog-header">Nothing here yet :(</h2></div>
    {{ else }}
        {{ template "partial/posts.tpl" .}}
    {{ end }}
    {{ template "partial/footer.tpl" .}}
{{ end }}
