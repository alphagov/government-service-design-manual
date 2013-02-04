---
layout: gsdm
title: Building a team
subtitle: Who to hire, your working environment and what your team should look like
categories : [guidance, getting-started]
status: draft
---
    
<div class="topic">
  <h3>{{ page.title }}</h3>
  <ul>
{% for post in site.categories[page.category] %}
    <li>
      <a href="{{ post.url }}">
        <span class="title">{{ post.title }}</span>
        <span class="description">{{ post.subtitle }}</span>
        <!-- Filed under {{ post.categories | category_links }} -->
      </a>
    </li>
{% endfor %}
  </ul>
</div>
