Written by Andrew Hunt and David Thomas. Originally published October 1999. 25th printing, 10th Feb version. 320 Pages. ISBN: 978-0-2016-1622-4. Publisher: Addison-Wesley.

The book invites the reader to progress in whatever order, as reading from front to back is not necessary.

Throughout the book you are prompted to visit other parts to expand on an topic. At the end of each section you are sometimes prompted with challenges that make you reflect on what you just read.

* Programming is a craft. At its simplest, it comes down to getting a computer to do what you want it to do (or what your user wants it to do).
* What Makes a Pragmatic Programmer?
    * Early adopter/fast adapter.
    * Inquisitive.
    * Critical thinker.
    * Realistic.
    * Jack of all trades.
* Care About Your Craft
* “Kaizen” is a Japanese term that captures the conceptof continuously making many small improvements

# CHAPTER 1 - A PRAGMATIC PHILOSOPHY
1. The Cat Ate My Source Code
    * TIP 3: Provide Options, Don’t Make Lame Excuses
2. Software Entropy
    * While software development is immune from almost all physical laws,entropy hits us hard
    * TIP 4: Don’t Live with Broken Windows […] Don’t let entropy win.
    * Even with a fire to put out, don’t be the first to break a window!
3. Stone Soup and Boiled Frogs
    * TIP 5: Be a Catalyst for Change
    * TIP 6: Remember the Big Picture
    * Challenge: Sometimes deception is needed to catalyse change, so don’t use it to do harm
4. Good-Enough Software
    * If trade-offs are possible ask for them: how good is this software meant to be?
    * TIP 7: Make Quality a Requirements Issue
    * Often great software today is better than perfect software tomorrow. (MVP, agile)
    * Challenge: Think about the software that you use. Is it perfect? Would you rather wait for the "perfect" version?
5. Your Knowledge Portfolio
    * An investment in knowledge always pays the best interest. (Benjamin Franklin)
    * Knowledge in software is an expiring asset
    * Knowledge is an asset and can be treated like an investment:
        * Learn regularly
        * Diversify your knowledge
        * Learn some new up-and-coming tech
        * Give up on old tech
    * Goals:
        * Learn one new language every year (seems a lot). Different languages will force you to think differently. E.g.: Rust forces you not to use nulls and think of return types as Option and Result.
        * Read a technical book each quarter
        * Take classes
        * Participate in local groups
        * Experiment with different environments
        * Stay current online and offline
    * It doesn’t matter what use you are going to make of your new knowledge, the process of learning will expand your thinking.
    * TIP 9: Critically Analyze What You Read and Hear.
    * How to ask for knowledge: be specific, frame the question carefully and politely, once your question is written down look one more time for the answer, be patient, give back.
6. Communicate!
    * "A good idea is an orphan without effective communication."
    * "Plan what you want to say. Write an outline"
    * Ask yourself about your audience the WISDOM acrostic (what, interest, sophisticated, detail, own, motivate)
    * Chose a style (short, long, factual, funny, etc)
    * Make it look good
    * if you want people to listen to you listen to them
    * TIP 10: It’s Both What You Say and the Way You Say It
    * Emails are forever

# Chapter 2 - A Pragmatic Approach
7. The Evils of Duplication
    * TIP 11: DRY—Don’t Repeat Yourself
    * "It isn’t a question of whether you’ll remember: it’s a question of when you’ll forget."
    * Duplication can be due to developers feeling there is no other way, without realising it, for laziness or because different developers produced by coincidence the same functionality.
        * Imposed duplication can be worked around by using a pre-processor.
        * Inadverted duplication can be due to unnormalised data.
    * Comments are knowledge too and take part in duplication.
    *  Performances may tempt developers to duplicate, however this is just a cheap solution (e.g. some calculated data can be stored due to the expense of being produced, however care needs to be taken to update it if any of the inputs changes)
    * TIP 12: Make It Easy to Reuse
    * Ease or reusing functionality is a step towards de-duplication.
