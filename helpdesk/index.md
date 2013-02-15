---
layout: default
title: Acting on user feedback
section: guidance
subsection: Maintaining services
status: draft
---

* Toc
{:toc}
    
#Acting on user feedback
Once you have created structures to deal with user feedback (see the “Helpdesk” article), you will want to figure out how to make the best use of user feedback to improve your service and the user experience.

##Treating contacts as defects##
One very ambitious improvement model that you can use treats each contact from a user as ultimately due to a defect in the service itself, either in your support model for it or in the communications processes you use to handle contacts. While this perhaps overstates the reality for some kinds of contacts--is a user’s email of thanks and appreciation genuinely a defect?--it is an excellent way to focus critical thinking on how you provide service.  

##Stratifying contact data##
Your ability to act quickly and constructively on user feedback will ultimately depend on the degree to which you can stratify contacts for analysis.  

Unless you have a very small number of contacts or a very large staff available for dedicated analysis, you will probably want to have some degree of stratification take place as the contact is handled. Depending on the complexity of your contact types and the sophistication of the systems you use to handle them, this might be almost completely manual (with metadata created for the contact by the staff handling it) or largely automated (with the software you use to aid contact handling adding most of the metadata automatically).

Minimally, you’ll want to stratify contacts by channel (phone, email, chat, social media, surface mail, etc.) and by the target group that can act on the feedback. For example, some feedback may be directly dealt with by your organisation, while other types may need to be sent to other departments or groups for their use.

##Potential dimensions for stratification##
In addition to the very basic channel and target stratifications mentioned above, you will likely want to consider the following categories as potentially applicable to your contacts/feedback:

###Enquiry type###
Is the contact a question, problem, complaint, FOI request, non-actionable rant, etc.?

###Requester details###
You will want to be very thoughtful about data collection of user details, given privacy concerns, but you will almost certainly want to gather minimal detail on the requesters of your contacts.

###Reply type###
Is a reply necessary/expected or not?

###Enquiry status###
Whether a contact is open, pending some other action, solved, etc.

###Enquiry category or categories###
Some internal sectional or functional categories along which you’ll want to be able to stratify--think of separate URLs or sections on a large website as a concrete example.

###Service category or categories###
Aspects of the handling processes you use that you should capture for further analysis--support level, priority, internal resolving group, resolving agent(s), day/date/time received, and day/date/time resolved are all examples you’ll likely want to have available for analysis.

###Root cause category or categories###
The ultimate reason for the contact; for example, page failing to load, Db down, forgotten password, user error, software bug, etc.  As with most of these categories, you will tailor these categories based on the specifics of your product and support model.

##Sending the data to appropriate groups##
Once you have gathered data for contact stratification, you’ll need to decide how to share it.  If all the data can be handled and used directly within your organisation then this step may be easier, however there will likely be different teams that will act on different types of problems, some of them outside of your department.

Depending on the nature of your contact handling systems, you may be able to use automation to do most of this routing, and you should plan for that if possible.

##Improvement projects##
Once the contact-level improvement data has been sent to the right group, you can begin doing analysis for process improvement. You may want to marry your contact level data with cost data to help prioritise which areas to target first, but you may also simply use the number (or percentage) of affected users as a prioritisation mechanism. 

In any case, you will want to have access to staff with experience in process analysis and improvement, ideally with relevant contact centre or technology experience. Directly involving the individuals who handle contacts and those who built the service into improvement projects will yield better and faster results.


#Helpdesk
In order to provide high-quality service, you will very likely want a dedicated group of specialists - or current staff dedicating some of their time - to handle user enquiries and to help direct them to the information they want. We’ll use the term “Helpdesk” to refer to them.

Your planning will be greatly aided by drawing on the traditional contact centre management and planning skill sets; if you do not have access to staff with these skills, you may have to work closely with other groups who do in a consultancy capacity.

##Volume Forecast: What kind and how many contacts will we get?##
You’ll need to have some idea of the volume and type of contacts that you will receive for the service you’re supporting. In many cases, you can use historical data for similar services as a baseline. This element of your planning will also need to take into account the contact channels you intend to support (e.g., email, phone, chat, Twitter, Facebook, surface mail, etc.).  

Ideally, each of these channels will have a separate contact forecast to aid your planning. This portion of your planning will also be informed by or will itself drive decisions about the technology you’ll use to route, handle, and store historical data from these contacts.

##Handling Time: How long does it take to deal with the contacts?##
You’ll also want to have at least some idea of the average handling time (AHT) and variance for each type of contact. If you’re building off of deep historical data, you may be able to model the AHT separately for each kind of contact. Minimally, you should use the best historical data available to calculate an average number of contacts per day that can be handled by a single agent doing a representative mix of contacts.

##Service Level: How quickly must we respond to user contacts?##
You’ll also need to have an idea of the service level that should be maintained by your Helpdesk.  Generally, service level (“SL” or “SLA”) is expressed as x percent of contacts resolved in y time units. As an example, you might plan to answer 80% of your incoming email contacts in 24 hours.  

You’ll need to know the desired SL to plan how many agents you’ll need for your Helpdesk.  Conversely, if your staffing is constrained by other factors, you may use this value as a constant and solve for the best SL that can be achieved at a certain volume level.

##Staffing and Scheduling: How do we need to staff and schedule to achieve this?##
Once you’ve planned for volume, handling time, and SL, you can finalise your demand forecast for contacts and plan to add staff as necessary to handle these contacts. The same historical volume you used to create your contact forecast will likely give you an idea of how the contacts will be received throughout the planning period (likely a week). This, in turn, will give you good direction on how to schedule your Helpdesk staff to achieve the SL you’ve planned for.

##Data Collection and Reporting: How are things going now that we’re live?##
Once your Helpdesk is operational and actually supporting users, you can collect data on all the areas for which you planned. This data can then be used to improve and refine the forecasts you made previously and to create longer-term projections about the performance of your Helpdesk.  

You will want to develop at least minimal reporting to allow you to evaluate how the group itself and the individuals in it are performing. Depending on the size of your Helpdesk and the details of your contact mix, this might be very detailed and sophisticated or very simple, but you’ll want to plan for how to measure these elements from the very beginning. 

##Going Further##
The points above are good starting points to consider as you plan to support new digital products and initiatives. There is a quite extensive body of online information that expands on the points made here, including some free planning tools. However, as with all areas of specialised knowledge, you should ideally access the services of an experienced analyst rather than depending entirely on online sources for guidance.
