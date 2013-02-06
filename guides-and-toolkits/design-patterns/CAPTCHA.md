[CAPTCHA](http://en.wikipedia.org/wiki/CAPTCHA) stands for "Completely Automated Public Turing test to tell Computers and Humans Apart". These are usually images of jumbled up text that a user needs to decipher and enter before submitting a form.  They are usually used to prevent bots (automated software) from completing a form or accessing a system.

**CAPTCHAs should not be used unless absolutely necessary.**

CAPTCHAs introduce significant problems to online services:
* Usability - they put the burden of detecting bots on the user rather than the system. As CAPTCHAs are designed to be hard to read and understand, this makes the service much more difficult to use
* Accessibility - they are inaccessible by design. This effectively makes the service unusable by people with certain disabilities. Even CAPTCHAs that provide audio versions do not completely resolve this issue.

Additionally, if a 3rd party CAPTCHA service is used, there are further problems to consider:
* Privacy - 3rd party CAPTCHA services set cookies, collect analytics and can track users across multiple sites. This introduces significant privacy concerns.
* Performance - use of a 3rd party CAPTCHA service ties your performance to theirs. If their service goes offline, so does access to your service.
* Security - the security of your service is tied to that of the 3rd party. If they are compromised, so is your service and it's users.

Many of the risks that CAPTCHAs are aimed to mitigate can be addressed in other ways. Rate and connection limiting, the use of honey pots and protective monitoring should all be considered.

It's important to note that even with a CAPTCHA in place bots will still get through due to advances in computer imaging and the use of CAPTCHA farms. A combination of different approaches generally gives the best results.

Further reading:
* http://coding.smashingmagazine.com/2011/03/04/in-search-of-the-perfect-captcha/
* http://www.bbc.co.uk/blogs/bbcinternet/2010/10/captcha_and_bbc_id.html
* http://thenextweb.com/insider/2013/01/30/good-news-music-fans-ticketmaster-is-ditching-its-captcha-conundrums-for-something-simpler/