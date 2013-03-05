---
layout: design-patterns
title: Buttons
subtitle: How to word them, how to code them, when to use them   
section: guidance
subsection: Design patterns
type: resource
status: draft
page_class: buttons
---

Buttons are used to signify actions that the user can perform. If it's not an action it should probably be a link. Here's how to create buttons in the GOV.UK style.

[The Sass mixin for making these is available here](https://github.com/alphagov/prototyping/blob/master/_includes/stylesheets/content-and-design/design-and-development-resources/_buttons.scss).

## Default examples
<div class="pattern-example">
  <div class="inner">

  <p>
    <a href="#" class="btn">Primary action</a> 
    <a href="#" class="btn-secondary">Secondary action</a> 
    <a href="#" class="btn-warning">Warning action</a>
  </p>

  </div>
  <pre><code>&lt;p&gt;
  &lt;a href="#" class="btn"&gt;Primary action&lt;/a&gt; 
  &lt;a href="#" class="btn-secondary"&gt;Secondary action&lt;/a&gt; 
  &lt;a href="#" class="btn-warning"&gt;Warning action&lt;/a&gt;
&lt;/p&gt;
</code></pre>
</div>

Primary actions move users onwards through a transaction. Typically there would only be one of these per screen.

Secondary actions modify the current transaction screen in some way, but don't move users on.

Warning actions are those that are irreversable ('Delete account' for example). Use these sparingly and consider ways of making actions reversable first.

## Launching transactions

The arrow symbol is used for action buttons that initiate a transaction. If the action takes users to a non GOV.UK website they must be informed of this.

<div class="pattern-example">
  <div class="inner">
      <p>
        <a href="#" class="btn" title="Get started on the HMRC website">Get started ❯</a> 
         on the HMRC website
      </p>
    </div>
    <pre><code>  &lt;p&gt;
    &lt;a href="#" class="btn" title="Get started on the HMRC website"&gt;Get started ❯&lt;/a&gt; 
    on the HMRC website
  &lt;/p&gt;
</code></pre>
</div>

## Disabled buttons
<div class="pattern-example">
  <div class="inner">
    <p>
      <a class="disabled btn">Primary action</a> 
      <a class="disabled btn-secondary">Secondary action</a> 
      <a class="disabled btn-warning">Warning action</a>
    </p>
  </div>
  <pre><code>&lt;p&gt;
  &lt;a class="disabled btn"&gt;Primary action&lt;/a&gt; 
  &lt;a class="disabled btn-secondary"&gt;Secondary action&lt;/a&gt; 
  &lt;a class="disabled btn-warning"&gt;Warning action&lt;/a&gt;
&lt;/p&gt;
</code></pre>
</div>


## Other sizes

Use the default size wherever possible. More complex interfaces may occasionally require smaller buttons however.

<div class="pattern-example">
  <div class="inner">

  <p>
    <a href="#" class="small btn">Primary action</a> 
    <a href="#" class="small btn-secondary">Secondary action</a> 
    <a href="#" class="small btn-warning">Warning action</a>
  </p>

  <p>
    <a href="#" class="x-small btn">Primary action</a> 
    <a href="#" class="x-small btn-secondary">Secondary action</a> 
    <a href="#" class="x-small btn-warning">Warning action</a>
  </p>

  </div>
  <pre><code>&lt;p&gt;
  &lt;a href="#" class="small btn"&gt;Primary action&lt;/a&gt; 
  &lt;a href="#" class="small btn-secondary"&gt;Secondary action&lt;/a&gt; 
  &lt;a href="#" class="small btn-warning"&gt;Warning action&lt;/a&gt;
&lt;/p&gt;
&lt;p&gt;
  &lt;a href="#" class="x-small btn"&gt;Primary action&lt;/a&gt; 
  &lt;a href="#" class="x-small btn-secondary"&gt;Secondary action&lt;/a&gt; 
  &lt;a href="#" class="x-small btn-warning"&gt;Warning action&lt;/a&gt;
&lt;/p&gt;
</code></pre>
</div>

## Applying button styles to different tags

<div class="pattern-example">
  <div class="inner">
    <p>
      <button class="btn">Button tag</button>
      <a class="btn">Link tag</a>
      <input class="btn" type="submit" value="Input button" /> 
    </p>
  </div>
<pre><code>&lt;p&gt;
  &lt;button class="btn"&gt;Button tag&lt;/button&gt;
  &lt;a class="btn"&gt;Link tag&lt;/a&gt;
  &lt;input class="btn" type="submit" value="Input button" /&gt; 
&lt;/p&gt;
</code></pre>
</div>

## Dependencies

The above classes are included when you import 'forms.scss'. To create your own button colours you'll need to import 'buttons.scss' and use it as follows:

    .btn-class{
      @include button($colour);
    }


## Button text

Don't use:

* technical terms - eg. use ‘delete’, not ‘form reset’
* lots of words - eye tracking shows users are less likely to read long labels

Do use:

* verbs and an active voice
* clear, informative and succinct language
* sentence case


### Common actions

* **Start now** - at the beginning of the service
* **Next** - not forward, go etc
* **Finish** - at the end of a service and you are taking them to a confirmation page 
* **Back** - not 'previous' etc
* **Sign in / sign out** - not log in / log out
* **Create an account** - not 'register'
* **Sign up** - only use this for mailing lists

### Compound actions

Sometimes you want a single button to perform more than one action. For example, 'Save and quit'.

It's worth trying to avoid this situation but if you can't, use common sense. If one of the actions is obvious or not important to know, don't mention it.

For example, if a button saves the current state and moves the user to the next screen, don't use 'Save and next', just use 'Next', because users will assume the former. When in doubt, test with real users.

### Using multiple buttons on one page

* Try to stick to a single primary action per page
* Make sure the button actions are clearly differentiated
* Keep hard-to-undo actions ('Delete' etc.) well away from other actions






