#Alpha/Beta warnings
For large-scale public alpha or beta releases, we recommend including a warning that alerts users that the service is still in production.

##Code/Templates
[The GOV.UK Beta warning](https://github.com/alphagov/static/blob/master/app/assets/javascripts/welcome.js) served an interstitial when users first arrived on GOV.UK during its public beta.

The message read:

>Welcome to GOV.UK. From 17 October this website will replace Directgov and Business Link as the best place to find government services and information.  

>Until then, you can explore the website by using this experimental trial 'beta' version.

>PLEASE BE AWARE: this is a test website. It may contain inaccuracies or be misleading. <a href='http://www.direct.gov.uk'>Directgov</a> and <a href='http://businesslink.gov.uk'>Business Link</a> remain the official websites for government information and services.

Users were then prompted to accept that they had read the warning, which set a cookie so the message was not displayed repeatedly. 

##Why we do this
Users of government services might not always be used to seeing incomplete or trial version of websites in the open, and it's important to make them aware that they may not be viewing the 'canonical' information they need.

##Further reading
- [The GOV.UK Beta warning](https://github.com/alphagov/static/blob/master/app/assets/javascripts/welcome.js)
- 