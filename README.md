# FlashHelper
Wikidoc about the app

# Research & Validation

## 1. Overview

FlashHelper is a proposed mobile application that helps students at Ho Chi Minh City University of Technology (HCMUT) connect with other students who can provide assistance with everyday tasks on campus.

The purpose of this research is to understand students' current help-seeking behaviors, identify barriers to finding assistance, evaluate existing alternatives, and determine which product features should be prioritized for the Minimum Viable Product (MVP).

This document combines primary research from a student survey with a preliminary analysis of existing alternatives. Product decisions are provisional and may change as additional evidence becomes available.

### Research Objectives

- Understand the situations in which students need help on campus.
- Identify the limitations of current help-seeking methods.
- Investigate students' concerns about trust, privacy, and safety.
- Analyze at least five existing alternatives.
- Translate research findings into evidence-based product decisions.

## 2. User Research

### 2.1. Research Method

The team conducted an online questionnaire to investigate students' experiences, preferences, and concerns regarding peer-to-peer assistance on campus.

| Item | Description |
|---|---|
| Research method | Online questionnaire |
| Number of responses | 20 |
| Data collection period | October 8–18, 2026 |
| Main target group | HCMUT students |
| Research focus | Help-seeking behavior, current alternatives, trust, privacy, and product preferences |

The questionnaire included questions about previous experiences, current ways of finding assistance, difficulties encountered, willingness to help other students, and expected features of a dedicated platform.

### 2.2. Participant Profile

| Participant characteristic | Responses | Percentage |
|---|---:|---:|
| HCMUT students | 19 | 95% |
| Students from other VNU-HCM universities | 1 | 5% |
| Fourth-year students | 14 | 70% |
| Third-year students | 4 | 20% |
| Second-year students | 2 | 10% |
| **Total** | **20** | **100%** |

The sample primarily represents senior HCMUT students. Therefore, the findings are useful for initial product discovery but should not be generalized to all university students without further research.

### 2.3. Key Findings

#### Finding 1: Students experience difficulties finding assistance

- 13 out of 20 respondents (65%) reported that they occasionally or frequently needed help on campus.
- 10 respondents (50%) reported that finding assistance typically took between 15 minutes and one hour.
- 5 respondents (25%) reported that it took several hours.
- 1 respondent (5%) reported that finding assistance was very difficult or sometimes unsuccessful.

Overall, 16 respondents (80%) reported that finding assistance took at least 15 minutes or was very difficult.

**Implication:** FlashHelper should investigate whether a dedicated, centralized platform can reduce the time and effort required to find an available helper.

#### Finding 2: Students primarily rely on existing personal relationships

The most common first point of contact was close friends, selected by 13 respondents (65%). Another 6 respondents (30%) selected classmates or people taking the same course.

Current channels included:

- Direct messages to friends and acquaintances: 19 respondents (95%).
- Asking people physically nearby: 7 respondents (35%).
- Posting in university or faculty Facebook groups: 5 respondents (25%).
- Class or cohort group chats on Zalo or Discord: 1 respondent (5%).

Respondents could select multiple channels.

**Implication:** FlashHelper should not simply duplicate existing messaging applications. Its potential value lies in helping students discover available helpers beyond their immediate social circles.

#### Finding 3: Uncertainty and social discomfort discourage public requests

When asked what made them hesitant to post public help requests:

- 13 respondents (65%) selected uncertainty about receiving a useful response.
- 8 respondents (40%) selected fear of being judged.
- 8 respondents (40%) selected concerns about personal information exposure.
- 7 respondents (35%) selected concern about disturbing other members of a shared group.

These findings suggest that students may need a more structured and comfortable way to request help.

**Implication:** The product should make requests easy to create, show their status clearly, and provide a controlled way to communicate without requiring users to publish personal details publicly.

#### Finding 4: Trust and privacy are important adoption considerations

When asked about the most important feature of a dedicated HCMUT platform:

| Feature | Responses | Percentage |
|---|---:|---:|
| HCMUT student account verification | 10 | 50% |
| Location-based matching | 5 | 25% |
| Reputation or post-task ratings | 4 | 20% |
| Notifications about nearby requests | 1 | 5% |

