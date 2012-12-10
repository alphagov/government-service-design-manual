#Defining KPIs
This guidance defines the four standard KPIs for digital by default services: completion rate, user satisfaction, digital take-up and cost per transaction. For details on how to implement the measurement of these KPIs, please refer to [Implementing completion rate](), [Implementing user satisfaction]() and [Implementing cost per transaction]().
##Completion rate
###Why measure completion rate?
Completion, or conversion, rate measures the extent to which users successfully achieve their goals in relation to a specific service. This metric is continually monitored by service managers, who take steps to improve and optimise the service over time. This is the ultimate measure of success for a service as it reflects the ability for users to easily and effectively get what they want. Completion rates fall when users are having problems getting what they want from, or navigating through, the digital service which can lead to avoidable contact through other channels, or completing transactions through other channels. This in turn leads to low levels of digital take-up and a high cost per transaction.

###Definition
The number of successfully completed tasks divided by the number of attempted tasks in relation to a digital service.

A task is defined as any pre-determined user outcome that the service manager has defined in relation to the transaction. For example, this might be registering to use the service, completing a part of an application process, or completing an entire transactional process end to end where it is possible to do so.

###An idealised transactional process flow
![idealised transaction process flow](transaction-nodel.png)

In this idealised scenario, the completion rate would be the number of transaction end page views divided by the number of transaction start page views excluding _bounces_, expressed as a percentage. A _bounce_ is where a user views the transaction start page but doesn’t click "Get Started". The start page should tell the user what information is needed to complete the transaction (e.g. credit card, reference number) to minimise the chances of clicking “Get Started” without the necessary information to hand.

###What happens if a user saves progress mid way through?
It is up to the service manager to decide whether saving progress midway through a task is considered to be a successful outcome for the user or not. For example, during alpha development, the manager may want to measure the completion rate in a lab-based test  specifying that users should attempt to complete tasks without saving progress. In that case, if a user saved progress it would be considered a failure. Conversely, if during a public beta, the service manager measures completion rate using digital analytics, they may specify that saving progress is an acceptable outcome for users and not count them as failures.

###What happens if a user starts mid way through the process?
It should not be possible to start a transaction mid way through, unless from a previously saved point. This means that pages beyond the start page should be hidden from search engines.

###What happens if there are multiple endpoints to a transaction?
That’s fine. Again, it’s up to the service manager to define what a successful outcome for users is for the specific version of the product being tested and for the chosen method of measurement. More info on implementing completion rate measurement is available.

###What about offline fulfilment?
Some transactional services have parts which are digital and others which are non-digital. For example, when granting a lasting power of attorney users can start and finish the process online, but are required to print, sign and post a copy of the form in the middle of the process. In this situation it makes sense to treat the discrete digital parts of the service as separate tasks and measure completion rates for each. The purpose of measuring completion rate is to optimise and therefore each portion will need to optimised and therefore treated as a different task that users complete. Offline parts of the process can still be measured but this is likely to be done through qualitative feedback (e.g from surveys, diary studies, focus groups).

##User satisfaction

###Why measure user satisfaction?
A good service enables users to complete tasks successfully. A great service is also enjoyable to use. Completion rate measures whether a service is working _effectively_ for users. Satisfaction provides a qualitative measure of how _satisfying_ the experience is. Many government transactions are mandatory and therefore not inherently enjoyable - sometimes referred to as _grudge transactions_ - but we should endeavour to make them as pleasant as possible for users, who may be nervous or stressed when interacting with the government.

User satisfaction also provides a definitive measure of all the elements contributing to the overall user experience such as ease of use, navigation and design.

###Definition
The percentage of people who answered either “very satisfied” or “satisfied” on a five-point scale in response to the question:

Q:  Overall, how satisfied are you with your visit to the online [e.g. car tax] service today?

A: 
* Very satisfied
* Satisfied
* Neither satisfied or dissatisfied
* Dissatisfied
* Very dissatisfied

###What happens if a user exits the transaction mid way through?
A satisfaction survey should be targeted at 1 in N users in order to achieve a sufficiently large user base - a minimum of 400 - and should be made available to users either on completion or failure to complete a transaction (as defined by completion rate KPI).

###What about parts of the service that are completed offline?
Some services may have components which are completed or fulfilled offline (e.g. applying for a passport, lasting power of attorney). User satisfaction should cover the service as a whole and therefore should be measured at the point of completion. If the point of completion is offline - sometimes referred to as _offline fulfilment_ - then the survey should be triggered via an email to the user.

It is still recommended that the individual parts of the service are tested to identify where specific problems are.

##Further reading
Link within the Manual, or to other posts, that will help people to work on the tool.