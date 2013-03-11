---
layout: wide
title: Browse guides by topic
subtitle: A list of all guides
category: home
status: draft
---

<div class="link-list">
  <h2>Agile</h2>
  <ul>
  {% sorted_for p in site.pages sort_by:title %}
    {% if p.category == 'agile' %}
    {% if p.type == 'guide' %}
    <li> 
        <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
    </li>
    {% endif %}
    {% endif %}
  {% endsorted_for %}
  </ul>

  <h2>Assisted digital</h2>
  <ul>
  {% sorted_for p in site.pages sort_by:title %}
    {% if p.category == 'assisted-digital' %}
    {% if p.type == 'guide' %}
    <li> 
        <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
    </li>
    {% endif %}
    {% endif %}
  {% endsorted_for %}
  </ul>

  <h2>Content and design</h2>
  <ul>
  {% sorted_for p in site.pages sort_by:title %}
    {% if p.category == 'content-and-design' %}
    {% if p.type == 'guide' %}
    <li> 
        <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
    </li>
    {% endif %}
    {% endif %}
  {% endsorted_for %}
  </ul>

  <h2>Making software</h2>
  <ul>
  {% sorted_for p in site.pages sort_by:title %}
    {% if p.category == 'making-software' %}
    {% if p.type == 'guide' %}
    <li> 
        <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
    </li>
    {% endif %}
    {% endif %}
  {% endsorted_for %}
  </ul>

  <h2>Measurement</h2>
  <ul>
  {% sorted_for p in site.pages sort_by:title %}
    {% if p.category == 'measurement' %}
    {% if p.type == 'guide' %}
    <li> 
        <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
    </li>
    {% endif %}
    {% endif %}
  {% endsorted_for %}
  </ul>

  <h2>Operations</h2>
  <ul>
  {% sorted_for p in site.pages sort_by:title %}
    {% if p.category == 'operations' %}
    {% if p.type == 'guide' %}
    <li> 
        <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
    </li>
    {% endif %}
    {% endif %}
  {% endsorted_for %}
  </ul>

  <h2>Phases</h2>
  <ul>
  {% sorted_for p in site.pages sort_by:title %}
    {% if p.category == 'phases' %}
    {% if p.type == 'guide' %}
    <li> 
        <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
    </li>
    {% endif %}
    {% endif %}
  {% endsorted_for %}
  </ul>

  <h2>The Team</h2>
  <ul>
  {% sorted_for p in site.pages sort_by:title %}
    {% if p.category == 'the-team' %}
    {% if p.type == 'guide' %}
    <li> 
        <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
    </li>
    {% endif %}
    {% endif %}
  {% endsorted_for %}
  </ul>

  <h2>Users</h2>
  <ul>
  {% sorted_for p in site.pages sort_by:title %}
    {% if p.category == 'users' %}
    {% if p.type == 'guide' %}
    <li> 
        <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
    </li>
    {% endif %}
    {% endif %}
  {% endsorted_for %}
  </ul>

</div>