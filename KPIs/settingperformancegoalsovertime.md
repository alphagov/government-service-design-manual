---
layout: gsdm
title: Setting performance goals over time
section: guidance
subsection: KPIs
status: draft
---
    
#Setting performance goals over time
The digital by default standard sets four KPIs to measure the success of your service: completion rate, user satisfaction, digital take-up and cost per transaction. This guide aims to help you set sensible goals for these KPIs over time which you will then use to measure the performance of your service.

##What to aim for

<table>
<tr><th>KPI</th><th>Goal range</th><th>Reason for up or down weighting</th></tr>
<tr><td>Completion rate</td><td>70% - 100%</td><td>
[Online research into task completion rates](http://www.measuringusability.com/blog/task-completion.php) suggests that anything above 78% is above average. However, depending on how you have configured your analytics tool to measure completion rate, you may need to adjust.</td></tr>
<tr><td>User satisfaction</td><td>70% -  100%</td><td>[Consumer Focus research into public service satisfaction](http://www.consumerfocus.org.uk/files/2011/10/Public-sector-service-satisfaction-index.pdf) found that 79% is the mean average public sector score for satisfaction. However, this is likely to be higher for services with pleasant outcome for users (e.g. passports) and lower for services around regulation or enforcement (e.g. tax).</td></tr>
<tr><td>Digital take-up</td><td>70% -	80%</td><td>[Digital Landscape Research](http://publications.cabinetoffice.gov.uk/digital/research/#who-is-online-and-who-is-offline) found that for the UK population as a whole, 82% are online. However, if your service is for younger people you should aim towards the top end of the range, services for older people should aim for the lower end.</td></tr>
<tr><td>Cost per transaction</td><td>Depends</td><td>The [Transactions Explorer](http://transactionalservices.alphagov.co.uk/) contains the digital take-up and cost per transaction figures for high volume transactional services. The goal range for a specific service depends on these values.</td></tr>
</table>

Goals for specific services should be discussed and agreed with GDS during the team inception phase.

##When to achieve by

Completion rate and user satisfaction are _operational_ KPIs and the service should be close to or have achieved its goal prior to launch. Where services are falling short, a decision will need to be taken with GDS as to whether the service can launch. It may be judged acceptable to launch if there is a realistic plan to achieve the agreed performance within six months.

Digital take-up and cost per transaction are _strategic_ KPIs and services should look to achieve the agreed performance within five years of launch. However, the trajectory against which performance is measured will need to be agreed with GDS.

Trajectories will be either low, mid or high as illustrated in the following image:
![performance trajectories](trajectories.png)

##Digital take-up trajectory

Once you have agreed and set your digital take-up goal, which will depend on your users’ demographic profile, you will need to agree the performance trajectory. This will depend on the current level of digital take-up for the service. Use the following table as a guide:
<table>
<tr><th>Current digital take-up</th><th>Trajectory</th></tr>
<tr><td> less than 20%</td><td>Low</td></tr>
<tr><td>20% - 80%</td><td>Mid</td></tr>
<tr><td>> 80%</td><td>High</td></tr>
</table>

These figures reflect the findings from the Digital Efficiency Report, which plotted [digital take-up over time](http://publications.cabinetoffice.gov.uk/digital/efficiency/#fig-10) and found that it followed the typical S-curve of technology adoption.

##Cost per transaction trajectory

The first thing to do is set the goal for cost per transaction (A,,new). This will depend on:
 - the current overall cost per transaction (A)
 - the current digital cost per transaction (B)
 - the current digital take-up (X)
 - the digital take-up goal (Y)

The following formula will provide a sensible goal but assumes a constant digital cost per transaction.

Anew =A+((B-A)/(1-X))*(Y-X)

This should be seen as the starting point for a conversation with GDS to agree a five year cost per transaction goal.

Once you have set your cost per transaction goal, you will need to agree a low, mid or high trajectory. Again, this will depend on your current digital take-up. Use the following as a guide:
<table>
<tr><th>Current digital take-up</th><th>Trajectory</th></tr>
<tr><td> less than 40%</td><td>Low</td></tr>
<tr><td>40% - 80%</td><td>Mid</td></tr>
<tr><td>> 80%</td><td>High</td></tr>
</table>

40% is the point at which cost savings start to be realised more rapidly, as found in the Digital Efficiency Report, hence the switch of trajectories at this point.