8. Orthogonality
    * Often overlooked but core concept of other methods and techniques.
    * Refers to the decoupling of things so that moving on one line your position projected on the other doesn’t change.
    * TIP 13: Eliminate Effects Between Unrelated Things
    * Productivity gain
        * "Changes are localized, so development time and testing time are reduced”
        * Promote reuse.
        * It’s easier to combine orthogonal components.
    * Reduce Risk
        * Compartamentalise problems.
        * Robustness
        * Probably better tested as tests are easier to write.
        * 3rd party changes have less effect on your system.
    * Layering is a powerful way to achieve orthogonality.
    * Test for orthogonality are easy: ask yourself how many parts of the software would be affected if you had to change any other part. It should be one.
    * Unit testing and documentation is an easy winner when code is orthogonal.
    * OOP can provide with tools for more orthogonality, but also more chances for less of it.
9. Reversibility
    * "Requirements, users, and hardware change faster than we can get the software developed."
    * TIP 14: There Are No Final Decisions
    * "keep decisions soft and pliable"
10. Tracer Bullets
    * Tracer bullets are used to correct the aim of fire arms in the dark
    * TIP 15: Use Tracer Bullets to Find the Target
    * Like prototyping but keeping the code.
    * Build skeleton of a system
11. Prototypes and Post-it Notes
    * Prototypes don’t need to be on code. Could be on post-its or whiteboard.
    * TIP 16: Prototype to Learn
    * You can prototype anything new, risky, experimental, etc.
    * Prefer high level language, stick with an at least similar language if testing for performances.
    * Architecture wise prototyping is valuable to check for coupling, responsibilities and collaborations of components, interface definition, possible duplications, data access.
    * Make very clear to everybody that it’s disposable code.
12. Domain Languages
    * TIP 17: Program Close to the Problem Domain.
    * Come up with high level languages to help define a problem.
    * Different part of a problem may require their own language.
    * Define a syntax first with a notation such as BNF.
    * Tools such as bison and yacc can convert well defined languages to programming languages.
13. Estimating
    * Practice estimating to be able to get a feel for order go magnitudes.
    * TIP 18: Estimate to Avoid Surprises.
    * The first question when asked for an estimation should be "how accurate?”.
    * The smaller the granularity of the estimations the greater the assumption of degree of accuracy.

    Duration | Quote estimate in
    ---------|------------------
    1–15 days | days
    3–8 weeks | weeks
    8–30 weeks | months
    30+ weeks | think hard before giving an estimate

    * The easiest way to improve on you estimate is to draw from other peoples experience.
    * Build a model for the estimate e.g. variants, building blocks.
    * "Your experience will tell you when to stop refining.”
    * As the estimate becomes more complex you’ll want to hedge your answers.
    * Take notes of you estimate and the actual values. Investigate when there is a substantial difference.
    * TIP 19: Iterate the Schedule with the Code
    * "What to Say When Asked for an Estimate? You say 'I’ll get back to you.’”

# Chapter 3 - The Basic Tools
* Start with basic tools then adopt new ones. Go back to to the basics from time to time to sharpen your skills.
14. The Power of Plain Text
    * TIP 20: Keep Knowledge in Plain Text
    * Cons: slower to computer, takes more space.
    * Pros: easier to remember the meaning of data, flexible, easier to test (e.g.: change data without recompiling).
15. Shell Games
    * GUI are best for simple operation, shell are more powerful.
    * TIP 21: Use the Power of Command Shells.
16. Power Editing
    * TIP 22: Use a Single Editor Well.
    * Make sure that your editor of choice is available on multiple platforms/OS/shells.
17. Source Code Control
    * Source Code Control Systems allows for decentralisation, collaboration, backup, statistics, undoing changes, compartmentalisation,
    * TIP 23: Always Use Source Code Control
    * Use SCCS for everything, not just code. [This may be a bit anachronistic now that a similar effect can be achieved with Google Docs and the like]
    * SCCS allows for automated and repeatable builds.
18. Debugging
    * "Embrace the fact that debugging is just problem solving, and attack itas such.”
    * TIP 24: Fix the Problem, Not the Blame.
    * TIP 25: Don’t Panic.
    * Resist the urge to fix just the symptoms you see. […] Always try to discover the root cause of a problem, not just this particular appearance of it.
    * Bug reporting isn’t an exact science.
    * Accuracy in bug reports is further diminished when they come through a third party
    * Isolate the steps to reproduce the bug
    * In the process of elimination assume you are at fault before a third party. If a third party is at fault you’ll still probably need to isolate the problem as much as possible for a bug report.
    * TIP 26: “select” Isn’t Broken
    * TIP 27: Don’t Assume It—Prove It
    * If it took a long time to fix try to find ways to make it easier for the next time.
    * If the bug came from one person’s misunderstanding see if anyone else had that misunderstanding.
