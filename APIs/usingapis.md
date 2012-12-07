#Consuming/Using APIs
Don't do everything yourself (you can't). Many times functionality your service needs will be provided by other parts of your organisation, other government departments or by reliable third parties via APIs.

##Guidance/Tool
Most modern digital services are built on top of a wide range of APIs. This allows each part of the service to focus on its core responsibility rather than constantly reinventing the wheel.

### Code Integration
When consuming APIs you should be careful to keep the integration with your code clean and distinct. This is important to ensure that you can swap between providers or update to new versions of an API without making substantial changes to your core code. We encourage the use of adapter code that is entirely focussed on interfacing with the system and mapping code that will provide the linkage between your code's domain model and the concepts and services provided by the API.

### Testing
You should consider carefully how you intend to test your integration with the service. In day to day development you will want to be able to test your code without making (computationally or potentially financially) costly calls out to third party services so you should come up with a way of providing mock versions of those APIs. For full system tests, however, you will want to test the full flow including the third-party service so an automated mechanism should be built for that.
> Many of the GOV.UK publishing applications send emails to provide alerts for content designers. When running tests we don't want to send lots of fake emails so we swap the normal email adapter for one that logs the emails it would have sent. This lets us test our code is doing the right thing without depending on external services.
> Our "data insights" code involves significant interactions with Google Analytics. It wouldn't be practical to test this by sending events to google, waiting for them to be processed, and then reviewing the results. Our developers therefore built a mock service that can be run alongside tests and provides a dummy version of google's api that lets us check the right data is being sent.
> Our publishing systems make use of a single sign-on service. In most of our tests the interaction with that service are mocked so the applications' tests can be run in isolation, but we also have a suite of "smoke tests" that run in our preview environment and use dummy accounts to ensure that the full authentication and authorisation flow is working.
> The Licence Application Tool integrates with a number of third-party payment services. It makes use of test accounts with those services to verify it is able to communicate with them and is sending the right data to complete payments.

### Service Agreements and Resilience
By depending on a third party API you could very easily be tying your service's availability to that of the third party. In some cases that may be acceptable, but often you will want to ensure that you have a fallback plan in place.
The details of that fallback will vary according to your service. It may be that you will need to offer the user the opportunity to use an alternative service, or queue the action to take place later. That could be an automated queue with software that monitors it and retries transactions, or it could be a manual queue where someone follows up to collect further details.
You should be clear with your users about what is happening. If a third party payment provider isn't available you might queue the transaction to try again later. That will mean you can't offer users the same guarantee that their payment will be processed correctly and you should tell them so.

##Code/Templates
If you're giving people code r copy to cuta and paste then here is where it will go.

##Why we do this
This is the reasoning for the decision. You might want to link to relevant blog posts or legislation here.

##Further reading
Link within the Manual, or to other posts, that will help people to work on the tool.