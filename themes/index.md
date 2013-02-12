---
layout: gsdm
section: home
title: Guides by theme
---


<div class="topic">
  <h3>Getting started</h3>
<ul>
{% for p in site.pages %}
  {% if p.theme == 'getting-started' %}
  <li> 
      <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
  </li>
  {% endif %}
{% endfor %}
</ul>
</div>

<div class="topic">
  <h3>Understanding your audience</h3>
<ul>
{% for p in site.pages %}
  {% if p.theme == 'understanding-your-audience' %}
  <li> 
      <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
  </li>
  {% endif %}
{% endfor %}
</ul>
</div>

<div class="topic">
  <h3>Security and the law</h3>
<ul>
{% for p in site.pages %}
  {% if p.theme == 'security-and-the-law' %}
  <li> 
      <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
  </li>
  {% endif %}
{% endfor %}
</ul>
</div>

<div class="topic">
  <h3>Designing and building your service</h3>
<ul>
{% for p in site.pages %}
  {% if p.theme == 'design-and-build' %}
  <li> 
      <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
  </li>
  {% endif %}
{% endfor %}
</ul>
</div>

<div class="topic">
  <h3>Maintaining and improving your service</h3>
<ul>
{% for p in site.pages %}
  {% if p.theme == 'maintain-and-improve' %}
  <li> 
      <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
  </li>
  {% endif %}
{% endfor %}
</ul>
</div>