19. Text Manipulation
    * TIP 28: Learn a Text Manipulation Language
20. Code Generators
    * TIP 29: Write Code That Writes Code
    * Passive code generators produce source code that becomes part of the project. They don’t always have to be totally accurate. [e.g.: templating]
    * Active code generators helps to comply with the DRY principle  and are meant to be used continuously to feed changes from the input files to the output ones [e.g.: asp.net transforms]

# Chapter 4 - Pragmatic Paranoia
* TIP 30: You Can’t Write Perfect Software
* Be defensive of other peoples code, and of your own.
21. Design by Contract
    * "Dealing with computer systems is hard. Dealing with people is even harder.”
    * "Bertrand Meyer developed the concept of Design by Contractfor the language Eiffel.”
    * In DBC there are 3 expectations that needs to be met: preconditions, postconditions and class invariants.
    * TIP 31: Design with Contracts
    * “Be strict in what you will accept beforeyou begin, and promise as little as possible in return”
22. Dead Programs Tell No Lies
    * TIP 32: Crash Early
    * "many times, crashing your program is the bestthing you can do”
23. Assertive Programming
    * TIP 33: If It Can’t Happen, Use Assertions to Ensure That It Won’t
    * Most “assert” should be left in the code even in production.
    * Asserts should have no side effects.
24. When to Use Exceptions
    * "One of the problems with exceptions is knowing when to use them. Webelieve that exceptions should rarely be used as part of a program’snormal flow; exceptions should be reserved for unexpected events”
    * TIP 34: Use Exceptions for Exceptional Problems.
25. How to Balance Resources
    * TIP 35: Finish What You Start
    * Don’t couple code with shared resources
    * "Deallocate resources in the opposite order to that in which you allocate them.”
    * If a series of resources are allocated in more than one place try to allocate them in the same order to avoid deadlocks.
    * Is always a good idea to invest in writing code to check the resource balance, like wrappers that tracks resources.

# Chapter 5 - Bend, or Break
26. Decoupling and the Law of Demeter
    * Write ‘shy’ code: don’t reveal yourself to others and don’t interact with too many people.
    * Organise code in cells and limit their interaction like organisations that value privacy
    * TIP 36: Minimize Coupling Between Modules
    * "The Law of Demeter for functions states that any method of anobject should call only methods belonging to: itself, any parameters that were passed in to the method, any objects it created and any directly held component objects”
27. Metaprogramming
    * Get the details out of the code to make it soft and configurable.
    * Metadata is data about data.
    * TIP 37: Configure, Don’t Integrate
    * TIP 38: Put Abstractions in Code, Details in Metadata
    * Writing configuration out of the code allows for decoupling, abstraction, configuring application more easily and changing the configuration without compilation step
    * Decide if it’s worth implementing a mechanism to reload the configuration on the fly. Consider how long-running the application will be and how quick and easy restarting it would be.
28. Temporal Coupling
    * Time is often ignored in software architecture. Time aspects to consider are concurrency and ordering.
    * Workflow needs to be modelled with UML activity diagram or similar.
    * TIP 39: Analyze Workflow to Improve Concurrency.
    * The hungry consumer model consists of workers getting tasks from a queue.
    * TIP 40: Design Using Services.
    * TIP 41: Always Design for Concurrency.
    * Adding concurrency to a system that has not been designed as such is very difficult.
29. It’s Just a View
    * Use events to decouple.
    * Dispatching events from a single routine cause more coupling. Use the publish/subscribe pattern.
    * TIP 42: Separate Views from Models.
    * Use Model-View-Controller (MVC) to separate the model from both the GUI and the controls.
    * MVC typically exists in a GUI context, but the same concept on segregation of concerns can be applied to any kind of programming: Model is the data, the view is a way of interpreting the data, and the controller feeds the data to the view
30. Blackboards
    * A blackboard system is a data storage that allows actors to asynchronously work on shared data.
    * Blackboard systems became popular in artificial intelligence, where problems are large and complex.
    * TIP 43: Use Blackboards to Coordinate Workflow.

