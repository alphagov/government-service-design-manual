---
layout: browse
title: Browse guides by topic
subtitle: A list of all guides
category: home
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
---

<div class="article-container group">
  <div class="contents">
    <nav>
      <h2>Contents</h2>

      <ol>
        <li><a href="#agile">Agile</a></li>
        <li><a href="#assisted-digital">Assisted digital</a></li>
        <li><a href="#content-and-design">Content and design</a></li>
        <li><a href="#making-software">Making software</a></li>
        <li><a href="#measurement">Measurement</a></li>
        <li><a href="#operations">Operations</a></li>
        <li><a href="#phases">Phases</a></li>
        <li><a href="#the-team">The team</a></li>
        <li><a href="#users">Users</a></li>
      </ol>
    </nav>
  </div>

  <div class="inner">
    <div class="link-list">
      <h2 id="agile">Agile</h2>
      <ul>
      {% sorted_for p in site.pages sort_by:title %}
        {% if p.category == 'agile' %}
        {% if p.type == 'guide' %}
        <li> 
            <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
            <span class="subtitle">{{ p.subtitle }}</span>
        </li>
        {% endif %}
        {% endif %}
      {% endsorted_for %}
      </ul>

      <h2 id="assisted-digital">Assisted digital</h2>
      <ul>
      {% sorted_for p in site.pages sort_by:title %}
        {% if p.category == 'assisted-digital' %}
        {% if p.type == 'guide' %}
        <li> 
            <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
            <span class="subtitle">{{ p.subtitle }}</span>
        </li>
        {% endif %}
        {% endif %}
      {% endsorted_for %}
      </ul>

      <h2 id="content-and-design">Content and design</h2>
      <ul>
      {% sorted_for p in site.pages sort_by:title %}
        {% if p.category == 'content-and-design' %}
        {% if p.type == 'guide' %}
        <li> 
            <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
            <span class="subtitle">{{ p.subtitle }}</span>
        </li>
        {% endif %}
        {% endif %}
      {% endsorted_for %}
      </ul>

      <h2 id="making-software">Making software</h2>
      <ul>
      {% sorted_for p in site.pages sort_by:title %}
        {% if p.category == 'making-software' %}
        {% if p.type == 'guide' %}
        <li> 
            <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
            <span class="subtitle">{{ p.subtitle }}</span>
        </li>
        {% endif %}
        {% endif %}
      {% endsorted_for %}
      </ul>

      <h2 id="measurement">Measurement</h2>
      <ul>
      {% sorted_for p in site.pages sort_by:title %}
        {% if p.category == 'measurement' %}
        {% if p.type == 'guide' %}
        <li> 
            <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
            <span class="subtitle">{{ p.subtitle }}</span>
        </li>
        {% endif %}
        {% endif %}
      {% endsorted_for %}
      </ul>

      <h2 id="operations">Operations</h2>
      <ul>
      {% sorted_for p in site.pages sort_by:title %}
        {% if p.category == 'operations' %}
        {% if p.type == 'guide' %}
        <li> 
            <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
            <span class="subtitle">{{ p.subtitle }}</span>
        </li>
        {% endif %}
        {% endif %}
      {% endsorted_for %}
      </ul>

      <h2 id="phases">Phases</h2>
      <ul>
      {% sorted_for p in site.pages sort_by:title %}
        {% if p.category == 'phases' %}
        {% if p.type == 'guide' %}
        <li> 
            <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
            <span class="subtitle">{{ p.subtitle }}</span>
        </li>
        {% endif %}
        {% endif %}
      {% endsorted_for %}
      </ul>

      <h2 id="the-team">The team</h2>
      <ul>
      {% sorted_for p in site.pages sort_by:title %}
        {% if p.category == 'the-team' %}
        {% if p.type == 'guide' %}
        <li> 
            <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
            <span class="subtitle">{{ p.subtitle }}</span>
        </li>
        {% endif %}
        {% endif %}
      {% endsorted_for %}
      </ul>

      <h2 id="users">Users</h2>
      <ul>
      {% sorted_for p in site.pages sort_by:title %}
        {% if p.category == 'users' %}
        {% if p.type == 'guide' %}
        <li> 
            <a href="{{ p.url }}" title="{{ p.subtitle }}">{{ p.title }}</a>
            <span class="subtitle">{{ p.subtitle }}</span>
        </li>
        {% endif %}
        {% endif %}
      {% endsorted_for %}
      </ul>
    </div>
  </div>
</div>