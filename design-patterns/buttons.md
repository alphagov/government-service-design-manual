---
layout: design-patterns
title: Buttons
subtitle: Build to the GOV.UK style 
section: guidance
subsection: Design patterns
type: resource
status: draft
css: /assets/stylesheets/design-patterns/buttons.css
---

Here's how to create buttons in the GOV.UK style.

[The Sass mixin for making these is available here](https://github.com/alphagov/prototyping/blob/master/_includes/stylesheets/design-patterns/_buttons.scss).

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


#Button labels
<p>A button is the one place on the page that allows your user to move closer to their goal. Here, you will find out common conventions and style guidance. </p>

<p>If the user doesn’t need to do something &ndash; it should be  a link. Buttons are for actions only.</p>

##Guidance
<p>Don&rsquo;t:</p>
<ul>
<li>use technical terms &ndash;  use ‘delete’, not ‘form reset’
<li>make button text too long &ndash; but do make it specific</li>
</ul>
</p>

<p>Use:</p>
<ul>
<li>verbs and an active voice</li>
<li>sentence case</li>
</ul>


###Button labels
<p> Use these terms for common actions:</P>
<ul>
<li>start now &ndash; at the beginning of the service</li>
<li>next &ndash; not forward, go etc</li>
<li>save &ndash; only tell the user about the action that will lead to a result &ndash; so if the action is to save a page and then move on to the next one, only tell the user &lsquo;next&rsquo;.</li>
<li>finish &ndash; at the end of a service and you are taking them to a confirmation page</li> 
<li>back &ndash; not previous etc</li>
<li>sign in / sign out &ndash; not log in / log out</li>
<li>create an account &ndash; not register</li>
<li>sign up &ndash; only for mailing lists</li>
</ul>


###Multiple buttons
<p>If you have a number of actions on the page &ndash; make sure you are very clear. Don’t have similar actions too close.</p>

<p>For example: Clear and Finish</p>
<p>&lsquo;Clear&rsquo;, as a secondary action, should be a link; &lsquo;Finish&rsquo;, as the primary action, should be a button. </p>


##Code/Templates
If you're giving people code r copy to cuta and paste then here is where it will go.

##Why we do this
<p>Users want the maximum amount of information in the shortest time. In a service, users are already more action-orientated than when expecting to read flat copy. They will want to click a button &ndash; it means they will be one step closer to finishing their task.</p>

<p>Eye tracking shows users are less likely to read long labels. We don’t want them to get the action wrong &ndash; that will lead to frustration so we need to be clear, informative and succinct.</p>

<p>Users look at the left of a box and then look for the label &ndash; or down the labels. Making labels short uses fewer eye muscles. This in turn, makes it easier for the user.</p>

##Further reading
Link within the Manual, or to other posts, that will help people to work on the tool.


# See also

* [Transaction design](/guides/designing-your-service/startandendoftransactionpages.html)
* [Buttons design pattern](/guides/design-patterns/buttons.html)

