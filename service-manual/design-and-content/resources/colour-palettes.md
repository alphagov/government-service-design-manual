---
layout: wide
title: Colour palettes
subtitle: Hex values and Sass variable names for the GOV.UK colours
category: design-and-development-resources
type: resource
audience:
    primary: designers
    secondary: 
status: draft
phases:
  - alpha
  - beta
  - live
page_class: colours
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Design and content
    url: /service-manual/design-and-content
---

This is the standard [GOV.UK](https://www.gov.uk) colour palette. We recommend you use the Sass variables where possible in case the colour values are updated. The variables are defined in 'colours.scss' in the [GOV.UK Front-end Toolkit](https://www.gov.uk/service-manual/design-and-content/resources/sass-repositories.html).

## Semantic colour names

<div class="swatches">
  <div>
    <h3>Text</h3>
    <ul>
      <li class="text-colour"><p>$text-colour ($black)</p></li>
      <li class="secondary-text-colour"><p>$secondary-text-colour ($grey-1)</p></li>
      <li class="link-colour"><p>$link-colour (#2e3191)</p></li>
      <li class="link-visited-colour"><p>$link-visited-colour (#2e3191)</p></li>
      <li class="link-active-colour"><p>$link-active-colour (#2e8aca)</p></li>
      <li class="link-hover-colour"><p>$link-hover-colour (#2e8aca)</p></li>
    </ul>
  </div>
  <div>
    <h3>Backgrounds and borders</h3>
    <ul>
      <li class="border-colour"><p>$border-colour ($grey-2)</p></li>
      <li class="panel-colour"><p>$panel-colour ($grey-3)</p></li>
      <li class="canvas-colour"><p>$canvas-colour ($grey-4)</p></li>
      <li class="canvas-colour"><p>$highlight-colour ($grey-4)</p></li>
      <li class="page-colour"><p>$page-colour ($white)</p></li>
    </ul>
  </div>
</div>

## Standard colour palette

<div class="swatches">
  <div>
    <h3>Purple</h3>
    <ul>
      <li class="purple"><p>Sass: $purple</p><p>Hex: #2e358b</p></li>
      <li class="purple-50"><p>Sass: $purple-50</p><p>Hex: #9799c4</p></li>
      <li class="purple-25 legible"><p>Sass: $purple-25</p><p>Hex: #d5d6e7</p></li>
    </ul>
  </div>
  <div>
    <h3>Mauve</h3>
    <ul>
      <li class="mauve"><p>Sass: $mauve</p><p>Hex: #6f72af</p></li>
      <li class="mauve-50"><p>Sass: $mauve-50</p><p>Hex: #b7b9d7</p></li>
      <li class="mauve-25 legible"><p>Sass: $mauve-25</p><p>Hex: #e2e2ef</p></li>
    </ul>
  </div>
  <div>
    <h3>Fuschia</h3>
    <ul>
      <li class="fuschia"><p>Sass: $fuschia</p><p>Hex: #912b88</p></li>
      <li class="fuschia-50"><p>Sass: $fuschia-50</p><p>Hex: #c994c3</p></li>
      <li class="fuschia-25 legible"><p>Sass: $fuschia-25</p><p>Hex: #e9d4e6</p></li>
    </ul>
  </div>
  <div>
    <h3>Pink</h3>
    <ul>
      <li class="pink"><p>Sass: $pink</p><p>Hex: #d53880</p></li>
      <li class="pink-50"><p>Sass: $pink-50</p><p>Hex: #eb9bbe</p></li>
      <li class="pink-25 legible"><p>Sass: $pink-25</p><p>Hex: #f6d7e5</p></li>
    </ul>
  </div>
  <div>
    <h3>Baby Pink</h3>
    <ul>
      <li class="baby-pink"><p>Sass: $baby_pink</p><p>Hex: #f499be</p></li>
      <li class="baby-pink-50 legible"><p>Sass: $baby-pink-50</p><p>Hex: #faccdf</p></li>
      <li class="baby-pink-25 legible"><p>Sass: $baby-pink-25</p><p>Hex: #fdebf2</p></li>
    </ul>
  </div>
  <div>
    <h3>Red</h3>
    <ul>
      <li class="red"><p>Sass: $red</p><p>Hex: #b10e1e</p></li>
      <li class="red-50"><p>Sass: $red-50</p><p>Hex: #d9888c</p></li>
      <li class="red-25 legible"><p>Sass: $red-25</p><p>Hex: #efcfd1</p></li>
    </ul>
  </div>
  <div>
    <h3>Mellow Red</h3>
    <ul>
      <li class="mellow-red"><p>Sass: $mellow-red</p><p>Hex: #df3034</p></li>
      <li class="mellow-red-50"><p>Sass: $mellow-red-50</p><p>Hex: #ef9998</p></li>
      <li class="mellow-red-25 legible"><p>Sass: $mellow-red-25</p><p>Hex: #f9d6d6</p></li>
    </ul>
  </div>
  <div>
    <h3>Orange</h3>
    <ul>
      <li class="orange"><p>Sass: $orange</p><p>Hex: #f47738</p></li>
      <li class="orange-50 legible"><p>Sass: $orange-50</p><p>Hex: #fabb96</p></li>
      <li class="orange-25 legible"><p>Sass: $orange-25</p><p>Hex: #fde4d4</p></li>
    </ul>
  </div>
  <div>
    <h3>Brown</h3>
    <ul>
      <li class="brown"><p>Sass: $brown</p><p>Hex: #b58840</p></li>
      <li class="brown-50"><p>Sass: $brown-50</p><p>Hex: #dac39c</p></li>
      <li class="brown-25 legible"><p>Sass: $brown-25</p><p>Hex: #f0e7d7</p></li>
    </ul>
  </div>
  <div>
    <h3>Yellow</h3>
    <ul>
      <li class="yellow"><p>Sass: $yellow</p><p>Hex: #ffbf47</p></li>
      <li class="yellow-50"><p>Sass: $yellow-50</p><p>Hex: #ffdf94</p></li>
      <li class="yellow-25 legible"><p>Sass: $yellow-25</p><p>Hex: #fff2d3</p></li>
    </ul>
  </div>
  <div>
    <h3>Grass Green</h3>
    <ul>
      <li class="grass-green"><p>Sass: $grass-green</p><p>Hex: #85994b</p></li>
      <li class="grass-green-50"><p>Sass: $grass-green-50</p><p>Hex: #c2cca3</p></li>
      <li class="grass-green-25 legible"><p>Sass: $grass-green-25</p><p>Hex: #e7ebda</p></li>
    </ul>
  </div>
  <div>
    <h3>Green</h3>
    <ul>
      <li class="green"><p>Sass: $green</p><p>Hex: #006435</p></li>
      <li class="green-50"><p>Sass: $green-50</p><p>Hex: #7fb299</p></li>
      <li class="green-25 legible"><p>Sass: $green-25</p><p>Hex: #cce0d6</p></li>
    </ul>
  </div>
  <div>
    <h3>Turquoise</h3>
    <ul>
      <li class="turquoise"><p>Sass: $turquoise</p><p>Hex: #28a197</p></li>
      <li class="turquoise-50"><p>Sass: $turquoise-50</p><p>Hex: #95d0cb</p></li>
      <li class="turquoise-25 legible"><p>Sass: $turquoise-25</p><p>Hex: #d5ecea</p></li>
    </ul>
  </div>
  <div>
    <h3>Light Blue</h3>
    <ul>
      <li class="light-blue"><p>Sass: $light-blue</p><p>Hex: #2b8cc4</p></li>
      <li class="light-blue-50"><p>Sass: $light-blue-50</p><p>Hex: #96c6e2</p></li>
      <li class="light-blue-25 legible"><p>Sass: $light-blue-25</p><p>Hex: #d5e8f3</p></li>
    </ul>
  </div>
</div>

<div class="swatches">
  <div>
    <h2>Greys</h2>
    <ul>
      <li class="black"><p>#0B0C0C, $black</p></li>
      <li class="grey-1"><p>#6F777B, $grey-1</p></li>
      <li class="grey-2"><p>#BFC1C3, $grey-2</p></li>
      <li class="grey-3"><p>#DEE0E2, $grey-3</p></li>
      <li class="grey-4"><p>#F8F8F8, $grey-4</p></li>
      <li class="white"><p>#FFFFFF, $white</p></li>
    </ul>
  </div>
</div>

## Government department colours

<ul class="department-swatches">
  <li class="hm-government"><h4>HM Government</h4><p>Sass: $hm-government</p><p>Hex: #0076c0</p></li>
  <li class="treasury"><h4>Treasury</h4><p>Sass: $treasury</p><p>Hex: #af292e</p></li>
  <li class="cabinet-office"><h4>Cabinet Office</h4><p>Sass: $cabinet-office</p><p>Hex: #0078ba</p></li>
  <li class="department-for-education"><h4>Department for Education</h4><p>Sass: $department-for-education</p><p>Hex: #003a69</p></li>
  <li class="department-for-transport"><h4>Department for Transport</h4><p>Sass: $department-for-transport</p><p>Hex: #006c56</p></li>
  <li class="home-office"><h4>Home Office</h4><p>Sass: $home-office</p><p>Hex: #9325b2</p></li>
  <li class="department-of-health"><h4>Department of Health (NHS)</h4><p>Sass: $department-of-health</p><p>Hex: #00ad93</p></li>
  <li class="ministry-of-justice"><h4>Ministry of Justice</h4><p>Sass: $ministry-of-justice</p><p>Hex: #231f20</p></li>
  <li class="ministry-of-defence"><h4>Ministry of Defence</h4><p>Sass: $ministry-of-defence</p><p>Hex: #4d2942</p></li>
  <li class="foreign-and-commonwealth-office"><h4>Foreign and Commonwealth Office</h4><p>Sass: $foreign-and-commonwealth-office</p><p>Hex: #003e74</p></li>
  <li class="department-for-communities-and-local-government"><h4>Department for Communities and Local Government</h4><p>Sass: $department-for-communities-and-local-government</p><p>Hex: #00857e</p></li>
  <li class="department-of-energy-climate-change"><h4>Department for Energy and Climate Change</h4><p>Sass: $department-of-energy-climate-change</p><p>Hex: #009ddb</p></li>
  <li class="department-for-culture-media-sport"><h4>Department for Culture Media and Sport</h4><p>Sass: $department-for-culture-media-sport</p><p>Hex: #d40072</p></li>
  <li class="department-for-environment-food-and-rural-affairs"><h4>Department for Environment Food and Rural Affairs</h4><p>Sass: $department-for-environment-food-and-rural-affairs</p><p>Hex: #898700</p></li>
  <li class="department-for-work-and-pensions"><h4>Department for Work and Pensions</h4><p>Sass: $department-for-work-and-pensions</p><p>Hex: #00beb7</p></li>
  <li class="department-for-business-innovation-and-skills"><h4>Department for Business, Innovation and Skills</h4><p>Sass: $department-for-business-innovation-and-skills</p><p>Hex: #003479</p></li>
  <li class="department-for-international-development"><h4>Department for International Development</h4><p>Sass: $department-for-international-development</p><p>Hex: #002878</p></li>
  <li class="government-equalities-office"><h4>Government Equalities Office</h4><p>Sass: $government-equalities-office</p><p>Hex: #9325b2</p></li>
  <li class="attorney-generals-office"><h4>Attorney General's Office</h4><p>Sass: $attorney-generals-office</p><p>Hex: #9f1888</p></li>
  <li class="scotland-office"><h4>Scotland Office</h4><p>Sass: $scotland-office</p><p>Hex: #002663</p></li>
  <li class="wales-office"><h4>Wales Office</h4><p>Sass: $wales-office</p><p>Hex: #a33038</p></li>
</ul>

