---
layout: wide
section: home
---

{% assign resources = false %}
{% assign width = ' wide' %}
{% for p in site.pages %}{% if p.audience.primary contains page.audience and p.type == 'resource' %}
{% assign resources = true %}
{% assign width = '' %}
{% endif %}{% endfor %}

<div class="primary-links">
  <div class="hero link-list">

  {% for p in site.pages %}
    {% if p.title == page.hero %}
    <a href="{{ p.url }}">
      <h3>{{ p.title }}</h3>
      <p>{{ p.subtitle }}</p>
    </a>
    {% endif %}
  {% endfor %}

  </div>
  <div class="link-list{{ width }}">
    <h3>Guides</h3>
    <ul>
    {% for p in site.pages %}
      {% if p.audience.primary contains page.audience and p.type == 'guide' and p.title != page.hero %}
      <li> 
          <a href="{{ p.url }}">{{ p.title }}</a>
      </li>
      {% endif %}
    {% endfor %}
    </ul>
  </div>

{% if resources %}
  <div class="link-list">
    <h3>Resources</h3>
    <ul>
    {% for p in site.pages %}
      {% if p.audience.primary contains page.audience and p.type == 'resource' %}
      <li> 
          <a href="{{ p.url }}">{{ p.title }}</a>
      </li>
      {% endif %}
    {% endfor %}
    </ul>
  </div>
{% endif %}

</div>


<div class="link-list">
  <h3>More guides from the manual</h3>
<ul>
{% for p in site.pages %}
  {% if p.audience.secondary contains page.audience %}
  <li> 
      <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
  </li>
  {% endif %}
{% endfor %}
</ul>
</div>

{% include _role-links.html %}

<p><strong><a href="/all-guides">See all guides...</a></strong></p>
