<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta charset="UTF-8">
    <title>{% block title %}Ayanerer{% endblock %}</title>
    <link rel="stylesheet" href={{ "/assets/stylesheets/pure-min.css" | suffixAssetsUpdate }} media="all">
    <link rel="stylesheet" href={{ "/assets/stylesheets/font-awesome.css" | suffixAssetsUpdate }} media="all">
    <link rel="stylesheet" href={{ "/assets/stylesheets/application.css" | suffixAssetsUpdate }} media="all">
  </head>
  <body>
    <div id="content">
      {% block content %}{% endblock %}
    </div>
  </body>
</html>
