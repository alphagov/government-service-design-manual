---
layout: detailed-guidance-2
title: People's names
format: 'Design pattern'
category: design-and-development-resources
type: resource
audience:
  primary: content-designers, designers
type: guide
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User centered design
    url: /service-manual/user-centered-design  
  -
    title: Design patterns
    url: /service-manual/user-centered-design/resources/design-patterns.html
---

{:.intro}
Be sensitive to different cultural conventions when asking for people's names.
If you can, use a single name field.


## Single name field

A single text field where the user can enter their full name, including title if they wish.

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/single-name-field.png" alt="An example of a single name field">
</div>


### The good

* Can accomodate any kind of name
* Requires the least effort to use
* Avoids the issue of how to label multiple fields

### The bad

* Difficult if you need to parse out things like last name
* Can't use formal contractions in correspondence (like 'Mr. Smith')
* May be difficult to integrate this apporach with legacy back-end systems

## Multiple name fields

If you're constrained by business or technical requirements you can use multiple fields.

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/multiple-name-fields.png" alt="An example of multiple name fields">
</div>

### The good

Depending on your audience you'll be able to extract the parts of the name and do stuff with them.

### The bad

As soon as you adopt multiple fields you introduce the possibility that:

* a person's name won't fit the format you've chosen
* they enter their names in the wrong order
* they try to enter their full name in the first field



### Labels

If you need to use multiple name fields then the preferred labelling is 'First name', 'Last name'. 
Don't include 'Middle names' unless you absolutely have to and make sure it's optional (you don't need to mark it as optional as users will understand this).


### Titles

Avoid asking for someone's title without good reason. It's extra work for users and you're asking them to potentially reveal their gender and marital status as well, which they may not want to do. There are appropriate ways of addressing people in correspondence without using titles.

If you do implement a title field make sure users can leave it blank and then deal with that option sensibly in any resulting correspondence.


### Internationalisation

Remember, users of GOV.UK services come from many different cultural backgrounds, each with their own conventions regarding personal names.

If you can't use a single name field and your service is used by many users outside the UK you can use this format:

- Given name(s)
- Family name
- Any other names (eg maiden name)


## Further reading

For a more detailed overview of issues relating to personal names, read these articles:

* [http://baymard.com/blog/mobile-form-usability-single-input-fields](http://baymard.com/blog/mobile-form-usability-single-input-fields)
* [http://www.w3.org/International/questions/qa-personal-names](http://www.w3.org/International/questions/qa-personal-names)
* [http://www.kalzumeus.com/2010/06/17/falsehoods-programmers-believe-about-names/](http://www.kalzumeus.com/2010/06/17/falsehoods-programmers-believe-about-names/)


---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Peoples-names-mgFWXkwyPEt)

