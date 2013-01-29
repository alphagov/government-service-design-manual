---
layout: gsdm
title: Defining KPIs
section: guidance
subsection: KPIs
status: draft
---
    
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
That’s fine. Again, it’s up to the service manager to define what a successful outcome for users is for the specific version of the service being tested and for the chosen method of measurement. More info on implementing completion rate measurement is available.

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
 - Very satisfied
 - Satisfied
 - Neither satisfied or dissatisfied
 - Dissatisfied
 - Very dissatisfied

###What happens if a user exits the transaction mid way through?
A satisfaction survey should be targeted at 1 in N users in order to achieve a sufficiently large user base - a minimum of 400 - and should be made available to users either on completion or failure to complete a transaction (as defined by completion rate KPI).

###What about parts of the service that are completed offline?
Some services may have components which are completed or fulfilled offline (e.g. applying for a passport, lasting power of attorney). User satisfaction should cover the service as a whole and therefore should be measured at the point of completion. If the point of completion is offline - sometimes referred to as _offline fulfilment_ - then the survey should be triggered via an email to the user.

It is still recommended that the individual parts of the service are tested to identify where specific problems are.

##Digital take-up

###Why measure digital take-up?

Digital take-up is the fundamental measure of how digital a service is. It is a long term strategic measure of how well digital by default service transformation is working. You will monitor this on a monthly basis to track that it is on course with the desired performance trajectory. Further guidance on [setting performance goals over time](/handbook/138/) is available.

###Definition

The number of digital transactions in a calendar month divided by the total number of transactions in the same month, expressed as a percentage.

A digital transaction is one where the primary interaction between the user and the service has been through a digital user interface. There may be transactions completed partly offline, including back-end processing, but these should be recorded as digital transactions so that the overall digital take-up of the service can be measured against use of other channels (e.g. telephone, face to face, post).

###What level of digital take-up should I aim for?

The Digital Landscape research found that 82% of adults in the UK are online. It therefore seems sensible to aim for something in that region. However, about [a third](http://publications.cabinetoffice.gov.uk/digital/research/#fig-2) of those online have never accessed government information or services online so this will involve a significant channel shift for most services. 

###What about services for users who are more or less than 82% online?

The extent to which people are online varies by age, gender and socio-economic status. The digital landscape research has a [useful breakdown](http://publications.cabinetoffice.gov.uk/digital/research/#who-is-online-and-who-is-offline) that should act as a guide. For example, a service aimed at 16-24 year olds should aim for 97% digital take-up whereas for a service aimed at 55-64 year olds, 70% would be a more sensible figure. Use the digital landscape research to up or down weight your target level of digital take-up appropriately.

###When should this level of take-up be achieved by?

Digital take-up is a long term strategic measure and you should be looking to achieve the target level within five years of launch. More information on [setting performance goals over time](/handbook/138/) is available in the manual.

###What about assisted digital? 

Some users will never use the digital service but will still have to have access to the same level of service as those using the digital service. This is called _assisted digital_ and departments will need to determine the appropriate mix of channels to support these users. For more info on assisted digital refer to [this section]() of the manual.

##Cost per transaction

###Why measure cost per transaction?

The main reason for measuring cost per transaction is to get an understanding of the relative cost of channels. The Digital Efficiency Report [link] found that the average cost of a digital transaction is:
 * 20 times lower than the cost of a telephone transaction
 * 30 times lower than the cost of postal transaction
 * 50 times lower than a face-to-face transaction

Understanding cost per transaction for your service will help you to accurately forecast savings and build a strong business case for channel shift.

The other reason for measuring cost per transaction is to provide transparency around public services. Cost per transaction is part of the Quarterly Data Summary (QDS) efficiency data which government departments routinely report. It is also a Cabinet Office requirement to publish cost per transaction for high volume transactional services. A complete list of [government transactions](transactionalservices.alphagov.co.uk) is published online.

###Definition
The _overall cost per transaction_ is the total cost divided by the total number of completed transactions. The total cost includes all fixed and variable costs of the transaction through a given channel, including overheads. It does not include start up costs.

The _digital cost per transaction_ is the total cost of providing the digital service divided by the total number of digital transactions. The cost should include non-digital aspects of service provision where these are essential (e.g. providing wet signatures or manual processing costs).

###What about cost shared across channels?
Some transactions will have costs which are common to digital and non-digital channels (e.g. printing costs for passports or driving licences). These should be included in both the overall and digital cost per transaction. 

For the digital cost per transaction, you should apportion these common costs by share of volume (i.e. if digital take-up is 50% then include half of the common costs). This won’t change the relative cost comparison between channels, but it does give a truer reflection of actual costs for transparency purposes.

###What happens if a user exits the transaction mid way through?
Cost per transaction is calculated for completed transactions only. The cost of failed transactions is included in the overall figure and provides an incentive to drive down failure and improve completion rates.

###What exactly should and shouldn’t be included in the cost?
To be comprehensive, the full cost of the transaction should include:
 
*  accommodation, including capital charges for freehold properties
*  fixtures and fittings
*  maintenance
*  utilities
*  office equipment, including IT systems
*  postage, printing, telecommunications
*  total employment costs of those providing the service, including training
*  overheads, eg (shares of) payroll, audit, top management costs, legal services, etc
*  raw materials and stocks
*  research and development
*  depreciation of start up and one-off capital items
*  taxes: VAT, council tax, stamp duty, etc
*  capital charges (if they were not met separately when the service was established)
*  notional or actual insurance premiums
*  fees to sub-contractors
*  distribution costs, including transport
*  advertising
*  bad debts
*  provisions
 
but not
 
*  enforcement costs
*  replacement costs of items notionally insured
*  start up costs (those which can be capitalised in the accounts)

##Further reading
Link within the Manual, or to other posts, that will help people to work on the tool.
