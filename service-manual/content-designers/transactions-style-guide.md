---
layout: detailed-guidance
title: Content style guide for services
subtitle: Writing questions, wording for labels, addressing the user and more
category: user-centered-design
type: guide
audience:
  primary: service-managers, designers, developers, content-designers
  secondary: performance-analysts, user-researchers
status: draft
phases:
  - alpha
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Content designers
    url: /service-manual/content-designers
---

##Style guide for services

Using the government standards for style, logic and content design will help you build services that meet user needs.

Building services from the user’s perspective and then developing services based on [users’ needs](https://www.gov.uk/service-manual/user-centered-design/user-needs.html) is essential to your project’s success.  

This guide is itself in the [alpha phase](https://www.gov.uk/service-manual/phases/alpha.html). The guide’s continued success depends on service managers telling us what works with their users. Please share your ideas and research insights on the [mailing list for service designers](mailto: digital-service-designers@digital.cabinet-office.gov.uk)


## GOV.UK writing style


### Follow the style guide

Follow the [GOV.UK style guide](https://www.gov.uk/design-principles/style-guide) for language and style, eg tone of voice, plain English guidance, and any particular points of style like abbreviations, capitalisation etc.


### Be consistent

Language and labelling must be consistent throughout your service, eg don’t switch from ‘last name’ to ‘family name’.


### Verbs not nouns

Tell users exactly what you want them to do. Don’t invent words to ‘sum up’ the user action. 

Example:<br/>
Tell HMRC about a change to your company car. 

Instead of:<br/>
File an HMRC change request


## Naming your service

Name your service after the action the users perform, eg. 'Register to vote', not 'Electoral Registration Portal'.

The name of your service should be on every page.

Optimise the title of your service so users can easily find it through search engines like Google. 

[Read more on search engine optimisation](https://www.gov.uk/designprinciples/seo).








###Don’t assume the user has previous knowledge
Avoid any jargon and ‘insider’ words. Users won’t know what you’re talking about. When writing copy for your service always ask yourself ‘what does this actually mean?’ and ‘what information do I want the user to give me?’.

When you get close to the subject matter it’s sometimes easy to use ‘insider language’. User research will show you what users understand and where they stumble. 

###Think about the user’s learning curve
If you’re building a service that users log on to regularly, eg every week or month, they will soon know what you’re talking about. 

For these services you can use shorter names for things but you must still be specific. For example, you can use ‘Other taxes’ instead of ‘Other taxes you can register for’. But don’t use ‘Other services’. 

To make the learning curve easy for users, call things in your service what they are. For example, use ‘Other tax you need to register for’ instead of ‘Other services’

You can also:

* explain things that are specific to your service when users register or log on for the first time
* give the option to switch off explanations within the service 



## Logic flow

### Start with the majority of users

Find out how the majority of users will use the service and focus on this path first, then incorporate edge cases where needed. 

### User’s perspective
Approach the logic from a user’s perspective rather than how the information is processed later on. 

User research will give you insight into how the user would use the service. 


### Only ask what you need to know
Before you start building your service, write out a list of the information you need from the user and why you need it. For example, ‘need to know their income so that we can subtract income over £100 a week from benefit paid’. 

This will force you to think through each piece of information you’re asking the user to give you. Remember, just because it’s in your current service, doesn’t mean you need it. 

###Once is enough
Make sure you only ask for information once. This seems obvious but in a long service it’s easy to miss and users can lose confidence in a service that lacks consistency and logic. For example, don’t ask for the user’s birth date and then later in the service for their age.




##Addressing the user

###Addressing the user as ‘you’
Address the user as ‘you’ and use ‘your’ for headings and labelling where appropriate, eg ‘Your account’ or just ‘Account’ instead of ‘My account’

If the user has to give information about other people, be clear whose details you’re asking for, eg ‘Your address’ and ‘Your partner’s address’. Give a ‘Same address as you’ option to populate these fields dynamically. 

###Services with multiple users
Where multiple users have access to the same service, use the name of the person logged in, eg ‘Susan’s account details’. 

If users have roles and permissions, spell out what these are. 

Example:<br/>
Susan’s account details

You’re the accountant for Ladida Ltd.<br/>
You can send tax returns on behalf of Ladida Ltd.<br/>

###Referring to your organisation
At the start of the service make clear which organisation the user is dealing with, eg ‘Apply for funding from Defra’. After that you can refer to your organisation as ‘we’. 

If the circumstances are ambiguous or there is more than one party involved, use the name of your organisation throughout, eg ‘HMRC will get back to you in 7 working days’. 

Test with users if they understand who ‘we’ is. Email your insights to [digital-service-designers@digital.cabinet-office.gov.uk](mailto: digital-service-designers@digital.cabinet-office.gov.uk)

##Writing questions 

The same rules apply to writing questions or steps.  

###Making your questions easy to understand
Keep questions short and straightforward. After you’ve written them take a step back and ask yourself ‘what do I actually want the user to do?’. Check your questions for any jargon. Words you use on your project to describe things usually don’t mean anything for users. 

Ask users in research to describe what they’re doing in your service. Use the words they use. 

Usually you’ll have 3 types of questions in your service:

* simple questions: users will know the answers without help, eg ‘What’s your date of birth?’
* lookup questions: users don’t need help but they have to look up the information first, eg ‘What’s your bank account number?’
* complex questions: users need to think before they can answer them and will probably need help, eg ‘Did you not use your car at any time in the last 30 days? Only answer if it was for 5 days in a row.’

Your questions don’t have to be questions, you can also use statements, eg ‘What is your date of birth?’ or ‘Date of birth’. This depends on the format of your service. However, you have to be consistent and keep reading to a minimum. Don’t use ‘What’s your date of birth?’ followed by ‘Contact details’ on the same page. 

###Telling the user why you’re asking a question
When you’re asking for sensitive or additional information, tell the user why using help text. For example, when you’re asking about their partner’s details. Users often don’t understand what this has to do with to their own application.

Example:<br/>
What’s your partner’s income?

[help text] Your household income is your and your partner’s income. To calculate it correctly we have to ask about your partner’s finances. 

User research will show you which questions users have an issue with. Email your insights to [digital-service-designers@digital.cabinet-office.gov.uk](mailto: digital-service-designers@digital.cabinet-office.gov.uk)

###Number of questions 
Keep the number of questions to a minimum. Remember, users are most likely reluctant to use your service in the first place. Each question adds to their frustration and opens the door for mistakes to be entered. 

Often it’s possible to get the same amount of information with fewer questions. 

For example, instead of asking users 10 questions about their personal circumstances you can group these into a checklist. 

Example: Your circumstances

Use: Which of these apply to you?

Choose all that apply [from this checklist]:

* you have a spouse or partner.
* you’re registered for Self Assessment.
* your household income is £30,000 or more.

The checklist is simpler for users than 3 separate questions:

* are you married or in a civil partnership?
* are you registered for Self Assessment?
* is your household income £30,000 or more?

##Help text and error messages

###Help text
Keep it short and concise and speak to the user in their language. Users are unlikely to read anything longer than 3 lines. Focus on user action. Don’t give background information, eg ‘This used to be called X but in 2008 it was changed to Y’. 

Example:<br/>
Enter the years you’ve paid National Insurance.

[help text] These are the years you’ve worked and paid tax or the years when you’ve received benefits like Jobseeker’s Allowance. 

Instead of:<br/>
National Insurance contributions can consist of Class 1 payments for employed, Class 2 payments for self employed or voluntary Class 3 payments. There may be certain years when you’ve received benefits that cover some of your National Insurance contributions. If this is the case, you must also count those years. 

Only if you really can’t fit help text on the page or when it’s an edge case you should use help links, eg ‘What does this mean?’ or ‘What’s X?’. However, this isn’t a dumping ground for long ‘lazy’ help text. 

Don’t send the user out of the service to read long guidance documents. They won’t.




##Who is eligible and who isn’t
There isn’t a ‘one size fits all’ answer for this. How you present eligibility criteria depends on your service. 

Different service exemplars have tried out these things:

* upfront eligibility checker, eg a short smart answer 
* short list of the main eligibility rules on the start page
* keeping eligibility in mind when ordering your questions: try to cover the main eligibility related questions early on (but avoid turning this into a random bunch of questions)
* a dynamically generated order of questions: users only see questions relating to their specific eligibility group, eg users who aren’t working shouldn’t see any questions about employment
* validate questions within your service so that they kick out  users who aren’t eligible, eg ‘Enter a UK postcode’, validation: ‘You can’t register without a UK address’

Things that don’t work very well:

* long bullet lists of eligibility criteria on the start page
* links to ‘see guidance notes’ within a service that send users to long-winded documents
* when it’s not clear what’s the eligibility checker and what’s the actual application 
* eligibility checkers asking the same questions that will be asked in the service later on as users find this very frustrating

If you’re trying out other ideas, email them to [digital-service-designers@digital.cabinet-office.gov.uk](mailto: digital-service-designers@digital.cabinet-office.gov.uk)





Avoid words like ‘more’ or ‘further’ in headings, eg ‘About you’ followed by ‘More about you’. In this example ‘Personal details’ followed by ‘Contact details’ would be useful.

 
