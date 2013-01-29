---
layout: gsdm
title: Product phases
subtitle: How a Digital by Default project is structured
section: guidance
status: draft
css: /assets/stylesheets/phases.css
---

<script type="text/JavaScript">
$(document).ready(function (){
  $('#slider-nav a').click(function(){
    var integer = $(this).attr('rel');
    $('#myslide .wrapper').animate({left:-900*(parseInt(integer)-1)})
    $('#slider-nav a').each(function(){
      $(this).removeClass('active');
      if($(this).hasClass('nav-item-'+integer)){
        $(this).addClass('active')}
      });
      $('#myslide').attr('class', function(i, c) { return c.replace(/\bsection-\d+\b/g, 'section-'+integer); });
      return false;
  });
  // Next section arrow link
  $('.next-section').click(function(){
    var classes = $(this).parent().attr('class');
    var thisSection = classes.replace('slider section-', '');
    var nextSection = parseInt(thisSection) + 1;
    $('.nav-item-'+nextSection).click();
      return false;
  });
  // Previous section arrow link
  $('.prev-section').click(function(){
    var classes = $(this).parent().attr('class');
    var thisSection = classes.replace('slider section-', '');
    var nextSection = parseInt(thisSection) - 1;
    $('.nav-item-'+nextSection).click();
      return false;
  });
});
</script>


<p>Building a great digital service is a complex task, with many risks. To maximise the chances of success it is best to break the work down into discrete phases. Each phase can be though of as an iteration of the idea, increasing the level of detail, complexity and risk throughout. This iterative approach allows the team working on the service to  to start small, fail fast, and deliver value to our users as early as possible.</p>

<h2>Explore the service phases</h2>

  <div id="slider-nav" class="slider-nav">
     <a class="nav-item-1 active" rel="1" href="#phase1">1. Discovery</a>
     <a class="nav-item-2" rel="2" href="#phase2">2. Alpha</a>
     <a class="nav-item-3" rel="3" href="#phase3">3. Beta</a>
     <a class="nav-item-4" rel="4" href="#phase4">4. Live</a>
  </div>
 
  <div id="myslide" class="slider section-1">
    <a class="next-section" href="#" title="Go to next phase"><span>▶</span></a>
    <a class="prev-section" href="#" title="Go to previous phase"><span>◀</span></a>
    <div class="wrapper">

      <section id="phase1">
        <!-- <div class="infographic">
        </div> -->
        <h2>Phase 1: Discovery</h2>

        <div class="team">
          <h3>The team</h3>
          <ul>
            <li>Service manager</li>
            <li>Product manager</li>
            <li>Technical architect</li>
            <li>Business analyst</li>
            <li>Delivery manager</li>
            <li>User researcher</li>
          </ul>
        </div>

        <p>The objective of this phase is to gain an understanding of what the users of the service need, what the business requirements are and what technological or policy related constraints there might be.</p>

        <p>This is achieved through research, workshops and interviews. Typical outputs from this phase are a list of user needs, high level wireframes or prototypes and possibly user personas.</p>

        <p>A small team will be required and the phase should not take longer than a week. At the end of the phase a decision should be made whether to proceeed to the Alpha phase.</p>

        <ul class="outputs">
          <li>
            <div class="output">
            </div>
            User needs
          </li>
          <li>
            <div class="output">
            </div>
              Personas
          </li>
          <li>
            <div class="output">
            </div>
              Wireframes
          </li>
        </ul>

      </section>
      <section id="phase2">
        <div class="infographic">
        </div>
        <h2>Phase 2: Alpha</h2>

        <div class="team">
          <h3>The team</h3>
          <ul>
            <li>Service manager</li>
            <li>Product manager</li>
            <li>Technical architect</li>
            <li>Business analyst</li>
            <li>Delivery manager</li>
            <li>User researcher</li>
            <li>Software developer</li>
          </ul>
        </div>

        <p>The objective of this phase is to build a working prototype that can be used by stakeholders or a closed group of end users. This will enable you to trial technology and get early feedback on the design of the service.</p>

        <p>The team will work iteratively, with an inception followed by development, a showcase and a retrospective.</p>

        <p>You may need to bring more developers and designers into the team, but the whole phase should not last longer than about six to eight weeks. By the end of the phase you should have a clear idea of what's required to build a Beta.</p>

        <ul class="outputs">
          <li>
            <div class="output">
            </div>
            Prototype
          </li>
          <li>
            <div class="output">
            </div>
            Backlog
          </li>
        </ul>

      </section>
      <section id="phase3">
        <div class="infographic">
        </div>
        <h2>Phase 3: Beta</h2>

        <div class="team">
          <h3>The team</h3>
          <ul>
            <li>Service manager</li>
            <li>Product manager</li>
            <li>Technical architect</li>
            <li>Business analyst</li>
            <li>Delivery manager</li>
            <li>User researcher</li>
            <li>Software developer</li>
          </ul>
        </div>

        <p>The objective of this phase is to build a fully working prototype and make it available to all end users.</p>

        <p>This is achieved by delivering the user stories in the backlog created in the Alpha phase. This is the time to resolve any outstanding technical or process-related challenges, get the service accredited and plan to go live.</p>

        <p>You will now know what size team you need to deliver the service. The Beta phase should typically take a couple of months at most.</p>

        <ul class="outputs">
          <li>
            <div class="output">
            </div>
            Prototype
          </li>
          <li>
            <div class="output">
            </div>
            Backlog
          </li>
        </ul>
      </section>
      <section id="phase4">
        <div class="infographic">
        </div>
        <h2 >Phase 4: Live</h2>

        <div class="team">
          <h3>Roles required</h3>
          <ul>
            <li>Service manager</li>
            <li>Technical architect</li>
            <li>Software developer</li>
          </ul>
        </div>

        <p>The objective of this phase is to provide a fuly resilient service to all end users. The service should now meet all security and performance standards.
        </p>

        <p>This is not the end of the process though. The service should now be improved continuously, based on user feedback, analytics and further research.</p>

        <p>The team may need to be reconfigured to reflect the operational requirements of the service.</p>

        <ul class="outputs">
          <li>
            <div class="output">
            </div>
            Live service
          </li>
        </ul>
      </section>
    </div>
  </div>



  

<!--

#Structure of a Digital by Default project

Citizens increasingly web savvy, more demanding.

Industry has evolved techniques for digital delivery in web environments.

Digital by Default approach based on industry best practice.

Building a great digital service is a complex task, with many risks. To maximise the chances of success it is best to break this work down into discrete phases. 

Each phase can be though of as an iteration of the idea, increasing the level of detail, complexity and risk throughout.

Great digital services are ones that have evolved rapidly & in the open, with that evolution directed by evidence based on real user feedback.

This iterative approach allows the team working on the service to  to start small, fail fast, and deliver value to our users as early as possible.

Approach designed to provide rapid feedback, openness.

Progression between phases, delivering value early & taking on complexity & risk as capability matures.


* [Discovery phase](/guides-and-toolkits/phases/discovery.html)
* [Alpha phase](/guides-and-toolkits/phases/alpha.html)
* [Public beta phase](/guides-and-toolkits/phases/beta.html)
* [Live phase](/guides-and-toolkits/phases/live.html)
* [Iteration, improvement and operations](/guides-and-toolkits/phases/operational.html)
* [Decommisioning phase](/guides-and-toolkits/phases/decommisioning.html)

-->