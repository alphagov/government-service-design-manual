---
layout: wide
title: All guides
subtitle: A list of all guides
section: home
status: draft
---

<div class="link-list">
  <h3>All guides</h3>
<ul>
{% for p in site.pages %}
  {% if p.section == 'guidance' %}
  <li> 
      <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
  </li>
  {% endif %}
{% endfor %}
</ul>