# Chapter 6 - While You Are Coding
31. Programming by Coincidence
    * "We should avoid pro-gramming by coincidence—relying on luck and accidental successes—in favor of programming deliberately.”
    * There are several reasons why not to program by coincidence:
        1. May not be really working but only looking like it is.
        2. May be an accidental boundary condition and not work all the times.
        3. Undocumented behaviour subject to change.
        4. Possibly slower.
        5. More error prone.
    * "Assumptions that aren’t based on well-established facts are the bane of all projects.”
    * TIP 44: Don’t Program by Coincidence
    * To program deliberately you should:
        * Be aware of what you are doing.
        * Be aware of what you are going to do.
        * Proceed from a plan.
        * Don’t rely on accidents or assumptions. Assume the worse.
        * Document your assumptions (e.g.: design by contract).
        * Test your assumptions.
        * Focus on the important parts.
        * Don’t let existing code/functionality dictate what you are doing. Refactor if necessary.
32. Algorithm Speed
    * It’s important to estimate speed of algorithms when not linear.
    * The O notation determines the upper bounds of algorithms. Because of this it does incur in some precision loss.
    * The big O notation is also uses to model memory or other resources.
    * You should keep in mind the big O notation and ask yourself what scale are you working on to understand what the impact will be.
    * TIP 45: Estimate the Order of Your Algorithms.
    * If you are unsure on the efficiency of your code then run it a few time with different inputs and plot the results.
    * TIP 46: Test Your Estimates.
    * Sometimes is not all about the efficiency: if the inputs are known small it could be better to have easier to write and debug code.
    * Be wary of premature optimisations.
33. Refactoring
    * Building software is organic, much more like gardening than building constructions, but business people are much more comfortable with the latter metaphor.
    * "Rewriting, reworking, and re-architecting code is collectively known as refactoring.”
    * Don’t hesitate to refactor.
    *  Good reasons for refactoring: duplication, non-orthodoganl design, outdated knowledge, performance.
    * Time pressure is often used as an excuse for not refactoring.
    * TIP 47: Refactor Early, Refactor Often.
    * Martin Fowler tips on refactoring: don’t add functionality while refactoring; Make sure you are covered by tests and run them often; Take short, deliberate steps.
34. Code That’s Easy to Test
    * "The term “Software IC” (Integrated Circuit) seems to have been invented in 1986 by Cox and Novobilski in their Objective-C book Object-Oriented Programming” and it refers to the practice of writing modular software that can be tested just like individual chips that are then composed together on a board.
    * We can test against the contract.
    * TIP 48: Design to Test
    * Unit tests provide examples on how to use your code.
    * Run the tests often.
    * Use test harnesses (JUnit, xUnit, etc).
    * Debugging may be necessary in production environment, in this case log files, hot keys sequences or web servers to reveal status messages may be the best option.
    * Format log files correctly to help with parson and for readability.
    * "Testing is more cultural than technical"
    * TIP 49: Test Your Software, or Your Users Will
35. Evil Wizards
    * TIP 50: Don’t Use Wizard Code You Don’t Understand
    * Developers routinely use things that don’t understand, but wizard’s code (boilerplate) is not abstracted away like a library or a processor mechanism, it’s interwoven with the developer code.

# Chapter 7 - Before the Project
* There is a fine line between specifying too much (analysis paralysis) and not enough (not being ready).
36. The Requirements Pit
    * "Requirements rarely lie on the surface.Normally, they’re buried deep beneath layers of assumptions, miscon-ceptions, and politics."
    * TIP 51: Don’t Gather Requirements—Dig for Them
    * “A requirement is a statement of somethingthat needs to be accomplished”
    * Document policies should be separated from requirements. E.g. only employee’s supervisors may see employee’s records should be defined as the need for access control in the requirements and specify who can access what(the policies/metadata/configuration) in another document.
    * "It’s important to discover the underlying reason why users do a par-ticular thing, rather than just the way they currently do it. At the endof the day, your development has to solve their business problem, notjust meet their stated requirements. Documenting the reasons behind requirements will give your team invaluable information when makingdaily implementation decisions.”
    * A simple technique to get the requirements right is to do the work of your future users.
    * TIP 52: Work with a User to Think Like a User
    * Don’t over-specify. Requirements are needs.
    * TIP 53: Abstractions Live Longer than Details
    * TIP 54: Use a Project Glossary
    * Make the glossary easily available (e.g., web resource).
