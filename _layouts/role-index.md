---
layout: gsdm
section: home
---

<div class="home-page-promos">
  <div class="hero topic">

  {% for p in site.pages %}
    {% if p.title == page.hero %}
    <h3>{{ p.title }}</h3>
    <p><a href="{{ p.url }}">{{ p.subtitle }}</a></p>
    {% endif %}
  {% endfor %}

  </div>
  <div class="topic">
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
  <div class="topic">
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
</div>


<div class="topic">
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

<p><strong><a href="/all-guides">See all guides...</a></strong></p>