Regarding identity display:

- 12 respondents (60%) preferred anonymity combined with verified HCMUT student status.
- 5 respondents (25%) preferred a nickname.
- 3 respondents (15%) preferred their real names.

The main concerns about using a platform to connect with unfamiliar students included personal information exposure, personal safety during face-to-face interactions, fraudulent accounts, and the possibility of others taking advantage of their willingness to help.

**Implication:** Student verification, privacy controls, and a basic reputation mechanism deserve consideration before advanced features.

Verification should establish eligibility to participate without automatically exposing a user's real name, phone number, or other private information.

#### Finding 5: Location sharing should remain under user control

Regarding location sharing:

- 13 respondents (65%) were willing to share their approximate location only when they needed help or wanted to help someone.
- 5 respondents (25%) were willing to share their location.
- 2 respondents (10%) did not want to share their location for privacy reasons.

**Implication:** Location-based matching may provide value, but continuous location tracking should not be assumed to be acceptable.

The initial product should consider optional, task-specific sharing of approximate campus location rather than requiring permanent location access.

#### Finding 6: Willingness to help depends on practical constraints

When asked whether they would help another student:

- 10 respondents (50%) said their decision would depend on their availability and the task.
- 7 respondents (35%) said they were willing to help whenever the task was within their capabilities.
- 3 respondents (15%) said they would only help people they knew.

In addition, 18 respondents (90%) reported that when they could not find assistance, they generally tried to handle the problem themselves.

**Implication:** FlashHelper should make it easy to understand a task's expected duration, difficulty, and location before accepting it. Matching should account for helpers' availability rather than assuming that nearby students are always willing to help.

### 2.4. Research Limitations

The survey provides useful initial signals but has several limitations:

1. **Small sample size:** The survey contains only 20 responses.
2. **Limited demographic coverage:** 95% of respondents are HCMUT students, and 70% are fourth-year students.
3. **Self-reported behavior:** Responses describe participants' recollections and preferences rather than directly observed behavior.
4. **Hypothetical product evaluation:** Perceived usefulness does not establish actual adoption or continued usage.
5. **Limited behavioral validation:** The survey does not measure actual response times, successful matching rates, or repeat usage.
6. **Potential selection bias:** The recruitment process may have attracted students who were already interested in peer assistance.

The results should therefore be treated as preliminary evidence rather than definitive validation of the product.

### 2.5. Further Research

The next research phase should include:

- Interviews with students from different year levels and faculties.
- Interviews with students who are willing to help, not only students seeking assistance.
- Investigation of recent, concrete situations in which students could not find help quickly.
- Prototype testing to observe whether students can create requests and identify suitable helpers.
- A small pilot to measure response time, successful matches, task completion, and repeat usage.

The team should pay particular attention to whether students would use a dedicated application instead of their existing social networks.

## 3. Existing Solutions and Alternatives

FlashHelper operates in a space where students already have several ways to find assistance. These alternatives include service marketplaces, household-service platforms, and general-purpose community or messaging tools.

The following comparison is an initial desk-research assessment. It describes the general purpose of each alternative and identifies questions the team should investigate further. It does not represent a completed hands-on usability test or a verified comparison of every feature.

### 3.1. Taskrabbit

**Website:** https://www.taskrabbit.com/

Taskrabbit connects customers with people offering local services, including household tasks and other practical assistance.

**Strengths**
- Provides a dedicated environment for finding people to perform tasks.
- Focuses on matching customers with service providers.
- Offers a more task-oriented experience than a general-purpose social network.

**Limitations for FlashHelper's target context**
- Its service-marketplace model differs from informal, student-to-student assistance.
- Its suitability for small, spontaneous, campus-specific requests requires further investigation.
- Its operating availability and service model may not match the HCMUT student community.

**Opportunity for FlashHelper:** Focus on campus-specific requests, peer assistance, and low-friction coordination rather than attempting to reproduce a general service marketplace.

### 3.2. Airtasker

**Website:** https://www.airtasker.com/

Airtasker allows users to post tasks and connect with people who may be able to complete them.

**Strengths**
- Uses a task-centered workflow.
- Provides a relevant reference for task descriptions and the process of accepting work.
- Demonstrates an alternative to searching through personal contacts.

