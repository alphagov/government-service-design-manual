---
layout: gsdm
title: Service managers
subtitle: Everything you need to create great services for GOV.UK
section: home
status: draft
hero: The Digital by Default Standard
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
      {% if p.audience.primary contains 'service-manager' and p.type == 'guide' and p.title != page.hero %}
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
      {% if p.audience.primary contains 'service-manager' and p.type == 'resource' %}
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
  {% if p.audience.secondary contains 'service-manager' %}
  <li> 
      <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
  </li>
  {% endif %}
{% endfor %}
</ul>
</div>

**[See all guides...](/all-guides)**

<!--
<div class="home-page-promos">
  
  <div class="topic">
    <h3>Our design principles</h3>
    <ul><li><a href="guides-and-toolkits/design-principles">Everything we make is founded on these ten principles</a></li></ul>
  </div>

  <div class="topic">
    <h3>Service development</h3>
    <p><a href="guides-and-toolkits/design-principles">The four phases</a></p>
  </div>

  <div class="topic">
    <h3>Building your team</h3>
    <p><a href="guides-and-toolkits/design-principles">Who to hire and your working environment</a></p>
  </div>

    <div class="topic">
    <h3>Working in an agile way</h3>
    <p><a href="guides-and-toolkits/design-principles">What it is, why it works and how to do it</a></p>
  </div>

</div>
-->