37. Solving Impossible Puzzles
    * TIP 55: Don’t Think Outside the Box—Find the Box
    * Whenever faced with a puzzle that seems impossible ask yourself:
        * "Is there an easier way?
        * "Are you trying to solve the right problem, or have you been distracted by a peripheral technicality?"
        * "Why is this thing a problem?"
        * "What is it that’s making it so hard to solve?Does it have to be done this way?
        * "Does it have to be done at all?"
38. Not Until You’re Ready
    * TIP 56: Listen to Nagging Doubts—Start When You’re Ready
    * Listen to your instincts
    * If insure if waiting for the right moment or procrastinating then start a prototype, then you’ll know to start if you fee like you are wasting your time. If prototyping is creating new questions and you encounter snags on the way then you know you weren’t ready.
39. The Specification Trap
    * "Program specification is the process of taking a requirement and reducing it down to the point where a programmer’s skill can take over."
    * The specification is a record for future developers and an agreements with the user, an implicit contact.
    * Many designers find it difficult to stop specifying. Is naive to assume that the perfect specification can be achieved, and even if the client signs it off you can’t guarantee that it’s is what they actually need and they won’t want to change it.
    * TIP 57: Some Things Are Better Done than Described
    * Having a lose specification allows developers to work their way around certain problems.
    * "spec- ification and implementation are simply different aspects of the same process—an attempt to capture and codify a requirement."
    * Over-specifying incurs in diminishing or even negative returns.
40. Circles and Arrows
    * TIP 58: Don’t Be a Slave to Formal Methods
    * "Never underestimate the cost of adopting new tools and methods."
    * TIP 59: Expensive Tools Do Not Produce Better Designs

# Chapter 8 - Pragmatic Projects
41. Pragmatic Teams
    * "Teams as a whole should not tolerate broken windows” (see 1.2).
    * Code quality shouldn’t be delegated to one or some members of the team, it should be everybody’s responsibility.
    * The boiling frog problem is even more common in teams. The dilution of responsibilities makes it easier for the scope to increase and the time scales to decrease. Maybe appoint a person to be in charge of this.
    * Communication is important for individuals as much as for teams, which should comunicate as one. It help to generate a branding for teams, to give them identity.
    * A team member could be appointed as librarian to help with the duplication of documentation.
    * It’s a mistake to organise teams based on jobs
    * TIP 60: Organize Around Functionality, Not Job Functions
    * "The project needs at least two “heads”—one technical, the other administrative."
    * "A great way to ensure both consistency and accuracy is to automate everything the team does."
42. Ubiquitous Automation
    * "Manual procedures leave consistency up to chance; repeatabilityisn’t guaranteed, especially if aspects of the procedure are open to interpretation by different people."
    * TIP 61: Don’t Use Manual Procedures
    * Compiling the project should be reliable and repeatable.
    * Automation doesn’t only apply to code, but also administrative tasks and documentation.
43. Ruthless Testing
    * "Most developers hate testing."
    * TIP 62: Test Early. Test Often. Test Automatically.
    * TIP 63 Coding Ain’t Done ’Til All the Tests Run
    * Unit testing ensures the functionality of a self-contained unit of code.
    * Integration testings verifies the units of code works well together.
    * Validation and verification tests make sure you are building what the client wants
    * Resource exhaustion, errors and recovery needs to be tested too. There are several resources that you can test the lack of: screen resolution, network bandwidth, disk bandwidth, system clock precision, etc.
    * Performance testing include stress testing and load testing. Specialised hardware or software may be necessary.
    * Usability testing are run by real users in real environments.
    * Code quality can be partially tested by measuring cyclomatic complexity, inheritance fan-in and fan-out, response set (number of functions directly invokedby methods of the class) and class coupling ratios. Some of these can only be used by comparison to other pieces of code and so standard deviation is often used.
    * Regression testing compares the outcome of the current test to previous ones or known values. E.g.: making sure bugs don’t resurface.
    * Test data may be real-world one and be collected by competitors systems, or synthetic one that is generated to obtain a big quantity of it or test boundary conditions.
    * GUI tests can be complicated to maintain if the interface is being worked on. In this situation having decoupled code help in reducing the scope of the GUI tests.
    * TIP 64: Use Saboteurs to Test Your Testing
    * Code coverage can help understand if enough have been tested. However, testing 100% of the code lines doesn’t mean perfect testing.
    * TIP 65: Test State Coverage, Not Code Coverage
    * Code should be tested as soon as any production code exists.
    * Test should be run before checking the code in the source repository.
    * When bugs are found a test should be written to avoid it resurfacing.
    * TIP 66: Find Bugs Once
