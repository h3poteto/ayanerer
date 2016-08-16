{% extends "layouts/public.html.tpl" %}

{% block title %}{{ title }}{% endblock %}

{% block content %}

<div class="image-board">
  <article>
    <div class="title">
      <h3>Sign In</h3>
    </div>
    <div class="content">
      <div class="sign-in-board">
        <span class="pure-button pure-button-primary session-button"><i class="fa fa-twitter" aria-hidden="true"></i>
 Sign In with Twitter</span>
      </div>
    </div>
  </article>
</div>

{% endblock %}
