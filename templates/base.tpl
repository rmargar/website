<!DOCTYPE html>
<html>

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.">
  <title>Ricardo Martinez's blog</title>

  <link rel="stylesheet" href="/static/stylesheets/blog.css" type="text/css" />
  <link rel="stylesheet" href="/static/stylesheets/main.css" type="text/css" />
  <link
    href="https://fonts.googleapis.com/css2?family=Montserrat&family=Space+Grotesk:wght@300;400;500;600;700&family=Space+Mono:wght@400;700&display=swap"
    rel="stylesheet">

  <link rel="stylesheet" href="https://dhbhdrzi4tiry.cloudfront.net/cdn/sites/foundation.min.css">
  <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
  {{ block "head_extra" . }}{{ end }}
</head>

<body>

  <!-- Start Top Bar -->
  <div class="top-bar blog-top-bar">
    <ul class="menu">
      <li><a href="/blog">Blog</a></li>
      <li><a href="/">Home</a></li>
    </ul>
  </div>
  <!-- End Top Bar -->

  {{ block "content" . }}{{ end }}


  <script src="https://code.jquery.com/jquery-2.1.4.min.js"></script>
  <script src="https://dhbhdrzi4tiry.cloudfront.net/cdn/sites/foundation.js"></script>
  <script>
    $(document).foundation();
  </script>
  <script type="module" src="https://md-block.verou.me/md-block.js"></script>
</body>

</html>