44. It’s All Writing
    * "Pragmatic Programmers embrace documentation as an integral part ofthe overall development process."
    * TIP 67: Treat English as Just Another Programming Language.
    * TIP 68: Build Documentation In, Don’t Bolt It On.
    * "Code should have comments,but too many comments can be just as bad as too few."
    * "Even worse than meaningless names are misleading names."
    * Don’t add to the comments information that can be determined by tools. Like usages, and revision history.
45. Great Expectations
    * A project is successful if it meets the expectations of the clients.
    * TIP 69: Gently Exceed Your Users’ Expectations
    * Surprise the users, in a good way.
46. Pride and Prejudice
    * TIP 70: Sign Your Work
    * Take pride in your work but don’t become prejudiced of it in favour of your coworkers.

# Quick reference guide

1. Care About Your Craft

Why spend your life developing software unless you care about doing it well?

2. Think! About Your Work

Turn off the autopilot and take control. Constantly critique and appraise your work.

3. Provide Options, Don’t Make Lame Excuses

Instead of excuses, provide options. Don’t say it can’t be done; explain what can be done.

4. Don’t Live with Broken Windows

Fix bad designs, wrong decisions, and poor code when you see them.

5. Be a Catalyst for Change

You can’t force change on people. Instead, show them how the future might be and help them participate in creating it.

6. Remember the Big Picture

Don’t get so engrossed in the details that you forget to check what’s happening around you.

7. Make Quality a Requirements Issue

Involve your users in determining the project’s real quality requirements.

8. Invest Regularly in Your Knowledge Portfolio

Make learning a habit.

9. Critically Analyze What You Read and Hear

Don’t be swayed by vendors, media hype, or dogma. Analyze information in terms of you and your project.

10. It’s Both What You Say and theWay You Say It

There’s no point in having great ideas if you don’t communicate them effectively.

11. DRY—Don’t Repeat Yourself

Every piece of knowledge must have a single, unambiguous, authoritative representation within a system.

12. Make It Easy to Reuse

If it’s easy to reuse, people will. Create an environment that supports reuse.

13. Eliminate Effects Between Unrelated Things

Design components that are self-contained, independent, and have a single, well-defined purpose.

14. There Are No Final Decisions

No decision is cast in stone. Instead, consider each as being written in the sand at the beach, and plan for change.

15. Use Tracer Bullets to Find the Target

Tracer bullets let you home in on your target by trying things and seeing how close they land.

16. Prototype to Learn

Prototyping is a learning experience. Its value lies not in the code you produce, but in the lessons you learn.

17. Program Close to the Problem Domain

Design and code in your user’s language.

18. Estimate to Avoid Surprises

Estimate before you start. You’ll spot potential problems up front. 19. Iterate the Schedule with the Code Use experience you gain as you implement to refine the project time scales.

20. Keep Knowledge in Plain Text

Plain text won’t become obsolete. It helps leverage your work and simplifies debugging and testing

21. Use the Power of Command Shells

Use the shell when graphical user interfaces don’t cut it.

22. Use a Single Editor Well

The editor should be an extension of your hand; make sure your editor is configurable, extensible, and programmable.

23. Always Use Source Code Control

Source code control is a time machine for your work—you can go back.

24. Fix the Problem, Not the Blame

It doesn’t really matter whether the bug is your fault or someone else’s—it is still your problem, and it still needs to be fixed.

25. Don’t Panic When Debugging

Take a deep breath and THINK! about what could be causing the bug.

26. “select” Isn’t Broken

It is rare to find a bug in the OS or the compiler, or even a third-party product or library. The bug is most likely in the application.

27. Don’t Assume It—Prove It

Prove your assumptions in the actual environment— with real data and boundary conditions.

28. Learn a Text Manipulation Language

You spend a large part of each day working with text. Why not have the computer do some of it for you?

29. Write Code That Writes Code

Code generators increase your productivity and help avoid duplication.

30. You Can’t Write Perfect Software

Software can’t be perfect. Protect your code and users from the inevitable errors.

31. Design with Contracts

Use contracts to document and verify that code does no more and no less than it claims to do.

32. Crash Early

A dead program normally does a lot less damage than a crippled one.

33. Use Assertions to Prevent the Impossible

Assertions validate your assumptions. Use them to protect your code from an uncertain world.