**Limitations for FlashHelper's target context**
- Its broader task marketplace is not specifically designed around verified HCMUT students.
- The suitability of its task and payment model for informal campus assistance needs further investigation.

**Opportunity for FlashHelper:** Simplify the process for small campus tasks and prioritize student identity verification and contextual matching.

### 3.3. bTaskee

**Website:** https://www.btaskee.com/

bTaskee provides household services through a dedicated service platform.

**Strengths**
- Offers a structured way to request household assistance.
- Provides a useful reference for service requests and the coordination of in-person work.
- Is relevant to the Vietnamese market.

**Limitations for FlashHelper's target context**
- Its household-service focus differs from the broader range of student needs, such as directions, finding classrooms, borrowing study materials, or collecting items.
- A professional service model may not be suitable for every informal peer-to-peer request.

**Opportunity for FlashHelper:** Focus on lightweight, everyday student needs that may not justify booking a conventional service.

### 3.4. Facebook Groups

**Website:** https://www.facebook.com/groups/

Students can post requests in university, faculty, or student community groups.

**Strengths**
- Students may already participate in relevant communities.
- Posting can reach people beyond immediate personal contacts.
- Existing groups reduce the need to introduce an entirely new communication channel.

**Limitations**
- Requests may compete with unrelated posts.
- Public requests can expose students to unwanted attention or judgment.
- A requester may not know whether anyone is available to help.
- Finding an appropriate response may require manual follow-up.

These limitations are hypotheses to validate through further observation of actual student groups.

**Opportunity for FlashHelper:** Organize requests by task, location, availability, and status, while limiting unnecessary public exposure.

### 3.5. Zalo and Messenger

**Websites:**
- https://zalo.me/
- https://www.messenger.com/

Students use messaging applications to contact friends, classmates, and other acquaintances directly.

**Strengths**
- Familiar communication workflows.
- Direct conversations make it easy to clarify details.
- Existing social relationships can provide an initial basis for trust.

**Limitations**
- Finding an available helper depends heavily on whom the requester already knows.
- Requests may need to be sent to multiple contacts.
- Messages do not inherently create a shared, searchable directory of available helpers for a campus community.

The extent of these limitations depends on how each student currently uses messaging groups and contacts.

**Opportunity for FlashHelper:** Help users discover potential helpers first, then provide a clear and controlled way to communicate and coordinate.

### 3.6. Comparative Summary

| Alternative | Primary strength | Potential gap for FlashHelper's context |
|---|---|---|
| Taskrabbit | Local task and service matching | Student-specific community and informal campus requests |
| Airtasker | Task posting and task-provider discovery | Verified, campus-focused peer assistance |
| bTaskee | Structured household-service requests | Small, diverse, informal student tasks |
| Facebook Groups | Existing community reach | Structured matching, request status, and privacy controls |
| Zalo / Messenger | Familiar direct communication | Discovering available helpers outside existing contacts |

These alternatives should not be interpreted as having no relevant capabilities. The comparison identifies potential opportunities for differentiation, not proven competitive advantages.

### 3.7. Research Implications

Three areas deserve further investigation:

1. **Discovery:** Can FlashHelper help students find an available and suitable helper more efficiently than their current methods?
2. **Trust:** Does verified student participation increase willingness to request or provide assistance?
3. **Convenience and privacy:** Can students coordinate small tasks without publishing unnecessary personal information or sharing their location continuously?

FlashHelper should not attempt to compete with established platforms on the breadth of services. Its proposed differentiation is a focused, student-centered experience for campus assistance.

## 4. Evidence-Based Product Decisions

The following decisions connect the current survey evidence to the proposed product direction. They remain provisional because the survey is small and the product has not yet been validated through a live pilot.