34. Use Exceptions for Exceptional Problems

Exceptions can suffer from all the readability and maintainability problems of classic spaghetti code. Reserve exceptions for exceptional things.

35. Finish What You Start

Where possible, the routine or object that allocates a resource should be responsible for deallocating it.

36. Minimize Coupling Between Modules

Avoid coupling by writing “shy” code and applying the Law of Demeter.

37. Configure, Don’t Integrate

Implement technology choices for an application as configuration options, not through integration or engineering.

38. Put Abstractions in Code, Details in Metadata

Program for the general case, and put the specifics outside the compiled code base.

39. Analyze Workflow to Improve Concurrency

Exploit concurrency in your user’s workflow.

40. Design Using Services

Design in terms of services—independent, concurrent objects behind well-defined, consistent interfaces.

41. Always Design for Concurrency

Allow for concurrency, and you’ll design cleaner interfaces with fewer assumptions.

42. Separate Views from Models

Gain flexibility at low cost by designing your application in terms of models and views.

43. Use Blackboards to Coordinate Workflow

Use blackboards to coordinate disparate facts and agents, while maintaining independence and isolation among participants.

44. Don’t Program by Coincidence

Rely only on reliable things. Beware of accidental complexity, and don’t confuse a happy coincidence with a purposeful plan.

45. Estimate the Order of Your Algorithms

Get a feel for how long things are likely to take before you write code.

46. Test Your Estimates

Mathematical analysis of algorithms doesn’t tell you everything. Try timing your code in its target environment.

47. Refactor Early, Refactor Often

Just as you might weed and rearrange a garden rewrite, rework, and re-architect code when it needs it. Fix the root of the problem.

48. Design to Test

Start thinking about testing before you write a line of code.

49. Test Your Software, or Your Users Will

Test ruthlessly. Don’t make your users find bugs for you.

50. Don’t UseWizard Code You Don’t Understand

Wizards can generate reams of code. Make sure you understand all of it before you incorporate it into your project.

51. Don’t Gather Requirements—Dig for Them

Requirements rarely lie on the surface. They’re buried deep beneath layers of assumptions, misconceptions, and politics.

52. Work with a User to Think Like a User

It’s the best way to gain insight into how the system will really be used.

53. Abstractions Live Longer than Details

Invest in the abstraction, not the implementation. Abstractions can survive the barrage of changes from different implementations and new technologies.

54. Use a Project Glossary

Create and maintain a single source of all the specific terms and vocabulary for a project.

55. Don’t Think Outside the Box—Find the Box

When faced with an impossible problem, identify the real constraints. Ask yourself: “Does it have to be done this way? Does it have to be done at all?”

56. Start When You’re Ready

You’ve been building experience all your life. Don’t ignore niggling doubts.

57. Some Things Are Better Done than Described

Don’t fall into the specification spiral—at some point you need to start coding.

58. Don’t Be a Slave to Formal Methods

Don’t blindly adopt any technique without putting it into the context of your development practices and capabilities.

59. Costly Tools Don’t Produce Better Designs

Beware of vendor hype, industry dogma, and the aura of the price tag. Judge tools on their merits.

60. Organize Teams Around Functionality

Don’t separate designers from coders, testers from data modelers. Build teams the way you build code.

61. Don’t Use Manual Procedures

A shell script or batch file will execute the same instructions, in the same order, time after time.

62. Test Early. Test Often. Test Automatically

Tests that run with every build are much more effective than test plans that sit on a shelf.

63. Coding Ain’t Done ’Til All the Tests Run

’Nuff said.

64. Use Saboteurs to Test Your Testing

Introduce bugs on purpose in a separate copy of the source to verify that testing will catch them.

65. Test State Coverage, Not Code Coverage

Identify and test significant program states. Just testing lines of code isn’t enough.

66. Find Bugs Once

Once a human tester finds a bug, it should be the last time a human tester finds that bug. Automatic tests should check for it from then on.

67. English is Just a Programming Language

Write documents as you would write code: honor the DRY principle, use metadata, MVC, automatic generation, and so on.

68. Build Documentation In, Don’t Bolt It On

Documentation created separately from code is less likely to be correct and up to date.

69. Gently Exceed Your Users’ Expectations

Come to understand your users’ expectations, then deliver just that little bit more.

70. Sign Your Work

Craftsmen of an earlier age were proud to sign their work. You should be, too.