| Product assumption or question | Evidence | Decision | Product implication |
|---|---|---|---|
| Students experience difficulty finding assistance | 16/20 reported needing at least 15 minutes or finding it very difficult to find help | Keep | Continue investigating faster discovery and matching |
| Students need a dedicated student community | 19/20 respondents were HCMUT students, but most already rely on personal contacts | Investigate further | Test whether a dedicated platform provides enough value beyond existing groups |
| Student verification is important | 10/20 selected verification as the most important feature | Keep | Prioritize a feasible student-verification approach |
| Public identity exposure may discourage requests | 12/20 preferred anonymity with student verification; 5/20 preferred a nickname | Keep | Provide controlled identity display and minimize public personal information |
| Location-based matching is useful | 5/20 selected location-based matching as the most important feature; 13/20 preferred conditional location sharing | Revise | Make approximate location sharing optional and task-specific |
| Ratings can improve trust | 4/20 selected post-task reputation ratings as the most important feature | Investigate further | Consider basic feedback after completion; test whether ratings improve trust |
| Students will always be available to help | 10/20 said willingness depends on availability and task requirements | Reject | Do not assume nearby users are available; let helpers choose whether to accept requests |
| Students will adopt a dedicated app | 14/20 rated the proposed idea 4 or 5 out of 5 | Investigate further | Test actual request creation, response, and repeat usage before claiming adoption |
| Advanced functionality is needed from the start | Respondents identified privacy and complexity as barriers | Revise | Keep the MVP small and easy to use |

### 4.1. Current Product Direction

Based on the available evidence, FlashHelper should initially focus on the following value proposition:

**Help HCMUT students find and coordinate peer assistance for everyday campus tasks through a simple, privacy-conscious, student-focused platform.**

The initial product direction should prioritize:

- A clear workflow for creating and browsing help requests.
- Student identity verification or another credible way to establish community membership.
- Basic task information, including location, expected effort, and relevant requirements.
- A way for helpers to indicate interest and for requesters to coordinate with them.
- User-controlled communication and location sharing.
- Basic safeguards against inappropriate behavior and misuse.

These are proposed priorities, not a claim that every capability has already been implemented or validated.

### 4.2. Decisions Deferred Until Further Evidence

The team should avoid committing to the following until additional research supports them:

- Continuous background location tracking.
- A complex rewards or points system.
- Monetary payments as a core requirement for every task.
- Advanced reputation algorithms.
- A large number of independent task categories.
- A guarantee that every request will receive a response within a fixed period.

The MVP should test the central assumption that students can successfully find and coordinate with suitable helpers through a dedicated campus platform.

## 5. MVP Validation Plan

The next stage should test whether the proposed workflow creates measurable value.

| Hypothesis | Validation method | Suggested measure |
|---|---|---|
| Students can create requests without difficulty | Moderated prototype usability test | Task completion and observed difficulties |
| A dedicated platform improves helper discovery | Compare pilot experiences with current methods | Time to first suitable response |
| Students are willing to accept requests from verified peers | Small campus pilot | Acceptance rate |
| Requests can lead to successful assistance | Track completed requests | Completion rate |
| Privacy controls make users comfortable | Usability testing and follow-up interviews | Reported concerns and observed abandonment |
| The platform provides value beyond existing channels | Follow-up interviews and repeat-use observation | Repeat usage and reasons for choosing FlashHelper |

The team should define measurement rules before the pilot. For example, a successful match should require a suitable helper to respond and the requester to confirm that the response is relevant.

A small pilot can provide directional evidence, but it should not be presented as proof of broad market demand.

## 6. Traceability: Evidence to Product Decisions

The research should connect directly to subsequent design and implementation work.

| Evidence | User need | Proposed requirement | Candidate MVP feature |
|---|---|---|---|
| Finding available assistance can take time | Discover suitable helpers efficiently | Requests and potential helpers must be easy to discover | Help-request feed |
| Students rely on personal contacts | Reach people beyond existing relationships | Requests should be visible to eligible community members | Campus request browsing |
| Student verification is a leading feature preference | Establish community trust | The platform should provide a credible verification mechanism | Student verification |
| Many respondents prefer controlled identity display | Request assistance without unnecessary exposure | Personal details should not be publicly exposed by default | Privacy and identity settings |
| Location sharing is usually conditional | Find nearby assistance without continuous tracking | Users should control when location is shared | Optional approximate location |
| Helpers consider time and task difficulty | Understand the commitment before accepting | Requests should describe task details and expected effort | Task details and acceptance flow |

Each proposed feature should be converted into a user story and a development task only when it is included in the agreed MVP scope.